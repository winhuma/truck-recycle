package services

type ServiceTruck interface {
	TruckProfileGet(truckID int) (interface{}, string, error)
}
