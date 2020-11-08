
exports.up = function(knex) {
    return knex.schema.createTable('warranties', function(table) {
        table.increments('id').unsigned().primary();
        table.string('state').defaultTo(1).notNullable();
        table.string('product_type').notNullable();
        table.string('car_plate').notNullable();
        table.timestamp('createdAt').notNullable().defaultTo(knex.fn.now());
        table.timestamp('updatedAt').notNullable().defaultTo(knex.fn.now());
    });
};

exports.down = function(knex) {
    return knex.schema.dropTable('warranties');
};
