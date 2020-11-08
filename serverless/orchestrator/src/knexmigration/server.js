const express = require('express');
const PORT = process.env.PORT || 3307;
const knex = require('./knex/knex.js');
const app = express();

app.get('/tasks', (req, res) => {
  // use the knex variable above to create dynamic queries
});

app.listen(PORT, () => {
  console.log(`Listening on port: ${PORT}`);
});