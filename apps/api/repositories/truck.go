package repositories

import (
	"truck-unii-app/apps/myconfig/mymodels"
	"truck-unii-app/pkg/myfunc"

	"gorm.io/gorm"
)

type rTruck struct {
	DB *gorm.DB
}

func NewRepoTruck(db *gorm.DB) RepoTruck {
	return &rTruck{
		DB: db,
	}
}

func (r *rTruck) TruckProfileGet(truckID int) (mymodels.DBTruckProfile, error) {
	var data mymodels.DBTruckProfile
	if err := r.DB.Table(mymodels.DBTruckProfile.TableName(mymodels.DBTruckProfile{})).
		Where("id=?", truckID).
		Scan(&data).Error; err != nil {
		return data, myfunc.MyErrFormat(err)
	}
	return data, nil
}
