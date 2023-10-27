package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Dashboard() (int64, int) {
	usersCount := CountDocuments()
	birthdays_list := Fi22ndBirthdays()
	return usersCount, len(birthdays_list)
}
func Fi22ndBirthdays() []Users {
	today := time.Now()
	filter := bson.M{}
	cursor := Find(filter, "users")
	var users []Users
	err := cursor.All(context.TODO(), &users)
	if err != nil {
		log.Println("=84ce91=", err)
	}

	var birthdays_list []Users
	for _, user := range users {
		if user.DateOfBirth.Day() == today.Day() && user.DateOfBirth.Month() == today.Month() {
			result := CreateLog(user)
			log.Println("=56eccc=", result)
			birthdays_list = append(birthdays_list, user)
			// log.Println("=afdf3c=", user.DateOfBirth.Day())
		}
	}

	return birthdays_list
}

func CreateLog(user Users) int64 {
	today := time.Now()
	filter := bson.M{
		"E-mail": user.Email,
	}
	update := bson.M{"$setOnInsert": bson.M{
		"Имя":           user.FirstName,
		"Фамилия":       user.LastName,
		"Отчество":      user.MiddleName,
		"Дата рождения": user.DateOfBirth,
		"E-mail":        user.Email,
		"dateCreate":    today,
	}}
	result := InsertIfNotExists(user, filter, update, "logs").UpsertedCount
	return result
}

// func FindLogs() []Users {
// 	today := time.Now()
// 	twentyFourHoursAgo := today.Add(-24 * time.Hour)
// 	// filter := bson.M{}
// 	filter := bson.M{"Дата рождения": bson.M{"$gte": twentyFourHoursAgo}}
// 	cursor := Find(filter, "logs")
// 	var users_logs []Users
// 	err := cursor.All(context.TODO(), &users_logs)
// 	if err != nil {
// 		log.Println("=84ce91=", err)
// 	}

// 	// var logs_list []Users
// 	// for _, user := range users_logs {
// 	// 	if user.DateOfBirth.Day() == today.Day() && user.DateOfBirth.Month() == today.Month() {
// 	// 		birthdays_list = append(birthdays_list, user)
// 	// 	}
// 	// }

// 	return
// }
