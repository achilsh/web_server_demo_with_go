package subpub

import (
	"fmt"
	"time"

	"go-micro.dev/v4/broker"
)

type SubCallBack struct {
	handler   broker.Handler
	suberItem broker.Subscriber
}

// 定义内存级别的，进程内消息订阅和发布的消息队列。
type MemBroker struct {
	BHandler broker.Broker
	Suber    map[string]*SubCallBack
}

func NewMembroker() *MemBroker {
	r := &MemBroker{
		BHandler: broker.NewMemoryBroker(),
		Suber:    make(map[string]*SubCallBack),
	}
	r.BHandler.Connect()

	return r
}
func (m *MemBroker) UnSub(topic string) {
	if sub, ok := m.Suber[topic]; ok {
		sub.suberItem.Unsubscribe()
	}
}
func (m *MemBroker) UnAllSub() {
	for k := range m.Suber {
		m.UnSub(k)
	}
	m.BHandler.Disconnect()
}

func (m *MemBroker) Pub(topic string, data []byte, header map[string]string) {
	msg := &broker.Message{
		Header: header,
		Body:   []byte(data),
	}
	m.BHandler.Publish(topic, msg)
}

func (m *MemBroker) Sub(topic string, cb broker.Handler) {
	if _, ok := m.Suber[topic]; !ok {
		m.Suber[topic] = &SubCallBack{
			handler: cb,
		}
	}
	b, e := m.BHandler.Subscribe(topic, cb)
	if e != nil {
		fmt.Printf("sub fail on topic: %v, e: %v\n", topic, e)
		delete(m.Suber, topic)
		return
	}
	m.Suber[topic].suberItem = b
}

func PubSub() {
	d := NewMembroker()
	defer d.UnAllSub()

	d.Sub("test1", func(b broker.Event) error {
		fmt.Println("recv msg: ", string(b.Message().Body))
		return nil
	})
	d.Sub("test1", func(b broker.Event) error {
		fmt.Println("recv msg2: ", string(b.Message().Body))
		return nil
	})

	ch := make(chan bool)
	go func() {
		stop := false
		time.AfterFunc(10*time.Second, func() {
			fmt.Println("stop pub logic....")
			stop = true
			ch <- true
		})
		for {
			if !stop {
				d.Pub("test1", []byte("this is demo"), nil)
				time.Sleep(1 * time.Second)
			} else {
				fmt.Println("break pub logic")
				break
			}
		}
	}()

	<-ch

	fmt.Println("stop all sub and pub logic ")
	
}
