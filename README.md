# DO IT

1. search your machine ip locally for that prometheus container can to access the app url
2. change prometheus/prometheys.yml file in scrape_configs.static_configs.targets with your machine ip and port of the app

#### execute

``` 
docker-compose up -d
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp
go run cmd\main.go
```