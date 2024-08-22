package database

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var postgresConn = os.Getenv("WL_POSTGRES_CONN")

type User struct {
	Id        string
	Name      string
	Snatch    int
	CleanJerk int
	Total     int
}

func NewDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, postgresConn)
	if err != nil {
		log.Printf("Unable to create database pool: %v", err)
		return nil, err
	}
	return pool, nil
}

func CreateUser(ctx context.Context, pool *pgxpool.Pool, user *User) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.Printf("Failed to generate user ID: %v\n", err)
		return err
	}
	user.Id = hex.EncodeToString(b)
	user.Total = user.Snatch + user.CleanJerk
	_, err = conn.Exec(ctx, "INSERT INTO leaderboard (id, name, snatch, cleanjerk, total) VALUES ($1, $2, $3, $4, $5);", user.Id, user.Name, user.Snatch, user.CleanJerk, user.Total)
	if err != nil {
		log.Printf("Failed to insert user %s: %v", user.Id, err)
		return err
	}
	log.Println("ADDING USER: " + user.Name + " " + user.Id)
	return nil
}

func GetUsers(ctx context.Context, pool *pgxpool.Pool) ([]*User, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT id, name, snatch, cleanjerk, total FROM leaderboard;")
	if err != nil {
		log.Printf("Failed to retrieve users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Snatch, &user.CleanJerk, &user.Total)
		if err != nil {
			log.Printf("Failed to parse user: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}
	log.Println("RETURNING LIST OF USERS")
	return users, nil
}

func UpdateUser(ctx context.Context, pool *pgxpool.Pool, user *User) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "UPDATE leaderboard SET name = $2, snatch = $3, cleanjerk = $4, total = $5 WHERE id = $1;", user.Id, user.Name, user.Snatch, user.CleanJerk, user.Total)
	if err != nil {
		log.Printf("Failed to update user %s: %v", user.Id, err)
		return err
	}
	log.Println("UPDATING USER: " + user.Name + " " + user.Id)
	return nil
}

func DeleteUser(ctx context.Context, pool *pgxpool.Pool, id string) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "DELETE FROM leaderboard WHERE id = $1;", id)
	if err != nil {
		log.Printf("Failed to delete user %s, :%v", id, err)
		return err
	}
	log.Println("DELETING USER: " + id)
	return nil
}
