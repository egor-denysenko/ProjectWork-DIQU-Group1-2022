---
id: Project_Cloud
title: Cloud
---

## Receiving data
Cloud receive data from raspberry py with MQTT protocol.
The cloud subscribes to the VerneMQ broker topic where the gateway publishes information; this way the data can be saved to the database and displayed.

## Data management
The node.js application receives the data and sends it to InfluxDB to storage it in the database.

## Gallery
