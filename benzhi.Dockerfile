FROM docker.m.daocloud.io/library/golang:1.23
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build ./...
ENV GOPROXY=off GOSUMDB=off
CMD ["bash"]
