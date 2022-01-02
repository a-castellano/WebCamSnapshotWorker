package jobprocessor

import (
	jobs "github.com/a-castellano/WebCamSnapshotWorker/jobs"
	snapshot "github.com/a-castellano/WebCamSnapshotWorker/snapshot"
)

func ProcessJob(data []byte, snapshotHandler snapshot.SnapshotInterface) ([]byte, bool, error) {

	var die bool = false
	receivedJob, _ := jobs.DecodeJob(data)
	var job jobs.SnapshotJob = receivedJob

	if job.ID == "die" {
		die = true
	} else {
		snapshotpath, snatpshotErr := snapshotHandler.DoSnapshot(job)
		if snatpshotErr != nil {
			return data, die, snatpshotErr
		}
		job.SnapshotPath = snapshotpath
	}
	processedJob, _ := jobs.EncodeJob(job)
	return processedJob, die, nil

}
