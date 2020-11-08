const databaseConnection = require('knex')({
    client: 'mysql',
    connection: {
        host     : 'terraform-20201028182302976100000001.cgrpasjjlw1k.us-west-2.rds.amazonaws.com',
        user     : 'newarchitectures',
        password : 'newarchitectures',
        database : 'recommendations',
    }
});

exports.lambdaHandler = async (event, context) => {
    const {
        pathParameters: {
            proxy: carPlate
        }
    } = event;

    const recommendations = await databaseConnection('recommended_services').select()
        .where('car_plate', carPlate);

    return {
        "statusCode": 200,
        "body": JSON.stringify(recommendations),
        "isBase64Encoded": false
    };
};