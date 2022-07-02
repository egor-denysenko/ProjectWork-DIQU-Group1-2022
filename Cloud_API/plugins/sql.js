"use strict";
const fp = require("fastify-plugin");
const db = require("pg-promise")();
module.exports = fp(async function (fastify, opts) {
  const cockroach = db(process.env.DATABASE_URL);

  fastify.decorate("cockroachDB", cockroach);
});
