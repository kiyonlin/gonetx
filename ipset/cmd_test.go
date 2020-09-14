package ipset

import (
	"testing"

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
