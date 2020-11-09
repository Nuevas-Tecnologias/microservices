
exports.up = function(knex) {
    return knex.schema.createTable('recommended_services', function(table) {
        table.increments('id').unsigned().primary();
        table.string('service_name').notNullable();
        table.string('service_sku').notNullable();
        table.string('car_plate').notNullable();
        table.text('description').notNullable();
        table.timestamp('recommendation_date').notNullable();
        table.timestamp('createdAt').notNullable().defaultTo(knex.fn.now());
        table.timestamp('updatedAt').notNullable().defaultTo(knex.fn.now());
    });
};

exports.down = function(knex) {
    return knex.schema.dropTable('recommended_services');
};
