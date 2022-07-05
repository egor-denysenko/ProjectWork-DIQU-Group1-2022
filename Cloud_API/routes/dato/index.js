"use strict";
const { InfluxDB,FluxTableMetaData } = require("@influxdata/influxdb-client");
const { default: fp } = require("fastify-plugin");

module.exports = async function (fastify, opts) {
<<<<<<< HEAD
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
=======
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
>>>>>>> 1773b6ced36dae2e35cd88a27a7ba5425499ab42
};
