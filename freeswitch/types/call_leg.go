package types

import (
	"fmt"
	"github.com/google/uuid"
)

type Leg struct {
	Uid      string
	Endpoint string
	Vars     map[string]interface{}
}

func (l *Leg) Valid() bool {
	if len(l.Endpoint) == 0 {
		return false
	}

	return true
}

func (l *Leg) DialString() string {
	if l.Uid != "" {
		if _, err := uuid.Parse(l.Uid); err != nil {
			l.Uid = ""
		}
	}
	if l.Vars == nil {
		l.Vars = make(map[string]interface{})
	}
	if l.Uid != "" {
		l.Vars["origination_uuid"] = l.Uid
	}

	vars := ""
	for k, v := range l.Vars {
		vars += fmt.Sprintf("%s=%v", k, v)
	}
	if vars != "" {
		vars = fmt.Sprintf("{%v}", vars)
	}

	return fmt.Sprintf("%v%v", vars, l.Endpoint)
}

func (l *Leg) String() string {
	return l.DialString()
}

func (l *Leg) GoString() string {
	return l.DialString()
}
