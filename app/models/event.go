package models

import (
	"github.com/revel/revel"
    "github.com/jinzhu/gorm"
)

type EventMessage struct {
  gorm.Model;
	Handled bool;
	EType string;
	Data string;
}

func (e *EventMessage) Validate(v *revel.Validation) {
	v.Required(e.EType);
	v.Required(e.Data);
}
