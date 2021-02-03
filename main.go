package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

type OnceA struct {
	done uint32
}

func (o *OnceA) Do(f func()) {
	if !atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		return
	}
	f()
}

func main() {
	var once OnceA
	var conn net.Conn
	go func() {
		fun1 := func() {
			time.Sleep(5 * time.Second) //模拟初始化的速度很慢
			conn, _ = net.DialTimeout("tcp", "baidu.com:80", time.Second)
		}
		once.Do(fun1)
	}()
	time.Sleep(500 * time.Millisecond)
	fun2 := func() {
		fmt.Println("执行fun2")
		conn, _ = net.DialTimeout("tcp", "baidu.com:80", time.Second)
	}
	//再调用do已经检查到done为1了
	once.Do(fun2)
	_, err := conn.Write([]byte("\"GET / HTTP/1.1\\r\\nHost: baidu.com\\r\\n Accept: */*\\r\\n\\r\\n\""))
	if err != nil {
		fmt.Println("err:", err)
	}
}
