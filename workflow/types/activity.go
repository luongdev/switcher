package types

import "context"

type ActivityFunc func(context.Context, *ActivityInput) (*ActivityOutput, error)

type Activity interface {
	HandlerFunc() ActivityFunc
}
