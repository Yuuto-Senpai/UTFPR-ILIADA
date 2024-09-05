@echo off
echo Construindo a imagem Docker...
docker build -t desafio-resolvedor .

echo Executando o container Docker...
docker run desafio-resolvedor

echo Desafio resolvido automaticamente!
pause

REM Path: CPQ-UTFPR/desafio2/executar%20windows.bat