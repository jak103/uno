FROM node:12-slim

RUN npm install -g @vue/cli
RUN npm install -g @vue/cli-service-global
RUN yarn add @vue/cli-service-global

COPY package.json .
RUN npm install