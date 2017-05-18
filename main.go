package main

import (
  "fmt"
  "log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
)

type Employee struct {
  Name string `bson:"name"`
  Phone string `bson:"phone"`
  Email string `bson:"email"`
  Department string `bson:"department"`
  Manager string `bson:"manager"`
  EmailGroups []string `bson:"emailgroups"`
  Assets map[string]Assets `bson:",inline"`  
}

type Assets struct {
  AssetType string `bson:"assettype"`
  AssetProperties map[string]string `bson:"assetproperties"`
}

type AssetProperties struct {
  AssetTag string `bson:"serialnumber"`
  Encrypted int `bson:"encrypted"`
}
   
func main() {
  session,err := mgo.Dial("127.0.0.1")
  if err != nil {
    fmt.Println("Mongo Dialer")
    panic(err)
  }

session.SetMode(mgo.Monotonic,true)

c := session.DB("sic").C("employees")

df := `
{
  "name": "Emp Name",
  "phone": "5555555555",
  "email": "email@address.com",
  "department": "Joker",
  "manager": "Some clueless Pleb",
  "emailgroups": [
    "test",
    "foo"
    ],
  "assets": {
      "Laptop": {
        "asset_tag": "1234",
        "encrypted": "1" 
    }
  }
}`

emp := Employee{}
err = json.Unmarshal([]byte(df), &emp)
fmt.Println(&emp)
c.Insert(emp)


err = c.Find(bson.M{"name": "Emp Name"}).One(&emp)
err = nil
if err != nil {
  fmt.Println("Insert")
  log.Fatal(err)
  }

fmt.Println("Phone:" , emp.Phone)


}

