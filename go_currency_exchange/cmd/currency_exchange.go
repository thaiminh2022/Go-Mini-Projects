package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/thaiminh2022/go_currency_exchange/internal"
	"golang.org/x/text/message"
)

func init() {
    err:=godotenv.Load()
    if err != nil {
        log.Fatal(err)
    }

}
func main() {
	internal.Ask()
	rate, err := internal.Answer()
	if err != nil {
		log.Fatal(err)
	}

    locate, err := GetLocale()
    if err != nil {
        fmt.Println(err)
        fmt.Println("Use default vietnamese locate: (, for decimal)(. for seperation)")
        locate = "en"

    }
    fmt.Println("Your locate is: ", locate)
    p := message.NewPrinter(message.MatchLanguage(locate))
    p.Println("Value: ",rate)

}

func GetLocale() (string, error) {
  // Check the LANG environment variable, common on UNIX.
  // XXX: we can easily override as a nice feature/bug.
  envlang, ok := os.LookupEnv("LANG")
  if ok {
    return strings.Split(envlang, ".")[0], nil
  }

  // Exec powershell Get-Culture on Windows.
  cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
  output, err := cmd.Output()
  if err == nil {
    return strings.Trim(string(output), "\r\n"), nil
  }

  return "", fmt.Errorf("cannot determine locale")
}
