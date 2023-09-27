package services

import (
	"truck-unii-app/apps/api/repositories"
	"truck-unii-app/apps/myconfig/myvar"
)

type sTruck struct {
	rt repositories.RepoTruck
}

func NewServiceTruck(rt repositories.RepoTruck) ServiceTruck {
	return &sTruck{
		rt: rt,
	}
}

func (s *sTruck) TruckProfileGet(truckID int) (interface{}, string, error) {
	var failMsg string
	mydata, err := s.rt.TruckProfileGet(truckID)
	if err != nil {
		return nil, failMsg, err
	}
	if mydata.ID == 0 {
		failMsg = myvar.MsgNotFoundData
		return nil, failMsg, nil
	}
	return mydata, failMsg, nil
}
