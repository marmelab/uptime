# Uptime

## Setup

In order to make the project functional, you have to run following command:

``` sh
make setup
```

In order to initialise the database, you have to run following command:

```sh
make init_db
```

You can adding default destination to ping with the command:

```sh
make load_fixtures
```

Then you can run the app with:

```sh
make run 
```

Note: you need [docker-compose](https://docs.docker.com/compose/#installation-and-set-up) to get the full architecture working.
