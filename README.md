
# Ferrari Trunk

Backends for storing jobs information for the ferrari project

[![Build Status](https://travis-ci.org/ottogiron/ferraritrunk.svg?branch=master)](https://travis-ci.org/ottogiron/ferraritrunk)

## Backends

Backends are in charge of storing and retreiving job results

```go
//Backend defines a data store for jobs to be persisted
type Backend interface {
    Persist(jobResults []*worker.JobResult) error
    JobResults(workerId string) ([]*worker.JobResult, error)
    Job(jobID string) (*worker.JobResult, error)
}
```

## Available Backends

* [Elastic](backend/elastic)
