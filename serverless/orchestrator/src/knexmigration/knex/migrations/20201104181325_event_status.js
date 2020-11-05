
exports.up = function(knex) {
    return knex.schema.createTable('event_status', function(table) {
        table.increments('id').unsigned().primary();
        table.integer('correlation_id').nullable();
        table.string('status').notNull();
        table.text('description').notNull();
        table.text('error').notNull();
        table.timestamp('createdAt').notNull().defaultTo(knex.fn.now());
        table.timestamp('updatedAt').notNull().defaultTo(knex.fn.now());
    });
};

exports.down = function(knex) {
    return knex.schema.dropTable('event_status');
};
