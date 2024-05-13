package pkg

import (
	"github.com/luongdev/switcher/workflow/internal"
	"github.com/luongdev/switcher/workflow/types"
)

func NewRegistry() types.Registry {
	return internal.NewRegistry()
}
