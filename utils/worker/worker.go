package worker

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"
	"mudiralmaham/models"
	"mudiralmaham/utils/database"
	"time"
)

func SendNotification(newTask chan models.Task) {
	var tasks []models.Task

	cursor, err := database.DB.Collection("task").Find(context.TODO(), bson.M{})
	if err != nil {
		print("problem in worker")
		return
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		err = cursor.Decode(&task)
		if err != nil {
			return
		}
		tasks = append(tasks, task)
	}
	for {
		select {
		case task := <-newTask:
			tasks = append(tasks, task)
		default:
			for idx := 0; idx < len(tasks); idx++ {
				due, err := time.Parse("Jan 2, 2006 15:04:05", tasks[idx].DueDate)
				now := time.Now()
				if err != nil {
					due, err = time.Parse("Jan 2, 2006 3:04:05 PM", tasks[idx].DueDate)
					if err != nil {
						println(err.Error())
						continue
					}
				}
				if idx >= len(tasks) {
					break
				}
				if isAfter(now, due) && !tasks[idx].IsOver {

					tasks[idx].IsOver = true

					m := gomail.NewMessage()
					m.SetHeader("From", "magaroojoo@gmail.com")
					m.SetHeader("To", tasks[idx].Owner)
					m.SetHeader("Subject", "task due time")
					m.SetBody("text/html",
						"<body>Hi there "+tasks[idx].Owner+"!<br><p> your task "+tasks[idx].Name+" has reached its due time, be sure to have completed it! <br>Good Day :)</p></body>")
					d := gomail.NewDialer("smtp.gmail.com", 587, "magaroojoo@gmail.com", "majid77??")
					if err := d.DialAndSend(m); err != nil {
						println(err.Error())
						continue
					}
					_, err = database.DB.Collection("task").UpdateOne(context.TODO(), bson.M{"name": tasks[idx].Name}, bson.M{"$set": bson.M{"isover": true}})
					if err != nil {
						println(err.Error())
					}
				}
			}
		}
	}
}

func isAfter(first time.Time, second time.Time) bool {
	if first.Year() > second.Year() {
		return true
	} else if first.Year() < second.Year() {
		return false
	}
	if first.Month() > second.Month() {
		return true
	} else if first.Month() < second.Month() {
		return false
	}
	if first.Day() > second.Day() {
		return true
	} else if first.Day() < second.Day() {
		return false
	}
	if first.Hour() > second.Hour() {
		return true
	} else if first.Hour() < second.Hour() {
		return false
	}
	if first.Minute() > second.Minute() {
		return true
	} else if first.Minute() < second.Minute() {
		return false
	}
	return true
}
