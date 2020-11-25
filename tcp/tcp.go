package tcp

import (
	"net"
	"time"
)

// Detect checks if a tcp address is connectable with specific timeout
func Detect(addr string, timeout time.Duration) (bool, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if conn != nil {
		// Discard any unsent or unacknowledged data and ignore error
		_ = conn.(*net.TCPConn).SetLinger(0)
		// Close conn and ignore error
		_ = conn.Close()

		return true, nil
	}

	return false, err
}
