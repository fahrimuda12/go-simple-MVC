package main

// func main(){
// 	config := &kafka.ConfigMap{
// 		"bootstrap.servers": "localhost:9092",
// 	}

// 	producer, err := kafka.NewProducer(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer producer.Close()

// 	topic := "helloworld"

// 	for i := 0; i < 10; i++ {
// 		fmt.Println("Send message to kafka %d \n", i)

// 		key := strconv.Itoa(i)
// 		value := fmt.Sprintf("Hello Go! %d", i)

// 		message := &kafka.Message{
// 			TopicPartition: kafka.TopicPartition{
// 				Topic: &topic,
// 				Partition: kafka.PartitionAny
// 			},
// 			Key: []byte(key),
// 			Value: []byte(value),
// 		}

// 		err := producer.Produce(message, nil)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	producer.Flush(15 * 1000)
// }