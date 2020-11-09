
exports.up = function(knex) {
    return knex.schema.createTable('car_replacements', function(table) {
        table.string('sku').primary();
        table.string('name').notNullable();
        table.string('producer').notNullable();
        table.string('car_brand').notNullable();
        table.string('car_model').notNullable();
        table.string('warranty_id').notNullable();
        table.timestamp('createdAt').notNullable().defaultTo(knex.fn.now());
        table.timestamp('updatedAt').notNullable().defaultTo(knex.fn.now());
    });
};

exports.down = function(knex) {
    return knex.schema.dropTable('car_replacements');
};
