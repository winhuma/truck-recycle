package mymodels

type (
	DBTruckProfile struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CarDetail string `json:"car_detail"`
	}
)

func (DBTruckProfile) TableName() string {
	return "truck_profile"
}
