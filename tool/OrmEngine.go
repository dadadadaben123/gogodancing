package tool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
)

// 供全局使用
var DbOrm = new(Orm)

type Orm struct {
	*xorm.Engine
}

func OrmEngine(dbConfig *DatabaseConfig) (*Orm, error) {
	connStr := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DbName
	db, err := xorm.NewEngine("mysql", connStr)

	db.ShowSQL(dbConfig.ShowSql)

	err = db.Sync2(new(model.Member))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = db

	DbOrm.Engine = db

	return orm, nil
}
