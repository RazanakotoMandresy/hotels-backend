<h1 align="center"> <i>An rest api of an hotels's reservation </i></h1>

## if it's the the first time you want to run the app

#### using the script sh

```bash
$ sh first_running.sh
```

#### manual configuration

<i> export your postgres url </i>

```bash
$ export POSTGRESQL_URL='postgres://postgres:yourpasssword@localhost:5432/the_name_of_the_db?sslmode=disable'
```

<p>  run migration files by executing the command: </p>

```bash
$ migrate -database ${POSTGRESQL_URL} -path migrations up
```

#### To run the backend

<p> run it with the make commandes </p>

```bash
$ make run
```

<h2 align="center"> functionnality </h2>

#### user :

il y a deux type d'utilisateur :

- user proprietaire : les user proprietaire peuvent create , update , delete , read , breff manipuler les hotels .
- user simple : peut voir tous les hotels , et reserver une hotel si il est encore disponnible avec 10% du prix d'avance lors de la reservation , peut noter les hotels , peut annuler les reservation.

#### hotels :

lorsqu'un user proprietaire cree un hotels les fonctionnaliter qu'il peut lui attribuer sont :

- cree les hotels
- y ajouter des images
- mettre a jours les information consernant l'hotels
- supprimer et restaurer les hotels supprimer

du points de vu de l'utilisateur simple une fois connecter:

- voir tous les hotels
- faire des reservation
- filtrer les resultat
