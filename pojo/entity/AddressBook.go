package entity

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AddressBook 地址簿
type AddressBook struct {
	ID           int64          `json:"id"`
	UserID       int64          `json:"userId"`
	Consignee    string         `json:"consignee"`
	Phone        string         `json:"phone"`
	Sex          string         `json:"sex"` // 0 女 1 男
	ProvinceCode string         `json:"provinceCode"`
	ProvinceName string         `json:"provinceName"`
	CityCode     string         `json:"cityCode"`
	CityName     string         `json:"cityName"`
	DistrictCode string         `json:"districtCode"`
	DistrictName string         `json:"districtName"`
	Detail       string         `json:"detail"`
	Label        FlexibleString `json:"label"`
	IsDefault    int            `json:"isDefault"` // 0否 1是
}

type FlexibleString string

func (f *FlexibleString) UnmarshalJSON(data []byte) error {
	if data[0] == '"' && data[len(data)-1] == '"' {
		*f = FlexibleString(strings.Trim(string(data), "\""))
		return nil
	}
	var num float64
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}
	*f = FlexibleString(fmt.Sprintf("%v", num))
	return nil
}
