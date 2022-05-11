### Generate New Certs
TODO

### Configure Response
Use the following query params to modify the server response

#### `push`
Enable different response strategies
```js
const PushModes = {
    OFF: '',
    ON: 'push',
    CHUNKED: 'chunked'
};
```

#### `size`
Adjust how many total results are returned
