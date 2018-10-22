## File Indexer

This is a simple project which reads files in the directory mentioned in the config file and indexes it. This then can be used to search the files present in the directory.



#### Motivation

This is a way of trying to index the logs file in dgplug so that we can easily find the files in which a word has occured in the file.



#### Dependency

This assumes that you have `Go` installed and setup.

All dependecies are added in `dep` and we need to make sure we install `fresh` although it is already added to the lock.



#### Development

We need to make a directory that has to be indexed and put few files in it. 

Put that as the `RootDirectory` in the `config.json` file.

We create a `log` directory and put it in the conf.

For development purpose you can run `go run main.go`

Now you need to create the index first so visit `localhost:8000/index`

Then you can search the file with `localhost:8000/search/{query}`

For ease of use we have included a `Makefile` with the following targets

```bash

make clean # for cleaning up  
make dev # for development purposes
make build # for normal usage

```

#### Usage
`cd` in the directory and run`fresh -c fresh.conf`
This will run the server and start watching for any changes in the file, it does two things
hot-reloading and auto-indexing.