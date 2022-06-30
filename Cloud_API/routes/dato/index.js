"use strict";
const { InfluxDB } = require("@influxdata/influxdb-client");
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
						LightOn: { type: "boolean" },
					},
				},
			},
		},
		handler: async function (request, reply) {
			//const influxClient = fastify.chiamainf();
			console.log("sto provando");
			//API token
			const url =
				"";
			const token =
				"";
			const org = "";
			const bucket = "Trainly";

			const client = new InfluxDB({ url, token: token }).getQueryApi(org);
			const fluxQuery =
				'from(bucket:"Trainly") |> filter(fn: (r) => r._measurement == "MEM")';

			const fluxObserver = {
				next(row, tableMeta) {
					const o = tableMeta.toObject(row);
					/*console.log(
						`${o._time} ${o._measurement} in ${o.region} (${o.sensor_id}): ${o._field}=${o._value}`
					);*/
				},
				error(error) {
					console.error(error);
					console.log("\nFinished ERROR");
				},
				complete() {
					console.log("\nFinished SUCCESS");
				},
			};

			/** Execute a query and receive line table metadata and rows. */
			client.queryRows(fluxQuery, fluxObserver);
			return 42;
		},
	});
};
