package types

type Message struct {
	Message string            `json:"message"`
	Routes  map[string]string `json:"routes"`
}
