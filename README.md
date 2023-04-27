# О приложении
TODO

# Развертывание 
В этом разделе описано как можно самостоятельно развернуть все компоненты системы по отдельности

### Развертывание БД
Для развертывания БД с нуля, достаточно запустить скрипт `db/check_install.sh`.

Предварительно следует изменить параметры определенные по дефолту внутри этого скрипта:
```txt
6  |  # Parameters
7  |  HOST='192.168.0.108'
8  |  PORT='5432'
9  |  POSTGRES='postgres'
10 |  POSTGRES_PASSWORD='postgres'
11 |  DB_ADMIN='mp_admin'
12 |  DB_ADMIN_PASSWORD='mp_admin'
13 |  DB_NAME="mind_palace"
```
После исполнения скрипта, логи его работы будут записаны в отдельный файл по пути `db/log/*`