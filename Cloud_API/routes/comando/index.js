"use strict";
const { InfluxDB } = require("@influxdata/influxdb-client");
module.exports = async function (fastify, opts) {
  fastify.post("/temperature", {
    schema: {
      tags: ["Commands"],
      description:
        "Recive command to send the selected temperature to the gateway",
      body: {
        type: "object",
        properties: {
          TrainId: { type: "integer" },
          WagonId: { type: "integer" },
          WagonCommand: { type: "integer" },
          Temperature: { type: "integer" },
        },
      },
    },
    response: {
      200: {
        type: "object",
        properties: {
          status: { type: "boolean" },
        },
      },
    },
    handler: async function (request, reply) {
      const { TrainId } = request.body;

      try {
        delete request.body["TrainId"];
        fastify.SendToTopic(TrainId, request.body.toString());
        console.log("mando temp");
        return true;
      } catch (err) {
        console.log(err);
        throw err;
      }
    },
  });

  fastify.post("/lights", {
    schema: {
      tags: ["Commands"],
      description:
        "Recive command to send the selected light mode to the gateway",
      body: {
        type: "object",
        properties: {
          TrainId: { type: "integer" },
          WagonId: { type: "integer" },
          WagonCommand: { type: "integer" },
          LightMode: { type: "boolean" },
        },
      },
    },
    response: {
      200: {
        type: "object",
        properties: {
          status: { type: "boolean" },
        },
      },
    },
    handler: async function (request, reply) {
      const { TrainId } = request.body;
      try {
        delete request.body["TrainId"];
        fastify.SendToTopic(TrainId, request.body.toString());
        return true;
      } catch (err) {
        console.log(err);
        throw err;
      }
    },
  });
};
