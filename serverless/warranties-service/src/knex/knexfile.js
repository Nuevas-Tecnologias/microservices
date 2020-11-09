// Update with your config settings.

module.exports = {

  development: {
    client: 'mysql',
    connection: {
      host : 'terraform-20201028182302976100000001.cgrpasjjlw1k.us-west-2.rds.amazonaws.com',
      user : 'newarchitectures',
      password : 'newarchitectures',
      database : 'warranties',
      charset: 'utf8'
    },
    migrations: {
      directory: './migrations',
    },
    seeds: {
      directory: './seeds'
    }
  }

};
