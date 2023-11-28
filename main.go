package main

import (
	"fmt"
	"os"
	"strconv"

	logger "github.com/davrepo/DS_assignment_5_election/logger"
	replica "github.com/davrepo/DS_assignment_5_election/replicamanager"
)




func main() {
	logger.ClearLog("log")
	logger.LogFileInit("main")

	replicaNumber, _ := strconv.Atoi(os.Args[1])

	go replica.Start(int32(replicaNumber), (int32(replicaNumber)))

	bl := make(chan bool)
	<-bl
}

func makePortListForFrontEnd(numberOfReplicas int) {
	logger.ClearLog("replicamanager/portlist")
	logger.MakeLogFolder("replicamanager/portlist")

	f, err := os.OpenFile("replicamanager/portlist/listOfReplicaPorts.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//
	}
	defer f.Close()
	for i := 1; i <= numberOfReplicas; i++ {
		if _, err := f.WriteString(fmt.Sprintf("%v\n", 6000+i)); err != nil {
			//
		}
	}
}
