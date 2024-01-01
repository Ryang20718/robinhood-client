package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// model "github.com/Ryang20718/robinhood-client/models"
	client "github.com/Ryang20718/robinhood-client/client"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter MFA please: ")
	mfa, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read stdin")
		return
	}
	mfa = strings.TrimSuffix(mfa, "\n")
	cli, _ := client.Dial(&client.OAuth{
		Username: os.Getenv("ROBINHOOD_USERNAME"),
		Password: os.Getenv("ROBINHOOD_PASSWORD"),
		MFA:      mfa,
	})
	res, err := cli.GetInstrumentForSymbol("ZNGA")
	if err != nil {
		fmt.Println("HERE", err.Error())
	}
	fmt.Println(res)
}
