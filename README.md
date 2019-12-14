# GoWeb - Searching Job in NAB

## System requirements

- Go 1.12
- Elasticsearch 6.8 (Important)
- Dep

## Install Dependencies

Run `dep ensure`

## Distributed Architecture
![architecture](https://res.cloudinary.com/deey9oou3/image/upload/v1576292348/Crawler.png)

## Before You Run

Make sure Elasticsearch is running on your localhost port 9200

## Run Simple Version

WebSpider: `go run main.go`

## Run Distributed Version

ItemSaver:  `go run /crawler_distributed/itemsaver.go --port 8080` (could be any int)

Worker:  `go run /crawler_distributed/worker.go --port 8080` (could be any int)

Client: `go run main.go --itemSaverHost ":8080" --workerHosts ":9000,:90001"` (ports should be string with ":" prefix)

## View Job list

Run Webspider first then run frontend to check joblist

Frontend: `go run frontend/starter.go`

visit http://localhost:8888/search