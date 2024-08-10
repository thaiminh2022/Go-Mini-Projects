package tools

import "log"

type DatabaseInterface interface {
	GetUrlAtKey(key string) (string, error)
	StoreKeyUrl(key string, url string) error
	UrlExists(url string) (string, bool)
	LoadDB() error
	SaveDB() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &mockDB{make(map[string]string)}
	err := db.LoadDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &db, nil
}
