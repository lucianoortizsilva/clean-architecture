#!/bin/sh

while ! nc -z container-mongodb 27017; do
    echo "###############################################"
    echo "######## Aguardando container-mongodb ########"
    echo "###############################################"
    sleep 15
done

/frete-api