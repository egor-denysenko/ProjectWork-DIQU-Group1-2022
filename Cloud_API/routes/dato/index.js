"use strict";
const { InfluxDB, FluxTableMetaData } = require("@influxdata/influxdb-client");
const { default: fp } = require("fastify-plugin");

module.exports = async function (fastify, opts) {
  fastify.get("/", {
    tags: ["Dati"],
    description:
      "Get data about temperature, humidity, door closing status and toilette occupancy.",
    response: {
      200: {
        type: "array",
        items: {
          type: "object",
        },
      },
    },
    handler: async function (request, reply) {
      console.log("ti ritorno dato");
      return await fastify.query();
    },
  });
};
