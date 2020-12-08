// Load the AWS SDK for Node.js
const AWS = require('aws-sdk');
// Set the region
AWS.config.update({region: 'us-west-2'});

// Create an SQS service object
const sqs = new AWS.SQS({apiVersion: '2012-11-05'});

exports.lambdaHandler = async (event, context) => {
    const {
        body,
    } = event;

    await new Promise((resolve, reject) => sqs.sendMessage({
        QueueUrl: "https://sqs.us-west-2.amazonaws.com/881619806726/process-tech-revision-format-command.fifo",
        MessageBody: body,
        MessageGroupId: "Tech-revision-format"
    }, function (err, data) {
        if (err) {
            console.log("Error", err);
            reject()
        } else {
            resolve()
        }
    }));

    return {
        "statusCode": 200
    };
};