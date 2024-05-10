package commands

import (
	"fmt"
	"github.com/google/uuid"
)

type UIdCommand struct {
	uid      string
	ignoreId bool
}

func (c *UIdCommand) Validate() error {
	if !c.ignoreId {
		if len(c.uid) == 0 {
			return fmt.Errorf("uid is required")
		}

		if _, err := uuid.Parse(c.uid); err != nil {
			return err
		}
	}

	return nil
}
