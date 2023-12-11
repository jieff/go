# Use a imagem Alpine do Go como base
FROM golang:alpine

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie apenas os arquivos necessários para o diretório de trabalho no contêiner
COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod download

# Copie os arquivos restantes para o diretório de trabalho no contêiner
COPY . .

# Comando para compilar o arquivo main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Comando padrão a ser executado quando o contêiner for iniciado
CMD ["./main"]

