# crypto-prices-web-app
A web app to track crypto prices.

![screenshot](/.github/capture.png)

It spawns a thread to update price data every 5 minutes using [CoinGecko](https://www.coingecko.com/en/api/documentation)'s API.

The template is executed once after each price data update. 

HTML content is held in memory, so users can instantly see the content.


## Setup Development Environment
This repo uses [Dev Containers](https://containers.dev/) to setup development environments easily.

## Generate `index.css` file
Run the command below to generate `index.css` file.
```sh
tailwindcss -o index.css -m --content template.html
```

## Start Server
Run the command below to start the server.
```sh
go run main.go
```
Now you can visit [localhost:8080/](http://localhost:8080/) to see the result.