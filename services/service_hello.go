package services

import (
	"fmt"

	"github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/api"
)

// this file is where you put all bussines logic seperate by file if you have more than one service
// i.e auth/middleware/calculation and many more

type Hello struct {
	db IDB
}

// here's you should define on main.go to set what DB you use
func (hello *Hello) SetDB(db IDB) {
	hello.db = db
}

// =====================

// here's request struct used for a function below
type requestHello struct {
	Name string `json:"name"`
}

// here's response struct used for a function below
type responseHello struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// this the function who using struct above
// remember you should define struct above the function
// that's make your code easy to read and search
func (hello *Hello) SaidHello(reqHello *requestHello) (responseHello, error) {
	getResult, err := hello.db.SelectSomething(&api.QuerySomething{Name: reqHello.Name})
	fmt.Println(getResult)
	if err != nil {
		return responseHello{Code: 500, Message: "failed " + err.Error(), Data: []byte{}}, err
	}
	return responseHello{Code: 200, Message: "hello " + reqHello.Name + " with baseUrl : " + config.BaseUrl, Data: getResult}, nil
}
