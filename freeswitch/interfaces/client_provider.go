package interfaces

const DefaultClient = "default"

type ClientProvider interface {
	Get(k string) Client
}
