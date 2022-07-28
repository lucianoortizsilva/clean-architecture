#!/bin/sh

while ! nc -z container-mongodb 27017; do
    echo "###############################################"
    echo "######## Aguardando container-mongodb ########"
    echo "###############################################"
    sleep 30
done

/frete-api