"use strict";

const fp = require("fastify-plugin");
const mqtt = require('mqtt')
let mqttConnection


function ConnectBroker() {
    mqttConnection = mqtt.connect(`mqtt://${process.env.MQTT_BROKER_ADDRESS}:${process.env.MQTT_BROKER_PORT}`, { protocolVersion: 5, })
}
ConnectBroker()
console.log("mi sono connesso")
module.exports = fp(async function (fastify, opts) {
    /*function CloseConnectionBroker(){
        mqttConnection.Close()
    }*/
    function SendToTopic(TrainId, message) {
        mqttConnection.publish('trainly/'+ TrainId + '/command', message, { qos: 2, retain: false }, function (error) {
            if (error) {
                console.log(error)
                return error
            } else {
                console.log('Published')
            }
        })
    }
/*
    function SendToLightTopic(TrainId, message) {
        client.publish('trainly/'+ TrainId + '/command', message, { qos: 2, retain: false }, function (error) {
            if (error) {
                console.log(error)
            } else {
                console.log('Published')
            }
        })
    }
*/
    fastify.decorate("SendToTopic", SendToTopic);
});
