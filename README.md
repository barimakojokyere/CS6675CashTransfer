# CS6675CashTransfer
This is the final project for OMSCS6675. It is a prototype of a money transfer program that will allow transfer of money from Ghana to the United States. A Demo of it can be accessed from here - https://drive.google.com/file/d/12jFzKaSLHHLhawFZwaF0M53c8pJGWzd1/view?usp=sharing

## Running the project
Running the project requires five components: MongoDB database, Server, PayPal Server, MTN Mobile Money (MoMo) Server and Clients. Make sure that the following ports are open on the machine you run it on: 8080, 8081, 8082 and 27017. See below for how to start the various components.

### Note that you need to start the database server and all other servers for the prototype to work properly

MongoDB database
- The quickest way to run a MongoDB for the project is to use a docker container running on port 0.0.0.0:27017.
- To create the MongoDB docker container, do the following:
  - Pull mongodb image - sudo docker pull mongo
  - Create container - sudo docker run -p 0.0.0.0:27017:27017/tcp --name cashtransfermongo -d mongo
  
Server
- Open a new terminal
- Change directory into /dev/bin directory
- Run command - ./server

PayPal Server
- Open a new terminal
- Change directory into /dev/bin directory
- Run command - ./paypalserver

MoMo Server
- Open a new terminal
- Change directory into /dev/bin directory
- Run command - ./momoserver

Client
You will need two or more clients to run to simulate the transfers. To run each client, do the following:
- Open a new terminal
- Change directory into /dev/bin directory
- Run command - ./client

To transfer money between clients the first time the database is created, you should make sure the following are done.
- After starting components and clients:  
  - Create PayPal and MoMo accounts for each of the clients
  - Create an account for the cash transfer app for each of the clients
