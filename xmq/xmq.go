package xmq

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

var connection *amqp.Connection

func Init(url string) error {
	if connection == nil || connection.IsClosed() {
		mq, err := amqp.Dial(url)
		if err != nil {
			return err
		}
		connection = mq
	}

	return nil
}

type XMQ struct {
	Connection *amqp.Connection
	queues     map[string]*Worker
}

func New(url string) (*XMQ, error) {
	if err := Init(url); err != nil {
		return nil, err
	}

	return &XMQ{
		Connection: connection,
		queues:     map[string]*Worker{},
	}, nil
}

func (xmq *XMQ) Push(queue string, data interface{}) (string, error) {
	ch, err := xmq.Connection.Channel()
	if err != nil {
		return "", err
	}
	defer ch.Close()
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	msgID := uuid.New().String()

	return msgID, ch.Publish("", queue, false, false, amqp.Publishing{
		MessageId:   msgID,
		ContentType: "text/plain",
		Body:        dataBytes,
	})
}

func (xmq *XMQ) Register(queue string, consumer Consumer, options ...WorkerOption) {
	w := &Worker{
		queue:    queue,
		consumer: consumer,
		number:   1,
	}
	for _, o := range options {
		o(w)
	}
	xmq.queues[queue] = w
}

func (xmq *XMQ) Start(url string) error {
	mq, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	connection = mq
	for _, worker := range xmq.queues {
		go worker.Start()
	}

	return nil
}

type Consumer interface {
	Do(messageID string, messageBody []byte) error
}

type ConsumerDo func(messageID string, messageBody []byte) error

type Worker struct {
	queue    string
	consumer Consumer
	number   int
}

func (w *Worker) Start() {
	ch, err := connection.Channel()
	if err != nil {
		fmt.Println("获取channel错误" + err.Error())
		return
	}
	defer ch.Close()

	_ = ch.Qos(w.number, 0, false)

	msgs, err := ch.Consume(w.queue, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("开始获取消息错误" + err.Error())
		return
	}

	for msg := range msgs {
		msg := msg
		go func() {
			if err = w.consumer.Do(msg.MessageId, msg.Body); err != nil {
				fmt.Println("执行任务错误:" + err.Error())
				if err = msg.Nack(true, false); err != nil {
					fmt.Println("nack确认消息错误:" + err.Error())
				}
				return
			}
			if err = msg.Ack(false); err != nil {
				fmt.Println("ack确认消息错误:" + err.Error())
			}
		}()
	}
}

type WorkerOption func(*Worker)

func SetNumber(num int) WorkerOption {
	return func(w *Worker) {
		w.number = num
	}
}
