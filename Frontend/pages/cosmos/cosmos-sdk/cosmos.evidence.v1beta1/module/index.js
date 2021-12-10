// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSubmitEvidence } from "./types/cosmos/evidence/v1beta1/tx";
const types = [
    ["/cosmos.evidence.v1beta1.MsgSubmitEvidence", MsgSubmitEvidence],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "process.env.VUE_APP_API_TENDERMINT" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSubmitEvidence: (data) => ({ typeUrl: "/cosmos.evidence.v1beta1.MsgSubmitEvidence", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "process.env.VUE_APP_API_COSMOS" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
