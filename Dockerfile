FROM node:12-slim AS client
WORKDIR /client
COPY ./client/ ./
RUN npm install 
RUN npm run build

FROM golang:1.14.2 AS server
WORKDIR /server/
COPY ./server/* ./
RUN go build -o uno .


FROM scratch
WORKDIR /uno
COPY --from=server /server/uno ./uno
COPY --from=client /client/dist/* ./web/
ENV PATH=/uno
CMD ["uno"]