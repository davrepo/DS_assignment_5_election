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

var ports = [4]string{"3001", "3002", "3003", "3004"}
var primaryPort string
var replicaPort string

func main() {
	log.Print(os.Args[1])
	primaryPort = fmt.Sprintf("%v", os.Args[1])
	replicaPort = primaryPort
	Output(WelcomeMsg())
	go updateReplica()

	replicaConnect(primaryPort)

	//______________________

	bl := make(chan bool)
	<-bl
}

func connectServer(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("could not connect: %v", err)
	}

	//--------------------
	return conn, nil

}

func activeChat(client protos.AuctionhouseServiceClient) {
	active := true

	for active {
		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{"bid", "query", "leave"},
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
		} else if input == "query" {
			result(client)
		} else if input == "leave" {
			active = false
			os.Exit(3)
		}

	}
}

func bid(client protos.AuctionhouseServiceClient, amount int32) {
	// Your code for the bid function goes here
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.Bid(ctx, &protos.BidRequest{
		ClientId: 10,
		Amount:   amount,
	})
	if err != nil {
		// if the bid fails try another replica
		replicaConnect("")
	}

	log.Print("Your bid status: ", res.Status)
	log.Print("Highest bid: ", res.HighestBid)

}

func result(client protos.AuctionhouseServiceClient) (*protos.ResponseToQuery, error) {
	// Your code for the bid function goes here
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.Result(ctx, &protos.QueryResult{})
	if err != nil {
		log.Printf("could not get")
		return nil, err
	}

	log.Print("Auction status: ", res.Status)
	log.Print("Current highest bid: ", res.HighestBid)

	return res, nil
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

func replicaConnect(primaryPort string) {

	client := protos.NewAuctionhouseServiceClient(nil)

	for i := 0; i < len(ports); i++ {
		if isServerLive(ports[i]) {
			log.Printf("connecting to replica %v", ports[i])
			conn, err := connectServer(ports[i])
			client = protos.NewAuctionhouseServiceClient(conn)
			log.Printf("connected to replica %v", ports[i])
			activeChat(client)
			if err != nil {
				log.Printf("error connecting")
			}
			primaryPort = ports[i]
			replicaPort = ports[i+1]
			break
		}
	}
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

// this method is used to ping the server at see if it is active if not try next
func isServerLive(port string) bool {
	// Connect to the server
	conn, err := connectServer(port)
	if err != nil {
		log.Printf("could no connect")
		return false
	}

	// Create a new client
	client := protos.NewAuctionhouseServiceClient(conn)

	// Call a simple method from the service
	_, err = result(client)

	if err != nil {
		log.Printf("The server is on " + port + " is down")

		return false
	} else {
		log.Printf("The server is on " + port + " is up")

		return true
	}

}

func Output(input string) {
	log.Println(input)
}

func updateReplica() {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		// Create a connection to the replica
		conn, err := connectServer(primaryPort)

		if err != nil {
			log.Printf("Error connecting to replica: %v", err)
			continue
		}

		// Create a client
		client := protos.NewAuctionhouseServiceClient(conn)

		// Create a SendData stream
		stream, err := client.SendData(context.Background())
		if err != nil {
			log.Printf("Error semding SendData stream: %v", err)
			continue
		}

		curResult, err := stream.Recv()

		log.Printf("Current result: %v", curResult)

		// Send the current result to the replica
		// Close the stream
		stream.CloseSend()
		
		// Close the connection
		conn.Close()

		//Replica 
		// Create a connection to the replica
		connReplica, err := connectServer(replicaPort)

		if err != nil {
			log.Printf("Error connecting to replica: %v", err)
			continue
		}

		// Create a client
		clientReplica := protos.NewAuctionhouseServiceClient(connReplica)


		log.Printf("Current result: %v", curResult)

		// Send the current result to the replica
		// Close the stream
		stream.CloseSend()
		
		// Close the connection
		conn.Close()

		replicaUpdateAuction(clientReplica, curResult)

	}
}

func replicaUpdateAuction(clientReplica protos.AuctionhouseServiceClient, curResult *protos.SendDataResponse ) {
	// Your code for the bid function goes here
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := clientReplica.SendDataToReplica(ctx, &protos.GetDataRequestToReplica{
		TotalBids:               curResult.TotalBids,
		CurrentHighestBidsAmount: curResult.CurrentHighestBidsAmount,
		IsAuctionEnded:           curResult.IsAuctionEnded,
	})
	if err != nil {
		log.Fatalf("Error updating replica: %v", err)
	}

	log.Print("Replica status: ", res.Status)

}


