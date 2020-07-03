FROM node:12-slim AS client
WORKDIR /client
ENV NODE_ENV=production
COPY ./client/package*.json ./
RUN npm install 
COPY ./client/ ./
RUN npm run build


FROM golang:1.14.2 AS server
WORKDIR /server/
COPY ./server/* /server/
RUN go build -o uno .

FROM scratch
WORKDIR /uno
COPY --from=server /server/uno ./uno/uno
COPY --from=client /client/dist/* /uno/web/
ENV PATH=/uno
CMD ["uno"]