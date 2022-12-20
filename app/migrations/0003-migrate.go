package mixtures

import (
	"github.com/ezn-go/mixture"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
)

func init() {
	mx := &gormigrate.Migration{
		ID:       "0007",
		Migrate:  mixture.CreateTableM(&service.Courierreservations{}),
		Rollback: mixture.DropTableR(&service.Courierreservations{}),
	}

	mixture.Add(mixture.ForAnyEnv, mx)
}
