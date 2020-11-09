// Load the AWS SDK for Node.js
const AWS = require('aws-sdk');
// Set the region
AWS.config.update({region: 'us-west-2'});

// Create an SQS service object
const sqs = new AWS.SQS({apiVersion: '2012-11-05'});

const databaseConnection = require('knex')({
    client: 'mysql',
    connection: {
        host     : 'terraform-20201028182302976100000001.cgrpasjjlw1k.us-west-2.rds.amazonaws.com',
        user     : 'newarchitectures',
        password : 'newarchitectures',
        database : 'service_center',
    }
});

exports.lambdaHandler = async (event, context) => {

    await Promise.all(event.Records.map(async record => {
        const {
            correlation_id,
            service_center_name,
        } = JSON.parse(record.body);

        const serviceCenterId = await databaseConnection('service_center').select('id');

        const ackMessage = {
            type: "TechServiceCenter",
            correlation_id,
            service_center_id: serviceCenterId[0].id,
        };

        await new Promise((resolve, reject) => sqs.sendMessage({
            QueueUrl: "https://sqs.us-west-2.amazonaws.com/881619806726/tech-revision-ack.fifo",
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
    }));
};