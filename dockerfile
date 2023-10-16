# Estágio de Compilação
FROM golang:1.21.0 AS builder

WORKDIR /app

# Copie o código para o contêiner
COPY . .

# Compile o código
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Estágio da Imagem Final
FROM alpine:3.13

WORKDIR /app

# Copie o binário compilado do estágio de compilação
COPY --from=builder /app/main .

# Comando para executar o aplicativo
CMD ["./main"]
