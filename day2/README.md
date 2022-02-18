# DAY 2
Learning how IPFS works under the hood

### 🚀 A few more protocols (~ 60 min)
- [Content Routing](https://docs.google.com/presentation/d/15kzc0rEgOmFTKfcY17E6sjxRDGyqGt760wLTonTtomc/edit#slide=id.gca91fcfd49_0_0): How is content found in IPFS
- _(optional)_ [Content Exchange](https://docs.google.com/presentation/d/1VqduQ6bWMV_R9CQCd86vs1Ozw4WnA3bdO-h-wWilf_0/edit#slide=id.gca3c208903_0_0): How is data transferred in the IPFS network.
- _(optional)_ [DRAND](https://docs.google.com/presentation/d/1xDU1a7P_BkMhy-AkgOz0zDGqsGsE7HKx2bfwlzZ5fWc/edit): Randomness beacons
- [Mutable Content](https://docs.google.com/presentation/d/1M63MpZYBBUpN8gvvWjbuPjaeny3aFBb5Hdzx-mr2yIw/edit#slide=id.gcad439d6ee_0_346): How to represent mutable content in the IPFS network.

### 📛 IPNS (~ 20 min)
- Upload static website as a directory (use relative paths for assets).
- [Publishing mutable content in the IPFS network](./ipns)
- [DNSLink: Using IPNS and DNS for mutable content](https://dnslink.io).
  - See how `docs.ipfs.io` is actually in IPFS (like many other sites): `dig +noall +answer TXT \_dnslink.docs.ipfs.io`
- _optional_ [Mutable Filesystems and UnixFS](https://docs.ipfs.io/concepts/file-systems/#mutable-file-system-mfs)
        - Try `ipfs files ls`

### 🖥️ IPFS Cluster and private networks (~ 60 min)
- [Deploy a private network](./private_network.md)
- [IPFS Cluster](https://cluster.ipfs.io)
- [Architecture overview](https://cluster.ipfs.io/documentation/deployment/architecture/)
- [Hands on: Let's try to create our own IPFS Cluster](./ipfs_cluster.md)

### 💿 NFT.Storage / Web3.Storage (~ 100 min)
The simplest way to store data with some linked metadata.
- [NFT Storage docs](https://nft.storage/#docs)
  - [Client lib](https://nftstorage.github.io/nft.storage/client/)
  - [API docs](https://nft.storage/api-docs/)
- [Hands on NFT.Storage](./nft_storage): Choose one!
  - [NFT Storage examples](https://github.com/nftstorage/nft.storage/tree/main/packages/client/examples/node.js): Encoding blobs, CARs, etc.
  - (_exercise_) [NFT School: Build minting service](https://nftschool.dev/tutorial/minting-service/)
  - (_exercice_) [Mint an NFT locally with Minty](https://docs.ipfs.io/how-to/mint-nfts-with-ipfs/#minty)
- [Intro to CARs]
  - [IPFS Car](https://car.ipfs.io/)
  - [Working with CARs](·/cars.md)
- [Web3 storage docs](https://docs.web3.storage/examples/getting-started)
- [Hands on Web3.Storage](./web3_storage)
  - [API docs](https://docs.web3.storage/reference/http-api/#tag/Web3.Storage-HTTP-API)
  - [Web3.storage examples](https://github.com/web3-storage/web3.storage/tree/main/packages/client/examples/node.js)
  - [Example of Drand relay](https://github.com/alanshaw/drand-relay-w3s)
  - (_optional_) [Golang library](https://pkg.go.dev/github.com/web3-storage/go-w3s-client#Client)
    - [Example](https://github.com/web3-storage/go-w3s-client/tree/main/example)

### 📓 Summary: Interacting with IPFS
- [`go-ipfs`](https://github.com/ipfs/go-ipfs/)
- IPFS Desktop
- IPFS Companion
- Browsers (Opera/Brave)
- `ipfs-http-client`
- Pinata / Infura / Gateways
- NFT.Storage / Web3.Storage
- Ecosystem tools (Fleek, Textile)



## Resources
- [File exchange testbed](https://github.com/protocol/beyond-bitswap/tree/develop/probe) 
- [IPFS Car](https://car.ipfs.io/)
- [Chat application in IPFS](https://blog.ipfs.io/2021-06-10-guide-to-ipfs-connectivity-in-browsers/)

## Slides
- [DRAND](https://docs.google.com/presentation/d/1xDU1a7P_BkMhy-AkgOz0zDGqsGsE7HKx2bfwlzZ5fWc/edit): Randomness beacons
- [Content Routing](https://docs.google.com/presentation/d/15kzc0rEgOmFTKfcY17E6sjxRDGyqGt760wLTonTtomc/edit#slide=id.gca91fcfd49_0_0): How is content found in IPFS
- (optional) [Content Exchange](https://docs.google.com/presentation/d/1VqduQ6bWMV_R9CQCd86vs1Ozw4WnA3bdO-h-wWilf_0/edit#slide=id.gca3c208903_0_0): How is data transferred in the IPFS network.
- [Mutable Content](https://docs.google.com/presentation/d/1M63MpZYBBUpN8gvvWjbuPjaeny3aFBb5Hdzx-mr2yIw/edit#slide=id.gcad439d6ee_0_346): How to represent mutable content in the IPFS network.


