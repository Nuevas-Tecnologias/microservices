// Load the AWS SDK for Node.js
const AWS = require('aws-sdk');
// Set the region
AWS.config.update({region: 'us-west-2'});

// Create an SQS service object
const sqs = new AWS.SQS({apiVersion: '2012-11-05'});
const sns = new AWS.SNS({apiVersion: '2012-11-05'});

const databaseConnection = require('knex')({
    client: 'mysql',
    connection: {
        host     : 'terraform-20201028182302976100000001.cgrpasjjlw1k.us-west-2.rds.amazonaws.com',
        user     : 'newarchitectures',
        password : 'newarchitectures',
        database : 'tech_orders',
    }
});

exports.lambdaHandler = async (event, context) => {

    await Promise.all(event.Records.map(async record => {
        const {
            correlation_id,
            order,
        } = JSON.parse(record.body);

        console.log(`Creating order ${JSON.stringify(order)}`);
        const techOrderId = await databaseConnection('tech_orders').insert(order).returning('id')

        const ackMessage = {
            type: "TechOrderCreated",
            correlation_id,
            tech_order_id: techOrderId[0],
        };

        await new Promise((resolve, reject) => sqs.sendMessage({
            QueueUrl: "https://sqs.us-west-2.amazonaws.com/881619806726/tech-revision-ack.fifo",
            MessageBody: JSON.stringify(ackMessage),
            MessageGroupId: "Tech-revision-format"
        }, function (err, data) {
            if (err) {
                console.log("Error", err);
                reject()
            } else {
                resolve()
            }
        }));

        // Emulate order creation to trigger warranties and recommendations
        const snsMessage = {
            type: "ORDER_CREATED",
            order: {
                id: techOrderId[0],
                car_plate: order.car_plate,
                items: [
                    {
                        product_type: 1,
                        type: 'WARRANTY',
                    }
                ]
            },
        };

        await sns.publish({
            TopicArn: "arn:aws:sns:us-west-2:881619806726:technical-orders-topic",
            Message: JSON.stringify(snsMessage),
        }).promise();

        console.log(`Processed order ${techOrderId[0]}`)
    }));
};