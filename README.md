# ChatBot - UBS -- Gateway

<!--toc:start-->
- [ChatBot - UBS -- Gateway](#chatbot-ubs-gateway)
  - [Adicione em um arquivo *.env* as seguintes informações](#adicione-em-um-arquivo-env-as-seguintes-informações)
<!--toc:end-->

Essa é a parte do bot responsável por receber, tratar, salvar, enviar para o
fluxo de conversa e mandar de volta a mensagem para o usuário.

Ela também deve enviar ao banco de dados as mensagenm.

Para rodar o gateway você precisa ter o Go 1.22+ instalado em sua máquina.

Para rodar os testes (ambiente de desenvolvimento) você precisará de uma conta
na NGROK e na TWILIO

Na sua conta twilio, vá na aba de enviar uma mensagem de teste (algo como "try it out") e adicione seu número de celular no sandbox deles.

Gere o link do ngrok na mesma porta do DEV_LOCALHOST.
se for 8080 será -> ngrok http 8080

Depois, coloque o link do ngrok endpoint do sandbox.

## Adicione em um arquivo *.env* as seguintes informações

-- Seu secure ID da Twilio. --  
**TWILIO_ACCOUNT_SID=XXXXXX**

-- Seu Token de Autorização da Twilio --  
**TWILIO_AUTH_TOKEN=XXXXXX**

-- O LocalHost que você quer colocar (geralmente :8080) --  
-- Essa porta também deve ser a mesma da porta que você precisará colocar no ngrok (para expor a sua porta local na internet)--
**DEV_LOCALHOST=XXXXXX**

-- Definindo o BOTKIT Webhook (geralmente <http://localhost:3000/api/messages)-->  
**BOTKIT_WEBHOOK=XXXXXX**
