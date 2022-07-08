"use strict";
const { InfluxDB } = require("@influxdata/influxdb-client");
module.exports = async function (fastify, opts) {
  fastify.get("/", {
    tags: ["Dati"],
    description:
      "Get data about temperature, humidity, door closing status and toilette occupancy.",
    response: {
      200: {
        type: "object",
        properties: {
          data: { type: "array" },
        },
        // 	properties: {
        // 		_value: {type:"number"},
        // IdTrain: { type: "number" },
        // IdWagon: { type: "number" },
        // ADoorIO: { type: "boolean" },
        // ADoorB: { type: "boolean" },
        // ADoorC: { type: "boolean" },
        // ATemperatureMax: { type: "boolean" },
        // ATemperatureMin: { type: "boolean" },
        // ALight: { type: "boolean" },
        // AHumidity: { type: "boolean" },
        // Door1: { type: "boolean" },
        // Door2: { type: "boolean" },
        // Door3: { type: "boolean" },
        // Door4: { type: "boolean" },
        // DoorBath: { type: "boolean" },
        // DoorConduct: { type: "boolean" },
        // Humidity: { type: "number" },
        // Temperature: { type: "number" },
        // LightMode: { type: "boolean" },
        // LightOn: { type: "boolean" },
        // },
        // },
      },
    },
    handler: async function (request, reply) {
      //const influxClient = fastify.chiamainf();
      console.log("sto provando");
      //API token
      const url = "";
      const token =
        "xwiyD0JC_klvoBDTNT4IrT5F3TP8aqELXe0b0f1b3X_qWt0u6Nl3k36CbYvjjYBi8zvoYM6TnceJ9_IXHd5nxg==";
      const org = "andrea.prenot@stud.tecnicosuperiorekennedy.it";
      const bucket = "Trainly";

      const client = new InfluxDB({
        url: "https://eu-central-1-1.aws.cloud2.influxdata.com",
        token: token,
      });

      //const client = new InfluxDB({ url, token: token }).getQueryApi(org);
      const queryApi = client.getQueryApi(org);

      //const fluxQuery ='from(bucket:"Trainly") |> filter(fn: (r) => r._measurement == "MEM")';

      const query = `from(bucket: "Trainly")
			|> range(start:0)
			|> filter(fn: (r) => r["_measurement"] == "mem")
			|> filter(fn: (r) => r["_field"] == "temperature")
			|> yield(name: "last")`;
      let data = [];
      queryApi.queryRows(query, {
        next: (row, tableMeta) => {
          console.log("shrek nella palude");
          //data = tableMeta.toObject(row)
          data.push(tableMeta.toObject(row)["_value"]);
          //console.log(data)
          //console.log(`${data._time} ${data._measurement}: ${data._field}=${data._value}`)
        },
        error: (error) => {
          console.error(error);
          console.log("Finished ERROR");
        },
        complete: () => {
          console.log("Finished SUCCESS");
          console.log(data);
          let gne = { data: data };
          return gne;
        },
      });

      /** Execute a query and receive line table metadata and rows. */
      //client.queryRows(query);
      //return 42;
    },
  });
};
