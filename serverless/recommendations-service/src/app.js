
exports.lambdaHandler = async (event, context) => {
    event.Records.map(record => {
        console.log(record);
    });
};