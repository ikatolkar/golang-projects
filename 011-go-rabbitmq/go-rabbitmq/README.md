# go-rabbitmq

Simple rabbitmq producer consumer program.

RabbitMQ is run as a docker container, launched using docker-compose.

Sender sends a message over a queue, while receiver receives it.


```bash
$ docker-compose up

root@ik-ubuntu-mantic:~/ikatolkar/rabbitmq/go-rabbitmq# go run receive/receive.go
2024/03/27 04:55:01 [*] Waiting for messages, To exit press CTRL + Z
2024/03/27 04:55:12 Received a message: Hello world!

root@ik-ubuntu-mantic:~/ikatolkar/rabbitmq/go-rabbitmq# go run send/send.go
2024/03/27 04:55:12 [x] sent Hello world!
```
