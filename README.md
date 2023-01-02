# Dokla

Push/Pull feed ingestion service

## Features of this project:

- Implement simple push and pull ingestion service
- Use mongodb for persistence. Haven't used volumes, so if the mongo(shardsvr) container restarts, data will be lost
- Implement ingest service with support to ingest from different sources
- Implement fetch by Id and fetch by duration, only few cases are handled
- Service is eventual consistent, highly available (Data is replicated) see [here](https://hub.docker.com/r/bitnami/mongodb-sharded) for more details, reading is done from secondary db and writing to primary db
- Ingestion is idempotent, requests can be replayed
- Builder pattern for building the requests from different sources
- Adapter pattern for converting the requests from different sources to internal req structure
- Added postman tests collection


## TODO

- [ ] Add better validations and error handling
- [ ] Fetch by duration supports only fewer cases, can be made much more robust
- [ ] Batch insertion of feeds
- [ ] Use kafka, for consuming messages from, instead of exposing an endpoint



## Building

### Local

1. Clone Dokla

2. Build and Run Dokla

    ```
    sh run.sh
    ```

    Please refer github workflows for more details

---

- GitHub [@nrnc](https://github.com/nrnc)