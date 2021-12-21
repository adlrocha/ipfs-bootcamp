

- [Install cli](https://docs.ipfs.io/install/command-line/#official-distributions)
```sh
wget https://dist.ipfs.io/go-ipfs/v0.10.0/go-ipfs_v0.10.0_linux-amd64.tar.gz
tar -xvzf go-ipfs_v0.10.0_linux-amd64.tar.gz
cd go-ipfs
sudo bash install.sh
```
- See if the installation was successfully
```
ipfs version
```
- Start daemon
```
ipfs daemon
```
NOTE: You may need to `ipfs init` to initialize your node config.
