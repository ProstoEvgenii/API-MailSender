package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// type apiRequest struct {
// 	Email string `json:"email"`
// }

func Start() {

	http.HandleFunc("/", anyPage)
	http.HandleFunc("/api", ParseRequest)
	// http.HandleFunc("/api/updateBD", updateBD)
	http.HandleFunc("/api/upload", uploadHandler)
	http.ListenAndServe(":80", nil)

}

// func updateBD(rw http.ResponseWriter, request *http.Request) {
// 	if request.Method == "POST" {
// 		rw.Header().Set("Content-Type", "application/json")
// 		rw.Header().Set("Access-Control-Allow-Origin", "*")
// 		rw.Header().Set("Access-Control-Max-Age", "15")

// 		decoder := json.NewDecoder(request.Body)

// 		var data []Data
// 		err := decoder.Decode(&data)
// 		if err != nil {
// 			log.Println("=b570dd=", err)
// 		}
// 		for _, item := range data {
// 			InsertIfNotExists(item)
// 		}
// 		return
// 	}
// }

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	file, _, err := r.FormFile("jsonFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file) // Читаем содержимое файла в срез байтов
	if err != nil {
		// http.Error(w, "Не удалось прочитать файл", http.StatusInternalServerError)
		return
	}
	var users []UsersUpload //  Форматитирую срез байтов в структуру
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}
	var result int64

	for _, document := range users {
		filter := bson.M{
			"E-mail": document.Email,
		}
		dateBirth, _ := time.Parse("01/02/2006", document.Date_birth)
		update := bson.M{"$setOnInsert": bson.M{
			"Имя":           document.First_name,
			"Фамилия":       document.Last_name,
			"Отчество":      document.Middle_name,
			"Дата рождения": dateBirth,
			"E-mail":        document.Email,
		}}
		result += InsertIfNotExists(document, filter, update, "users").UpsertedCount
	}

	response := Response{
		Count: result,
	}

	usersAdded, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(usersAdded)
}
func ParseRequest(rw http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Max-Age", "15")

		usersCount, birthdaysListLen := Dashboard()
		response := Response{
			Count:         usersCount,
			CountBirtdays: birthdaysListLen,
		}

		itemCountJson, err := json.Marshal(response)
		if err != nil {
			fmt.Println("error:", err)
		}
		rw.Write(itemCountJson)
		return
	} else if request.Method == "POST" {
		// r.FormFile("userFile")
		// rw.Header().Set("Content-Type", "application/json")
		// rw.Header().Set("Access-Control-Allow-Origin", "*")
		// rw.Header().Set("Access-Control-Max-Age", "15")

		// decoder := json.NewDecoder(request.Body)

		// var data apiRequest
		// err := decoder.Decode(&data)
		// if err != nil {
		// 	panic(err)
		// 	// log.Fatal("Aborting", err)
		// }
		// log.Println("=90674d=", data.Email)
		// SendEmail(data.Email)
		return
	}
}

func anyPage(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(rw, "Hello")
}
