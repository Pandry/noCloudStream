package handlers

import (
	"auther/api"
)

type backpack struct {
	api.Backpack
}

// NewHttpBackpack returns a backpack containing the repository (database), the struct for operating signatures and
func NewHttpBackpack(b api.Backpack) backpack {
	return backpack{b}
}
