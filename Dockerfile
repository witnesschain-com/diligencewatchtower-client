# syntax=docker/dockerfile:1
  
FROM golang:1.22.1 as stage-build

ARG WATCHTOWER_VERSION
WORKDIR /watchtower

COPY . .
RUN ls
RUN go mod download

WORKDIR /watchtower/cmd/watchtower
RUN go build -o watchtower -ldflags="-X 'main.VERSION=${WATCHTOWER_VERSION}'" . 
RUN echo "Completed build1"


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
COPY --from=stage-build /watchtower/cmd/watchtower/watchtower .
RUN echo "Completed build2"
CMD ["./rundocker.sh"]