package resource

import (
	"io"
	"os"
)

////////////////////////////////////////////////////////////////////////////////

// Exist returns true if the resource exists.
func Exist(name string) bool {
	for _, s := range sources {
		if s.exist(name) {
			return true
		}
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////

// Open returns a ReadCloser corresponding to the resource. If the resource does
// not exists, it returns an error.
func Open(name string) (io.ReadCloser, error) {
	for _, s := range sources {
		f, err := s.open(name)
		if !os.IsNotExist(err) {
			return f, err
		}
	}
	return nil, os.ErrNotExist
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).