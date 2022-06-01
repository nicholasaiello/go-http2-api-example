const path = require('path');

module.exports = {
    output: {
        path: path.resolve(__dirname, '../assets')
    },
    mode: 'production',
    entry: './client.js'
};
