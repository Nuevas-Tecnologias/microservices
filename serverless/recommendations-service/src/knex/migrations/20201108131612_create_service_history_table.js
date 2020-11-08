
exports.up = function(knex) {
    return knex.schema.createTable('service_history', function(table) {
        table.integer('id').unsigned().primary();
        table.string('name').notNullable();
        table.string('car_plate').notNullable();
        table.integer('tech_center_id').notNullable();
        table.timestamp('createdAt').notNullable().defaultTo(knex.fn.now());
        table.timestamp('updatedAt').notNullable().defaultTo(knex.fn.now());
    });
};

exports.down = function(knex) {
    return knex.schema.dropTable('service_history');
};
