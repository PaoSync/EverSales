package models

import (
	u "../utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"strconv"
)

type Property struct {
	gorm.Model
	UserId uint `json:"user_id"`
	RoomSize float32 `json:"room_size"`
	Rooms	int `json:"rooms"`
	Bathrooms int `json:"bathrooms"`
	Parking int `json:"parking"`
	Floors int `json:"floors"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Pets bool `json:"pets"`
	PropertyType int `json:"property_type"`
	Price float32 `json:"price"`
	Rating int `json:"rating"`
	PrivateSecurity bool `json:"private_security"`
	Capacity int `json:"capacity"`
	Photos pq.StringArray `json:"photos" gorm:"type:varchar(100)[]"`
	Active bool `json:"active"`
}
func (property *Property) Validate() (map[string]interface{}, bool) {

	if property.RoomSize < 0 {
		return u.Message(false, "Room size cant be negative"), false
	}

	if property.Bathrooms < 0 {
		return u.Message(false, "Number of bathrooms can't be negative"), false
	}

	if property.Parking < 0 {
		return u.Message(false, "Number of parking can't be negative"), false
	}

	if property.Floors < 0 {
		return u.Message(false, "Price can't be negative"), false
	}

	if property.Price < 0 {
		return u.Message(false, "Number of rooms can't be negative"), false
	}

	if property.Capacity < 0 {
		return u.Message(false, "Number of occupants can't be negative"), false
	}

	if property.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (property *Property) Create() (map[string]interface{}) {

	if resp, ok := property.Validate(); !ok {
		return resp
	}
	property.Active = true

	GetDB().Create(property)

	resp := u.Message(true, "success")
	resp["contact"] = property
	return resp
}

func GetActiveProperties() ([]*Property){
	properties := make([]*Property, 0)
	err := GetDB().Table("properties").Where("active = true").Find(&properties).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return properties
}

func DeleteProperty(propertyID uint,userID uint) (map[string]interface{}){
	property := &Property{}
	err := GetDB().Table("properties").Where("id = ?", propertyID).Find(&property).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if property.UserId != userID {
		return u.Message(false, "User doesnt own the property")
	}
	if property.Active == true {
		return u.Message(false, "Can't delete property while active")
	}
	GetDB().Delete(&property)
	resp := u.Message(true, "Successfully deleted property")
	resp["Property"] = property
	return resp
}

func TogglePropertyStatus(propertyID uint,userID uint) (map[string]interface{}){
	property := &Property{}
	err := GetDB().Table("properties").Where("id = ?", propertyID).Find(&property).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if property.UserId != userID {
		return u.Message(false, "User doesnt own the property")
	}
	property.Active = !property.Active
	GetDB().Save(&property)
	resp := u.Message(true, "status Successfully toggled")
	resp["Property"] = property
	return resp
}

func GetPropertyByID(propertyID uint)(map[string]interface{}){
	property := &Property{}
	err := GetDB().Table("properties").Where("id = ?", propertyID).Find(&property).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resp := u.Message(true, "status Successfully toggled")
	resp["Property"] = property
	return resp
}

func (property *Property)  Modify(userID uint)(map[string]interface{}){
	propertyOld := &Property{}
	err := GetDB().Table("properties").Where("id = ?", property.ID).Find(&propertyOld).Error
	fmt.Println(userID)
	fmt.Println(propertyOld.UserId)
	if userID!=propertyOld.UserId {
		return u.Message(false, "User doesnt own the property")
	}
	property.UserId = userID
	if resp, ok := property.Validate(); !ok {
		fmt.Println(property.ID)
		return resp
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	propertyOld = property
	GetDB().Save(&propertyOld)
	resp := u.Message(true, "Property successfully modified")
	resp["Property"] = property
	return resp
}

func Search(minSize string,maxSize string,rooms string,bathrooms string,parking string,floors string,pets string,propertyType string,maxPrice string,minPrice string,privateSecurity string,capacity string)([]*Property){
	tx := GetDB().Table("properties")
	minSizeC, err := strconv.ParseFloat(minSize, 32)
	maxSizeC, err2 := strconv.ParseFloat(maxSize, 32)
	if minSize != "" && err==nil {
		if maxSize != "" && err2==nil {
			tx = tx.Where("room_size <= ?",maxSizeC).Where("room_size >= ?",minSizeC)
		}else{
			tx = tx.Where("room_size >= ?",minSizeC)
		}
	}else{
		if maxSize != "" && err2==nil {
			tx = tx.Where("room_size <= ?",maxSizeC)
		}
	}
	var roomsC int
	roomsC,err = strconv.Atoi(rooms)
	if rooms != "" && err== nil {
		tx = tx.Where("rooms = ?",roomsC)
	}
	var bathroomsC int
	bathroomsC,err = strconv.Atoi(bathrooms)
	if bathrooms != "" && err== nil {
		tx = tx.Where("bathrooms = ?",bathroomsC)
	}
	var parkingC int
	parkingC,err = strconv.Atoi(parking)
	if parking != "" && err== nil {
		tx = tx.Where("parking = ?",parkingC)
	}
	var floorsC int
	floorsC,err = strconv.Atoi(floors)
	if floors != "" && err== nil {
		tx = tx.Where("floors = ?",floorsC)
	}
	var petsC bool
	petsC,err = strconv.ParseBool(pets)
	if pets != "" && err== nil {
		tx = tx.Where("pets = ?",petsC)
	}
	var propertyTypeC int
	propertyTypeC,err = strconv.Atoi(propertyType)
	if propertyType != "" && err== nil {
		tx = tx.Where("property_type = ?",propertyTypeC)
	}
	var minPriceC,maxPriceC float64
	minPriceC, err = strconv.ParseFloat(minPrice, 32)
	maxPriceC, err2 = strconv.ParseFloat(maxPrice, 32)
	if minPrice != "" && err==nil {
		if maxPrice != "" && err2==nil {
			tx = tx.Where("price <= ?",maxPriceC).Where("price >= ?",minPriceC)
		}else{
			tx = tx.Where("price >= ?",minPriceC)
		}
	}else{
		if maxPrice != "" && err2==nil {
			tx = tx.Where("price <= ?",maxPriceC)
		}
	}
	var privateSecurityC bool
	privateSecurityC,err = strconv.ParseBool(privateSecurity)
	if privateSecurity != "" && err== nil {
		tx = tx.Where("private_security = ?",privateSecurityC)
	}
	var capacityC int
	capacityC,err = strconv.Atoi(capacity)
	if capacity != "" && err== nil {
		tx = tx.Where("capacity = ?",capacityC)
	}
	properties := make([]*Property, 0)
	err3 := tx.Find(&properties).Error
	if err3 != nil {
		fmt.Println(err3)
		return nil
	}
	return properties
}