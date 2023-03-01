# RabbitMQ consumer and sender

A simple docker container that will receive messages from a RabbitMQ queue and scale via KEDA.  The receiver will receive a single message at a time (per instance), and sleep for 1 second to simulate performing work.  When adding a massive amount of queue messages, KEDA will drive the container to scale out according to the event source (RabbitMQ).

