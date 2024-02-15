package tool

import (
	"fmt"
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
	if err != nil {
		return nil, err
	}
	db.ShowSQL(dbConfig.ShowSql)

	err = db.Sync2(new(model.Member), new(model.FoodCategory), new(model.Shop), new(model.Service), new(model.ShopService), new(model.Goods))

	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = db

	DbOrm.Engine = db
	//
	//InitShopData()
	//InitGoodsData()
	return orm, nil
}

func InitShopData() {

	shops := []model.Shop{
		model.Shop{Id: 1, Name: "张亮麻辣烫", Address: "上海市闵行区吴中路21栋201", Phone: "1585476789", Longitude: 5.43, Latitude: 10.33},
		model.Shop{Id: 2, Name: "上海本帮菜", Address: "上海市闵行区吴中路222栋201", Phone: "1585376789", Longitude: 6.43, Latitude: 8.33},
		model.Shop{Id: 12, Name: "东北菜", Address: "上海市闵行区万源路21栋201", Phone: "1585476389", Longitude: 20.43, Latitude: 80.33},
	}

	//创建一个事务
	session := DbOrm.NewSession()
	//defer关闭
	defer session.Close()
	//开启
	err := session.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	//插入
	for _, shop := range shops {
		_, err := session.Insert(&shop)
		if err != nil {
			fmt.Println(err)
			session.Rollback()
			return
		}
	}
	//提交
	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}

}

func InitGoodsData() {

	goods := []model.Goods{
		model.Goods{Name: "麻辣烫", Description: "番茄锅底，不辣", SellCount: 100, Price: 30, OldPrice: 40, ShopId: 1},
		model.Goods{Name: "番茄炒蛋", Description: "正宗本帮菜", SellCount: 2100, Price: 20, OldPrice: 30, ShopId: 2},
	}

	//创建一个事务
	session := DbOrm.NewSession()
	//defer关闭
	defer session.Close()
	//开启
	err := session.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	//插入
	for _, good := range goods {
		_, err := session.Insert(&good)
		if err != nil {
			fmt.Println(err)
			session.Rollback()
			return
		}
	}
	//提交
	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}

}
