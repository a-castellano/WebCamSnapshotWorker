package snapshot

import (
	"fmt"
	"os/exec"

	jobs "github.com/a-castellano/WebCamSnapshotWorker/jobs"
)

type SnapshotMockHandler struct {
	ffmpegPath string
}

func (s SnapshotMockHandler) DoSnapshot(job jobs.SnapshotJob) (string, error) {

	output, outputerr := generateRandomString()
	if outputerr != nil {
		return output, outputerr
	}

	outputFile := fmt.Sprintf("/tmp/%s.jpg", output)
	err := exec.Command("/usr/bin/touch", outputFile).Run()

	if err != nil {
		return "", err
	}
	return outputFile, nil

}
