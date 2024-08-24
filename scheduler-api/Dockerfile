FROM node:22

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY . .

ENTRYPOINT ["npm", "run", "start"]
