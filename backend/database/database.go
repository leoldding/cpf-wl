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
}

func NewDatabase() *Database {
	return &Database{[]*User{}}
}

func (database *Database) LoadMockData() {
	log.Println("LOADING MOCK DATA")
	database.AddUser(&User{Name: "Alfred A.", Snatch: 20, CleanJerk: 40})
	database.AddUser(&User{Name: "Benjamin B.", Snatch: 40, CleanJerk: 90})
	database.AddUser(&User{Name: "Charles C.", Snatch: 10, CleanJerk: 20})
	database.AddUser(&User{Name: "Damian D.", Snatch: 150, CleanJerk: 290})
}

func (database *Database) AddUser(user *User) {
	log.Println("ADDING USER: " + user.Name)
	b := make([]byte, 4)
	rand.Read(b)
	user.Id = hex.EncodeToString(b)
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
			log.Println("UPDATING USER: " + user.Name)
			return
		}
	}
}

func (database *Database) DeleteUser(id string) {
	for i, user := range database.Users {
		if user.Id == id {
			database.Users[i] = database.Users[len(database.Users)-1]
			database.Users = database.Users[:len(database.Users)-1]
			log.Println("DELETING USER: " + user.Name)
			return
		}
	}
}
