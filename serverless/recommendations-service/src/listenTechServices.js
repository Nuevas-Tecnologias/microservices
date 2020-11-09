const databaseConnection = require('knex')({
    client: 'mysql',
    connection: {
        host     : 'terraform-20201028182302976100000001.cgrpasjjlw1k.us-west-2.rds.amazonaws.com',
        user     : 'newarchitectures',
        password : 'newarchitectures',
        database : 'recommendations',
    }
});

const recommendations = [
    {
        service_name: 'Cambio de Aceite',
        service_sku: 'SER001',
        description: 'El aceite de motor es fundamental para el buen funcionamiento del vehículo. Su acción lubricante permite el movimiento de las piezas internas sin fricción lo cual reduce el desgaste entre ellas. Pero esta no es su única función, el aceite también es agente detergente y refrigerante debido a que en todo su recorrido recoge pequeñas impurezas que se depositan en el filtro y absorbe cierta parte del calor generado.',
    },
    {
        service_name: 'Revision de motor',
        service_sku: 'SER002',
        description: 'Observar posibles goteos tanto de aceite como de agua o de líquido refrigerante. Localizar la fuente del goteo y detenerlo es muy importante ya que estos siempre conducen a problemas mayores.',
    },
    {
        service_name: 'Alineación',
        service_sku: 'SER003',
        description: 'Básicamente, la alineación es el proceso en el que se ajustan las llantas de un vehículo para que miren hacia el frente, es decir, los neumáticos de tu auto deben quedar paralelos entre sí y perpendiculares al camino, (el tiempo óptimo para realizar este servicio, se encuentra en el manual de propietario)',
    },
    {
        service_name: 'Revision de frenos',
        service_sku: 'SER004',
        description: 'Encuentra Kit Cambio De Pastillas Frenos en Mercado Libre Colombia. Descubre la mejor forma de comprar online.',
    },
    {
        service_name: 'Latoneria y pintura',
        service_sku: 'SER005',
        description: 'El proceso de Latonería consiste en restaurar las piezas del vehículo y la pintura se hace con pintura anticorrosiva, que cumple la función de prevenir el óxido, mientras mejora la estética del carro.',
    }
];

exports.lambdaHandler = async (event, context) => {

    await Promise.all(event.Records.map(async record => {
        const {
            type,
        } = JSON.parse(record.Sns.Message);

        if (type === 'SERVICE_CREATED') {
            const {
                registroServicio: {
                    id,
                    nombreServicio: name,
                    placaVehiculo: car_plate,
                    idCentroServicio: tech_center_id,
                }
            } = JSON.parse(record.Sns.Message);

            await databaseConnection('service_history').insert({
                id,
                name,
                car_plate,
                tech_center_id,
            });

            // Emulate recommendation engine
            const min = 1;
            const max = 5;
            const seed = Math.floor(Math.random() * (max - min) + min);

            await databaseConnection('recommended_services').insert({
                ...recommendations[seed],
                car_plate,
                recommendation_date: '2021-01-01',
            });

        } else {
            console.info(`Skipping event ${type}`);
        }
    }));
};