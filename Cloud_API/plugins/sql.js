"use strict";
const fp = require("fastify-plugin");
const db = require("pg-promise")();
module.exports = fp(async function (fastify, opts) {
  const psql = db(process.env.DATABASE_URL);
  console.log(await psql.query("SELECT NOW()"));
  fastify.decorate("sql", psql);
});
