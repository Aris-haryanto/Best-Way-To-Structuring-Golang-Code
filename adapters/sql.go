package adapters

import (
	"log"
	"time"

	"github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/api"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Hello struct {
	Name string `gorm:"column:name"`
	ID   int64  `gorm:"primaryKey"`
}

type SqlDB struct {
	SqlConn *gorm.DB
}

func SqlConnection(username string, password string, host string, port string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/hello_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, connErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if connErr != nil {
		log.Fatalln(connErr)
	}

	sqlDB, poolErr := db.DB()

	if poolErr != nil {
		log.Fatalln(connErr)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.Debug().AutoMigrate(&Hello{})

	return db
}

func (sqldb *SqlDB) CloseSql() {
	c, _ := sqldb.SqlConn.DB()
	c.Close()
}

func (sqldb *SqlDB) SelectSomething(q *api.QuerySomething) (resp []api.ResponseSomething, err error) {
	response := []map[string]interface{}{}
	errSql := sqldb.SqlConn.Model(&Hello{}).Where("name = ?", q.Name).Find(&response)
	if errSql.Error != nil {
		return resp, errSql.Error
	}

	// you need to decode to api struct make it not dependency with the adapter
	// so you can change with other DB like a redis, elasticsearch, mongodb
	errDecodeToStruct := mapstructure.Decode(response, &resp)
	if errDecodeToStruct != nil {
		return resp, errDecodeToStruct
	}

	return resp, nil
}

func (sqldb *SqlDB) CreateSomething(ins *api.InsertSomething) (err error) {
	data := Hello{
		ID:   ins.ID,
		Name: ins.Name,
	}
	tx := sqldb.SqlConn.Create(data)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// ===================

type SqlDBTx struct {
	api.IDB
}

func TxConn(db *gorm.DB) *SqlDBTx {
	setToAdapter := &SqlDB{SqlConn: db}
	return &SqlDBTx{setToAdapter}
}

func (sqldb *SqlDB) WrapTx(fn func(api.IDB) error) error {

	// start transaction
	tx := sqldb.SqlConn.Begin()

	// triger rollback when panic recover
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// re-initiate SqlDB struct to SqlDBTx struct only for this transaction
	q := TxConn(tx)
	// set struct SqlDB to closure function
	err := fn(q)

	// check if closure function error
	if err != nil {
		// check error when rollback
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr.Error
		}

		return err
	}

	// return commit error any of nil or not
	return tx.Commit().Error

}
