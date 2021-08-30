package domain

type Employee struct {
	NoUser    int    `json:"nouser" bson:"nouser"`
	Email     string `json:"email" bson:"email"`
	Full_name string `json:"full_name" bson:"full_name"`
	Data1     string `json:"data1" bson:"data1"`
	Data2     string `json:"data2" bson:"data2"`
}
