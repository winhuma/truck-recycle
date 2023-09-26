package repositories

type Repositories struct {
}

type RepoTruck interface {
	TruckProfileGet(truckID int) (interface{}, error)
}
