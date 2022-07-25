// ��������https://lailin.xyz/post/observer.html
package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// Bus Bus
type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// AsyncEventBus �첽�¼�����
type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

// NewAsyncEventBus new
func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

// Subscribe ����
func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := bus.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handlers[topic] = handler

	return nil
}

// Publish ����
// �����첽ִ�У����Ҳ���ȴ����ؽ��
func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}

func sub1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("sub1, %s %s\n", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	fmt.Printf("sub2, %s %s\n", msg1, msg2)
}

func main() {
	bus := NewAsyncEventBus()
	bus.Subscribe("topic:1", sub1)
	bus.Subscribe("topic:1", sub2)
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)
}

//https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648470771&idx=1&sn=9ec1e90ae3e38762b6172ea8424c1823&chksm=f247529cc530db8a9696914995478a8f28b9e561d5a030cd22323e0b8afba436d665718cf134&scene=132#wechat_redirect
