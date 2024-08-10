package tools

import (
	"encoding/json"
	"errors"
	"os"
)

type mockDB struct {
	KeyUrls map[string]string `json:"keyUrls"`
}

func (db *mockDB) GetUrlAtKey(key string) (string, error) {
	value, ok := db.KeyUrls[key]
	if ok {
		return value, nil
	}
	return "", errors.New("key does not exist")
}
func (db *mockDB) StoreKeyUrl(key string, url string) error {
	_, ok := db.KeyUrls[key]
	if ok {
		return errors.New("key already exist")
	}

	// value already exist in db
	for _, v := range db.KeyUrls {
		if v == url {
			return nil
		}
	}
	db.KeyUrls[key] = url
	return nil
}

func (db *mockDB) LoadDB() error {
	data, err := os.ReadFile("database.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &db)
	if err != nil {
		return err
	}

	return nil
}

func (db *mockDB) SaveDB() error {
	data, err := json.Marshal(&db)
	if err != nil {
		return err
	}
	return os.WriteFile("database.json", data, 0664)
}

func (db *mockDB) UrlExists(url string) (string, bool) {
	for key, value := range db.KeyUrls {
		if value == url {
			return key, true
		}
	}
	return "", false
}
