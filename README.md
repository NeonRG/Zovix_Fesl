[![HitCount](http://hits.dwyl.io/Synaxis/bfheroesFesl.svg)](http://hits.dwyl.io/Synaxis/bfheroesFesl)
# Open Heroes Backend

# Setting up
Remember to configure GOPATH to match your github directory.

# Configuration
Enviroment (.env) variables You can look in `./config/config.go` for more details

| String Name          |Default value        |
|--------------------- |---------------------|
| `LOG_LEVEL`          |`INFO`               |
| `HTTP_BIND`          |`0.0.0.0:8080`       |
| `HTTPS_BIND`         |`0.0.0.0:443`        |
| `GAMESPY_IP`         |`0.0.0.0`(auto bind) |
| `THEATER_ADDR`       |`127.0.0.1`          |
| `LEVEL_DB_PATH`      |`_data/lvl.db`       |
| `DATABASE_USERNAME`  |`root`               |
| `DATABASE_PASSWORD`  |                     |
| `DATABASE_HOST`      |`127.0.0.1`          |
| `DATABASE_PORT`      |`3306`               |
| `DATABASE_NAME`      |`tutorialDB`         |

## Example `.env` file
```ini
DATABASE_NAME=tutorialDB
DATABASE_HOST=127.0.0.1
DATABASE_PASSWORD=dbPass
LOG_LEVEL=DEBUG
```
#Setup gameServer

1- Download the gameServer from here

open Hxd.exe and put BFheroes_w32ded.exe inside. Now search for 127.0.0.1. => put your backend ip in there(or don't change it for localhost)(same method for VPS's) SAVE IT.
2- Create a text file and put this inside it 

