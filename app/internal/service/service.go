package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return Service{db: db}
}

type Reservations struct {
	CourierId   uuid.UUID `json:"courierId" gorm:"type:uuid; not null"`
	OrderId     uuid.UUID `json:"orderId" gorm:"type:uuid; not null"`
	Destination string    `json:"destination"`
}

type Courier struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid; unique; primary_key;"`
	Name string    `json:"name" gorm:"type:string;"`
}

//Реализация методов обращения в базу данных

func (s *Service) GetFreeCourier() (Courier, error) {

	var result Courier

	err := s.db.Raw("SELECT * FROM couriers WHERE id NOT IN (SELECT courier_id FROM reservations) LIMIT 1").Scan(&result).Error

	if err != nil {
		return Courier{}, err
	}

	return result, nil
}

func (s *Service) Reserve(courier Courier, orderId uuid.UUID) error {
	err := s.db.Create(Reservations{
		OrderId:   orderId,
		CourierId: courier.Id,
	}).Error
	return err
}

func (s *Service) CancelReservation(orderId uuid.UUID) error {
	return s.db.Delete(&Reservations{OrderId: orderId}).Error
}
