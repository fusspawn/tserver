package controllers;


import (
	"github.com/revel/revel"
    "github.com/jinzhu/gorm"
    "github.com/fusspawn/tserver/app/models"
     _ "github.com/mattn/go-sqlite3"
	"database/sql"
)


var (
    Dbm *gorm.DB
)

type GormController struct {
    *revel.Controller	
    Txn *gorm.DB
}

func InitDB() {
    Dbm, err := gorm.Open("sqlite3", "gorm.db")
	Dbm.LogMode(true)

    if err != nil {
		panic(err)
	}

    Dbm.AutoMigrate(&models.EventMessage{})
	
    //revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
    //revel.InterceptMethod((*GormController).Commit, revel.AFTER)
    //revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)
}


func (c *GormController) Begin() revel.Result {
    txn := Dbm.Begin()
    if txn.Error != nil {
        panic(txn.Error)
    }
    c.Txn = txn
    return nil
}

// This method clears the c.Txn after each transaction
func (c *GormController) Commit() revel.Result {
    if c.Txn == nil {
        return nil
    }
    c.Txn.Commit()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

// This method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() revel.Result {
    if c.Txn == nil {
        return nil
    }
    c.Txn.Rollback()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

