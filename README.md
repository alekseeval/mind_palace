# О приложении

### Основные компоненты системы:

- CRUD HTTP-Server
- Telegram bot (TODO)
- БД в PostgreSQL

### Взаимодействие компонентов
TODO

# Развертывание БД
В этом разделе описано как можно самостоятельно развернуть PostgreSQL БД для проекта

### Развертывание БД запуском скрипта
Для развертывания БД с нуля, достаточно запустить скрипт `db/check_install.sh`. Например, находясь в корне проекта, командой:
```bash
bash db/check_install.sh
```

Предварительно следует передать в скрипт параметры подключения к базе и настройки admin-пользователя.
Сделать это можно изменив непосредственно дефолтные параметры в тексте скрипта:
```txt
15 |  host='192.168.0.108'
16 |  port='5432'
17 |  postgres='postgres'
18 |  postgres_password='postgres'
19 |  db_admin='mp_admin'
20 |  db_admin_password='mp_admin'
21 |  db_name="mind_palace"
```
Либо можно задать соответствующие переменные окружения (upper case) и запустить скрипт:
```bash
export HOST='192.168.0.108'
export PORT='5432'
export POSTGRES='postgres'
export POSTGRES_PASSWORD='postgres'
export DB_ADMIN='mp_admin'
export DB_ADMIN_PASSWORD='mp_admin'
export DB_NAME="mind_palace"
```

> Переменные окружения, если они установлены, считаются более приоритетными

После исполнения скрипта, логи его работы будут записаны в отдельный файл по пути `db/log/*`

### Развертывание БД через Docker-образ

БД можно развернуть запустив Docker-контейнер и передав в него переменные окружения описанные выше.

Ссылка на образ в Docker Hub - https://hub.docker.com/repository/docker/alekseeval/mp_db/general

# Развертывание сервера

Предполагается что сервер будет разворачиваться в Docker-контейнере.

Ссылка на образ в Docker Hub - https://hub.docker.com/repository/docker/alekseeval/mp_app/general

Конфигурационный файл должен располагаться по пути `/etc/mp_app/config.yaml`. Шаблон конфигурационного
файла сервера можно найти в репозитории по пути `internal/mindPalace/config.yaml`

# Сборка Docker-образов проекта
Для того чтобы локально собрать Docker-образ сервера, можно из корня проекта запустить команду:
```bash
docker build -t alekseeval/mp_app:1.0.0 -f internal/mindPalace/Dockerfile .
```
Для сборки Docker-образа для развертывания БД можно из директории `/db` запустить команду:
```bash
docker build . -t alekseeval/mp_db:1.0.0
```