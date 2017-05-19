package main

import (
  "fmt"
 // "log"
  "gopkg.in/mgo.v2"
 // "gopkg.in/mgo.v2/bson"
  "encoding/json"


  "log"
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

func main() {
  session, _ := mgo.Dial("127.0.0.1")
  defer session.Close()

  session.SetMode(mgo.Monotonic, true)
  c:= session.DB("sic").C("employees")


  df := `
{
  "name": "Emp Name",
  "phone": "5555555555",
  "email": "email@address.com",
  "department": "Joker",
  "manager": "Some clueless Pleb",
  "email_groups": [
    "test",
    "foo"
    ],
  "assetlist": {
      "Laptop": {
        "asset_tag": "1234",
        "encrypted": 1
    },
      "Tablet": {
        "asset_tag": "4422",
        "encrypted": 0
     }

  }
}`

emp := Employee{}
json.Unmarshal([]byte(df), &emp)
fmt.Println(&emp)
err := c.Insert(emp)
if err != nil {
  log.Fatal(err.Error())
}

/*
err := c.Find(bson.M{"name": "Emp Name"}).One(&emp)

if err != nil {
  fmt.Println("Insert")
  log.Fatal(err)
  }
*/

fmt.Println("Phone:" , emp.Phone)


}

