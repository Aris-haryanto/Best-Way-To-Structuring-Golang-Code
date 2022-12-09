package api

type QuerySomething struct {
	Name string
}
type ResponseSomething struct {
	Namanya string `json:"namanya" mapstructure:"name"`
}
