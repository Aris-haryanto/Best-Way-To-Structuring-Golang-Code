package services

import (
	"errors"
	"fmt"

	"github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/api"
)

// this file is where you put all bussines logic seperate by file if you have more than one service
// i.e auth/middleware/calculation and many more

type Hello struct {
	db api.IDB
}

// here's you should define on main.go to set what DB you use
func (hello *Hello) SetDB(db api.IDB) {
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

// here's request struct used for a function below
type requestCreate struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (hello *Hello) failTx() (err error) {
	return errors.New("this transaction will fail")
}

// example function with db transaction
func (hello *Hello) CreateHello(reqCreate *requestCreate) (responseHello, error) {
	err := hello.db.WrapTx(func(db api.IDB) error {
		errCreateOne := db.CreateSomething(&api.InsertSomething{ID: reqCreate.ID, Name: reqCreate.Name})

		if errCreateOne != nil {
			// if error this transaction will fail and rollback
			return errCreateOne
		}

		errCreateTwo := db.CreateSomething(&api.InsertSomething{ID: reqCreate.ID, Name: reqCreate.Name})

		if errCreateTwo != nil {
			// if error this transaction will fail and rollback
			return errCreateTwo
		}

		// this tx will fail if this outside function return error
		if errFromOther := hello.failTx(); errFromOther != nil {
			return errFromOther
		}

		// if success this transaction will auto commit
		return nil
	})

	if err != nil {
		return responseHello{Code: 500, Message: "failed " + err.Error(), Data: []byte{}}, err
	}
	return responseHello{Code: 200, Message: "hello " + reqCreate.Name + " with baseUrl : " + config.BaseUrl, Data: nil}, nil
}
