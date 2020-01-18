package tunnel

import "net"

// Flow 双向流动
func Flow(left, right net.Conn) {
	t := GetTunnel()
	t.SetLeft(left)
	t.SetRight(right)
	t.Flow()
	t.Close()
	PutTunnel(t)
}
