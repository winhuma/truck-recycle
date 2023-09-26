package services

import "truck-unii-app/apps/api/repositories"

type sTruck struct {
	rt repositories.RepoTruck
}

func NewServiceTruck(rt repositories.RepoTruck) ServiceTruck {
	return &sTruck{
		rt: rt,
	}
}

func (s *sTruck) TruckProfileGet(truckID int) (interface{}, error) {

	return nil, nil
}
