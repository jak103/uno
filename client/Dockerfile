FROM node:12-slim

WORKDIR /client

RUN npm install -g @vue/cli

COPY package*.json ./

RUN npm install

RUN npm install vue-confetti --save
