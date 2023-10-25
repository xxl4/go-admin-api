# Docker Install go-admin
> Build a full-featured administrative interface quickly Go-Admin

### 1、Ready

> docker-compose.yaml (MySQL Version)
```
version: '3.8'
services:
  go-admin-api:
    container_name: go-admin-api
    image: registry.ap-southeast-1.aliyuncs.com/kuops/go-admin-api:latest
    privileged: true
    restart: always
    ports:
      - 8000:8000
    volumes:
      - ./config/:/go-admin-api/config/
      - ./static/:/go-admin-api/static/
      - ./temp/:/go-admin-api/temp/
    networks:
      - kuops
  mysql:
    container_name: mysql
    image: mysql:5.7.26
    hostname: mysql
    restart: always
    command: mysqld
    environment:
      - TZ=Asia/Shanghai
      - MYSQL_ROOT_PASSWORD=GoAdmin
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - ./data:/var/lib/mysql/data
      - ./data/my.cnf:/etc/mysql/my.cnf
    networks:
      - kuops
networks:
  kuops:
    driver: bridge
```

> docker-compose.yaml (PG Version)
```
version: '3.8'
services:
  go-admin-api:
    container_name: go-admin-api
    image: registry.ap-southeast-1.aliyuncs.com/kuops/go-admin-api:latest
    privileged: true
    restart: always
    ports:
      - 8000:8000
    volumes:
      - ./config/:/go-admin-api/config/
      - ./static/:/go-admin-api/static/
      - ./temp/:/go-admin-api/temp/
    networks:
      - kuops
  postgres:
    image: postgres:14.1-alpine
    healthcheck:
      test: [ "CMD", "pg_isready", "-q", "-d", "postgres", "-U", "root" ]
      timeout: 45s
      interval: 10s
      retries: 10
    restart: always
    environment:
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=GoAdmin
      - APP_DB_USER=GoAdmin
      - APP_DB_PASS=GoAdmin
      - APP_DB_NAME=GoAdmin
    volumes:
      - ./db:/docker-entrypoint-initdb.d/
    expod:
      - 5432:5432
networks:
  kuops:
    driver: bridge
```
### 2、Configure

[Configure Docs](https://nicelizhi.github.io/go-admin-api/guide/configure/)

### 3、Start It

> You need have setting.yml file in ./config dir
```
docker run --name api-admin -p 8000:8000 -v ./config:/go-admin-api/config/ -d registry.ap-southeast-1.aliyuncs.com/kuops/go-admin-api:latest
```

### 4、Test it

```
docker exec -it api-admin bash 
netstat -an | grep 8000
```

### Issue Submit
https://github.com/nicelizhi/go-admin-api/issues

