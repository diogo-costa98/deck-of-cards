# Deck of Cards
The Deck of Cards application is a simple command-line tool written in Go that allows you to create and manipulate a deck of playing cards. With this tool, you can create a new deck of cards (shuffled or unshuffled), draw cards from the deck and open an existing deck.

## Directory Organization
This directory organization is a common structure used in Go projects that use the MVC (Model-View-Controller) pattern to organize their code.  
* `api` - contains the application's main source code, including handlers, models, and routes. 
  * `handlers` - contains all the HTTP request handlers, including deck_handler.go, which defines the deck-related endpoints. 
  * `models` - contains the domain models of the application, including card.go, which defines the properties and behaviors of a card in a deck, and deck.go, which defines the properties and behaviors of a deck of cards. 
  * `routes` - contains the route definitions of the application, including deck_routes.go, which defines the endpoints for the deck-related operations.
  * `server.go` - contains the code to start and stop the HTTP server.
* `test` - contains all the tests of the application. 
  * `bdd` - contains the BDD (Behavior-Driven Development) tests, including bdd_suite_test.go, which initializes the Ginkgo test suite, and deck_test.go, which contains the tests for the deck-related operations.
  * `tdd` - contains card_test.go and deck_test.go files, which contain unit tests for the card and deck.
* `main.go` - is the entry point of the application. It sets up the configuration of the server and starts the server. 

This directory organization helps to organize the application code into logical groups and makes it easier to maintain and test.

## Installation
To use the Deck of Cards application, you must first install [Go](https://go.dev/dl/).

Once you have installed Go, you can download and install the Deck of Cards application by running the following command:
```sh
git clone https://github.com/diogo-costa98/deck-of-cards.git
```
Open the application folder with the command:
```sh
cd .\deck-of-cards\ #Windows Powershell
cd ./deck-of-cards/ #Linux and macOS
```
Finally, run:
```sh
go build
```
This will build the application and create an executable file.

## Usage
To use the Deck of Cards application, run the following command to setup a server at port 8080:

```sh
.\deck-of-cards.exe #Windows
./deck-of-cards     #Linux and macOS
```

## Examples
Here are some of request commands and their expected output:

### Create unshuffled deck
`POST localhost:8080/deck`

Expected response:
```json
{
    "deck_id": "b3e7ab53-33c9-4201-8307-230e89845310",
    "shuffled": false,
    "remaining": 52
}
```

### Create shuffled deck
`POST localhost:8080/deck?shuffle=true`

Expected response:
```json
{
    "deck_id": "d5f17b57-97f0-44a8-98e7-e4ee9fe41461",
    "shuffled": true,
    "remaining": 52
}
```

### Create partial deck
`POST localhost:8080/deck?cards=AS,KD,AC,2C,KH`

Expected response:
```json
{
    "deck_id": "429a86fa-8cc4-4c0a-bcfe-ca01fa7dc32d",
    "shuffled": false,
    "remaining": 5
}
```

### Open a created deck
`GET localhost:8080/deck/429a86fa-8cc4-4c0a-bcfe-ca01fa7dc32d`

Expected response:
```json
{
    "deck_id": "429a86fa-8cc4-4c0a-bcfe-ca01fa7dc32d",
    "shuffled": false,
    "remaining": 5,
    "cards": [
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
        {
            "value": "KING",
            "suit": "DIAMONDS",
            "code": "KD"
        },
        {
            "value": "ACE",
            "suit": "CLUBS",
            "code": "AC"
        },
        {
            "value": "2",
            "suit": "CLUBS",
            "code": "2C"
        },
        {
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        }
    ]
}
```

### Draw from a deck
`POST localhost:8080/deck/429a86fa-8cc4-4c0a-bcfe-ca01fa7dc32d/draw?count=2`

Expected response:
```json
{
    "cards": [
        {
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        },
        {
            "value": "2",
            "suit": "CLUBS",
            "code": "2C"
        }
    ]
}
```

## Testing
To run the tests for the Deck of Cards application with feedback, run the following command:

```sh
go test ./... -v
```

This will execute all of the tests in the project and display the results in the console.
If any tests fail, the output will indicate which tests failed and why.
