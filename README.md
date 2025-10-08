# YTS Go Client
Go library to access YTS.mx API

API documentation: https://yts.mx/api


## CLI Example
```sh
  ./ytsgo -t bluray -q 1080p list "star wars"
````

Output
```output
"Lego Star Wars: The Padawan Menace" (2011)
        Quality: 1080p, Type: bluray, Seeds: 17 Peers: 1 Size: 416.65 MB
"Star Wars: Star Warriors" (2007)
        Quality: 1080p, Type: bluray, Seeds: 8 Peers: 4 Size: 1.4 GB
"Star Wars: The Last Jedi Cast Live Q&A" (2017)
        Quality: 1080p, Type: bluray, Seeds: 6 Peers: 1 Size: 2.8 GB
"Rogue One: A Star Wars Story" (2016)
        Quality: 1080p, Type: bluray, Seeds: 70 Peers: 0 Size: 2.47 GB
"Star Wars: Episode I - The Phantom Menace" (1999)
        Quality: 1080p, Type: bluray, Seeds: 53 Peers: 8 Size: 2.51 GB
"Robot Chicken: Star Wars III" (2010)
        Quality: 1080p, Type: bluray, Seeds: 25 Peers: 6 Size: 850.59 MB
"Zoom In Solo A Star Wars Story" (2018)
```

## Usage Example
```go
  c, err := ytsgo.New()
  if err != nil {
    log.Fatalf("Failed to create ytsgo client: %v", err)
  }
  m, err := c.Movie(10)
  if err != nil {
    log.Fatalf("Failed to fetch movie id:%v :%v", id, err)
  }
  fmt.Println(m.Title)
```

[ytsgo.go](https://github.com/onyxg/ytsgo/blob/master/cmd/ytsgo.go) is a simple CLI tool to fetch and search movies from YTS.MX and show manget links to all torrents.


