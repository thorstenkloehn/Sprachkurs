package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

type Vokabelm struct {
	Id              int       `json:"id"`
	VokabelEnglisch string    `json:"Vokabel_Englisch"`
	VokabelDeutsch  string    `json:"Vokabel_Deutsch"`
	Lauschrift      string    `json:"Lauschrift"`
}

func main() {

	if _, err := os.Stat("vokabeln.db"); os.IsNotExist(err) {
		fmt.Println("file does not exist")

		db, err := gorm.Open(sqlite.Open("vokabeln.db"), &gorm.Config{})
		if err != nil {
			panic("Fehler bei der Verbindung zur Datenbank")
		}
		db.AutoMigrate(&Vokabelm{})
	}

}
