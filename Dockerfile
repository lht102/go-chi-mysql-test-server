FROM golang:alpine

WORKDIR /app

COPY . .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait

RUN chmod +x /wait

RUN go build -o main .

CMD /wait && ./main
