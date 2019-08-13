## Implementing a Kafka producer consumer using Golang
sarama-cluster, a balanced consumer implementation built on top the existing sarama client library by Shopify

First Install using :-  go get github.com/Shopify/sarama

1) bin/zookeeper-server-start.sh config/zookeeper.properties 
2) bin/kafka-server-start.sh config/server.properties 
3) bin/kafka-topics.sh --create --topic newest --zookeeper localhost:2181 --replication-factor 1 --partitions 1
4) go run consumer.go
5) go run producer.go




























references :-  [Golang: Implementing kafka consumers using sarama-cluster](https://dev.to/davidsbond/golang-implementing-kafka-consumers-using-sarama-cluster-4fko)
