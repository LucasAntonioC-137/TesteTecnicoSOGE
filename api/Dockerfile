# Usa imagem base oficial do Go
FROM golang:1.23-alpine

# Instala Git (algumas libs precisam)
RUN apk add --no-cache git

# Define a pasta de trabalho
WORKDIR /app

# Copia tudo da pasta api/ para o container
COPY . .

# Instala as dependências
RUN go mod tidy

# Compila o projeto
RUN go build -o main main.go

# Expõe a porta
EXPOSE 5000

# Comando padrão ao rodar o container
CMD ["/app/main"]
