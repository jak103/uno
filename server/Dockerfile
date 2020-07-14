# use latest Node LTS slim image
FROM golang:1.14

# install gin
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

RUN go get github.com/codegangsta/gin

ENV GOOGLE_APPLICATION_CREDENTIALS="/server/USUDevOps-5ed3f59927e8.json"
