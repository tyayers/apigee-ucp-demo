FROM node:24-alpine
RUN mkdir -p /opt/app

WORKDIR /opt/app

COPY . .

RUN npm i
RUN npm run build

EXPOSE 8080

CMD [ "node", "server.ts" ]
