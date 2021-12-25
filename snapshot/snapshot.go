package snapshot

import (
	"crypto/rand"
	"fmt"
	"os/exec"

	jobs "github.com/a-castellano/WebCamSnapshotWorker/jobs"
)

type SnapshotInterface interface {
	DoSnapshot(job jobs.SnapshotJob) (string, error)
}

type SnapshotHandler struct {
	ffmpegPath string
}

func generateRandomString() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return string(randomBytes), nil
}

func (s SnapshotHandler) DoSnapshot(job jobs.SnapshotJob) (string, error) {

	output, outputerr := generateRandomString()
	if outputerr != nil {
		return output, outputerr
	}

	streamUrl := fmt.Sprintf("rtsp://%s:%s@%s:%d%s", job.User, job.Password, job.IP, job.Port, job.SnapshotPath)
	outputFile := fmt.Sprintf("/tmp/%s.jpg", output)
	snapshotCommand := fmt.Sprintf("%s -y -i %s -vframes 1 %s", s.ffmpegPath, streamUrl, outputFile)
	_, err := exec.Command(snapshotCommand).Output()

	if err != nil {
		return "", err
	}
	return outputFile, nil

}
