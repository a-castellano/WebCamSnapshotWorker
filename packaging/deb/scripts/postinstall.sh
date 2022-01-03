#!/bin/sh

mkdir -p /etc/windmaker-webcam-snapshot-worker

echo "### NOT starting on installation, please execute the following statements to configure windmaker-webcam-snapshot-worker to start automatically using systemd"
echo " sudo /bin/systemctl daemon-reload"
echo " sudo /bin/systemctl enable windmaker-webcam-snapshot-worker"
echo "### You can start windmaker-webcam-snapshot-worker by executing"
echo " sudo /bin/systemctl start windmaker-webcam-snapshot-worker"
