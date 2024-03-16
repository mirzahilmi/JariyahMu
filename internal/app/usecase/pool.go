package usecase

import (
	"sync"

	"github.com/go-sql-driver/mysql"
)

var mysqlErrPool = sync.Pool{
	New: func() any {
		err := mysql.MySQLError{}
		return &err
	},
}
