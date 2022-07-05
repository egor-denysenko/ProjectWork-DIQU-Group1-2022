---
id: Project_Firmware
title: Firmware
---

## Collect and sending data
The firmware part include data collection and sending data; data collects by the sensors are temperature, humidity, status of wagons and toilette occupancy. Datas are collect by BUS RS485 that sent them to the gateway. 

We use PicSimLab to simulate the electronic board that due to technical problems we could not get working.

## Gallery
The following picture shows the main screen of the LCD display.

![](../../static/img/1Embedded.jpg)

RB3 button opens the doors.

![](../../static/img/2Embedded.jpg)

RB4 button closes the doors.

![](../../static/img/3Emdedded.jpg)

RB5 button sets the bathroom as busy.

![](../../static/img/4Embedded.jpg)


## Behind the scene
In this section we are going to give a brief explanation about the used code.

This function has the aim to read the data (like temperature) and return a value between 0 and 1024.

![](../../static/img/letturaDatoAnalogico.jpg)

This function transform a int value in a string value to show it in the LCD display.

![](../../static/img/ConversioneIntString.jpg)


This function has to send data or command to LCD display.

![](../../static/img/InvioDatoComandoLCD.jpg)

This function initialize the LCD display.

![](../../static/img/InizializzazioneLCD.jpg)