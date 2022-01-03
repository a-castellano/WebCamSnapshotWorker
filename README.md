# WebCamSnapshotWorker

Worker process that takes snapshots from required webcams.

## Install

Add Widmaker repo ans install **windmaker-webcam-snapshot-worker**:
```bash
wget -O - https://packages.windmaker.net/WINDMAKER-GPG-KEY.pub | sudo apt-key add -
sudo add-apt-repository "deb http://packages.windmaker.net/ focal main"
sudo apt-get update
sudo apt-get install windmaker-webcam-snapshot-worker
```

## Configuration

This service uses a config file which folder location is defined by environment variable ** WEBCAM_SNAPSHOT_WORKER_SERVICE_CONFIG_FILE_LOCATION**, inside this folder it must exists a file called **config.toml**.

```toml
[server]
host = "localhost"
port = 5672
user = "guest"
password = "guest"

[incoming]
name = "sendsnapshotjobs"

[outgoing]
name = "outgoing"

[ffmpeg]
path = "/usr/bin/ffmpeg"
```

Config files must include the following sections:
### server

Defines rabbitmq server  config:
* host -> rabbitmq ip or name
* port -> rabbitmq port
* user -> rabbitmq user
* password -> rabbitmq ip or name

### incoming

Define incoming jobs queue names. 
* name -> Queue name

### outgoing

Define outgoing jobs queue names. 
* name -> Queue name

### ffmpeg

Define ffmpeg config. 
* path -> ffmpeg binary path

## Systemd service setup

After saving config in **/etc/windmaker-webcam-snapshot-worker/config.toml** systemd service can be enabled:
```bash
sudo /bin/systemctl daemon-reload
sudo /bin/systemctl enable windmaker-webcam-snapshot-worker
sudo /bin/systemctl start windmaker-webcam-snapshot-worker
```

