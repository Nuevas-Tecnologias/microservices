var { _hash } = require("./lib");
var { TP_NAMESPACE } = require("./constants");

class SimpelStoreState {
    constructor(context) {
        this.context = context;
        this.timeout = 500;
        this.stateEntries = {};
    }

    setValue(value) {
        var address = makeAddress(value);
        var stateEntriesSend = {}
        stateEntriesSend[address] = Buffer.from(value);
        return this.context.setState(stateEntriesSend, this.timeout).then(function(result) {
            console.log("Success", result)
            return result;
        }).catch(function(error) {
            console.error("Error", error)
        })
    }

    getValue(value) {
        var address = makeAddress(value);
        return this.context.getState([address], this.timeout).then(function(stateEntries) {
            Object.assign(this.stateEntries, stateEntries);
            console.log("Success", this.stateEntries[address].toString())
        }.bind(this))
    }
}

const makeAddress = (x, label) => TP_NAMESPACE + _hash(x)

module.exports = SimpelStoreState;