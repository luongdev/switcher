package enums

type Direction string

const (
	Inbound  Direction = "inbound"
	Outbound Direction = "outbound"
	Internal Direction = "internal"
	Unknown  Direction = "unknown"
)

func (d Direction) String() string {
	return string(d)
}

func Parse(s string) Direction {
	switch s {
	case Inbound.String():
		return Inbound
	case Outbound.String():
		return Outbound
	case Internal.String():
		return Internal
	}

	return Unknown
}
