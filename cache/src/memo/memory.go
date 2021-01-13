package memo

import "sync"

type Memory struct {		// 缓存类
	f Func		// 当缓存仓库没有key对应的数据时就会调用这个f方法从数据源获取数据
	cache map[string]*entry	// 缓存仓库results,这个仓库中存放多个entry对象，entry包括我们想获取的结果和一个key对应的channel
	repeat_lock *sync.Mutex		// 用于并发时，防止重复请求相同url的锁
}

// Func类型是一个函数类型，这种函数是一种能从数据源获取数据并存到Memory的函数
type Func func(key string) (value interface{}, err error)

type entry struct{
	res *result
	waiting_req chan struct{}
}

type result struct{		// 某个key对应的value
	value interface{}		// value可以是任何类型
	err error
}

// 实例化一个Memory对象
func New(getFromSource Func) *Memory{
	return &Memory{
		f: getFromSource,
		cache: map[string]*entry{},
		repeat_lock: &sync.Mutex{},
	}
}

// 获取一个key, 返回一个result
func (m *Memory) Get(key string) (value interface{}, err error){
	m.repeat_lock.Lock()
	ety, ok := m.cache[key]
	if !ok {		// 如果缓存仓库中不存在key数据，则从数据源获取
		ety = &entry{
			res: &result{},
			waiting_req: make(chan struct{}),	// 必须有这句哦，否则会隐式的为waiting_req成员设置一个零值channel。对零值channel接收会马上返回不会阻塞的
		}
		m.cache[key] = ety	// 第一个拿到锁并且检测到没有缓存的goroutine会在请求url前设置一个零值缓存，这样后面请求相同url的goroutine就会走到else的代码块中
		m.repeat_lock.Unlock()

		// 请求的过程会阻塞几百毫秒，此时其他请求相同的url也会被else代码块的channel阻塞
		m.cache[key].res.value, m.cache[key].res.err = m.f(key)
		close(ety.waiting_req)		// 广播给其他请求相同的url请求已结束
	}else{
		m.repeat_lock.Unlock()
		<-ety.waiting_req	// 走到这里会分两种情况：1.url已经请求过了，此时waiting_req在很早之前就已经被关闭了，因此不会阻塞；2.url正在被请求，waiting_req阻塞，直到请求结束，waiting_req被关闭，阻塞解除
	}

	return ety.res.value, ety.res.err
}