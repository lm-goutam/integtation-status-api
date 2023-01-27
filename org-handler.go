package main

import (
	"encoding/json"
	"net/http"
)

func CreateOrg(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var org Org
	json.NewDecoder(r.Body).Decode(&org)
	Database.Create(&org)
	json.NewEncoder(w).Encode(org)
}

/*func GetOrg(w http.ResponseWriter,r *http.Request){

}

func GetOrgByID(w http.ResponseWriter,r *http.Request){


}*/
