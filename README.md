# awesome-cli
First CLI app in Go. Objective to learn CLI app dev and Go as a language.

Skills learned:
Go, building CLI apps
Cobra
Viper
net/http
task + Taskfiles


Completely following along with https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43

Started Apr-16-25


## Usage
### Raw:
```
go run main.go {gopherName}
```

### Taskfile built (added to your /bin):
```
task build
```
Verify successful build
```
ll bin/awesome-cli
```
```
./bin/awesome-cli get {gopherName}
```