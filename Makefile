

build:
	mkdir -p ./bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/ytsgo_amd64 main.go
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/ytsgo_arm64 main.go
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/ytsgo_darwin_arm64 main.go
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/ytsgo main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/ytsgo_win.exe main.go

test:
	go test -v ./...

download_movies:
	./bin/ytsgo download --limit 50 --output ./data/movies

search_movies:
	./bin/ytsgo -t bluray -q 1080p list "star wars"

all_movies:
	./bin/ytsgo -t bluray -q 1080p -f list
