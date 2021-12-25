package snapshot

import (
	"errors"
	"os"
	"testing"

	jobs "github.com/a-castellano/WebCamSnapshotWorker/jobs"
)

func TestSnapshotMock(t *testing.T) {

	handler := SnapshotMockHandler{ffmpegPath: "foo"}
	emptyJob := jobs.SnapshotJob{}
	snapshotFile, err := handler.DoSnapshot(emptyJob)
	if err != nil {

		t.Errorf("SnapshotMockHandler should not fail. %s", err.Error())
	}
	if _, err := os.Stat(snapshotFile); errors.Is(err, os.ErrNotExist) {
		t.Errorf("SnapshotMockHandler should create file %s but it does not exist.", snapshotFile)
	}

}
