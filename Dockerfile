FROM golang

WORKDIR /usr/app

COPY go.mod /Users/eugene/Projects/http-server/

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ./main