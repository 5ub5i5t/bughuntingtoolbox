FROM node:20 AS builder

WORKDIR /app
COPY package.json /app/
COPY . /app/
RUN npm install
RUN npm run build

EXPOSE 3000

RUN cd /app

CMD ["npm", "start"]