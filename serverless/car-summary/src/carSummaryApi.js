const axios = require('axios');

const apiGateway = 'https://hsl8ggt76k.execute-api.us-west-2.amazonaws.com/prod';

exports.lambdaHandler = async (event, context) => {
    const {
        pathParameters: {
            proxy: carPlate
        }
    } = event;

    const orders = await axios.get(`${apiGateway}/orders/${carPlate}`);
    const warranties = await axios.get(`${apiGateway}/warranties/${carPlate}`);
    const recommendations = await axios.get(`${apiGateway}/recommendations/${carPlate}`);

    const summary = {
        car_plate: carPlate,
        orders: orders.data,
        warranties: warranties.data,
        recommendations: recommendations.data,
    };

    return {
        "statusCode": 200,
        "body": JSON.stringify(summary),
        "isBase64Encoded": false
    };
};