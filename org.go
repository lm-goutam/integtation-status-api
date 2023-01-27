package main

type Org struct {
	//gorm.Model
	OrgId   int    `json:"org_id"`
	OrgName string `json:"org_name"`
}
