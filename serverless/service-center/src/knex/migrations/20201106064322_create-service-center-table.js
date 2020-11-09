
exports.up = function(knex) {
    return knex.schema.createTable('service_center', function(table) {
        table.increments('id').unsigned().primary();
        table.string('name').notNullable();
        table.timestamp('createdAt').notNullable().defaultTo(knex.fn.now());
        table.timestamp('updatedAt').notNullable().defaultTo(knex.fn.now());
    });
};

exports.down = function(knex) {
    return knex.schema.dropTable('service-center');
};
