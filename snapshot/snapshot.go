package snapshot

import (
	"crypto/rand"
	"encoding/base32"
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
	return base32.StdEncoding.EncodeToString(randomBytes)[:32], nil
}

func (s SnapshotHandler) DoSnapshot(job jobs.SnapshotJob) (string, error) {

	output, outputerr := generateRandomString()
	if outputerr != nil {
		return output, outputerr
	}

	streamUri := fmt.Sprintf("rtsp://%s:%s@%s:%d%s", job.User, job.Password, job.IP, job.Port, job.SnapshotPath)
	outputFile := fmt.Sprintf("/tmp/%s.jpg", output)

	err := exec.Command(s.ffmpegPath, "-y", "-i", streamUri, "-vframes", "1", outputFile).Run()

	if err != nil {
		return "", err
	}
	return outputFile, nil

}

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
