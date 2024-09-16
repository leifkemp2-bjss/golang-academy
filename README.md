# Leif's Go Academy Work
This is my attempt at the BJSS Go Academy, it is split into 3 different sections:
- Hello Go, this is just a quick Hello World script I wrote as my first code in Go.
- Assignments, this is where assignments 1 to 10 can be found
- To Do App, this is the To Do application, with the CLI and Web App

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

##### Parameters
- contents / c - The contents parameter
- status / s - The status parameter
- id / i - The ID parameter
#### Web App
To run the webapp, run the part2.exe file in the todoapp/part2/web folder, the app can then be accessed from <b>localhost:8080</b>

## File Storage
Some programs in this repository create files that are stored in either assignments/files, or todoapp/files.

## Credits
Alex Edwards - Simple Flash Messages in Golang
https://www.alexedwards.net/blog/simple-flash-messages-in-golang