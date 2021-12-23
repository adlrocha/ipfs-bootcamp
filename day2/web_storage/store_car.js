
import { createReadStream } from 'fs'
import { CarReader } from '@ipld/car'

import process from 'process'
import minimist from 'minimist'
import { Web3Storage } from 'web3.storage'

function makeStorageClient(token) {  return new Web3Storage({ token: token })}

async function storeCarFile(filename) {
  const inStream = createReadStream(filename)
  const car = await CarReader.fromIterable(inStream)
  
  const client = makeStorageClient()
  const cid = await client.putCar(car)
  console.log('Stored CAR file! CID:', cid)
}

async function main(){
        // TODO
}

