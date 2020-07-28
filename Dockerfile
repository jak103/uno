FROM node:12-slim AS client
WORKDIR /client
COPY ./client/ ./
RUN npm install 
RUN npm run build

FROM golang:1.14.2 AS server
WORKDIR /server/
COPY ./server ./
#RUN go build -o uno .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o uno .

FROM scratch
WORKDIR /uno
COPY --from=server /server/uno /uno/uno
COPY --from=client /client/dist /client/dist
#ENV PATH=$PATH:/uno
CMD ["/uno/uno"]