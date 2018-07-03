# Run locally
```
# go build
# ./icanhazip-go
```

# Run locally with docker
```
# docker build -t icanhazip-go .
# docker run -e EXTERNAL_WHATISMYIP_SERVICE_HOST=ipv4.icanhazip.com -e EXTERNAL_WHATISMYIP_SERVICE_PORT=80 --rm -p 8080:8080 icanhazip-go
```
