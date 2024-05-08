# syntax=docker/dockerfile:1
  
FROM golang:1.20 as stage-build

WORKDIR /watchtower/src/client/

COPY . ./
RUN go mod download


RUN go build -o watchtower .
RUN chmod +x ./rundocker.sh

FROM debian:stable 
WORKDIR /watchtower/src/client

RUN mkdir certs
COPY certs/ certs/.
COPY generate-certificate.sh .
COPY run .
RUN touch config.json
COPY rundocker.sh .
COPY LICENSE /
RUN apt update && apt install certbot curl -y
RUN chmod +x ./rundocker.sh
COPY --from=stage-build /watchtower/src/client/watchtower .
CMD ["./rundocker.sh"]
