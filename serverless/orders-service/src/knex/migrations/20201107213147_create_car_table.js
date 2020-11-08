
exports.up = function(knex) {
    return knex.schema.createTable('cars', function(table) {
        table.string('car_plate').primary();
        table.string('brand').nullable();
        table.string('model').nullable();
        table.string('line').nullable();
        table.string('color').nullable();
        table.timestamp('createdAt').notNullable().defaultTo(knex.fn.now());
        table.timestamp('updatedAt').notNullable().defaultTo(knex.fn.now());
    });
};

exports.down = function(knex) {
    return knex.schema.dropTable('cars');
};
