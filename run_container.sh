#!/bin/bash

docker run -e FROM_EMAIL=andremendes0113@gmail.com \
           -e FROM_PASSWORD="03DDD7F3735A6E0F687997B4F2D89433D469" \
           -e TO_EMAIL=robertaspigolon.aluno@unipampa.edu.br \
           my-mail-sender

echo "A execução do script funcionou!"