FROM golang:1.19

# Define o diretório de trabalho
WORKDIR /go/src

# Configura o PATH e ativa o CGO
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=1

# Instala dependências do Go
RUN go install github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest

# Instala SQLite3 e suas dependências
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev gcc libc-dev build-essential

# Ajusta permissões e define o usuário
RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go /var/www/.cache


# Retorna ao usuário www-data
USER www-data

# Mantém o contêiner ativo
CMD ["tail", "-f", "/dev/null"]