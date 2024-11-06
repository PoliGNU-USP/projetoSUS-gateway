FROM node:alpine

WORKDIR /app

RUN apk add git

RUN git clone --progress https://github.com/PoliGNU-USP/projetoSUS-fluxo.git .

RUN npm install -g nodemon

RUN npm install

EXPOSE 3000

CMD ["npm", "start"]
