# Altair Technical Test
##Overview

Simple example to merge two different types of structs, and sorted by age from youngest to oldest.

## Dependencies

### GO
![Golang](https://golang.org/lib/godoc/images/go-logo-blue.svg)

The main language for the application

### Docker
![Docker](https://avatars0.githubusercontent.com/u/5429470?s=200&v=4)

We used docker to deploy the app

## Run application

To run the application, a Makefile is provided. To used it, on terminal write the following command:

```bash
make
```

Once is done, you will see the following screen, which show you the different options:

```bash
Welcome to Altair Technical Test. Choice one of the following options to start:
help:  Show this help.
app-up:  Command to run the application
app-down:  Command to shut down the application
```
To execute one of the options mentioned above, just write *make* on terminal followed by one of the options; example:

```bash
make app-up
```
