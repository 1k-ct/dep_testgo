FROM golang:latest

COPY src/api /go/src/api

WORKDIR /go/src/api/

RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
#RUN go get github.com/gin-contrib/sessions
#RUN go get github.com/gin-contrib/sessions/cookie
#RUN go get golang.org/x/crypto/bcrypt
#RUN go get google.golang.org/api/googleapi/transport
#RUN go get google.golang.org/api/youtube/v3

RUN apt-get update -qq
RUN apt-get install -y mariadb-client

CMD ["sh", "./start_app.sh"]