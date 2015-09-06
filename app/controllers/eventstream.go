package controllers;


import (
	"github.com/revel/revel"
	"github.com/fusspawn/tserver/app/models"
)

type EventStream struct {
	*GormController
}

func (c EventStream) NewForm() revel.Result {
	return c.Render()
}

func (c EventStream) CreateEvent(event_type, event_json string) revel.Result {
	ins := &models.EventMessage{EType:event_type,  Data:event_json,  Handled:false}
	//c.Txn.NewRecord(ins)
	//c.Txn.Create(ins)
	return c.RenderJson(ins)
}
