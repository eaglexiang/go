package tunnel

import (
	"net"
	"time"
)

// Flow 双向流动
func Flow(left, right net.Conn, timeout ...time.Duration) {
	t := GetTunnel()
	t.SetLeft(left)
	t.SetRight(right)

	if len(timeout) > 0 {
		t.SetTimeout(timeout[0])
	}

	t.Flow()
	t.Close()

	PutTunnel(t)
}
