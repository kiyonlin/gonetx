package ipset

import (
	"fmt"
	"sync"
)

const (
	_create  = "create"
	_add     = "add"
	_del     = "del"
	_test    = "test"
	_destroy = "destroy"
	_list    = "list"
	_save    = "save"
	_restore = "restore"
	_flush   = "flush"
	_rename  = "rename"
	_swap    = "swap"
	_version = "version"
)

type cmd struct {
	action  string
	name    string
	entry   string
	setType SetType
	out     []byte
}

func (c *cmd) buildArgs(opts ...Option) (args []string) {
	if c.action == _create {
		args = append(args, c.action, c.name, string(c.setType))
	} else if c.action == _list {
		args = append(args, c.action, c.name)
	} else {
		args = append(args, c.action, c.name, c.entry)
	}

	o := getOptions().apply(opts...)
	defer optionsPool.Put(o)

	if !o.disableExist {
		args = append(args, "-exist")
		o.disableExist = false
	}

	if (c.action == _list || c.action == _save) && o.resolve {
		args = append(args, "-resolve")
		o.resolve = false
	}

	return
}

func (c *cmd) exec(opts ...Option) error {
	out, err := execCommand(ipsetPath, c.buildArgs(opts...)...).
		CombinedOutput()

	if err != nil {
		if c.action == _list || c.action == _save || c.action == _destroy {
			return fmt.Errorf("ipset: can't %s %s: %s", c.action, c.name, out)
		}

		target := c.entry
		if target == "" {
			target = string(c.setType)
		}

		return fmt.Errorf("ipset: can't %s %s %s: %s", c.action, c.name, target, out)
	}

	if c.action == _list || c.action == _save {
		c.out = out
	}

	return nil
}

var cmdPool = sync.Pool{
	New: func() interface{} {
		return &cmd{}
	},
}

func getCmd(action, name string, setType SetType, entry ...string) *cmd {
	c := cmdPool.Get().(*cmd)
	c.action = action
	c.name = name
	c.setType = setType
	if len(entry) > 0 {
		c.entry = entry[0]
	}
	return c
}

func putCmd(c *cmd) {
	c.out = nil
	cmdPool.Put(c)
}
