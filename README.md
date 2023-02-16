# integtation-status-api
# integtation-status-api



package main

//importing all necessary package //
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// structure for organisation//
type Org struct {
	Org_id   int    `json : "org_id"`
	Org_name string `json : "org_name"`
}

//structure for organisation//

type Cms struct {
	Cms_id   int    `json : "cms_id"`
	Cms_name string `json : "cms_name"`
}

//structure for organisation//

type Status struct {
	Stat_id   int    `json : "stat_id"`
	Stat_name string `json : "stat_name"`
}

//structure for organisation//

type App struct {
	App_id   string `json : "app_id"`
	App_name string `json : "app_name"`
}

//structure for organisation//

type Intgs struct {
	I_id     int    `json : "i_id"`
	App_url  string `json : "app_url"`
	Comment  string `json : "comment"`
	I_app    string `json : "i_app"`
	I_cms    int    `json :"i_cms"`
	I_org    int    `json : "i_org"`
	I_status int    `json : "i_status"`
}

//Sql connection

var db *sql.DB
var err error

// main method
func main() {
	//sql connection
	db, err = sql.Open("mysql", "root:admin@123@tcp(127.0.0.1:3306)/intg_stat")
	if err != nil {
		panic(err.Error())
	}
	//functions for apis//
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/org", GetAllOrg).Methods("GET")
	router.HandleFunc("/status", GetAllStat).Methods("GET")
	router.HandleFunc("/cms", GetAllCms).Methods("GET")
	router.HandleFunc("/app", GetAllApp).Methods("GET")
	router.HandleFunc("/intgs", GetAllIntgs).Methods("GET")
	router.HandleFunc("/intgs", PostAllIntgs).Methods("POST")
	router.HandleFunc("/intgs/{i_id}", DeleteIntgsById).Methods("DELETE")
	router.HandleFunc("/intgs/{i_id}", UpdateIntgsById).Methods("PUT")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("Server running on port : 8010")
	http.ListenAndServe(":8010", handler)
}
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,GET,GET,GET,GET,OPTIONS,POST")
}

// Access the Data from organization
func GetAllOrg(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM org")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var org_id []Org
	for result.Next() {
		var org Org
		err := result.Scan(&org.Org_id, &org.Org_name)
		if err != nil {
			panic(err.Error())
		}
		org_id = append(org_id, org)

	}
	json.NewEncoder(w).Encode(org_id)
}

// Access the Data from status table

func GetAllStat(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM status")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var stat_id []Status
	for result.Next() {
		var status Status
		err := result.Scan(&status.Stat_id, &status.Stat_name)
		if err != nil {
			panic(err.Error())
		}
		stat_id = append(stat_id, status)

	}
	json.NewEncoder(w).Encode(stat_id)
}

// Access the Data from app table

func GetAllApp(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM app")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var app_id []App
	for result.Next() {
		var app App
		err := result.Scan(&app.App_id, &app.App_name)
		if err != nil {
			panic(err.Error())
		}
		app_id = append(app_id, app)

	}
	json.NewEncoder(w).Encode(app_id)
}

// Access the Data from cms table

func GetAllCms(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM cms")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var cms_id []Cms
	for result.Next() {
		var cms Cms
		err := result.Scan(&cms.Cms_id, &cms.Cms_name)
		if err != nil {
			panic(err.Error())
		}
		cms_id = append(cms_id, cms)

	}
	json.NewEncoder(w).Encode(cms_id)
}

// Access the Data from intgs table

func GetAllIntgs(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM intgs")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var i_id []Intgs
	for result.Next() {
		var intgs Intgs
		err := result.Scan(&intgs.I_id, &intgs.I_org, &intgs.I_cms, &intgs.I_status, &intgs.I_app, &intgs.App_url, &intgs.Comment)
		if err != nil {
			panic(err.Error())
		}
		i_id = append(i_id, intgs)

	}
	json.NewEncoder(w).Encode(i_id)
}

// Posting sending the Data to database intgs table

func PostAllIntgs(w http.ResponseWriter, r *http.Request) {

	//EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO `intgs`(`i_org`,`i_cms`,`i_status`,`i_app`,`app_url`,`comment`) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]int)
	json.Unmarshal(body, &keyVal)
	i_org := keyVal["i_org"]
	i_cms := keyVal["i_cms"]
	i_status := keyVal["i_status"]
	keyVal1 := make(map[string]string)
	json.Unmarshal([]byte(body), &keyVal1)
	app_url := keyVal1["app_url"]
	comment := keyVal1["comment"]
	i_app := keyVal1["i_app"]

	_, err = stmt.Exec(i_org, i_cms, i_status, i_app, app_url, comment)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Fprintf(w, "Congratulations! Data is added successfully.... \n")
	resp := make(map[string]interface{})
	resp["message"] = "Success"
	resp["status"] = 200
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func DeleteIntgsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("DELETE FROM intgs WHERE i_id=?;", params["i_id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
}

func UpdateIntgsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("UPDATE intgs SET i_org=?,i_cms=?,i_status=?,i_app=?,app_url=?,comment=? WHERE i_id=?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]int)
	json.Unmarshal(body, &keyVal)
	i_org := keyVal["i_org"]
	i_cms := keyVal["i_cms"]
	i_status := keyVal["i_status"]
	keyVal1 := make(map[string]string)
	json.Unmarshal([]byte(body), &keyVal1)
	app_url := keyVal1["app_url"]
	comment := keyVal1["comment"]
	i_app := keyVal1["i_app"]

	_, err = stmt.Exec(i_org, i_cms, i_status, i_app, app_url, comment, params["i_id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Congratulations! Data is updated successfully.... \n")
	resp := make(map[string]interface{})
	resp["message"] = "Success"
	resp["status"] = 200
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

/*func GetOrgById(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("select * from org where org_id=?;", params["org_id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var org_id []Org
	for result.Next() {
		var org Org
		err := result.Scan(&org.Org_id, &org.Org_name)
		if err != nil {
			panic(err.Error())
		}
		org_id = append(org_id, org)

	}
	json.NewEncoder(w).Encode(org_id)
}
func GetStatById(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("Select * FROM status where stat_id=?;", params["stat_id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var stat_id []Status
	for result.Next() {
		var status Status
		err := result.Scan(&status.Stat_id, &status.Stat_name)
		if err != nil {
			panic(err.Error())
		}
		stat_id = append(stat_id, status)

	}
	json.NewEncoder(w).Encode(stat_id)
}
func GetAppById(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("Select * FROM app where app_id=?;", params["app_id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var app_id []App
	for result.Next() {
		var app App
		err := result.Scan(&app.App_id, &app.App_name)
		if err != nil {
			panic(err.Error())
		}
		app_id = append(app_id, app)

	}
	json.NewEncoder(w).Encode(app_id)
}

// Access the Data from cms table by id

func GetCmsById(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("Select * FROM cms where cms_id=?;", params["cms_id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var cms_id []Cms
	for result.Next() {
		var cms Cms
		err := result.Scan(&cms.Cms_id, &cms.Cms_name)
		if err != nil {
			panic(err.Error())
		}
		cms_id = append(cms_id, cms)

	}
	json.NewEncoder(w).Encode(cms_id)
}
*/
