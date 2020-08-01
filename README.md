# Hotel Management System

## Description
Hotel management system is a simple service to manage hotel operations.

## Features
As a customer
1. Check room availability of all hotel
2. Reserve available room on a hotel

As an admin
1. Check-in guest
2. Add hotel registry

## Installation
- Run mysql on docker-compose
```
docker-compose up --build
```
- Run service using docker
```
docker build -t hotel_mgmt_svc .
docker run --network host hotel_mgmt_svc
```

## Postman Collection
https://documenter.getpostman.com/view/4226084/T1Dv8Ex2?version=latest