package main

import (
    "context"
    "fmt"
    "time"
)

func main(){
	tr := NewTracker() //创建一个埋点
	go tr.Run()  //埋点去工作

	_ = tr.Event(context.Backgroup(), "test1") // 发送应用事件给埋点
	_ = tr.Event(context.Backgroup(), "test2") // 发送应用事件给埋点
	_ = tr.Event(context.Backgroup(), "test3") // 发送应用事件给埋点

	ctx, cancel := context.WithDeadline(context.Backgroup(), time.Now().Add(2 * time.Second))  //定时器
	defer cancel()

	tr.Shutdown(ctx)  // 应用发送数据完毕,关闭通道ch;  并等待定时器设置的时间后,看在规定时间内,埋点是否给应用程序发送已经关闭的信号


}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10)  // 构造埋点, 缓存10,指针
	}
}

// Tracker knows how to track events for the application.
type Tracker struct {
	ch chan string  		// 埋点, 结构体,  将要被在子线程中使用的类型值.  自带传输数据的通道,可以在子线程的出生地(应用程序), 和子线程质检传输数据,数据传输方向,从应用程序发送到子线程. 主线程发送数据,子线程接受数据,并处理数据.
	stop chan struct{}	// 当子线程处理完毕数据后,发个信号给主线程,告诉主线程,数据处理完毕,它要结束了.
}

func (t *Tracker) Event(ctx context.context, data string) error {
	select {
	case t.ch <- data:   // 主线程向子线程发送数据
		return nil 
	case <- ctx.Done():	// 给发送行为进行时间控制,不能一直阻塞等待发送.
		return ctx.Err()
	}
}

func (t *Tracker)Run(){
	for data := range t.ch {  	// 子线程,一直从通道ch中获取数据,然后休息1秒.打印数据(处理数据)
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}

	t.stop <- struct{}{}		// 当子线程处理完毕数据,那就发送一个信号到stop通道   命令控制通道.

}


func (t *Tracker)Shutdown(ctx context.Context){
	close(t.ch)			// 主线程发送完毕,通知子线程可以关闭了.其实就是关闭      数据传输通道.
	select {
	case <- t.stop:			// 阻塞等待子线程发送数据处理完毕信号,知道子线程处理完毕了,可以关闭了.
	case <- ctx.Done():		// 设置阻塞的时间,不能一直等待数据阻塞.
	}
}

