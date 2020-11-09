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
        database : 'service-center',
    }
});

exports.lambdaHandler = async (event, context) => {

    await Promise.all(event.Records.map(async record => {
        const {
            correlation_id,
            service_center,
        } = JSON.parse(record.body);

        console.log(`Creating service center ${JSON.stringify(service_center)}`);
        const serviceCenterId = await databaseConnection('service-center').insert(service_center).returning('id')

        const ackMessage = {
            type: "ServiceCenterCreated",
            correlation_id,
            service_center_id: serviceCenterId[0],
        };

        await new Promise((resolve, reject) => sqs.sendMessage({
            QueueUrl: "https://sqs.us-west-2.amazonaws.com/881619806726/save-tech-order-command.fifo",
            //QueueUrl: "https://sqs.us-west-2.amazonaws.com/881619806726/service-center-ack.fifo",
            MessageBody: JSON.stringify(ackMessage),
            MessageGroupId: "Service-center-format"
        }, function (err, data) {
            if (err) {
                console.log("Error", err);
                reject()
            } else {
                resolve()
            }
        }));

        // Emulate service center creation
        const snsMessage = {
            type: "SERVICE_CENTER_CREATED",
            serviceCenter: {
                id: serviceCenterId[0],
                name: serviceCenter.name,
            },
        };

        await sns.publish({
            TopicArn: "arn:aws:sns:us-west-2:881619806726:service-center-topic",
            Message: JSON.stringify(snsMessage),
        }).promise();

        console.log(`Processed service center ${serviceCenterId[0]}`)
    }));
};