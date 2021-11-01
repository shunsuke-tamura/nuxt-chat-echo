FROM node:16-alpine

LABEL maintainer=shunsuke

ENV HOST 0.0.0.0

WORKDIR /app
# ADD . /app

RUN npm install -g npm@8.1.2

CMD PORT=8080 npm run dev
