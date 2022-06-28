var mqtt = require('mqtt')
var client  = mqtt.connect('mqtt://test.mosquitto.org')
require('dotenv').config()

//subscription on the topic
client.on('connect', function () {
  client.subscribe('trainly/+/+/status')
//                 Trainly/"IdTreno"/"idCarrozza"/status
})

//save_influx()

//read data from the tail
client.on('message', function (topic, message) {
  // message is Buffer

  console.log(message.toString())
  try {
    message.toString()
    jsonmss = JSON.parse(message.toString())
    console.log(jsonmss)
    save_influx(jsonmss)
  } catch (error) {
    console.log(error)
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
    const client = new InfluxDB({url: 'http://localhost:8456', token: token})
    
    const {Point} = require('@influxdata/influxdb-client')
        const writeApi = client.getWriteApi(org, bucket)
        writeApi.useDefaultTags({host: 'host1'})

        //parse the json file recived from mosquitto
        //console.log(obj.date)
        const data = new Point('mem')
        //train data
        .intField("IdTrain", /*254*/parsedMessage.IdTrain)
        .intField("IdWagon",/* 253*/parsedMessage.IdWagon)
        //allarms
        .booleanField("ADoorIO", /* false*/parsedMessage.AdoorIO)
        .booleanField("ADoorB", /*false*/parsedMessage.ADoorB)
        .booleanField("ADoorC", /*false*/ parsedMessage.ADoorC)
        .booleanField("ATemperatureMax", /*false*/parsedMessage.ATemperatureMax)
        .booleanField("ATemperatureMin",/* false*/parsedMessage.ATemperatureMin)
        .booleanField("ALight", /*false*/ parsedMessage.ALight)
        .booleanField("AHumidity", /*false*/parsedMessage.AHumidity)
        //Doors
        .booleanField("Door1", /*true*/parsedMessage.Door1)
        .booleanField("Door2", /*true*/parsedMessage.Door2)
        .booleanField("Door3", /*true*/ parsedMessage.Door3)
        .booleanField("Door4", /*true*/ parsedMessage.Door4)
        .booleanField("DoorBath", /*true*/parsedMessage.DoorBath)
        .booleanField("DoorConduct", /*true*/parsedMessage.DoorConduct)
        //Humidity
        .intField("Humidity", /*20*/parsedMessage.Humidity)
        //temperature
        .intField("temperature", /*20*/parsedMessage.temperature)
        //Lights
        .booleanField("LightMode", /*true*/parsedMessage.LightMode)
        .booleanField("LightOn", /*true*/parsedMessage.LightOn)
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
    