# Kafka 

A distributed streaming platform
Getting starting with Kafka

1)  sudo apt update
2)  Check if have a jdk-8 available
   1) java -version
   2) sudo apt install default-jre            
      sudo apt install openjdk-11-jre-headless
      sudo apt install openjdk-8-jre-headless 
3)  Download a scala : https://www.scala-lang.org/download/all.html
      1) sudo dpkg -i scala-2.11.4.deb
      2) scala -version
4)  wget http://www-us.apache.org/dist/kafka/2.2.1/kafka_2.11-2.2.1.tgz
5)  tar xzf kafka_2.11-2.2.1.tgz
6)  sudo mkdir /home/user/kafka
7)  sudo mv kafka_2.11-2.2.1 /home/user/kafka
8)  Inside kafka_2.11-2.2.1
    bin  config  libs  LICENSE  NOTICE  site-docs
    
    ### Setiing up a kafka cluster with a single zookeeper server in a single kafka broker (zookeeper instance and broker :- main components
    
## 1) Starting Zookaaeper
      (zookeeper needs configuration file to know how zookeeper should behave once started )
      
      bin/zookeeper-server-start.sh config/zookeeper.properties
      
      now waiting there to processes to connect to it...
      
      check $ telnet localhost 2181 ---> stat
        
## 2) Starting a broker 

     bin/kafka-server-start.sh config/server.properties
     
## 3) Creating a topic
     
     bin/kafka-topics.sh --create --topic my_topic --zookeeper localhost:2181 --replication-factor 1 --partitions 1
      
     we are specifying zookeeper servers because there could b emultiple zokeeper serviers managing there own clusters
     It is the zookeeper server who assinging broker to be responsible for the topic (index file and log files are created)
            
     ** list of the topics availble
     bin/kafka-topics.sh --list --zookeeper localhost:2181

## 4) Instatiate a producer 
   
     bin/kafka-console-producer.sh --broker-list localhost:9092 --topic mytopic

## 5) Instantiate a consumer 
  
     bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --from-beginning --topic mytopic
  
   * to start multiple brokers change the broker.id and listerners in property file of broker (server.properties)
      and log.dirs --> where kafka stores all the messages received
      
   * zookeeper.connect (zookeeper.properties)  where the broker can find the zookeeper to connect to 
 
 
 Messages trasnmitted to and from kafa known as records ( key(anything), value(anything), timestamp(timestamp format)) --> size 1 MB
 
 Partitions are used to split the load acrosss the broker (spread messages evenly). Replicated partitioning is used to ensure the integrity of your data if one of your broker goes down 
 
 Producer is an application which creates and transmits events to the kafka cluster. Serves as adapter for transmitting data (serialization, network connection, encryption --> low level details are handled)
