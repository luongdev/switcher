package interfaces

type ClientStore interface {
	Get(k string) (Client, bool)
	Set(k string, v Client)
	Del(k string)
}
