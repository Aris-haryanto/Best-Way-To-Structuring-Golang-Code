package api

// this is file for define your interface to adapter
// why we define in service
// like dave cheney said "the consumer should be define the interface"
// this service as a consumer

// for easy way to doing sql transact in abstraction
// we need to move it from service to api
// if you keep adapter.go in service that will be cycling import
// this is not break the rule

type IDB interface {
	SelectSomething(q *QuerySomething) (resp []ResponseSomething, err error)
	CreateSomething(ins *InsertSomething) (err error)
	WrapTx(fn func(IDB) error) error
}
