FROM golang:latest as builder

WORKDIR /dbackupdul

ENV DB=
ENV DB_USERNAME=
ENV DB_PASSWORD=
ENV DB_HOST=
ENV DB_PORT=
ENV DB_NAME=
ENV PATH_BACKUP=
ENV SCHEDULE=

COPY . .
RUN go mod download
RUN go build -o /dbackupdul/dbackupdul main.go

FROM ubuntu:latest
WORKDIR /app

ENV DB=${DB}
ENV DB_USERNAME=${DB_USERNAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_NAME=${DB_NAME}
ENV PATH_BACKUP=${PATH_BACKUP}
ENV SCHEDULE=${SCHEDULE}

EXPOSE ${DB}
EXPOSE ${DB_USERNAME}
EXPOSE ${DB_PASSWORD}
EXPOSE ${DB_HOST}
EXPOSE ${DB_PORT}
EXPOSE ${DB_NAME}
EXPOSE ${PATH_BACKUP}
EXPOSE ${SCHEDULE}

RUN apt-get update && apt-get install mariadb-client -y && apt-get install postgresql-client -y

COPY --from=builder /dbackupdul/dbackupdul .

ENTRYPOINT ["./dbackupdul"]


