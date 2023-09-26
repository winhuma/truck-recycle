package repositories

import "gorm.io/gorm"

type rTruck struct {
	DB *gorm.DB
}

func NewRepoTruck(db *gorm.DB) RepoTruck {
	return &rTruck{
		DB: db,
	}
}

func (r *rTruck) TruckProfileGet(truckID int) (interface{}, error) {
	return nil, nil
}
