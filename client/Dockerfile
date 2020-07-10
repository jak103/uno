FROM node:12-slim

WORKDIR /client

RUN mkdir /client/node_modules && \
    npm install -g @vue/cli && \
    npm install -g @vue/cli-service-global && \
    yarn add @vue/cli-service-global

COPY package*.json ./

RUN npm install
