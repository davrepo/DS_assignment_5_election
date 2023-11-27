package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/davrepo/DS_assignment_5_election/database"
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
}

func Start(id int32, po int32) {
	port := po
	print(port)
	logger.LogFileInit("replica", id)

	s := &Server{
		clientBids:     make(map[int32]int32),
		isAuctionEnded: false,
		totalBids:      0,
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

	if s.totalBids == 5 {
		s.isAuctionEnded = true
		return &protos.StatusOfBid{
			Status:     protos.Status_AUCTION_ENDED,
			HighestBid: s.currentHighestBidsAmount,
		}, nil
	} else {
		s.totalBids += 1

		s.clientBids[req.ClientId] = req.Amount
		database.WriteToCSV(fmt.Sprintf("%d", req.Amount))

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
		ports = append(ports, line)
	}

	return ports, nil
}
