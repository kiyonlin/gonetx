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

	testSetTypes = []SetType{
		BitmapIp,
		BitmapIpMac,
		BitmapPort,
		HashIp,
		HashMac,
		HashIpMac,
		HashNet,
		HashNetNet,
		HashIpPort,
		HashNetPort,
		HashIpPortIp,
		HashIpPortNet,
		HashIpMark,
		HashNetPortNet,
		HashNetIface,
		ListSet,
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

func Test_Options_Packets(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without packets", func(t *testing.T) {
			args := c.appendArgs(nil, Packets(0))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need packets", func(t *testing.T) {
				args := c.appendArgs(nil, Packets(1))
				assert.Equal(t, _packets, args[0])
				assert.Equal(t, "1", args[1])
			})
		} else {
			t.Run(action+" ignore packets", func(t *testing.T) {
				args := c.appendArgs(nil, Packets(1))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Bytes(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without bytes", func(t *testing.T) {
			args := c.appendArgs(nil, Bytes(0))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need bytes", func(t *testing.T) {
				args := c.appendArgs(nil, Bytes(1))
				assert.Equal(t, _bytes, args[0])
				assert.Equal(t, "1", args[1])
			})
		} else {
			t.Run(action+" ignore bytes", func(t *testing.T) {
				args := c.appendArgs(nil, Bytes(1))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Comment(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without comment", func(t *testing.T) {
			args := c.appendArgs(nil, Comment(false))
			assert.Len(t, args, 0)
		})

		if c.onlyCreate() {
			t.Run(action+" need comment", func(t *testing.T) {
				args := c.appendArgs(nil, Comment(true))
				assert.Equal(t, _comment, args[0])
			})
		} else {
			t.Run(action+" ignore comment", func(t *testing.T) {
				args := c.appendArgs(nil, Comment(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_CommentContent(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without comment content", func(t *testing.T) {
			args := c.appendArgs(nil, CommentContent(""))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need comment content", func(t *testing.T) {
				args := c.appendArgs(nil, CommentContent("comment"))
				assert.Equal(t, _comment, args[0])
				assert.Equal(t, "comment", args[1])
			})
		} else {
			t.Run(action+" ignore comment content", func(t *testing.T) {
				args := c.appendArgs(nil, CommentContent("comment"))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Skbinfo(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without skbinfo", func(t *testing.T) {
			args := c.appendArgs(nil, Skbinfo(false))
			assert.Len(t, args, 0)
		})

		if c.onlyCreate() {
			t.Run(action+" need skbinfo", func(t *testing.T) {
				args := c.appendArgs(nil, Skbinfo(true))
				assert.Equal(t, _skbinfo, args[0])
			})
		} else {
			t.Run(action+" ignore skbinfo", func(t *testing.T) {
				args := c.appendArgs(nil, Skbinfo(true))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Skbmark(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without skbmark", func(t *testing.T) {
			args := c.appendArgs(nil, Skbmark(""))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need skbmark", func(t *testing.T) {
				args := c.appendArgs(nil, Skbmark("skbmark"))
				assert.Equal(t, _skbmark, args[0])
				assert.Equal(t, "skbmark", args[1])
			})
		} else {
			t.Run(action+" ignore skbmark", func(t *testing.T) {
				args := c.appendArgs(nil, Skbmark("skbmark"))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Skbprio(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without skbprio", func(t *testing.T) {
			args := c.appendArgs(nil, Skbprio(""))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need skbprio", func(t *testing.T) {
				args := c.appendArgs(nil, Skbprio("skbprio"))
				assert.Equal(t, _skbprio, args[0])
				assert.Equal(t, "skbprio", args[1])
			})
		} else {
			t.Run(action+" ignore skbprio", func(t *testing.T) {
				args := c.appendArgs(nil, Skbprio("skbprio"))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Skbqueue(t *testing.T) {
	for _, action := range testActions {
		c := getFakeCmd(action)
		t.Run(action+" without skbqueue", func(t *testing.T) {
			args := c.appendArgs(nil, Skbqueue(0))
			assert.Len(t, args, 0)
		})

		if c.onlyAdd() {
			t.Run(action+" need skbqueue", func(t *testing.T) {
				args := c.appendArgs(nil, Skbqueue(1))
				assert.Equal(t, _skbqueue, args[0])
				assert.Equal(t, "1", args[1])
			})
		} else {
			t.Run(action+" ignore skbqueue", func(t *testing.T) {
				args := c.appendArgs(nil, Skbqueue(1))
				assert.Len(t, args, 0)
			})
		}
	}
}

func Test_Options_Nomatch(t *testing.T) {
	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" without nomatch", func(t *testing.T) {
				args := c.appendArgs(nil, Nomatch(false))
				assert.Len(t, args, 0)
			})

			if c.needNomatch() {
				t.Run(action+" need nomatch", func(t *testing.T) {
					args := c.appendArgs(nil, Nomatch(true))
					assert.Equal(t, _nomatch, args[0])
				})
			} else {
				t.Run(action+" ignore nomatch", func(t *testing.T) {
					args := c.appendArgs(nil, Nomatch(true))
					assert.Len(t, args, 0)
				})
			}
		}
	}
}

func Test_Options_Family(t *testing.T) {
	for _, action := range testActions {
		for _, setType := range testSetTypes {
			c := getFakeCmd(action, setType)
			t.Run(action+" without family", func(t *testing.T) {
				args := c.appendArgs(nil, Family(""))
				assert.Len(t, args, 0)
			})

			if c.needFamily() {
				t.Run(action+" need family", func(t *testing.T) {
					args := c.appendArgs(nil, Family("inet"))
					assert.Equal(t, _family, args[0])
					assert.Equal(t, "inet", args[1])
				})
			} else {
				t.Run(action+" ignore family", func(t *testing.T) {
					args := c.appendArgs(nil, Family("inet"))
					assert.Len(t, args, 0)
				})
			}
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
