"use strict";

const fp = require("fastify-plugin");

const { InfluxDB } = require("@influxdata/influxdb-client");

module.exports = fp(async function (fastify, opts) {
  function Query() {
    console.log("sto provando");
    //API token
    const token = process.env.INFLUX_Token;
    const org = process.env.INFLUX_Org;
    const client = new InfluxDB({ url: process.env.INFLUX_Url, token: token });
    const queryApi = client.getQueryApi(org);

    const query = `from(bucket: "Trainly") |> range(start: -1h)`;

    let dato = [];
    queryApi.queryRows(query, {
      next: (row, tableMeta) => {
        const o = tableMeta.toObject(row);
        dato.push(o);
        /*console.log(
					`${o._time} ${o._measurement} in '${o.location}' (${o.example}): ${o._field}=${o._value}`
				  )*/
      },
      error: (error) => {
        console.error(error);
        console.log("\nFinished ERROR");
      },
      complete: () => {
        console.log("\nFinished SUCCESS");
        console.log("gne");
        console.log(JSON.stringify(dato));
        return dato;
      },
    });
  }
  fastify.decorate("query", Query);
});
