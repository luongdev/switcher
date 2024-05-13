package commands

import (
	"fmt"
	"github.com/google/uuid"
)

type UId struct {
	uid          string
	allowMissing bool
}

func (c *UId) Validate() error {
	if !c.allowMissing {
		if len(c.uid) == 0 {
			return fmt.Errorf("uid is required")
		}
	}

	if len(c.uid) > 0 {
		if _, err := uuid.Parse(c.uid); err != nil {
			return fmt.Errorf("invalid uid: %v", err)
		}
	}

	return nil
}
