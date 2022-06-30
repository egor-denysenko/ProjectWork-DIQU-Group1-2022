'use strict'

module.exports = async function (fastify, opts) {
  fastify.get("/", {
    tags: ["Dati"],
    description: "Get data about temperature, humidity, door closing status and toilette occupancy.",
    response: {
      200: {
        type: "array",
        items: {
          type: "object",
          properties: {
            IdTrain: { type: "number" },
            IdWagon: { type: "number" },
            ADoorIO: { type: "boolean" },
            ADoorB: { type: "boolean" },
            ADoorC: { type: "boolean" },
            ATemperatureMax: { type: "boolean" },
            ATemperatureMin: { type: "boolean" },
            ALight: { type: "boolean" },
            AHumidity: { type: "boolean" },
            Door1: { type: "boolean" },
            Door2: { type: "boolean" },
            Door3: { type: "boolean" },
            Door4: { type: "boolean" },
            DoorBath: { type: "boolean" },
            DoorConduct: { type: "boolean" },
            Humidity: { type: "number" },
            Temperature: { type: "number" },
            LightMode: { type: "boolean" },
            LightOn: { type: "boolean" }
          },
        },
      },
    },
    handler: async function (request, reply) {
      fastify.chiamainf()      
    },
  });
}
