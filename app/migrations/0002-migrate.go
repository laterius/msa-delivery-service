package mixtures

import (
	"github.com/ezn-go/mixture"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
)

func (c Courier) TableName() string {
	return "couriers"
}

type Courier struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid; unique; primary_key;"`
	Name string    `json:"name" gorm:"type:string;"`
}

func init() {

	couriers := []Courier{
		{Id: uuid.New(), Name: "courier 1"},
		{Id: uuid.New(), Name: "courier 2"},
		{Id: uuid.New(), Name: "courier 3"},
	}

	mx := &gormigrate.Migration{
		ID:       "0006",
		Migrate:  mixture.CreateBatchM(couriers),
		Rollback: mixture.DeleteBatchR(couriers),
	}

	mixture.Add(mixture.ForAnyEnv, mx)
}
