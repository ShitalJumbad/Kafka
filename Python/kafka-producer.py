from kafka import KafkaProducer
import time
import json

producer = KafkaProducer(bootstrap_servers='localhost:9092',value_serializer=lambda v: json.dumps(v).encode('utf-8'))
i=1
while True:
    keyId = "Id" + str(i)
    producer.send('python-test',key=keyId, value="value")
    i=i+10
    time.sleep(1)
