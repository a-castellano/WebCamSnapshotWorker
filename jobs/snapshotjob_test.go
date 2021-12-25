package jobs

import (
	"testing"
)

func TestEncodeAndDecodeJobs(t *testing.T) {

	var job SnapshotJob
	job.ID = "asdasd1221"
	job.IP = "10.10.12.12"
	test, _ := EncodeJob(job)
	result, _ := DecodeJob(test)

	if result.ID != job.ID {
		t.Errorf(`Encode failed.`)
	}

	if result.IP != job.IP {
		t.Errorf(`Encode failed.`)
	}

}

func TestDecodeEmptyDataJobs(t *testing.T) {

	var emptyData []byte
	_, err := DecodeJob(emptyData)
	if err == nil {
		t.Error("Empty data decoding should fail.")
	}
}
