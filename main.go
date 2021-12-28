package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"

	config "github.com/a-castellano/WebCamSnapshotWorker/config_reader"
	queues "github.com/a-castellano/WebCamSnapshotWorker/queues"
	"github.com/a-castellano/WebCamSnapshotWorker/snapshot"
)

func main() {

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "security-cam-bot")
	if e == nil {
		log.SetOutput(logwriter)
		// Remove date prefix
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	}

	log.Println("Reading config")
	serviceConfig, err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		log.Println("Config readed successfully.")

		snapshotHandler := snapshot.SnapshotHandler{ExecPath: serviceConfig.FfmpegPath}

		jobManagementError := queues.StartJobManagement(serviceConfig, snapshotHandler)

		if jobManagementError != nil {
			fmt.Println(jobManagementError)
		}
	}

}
