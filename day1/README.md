# DAY 1: The IPFS Stack (I)
Learning how IPFS works under the hood.

### 🤷‍♂️ Introduction: Why IPFS? (10 min)
### 🖇️ Content addressing in IPFS(~ 100 min)
- [Anatomy of a Cid](https://cid.ipfs.io/)

#### Installing IPFS
- [Installing IPFS Desktop](https://docs.ipfs.io/install/ipfs-desktop/)
- [Installing IPFS Companion](https://docs.ipfs.io/install/ipfs-companion/)
- [Installing CLI](./ipfs/ipfs_install.md)

#### Adding files to IPFS
- Adding a simple file over IPFS using IPFS Desktop.
- Finding it in your browser with IPFS companion (or Brave)
  - https://ipfs.io/ipfs/Qmf17K1hK6DnaLXK1BoumJLzEQPXtcuHu6FXZTtQ9H7Vav
- [Creating a simple website that uploads a file to IPFS](./ipfs/ipfs-example)
  - Local node
  - Infura

#### Pinning files
- Updating with Pinata and pinning a file
  - [Pinata docs](https://docs.pinata.cloud/)
    - [Create API Key](https://app.pinata.cloud/keys)
    - Add and pin from UI
    - Pin using API
    - Link metadata to file. 
   - [Pinning directory script using API](./ipfs/pinata)
  - Why pinning a file?
  - How gateways work?
  - Differences between pinning and adding.
  - _(exercise)_ [Unpin content from Pinata](https://docs.pinata.cloud/api-pinning/unpin)

#### Getting files
- Getting a file from IPFS
  - Curl: `curl -X POST "http://127.0.0.1:5001/api/v0/dag/get?QmedbwgDCW1xgEWsBDSCQStfar3e6zY3La7Xpj9ffWWciU`
    - Local node / IPFS Gateway / Pinata / etc.
  - Node IPFS API

#### Introduction to go-ipfs
- Introduction to go-ipfs
  - `ipfs swarm`: Networking cmds
  - `ipfs pubsub`: Use gossipsub to broadcast messages and create pubsub topic.
  - Download a dataset from [IPFS Awesome](http://awesome.ipfs.io.ipns.localhost:8080/datasets/)
    - `ipfs get <cid>
    - `ipfs cat <cid>`
    - `ipfs ls <cid>` for directories: `ipfs ls /ipfs/QmXHMEB9C3Q4jAAqsrDYXj1kbqmhrDqFmkWaaMH73z6mdE`
  - Pin a file or directory.
    - `ipfs pin`
  - Interact with IPFS as if it was the local filesystem.
    - `ipfs files`


### ☕ (20 min)
### 📂 IPLD (~ 50 min)
- Introduction to IPLD
  - [IPLD docs](https://ipld.io/docs/)
- Hands on with IPLD: [Schemas / LinkSystem](./ipld)
- [CARs (Content Addressable aRchives)](https://ipld.io/specs/transport/car/)
- Filecoin sectors are represented as CARs
- Advanced concepts
  - [How IPFS Web Gateways Work](https://ipld.io/docs/synthesis/how-ipfs-web-gateways-work/)
  - [IPLD Encryption](https://ipld.io/docs/synthesis/encryption/)

### Libp2p (~ 50 min)
- Introduction to Libp2p (20 min)
        - Hands on with libp2p. (30 min)
- Break (10 min)


## Slides
- (optional) [Welcome to Web3](https://docs.google.com/presentation/d/1LvyOH1cqRNefbdLfVwg7raHMqbas0zE6UuQLT2MOTPI/edit#slide=id.gc7c539c017_1_0)
- [Content Addressing in IPFS](https://docs.google.com/presentation/d/1Ym2jGkQAnK4NftPYJPsffQKsxZoh5hf9o-PPsAxoAnw/edit#slide=id.gcbc13f3623_1_65): How content is represented in the IPFS network.
- [Content Routing](https://docs.google.com/presentation/d/15kzc0rEgOmFTKfcY17E6sjxRDGyqGt760wLTonTtomc/edit#slide=id.gca91fcfd49_0_0): How is content found in IPFS
- (optional) [Content Exchange](https://docs.google.com/presentation/d/1VqduQ6bWMV_R9CQCd86vs1Ozw4WnA3bdO-h-wWilf_0/edit#slide=id.gca3c208903_0_0): How is data transferred in the IPFS network.
- [Mutable Content](https://docs.google.com/presentation/d/1M63MpZYBBUpN8gvvWjbuPjaeny3aFBb5Hdzx-mr2yIw/edit#slide=id.gcad439d6ee_0_346): How to represent mutable content in the IPFS network.
- (optional) [Libp2p](https://docs.google.com/presentation/d/190-e2PvZ9OPu3oLrT1j2Qf5RmWygV-7txpYrrcnip04/edit#slide=id.gd4931e17b8_0_4)
- (optional) [IPLD](https://docs.google.com/presentation/d/1-ZscY84fI_gncQn6H3IOLnL8Icr06a9aun8dgvKUGtM/edit#slide=id.gd94be6831b_1_315)

