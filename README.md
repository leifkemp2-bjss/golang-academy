# Leif's Go Academy Work
This is my attempt at the BJSS Go Academy, based on the Kepler version of the exercise.
There are 3 different sections:
- Assignments, this is where assignments 1 to 10 can be found
- To Do App, this is the To Do application, split into the CLI and the Web App. The Web App can function with either In Memory Storage, or use a PostgreSQL docker container.
- Thread Safe Test, this is a test at writing a server that could concurrently handle requests in a thread safe manner

## Running each application
### Assignments
To run this code, simply run the assignments.exe file in the folder.
### To-Do Application, Part 1
Run the part1.exe file in the todoapp/part1 folder.
### To-Do Application, Part 2
#### CLI
To run the CLI, run cli.exe in the todoapp/part2/cli folder using this list of commands:
##### Commands
- help - Lists all commands
- command {command} {params} - Runs a specific command with given parameters
    - help, read, list, create, update, delete
- l - Lists all To Dos
- r - Reads a To Do by ID
- a - Creates a To Do, contents are required, status is optional
- d - Deletes a To Do by ID
- u - Updates a To Do with given ID, provided either contents or status are given
### Thread Safe Test
Run the threadsafetest.exe file in the threadsafetest folder, the API can then be called using CURL or postman at <b>localhost:8000</b>

##### Parameters
- contents / c - The contents parameter
- status / s - The status parameter
- id / i - The ID parameter
#### Web App
##### In Memory Storage
To run the webapp, run the part2.exe file in the todoapp/part2/web folder, the app can then be accessed from <b>localhost:8080</b>
##### PostgreSQL Storage
To run the database, Docker is required.

To run the webapp utilising a PostgreSQL database, run ```docker compose up -d``` on the docker-compose.yml file. Then run the part2.exe file with the ```-db``` flag to tell the application to use the database.

To tear down the container after finishing with the program, use ```docker compose down -v``` to also remove the volumes the container creates.

## File Storage
Some programs in this repository create files that are stored in either assignments/files, or todoapp/files.

## Future Plans
If I get the time to work on this repository after the academy finishes, I'd like to implement an SQL database of some form, though I'm currently unsure how to do so.
I would also like to explore more black box testing, which I attempted with the first version of the thread safe test code, but it was a struggle to get it to work properly.

## Credits
Alex Edwards - Simple Flash Messages in Golang
https://www.alexedwards.net/blog/simple-flash-messages-in-golang

Oli Hathaway - KV Store in Golang
https://github.com/labiraus/go-kvstore

BugBytes - Go - SQL Databases in Golang with the database/sql package
https://www.youtube.com/watch?v=Y7a0sNKdoQk