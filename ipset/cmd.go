package ipset

import (
	"fmt"
	"strconv"
	"strings"
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
	_exist    = "-exist"
	_resolve  = "-resolve"
	_timeout  = "timeout"
	_counters = "counters"
	_packets  = "packets"
	_bytes    = "bytes"
	_comment  = "comment"
	_skbinfo  = "skbinfo"
	_skbmark  = "skbmark"
	_skbprio  = "skbprio"
	_skbqueue = "skbqueue"
	_nomatch  = "nomatch"
	_family   = "family"
	_hashsize = "hashsize"
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
	o := acquireOptions().apply(opts...)
	defer releaseOptions(o)

	if o.timeout > 0 && c.needTimeout() {
		args = append(args, _timeout, i2str(int64(o.timeout.Seconds())))
	}

	if o.exist && c.needExist() {
		args = append(args, _exist)
	}

	if o.resolve && c.needResolve() {
		args = append(args, _resolve)
	}

	if o.counters && c.needCounters() {
		args = append(args, _counters)
	}

	if o.countersPackets > 0 && c.onlyAdd() {
		args = append(args, _packets, i2str(int64(o.countersPackets)))
	}

	if o.countersBytes > 0 && c.onlyAdd() {
		args = append(args, _bytes, i2str(int64(o.countersBytes)))
	}

	if o.comment && c.onlyCreate() {
		args = append(args, _comment)
	}

	if o.commentContent != "" && c.onlyAdd() {
		args = append(args, _comment, o.commentContent)
	}

	if o.skbinfo && c.onlyCreate() {
		args = append(args, _skbinfo)
	}

	if o.skbmark != "" && c.onlyAdd() {
		args = append(args, _skbmark, o.skbmark)
	}

	if o.skbprio != "" && c.onlyAdd() {
		args = append(args, _skbprio, o.skbprio)
	}

	if o.skbqueue != 0 && c.onlyAdd() {
		args = append(args, _skbqueue, i2str(int64(o.skbqueue)))
	}

	if o.nomatch && c.needNomatch() {
		args = append(args, _nomatch)
	}

	if o.family != "" && c.needFamily() {
		args = append(args, _family, string(o.family))
	}

	if o.hashSize != 0 && c.needHashSize() {
		args = append(args, _hashsize, i2str(int64(o.hashSize)))
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

func (c *cmd) needExist() bool {
	return c.action == _create || c.action == _add || c.action == _del
}

func (c *cmd) needTimeout() bool {
	return c.action == _create || c.action == _add
}

func (c *cmd) needResolve() bool {
	return c.action == _list || c.action == _save
}

func (c *cmd) needCounters() bool {
	return c.action == _create
}

func (c *cmd) onlyAdd() bool {
	return c.action == _add
}

func (c *cmd) onlyCreate() bool {
	return c.action == _create
}

func (c *cmd) needNomatch() bool {
	return c.action == _add &&
		(c.setType == HashNet || c.setType == HashNetNet ||
			c.setType == HashNetPort || c.setType == HashIpPortNet ||
			c.setType == HashNetPortNet || c.setType == HashNetIface)
}

func (c *cmd) needFamily() bool {
	return c.action == _create && c.setType != HashMac && strings.HasPrefix(string(c.setType), "hash")
}

func (c *cmd) needHashSize() bool {
	return c.action == _create && strings.HasPrefix(string(c.setType), "hash")
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

func i2str(i int64) string {
	return strconv.FormatInt(i, 10)
}
