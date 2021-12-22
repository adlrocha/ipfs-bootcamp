# DAY 2
Learning how IPFS works under the hood

### A few more protocols (~ 60 min)
- (_optional_) Context Exchange
- Protocols to handle mutable content
- (_optional) DRAND

### IPNS (~ 20 min)
- [Publishing mutable content in the IPFS network](./ipns)
- [DNSLink: Using IPNS and DNS for mutable content](https://dnslink.io).
  - See how `docs.ipfs.io` is actually in IPFS (like many other sites): `dig +noall +answer TXT \_dnslink.docs.ipfs.io`
- _optional_ [Mutable Filesystems and UnixFS](https://docs.ipfs.io/concepts/file-systems/#mutable-file-system-mfs)
        - Try `ipfs files ls`

### IPFS Cluster and private networks (~ 60 min)
- [Deploy a private network](./private_network.md)
- [IPFS Cluster](https://cluster.ipfs.io)
- [Architecture overview](https://cluster.ipfs.io/documentation/deployment/architecture/)
- [Hands on: Let's try to create our own IPFS Cluster](./ipfs_cluster.md)

### NFT.Storage / Web3.Storage (~ 60 min)

### Summary: Interacting with IPFS
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

## Slides
- [DRAND](https://docs.google.com/presentation/d/1xDU1a7P_BkMhy-AkgOz0zDGqsGsE7HKx2bfwlzZ5fWc/edit): Randomness beacons
- [Content Routing](https://docs.google.com/presentation/d/15kzc0rEgOmFTKfcY17E6sjxRDGyqGt760wLTonTtomc/edit#slide=id.gca91fcfd49_0_0): How is content found in IPFS
- (optional) [Content Exchange](https://docs.google.com/presentation/d/1VqduQ6bWMV_R9CQCd86vs1Ozw4WnA3bdO-h-wWilf_0/edit#slide=id.gca3c208903_0_0): How is data transferred in the IPFS network.
- [Mutable Content](https://docs.google.com/presentation/d/1M63MpZYBBUpN8gvvWjbuPjaeny3aFBb5Hdzx-mr2yIw/edit#slide=id.gcad439d6ee_0_346): How to represent mutable content in the IPFS network.


