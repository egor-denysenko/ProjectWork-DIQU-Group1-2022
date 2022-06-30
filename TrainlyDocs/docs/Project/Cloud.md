---
id: Project_Cloud
title: Cloud
---

## Receiving data
Cloud receive data from raspberry py with MQTT protocol.
The cloud subscribes to the VerneMQ broker topic where the gateway publishes information; this way the data can be saved to the database and displayed.

## Data management
The node.js application receives the data and sends it to InfluxDB to storage it in the database.

## Web app
We created a node.js web app that receives data from InfluxDB and shows them in a web interface; the aim of this web app is to send commands to the gateway that sends in turns to the microcontroller that modify the value of temperaturea and the status of the doors (open/close).

## Gallery
