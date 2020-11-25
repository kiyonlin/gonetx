package tcp

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Tcp_Detect(t *testing.T) {
	var (
		addr    = "127.0.0.1:65443"
		timeout = time.Millisecond * 10
		ok      bool
		err     error
	)

	ok, err = Detect(addr, timeout)

	assert.False(t, ok)
	assert.NotNil(t, err)

	var ln net.Listener
	ln, err = net.Listen("tcp", "127.0.0.1:65443")
	assert.Nil(t, err)
	defer func() { _ = ln.Close() }()

	ok, err = Detect(addr, timeout)
	assert.True(t, ok)
	assert.Nil(t, err)
}
