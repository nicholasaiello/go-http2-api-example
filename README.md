### Generate New Certs
Run the following, and specify `localhost` for the Company Name:
```sh
$ openssl req -new -newkey rsa:2048 -days 3650 -nodes -x509 -keyout server.key -out server.crt
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
