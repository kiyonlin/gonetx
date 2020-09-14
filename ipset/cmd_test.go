package ipset

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testActions = []string{
		_create,
		_add,
		_del,
		_test,
		_destroy,
		_list,
		_save,
		_restore,
		_flush,
		_rename,
		_swap,
	}
)

func Test_Options_Timeout(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without timeout", func(t *testing.T) {
			args := c.appendArgs(nil, Timeout(0))
			assert.Len(t, args, 0)
		})

		if c.needTimeout() {
			t.Run(action+" need timeout", func(t *testing.T) {
				args := c.appendArgs(nil, Timeout(time.Second))
				assert.Equal(t, _timeout, args[0])
				assert.Equal(t, "1", args[1])
			})
		} else {
			t.Run(action+" ignore timeout", func(t *testing.T) {
				args := c.appendArgs(nil, Timeout(time.Second))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Exist(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without exist", func(t *testing.T) {
			args := c.appendArgs(nil, Exist(false))
			assert.Len(t, args, 0)
		})

		if c.needExist() {
			t.Run(action+" need exist", func(t *testing.T) {
				args := c.appendArgs(nil, Exist(true))
				assert.Equal(t, _exist, args[0])
			})
		} else {
			t.Run(action+" ignore exist", func(t *testing.T) {
				args := c.appendArgs(nil, Exist(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Resolve(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without resolve", func(t *testing.T) {
			args := c.appendArgs(nil, Resolve(false))
			assert.Len(t, args, 0)
		})

		if c.needResolve() {
			t.Run(action+" need resolve", func(t *testing.T) {
				args := c.appendArgs(nil, Resolve(true))
				assert.Equal(t, _resolve, args[0])
			})
		} else {
			t.Run(action+" ignore resolve", func(t *testing.T) {
				args := c.appendArgs(nil, Resolve(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Counters(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without counters", func(t *testing.T) {
			args := c.appendArgs(nil, Counters(false))
			assert.Len(t, args, 0)
		})

		if c.needCounters() {
			t.Run(action+" need counters", func(t *testing.T) {
				args := c.appendArgs(nil, Counters(true))
				assert.Equal(t, _counters, args[0])
			})
		} else {
			t.Run(action+" ignore counters", func(t *testing.T) {
				args := c.appendArgs(nil, Counters(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func getFakeCmd(action string, setType ...SetType) *cmd {
	st := HashIp
	if len(setType) > 0 {
		st = setType[0]
	}
	return &cmd{
		action:  action,
		name:    "test",
		setType: st,
	}
}
