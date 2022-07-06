"use strict";

const fp = require("fastify-plugin");

const { InfluxDB } = require("@influxdata/influxdb-client");

module.exports = fp(async function (fastify, opts) {
  async function Query() {
    console.log("sto provando");
    //API token
    const token = process.env.INFLUX_Token;
    const org = process.env.INFLUX_Org;
    const client = new InfluxDB({ url: process.env.INFLUX_Url, token: token });
    const queryApi = client.getQueryApi(org);

    const query = `from(bucket: "Trainly")
	|> range(start: -24h)
	|> filter(fn: (r) => r["_field"] == "Humidity" or r["_field"] == "temperature")
	|> yield(name: "last")`;
    return queryApi
      .collectRows(query /*, you can specify a row mapper as a second arg */)
      .then((data) => {
        if (data._field == "temperature" || data._field == "Humidity") {
        	datoObject._time = o._time;
        	datoObject._field = o._field;
        	datoObject._value = o._value;
        	dato.push(datoObject);
        }
        //data.forEach(x => console.log(JSON.stringify(x)))
        console.log("\nCollect ROWS SUCCESS");
		return JSON.stringify(data)
      })
      .catch((error) => {
        console.error(error);
        console.log("\nCollect ROWS ERROR");
      });
  }

  // function Query() {

  // 	var dato = [];
  // 	let datoObject = new Object();
  // 	queryApi.queryRaw(query).then(result=> {
  // 		return resolve(result)
  //})
  // queryApi.queryRows(query, {
  //   next: (row, tableMeta) => {
  // 	const o = tableMeta.toObject(row);
  // 	if (o._field == "temperature" || o._field == "Humidity") {
  // 	  datoObject._time = o._time;
  // 	  datoObject._field = o._field;
  // 	  datoObject._value = o._value;
  // 	  dato.push(datoObject);
  // 	}

  // 	/*console.log(
  // 				`${o._time} ${o._measurement} in '${o.location}' (${o.example}): ${o._field}=${o._value}`
  // 			  )*/
  //   },
  //   error: (error) => {
  // 	console.error(error);
  // 	console.log("\nFinished ERROR");
  // 	reject(error);
  //   },
  //   complete: () => {
  // 	console.log("\nFinished SUCCESS");
  // 	return resolve(dato);
  //   },
  // });
  //}

  fastify.decorate("query", Query);
});
