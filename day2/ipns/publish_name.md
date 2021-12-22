# IPNS

### Publish name with go-ipfs.
- Mutable content with IPNS
```
ipfs add day2/ipns/website.html
ipfs cat QmdZc8QqFUNkQVpM5Zk1cCSomvksECBeffPuoe3yWfqt4Z
ipfs name publish /ipfs/QmUVTKsrYJpaxUT7dr9FpKq6AoKHhEM7eG1ZHGL56haKLG

# See where you ipns name is pointing
ipfs name resolve
```
- Retrieve using IPNS name: https://gateway.ipfs.io/ipns/k2k4r8ll7mvy4alzl45nquz4tkscbnz3x10k0qv49ne309rlyr5it5nm
- _exercise_ Modify content and re-publish
  - Change website
  - Publish name
  - See if it changes

- Generate a new key for an additional website
```
ipfs key gen KeyForSecondWeb
ipfs name publish --key=KeyForSecondWeb /ipfs/Qm..
```

### Publish name programmatically
- `npm install -g ipfs`
```js
var ipfs = require('ipfs`)
// The address of your files.
const addr = '/ipfs/QmbezGequPwcsWo8UL4wDF6a8hYwM1hmbzYv2mnKkEWaUp'

ipfs.name.publish(addr).then(function (res) {
  // You now receive a res which contains two fields:
  //   - name: the name under which the content was published.
  //   - value: the "real" address to which Name points.
  console.log(`https://gateway.ipfs.io/ipns/${res.name}`)
})
```
