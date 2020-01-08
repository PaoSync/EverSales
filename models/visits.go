package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "../utils"
)

type Visit struct {
	gorm.Model
	Status int `json:"status"`
	RequesterID uint `json:"requester_id"`
	PropertyID uint `json:"owner_id"`
	OwnerID uint `json:"owner_id"`
}

func RequestVisit(userID uint,propertyID uint)(map[string]interface{}){
	property := &Property{}
	err := GetDB().Table("properties").Where("id = ?", propertyID).Find(&property).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	visit := &Visit{}
	visit.OwnerID = property.UserId
	visit.Status = 0
	visit.RequesterID = userID
	visit.PropertyID = propertyID
	GetDB().Create(visit)
	resp:= u.Message(true, "success")
	resp["visit"] = visit
	return resp
}

func GetVisitsForUser(userID uint)([]*Visit){
	visits := make([]*Visit, 0)
	err := GetDB().Table("visits").Where("owner_id = ?",userID).Find(&visits).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return visits
}

func GetVisitsForProperty(propertyID uint)([]*Visit){
	visits := make([]*Visit, 0)
	err := GetDB().Table("visits").Where("property_id = ?",propertyID).Find(&visits).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return visits
}

func ApproveVisit(visitID uint)(map[string]interface{}){
	visit := &Visit{}
	err := GetDB().Table("visits").Where("id = ?", visitID).Find(&visit).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	visit.Status = 1
	GetDB().Save(&visit)
	resp := u.Message(true, "success")
	return resp
}

func DenyVisit(visitID uint)(map[string]interface{}){
	visit := &Visit{}
	err := GetDB().Table("visits").Where("id = ?", visitID).Find(&visit).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	visit.Status = 2
	GetDB().Save(&visit)
	resp := u.Message(true, "success")
	return resp
}