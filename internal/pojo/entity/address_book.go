package entity

// AddressBook 地址簿实体类（对应 Java 版 AddressBook）。sex: '0'女 '1'男。
type AddressBook struct {
	ID           int64   `json:"id" db:"id"`
	UserID       int64   `json:"userId" db:"user_id"`
	Consignee    string  `json:"consignee" db:"consignee"`
	Phone        string  `json:"phone" db:"phone"`
	Sex          string  `json:"sex" db:"sex"`
	ProvinceCode *string `json:"provinceCode" db:"province_code"`
	ProvinceName *string `json:"provinceName" db:"province_name"`
	CityCode     *string `json:"cityCode" db:"city_code"`
	CityName     *string `json:"cityName" db:"city_name"`
	DistrictCode *string `json:"districtCode" db:"district_code"`
	DistrictName *string `json:"districtName" db:"district_name"`
	Detail       string  `json:"detail" db:"detail"`
	Label        *string `json:"label" db:"label"`
	IsDefault    int     `json:"isDefault" db:"is_default"`
}
