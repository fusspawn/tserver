package controllers;


import (
	"github.com/revel/revel"
    "github.com/jinzhu/gorm"
    "github.com/fusspawn/tserver/app/models"
     _ "github.com/mattn/go-sqlite3"
)


var (
    Dbm *gorm.DB
)

type GorpController struct {
    *revel.Controller
}

func InitDB() {
    Dbm, err := gorm.Open("sqlite3", "gorm.db")
    if err != nil {
		panic(err)
	}

    Dbm.AutoMigrate(&models.EventMessage{})
}
