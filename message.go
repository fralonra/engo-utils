package utils

type Message struct {
	Name string
}

func (m Message) Type() string {
	return m.Name
}
