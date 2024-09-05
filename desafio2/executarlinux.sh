#!/bin/bash

# Nome da imagem Docker
IMAGE_NAME="desafio-resolvedor"

# Passo 1: Build da imagem Docker
echo "Construindo a imagem Docker..."
docker build -t $IMAGE_NAME .

# Passo 2: Executar o container Docker
echo "Executando o container Docker..."
docker run $IMAGE_NAME

echo "Desafio resolvido automaticamente!"
