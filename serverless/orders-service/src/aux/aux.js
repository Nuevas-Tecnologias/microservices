const cbor = require('cbor');
const axios = require('axios');

const HOST = 'http://localhost:8008';

axios.get(`${HOST}/transactions/df22567e12ab9f23252afa01662120d890256a45b5777674ea06158a07afffe50a48bf3655706a0e8ba7485715726e4aa7c7e199018a88c99df02d0ad3c29870`)
    .then((response) => Buffer.from(response.data.data.payload, 'base64').toString('ascii'))
    .then((data) => console.log(data.split('@')[1]))
