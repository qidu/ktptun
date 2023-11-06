# build proxy client and server
/usr/local/go/bin/go build -o proxy_server server/main.go server/config.go server/signal.go
/usr/local/go/bin/go build -o proxy_client client/main.go client/config.go client/dial.go client/signal.go

# test
# ./proxy_server -t 127.0.0.1:80 --key 1234
# ./proxy_client -l :7890 -r 127.0.0.1:29900 --key 1234
# curl http://localhost:7890/
