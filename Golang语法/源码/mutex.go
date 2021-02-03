package 源码

import "sync/atomic"

const (
	mutexLocked = 1 << iota
)

//独占锁
type MyMutex struct {
	// 信号量
	sema uint32
	// 锁的状态,包括四部分:waitersCount，starving，woken，locked
	state uint32
}

func (loc *MyMutex) Lock() {
	if atomic.CompareAndSwapUint32(&loc.state, 0, mutexLocked) {
		// 获取锁成功
		return
	}
	// 进入自旋状态
	loc.slowlock()
}

// 大for循环，判断当前方法能否通过自旋来等待锁的释放
func (loc *MyMutex) slowlock() {

}
