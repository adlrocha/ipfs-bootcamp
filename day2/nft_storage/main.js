
import fs from 'fs'
import { NFTStorage, File } from 'nft.storage'

const endpoint = 'https://api.nft.storage' // the default
const token = 'YOUR_API_KEY'

async function main() {
  const storage = new NFTStorage({ endpoint, token })
  const metadata = await storage.store({
    name: 'nft.storage store test',
    description:
      'Using the nft.storage metadata API to create ERC-1155 compatible metadata.',
        // TODO: Try adding as much metadata as you need to.
    someOtherMetadata: "the metadata value",
    image: new File([await fs.promises.readFile('my_lego_nft.jpg')], 'my_lego_nft.jpg', {
      type: 'image/jpg',
    }) 
  })
  console.log('IPFS URL for the metadata:', metadata.url)
  console.log('metadata.json contents:\n', metadata.data)
  console.log(
    'metadata.json contents with IPFS gateway URLs:\n',
    metadata.embed()
  )
}
main()
