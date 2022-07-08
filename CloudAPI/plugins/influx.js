"use strict";

const fp = require("fastify-plugin");

const { InfluxDB } = require("@influxdata/influxdb-client");

module.exports = fp(async function (fastify, opts) {
  function ConnectToInflux() {
    //API token
    const token =
      "uIBADP_H45QDzWuiBIkZNubQYDaP5BhZDrCA8rxDAumb1YVGQGYu33yafHEGGbr9KhOjsJdqhM5fDD2-X7TIRA==";
    const org = "trainly";
    const bucket = "Trainly";

    const client = new InfluxDB({ url: "http://localhost:8456", token: token });
  }

  fastify.decorate("chiamainf", ConnectToInflux);
});

const queryApi = client.getQueryApi(org);

const query = `from(bucket: "Trainly") |> range(start: -1h)`;
queryApi.queryRows(query, {
  next(row, tableMeta) {
    const o = tableMeta.toObject(row);
    console.log(`${o.ADoorB}`);
    //    console.log(`${o.ADoorB} ${o._measurement}: ${o._field}=${o._value}`)
  },
  error(error) {
    console.error(error);
    console.log("Finished ERROR");
  },
  complete() {
    console.log("Finished SUCCESS");
  },
});
