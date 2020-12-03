// Copyright 2020 Kiyon Lin All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package tcp is a library providing utility functions of tcp.
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
