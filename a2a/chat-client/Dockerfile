FROM node:24-alpine
RUN mkdir -p /opt/app

WORKDIR /opt/app

COPY . .

RUN npm i

EXPOSE 3000

CMD [ "npm", "start" ]
