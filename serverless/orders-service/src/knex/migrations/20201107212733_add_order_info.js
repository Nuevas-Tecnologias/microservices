
exports.up = function(knex) {
    return knex.schema.table('tech_orders', function (t) {
        t.specificType('state', 'char(10)').nullable();
        t.specificType('car_plate', 'char(10)').nullable();
        t.specificType('transaction_id', 'char(128)').nullable();
    });
};

exports.down = function(knex) {
    return knex.schema.table('tech_orders', function (t) {
        t.dropColumn('state');
        t.dropColumn('car_plate');
        t.dropColumn('transaction_id');
    });
};