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

/*
	Package ipset is a library providing a wrapper to the iptables ipset user space utility
*/
package ipset

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

const (
	minMajorVersion = 6
)

var (
	ipsetPath              string
	ErrNotFound            = errors.New("ipset utility not found")
	ErrVersionNotSupported = errors.New("ipset utility version is not supported, requiring version >= 6.0")
)

var (
	execCommand  = exec.Command
	execLookPath = exec.LookPath
)

type IPSet interface {
	// List dumps header data and the entries for the set to an
	// *Info instance. The Resolve option can be used to force
	// action lookups(which may be slow).
	List(options ...Option) (*Info, error)

	// List dumps header data and the entries for the set to the
	// specific file. The Resolve option can be used to force
	// action lookups(which may be slow).
	ListToFile(filename string, options ...Option) error

	// Name returns the set's name
	Name() string

	// Rename the set's action and the new action must not exist.
	Rename(newName string) error

	// Add adds a given entry to the set. If the Exist option is
	// specified, ipset ignores the error if the entry already
	// added to the set.
	Add(entry string, options ...Option) error

	// Del deletes an entry from a set. If the Exist option is
	// specified and the entry is not in the set (maybe already
	// expired), then the command ignores the error.
	Del(entry string, options ...Option) error

	// Test tests whether an entry is in a set or not.
	Test(entry string) (bool, error)

	// Flush flushed all entries from the the set.
	Flush() error

	// Destroy removes the set from kernel.
	Destroy() error

	// Save dumps the set data to a io.Reader in a format that restore
	// can read.
	Save(options ...Option) (io.Reader, error)

	// SaveToFile dumps the set data to s specific file in a format
	// that restore can read.
	SaveToFile(filename string, options ...Option) error

	// Restore restores a saved session from io.Reader generated by
	// save.
	Restore(r io.Reader) error

	// RestoreFromFile restores a saved session from a specific file
	// generated by save.
	RestoreFromFile(filename string) error
}

// New create a set identified with setname and specified type.
// The type may require type specific options. If the Exist
// option is specified, ipset ignores the error when the same set
// (setname and create parameters are identical) already exists.
func New(name string, setType SetType, options ...Option) (IPSet, error) {
	c := getCmd(_create, name, setType, string(setType))
	defer putCmd(c)
	if err := c.exec(options...); err != nil {
		return nil, err
	}
	return &set{name, setType}, nil
}

// Flush all entries from the specified set or flush all sets if none
// is given.
func Flush(names ...string) error {
	if len(names) > 0 {
		for _, name := range names {
			if err := flush(name); err != nil {
				return err
			}
		}
	}
	return flushAll()
}

// flush flushes specific set
func flush(name string) error {
	if out, err := execCommand(ipsetPath, _flush, name).
		CombinedOutput(); err != nil {
		return fmt.Errorf("ipset: can't flush set %s: %s", name, out)
	}
	return nil
}

// flushAll flushes all set
func flushAll() error {
	if out, err := execCommand(ipsetPath, _flush).
		CombinedOutput(); err != nil {
		return fmt.Errorf("ipset: can't flush all set: %s", out)
	}
	return nil
}

// Destroy removes the specified set or all the sets if none is given.
// If the set has got reference(s), nothing is done and no set destroyed.
func Destroy(names ...string) error {
	if len(names) > 0 {
		for _, name := range names {
			if err := destroy(name); err != nil {
				return err
			}
		}
	}
	return destroyAll()
}

// destroy removes specific set
func destroy(name string) error {
	if out, err := execCommand(ipsetPath, _destroy, name).
		CombinedOutput(); err != nil {
		return fmt.Errorf("ipset: can't destroy set %s: %s", name, out)
	}
	return nil
}

// destroyAll removes all set
func destroyAll() error {
	if out, err := execCommand(ipsetPath, _destroy).
		CombinedOutput(); err != nil {
		return fmt.Errorf("ipset: can't destroy all set: %s", out)
	}
	return nil
}

// Swap swaps the content of two sets, or in another words,
// exchange the action of two sets. The referred sets must
// exist and compatible type of sets can be swapped only.
func Swap(from, to string) error {
	if out, err := execCommand(ipsetPath, _swap, from, to).
		CombinedOutput(); err != nil {
		return fmt.Errorf("ipset: can't swap from %s to %s: %s", from, to, out)
	}
	return nil
}

//Check checks whether there is an ipset command in the system.
// If so, check if the version is legal.
func Check() error {
	if ipsetPath != "" {
		return nil
	}

	path, err := execLookPath("ipset")
	if err != nil {
		return ErrNotFound
	}
	ipsetPath = path

	var supported bool
	if supported, err = isSupported(); err != nil {
		return fmt.Errorf("ipset: can't check version : %s", err)
	}

	if supported {
		return nil
	}
	return ErrVersionNotSupported
}

//// Refresh is used to to overwrite the set with the specified entries.
//// The ipset is updated on the fly by hot swapping it with a temporary set.
//func (s *IPSet) Refresh(entries ...string) (err error) {
//	if len(entries) == 0 {
//		return nil
//	}
//
//	tempName := s.action + "-temp"
//	_, err = s.createHashSet(tempName)
//	if err != nil {
//		return err
//	}
//
//	defer func() {
//		if e := Destroy(tempName); err == nil || e != nil {
//			err = e
//		}
//	}()
//
//	for _, entry := range entries {
//		out, err := exec.Command(ipsetPath, "add", tempName, entry, "-exist").CombinedOutput()
//		if err != nil {
//			return fmt.Errorf("ipset: can't add entry %s to set %s: %s", entry, tempName, out)
//		}
//	}
//
//	return Swap(tempName, s.action)
//}
//

func isSupported() (bool, error) {
	if out, err := execCommand(ipsetPath, _version).
		CombinedOutput(); err != nil {
		return false, err
	} else {
		return getMajorVersion(out) >= minMajorVersion, nil
	}
}

func getMajorVersion(version []byte) int {
	vIndex := bytes.IndexByte(version, 'v')
	dotIndex := bytes.IndexByte(version, '.')
	var majorVersion int
	for i := vIndex + 1; i < dotIndex; i++ {
		if c := version[i]; c >= '0' && c <= '9' {
			majorVersion = majorVersion*10 + int(c-'0')
		} else {
			return 0
		}
	}

	return majorVersion
}
