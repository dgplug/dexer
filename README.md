## File Indexer

This is a simple project which reads files in the directory mentioned in the config file and indexes it. This then can be used to search the files present in the directory.



#### Motivation

This is a way of trying to index the logs file in [dgplug](https://dgplug.org/) so that we can easily find the files in which a word has occured in the file.



#### Dependency

This assumes that you have `Go` installed and setup.

`pip3` and `python3` is what we need for development purpose.


### How to install

1. [Install Go](https://golang.org/doc/install)
2. Run the following commands

```bash
$ make all
$ dexer
```

Open http://localhost:8000 in your browser to open the user interface.

The default configurations can be changed by editing the `config.json` file.
```
{
    "RootDirectory": "logs",
    "IndexFilename": "irclogs.bleve",
    "Port": ":8000"
}
```

### API

Once everything is working fine install the postman plugin for your browser. And from that plugin you need to hit the endpoint as:

`localhost:8000/search/american`

Here american is the query word that I passed, make sure to open any file in the logs/ directory and find a word to search. It will look like this:

![Missing screenshot](https://image.ibb.co/cWP5TA/postman-query.png "Postman Screenshot")


### Run locally using docker

You can run the application using Docker in your local machine. It will use the `Dockerfile` instructions. Make sure you have [Docker](https://www.docker.com/) installed in your machine.

Run the following commands to build and run the docker image.

```bash
$ make docker-build # builds docker image
$ make docker-run # runs the image in new container
```