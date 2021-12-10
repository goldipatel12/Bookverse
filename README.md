# Bookverse
 Bookverse empowers book authors by providing a unique community of digital readers a Non-fungible token (NFT) marketplace for books (Think of Bookverse as an online NFT book shop.)

## What you can do on Bookverse

Bookverse is world’s very first marketplace where anyone can :

1. Authors can **Mint their Books as NFTs** 
2. Authors can **Sell their Books as NFTs** to the readers by fixed-price
3. Readers can **Buy Books as NFT from different different Authors** directly by paying a fixed-price.
4. Readers,  Authors both can **Transfer the NFT books (on-chain/cross-chain thanks to Cosmos IBC)** directly to a friend or someone who is interested in P2P transfer.

*Bookverse leverages the power of IPFS to store token metadata and the power of Cosmos IBC to be able to transfer ownership of NFT Book on cross-chains (The Cross-chains part is not included in the demo because it is just DEV Done)*



## How we built it

I have used **Cosmos SDK**, which is the world’s most popular framework for building application-specific blockchains. The basic codes, like different modules, messages, types with CRUD operations, Inter Blockchain Communication packets tor transfer ownership of NFT Books on cross-chains, all this code  was scaffolded using **Starport**. Since Starport and the Cosmos SDK modules are written in the Go programming language, hence I have got a grasp of Golang during this Hackathon. It was a very fun experience working on these Technologies and building my own business logic for Bookverse, And for the Frontend part I was working on ReactJs and NextJs for a while hence got a good grasp over it so it was not a big task for me to create a better UI and UX with these Technologies. Besides all these great technologies of Cosmos Ecosystem, I have used IPFS to store metadata of each Book NFT, the specific service that I have used is called nft.storage which is built specifically for storing off-chain NFT data. Data is stored decentralized on IPFS, and is referenced using content-addressed IPFS URIs.
