const AWS = require('aws-sdk');

// Create API_KEY in app.fleek.
const s3 = new AWS.S3({
        apiVersion: '2006-03-01',
        accessKeyId: '[[apiKey]]',
        secretAccessKey: '[[apiSecret]]',
        endpoint: 'https://storageapi.fleek.co',
        region: 'us-east-1',
        s3ForcePathStyle: true
});

function listBuckets(){
        s3.listBuckets(function (err, data) {
                if (err) {
                        console.log("Error when listing buckets", err);
                } else {
                        console.log("Success when listing buckets", data);
                }
        });
}

func listingFiles(){
        const params = {
                Bucket: "my-bucket",
                MaxKeys: 20
        };

        s3.listObjectsV2(params, function (err, data) {
                if (err) {
                        console.log("Error when listing objects", err);
                } else {
                        console.log("Success when listing objects", data);
                }
        });
}

function uploadFile(){
        const params = {
                Bucket: 'my-team-bucket',
                Key: 'folder/my-picture',
                ContentType: 'image/png',
                Body: myPictureFile,
                ACL: 'public-read',
        };

        const request = s3.putObject(params);
        request.send();
}

function getHash(){
        const request = s3.putObject(params);

        request.on('httpHeaders', (statusCode, headers) => {
                const ipfsHash = headers['x-fleek-ipfs-hash'];
                // Do stuff with ifps hash....
                const ipfsHashV0 = headers['x-fleek-ipfs-hash-v0'];
                // Do stuff with the short v0 ipfs hash... (appropriate for storing on blockchains)
        }).send();
}

// Fetching file from hash
// curl https://ipfs.io/ipfs/Qmaisz6NMhDB51cCvNWa1GMS7LU1pAxdF4Ld6Ft9kZEP2a
