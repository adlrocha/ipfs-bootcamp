- Install nvm to install latest version of nodejs.
```
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash
```
- Install nodejs
```
nvm install v18.12.1
```
- Install dependencies
```
npm install
```
- Run
```
npm start
```
- Configure CORS in IPFS if needed
```
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["*"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["GET", "POST"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Headers '["Authorization"]'
ipfs config --json API.HTTPHeaders.Access-Control-Expose-Headers '["Location"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Credentials '["true"]'
```
- See in your IPFS desktop or cli that your image is available and has been added.
