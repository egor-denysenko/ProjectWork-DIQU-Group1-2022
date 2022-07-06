---
id: Project_Brief
title: Project Brief
custom_edit_url: https://github.com/egor-denysenko/ProjectWork-DIQU-Group1-2022/docs/docs/Introduction/ProjectBrief.md
---
## Purpose of the project
Create a system to monitor the status of wagons and to control the operation both from internal train workstation and remotely.

## Context of the project
- The system connect every wagon to a microcontroller board that detect temperature, humidity, door closing status and toilette occupancy
- The system can lead the climatization indipendetely for each wagon, doors opening and emergency signalers
- Every wagon board must have a LCD display that can be able shows messages sent by a remote controller, reads the sensors and modify functioning parameters

## Project planning
### General architecture
![](../../static/img/GeneralArchitecture.jpg)

As can be seen in the image, the sensors send data to the gateway (in this case Raspberry) which sends it to the cloud which displays it to the user.

### Specific architecture
![](../../static/img/SpecificArchitecture.jpg)

As can be seen in the image, the sensors collect and send data to Raspberry; then Raspberry queues data in a non-permanent Redis queue and, in the event that there is no connection, it shows them on the GUI. If there is connection, Raspberry sends a file json to the node.js application that sends data in turn to InfluxDB.

### Initial brainstorming
In our initial brainstorming we have collected as much informations as possible to rearrange ideas. 
During this phase we created the architecture of the project, talked about which technologies use and how to divide the jobs (Look the chapter 'Group Members').

### Trello
![](../../static/img/Trello.jpg)

We used a Trello Kanban to manage and visualize all tasks we need to do to complete the project. Each time a task was completed it was moved to the "done" column.
