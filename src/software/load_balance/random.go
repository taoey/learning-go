package load_balance

import "math/rand"

// 随机负载

type RandomBalance struct {
	curIndex int      // 当前服务器index
	rss      []string // 服务器列表
}

func (rb *RandomBalance) Add(params ...string) error {
	rb.rss = append(rb.rss, params...)
	return nil
}

func (rb *RandomBalance) Get(string) (string, error) {
	return rb.next(), nil
}

func (rb *RandomBalance) next() string {
	if len(rb.rss) == 0 {
		return ""
	}
	rb.curIndex = rand.Intn(len(rb.rss))
	return rb.rss[rb.curIndex]
}
