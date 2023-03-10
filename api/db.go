package api

type QuerySomething struct {
	Name string
}
type ResponseSomething struct {
	Namanya string `json:"namanya" mapstructure:"name"`
}

type InsertSomething struct {
	ID   int64
	Name string
}
