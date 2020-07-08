FROM node:12-slim

WORKDIR /client

RUN apt-get update && \
    apt-get install -y libgtk2.0-0 \
    libgtk-3-0 \
    libnotify-dev \
    libgconf-2-4 \
    libnss3 \
    libxss1 \
    libasound2 \
    libxtst6 \
    xauth \
    xvfb \
    chromium

RUN npm install -g @vue/cli && \
    npm install -g @vue/cli-service-global && \
    yarn add @vue/cli-service-global

RUN npm install cypress && \
    npm install -g cypress
