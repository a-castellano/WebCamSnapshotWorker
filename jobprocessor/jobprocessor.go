package jobprocessor

import (
	jobs "github.com/a-castellano/WebCamSnapshotWorker/jobs"
	snapshot "github.com/a-castellano/WebCamSnapshotWorker/snapshot"
)

func ProcessJob(data []byte, snapshotHandler snapshot.SnapshotInterface) ([]byte, error) {

	receivedJob, _ := jobs.DecodeJob(data)
	var job jobs.SnapshotJob = receivedJob

	snapshotpath, snatpshotErr := snapshotHandler.DoSnapshot(job)
	if snatpshotErr != nil {
		return data, snatpshotErr
	}
	job.SnapshotPath = snapshotpath
	processedJob, _ := jobs.EncodeJob(job)
	return processedJob, nil

}
