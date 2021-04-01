package simple_pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// 连接池
// 功能点：
// 连接池容量，最大连接数，连接存活时间，获取连接，释放连接，关闭连接池

type Pool struct {
	m       sync.Mutex                // 保证多个goroutine访问时候，closed的线程安全
	res     chan io.Closer            //连接存储的chan
	factory func() (io.Closer, error) //新建连接的工厂方法
	closed  bool                      //连接池关闭标志
}

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了。")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

//从资源池里获取一个资源
func (p *Pool) Get() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, errors.New("获取失败")
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

//关闭资源池，释放资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	//关闭通道，不让写入了
	close(p.res)
	//关闭通道里的资源
	for r := range p.res {
		r.Close()
	}
}
func (p *Pool) Release(r io.Closer) {
	//保证该操作和Close方法的操作是安全的
	p.m.Lock()
	defer p.m.Unlock()
	//资源池都关闭了，就省这一个没有释放的资源了，释放即可
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了，释放这个资源吧")
		r.Close()
	}
}
