const { TransactionHandler } = require('sawtooth-sdk/processor/handler')
const { InvalidTransaction, InternalError } = require('sawtooth-sdk/processor/exceptions')
const cbor = require('cbor')
const SimpleStoreState = require('./state');
var { TP_FAMILY, TP_NAMESPACE } = require("./constants");

class SimpleStoreHandler extends TransactionHandler {
    constructor() {
        super(TP_FAMILY, ['1.0'], [TP_NAMESPACE])
    }

    apply(transactionProcessRequest, context) {
        let payload = cbor.decode(transactionProcessRequest.payload);
        let simpleStoreState = new SimpleStoreState(context);

        if (payload.action === 'get') {
            console.log(`getting ${payload.data}`);
            return simpleStoreState.getValue(payload.data)
        } else  if (payload.action === 'set') {
            console.log(`setting ${payload.data}`);
            return simpleStoreState.setValue(payload.data)
        } else {
            throw  new InvalidTransaction(
                `Action must be get or set not ${payload.action}`
            )
        }
    }
}

module.exports = SimpleStoreHandler;