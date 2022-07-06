---
id: Project_Gateway
title: Gateway
---

## Receiving data
Data are received via serial and are queued in a non-permanent local Redis queue.

## Receiving data from Redis
Gateway receives data from redis local queue and push them .n the topic to which you subscribe the cloud.

## Receiving command from Web App
Gateway subscribes to another topic to receive commands.

## Send commands to the serial
As the last thing the gateway send commands to the serial.

## Gallery