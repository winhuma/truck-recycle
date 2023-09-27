package repositories

import "truck-unii-app/apps/myconfig/mymodels"

type Repositories struct {
}

type RepoTruck interface {
	TruckProfileGet(truckID int) (mymodels.DBTruckProfile, error)
}
