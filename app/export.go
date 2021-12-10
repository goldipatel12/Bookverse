package app

import (
	"encoding/json"
	"log"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk//slashing/types"
	"github.com/cosmos/cosmos-sdk//staking"
	stakingtypes "github.com/cosmos/cosmos-sdk//staking/types"
)

// EportAppStateAndValidators eports the state of the application for a genesis
// file.
func (app *App) EportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string,
) (servertypes.EportedApp, error) {

	// as if they could withdraw from the start of the net block
	ct := app.NewContet(true, tmproto.Header{Height: app.LastBlockHeight()})

	// We eport at last height + 1, because that's the height at which
	// Tendermint will start InitChain.
	height := app.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0
		app.prepForZeroHeightGenesis(ct, jailAllowedAddrs)
	}

	genState := app.mm.EportGenesis(ct, app.appCodec)
	appState, err := json.MarshalIndent(genState, "", "  ")
	if err != nil {
		return servertypes.EportedApp{}, err
	}

	validators, err := staking.WriteValidators(ct, app.StakingKeeper)
	if err != nil {
		return servertypes.EportedApp{}, err
	}
	return servertypes.EportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: app.BaseApp.GetConsensusParams(ct),
	}, nil
}

// prepare for fresh start at zero height
// NOTE zero height genesis is a temporary feature which will be deprecated
//      in favour of eport at a block height
func (app *App) prepForZeroHeightGenesis(ct sdk.Contet, jailAllowedAddrs []string) {
	applyAllowedAddrs := false

	// check if there is a allowed address list
	if len(jailAllowedAddrs) > 0 {
		applyAllowedAddrs = true
	}

	allowedAddrsMap := make(map[string]bool)

	for _, addr := range jailAllowedAddrs {
		_, err := sdk.ValAddressFromBech32(addr)
		if err != nil {
			log.Fatal(err)
		}
		allowedAddrsMap[addr] = true
	}

	/* Just to be safe, assert the invariants on current state. */
	app.CrisisKeeper.AssertInvariants(ct)

	/* Handle fee distribution state. */

	// withdraw all validator commission
	app.StakingKeeper.IterateValidators(ct, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
		_, err := app.DistrKeeper.WithdrawValidatorCommission(ct, val.GetOperator())
		if err != nil {
			panic(err)
		}
		return false
	})

	// withdraw all delegator rewards
	dels := app.StakingKeeper.GetAllDelegations(ct)
	for _, delegation := range dels {
		_, err := app.DistrKeeper.WithdrawDelegationRewards(ct, delegation.GetDelegatorAddr(), delegation.GetValidatorAddr())
		if err != nil {
			panic(err)
		}
	}

	// clear validator slash events
	app.DistrKeeper.DeleteAllValidatorSlashEvents(ct)

	// clear validator historical rewards
	app.DistrKeeper.DeleteAllValidatorHistoricalRewards(ct)

	// set contet height to zero
	height := ct.BlockHeight()
	ct = ct.WithBlockHeight(0)

	// reinitialize all validators
	app.StakingKeeper.IterateValidators(ct, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
		// donate any unwithdrawn outstanding reward fraction tokens to the community pool
		scraps := app.DistrKeeper.GetValidatorOutstandingRewardsCoins(ct, val.GetOperator())
		feePool := app.DistrKeeper.GetFeePool(ct)
		feePool.CommunityPool = feePool.CommunityPool.Add(scraps...)
		app.DistrKeeper.SetFeePool(ct, feePool)

		app.DistrKeeper.Hooks().AfterValidatorCreated(ct, val.GetOperator())
		return false
	})

	// reinitialize all delegations
	for _, del := range dels {
		app.DistrKeeper.Hooks().BeforeDelegationCreated(ct, del.GetDelegatorAddr(), del.GetValidatorAddr())
		app.DistrKeeper.Hooks().AfterDelegationModified(ct, del.GetDelegatorAddr(), del.GetValidatorAddr())
	}

	// reset contet height
	ct = ct.WithBlockHeight(height)

	/* Handle staking state. */

	// iterate through redelegations, reset creation height
	app.StakingKeeper.IterateRedelegations(ct, func(_ int64, red stakingtypes.Redelegation) (stop bool) {
		for i := range red.Entries {
			red.Entries[i].CreationHeight = 0
		}
		app.StakingKeeper.SetRedelegation(ct, red)
		return false
	})

	// iterate through unbonding delegations, reset creation height
	app.StakingKeeper.IterateUnbondingDelegations(ct, func(_ int64, ubd stakingtypes.UnbondingDelegation) (stop bool) {
		for i := range ubd.Entries {
			ubd.Entries[i].CreationHeight = 0
		}
		app.StakingKeeper.SetUnbondingDelegation(ct, ubd)
		return false
	})

	// Iterate through validators by power descending, reset bond heights, and
	// update bond intra-t counters.
	store := ct.KVStore(app.keys[stakingtypes.StoreKey])
	iter := sdk.KVStoreReversePrefiIterator(store, stakingtypes.ValidatorsKey)
	counter := int16(0)

	for ; iter.Valid(); iter.Net() {
		addr := sdk.ValAddress(iter.Key()[1:])
		validator, found := app.StakingKeeper.GetValidator(ct, addr)
		if !found {
			panic("epected validator, not found")
		}

		validator.UnbondingHeight = 0
		if applyAllowedAddrs && !allowedAddrsMap[addr.String()] {
			validator.Jailed = true
		}

		app.StakingKeeper.SetValidator(ct, validator)
		counter++
	}

	iter.Close()

	if _, err := app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ct); err != nil {
		panic(err)
	}

	/* Handle slashing state. */

	// reset start height on signing infos
	app.SlashingKeeper.IterateValidatorSigningInfos(
		ct,
		func(addr sdk.ConsAddress, info slashingtypes.ValidatorSigningInfo) (stop bool) {
			info.StartHeight = 0
			app.SlashingKeeper.SetValidatorSigningInfo(ct, addr, info)
			return false
		},
	)
}
