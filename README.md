# goWordAnalysis
## Web application that counts the frequency of words of input text file.

## Project dependencies
- Go (1.12+) + std libs
- github.com/gorrilla/mux
- github.com/stretchr/testify/assert
## Dev Installation
1. clone project into go/src/ 
2. run commands to install ext library dependencies: 
> - go get github.com/gorrilla/mux
> - go get github.com/stretchr/testify/assert
3. run command to execute program
> - go run main.go
### Testing
Tests follow the structure of {name}_test.go

Full coverage tests:
> go tests ./...
## User installation (for linux)
1. clone project 
2. run
> - ./goWordAnalysis
## Project structure

|Folder          |Function                       |
|----------------|-------------------------------|
|assets			 |`contains all static files such as sample test data and externally downloaded datasets`            |
|templates       |`front-end html files`            |
|uploads         |`directory used to store temp data files`|
|vender | `holds all abstracted out utility functions + tests programs`
|main.go				 | `backend + integration program: contains server + routing + utilities `
|goWordAnalysis | `compiled binary for linux64`

##Program Design

Quick overview: 

Go is chosen for this project because it highlights my ability to adapt to new technologies. I did not know how to program in go before this project. Though it's benefits in speed and ease of use certainly helped.

Program was designed with an emphasis on being unit testable, though an integration test is missing due to time constraint. Along with some needed refactoring on main.go. This project does follows the MVC architecture (lightly). Go's internal templating structure allows for decoupling of the front and the backend. The front-end is composed of index.html with materializecsss, holding only references to data. Main.go acts as a controller most of the time, dictacting flow of data, while vendor contains the different data structures and the associate methods.  There are cases of DRY as a result of lack of generics in Go.

Also, internal libraries are used as much as possible to showcase the go language itself. Thus the word counting, stemming, uniqifying, and sorting functions are all self implemented as opposed to using magic functions and provided datastructures in Python or similar languages.