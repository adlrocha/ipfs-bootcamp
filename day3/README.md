# DAY 3

## 🪙 Introduction to Filecoin for builders (~50min)
- [Intro to Filecoin](https://docs.google.com/presentation/d/1Kjqa5_wZfxQaoKyE3UsrmA-eYWL78bSz5-OXBN18s-4/edit?usp=sharing)
  - [Filecoin Spec](https://spec.filecoin.io/) 
- [Filecoin Implementations](https://docs.filecoin.io/get-started/#filecoin-implementations)
- [Getting started with Filecoin](https://protocollabs.notion.site/Getting-started-with-IPFS-Filecoin-173c73d4d8d64765a42058594bc46bb7)

## 🛰️ Interacting with the Filecoin network
### Lotus
- [Install Lotus](https://lotus.filecoin.io/docs/set-up/install/)
  ```
  git clone https://github.com/filecoin-project/lotus/
  cd lotus
  make clean all
  sudo make install
  ```
  - Create an identity
    ```
    # In your own full-node
    make lotus-keygen
    ./lotus-keygen -t secp256k1
    lotus wallet import --format=json-lotus f16ejkfhnmmcfhwvr4wjrqd6z6wzchruhq2hzd4ci.key

    # With any lotus distribution
    lotus wallet new
    lotus wallet export f1... > my_address.key
    ```
  - [Lotus lite](https://lotus.filecoin.io/docs/set-up/lotus-lite/)
  - `FULLNODE_API_INFO=wss://api.chain.love lotus daemon --lite`
- [Install development network](https://docs.filecoin.io/build/local-devnet/#manual-set-up)
  - Storage onboarding.
  - Proofs
  - (BONUS) [Textile Docker environment](https://github.com/textileio/lotus-devnet)
    - Support for faster sealing and expensive operations.
  - [Basic commands](./filecoin.md)
- [Store and Retrieve data](https://lotus.filecoin.io/docs/tutorials/store-and-retrieve/)
  - [Filecoin Plus](https://plus.fil.org/miners/)
  ```bash
  lotus client import <some_file>
  lotus client deal
  lotus client list-deals --show-failed
  lotus client list-transfers
  ```
  - Should we try to perform a verified deal in mainnet? (it may take long)
- [Retrieve Data](https://lotus.filecoin.io/docs/developers/retrieve-data/)
```
FULLNODE_API_INFO=wss://api.chain.love lotus daemon --lite
# Look verified miners in https://plus.fil.org/miners/
# Look for miner in https://filecoin.tools
# Find a miner with free retrievals.
lotus client retrieve --miner f07709 mAVWg5AIgFw51hfKzfy8nRsKHlMtT8/DPBJhn1f9eFyOSeldlAiE output-file
```
- [Payment channels](https://lotus.filecoin.io/docs/developers/payment-channels/)
  - [Spec](https://spec.filecoin.io/#section-systems.filecoin_token.payment_channels)
  - Let's try our own payment channel
  ```bash
  # Start payment channel
  lotus paych add-funds <from_addr> <to_addr> 10
  # Create voucher from one side
  lotus paych voucher create <channel addr> 2
  # Send it to provider
  lotus paych voucher add <channel addr> <voucher>
  lotus paych voucher create <channel addr> 4
  # Add more funds
  lotus paych add-funds <client addr> <provider addr> 5
  # Settle payment channel
  lotus paych settle <channel addr>
  # Check best voucher
  lotus paych voucher best-spendable <channel addr>
  # Submit on-chain
  lotus paych voucher submit <channel addr> <voucher>
  # Wait 12 hours for final settling and collect pending funds
  lotus paych collect <channel addr>
  ```
- [Importing data from IPFS](https://lotus.filecoin.io/docs/developers/import-data-from-ipfs/)
  - Hot/Cold copy paradigm.
- [Interact with mainnet without a node](https://lotus.filecoin.io/docs/developers/hosted-lotus/)
  - Lotus-lite
  - Infura
    - `curl -X POST -H "Content-Type: application/json" --url <infura_endpoint> --data '{ "id": 0, "jsonrpc": "2.0", "method": "Filecoin.ChainHead", "params": [] }'`
  - api.chain.love 

### CDN
- [Filecoin Station App](https://www.filstation.app/)

### Estuary
- [Estuary](https://estuary.tech/)
  - [Deploy estuary locally](https://github.com/application-research/estuary)
    - `./estuary`
  - [Getting started](https://docs.estuary.tech/tutorial-get-an-api-key)
    - Let's go through the API

## 💸 Wallets and explorers
- [Glif Wallet](https://wallet.glif.io/?network=f)
- [Filfox](https://filfox.info/en)
- [Spacegap](https://spacegap.github.io/#/)
- [Starboard Solutions](https://www.starboard.ventures/solutions)
  - Miner seals sector and submits miner.PreCommitSector
  - Sector ProveCommit is the first time power is proven to the network and hence power is first added upon successful sector ProveCommit.
- [Protocol Revenue Analysis](https://observablehq.com/@starboard/filecoin-protocol-revenue-analysis)

And a few more explorers (down at the time of writing):
- [Filscout](https://filscout.com/)
- [Filscan](https://filscan.io/)

### The Crypto behind Filecoin
- [Proto.School Filecoin Module](https://proto.school/verifying-storage-on-filecoin)
# Slides
- [Intro to Filecoin](https://docs.google.com/presentation/d/1Kjqa5_wZfxQaoKyE3UsrmA-eYWL78bSz5-OXBN18s-4/edit?usp=sharing)
- [Filecoin Implementations](https://docs.filecoin.io/get-started/#filecoin-implementations)

# Resources
- [Filecoin Ecosystem](https://develop.ecosystem.filecoin.io/?filters=enabled) 
- [Glif Wallet](https://wallet.glif.io/?network=f)
- [Filecoin Spec](https://spec.filecoin.io/) 
