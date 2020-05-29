# use latest Node LTS slim image
FROM node:lts-slim

RUN mkdir /usr/share/man/man1/
RUN apt-get update && apt-get install -y --no-install-recommends \
    openjdk-8-jdk \
    ant; \
    rm -rf /var/lib/apt/lists/*

# set user to avoid permission issues
# (see https://github.com/nodejs/node-gyp/issues/1236)
USER node
RUN mkdir /home/node/.npm-global
ENV PATH=/home/node/.npm-global/bin:$PATH
ENV NPM_CONFIG_PREFIX=/home/node/.npm-global

# install Firebase CLI
RUN npm install -g firebase-tools

# reset user back to root
# USER root