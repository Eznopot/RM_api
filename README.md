# RM_api
RM_api est l'API rest permettant a l'application web [RM_app](https://github.com/Eznopot/RM_app) de fonctionner.

## Installation
Pour faire fonctionner l'API il faut dans un premier temps installer [golang](https://go.dev/doc/install), ainsi que [MariaDB](https://mariadb.com/kb/en/getting-installing-and-upgrading-mariadb/).

## Base de données
Pour que l'API fonctionne correctement, vous devez creer dans MariaDB une base de données nommée `db_RMS`.

```Bash
mysql -u username -p
CREATE DATABASE db_RMS;
```

Ensuite vous devez y importer le fichier [export1.0.sql](./sql/export1.0.sql) qui se trouve dans le dossier [sql](./sql).

```Bash
mysql -u username -p db_RMS < ./sql/export1.0.sql
```

## Compilation et Execution
Pour pouvoir executer le code de l'API, il faut dans un premier temps installer les dépendances du projet.
```Bash
go get .
```

On peut ensuite lancer l'api en debug mode en utilisant la commande:
```Bash
go run .
```

Pour lancer en production l'API il est necessaire de compiler le code a l'aide de la commande:
```Bash
go build .
```

Puis de lancer le binaire généré:
```Bash
./RM_api
```

## Configuration
La configuration du projet se fait dans le fichier [config.yml](./config.yml).<br>
Il est necessaire de modifier les valeurs user et pass de la section database pour que l'API puisse se connecter a votre base de données.

```YAML
database:
  user: "User"
  pass: "password"
  ip: "localhost"
  port: "3306"
  name: "db_RMS"

logger:
  path: "log"
```

## Developpement
