package types

const DefaultClient = "default"

type ClientProvider interface {
	Get(k string) (Client, bool)
}
