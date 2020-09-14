package ipset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Options_Exist(t *testing.T) {
	tt := []struct {
		action       string
		disableExist bool
		containExist bool
	}{
		{_create, false, true},
		{_create, true, false},
		{_add, false, true},
		{_add, true, false},
		{_del, false, true},
		{_del, true, false},
	}

	for _, tc := range tt {
		c := getFakeCmd(tc.action)
		var args []string
		if tc.disableExist {
			args = c.appendArgs(nil, DisableExist(true))
		} else {
			args = c.appendArgs(nil)
		}
		if tc.containExist {
			assert.Len(t, args, 1,
				"action %s should contain %s option when disable exist is %t",
				tc.action, _exist, tc.disableExist)
			assert.Equal(t, _exist, args[0])
		} else {
			assert.Len(t, args, 0,
				"action %s should not contain %s option when disable exist is %t",
				tc.action, _exist, tc.disableExist)
		}
	}
}

func Test_Options_Exist_Ignore(t *testing.T) {
	actions := []string{
		_test,
		_destroy,
		_list,
		_save,
		_restore,
		_flush,
		_rename,
		_swap,
	}

	for _, action := range actions {
		c := getFakeCmd(action)
		args := c.appendArgs(nil, DisableExist(true))
		assert.Len(t, args, 0,
			"action %s should not contain %s option when disable exist is %t",
			action, _exist, true)

		args = c.appendArgs(nil, DisableExist(false))
		assert.Len(t, args, 0,
			"action %s should not contain %s option when disable exist is %t",
			action, _exist, false)
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
