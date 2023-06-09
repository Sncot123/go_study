package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	//发送完数据要等leader和followers确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	//新选出一个partition
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//成功交付的消息将在success_chanel返回
	config.Producer.Return.Successes = true
	//创建一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,err:", err)
		return
	}
	fmt.Printf("pid:%v  offset:%v\n", pid, offset)
}
