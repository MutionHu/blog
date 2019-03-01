package models

import (
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	//引入sqlite3的驱动
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	db *gorm.DB
}

func (db *DB) Begin() {
	db.db = db.db.Begin()
}

func (db *DB) Rollback() {
	db.db = db.db.Rollback()
}

func (db *DB) Commit() {
	db.db = db.db.Commit()
}

var (
	db *gorm.DB
)

func NewDB() *DB {
	return &DB{db: db}
}

func init() {
	var err error
	if err := os.MkdirAll("data", 0777); err != nil {
		panic("failed to connect database," + err.Error())
	}
	if err = initDB(); err != nil {
		panic("failed to connect database," + err.Error())
	}
	db.SetLogger(logs.GetLogger("orm"))
	db.LogMode(true)
	//自动同步表
	db.AutoMigrate(&User{}, &Note{}, &Message{}, &PraiseLog{})
	var count int

	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		//新增
		db.Create(&User{
			Name:   "admin",
			Email:  "532889507@qq.com",
			Pwd:    "123456@x",
			Avatar: "/static/images/info-img.png",
			Role:   0,
		})
	}
}

func initDB() error {
	var err error
	dbconf, err := beego.AppConfig.GetSection("database")
	if err != nil {
		logs.Error(err)
		dbconf = map[string]string{
			"type": "sqlite3",
		}
	}
	switch dbconf["type"] {
	case "mysql":
		db, err = gorm.Open("mysql", dbconf["url"])
	default:
		db, err = gorm.Open("sqlite3", "data/data.db")
	}
	if err != nil {
		return err
	}
	return nil
}

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createtime"`
	UpdatedAt time.Time  `json:"updatetime"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (db *DB) GetDBTime() *time.Time {
	var t *time.Time
	row, err := db.db.DB().Query("select NOW()")
	if err != nil {
		logs.Error(err)
		return nil
	}
	row.Scan(t)
	return t
}
