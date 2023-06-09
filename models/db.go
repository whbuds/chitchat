package models

import (
	. "chitchat/config"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	config := LoadConfig()
	driver := config.Db.Driver
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true", config.Db.User, config.Db.Password,
		config.Db.Address, config.Db.Database)
	Db, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// createUUID Create a random UUID with from RFC 4122
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] | 0xF) & (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
