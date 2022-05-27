package Config

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
