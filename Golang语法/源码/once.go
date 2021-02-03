package 源码

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(f func()) {
	// 错误实现,无法保证只执行一次
	// 并发情况下,A执行时间较慢同时初始化将done制为1，B想执行却无法获得资源
	//if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
	//	f()
	//}
	//正确,使用互斥锁,保证只有一个g初始化，同时双检查done是否是0
	if atomic.LoadUint32(&o.done) == 0 {
		o.doslow(f)
	}
}

func (o *Once) doslow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
