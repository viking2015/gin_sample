package conf

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DB() gorm.DB {
	sqlConnection := DB_USER_NAME + ":" + DB_PASSWORD + "@tcp(127.0.0.1:3306)/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
	}

	return db
}

const (
	DefaultKey  = "conf/db"
	errorFormat = "[gorm] ERROR! %s\n"
)

func DBInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlConnection := "manatini_dba:123456@tcp(127.0.0.1:3306)/manatini_db?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open("mysql", sqlConnection)
		if err != nil {
			panic(err)
		}
		c.Set(DefaultKey, db)
		c.Next()
	}
}

// shortcut to get DB
func DefaultDB(c *gin.Context) gorm.DB {
	return c.MustGet(DefaultKey).(gorm.DB)
}