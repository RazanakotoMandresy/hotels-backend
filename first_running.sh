#!/bin/bash
echo -e "
welcome to the first running of hotels backend 🏨 \n
if you have some issues please contact RazanakotoMandresy 🗿 😉 \n
this script is  used if it's the first time you want to run hotel's api in your pc 🖥️💻\n
first let's configure the database ⚙️  \n
enter your name in your postgres db \n 
press on "d" for the default postgres username \n"
read userDb
echo "enter your  postgres db passwords"
read dbPassword
echo "enter the name of the db you want"
read dbName
if [ $userDb == "d" ]; then
    echo "default userDb name"
    export POSTGRESQL_URL='postgres://postgres:'"$dbPassword"'@localhost:5432/hotels?sslmode=disable'
else
    echo -e " user's db name:$userDb , postgres's db name:$dbName "
    export POSTGRESQL_URL='postgres://'"$userDb"":$dbPassword"'@localhost:5432/'"$dbName"'?sslmode=disable'
fi
migrate -database ${POSTGRESQL_URL} -path migrations up
make run
