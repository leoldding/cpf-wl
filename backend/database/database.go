package database

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

type Database struct {
	Users []*User
}

type User struct {
	Id        string
	Name      string
	Snatch    int
	CleanJerk int
	Total     int
}

func NewDatabase() *Database {
	return &Database{[]*User{}}
}

func (database *Database) LoadMockData() {
	log.Println("LOADING MOCK DATA")
	database.CreateUser(&User{Name: "Alfred Alvarado", Snatch: 20, CleanJerk: 40})
	database.CreateUser(&User{Name: "Benjamin Bolognese", Snatch: 40, CleanJerk: 90})
	database.CreateUser(&User{Name: "Charles Charleston", Snatch: 10, CleanJerk: 20})
	database.CreateUser(&User{Name: "Damian Dog", Snatch: 150, CleanJerk: 290})
}

func (database *Database) CreateUser(user *User) {
	log.Println("ADDING USER: " + user.Name)
	b := make([]byte, 4)
	rand.Read(b)
	user.Id = hex.EncodeToString(b)
	user.Total = user.Snatch + user.CleanJerk
	database.Users = append(database.Users, user)
}

func (database *Database) GetUsers() []*User {
	log.Println("RETURNING LIST OF USERS")
	return database.Users
}

func (database *Database) UpdateUser(user *User) {
	for i, dbUser := range database.Users {
		if user.Id == dbUser.Id {
			database.Users[i] = user
			log.Println("UPDATING USER: " + user.Id)
			return
		}
	}
}

func (database *Database) DeleteUser(id string) {
	for i, user := range database.Users {
		if user.Id == id {
			database.Users[i] = database.Users[len(database.Users)-1]
			database.Users = database.Users[:len(database.Users)-1]
			log.Println("DELETING USER: " + user.Id)
			return
		}
	}
}
