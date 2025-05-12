package models

type Countries struct {
	Id          uint    `json:"id" gorm:"primarykey;autoIncrement"`
	Country     *string `json:"country" gorm:"not null"`
	CountryCode *string `json:"countryCode" gorm:"not null"`
}
