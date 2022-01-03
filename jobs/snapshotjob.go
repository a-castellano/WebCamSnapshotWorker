package jobs

import (
	"bytes"
	"encoding/gob"
)

type SnapshotJob struct {
	ID           string `json:"id"`
	Errored      bool   `json:"errored"`
	Finished     bool   `json:"finished"`
	SnapshotPath string `json:"snapshotpath"`
	IP           string `json:"ip"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Port         int    `json:"port"`
	StreamPath   string `json:"streampath"`
	Sender       int    `json:"sender"`
}

func EncodeJob(job SnapshotJob) ([]byte, error) {
	var encodedJob []byte
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(job)
	if err != nil {
		return encodedJob, err
	}
	encodedJob = network.Bytes()
	return encodedJob, nil
}

func DecodeJob(encoded []byte) (SnapshotJob, error) {
	var job SnapshotJob
	network := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&job)
	if err != nil {
		return SnapshotJob{}, err
	}
	return job, nil
}
