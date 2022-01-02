package jobprocessor

import (
	"testing"

	jobs "github.com/a-castellano/WebCamSnapshotWorker/jobs"
	snapshot "github.com/a-castellano/WebCamSnapshotWorker/snapshot"
)

func TestProcessJob(t *testing.T) {

	handler := snapshot.SnapshotMockHandler{ExecPath: "foo"}
	newJob := jobs.SnapshotJob{ID: "asdasd1221", IP: "10.10.12.12", SnapshotPath: "foo"}
	encodedJob, _ := jobs.EncodeJob(newJob)

	resultJob, die, err := ProcessJob(encodedJob, handler)
	if err != nil {
		t.Errorf("decodedJob should not fail: %s", err)
	}

	decodedJob, _ := jobs.DecodeJob(resultJob)
	if decodedJob.SnapshotPath == "foo" {
		t.Errorf("decodedJob SnapshotPath should change after being processed.")
	}

	if die == true {
		t.Errorf("die value should be false.")
	}

}

func TestProcessJobWithDie(t *testing.T) {

	handler := snapshot.SnapshotMockHandler{ExecPath: "foo"}
	newJob := jobs.SnapshotJob{ID: "die", IP: "10.10.12.12", SnapshotPath: "foo"}
	encodedJob, _ := jobs.EncodeJob(newJob)

	resultJob, die, err := ProcessJob(encodedJob, handler)
	if err != nil {
		t.Errorf("decodedJob should not fail: %s", err)
	}

	decodedJob, _ := jobs.DecodeJob(resultJob)
	if decodedJob.SnapshotPath != "foo" {
		t.Errorf("decodedJob SnapshotPath shouldn't change after because of die.")
	}

	if die == false {
		t.Errorf("die value should be true.")
	}

}
