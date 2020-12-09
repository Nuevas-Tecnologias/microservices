
const { createContext, CryptoFactory } = require('sawtooth-sdk/signing')
const { protobuf } = require('sawtooth-sdk')
const cbor = require('cbor');
const axios = require('axios');
const crypto = require('crypto')

const TP_FAMILY = 'simplestore';

const context = createContext('secp256k1')
const privateKey = context.newRandomPrivateKey()
const signer = (new CryptoFactory(context)).newSigner(privateKey)

const HOST = 'http://sawtooth-rest-api-default:8008';

const sendRequest = (payload) => {

    const payloadBytes = cbor.encode(payload);

    const transactionHeaderBytes = protobuf.TransactionHeader.encode({
        familyName: TP_FAMILY,
        familyVersion: '1.0',
        inputs: ['917479'],
        outputs: ['917479'],
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

const createTransaction = ({data}) => {
    const payload = {
        action: 'set',
        data,
    };
    return sendRequest(payload);
}

const getBatchInfo = (batchLink) => {
    return axios.get(batchLink)
        .then((response) => response.data.data[0]);
}

const getTransactions = (batch) => {
    return axios.get(`${HOST}/batches/${batch}`)
        .then((response) => response.data.data.header.transaction_ids);
}

module.exports = {
    createTransaction,
    getTransactions,
    getBatchInfo,
}