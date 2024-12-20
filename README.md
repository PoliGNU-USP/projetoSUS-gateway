# Gateway

Aqui está o código do gateway do projeto SUSBot.

Este serviço é responsável por receber, tratar, armazenar as mensagens vindas do
WhatsApp e também enviar para o Botkit e encaminhar a resposta ao usuário.

## Estou desenvolvendo, como rodar o chatbot na minha maquina para testes?

### Configurando Twilio + Ngrok

A [Twilio](https://www.twilio.com/en-us) é o nosso broker de Whatsapp. Ela que
intermedia as mensagens entre nós e o usuário.

O [Ngrok](https://ngrok.com) é útil durante o desenvolvimento, ele que expõe
o nosso gateway para a internet o que nos possibilita
receber POST requests no gateway.

Crie sua conta nos dois serviços.

Na sua conta twilio, vá na aba de enviar uma mensagem de teste (algo como
"try it out") e adicione seu número de celular no sandbox deles.

### Rodando com o Docker

Você precisa ter o [Docker](https://www.docker.com) instalado em seu computador.

Defina o [seu .env](#adicione-em-um-arquivo-env-as-seguintes-informações) e deixe
na pasta base do projeto (ao lado da main.go)

No diretório **remote** use o comando:

```shell
docker-compose up -d -build
```

Depois, vá no seu navegador e entre em **localhost:4040** e copie o link
fornecido pelo NGROK.

Ele é longo, parecido com "<https://XXXX-XXX-X-XXX-XXX.ngrok-free.app>"

Na Twilio, entre em "Develop" depois "Messaging" e "Try it out" e,
por fim, "Send a Whatsapp message"

Na aba sandbox adicione o seu celular no sandbox do bot.

Pegue o link fornecido no ngrok e adicione no seu
sandbox settings na campo "when a message comes in" e defina o método para post

Agora teste! Mande mensagem pelo seu celular e veja o SUSBot vivo!

## Adicione em um arquivo *.env* as seguintes informações

\# -- Seu Token de Autorização do Ngrok --  
**NGROK_AUTHTOKEN=XXXXXX**
