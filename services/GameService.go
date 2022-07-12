package services

import (
	"aatrox/chess/database"
	"aatrox/chess/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var db = database.DB

type GameService struct{}

func init() {
	var err error

	// create table
	fmt.Printf("creating \"games\" table\n")

	err = db.AutoMigrate(&models.Game{})
	if err != nil {
		fmt.Printf("db.AutoMigrate err: %s\n", err)
	}

	var hasTable bool = db.Migrator().HasTable(&models.Game{})
	if !hasTable {
		fmt.Printf("Table \"games\" doesn't exist")
		return
	}

	// initialize data
	fmt.Printf("initializing \"games\" data\n")

	db.Exec("DELETE FROM games")
	db.Exec("ALTER TABLE games AUTO_INCREMENT = 1")

	var file *os.File
	file, err = os.Open("./database/games.csv")
	if err != nil {
		fmt.Printf("os.Open err: %s\n", err)
		return
	}
	defer file.Close()

	var row_turns int64
	var scanner = bufio.NewScanner(file)
	//ignore first line
	scanner.Scan()
	for scanner.Scan() {
		var row = strings.Split(scanner.Text(), ",")

		row_turns, err = strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			fmt.Printf("strconv.ParseInt err: %s\n", err)
			continue
		}

		var result = db.Create(&models.Game{
			Turns:   row_turns,
			Status:  strings.Trim(row[1], " "),
			Winner:  strings.Trim(row[2], " "),
			WhiteID: strings.Trim(row[3], " "),
			BlackID: strings.Trim(row[4], " "),
			Moves:   strings.Trim(row[5], " "),
			ECOCode: strings.Trim(row[6], " "),
			Opening: strings.Trim(row[7], " "),
		})

		if result.Error != nil {
			fmt.Printf("db.Create err: %s\n", result.Error)
			return
		}
	}
}

func (s *GameService) GetAll() *[]models.Game {
	var games []models.Game
	var result = db.Find(&games)
	if result.Error != nil {
		fmt.Printf("db.Find err: %s\n", result.Error)
	}
	return &games
}

func (s *GameService) GetByID(id int64) *models.Game {
	var game models.Game
	var result = db.First(&game, id)
	if result.Error != nil {
		fmt.Printf("db.Find err: %s\n", result.Error)
	}
	return &game
}
