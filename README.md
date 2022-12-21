Spring Boot 3 VS GO
=================

[![Author](https://img.shields.io/badge/author-@nao99-green.svg)](https://github.com/nao99)
[![Source Code](https://img.shields.io/badge/source-nao99/spring_boot_3_vs_go-blue.svg)](https://github.com/nao99/spring-boot-3-vs-go/tree/master)

## What is it?
This project contains two identical REST applications written in Java and GO

Application in Java uses:
 - Spring Framework 6
 - Spring Boot 3
 - Other Spring-based dependencies (web, data)

Application in GO uses:
 - Nothing interesting


The purpose of this project is: Compare performance between compiled Java and GO web applications

## Disclaimer
Please, keep in mind, there are no production ready services <br>
They are only for testing and nothing more

## Main parts
Application stuff is pretty easy: There are 3 REST API endpoints to "manage" products in a database

Get 20 last products:
```
GET /api/v1/products
Accept: */*
Content-Type: application/json
Cache-Control: no-cache
```

Get a product by id:
```
GET /api/v1/products/{productId}
Accept: */*
Content-Type: application/json
Cache-Control: no-cache
```

Create a new product:
```
POST /api/v1/products
Accept: */*
Content-Type: application/json
Cache-Control: no-cache

{
"category": "GAME",
"name": "God Of War",
"description": "A super game of the 2022 year!"
}
```

## Environment variables

### Java:

| Variable        | Description                     | Example                                  |
|-----------------|---------------------------------|------------------------------------------|
| APP_DB_URL      | URL to Postgresql database      | postgresql://172.30.4.2/app?charset=utf8 |
| APP_DB_USERNAME | Username to Postgresql database | dbuser                                   |
| APP_DB_PASSWORD | Password to Postgresql database | dbuser                                   |

### GO:

| Variable | Description                           | Example                                                   |
|----------|---------------------------------------|-----------------------------------------------------------|
| APP_DSN  | Information about database connection | postgresql://dbuser:dbuser@172.30.4.2/app?sslmode=disable |

## If you want to run
### Java:
0. Install SdkMan: https://sdkman.io/install
1. Install GraalVM

```
sudo apt update
sdk install java 22.3.r17-grl
```

2. Install UPX:

```
sudo apt install upx
```

3. Build application:

```
./gradlew clean :products-spring-boot:nativeCompile
```

4. Compress application:

```
upx products-spring-boot/build/native/nativeCompile/products-spring-boot
```

5. Export all required ENV variables:

```
export APP_DB_URL=your_value
export APP_DB_PASSWORD=your_value
export APP_DB_PASSWORD=your_value
export APP_PORT=your_value
```

6. Run application:

```
./products-spring-boot/build/native/nativeCompile/products-spring-boot
```

![Products Spring Boot Compiled](images/products_spring_boot_3_compiled_running.png)

### GO:
1. Build application:

```
cd products-go
go build -o build/products-go cmd/app/main.go
```

2.  Export all required ENV variables:

```
export APP_DSN=your_value
export APP_PORT=your_value
```

2. Run application:

```
./build/products-go
```

![Products GO Compiled](images/products_go_compiled_running.png)
