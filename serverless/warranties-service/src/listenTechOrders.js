// Load the AWS SDK for Node.js
const AWS = require('aws-sdk');
// Set the region
AWS.config.update({region: 'us-west-2'});

// Create an SQS service object
const sns = new AWS.SNS({apiVersion: '2012-11-05'});

const databaseConnection = require('knex')({
    client: 'mysql',
    connection: {
        host     : 'terraform-20201028182302976100000001.cgrpasjjlw1k.us-west-2.rds.amazonaws.com',
        user     : 'newarchitectures',
        password : 'newarchitectures',
        database : 'warranties',
    }
});

exports.lambdaHandler = async (event, context) => {

    await Promise.all(event.Records.map(async record => {
        const {
            type,
        } = JSON.parse(record.Sns.Message);

        if (type === 'ORDER_CREATED') {
            const {
                order: {
                    items,
                    car_plate
                },
            } = JSON.parse(record.Sns.Message);

            await Promise.all(items.map(async item => {
                if (item.type === 'WARRANTY') {
                    const warrantyData = {
                        product_type: item.product_type,
                        car_plate,
                        state: 'NEW_REQUEST'
                    };
                    console.log(`Creating warranty ${JSON.stringify(warrantyData)}`);
                    const warranty = await databaseConnection('warranties').insert(warrantyData).returning('*');

                    const snsMessage = {
                        type: "WARRANTY_CREATED",
                        warranty,
                    };

                    await sns.publish({
                        TopicArn: "arn:aws:sns:us-west-2:881619806726:technical-orders-topic",
                        Message: JSON.stringify(snsMessage),
                    }).promise();
                } else {
                    console.log(`Skipping item ${item}`);
                }
            }));
        } else {
            console.info(`Skipping event ${type}`);
        }
    }));
};