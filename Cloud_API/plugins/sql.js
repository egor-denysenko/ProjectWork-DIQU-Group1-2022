"use strict";
const fp = require("fastify-plugin");
const {Client} = require("pg");
const client = new Client(process.env.DATABASE_URL);
module.exports = fp(async function (fastify, opts) {
  (async () => {
    await client.connect();
    try {
      const results = await client.query("SELECT NOW()");
      console.log(results);
    } catch (err) {
      console.error("error executing query:", err);
    } finally {
      client.end();
    }
  })();
});
