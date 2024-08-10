package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Database struct {
	ExchangeRates map[string]float64 `json:"exchangeRates"`
	TimeLastSave  time.Time          `json:"timeLastSave"`
}

func (db *Database) Exchange(from string, to string, amount float64) (float64, error) {
    fmt.Println("Converting", amount, from, "to", to)
	fromRate, ok := db.ExchangeRates[from]
	if !ok {
		return -1, errors.New("Cannot find from exchange")
	}
	toRate, ok := db.ExchangeRates[to]
	if !ok {
		return -1, errors.New("Cannot find to exchange")
	}

    fmt.Println("rates relative to USD:", fromRate, "->", toRate)
	return (amount / fromRate) * toRate, nil

}

func getDbFilePath() string {
    ex, err := os.Executable();
    if err != nil {
        log.Fatal(err)

    }
    exPath := filepath.Dir(ex)

    return exPath + "/database.json"

}

func (db *Database) LoadDatabase() error {
	data, err := os.ReadFile(getDbFilePath())
	if err != nil {
		return err
	}
	json.Unmarshal(data, db)
	return nil
}

func (db *Database) CheckDatabaseDate() {
	since := time.Since(db.TimeLastSave).Minutes()

	if since > 60 || len(db.ExchangeRates) == 0 {
		rates := RequestMoreData()
		db.ExchangeRates = rates
	}
	db.SaveDatabase()
}

func (db *Database) NewDatabase() error {
	db.TimeLastSave = time.Now()
	db.ExchangeRates = make(map[string]float64)

	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = os.WriteFile(getDbFilePath(), data, 0644)
	if err != nil {
		return err
	}

	return nil
}
func (db *Database) SaveDatabase() error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = os.WriteFile(getDbFilePath(), data, 0644)
	if err != nil {
		return err
	}

	return nil
}
