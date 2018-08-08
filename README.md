## File Indexer

This is a simple project which reads files in the directory mentioned in the
config file and indexes it. This then can be used to search the files present in
the directory.

#### Motivation

This is a way of trying to index the logs file in dgplug so that we can easily
find the files in which a word has occured in the file.

#### Dependency

This assumes that you have `Go` installed and setup

We need to install `mux`, `bleve` and `handler`.

`go get -u github.com/gorilla/mux`

`go get -u github.com/blevesearch/bleve`

`go get -u github.com/gorilla/handlers`

#### Development

We need to make a directory that has to be indexed and put few files in it.

Put that as the `RootDirectory` in the `config.json` file.

For development purpose you can run `go run main.go`

Now you need to create the index first so visit `localhost:8000/index`

Then you can search the file with `localhost:8000/search/{query}`

#### Usage

`cd` in the directory and run

`go build`

This creates a binary in the directory now we can just run it as `./file-indexer`