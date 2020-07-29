FROM node:12-slim AS client
ENV NODE_ENV production
WORKDIR /client
COPY ./client/ ./
RUN npm install 
RUN npm run build

FROM golang:1.14.2 AS server
WORKDIR /server/
COPY ./server ./
#RUN go build -o uno .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o uno .

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /uno
COPY --from=server /server/uno /uno/uno
COPY --from=client /client/dist /client/dist
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD ["/uno/uno"]