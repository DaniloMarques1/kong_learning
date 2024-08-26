FROM node:22

WORKDIR /app

COPY package.json .
COPY package-lock.json .

COPY . .

RUN npm install

ENV QUEUE_URL="amqp://fitz:fitz@rabbitmq:5672"

ENTRYPOINT ["npm", "run", "start"]
