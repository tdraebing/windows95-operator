package controller

import (
	"win95-op/win95-operator/pkg/controller/win95"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, win95.Add)
}
