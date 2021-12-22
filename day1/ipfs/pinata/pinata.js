//imports needed for this function
const axios = require('axios');
const fs = require('fs');
const FormData = require('form-data');
const recursive = require('recursive-fs');
const basePathConverter = require('base-path-converter');

console.log("=== Pinning new file ==")

const pinDirectoryToIPFS = (pinataApiKey, pinataSecretApiKey) => {
    const url = `https://api.pinata.cloud/pinning/pinFileToIPFS`;
    const src = './dir';

    //we gather the files from a local directory in this example, but a valid readStream is all that's needed for each file in the directory.
    recursive.readdirr(src, function (err, dirs, files) {
        let data = new FormData();
        files.forEach((file) => {
            //for each file stream, we need to include the correct relative file path
            data.append(`file`, fs.createReadStream(file), {
                filepath: basePathConverter(src, file)
            });
        });

        // Adding some metadata
        const metadata = JSON.stringify({
            name: 'myDirectory',
            keyvalues: {
                author: 'adlrocha'
            }
        });
        data.append('pinataMetadata', metadata);

        // Interaction with the Pinata API
        return axios
            .post(url, data, {
                maxBodyLength: 'Infinity', //this is needed to prevent axios from erroring out with large directories
                headers: {
                    'Content-Type': `multipart/form-data; boundary=${data._boundary}`,
                    pinata_api_key: pinataApiKey,
                    pinata_secret_api_key: pinataSecretApiKey
                }
            })
            .then(function (response) {
                //handle response here
                console.log(`> Pinned directory with CID: ${response.data.IpfsHash}`)
            })
            .catch(function (error) {
                //handle error here
                console.log(error)
            });
    });
};

// Yikes! Super insecure.
pinDirectoryToIPFS("65ced690e93a75c0e08b", "9a7fa58c1bea344dd1f1cd7678bb663f0688128ba9741ac7f460bb08dd96afad");
