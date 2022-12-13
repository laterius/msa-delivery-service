package mixtures

import (
	"github.com/ezn-go/mixture"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
)

func init() {
	mx := &gormigrate.Migration{
		ID:       "0003",
		Migrate:  mixture.CreateTableM(&service.Reservations{}),
		Rollback: mixture.DropTableR(&service.Reservations{}),
	}

	mixture.Add(mixture.ForAnyEnv, mx)
}
