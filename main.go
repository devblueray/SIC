package main

import (
  "fmt"
 "gopkg.in/mgo.v2"
// "log"
// "gopkg.in/mgo.v2/bson"
"encoding/json"
"github.com/gorilla/mux"
"net/http"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type Employee struct {
  Name string `json:"name"`
  Phone string `json:"phone"`
  Email string `json:"email"`
  Department string `json:"department"`
  Manager string `json:"manager"`
  EmailGroups []string `json:"email_groups"`
  Assets map[string]Asset `json:"assetlist"`
}

type Asset struct {
  Tag string `json:"asset_tag"`
  Encrypted int `json:"encrypted"`
}

type mgoEmployee struct {
	Name        string `bson:"name"`
	Phone       string `bson:"phone"`
	Email       string `bson:"email"`
	Department  string `bson:"department"`
	Manager     string `bson:"manager"`
	EmailGroups []string `bson:"emailgroups"`
	Assets      map[string]mgoAsset `bson:"assets"`
}

type mgoAsset struct {
	Tag string `bson:"tag"`
	Encrypted bool`bson:"encrypted"`
}


func mongoSetup(database string, collection string) *mgo.Collection {
  session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
  session.SetMode(mgo.Monotonic, true)
  c := session.DB(database).C(collection)
  return c
}

func mgoInsert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	c := mongoSetup("sic", "employees")

	err := c.Insert(emp)
	if err != nil {
		result,_ := json.Marshal("Could not insert data")
		fmt.Fprintln(w,result)
	}
	var mgoEmp mgoEmployee
	err = c.Find(bson.M{"name": &emp.Name}).One(&mgoEmp)
	success,_ := json.Marshal(&mgoEmp)
	w.Write(success)
	//a,_:= json.Marshal(&emp)

}
func main() {
	m := mux.NewRouter()
	m.HandleFunc("/",mgoInsert).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080",m))



}

