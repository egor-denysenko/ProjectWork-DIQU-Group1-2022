---
id: Project_Tecnologies
title: Tecnologies
---

## Firmware/Embedded
* MPLAB X IDE = used to write the code for the microntroller
* PICSIMLAB = used to simulate the microcontroller code
* ARDUINO IDE = used to write the code for the phsical electronic board

## Gateway
* GOLANG = used to write the entire gateway code

## Cloud
* VERNEMQ = used to receive message and send command. We chose to use VerneMQ because is more scalable than MosquittoMQ
* INFLUXDB = used to receive and store datas
* NODE.JS = used to write the API to receive messages from InfluxDB and send to the web app in Angular.js

## WebApp
* ANGULAR.JS = used to write the code for the web app

