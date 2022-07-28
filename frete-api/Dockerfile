FROM golang:1.18.1-alpine3.14

LABEL authors="lucianoortizsilva@gmail.com"

#Criando um diretório chamado 'app'
#Isso instrui o Docker a usar esse diretório como destino padrão para todos os comandos subsequentes. 
#Dessa forma, não precisamos digitar caminhos de arquivo completos, mas podemos usar caminhos relativos com base nesse diretório.
WORKDIR /app

#O comando COPY  recebe dois parâmetros. 
#O primeiro parâmetro informa ao Docker quais arquivos você deseja copiar para a imagem. 
#O último parâmetro informa ao Docker para onde você deseja que esse arquivo seja copiado.
COPY go.mod /app
COPY go.sum /app
COPY /src /app/src
COPY .env /app
COPY main.go /app
COPY entrypoint.sh /app

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