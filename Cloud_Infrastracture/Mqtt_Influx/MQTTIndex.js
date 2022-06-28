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
    save_influx(jsonmss)

  } catch (error) {
    console.log("errore banana")
  }
})

function save_influx(message) {

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
        const obj = JSON.parse(message);
        console.log(obj.date)
        const data = new Point('mem')
        //train data
        .intField("IdTrain", /*254*/obj.IdTrain)
        .intField("IdWagon",/* 253*/obj.IdWagon)
        //allarms
        .booleanField("ADoorIO", /* false*/obj.AdoorIO)
        .booleanField("ADoorB", /*false*/obj.ADoorB)
        .booleanField("ADoorC", /*false*/ obj.ADoorC)
        .booleanField("ATemperatureMax", /*false*/obj.ATemperatureMax)
        .booleanField("ATemperatureMin",/* false*/obj.ATemperatureMin)
        .booleanField("ALight", /*false*/ obj.ALight)
        .booleanField("AHumidity", /*false*/obj.AHumidity)
        //Doors
        .booleanField("Door1", /*true*/obj.Door1)
        .booleanField("Door2", /*true*/obj.Door2)
        .booleanField("Door3", /*true*/ obj.Door3)
        .booleanField("Door4", /*true*/ obj.Door4)
        .booleanField("DoorBath", /*true*/obj.DoorBath)
        .booleanField("DoorConduct", /*true*/obj.DoorConduct)
        //Humidity
        .intField("Humidity", /*20*/obj.Humidity)
        //temperature
        .intField("temperature", /*20*/obj.temperature)
        //Lights
        .booleanField("LightMode", /*true*/obj.LightMode)
        .booleanField("LightOn", /*true*/obj.LightOn)
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
    