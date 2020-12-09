const databaseConnection = require('knex')({
    client: 'mysql',
    connection: {
        host     : 'terraform-20201028182302976100000001.cgrpasjjlw1k.us-west-2.rds.amazonaws.com',
        user     : 'newarchitectures',
        password : 'newarchitectures',
        database : 'tech_orders',
    }
});

const { getTransaction, _hash } = require('./blockchain-gateway');

exports.lambdaHandler = async (event, context) => {
    const {
        pathParameters: {
            proxy: carPlate
        }
    } = event;

    const orders = await databaseConnection('tech_orders').select().where('car_plate', carPlate);

    const validatedOrders = await Promise.all(orders.map((order => new Promise(async (resolve) => {
            const transaction = await getTransaction(order.transaction_id);
            const orderHash = transaction.split('@')[1];

            const orderData = {
                service_center_id: order.service_center_id,
                order_date: order.order_date,
                car_plate: order.car_plate,
            }

            resolve({
                ...order,
                trustworthy: _hash(JSON.stringify(orderData)) === orderHash
            });
        })
    )));

    return {
        "statusCode": 200,
        "body": JSON.stringify(validatedOrders),
        "isBase64Encoded": false
    };
};