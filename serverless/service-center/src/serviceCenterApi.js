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
    const {
        pathParameters: {
            proxy: name
        }
    } = event;

    const serviceCenter = await databaseConnection('service-center').select().where('name', name);

    return {
        "statusCode": 200,
        "body": JSON.stringify(serviceCenter),
        "isBase64Encoded": false
    };
};