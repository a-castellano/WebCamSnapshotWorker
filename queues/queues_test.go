// +build integration_tests

package queues

import (
	"log"
	"testing"

	config "github.com/a-castellano/WebCamSnapshotWorker/config_reader"
	jobs "github.com/a-castellano/WebCamSnapshotWorker/jobs"
	snapshot "github.com/a-castellano/WebCamSnapshotWorker/snapshot"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func TestSendDie(t *testing.T) {

	var queueConfig config.Config

	queueConfig.Server.Host = "rabbitmq"
	queueConfig.Server.Port = 5672
	queueConfig.Server.User = "guest"
	queueConfig.Server.Password = "guest"

	queueConfig.Incoming.Name = "incoming"
	queueConfig.Outgoing.Name = "outgoing"

	newJob := jobs.SnapshotJob{ID: "die", Errored: false, Finished: false, IP: "10.10.10.10", User: "user", Password: "password", Port: 80}
	encodedNewJob, _ := jobs.EncodeJob(newJob)

	handler := snapshot.SnapshotMockHandler{ExecPath: "foo"}

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel in TestSendDie")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"incoming", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue in TestSendDie")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         encodedNewJob,
		})

	jobManagementError := StartJobManagement(queueConfig, handler)
	if jobManagementError != nil {
		t.Errorf("StartJobManagement should return no errors when die is processed.")
	}

}
