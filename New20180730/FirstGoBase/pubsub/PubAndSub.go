package pubsub

import (
	"fmt"
	"sync"
	"time"
)
//构建了一个名为pubsub的发布订阅模型支持包
func init() {
	fmt.Println("pubsub 包下的开始运行了")
}

type (
	//订阅者为一个管道
	sub chan interface{}
	// 主题为一个过滤器
	topicFunc func(v interface{}) bool

)
//发布者 对象
type Publisher struct {
	/*
	RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。RWMutex可以创建为其他结构体的字段；零值为解锁状态。RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。
	 */
	m sync.RWMutex
	//订阅队列的缓存的大小
	buffer int
	// 发布超时的时间
	timeout time.Duration
	//订阅者的信息
	subscribers map[sub]topicFunc

}
//  构建一个发布者对象，可以设置发布的超时的时间和缓存队列的的长度
func NewPublisher(timeout time.Duration,buffer int) *Publisher  {

	return &Publisher{
		buffer:buffer,
		timeout:timeout,
		subscribers:make(map[sub]topicFunc),
	}
}


//添加一个新的订阅者，订阅全部的主题
func (p *Publisher) SubscibeAll() chan interface{} {
	//return p.subscribers(nil)
    return p.SubscribeTopic(nil)
}

// 添加一个新的订阅者，订阅过滤器筛选后的主题
func  (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{}  {
	ch:= make(chan interface{},p.buffer)
	p.m.Lock()
	p.subscribers[ch]=topic
	p.m.Unlock()
	return ch
}

//退出订阅者
func (p *Publisher) Evict(sub chan interface{}){
	p.m.Lock()
	defer p.m.Unlock()
    //删除map中的元素
	delete(p.subscribers,sub)
	close(sub)
}
//发布一个主题
func (p *Publisher) Publish(v interface{}){
	//RLock方法将rw锁定为读取状态，禁止其他线程写入，但不禁止读取。
	p.m.RLock()
	defer p.m.RUnlock()

	var   wg sync.WaitGroup
	for subs,topic := range p.subscribers{
		wg.Add(1)
		//发送主题
		go p.sendTopic(subs, topic, v, &wg)
	}
    wg.Wait()
}

// 关闭发布者的对象，同时关闭所有的订阅者管道
func (p *Publisher) Close()  {
	p.m.Lock()
	defer p.m.Unlock()

	for subs := range p.subscribers{
		delete(p.subscribers,subs)
		close(subs)
	}
}

//发送主题，可以容忍一定的超时
func (p *Publisher) sendTopic(sendsub sub,topicFunc2 topicFunc,v interface{},wg *sync.WaitGroup){
	defer wg.Done()
	if topicFunc2!=nil&& !topicFunc2(v) {
      return
	}

	select {
	case sendsub<- v:
	case <-time.After(p.timeout):

	}

}




