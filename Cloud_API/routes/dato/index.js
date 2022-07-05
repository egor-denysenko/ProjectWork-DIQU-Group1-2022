"use strict";
const { InfluxDB,FluxTableMetaData } = require("@influxdata/influxdb-client");
const { default: fp } = require("fastify-plugin");

module.exports = async function (fastify, opts) {
	fastify.get("/", {
		tags: ["Dati"],
		description:
			"Get data about temperature, humidity, door closing status and toilette occupancy.",
		response: {
			200: {
				type: 'object',
				properties: {
				  data: { type: 'object' }
				}
			},
		},
		handler: async function (request, reply) {
			const datoRic = fastify.query()
			return datoRic
			//const influxClient = fastify.chiamainf();
			/*console.log("sto provando");
			//API token
			const token = process.env.INFLUX_Token;
			const org = process.env.INFLUX_Org;

			const client = new InfluxDB({ url: process.env.INFLUX_Url, token: token })
		
				//const client = new InfluxDB({ url, token: token }).getQueryApi(org);
			const queryApi = client.getQueryApi(org)

			//const query = `from(bucket: "Trainly") |> range(start: -1h)`

			const query = `from(bucket: "Trainly")
			|> ran<ge(start: v.timeRangeStart, stop: v.timeRangeStop)
			|> filter(fn: (r) => r["_measurement"] == "mem")
			|> filter(fn: (r) => r["_field"] == "ADoorB" or r["_field"] == "ADoorC" or r["_field"] == "ADoorIO" or r["_field"] == "AHumidity" or r["_field"] == "ALight" or r["_field"] == "ATemperatureMax" or r["_field"] == "Door1" or r["_field"] == "ATemperatureMin" or r["_field"] == "Door2" or r["_field"] == "Door3" or r["_field"] == "Door4" or r["_field"] == "DoorBath" or r["_field"] == "DoorConduct" or r["_field"] == "Humidity" or r["_field"] == "IdTrain" or r["_field"] == "IdWagon" or r["_field"] == "LightMode" or r["_field"] == "LightOn" or r["_field"] == "temperature")
			|> aggregateWindow(every: v.windowPeriod, fn: last, createEmpty: false)
			|> yield(name: "last")`

			queryApi.queryRows(query, {
				next: (row, tableMeta) => {
				  // the following line creates an object for each row
				  const o = tableMeta.toObject(row)
				  // console.log(JSON.stringify(o, null, 2))
				  console.log(
					`${o._time} ${o._measurement} in '${o.location}' (${o.example}): ${o._field}=${o._value}`
				  )
				},
				error: (error) => {
				  console.error(error)
				  console.log('\nFinished ERROR')
				},
				complete: () => {
				  console.log('\nFinished SUCCESS')
				},
			  })

			  return "pippo"*/
			//Execute a query and receive line table metadata and rows.
			//client.queryRows(query);
			//return 42;
		},
	});
};
