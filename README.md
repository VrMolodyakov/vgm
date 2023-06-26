# VGM
![image](https://github.com/VrMolodyakov/vgm/assets/99216816/493b3ec6-6b19-4fb7-8bdc-e96256cc60fa)

## General info
This is a project that allows you to view information about new/released VGM albums.With the ability to get a youtube playlist consisting of random albums presented on the site.

The project is made on a microservice basis, where the services themselves communicate via grpc.

## Scheme
![image](https://github.com/VrMolodyakov/vgm/assets/99216816/a01105c3-3be2-4358-88a6-0b5cc19b0d10)

## Technologies
Project is created with:
* Golang version: 1.18
* React version: 18.2.0
* Grafana version: 6.1.6
* Prometheus version: 6.1.6
* Redis version: 6.2
* Postgres version: 12.0
* Docker
* Nats:2.8.4

## Build
To build this project, you need to run the following commands:

```
https://github.com/VrMolodyakov/stock-market.git
make build
```

## Start
To run this project, you need to run the following command:
```
make start
```

## Metrics

Prometheus UI:
```
http://localhost:9090
```

Jaeger UI:
```
http://localhost:16686
```
