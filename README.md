### Generate New Certs
Run the following, and specify your local domain `example.com` for the Company Name:
```sh
$ openssl req -new -newkey rsa:2048 -days 3650 -nodes -x509 -keyout server.key -out server.crt
```

### Build protobuf
Server
```sh
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/holdings.proto 
```

Client
```sh
$ protoc -I=./proto holdings.proto  \
    --js_out=import_style=commonjs:./client/js \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client/js
```

### Configure Response
Use the following query params to modify the server response

#### `push`
Enable different response strategies
```js
const PushModes = {
    OFF: '',
    ON: 'push',
};
```

#### `chunks`
Enable different response strategies
```js
const ChunkModes = {
    OFF: '',
    ON: 'chunks',
};
```

#### `size`
Adjust how many total results are returned

