package repository

import (
	"context"

	"github.com/rushbuilding/airlineseatmap/model/entity"
	"gorm.io/gorm"
)

type PassengerRepositoryImpl struct {
	db *gorm.DB
}

func NewPassengerRepository(db *gorm.DB) PassengerRepository {
	return &PassengerRepositoryImpl{db:db}
}

func (r *PassengerRepositoryImpl) Insert(ctx context.Context, passenger *entity.Passenger) (int64, error) {
	response := r.db.WithContext(ctx).Create(passenger);
	return response.RowsAffected, response.Error
}

func (r *PassengerRepositoryImpl) FindById(ctx context.Context, id string) entity.Passenger {
	var passenger entity.Passenger
	err := r.db.WithContext(ctx).Where("passengers.passenger_id = ?", id).
				Preload("EmailAddresses").
				Preload("PhoneNumbers").
				Joins("Document").
				First(&passenger).Error

	if err != nil {
		panic(err)
	}

	return passenger
}