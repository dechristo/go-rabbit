# Go Rabbit
Examples using the blasing fast Go language with RabbitMQ

## 1. Install the dependencies
`go get -d ./...`

## 2. Examples

### 2.1 Basic
The basic example demonstrates concurrent message queueing.

Steps:

1 - Open two console windows.

2 - On one start the consumer: `go run consumer.go` (Inside the consumer folder)

3 - On the other start the publisher: `go run publisher.go` (Inside the publisher folder).

Output example:

**Publisher**
    
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [+++++] Sent That's why you fail.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.
    2019/04/05 10:38:22  [-----] Sent Try not. Do or do not. There is no try.

**Consumer**

    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.
    2019/04/05 10:38:22 Received a message: That's why you fail.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.
    2019/04/05 10:38:22 Received a message: Try not. Do or do not. There is no try.