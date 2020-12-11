
exports.up = function(knex) {
    return knex.schema.createTable('tech_orders', function(table) {
        table.increments('id').unsigned().primary();
        table.string('service_center_id').notNullable();
        table.text('order_date').notNullable();
        table.timestamp('createdAt').notNullable().defaultTo(knex.fn.now());
        table.timestamp('updatedAt').notNullable().defaultTo(knex.fn.now());
    });
};

exports.down = function(knex) {
    return knex.schema.dropTable('tech_orders');
};
