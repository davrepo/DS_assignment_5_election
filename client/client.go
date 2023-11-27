package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	protos "github.com/davrepo/DS_assignment_5_election/proto"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Print(os.Args[1])
	port := fmt.Sprintf(":%v", os.Args[1])

	Output(WelcomeMsg())
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//--------------------
	client := protos.NewAuctionhouseServiceClient(conn)

	activeChat(client)

	//______________________

	bl := make(chan bool)
	<-bl
}

func activeChat(client protos.AuctionhouseServiceClient) {
	active := true

	for active {
		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{"bid", "leave"},
		}
		_, input, err := prompt.Run()
		if err != nil {
			log.Fatalf("could not get input: %v", err)
		}

		if input == "bid" {
			if active {
				prompt := promptui.Prompt{
					Label: "input your bid ",
					Validate: func(input string) error {
						if len(input) == 0 || len(input) > 128 {
							return fmt.Errorf("input must be <1000>")
						}
						return nil
					},
				}
				input, err := prompt.Run()
				if err != nil {
					log.Fatalf("could not get input: %v", err)
				}
				i, err := strconv.ParseInt(input, 10, 32)
				if err != nil {
					panic(err)
				}
				result := int32(i)
				fmt.Printf("Parsed int is %d\n", result)
				bid(client, result)

			}
		} else if input == "leave" {
			active = false
		}

	}
}

func bid(client protos.AuctionhouseServiceClient, amount int32) {
	// Your code for the bid function goes here
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := client.Bid(ctx, &protos.BidRequest{
		ClientId: 10,
		Amount:   amount,
	})
	if err != nil {
		log.Fatalf("could not place bid: %v", err)
	}

	log.Print(res)

}

func WelcomeMsg() string {
	return `
______________________________________________________
======================================================
    **>>> WELCOME TO AUCTIONHOUSE <<<**
======================================================
Here you can bid on different items.
` + Instrc()
}

func Help() {
	Output(Instrc())
}

func Instrc() string {

	return `
	This is the Auction House, here you can bid on different items.
	A certain amount of time is set off for clients to bid on an item.
	The time on the items are NOT displayed to the clients, so if you wanna bid do it fast.

	INPUTS
	----------------------------------------------------------------------------------------------------------------
		Bidding on an item: 
			To bid on an item just write the amount in the terminal, followed by enter, the bid must be a valid int.
			     bid <bid amount>
		Information about current item:
			To ask the auctioneer what item you are bidding on and what the highest bid is please write:
				query
			in the terminal, followed by enter.

		Quitting:
			To quit the auction please write:
				quit
			in the terminal, followed by enter.

		Help:
			To get the input explaination again please write:
				help
			in the terminal, followed by enter.
	------------------------------------------------------------------------------------------------------------------

		`
}




func Output(input string) {
	log.Println(input)
}
