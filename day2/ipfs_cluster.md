#IPFS Cluster
### Instal ipfs-cluster services
- Install from source (the easiest and faster one)
```
git clone https://github.com/ipfs/ipfs-cluster.git
cd ipfs-cluster
go install ./cmd/ipfs-cluster-service
go install ./cmd/ipfs-cluster-ctl
go install ./cmd/ipfs-cluster-follow
```
- [Optionally use docker](https://cluster.ipfs.io/download/)
- (_for devops folks_)[Deployment automations](https://cluster.ipfs.io/documentation/deployment/automations/)

### Deploy ipfs-cluster node
- Start the first node.
```
ipfs-cluster-service init --consensus crdt
# Check the secret and trusted_peers
vim ~/.ipfs-cluster/service.json
ipfs-cluster-service  # to see the help
ipfs-cluster-service daemon

# Check addresses of IPFS cluster
ipfs-cluster-ctl id
ipfs-cluster ctl add <file>
```

- Configure template to start new nodes.
  - Set `export CLUSTER_SECRET=secret` for a raw start
- Use existing `service.json` from previous start (remember to have multiaddress in other ports binded)
 - To start in another folder in the same environment: `ipfs-cluster-service -c someFolder init`
 - Replace service.json with template
 - Start to see that the template works before distribution: `ipfs-cluster-service -c someFolder daemon`
 - Check that it joins the cluster: `ipfs-cluster-ctl peers ls`
  
- Configuration template: `service.json`
  - You can copy the `service.json` from the first peer
- Upload configuration to IPFS with [ipns](./ipns)
  - Or share it through libp2p / IPFS: `ipfs-cluster-ctl add follower_service.json --name follower-config`
  - `ipfs-cluster service init <url of config>

- Start with the template configuration:
  - `ipfs-cluster-service init http://127.0.0.1:8080/ipfs/Qm....`
  - `ipfs-cluster-follow myCluster init http://127.0.0.1:8080/ipfs/Qm...`

- Adding bootstrap peers
```
echo "/dns4/cluster1.domain/tcp/9096/ipfs/QmcQ5XvrSQ4DouNkQyQtEoLczbMr6D9bSenGy6WQUCQUBt" >> ~/.ipfs-cluster/peerstore
ipfs-cluster-service daemon
```

# Adding and pinning files
```
ipfs-cluster-ctl add <file>
ipfs-cluster-ctl pin ls Qm.. # Check pin data
ipfs-cluster-ctl status Qm.. # Request status from every peer
```
```
ipfs-cluster-ctl pin add <cid/ipfs-path> --replication 3
# You can also use --replication_factor_max --replication_factor_min
ipfs-cluster-ctl pin rm <cid/ipfs-path>
ipfs-cluster-ctl pin ls         # List of pins
ipfs-cluster-ctl pin status <cid>    # Status of pins
```
- Metrics of the cluster
```
ipfs-cluster-ctl health
```
