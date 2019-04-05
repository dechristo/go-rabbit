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

### 2.2 Worker

Implements an worker example that demonstrates how to deal with long running tasks mixed with short running ones.

The quantity of **dots** in the user input is the length of time that the task will run.

Steps:

1 - Open three console windows.

2 - On one start the worker: `go run worker.go` (Inside the worker folder)

3 - On the other two start the publisher with a long task and a short one (Inside the publisher folder):

- `go run publisher.go yoda............................`
- `go run publisher.go luke..`

Note how the worker does not need to wait for the longer task to finish to process the shorter one:

**Publisher**

    λ go run publisher\publisher.go yoda............................
    2019/04/05 13:10:57  [--->] Sent yoda............................

    λ go run publisher\publisher.go luke..
    2019/04/05 13:11:06  [--->] Sent luke..

**Worker 1**

    2019/04/05 13:10:47  [*] Waiting for messages. To exit press CTRL+C
    2019/04/05 13:10:57 Processing...
    2019/04/05 13:11:22 Done!

**Worker 2**

    2019/04/05 13:10:54  [*] Waiting for messages. To exit press CTRL+C
    2019/04/05 13:11:06 Processing...
    2019/04/05 13:11:08 Done!


