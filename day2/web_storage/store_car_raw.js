import fs from 'fs'
import { Readable } from 'stream'
import { CarReader, CarWriter } from '@ipld/car'
import * as raw from 'multiformats/codecs/raw'
import { CID } from 'multiformats/cid'
import { sha256 } from 'multiformats/hashes/sha2'

async function getCar() {
        const bytes = new TextEncoder().encode('random meaningless bytes')
        const hash = await sha256.digest(raw.encode(bytes))
        const cid = CID.create(1, raw.code, hash)
        // create the writer and set the header with a single root
        const { writer, out } = await CarWriter.create([cid])
        Readable.from(out).pipe(fs.createWriteStream('example.car'))
        // store a new block, creates a new file entry in the CAR archive
        await writer.put({ cid, bytes })
        await writer.close()
        const inStream = fs.createReadStream('example.car')
        // read and parse the entire stream in one go, this will cache the contents of
        // the car in memory so is not suitable for large files.
        const reader = await CarReader.fromIterable(inStream)
        return reader
}

const car = await getCar()
const cid = await client.putCar(car)
