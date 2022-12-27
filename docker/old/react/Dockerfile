FROM node:16.13.0-alpine
WORKDIR /app/react
RUN npm install -g npm \
  && npm i --save prop-types \
  && npm i -g create-react-app 

COPY ./react/package*.json  ./
RUN npm install