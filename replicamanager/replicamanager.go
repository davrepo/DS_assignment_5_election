package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	logger "github.com/davrepo/DS_assignment_5_election/logger"
	protos "github.com/davrepo/DS_assignment_5_election/proto"

	"google.golang.org/grpc"
)

type Server struct {
	protos.UnimplementedAuctionhouseServiceServer
	clientBids               map[int32]int32
	currentHighestBidsAmount int32
	isAuctionEnded           bool
	totalBids                int32
	id                       string
}

func main() {
	logger.ClearLog("log")
	logger.LogFileInit("main")

	replicaNumber, _ := strconv.Atoi(os.Args[1])

	go Start(int32(replicaNumber), (int32(replicaNumber)))

	bl := make(chan bool)
	<-bl
}

func Start(id int32, po int32) {
	port := po
	print(port)

	s := &Server{
		clientBids:     make(map[int32]int32),
		isAuctionEnded: false,
		totalBids:      0,
		id:             fmt.Sprintf("%d", id),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", po))
	if err != nil {
		logger.InfoLogger.Printf(fmt.Sprintf("FATAL: Connection unable to establish. Failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer()

	protos.RegisterAuctionhouseServiceServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.ErrorLogger.Fatalf("FATAL: replica connection failed: %s", err)
		}
	}()
	logger.InfoLogger.Println("mis")
	Output(fmt.Sprintf("Replica connected on port: %v", port))

	bl := make(chan bool)
	<-bl
}

func (s *Server) Bid(ctx context.Context, req *protos.BidRequest) (*protos.StatusOfBid, error) {

	log.Printf("Bid request received from client %v", req.ClientId)
	log.Printf("Bid amount: %v", req.Amount)
	if s.totalBids == 5 {
		s.isAuctionEnded = true
		return &protos.StatusOfBid{
			Status:     protos.Status_AUCTION_ENDED,
			HighestBid: s.currentHighestBidsAmount,
		}, nil
	} else {
		s.totalBids += 1

		s.clientBids[req.ClientId] = req.Amount
		if req.Amount > s.currentHighestBidsAmount {
			s.currentHighestBidsAmount = req.Amount

			return &protos.StatusOfBid{
				Status:     protos.Status_NOW_HIGHEST_BIDDER,
				HighestBid: s.currentHighestBidsAmount,
			}, nil

		} else if req.Amount < s.currentHighestBidsAmount {
			return &protos.StatusOfBid{
				Status:     protos.Status_TOO_LOW_BID,
				HighestBid: s.currentHighestBidsAmount,
			}, nil
		} else {
			return &protos.StatusOfBid{
				Status:     protos.Status_EXCEPTION,
				HighestBid: s.currentHighestBidsAmount,
			}, nil

		}
	}

}

func (s *Server) Result(ctx context.Context, in *protos.QueryResult) (*protos.ResponseToQuery, error) {
	status := protos.Status_AUCTION_ACTIVE
	if s.isAuctionEnded {
		status = protos.Status_AUCTION_ENDED
	}

	return &protos.ResponseToQuery{
		Status:     status,
		HighestBid: s.currentHighestBidsAmount,
	}, nil
}

func Output(input string) {
	log.Println(input)
}

func ReadPorts() ([]string, error) {
	data, err := os.ReadFile("/home/hanan/DS_assignment_5_election/replicamanager/portlist/listOfReplicaPorts.txt")
	if err != nil {
		return nil, err
	}

	var ports []string
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		log.Println(line)
		strings.TrimSpace(line)

		ports = append(ports, line)
	}

	return ports, nil
}

func (s *Server) SendData(req protos.AuctionhouseService_SendDataServer) error {

	data := &protos.SendDataResponse{
		Amount:    int64(s.currentHighestBidsAmount),
		BackupIds: s.id,
	}

	if err := req.Send(data); err != nil {
		log.Printf("send error %v", err)
	}

	return nil
}
