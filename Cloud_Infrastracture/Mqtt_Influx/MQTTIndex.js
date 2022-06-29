let mqtt = require('mqtt')
//let client  = mqtt.connect('mqtt://test.mosquitto.org')
const client = mqtt.connect('mqtt://20.82.72.240:1883')
require('dotenv').config()

//subscription on the topic
client.on('connect', function () {
  client.subscribe('trainly/+/+/status')
})

//read data from the tail
client.on('message', function (topic, message) {
  try {
    console.log(message.toString())
    jsonmss = JSON.parse(message.toString())
    console.log(jsonmss)
    //jsonmss = message.toJSON();
    //console.log(jsonmss)
    save_influx(jsonmss)
  } catch (error) {
    console.error(error)
    console.log("errore formato del pacchetto non corretto")
  }
})

function save_influx(parsedMessage) {

  //info to write data to InfluxDB
  const {InfluxDB} = require('@influxdata/influxdb-client')
  const token = process.env.Influx_Token
  const org = process.env.Influx_Org
  const bucket = process.env.Influx_Bucket

  //connect to influxdb
  const client = new InfluxDB({url: 'https://eu-central-1-1.aws.cloud2.influxdata.com', token: token})

  const { Point } = require('@influxdata/influxdb-client')
  const writeApi = client.getWriteApi(org, bucket)
  writeApi.useDefaultTags({ host: 'host1' })

  const data = new Point('mem')
    //train data
    .intField("IdTrain", parsedMessage.IdTrain)
    .intField("IdWagon",parsedMessage.IdWagon)
    //allarms
    .booleanField("ADoorIO", parsedMessage.AdoorIO)
    .booleanField("ADoorB", parsedMessage.ADoorB)
    .booleanField("ADoorC",  parsedMessage.ADoorC)
    .booleanField("ATemperatureMax", parsedMessage.ATemperatureMax)
    .booleanField("ATemperatureMin",parsedMessage.ATemperatureMin)
    .booleanField("ALight", parsedMessage.ALight)
    .booleanField("AHumidity", parsedMessage.AHumidity)
    //Doors
    .booleanField("Door1", parsedMessage.Door1)
    .booleanField("Door2", parsedMessage.Door2)
    .booleanField("Door3", parsedMessage.Door3)
    .booleanField("Door4", parsedMessage.Door4)
    .booleanField("DoorBath", parsedMessage.DoorBath)
    .booleanField("DoorConduct", parsedMessage.DoorConduct)
    //Humidity
    .floatField("Humidity", parsedMessage.Humidity)
    //temperature
    .floatField("temperature", parsedMessage.temperature)
    //Lights
    .booleanField("LightMode", parsedMessage.LightMode)
    .booleanField("LightOn", parsedMessage.LightOn)
  writeApi.writePoint(data);

  writeApi
    .close()
    .then(() => {
      console.log('FINISHED')
    })
    .catch(e => {
      console.error(e)
      console.log('\\nFinished ERROR')
    })
}