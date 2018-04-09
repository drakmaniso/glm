// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

import (
	"github.com/drakmaniso/glam/internal"
)

type Float uint32

const noFloat = Float(maxID)

func NewFloat(name string) Float {
	_, ok := actions[name]
	if ok {
		//TODO: set error
		return noFloat
	}

	l := len(internal.Floats.Name)
	if l >= maxID {
		//TODO: set error
		return noFloat
	}

	actions[name] = Float(l)
	internal.Floats.Name = append(internal.Floats.Name, name)
	internal.Floats.Active = append(internal.Floats.Active, false)
	internal.Floats.Value = append(internal.Floats.Value, 0)

	return Float(l)
}

func (f Float) Name() string {
	return internal.Bools.Name[f]
}

func (f Float) Active() bool {
	return internal.Floats.Active[f]
}

func (f Float) Value() float32 {
	return internal.Floats.Value[f]
}