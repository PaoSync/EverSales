package models

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Rating struct {
	gorm.Model
	UserID uint `json:"user_id"`
	PropertyID uint `json:"property_id"`
	Rating int `json:"rating"`
	Comment string `json:"comment"`
}

func (rating *Rating)  Validate() (map[string]interface{}, bool) {
	if rating.Rating < 0 && rating.Rating > 5 {
		return u.Message(false, "Rating out of Range"), false
	}
	return u.Message(true, "success"), true
}

func (rating *Rating) Create() (map[string]interface{}) {
	if resp, ok := rating.Validate(); !ok {
		return resp
	}
	GetDB().Create(rating)
	var count int
	err := GetDB().Table("ratings").Where("property_id = ?",rating.PropertyID).Count(&count).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	property := &Property{}
	err = GetDB().Table("properties").Where("property_id = ?",rating.PropertyID).Find(&property).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if count == 0 {
		property.Rating = rating.Rating
	}else{
		property.Rating = (property.Rating*(count-1)+rating.Rating)/count
	}
	GetDB().Save(&property)
	resp := u.Message(true, "success")
	resp["rating"] = rating
	return resp
}
