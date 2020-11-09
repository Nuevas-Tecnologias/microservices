
exports.up = function(knex) {
    return knex.schema.table('service-center', function (t) {
        t.specificType('name', 'string').nullable();
    });
};

exports.down = function(knex) {
    return knex.schema.table('service-center', function (t) {
        t.dropColumn('name');
    });
};
