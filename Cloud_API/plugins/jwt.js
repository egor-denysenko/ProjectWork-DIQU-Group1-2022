"use strict";

const fp = require("fastify-plugin");
const jsonwebtoken = require("jsonwebtoken");

module.exports = fp(async function (fastify, opts) {
  fastify.decorate("jwt", jsonwebtoken);
});
