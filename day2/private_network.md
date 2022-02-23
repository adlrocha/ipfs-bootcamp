# Deploy private IPFS network
- Install priv key generator: `go install github.com/Kubuxu/go-ipfs-swarm-key-gen/ipfs-swarm-key-gen@latest`
  - (Add `/usr/local/go/bin/` to your $PATH)
- Generate the private key: `ipfs-swarm-key-gen > ~/.ipfs/swarm.key`
- Remove bootstraps from config
```
ipfs bootstrap rm --all
# or
vim ~/.ipfs/config
--> "Bootstrap": null
```

- Choose one of the nodes as bootstrap
```
ipfs net listen # To see the address
ipfs bootstrap add <multiaddr_of_bootstrap>
```
- Start your daemon

The problem with private networks is that there
is no pinning replication between all the peers in
the network. For pinning replication we can use IPFS Cluster.
Remember that IPFS Cluster uses an IPFS node as a sidecar.
