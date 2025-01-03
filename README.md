# Communication Service

The Leagueify communication service is responsible for managing user-to-user
communication and system-generated notifications. The Leagueify communication
service uses  [Go][go-download] using version `1.23.0`.

## Key Functions

- Facilitate direct messaging and group messaging.
- Send system-generated notifications for events like schedule changes and roster updates.
- Provide message management features.
- Support multiple communication channels, including email, SMS, and in-app messaging.

## Development

### Prerequisites

- [Docker][docker-download] is installed and running.

### Getting Started

Local development of the Leagueify communication service uses docker for a
consistent development environment. Running the Leagueify communication service
locally can be accomplished using commands found in the
[Makefile][repo-makefile]. During local development changes will trigger a live
reload, eliminating the requirement to restart the docker image manually. This
is accomplished by using the wonderful tool [air][github-air]. The most common
commands have been outlined below:

```bash
# start leagueify docker image
make start

# stop leagueify docker image removing attached volumes
make clean
```

The Leagueify communication service is ready for development once the banner
output is visible within the terminal. By default the Leagueify communication
service api docs are accessible at
[http://localhost:6506/communication/docs][service-url]. The banner below was
created using the [Text to ASCII Art Generator by Patorjk][patorjk-taag].

```
leagueify-communication-1  |
leagueify-communication-1  | '||'      '||''''|      |      ..|'''.|  '||'  '|' '||''''|  '||' '||''''| '||' '|'
leagueify-communication-1  |  ||        ||  .       |||    .|'     '   ||    |   ||  .     ||   ||  .     || |
leagueify-communication-1  |  ||        ||''|      |  ||   ||    ....  ||    |   ||''|     ||   ||''|      ||
leagueify-communication-1  |  ||        ||        .''''|.  '|.    ||   ||    |   ||        ||   ||         ||
leagueify-communication-1  | .||.....| .||.....| .|.  .||.  ''|...'|    '|..'   .||.....| .||. .||.       .||.
leagueify-communication-1  |
leagueify-communication-1  |   ..|'''.|  ..|''||   '||    ||' '||    ||' '||'  '|' '|.   '|' '||'   ..|'''.|     |     |''||''| '||'  ..|''||   '|.   '|'
leagueify-communication-1  | .|'     '  .|'    ||   |||  |||   |||  |||   ||    |   |'|   |   ||  .|'     '     |||       ||     ||  .|'    ||   |'|   |
leagueify-communication-1  | ||         ||      ||  |'|..'||   |'|..'||   ||    |   | '|. |   ||  ||           |  ||      ||     ||  ||      ||  | '|. |
leagueify-communication-1  | '|.      . '|.     ||  | '|' ||   | '|' ||   ||    |   |   |||   ||  '|.      .  .''''|.     ||     ||  '|.     ||  |   |||
leagueify-communication-1  |  ''|....'   ''|...|'  .|. | .||. .|. | .||.   '|..'   .|.   '|  .||.  ''|....'  .|.  .||.   .||.   .||.  ''|...|'  .|.   '|
leagueify-communication-1  |
```

[docker-download]: https://www.docker.com/get-started/
[github-air]: https://github.com/air-verse/air
[go-download]: https://go.dev/dl/
[patorjk-taag]: https://patorjk.com/software/taag/#p=display&f=Kban&t=LEAGUEIFY%0ACOMMUNICATION
[repo-makefile]: ./Makefile
[service-url]: http://localhost:6506/communication/docs
