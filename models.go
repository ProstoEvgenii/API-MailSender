package main

import "time"

type Users struct {
	FirstName   string    `bson:"Имя"`
	LastName    string    `bson:"Фамилия"`
	MiddleName  string    `bson:"Отчество"`
	DateOfBirth time.Time `bson:"Дата рождения"`
	Email       string    `bson:"E-mail"`
}
type Templates struct {
	Name      string `bson:"name"`
	IndexHTML string `bson:"indexHTML"`
}

type Logs struct {
	FirstName   string    `bson:"Имя"`
	LastName    string    `bson:"Фамилия"`
	MiddleName  string    `bson:"Отчество"`
	DateOfBirth time.Time `bson:"Дата рождения"`
	Email       string    `bson:"E-mail"`
	DateCreate  time.Time `bson:"dateCreate"`
}
type Dashboard_Params struct {
	SendAll string `json:"sendAll"`
}

type Response struct {
	DocumentsCount    int64  `json:"documentsCount"`
	CountBirtdays     int    `json:"countBirtdays"`
	CountLogs         int    `json:"сountNumber"`
	DocumentsInserted int64  `json:"documentsInserted"`
	SendEmail         string `json:"sendEmail"`
}
type UsersUpload struct {
	Last_name   string `json:"Фамилия" bson:"Фамилия"`
	First_name  string `json:"Имя" bson:"Имя"`
	Middle_name string `json:"Отчество" bson:"Отчество"`
	Date_birth  string `json:"Дата рождения" bson:"Дата рождения"`
	Email       string `json:"E-mail" bson:"E-mail"`
}

type Data struct {
	Last_name   string `json:"Фамилия" bson:"Фамилия"`
	First_name  string `json:"Имя" bson:"Имя"`
	Middle_name string `json:"Отчество" bson:"Отчество"`
	Date_birth  string `json:"Дата рождения" bson:"Дата рождения"`
	Email       string `json:"E-mail" bson:"E-mail"`
}
