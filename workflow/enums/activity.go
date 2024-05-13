package enums

type Activity string

func (a Activity) String() string {
	return string(a)
}

const (
	ActivityHttp Activity = "http"
)
