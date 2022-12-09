package services

// this is file for define your interface to adapter
// why we define in service
// like dave cheney said "the consumer should be define the interface"
// this service as a consumer

import "github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/api"

type IDB interface {
	SelectSomething(q *api.QuerySomething) (resp []api.ResponseSomething, err error)
}
