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

export const getOrder = async (carPlate) => {

    const orders = await databaseConnection('tech_orders').select().where('car_plate', carPlate);

    const validatedOrders = await Promise.all(orders.map((order =>
        new Promise(async (resolve) => {
            let blockchain = {}
            const {transaction_id, ...filteredOrder} = order;
            if (transaction_id) {
                try {
                    const transaction = await getTransaction(transaction_id);
                    const orderHash = transaction.split('@')[1];

                    const orderData = {
                        service_center_id: order.service_center_id,
                        order_date: order.order_date,
                        car_plate: order.car_plate,
                    }

                    blockchain = {
                        transaction_id,
                        trustworthy: _hash(JSON.stringify(orderData)) === orderHash
                    }
                } catch (err) {
                    blockchain = {
                        transaction_id,
                        trustworthy: false
                    }
                }
            }

            const validatedOrder = {
                ...filteredOrder,
                blockchain,
            }
            resolve(validatedOrder);
        })
    )));

    return validatedOrders;
};