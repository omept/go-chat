# Go Microservices Realtime Finance Chat Application with RabbitMQ, Websockets, JWT, Mux, Gorm(MySQL), And React

The goal of this exercise is to create a browser-based chat application using Go. This application allows several users to talk in one or more chatrooms and also to get stock quotes from an API using a specific command.

## Requirement
The following tools are required to run the project
<ul>
<li> Go (Golang) </li> 
<li>RabbitMQ</li>
<li> MySQL</li>
<li> Docker(optional)</li>
</ul>

## Context
I use the poly repo ideology, hence, the project works in collaboration with two other repositories:

- go-stockbot : A decoupled service that calls an API using a "stock_code" as a parameter (<a href="​https://stooq.com/q/l/?s=aapl.us&f=sd2t2ohlcv&h&e=csv">​https://stooq.com/q/l/?s=aapl.us&f=sd2t2ohlcv&h&e=csv</a>​, here ​aapl.us is the stock_code​) and recieves a csv file. The  bot parses the received CSV file and then sends a message to the chatroom service using a message broker (RabbitMQ).  
  <b>Repo Link : https://github.com/ong-gtp/go-stockbot-rabbitmq.git</b>


- React client app : The frontend application that allows registered users to log in and talk with other users in a chatroom. It also allows users to post messages as commands into a chatroom with the following format `/stock=stock_code` <br />
  <b>Repo Link : https://github.com/ong-gtp/go-chat-react.git</b>

<br />

This service reads published stock bot requests to env.STKBT_RECEIVER_QUEUE queue on RabbitMQ and reads the processed stock response on env.STKBT_PUBLISHER_QUEUE queue on RabbitMQ.

<br />

## <u>Starting The App</u>

<br />

## Step 1 : Start RabbitMQ
RabbitMQ is the message broker between the bot service and the chat service. To run it with docker, run the folling command in your terminal: <br />
### `docker run -d --hostname rabbitmq-svc --name rbbtmq -p 15672:15672 -p 5672:5672 rabbitmq:3.11.3-management`
<br />

## Step 2 : Update Env file
 copy the `.env.example` to `.env` and update the entries. 
<br />

## Step 3 : Starting the app
### run the folling command in your terminal `go run cmd/main.go`
<br />

## Step 4 : Start Bot service app
### visit this repo: `https://github.com/ong-gtp/go-stockbot-rabbitmq.git` on steps for stating the app

<br />

## Step 5 : Start React frontend app
### visit this repo `https://github.com/ong-gtp/go-chat-react.git` on steps for stating the react frontend app

<br />

## <u>REST Endpoints</u>
```bash
The table below describes the endpoints available on the app:

#### Routes ⚡
| Routes                     | HTTP Methods | Params                         | Description                                      |
| :------------------------- | :----------- | :----------------------------- | :----------------------------------------------- |
| /v1/api/auth/signup        | POST         | `email` `password` `user_name` | Creates a new user and returns jwt session token |
| /v1/api/auth/login         | POST         | `email` `password`             | Logs in a user and returns the jwt session token |
| /v1/api/chat/create        | POST         | `name`                         | Creates a new chat room with the name provided   |
| /v1/api/chat/rooms         | POST         | none                           | returns a list of chat rooms                     |
| /v1/api/chat/room-messages | POST         | `roomId`                       | Returns the latest 50 messages in a chat room    |
| /v1/ws                     | GET          | none                           | websocket connection url                         |
```

## App tests
Run the command below to execute tests
### `go test github.com/ong-gtp/go-chat/pkg/services`


## Useful Links
- https://blog.questionable.services/article/guide-logging-middleware-go/
- https://gorm.io
- https://stackoverflow.com/questions/47637308/create-unit-test-for-ws-in-golang
