// package main

// import (
// 	"bufio"
// 	"context"
// 	"fmt"
// 	"os"
// 	"strings"

// 	// model "github.com/Ryang20718/robinhood-client/models"
// 	client "github.com/Ryang20718/robinhood-client/client"
// )

// func main() {
// 	ctx := context.Background()

//		reader := bufio.NewReader(os.Stdin)
//		fmt.Print("Enter MFA please: ")
//		mfa, err := reader.ReadString('\n')
//		if err != nil {
//			fmt.Println("Failed to read stdin")
//			return
//		}
//		mfa = strings.TrimSuffix(mfa, "\n")
//		cli, _ := client.Dial(&client.OAuth{
//			Username: os.Getenv("ROBINHOOD_USERNAME"),
//			Password: os.Getenv("ROBINHOOD_PASSWORD"),
//			MFA:      mfa,
//		})
//		res, _ := cli.GetOptionsOrders(ctx)
//		tes := 0.0
//		for _, val := range *res {
//			if strings.Split(val.CreatedAt, "-")[0] != "2023" {
//				continue
//			}
//			if val.Status != "Expired" {
//				fmt.Println(val)
//				continue
//			}
//			if val.TransactionType == "STO" || val.TransactionType == "STC" {
//				tes += (val.Qty * val.UnitCost * 100)
//			} else if val.TransactionType == "BTO" || val.TransactionType == "BTC" {
//				tes -= (val.Qty * val.UnitCost * 100)
//			}
//		}
//	}