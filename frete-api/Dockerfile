FROM golang:1.18.1-alpine3.14

LABEL authors="lucianoortizsilva@gmail.com"

WORKDIR /app

COPY /adapters /app/adapters
COPY /drivers /app/drivers
COPY /entities /app/entities
COPY /usecases /app/usecases

COPY entrypoint.sh /app
COPY go.mod /app
COPY go.sum /app
COPY main.go /app

#Para instalar os módulos Go em um diretório dentro da imagem
RUN go mod download

#Criando um arquivo binário com nome de 'frete-api'
RUN go build -o /frete-api

#Porta
EXPOSE 5000

#Informando qual arquivo executar ao iniciar im container
#CMD ["/frete-api"], comentado para executar no entrypoint.sh
# Comando para criar imagem:
# docker build --tag frete-api .