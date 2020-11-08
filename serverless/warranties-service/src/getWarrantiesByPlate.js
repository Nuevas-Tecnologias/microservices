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
    const {
        pathParameters: {
            proxy: carPlate
        }
    } = event;

    const warranties = await databaseConnection('warranties').select()
        .join('car_replacements', 'car_replacements.warranty_id', 'warranties.id')
        .where('car_plate', carPlate);

    return {
        "statusCode": 200,
        "body": JSON.stringify(warranties),
        "isBase64Encoded": false
    };
};