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
Here is all the route of the API:

```Bash
No rights route
POST   /api/test                 --> github.com/Eznopot/RM_api/src/Function.Test (5 handlers)
POST   /user/login               --> github.com/Eznopot/RM_api/src/Function.Login (5 handlers)
GET    /getDayOff                --> github.com/Eznopot/RM_api/src/Function.GetDayOff (5 handlers)
GET    /calendar/getEventTypes   --> github.com/Eznopot/RM_api/src/Function.GetEventTypes (5 handlers)
GET    /calendar/getOtherEventTypes --> github.com/Eznopot/RM_api/src/Function.GetOtherEventTypes (5 handlers)
GET    /calendar/getAbsenceEventTypes --> github.com/Eznopot/RM_api/src/Function.GetAbsenceEventTypes (5 handlers)
DELETE /user/logout              --> github.com/Eznopot/RM_api/src/Function.Logout (5 handlers)
POST   /candidat/add             --> github.com/Eznopot/RM_api/src/Function.AddCandidat (5 handlers)
GET    /info/roles               --> github.com/Eznopot/RM_api/src/Function.GetRoles (5 handlers)

User route
GET    /session/user/getInfo     --> github.com/Eznopot/RM_api/src/Function.GetInfo (6 handlers)
GET    /session/user/getPages    --> github.com/Eznopot/RM_api/src/Function.GetPages (6 handlers)
GET    /session/user/checkSession --> github.com/Eznopot/RM_api/src/Function.CheckSession (6 handlers)
POST   /session/user/addCalendarEvent --> github.com/Eznopot/RM_api/src/Function.AddCalendarEvent (6 handlers)
GET    /session/user/getCalendarEvents --> github.com/Eznopot/RM_api/src/Function.GetCalendarEvents (6 handlers)
POST   /session/user/autoPresenceCalendarEvents --> github.com/Eznopot/RM_api/src/Function.AutoPresenceCalendarEvents (6 handlers)
POST   /session/user/modifyCalendarEvent --> github.com/Eznopot/RM_api/src/Function.ModifyCalendarEvent (6 handlers)
POST   /session/user/deleteCalendarEvent --> github.com/Eznopot/RM_api/src/Function.DeleteCalendarEvent (6 handlers)
GET    /session/user/getHollidayRequest --> github.com/Eznopot/RM_api/src/Function.GetHollidayRequest (6 handlers)
POST   /session/user/addHollidayRequest --> github.com/Eznopot/RM_api/src/Function.AddHollidayRequest (6 handlers)
POST   /session/user/deleteHollidayRequest --> github.com/Eznopot/RM_api/src/Function.DeleteHollidayRequest (6 handlers)
GET    /session/user/getOwnCRAM  --> github.com/Eznopot/RM_api/src/Function.GetOwnCram (6 handlers)
GET    /session/info/get         --> github.com/Eznopot/RM_api/src/Function.GetAdminString (6 handlers)

Manager route
GET    /sessionManager/holliday/getAllHollidayRequest --> github.com/Eznopot/RM_api/src/Function.GetAllHollidayRequest (6 handlers)
POST   /sessionManager/holliday/declineHollidayRequest --> github.com/Eznopot/RM_api/src/Function.DeclineHollidayRequest (6 handlers)
POST   /sessionManager/holliday/acceptHollidayRequest --> github.com/Eznopot/RM_api/src/Function.AcceptHollidayRequest (6 handlers)
POST   /sessionManager/holliday/deleteOtherHollidayRequest --> github.com/Eznopot/RM_api/src/Function.DeleteOtherHollidayRequest (6 handlers)
GET    /sessionManager/candidat/search --> github.com/Eznopot/RM_api/src/Function.SearchCandidat (6 handlers)
GET    /sessionManager/candidat/searchByEmail --> github.com/Eznopot/RM_api/src/Function.SearchCandidatByEmail (6 handlers)
GET    /sessionManager/candidat/loadSome --> github.com/Eznopot/RM_api/src/Function.LoadSomeCandidat (6 handlers)
GET    /sessionManager/RDV/getAll --> github.com/Eznopot/RM_api/src/Function.GetRDVEvent (6 handlers)
POST   /sessionManager/RDV/add   --> github.com/Eznopot/RM_api/src/Function.AddRDVEvent (6 handlers)
POST   /sessionManager/RDV/modify --> github.com/Eznopot/RM_api/src/Function.ModifyRDVEvent (6 handlers)
POST   /sessionManager/RDV/delete --> github.com/Eznopot/RM_api/src/Function.DeleteRDVEvent (6 handlers)
POST   /sessionManager/RDV/saveNote --> github.com/Eznopot/RM_api/src/Function.InsertAppreciation (6 handlers)

Admin route
GET    /sessionPlus/user/getAllUser --> github.com/Eznopot/RM_api/src/Function.GetAllUser (6 handlers)
POST   /sessionPlus/user/updateRole --> github.com/Eznopot/RM_api/src/Function.UpdateRole (6 handlers)
POST   /sessionPlus/user/register --> github.com/Eznopot/RM_api/src/Function.Register (6 handlers)
POST   /sessionPlus/info/add     --> github.com/Eznopot/RM_api/src/Function.AddAdminString (6 handlers)
POST   /sessionPlus/info/modify  --> github.com/Eznopot/RM_api/src/Function.ModifyAdminString (6 handlers)
POST   /sessionPlus/info/delete  --> github.com/Eznopot/RM_api/src/Function.DeleteAdminString (6 handlers)
GET    /sessionPlus/calendar/getAllCRAM --> github.com/Eznopot/RM_api/src/Function.GetAllCram (6 handlers)
GET    /sessionPlus/user/getCRAMByEmail --> github.com/Eznopot/RM_api/src/Function.GetCrambyEmail (6 handlers)
GET    /sessionPlus/user/cv/get  --> github.com/Eznopot/RM_api/src/Function.GetUserCv (6 handlers)
```