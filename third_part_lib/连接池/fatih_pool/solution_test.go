package fatih_pool

import (
	"fmt"
	"github.com/fatih/pool"
	"net"
	"testing"
)

func TestSolution(t *testing.T) {
	factory := func() (net.Conn, error) { return net.Dial("tcp", "127.0.0.1:4000") }

	p, _ := pool.NewChannelPool(5, 30, factory)
	conn, _ := p.Get()
	conn.Close()
	if pc, ok := conn.(*pool.PoolConn); ok {
		pc.MarkUnusable()
		pc.Close()
	}
	p.Close()
	current := p.Len()
	fmt.Println(current)
}
