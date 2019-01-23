# MutantOrNot (mutant detection system)

The purpose of this ecosystem is to provide humanity with the most advanced tool for mutant detection.

## Description:
This system is componed mainly by two microservices written in golang

## Necessary components:
- gin-gonic framework
- Mongodb
- RabbitMQ

## How to run it:
For you to run it locally, you will need docker installed in your system.

1. Clone this repo in your local machine
2. Install docker if you don't have it already installed
3. Install `docker-compose`
4. Run `docker-compose up`

After that you will find two microservices up and running
1. `mutantornot` serving in port 8081
2. `dnastats` serving in port 8082

You will also find an instance of rabbitmq, and mongodb running each one in it's own container.
