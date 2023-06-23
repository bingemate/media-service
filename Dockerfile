# build stage
FROM golang:1.20 AS build

ENV GO111MODULE=on

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -x -ldflags "-s -w" -o main .

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/
COPY --from=build /app/main .

# Define your environment variables here
ENV TZ=Europe/Paris \
    PORT=8080 \
    LOG_FILE=/app/logs/golang-app.log \
    MOVIE_TARGET_FOLDER=/app/media-target \
    TV_TARGET_FOLDER=/app/media-target \
    TMDB_API_KEY="" \
    DB_SYNC=true \
    DB_HOST=127.0.0.1 \
    DB_PORT=5432 \
    DB_USER=bingemate \
    DB_PASSWORD=bingemate \
    DB_NAME=bingemate \
    S3_ENDPOINT="http://localhost:9000" \
    S3_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxxxxx \
    S3_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx \
    S3_BUCKET_NAME=media \
    REDIS_HOST="localhost:6379" \
    REDIS_PASSWORD=""

# Expose the port on which the application will listen
EXPOSE $PORT

VOLUME /var/logs/app \
         /app/media-target

USER 1000:100

# Start the application
ENTRYPOINT ["/app/main"]
