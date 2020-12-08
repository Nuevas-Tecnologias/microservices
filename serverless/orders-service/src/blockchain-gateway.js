
const { createContext, CryptoFactory } = require('sawtooth-sdk/signing')
const { protobuf } = require('sawtooth-sdk')
const cbor = require('cbor');
const axios = require('axios');
const crypto = require('crypto')

const hash = (x) =>
    crypto.createHash('sha512').update(x).digest('hex').toLowerCase()

const INT_KEY_FAMILY = 'intkey'
const INT_KEY_NAMESPACE = hash(INT_KEY_FAMILY).substring(0, 6)
const address = INT_KEY_NAMESPACE + hash('foo').slice(-64)

const context = createContext('secp256k1')
const privateKey = context.newRandomPrivateKey()
const signer = (new CryptoFactory(context)).newSigner(privateKey)

const HOST = 'http://sawtooth-rest-api-default:8008';

const createTransaction = ({data}) => {

    const payload = {
        action: 'set',
        data,
    };

    const payloadBytes = cbor.encode(payload);

    const transactionHeaderBytes = protobuf.TransactionHeader.encode({
        familyName: 'simplestore',
        familyVersion: '1.0',
        inputs: [address],
        outputs: [address],
        signerPublicKey: signer.getPublicKey().asHex(),
        // In this example, we're signing the batch with the same private key,
        // but the batch can be signed by another party, in which case, the
        // public key will need to be associated with that key.
        batcherPublicKey: signer.getPublicKey().asHex(),
        dependencies: [],
        payloadSha512: crypto.createHash('sha512').update(payloadBytes).digest('hex'),
        nonce:"hey4"
    }).finish();

    let signature = signer.sign(transactionHeaderBytes);

    const transaction = protobuf.Transaction.create({
        header: transactionHeaderBytes,
        headerSignature: signature,
        payload: payloadBytes
    });

    const batchHeaderBytes = protobuf.BatchHeader.encode({
        signerPublicKey: signer.getPublicKey().asHex(),
        transactionIds: [transaction].map((txn) => txn.headerSignature),
    }).finish();

    signature = signer.sign(batchHeaderBytes)

    const batch = protobuf.Batch.create({
        header: batchHeaderBytes,
        headerSignature: signature,
        transactions: [transaction]
    });

    const batchListBytes = protobuf.BatchList.encode({
        batches: [batch]
    }).finish()

    return axios.post(`${HOST}/batches`, batchListBytes, {
        headers: {'Content-Type': 'application/octet-stream'}
    });
}

const getTransaction = ({ address }) => {
    axios({
        method: 'get',
        url: `${HOST}/state/${address}`,
        headers: {'Content-Type': 'application/json'}
    })
        .then(function (response) {
            let base = Buffer.from(response.data.data, 'base64');
            let stateValue = cbor.decodeFirstSync(base);
            console.log(stateValue);
        })
        .catch(err => {
            console.log(err);
        });
}

module.exports = {
    createTransaction,
    getTransaction,
}