FROM node:alpine

WORKDIR /app

RUN apk add git=2.47.1-r0 --no-cache

RUN git clone --progress https://github.com/PoliGNU-USP/projetoSUS-fluxo.git . &&\
  npm install -g nodemon@latest &&\
  npm install

CMD ["npm", "start"]

