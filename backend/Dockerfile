FROM golang:1.23rc2-alpine AS builder

WORKDIR /app
COPY go.* ./

RUN apk -U --no-cache add ca-certificates openssl
ARG cert_location=/usr/local/share/ca-certificates
COPY ./mitmproxy.crt ${cert_location}/mitmproxy.crt
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt 
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/proxy.golang.crt 
RUN update-ca-certificates

RUN go mod download

COPY . ./
RUN go build -v -o server

EXPOSE 8000

CMD ["/app/server"]