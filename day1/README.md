# DAY 1: The IPFS Stack (I)
Learning how IPFS works under the hood.

### 🤷‍♂️ Introduction: Why IPFS?
- [Welcome to Web3][https://docs.google.com/presentation/d/19rOP7RlMsGuHlT9ozdwpzT9Vo3DNKj4GLHbKRc1_ae8/edit?usp=sharing]

### 🖇️ Content addressing in IPFS(~ 100 min)
- [Content Addressing in IPFS](https://docs.google.com/presentation/d/1Ym2jGkQAnK4NftPYJPsffQKsxZoh5hf9o-PPsAxoAnw): How content is represented in the IPFS network.
- [Anatomy of a Cid](https://cid.ipfs.io/)
  - [Proto.school tutorial](https://proto.school/anatomy-of-a-cid/06)

#### Installing IPFS
- [Installing IPFS Desktop](https://docs.ipfs.io/install/ipfs-desktop/)
  - Visual overview of IPFS' operation.
  - Modules from the Web3 stack used by IPFS.
- [Installing IPFS Companion](https://docs.ipfs.io/install/ipfs-companion/)
- [Installing CLI](./ipfs/ipfs_install.md)
  - [Building from source](https://github.com/ipfs/kubo#download-and-compile-ipfs)

#### Adding files to IPFS
- Adding a simple file over IPFS using IPFS Desktop.
  - `ipfs add` / `ipfs cat`
- Finding it in your browser with IPFS companion (or Brave)
  - https://ipfs.io/ipfs/QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc/readme
  - Try with from your browser with IPFS companion or Brave.
  - Try with your own file through CID.
- [Creating a simple website that uploads a file to IPFS](./ipfs/ipfs-example)
  - Connect to local node
  - Connect to Infura Gateway
  - Upload the website to IPFS: `ipfs add -R build`
- Let's try to connect our nodes!!🙌
  - `ipfs id`: Get your addresses.
  - `ipfs swarm connect /ip4/127.0.0.1/tcp/36181/p2p/QmbRESPk9EyYPdrmmmcMKsyxnioQFbuAUBTqbpxJ2DKMqM`

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
  - Configuring our go-ipfs node to use Pinata as pinning service.
    - Ensure there is another node pinning content.
      - `ipfs pin remote service add pinata https://api.pinata.cloud/psa YOUR_JWT`
      - `ipfs pin remote add --service=pinata --name=war-and-peace.txt bafybeib32tuqzs2wrc52rdt56cz73sqe3qu2deqdudssspnu4gbezmhig4`
      - `ipfs pin remote ls --service=pinata`
      - `ipfs pin remote ls --service=pinata --status=queued,pinning,failed`
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
    - `ipfs get <cid>`
    - `ipfs cat <cid>`
    - `ipfs ls <cid>` for directories: `ipfs ls /ipfs/QmXHMEB9C3Q4jAAqsrDYXj1kbqmhrDqFmkWaaMH73z6mdE`
  - Pin a file or directory.
    - `ipfs pin`
  - Interact with IPFS as if it was the local filesystem.
    - `ipfs files`
  - Interact via [API](https://docs.ipfs.io/reference/http/api/#api-v0-dag-get):
    - Get a file: `curl -X POST "http://127.0.0.1:5001/api/v0/dag/get?arg=QmaJk4KMwvfthjWhQZEe73DhfT277gUKJVH18eAyRQd58b&output-codec=dag-json"`

### Content Routing in IPFS (~ 30 min)
- [Content Routing](https://docs.google.com/presentation/d/15kzc0rEgOmFTKfcY17E6sjxRDGyqGt760wLTonTtomc): How is content found in IPFS

### 📂 IPLD (~ 50 min)
- [Introduction to IPLD](https://docs.google.com/presentation/d/1-ZscY84fI_gncQn6H3IOLnL8Icr06a9aun8dgvKUGtM/)
  - [IPLD docs](https://ipld.io/docs/)
- Hands on with IPLD: [Schemas / LinkSystem](./ipld)
- [CARs (Content Addressable aRchives)](https://ipld.io/specs/transport/car/)
  - Filecoin sectors are represented as CARs
- Advanced concepts
  - [How IPFS Web Gateways Work](https://ipld.io/docs/synthesis/how-ipfs-web-gateways-work/)
  - [IPLD Encryption](https://ipld.io/docs/synthesis/encryption/)

### 🧑‍💻 Libp2p (~ 50 min)
- [Introduction to Libp2p](https://docs.google.com/presentation/d/190-e2PvZ9OPu3oLrT1j2Qf5RmWygV-7txpYrrcnip04/)
- [Hands on with libp2p](./libp2p)
  - _exercise_ Send marshalled data from an IPLD node
- [Run a few go-libp2p examples](https://github.com/libp2p/go-libp2p/tree/master/examples)
- Alternatively, [run them in javascript](https://github.com/libp2p/js-libp2p/tree/master/examples)

## Resources
- [Libp2p boilerplate](https://github.com/adlrocha/libp2p-boilerplate)
- [Libp2p examples](https://github.com/libp2p/go-libp2p/tree/master/examples)

## Slides
- [Content Addressing in IPFS](https://docs.google.com/presentation/d/1Ym2jGkQAnK4NftPYJPsffQKsxZoh5hf9o-PPsAxoAnw/): How content is represented in the IPFS network.
- [Libp2p](https://docs.google.com/presentation/d/190-e2PvZ9OPu3oLrT1j2Qf5RmWygV-7txpYrrcnip04/)
- [IPLD](https://docs.google.com/presentation/d/1-ZscY84fI_gncQn6H3IOLnL8Icr06a9aun8dgvKUGtM/)
- [Content Routing](https://docs.google.com/presentation/d/15kzc0rEgOmFTKfcY17E6sjxRDGyqGt760wLTonTtomc/): How is content found in IPFS
