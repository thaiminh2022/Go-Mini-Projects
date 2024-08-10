package internal

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

func getCurrency(title string, value *string) huh.Field {
	return huh.NewSelect[string]().
		Options(huh.NewOptions(CurrencyOptions...)...).
        Height(12).
		Value(value).
		Title(title)
}

var from, to string
var amountText string
var confirmed bool

func Ask() bool {
	form := huh.NewForm(
		huh.NewGroup(
			getCurrency("Welcome to currency exchange, please choose a FROM currency", &from),
			getCurrency("Please choose a TO currency", &to),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Amount of exchange").
				Validate(func(s string) error {
					value, err := strconv.
						ParseFloat(strings.TrimSpace(amountText), 64)

					if err != nil {
						return err
					}

					if value <= 0 {
						return errors.New("Cannot input 0 or negative number")
					}

					return nil
				}).
				Value(&amountText),
			huh.NewConfirm().
				Title("Confirm value").
				Affirmative("Cool").
				Negative("Take me back").
				Value(&confirmed),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	return confirmed
}

func Answer() (float64, error) {
	db := Database{}
	err := db.LoadDatabase()
	if err != nil {
		fmt.Println("Cannot find db, create new")
		err := db.NewDatabase()
		if err != nil {
			log.Fatal(err)
		}
	}

	db.CheckDatabaseDate()
	amount, err := strconv.ParseFloat(strings.TrimSpace(amountText), 64)
	if err != nil {
		log.Fatal("Conversion pass validation somehow")
	}

	return db.Exchange(from, to, amount)
}
