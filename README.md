# MutantOrNot (mutant detection system)

The purpose of this ecosystem is to provide humanity with the most advanced tool for mutant detection.

## Description:
This system is componed mainly by two microservices written in golang

## Necessary golang libraries:
- gin-gonic framework
- Mongodb
- RabbitMQ

## How to run it:
For you to run it locally, you will need docker installed in your system.

1. Clone this repo in your local machine
2. Install docker if you don't have it already installed
3. Install `docker-compose`
4. Run `docker-compose build`
4. Run `docker-compose up`

After that you will be able to hit the API through two endpoints: 
1. http://localhost/mutant/ (POST ONLY)
2. http://localhost/stats (GET ONLY)

You will also find an instance of rabbitmq, and mongodb running each one in it's own container.

For testing the API, you can hit both endpoints deployed in localhost or the remote ones:

For localhost testing
=====================

*__Endpoint for getting stats__*
GET http://localhost/stats

*__Posible responses__*: 
1. 200 - Ok

*__Response Body__*:
{
    "count_mutant_dna": 2,
    "count_human_dna": 1,
    "ratio": 2
}

*__Endpoint for sending dna matrix__*
POST http://localhost/mutant
{
    Dna: ["AAAA", "CCCC", "GGGG", "TTTT"]
}
*__Posible responses__*: 
1. Ok - 200 (if Dna is mutant) 
2. 403 - Forbidden (if the Dna is human)
3. 404 - The DNA matrix is invalid
4. 404 - Invalid request (the POST body is not formed properly)

*(NOTE: The dna matrix has to be a NxN matrix with N >= 4, the posible
letters can be only `A`, `C`, `G` and `T`. The API is case sensitive, so
`a` is not the same than `A`, and `a` is an invalid input.)*

For remote testing:
===================
Remote links will be provided in the PDF documentation
