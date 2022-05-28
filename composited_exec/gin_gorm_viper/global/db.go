package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	DBLink *gorm.DB
)

func SetupDBLink() error {
	var err error
	DBLink, err = gorm.Open(DatabaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			DatabaseSetting.UserName,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.DBName,
			DatabaseSetting.Charset,
			DatabaseSetting.ParseTime,
		))
	//gorm.Open("mysql", "root:root_mysql@tcp(127.0.0.1:3307)/gorm_exec_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	if ServerSetting.RunMode == "debug" {
		DBLink.LogMode(true)
	}
	DBLink.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DBLink.DB().SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	DBLink.DB().SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	//otgorm.AddGormCallbacks(db)
	return nil
}
