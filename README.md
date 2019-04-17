<table>
        <tr>
            <td><img width="20" src="https://cdnjs.cloudflare.com/ajax/libs/octicons/8.5.0/svg/archive.svg" alt="archived" /></td>
            <td><strong>Archived Repository</strong><br />
            This code is no longer maintained. Feel free to fork it, but use it at your own risks.
        </td>
        </tr>
</table>

# Uptime

:warning: Before installing this project you must know that uptime is outdated and not maintained anymore by marmelab's team.

If you are interested to maintain this project, please [open an issue](https://github.com/marmelab/uptime/issues).

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
