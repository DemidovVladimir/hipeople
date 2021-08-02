FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /publishfiles

EXPOSE 8080

CMD [ "/publishfiles", "8080" ]