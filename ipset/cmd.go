package ipset

import (
	"fmt"
	"sync"
)

// Commands
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

// Options
const (
	_exist   = "-exist"
	_resolve = "-resolve"
)

type cmd struct {
	action  string
	name    string
	entry   string
	setType SetType
	out     []byte
}

func (c *cmd) buildArgs(opts ...Option) (args []string) {
	args = append(args, c.action, c.name)
	if !c.isTwoArgs() {
		args = append(args, c.entry)
	}

	return c.appendArgs(args, opts...)
}

func (c *cmd) appendArgs(args []string, opts ...Option) []string {
	o := getOptions().apply(opts...)
	defer optionsPool.Put(o)

	if !o.disableExist && c.needExist() {
		args = append(args, _exist)
	} else {
		o.disableExist = false
	}

	if o.resolve && c.needResolve() {
		args = append(args, _resolve)
		o.resolve = false
	}

	return args
}

func (c *cmd) exec(opts ...Option) error {
	out, err := execCommand(ipsetPath, c.buildArgs(opts...)...).
		CombinedOutput()

	if err != nil {
		if c.isTwoArgs() {
			return fmt.Errorf("ipset: can't %s %s: %s", c.action, c.name, out)
		}

		return fmt.Errorf("ipset: can't %s %s %s: %s", c.action, c.name, c.entry, out)
	}

	if c.needResolve() {
		c.out = out
	}

	return nil
}

func (c *cmd) isTwoArgs() bool {
	return c.action == _list || c.action == _save ||
		c.action == _destroy || c.action == _flush
}

func (c *cmd) needResolve() bool {
	return c.action == _list || c.action == _save
}

func (c *cmd) needExist() bool {
	return c.action == _create || c.action == _add || c.action == _del
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
