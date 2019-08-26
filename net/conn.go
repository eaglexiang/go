package net

import (
	"net"
	"strings"
)

// GetIPOfConnRemote 获取Conn的远端IP
func GetIPOfConnRemote(conn net.Conn) string {
	return strings.Split(conn.RemoteAddr().String(), ":")[0]
}
