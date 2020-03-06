package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "go-contacts/utils"
)

type Document struct {
	gorm.Model
	Number string `json:"number"`
	Phone string `json:"phone"`
	Description string `json:"description"`
	UserId uint   `json:"user_id"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (document *Document) Validate() (map[string]interface{}, bool) {

	if document.Number == "" {
		return u.Message(false, "Document number should be on the payload"), false
	}

	if document.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (document *Document) Create() (map[string]interface{}) {

	if resp, ok := document.Validate(); !ok {
		return resp
	}

	GetDB().Create(document)

	resp := u.Message(true, "success")
	resp["document"] = document
	return resp
}

func GetDocument(id uint) (*Document) {

	document := &Document{}
	err := GetDB().Table("documents").Where("id = ?", id).First(document).Error
	if err != nil {
		return nil
	}
	return document
}

func UpdateDocument(id uint, document *Document) (*Document) {

	err := GetDB().Table("documents").Where("id = ?", id).Updates(document).Error
	if err != nil {
		return nil
	} 
	return document
}

func DeleteDocument(id uint, document *Document) (*Document) {

	err := GetDB().Table("documents").Where("id = ?", id).Delete(document).Error
	if err != nil {
		return nil
	} 
	return document
}

func GetDocuments(user uint) ([]*Document) {

	documents := make([]*Document, 0)
	err := GetDB().Table("documents").Where("user_id = ?", user).Find(&documents).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return documents
}
