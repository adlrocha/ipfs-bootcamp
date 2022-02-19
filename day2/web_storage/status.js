
import process from 'process'
import minimist from 'minimist'
import { Web3Storage } from 'web3.storage'

function makeStorageClient(token) {  return new Web3Storage({ token: token })}
async function checkStatus(cid) {
        const client = makeStorageClient()
        const status = await client.status(cid)
        console.log(status)
        if (status) {
                // TODO: your code to do something fun with the status info here
        }
}

async function main(){
        const args = minimist(process.argv.slice(2))
        const token = args.token

        if (!token) {
                console.error('A token is needed. You can create one on https://web3.storage')
                return
        }

        // TODO: replace with your own CID to see info about your uploads!
        checkStatus('YOUR_CID')
}
