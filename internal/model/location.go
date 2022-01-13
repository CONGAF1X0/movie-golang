package model

type UserLocation struct {
	Model
	UID         uint64 `json:"uid"`
	CurrNation string `json:"curr_nation"`
	CurrProvince string `json:"curr_province"`
	CurrCity	string `json:"curr_city"`
	CurrDistrict string `json:"curr_district"`
	Location string `json:"location"`
	Longitude float64 `json:"longitude"` 	//经度
	Latitude float64 `json:"latitude"`		//纬度
}
