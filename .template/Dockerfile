# syntax=docker/dockerfile:1

FROM golang:1.19-bullseye
WORKDIR /app
COPY . ./
RUN go build .

FROM debian:bullseye-slim
WORKDIR /app
RUN apt-get update -qq && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -yq --no-install-recommends \
    ca-certificates
    
COPY --from=0 /app/PROJECTNAME /app
COPY ./data /app/data

CMD ["/app/PROJECTNAME"]