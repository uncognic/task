# Task
## 📝 To-do CLI written in Go
This applet was made as a beginner project

Tasks are saved to a `tasks.json` file in the app's directory

## Building
To build this project you need to [install Go](https://go.dev/doc/install)
```
$ git clone github.com/dacixn/task
$ cd task
$ go build .
```

Now simply run the app!
```
$ ./task
```

## Features
Task currently supports a few basic operations:
* Adding tasks
* Deleting tasks individually
* Changing a task's completion state
* Clearing tasks.json

Advanced features are in progress, along with a refactor of the codebase

### Planned features
* Editing tasks
* Categorizing tasks
* Task timestamps
* Prettier