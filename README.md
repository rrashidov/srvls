# srvls

**srvls** is a serverless platform written in Go.

## Goals

The project has the following goals:

- help me get started with [Go](https://go.dev)
- help me get into the world of serverless by simply trying to implement it myself

## Features

The platform will provide the following features:

- function management
    - create a function
    - modify a function
    - list/get a function
    - delete a function
- function execution
    - trigger a function execution
    - list/get a function execution

## REST API

| Method | URL                                                             | Description                    |
| -------| ----------------------------------------------------------------| -------------------------------|
| GET    | /api/v1/:tenant/functions                                       | List functions                 |
| GET    | /api/v1/:tenant/functions/:function_id                          | Get function details           |
| POST   | /api/v1/:tenant/functions                                       | Create new function            |
| PUT    | /api/v1/:tenant/functions/:function_id                          | Update an existing function    |
| DELETE | /api/v1/:tenant/functions/:function_id                          | Delete an existing function    |
| GET    | /api/v1/:tenant/functions/:function_id/executions               | List function executions       |
| GET    | /api/v1/:tenant/functions/:function_id/executions/:execution_id | Get function execution details |
| POST   | /api/v1/:tenant/functions/:function_id/executions               | Trigger function execution     |

## Data model 

- tenant
  - id
  - name
  - description
  - function
    - id
    - name
    - description
    - container image
    - function execution
      - date time triggered
      - date time started
      - date time finished
      - status

## Boundary conditions

- the platform will initially work only with mysql db;
- the platform will initially work only with k8s as an execution runtime;

## Development

The platform could be tested locally. In order to run and test it locally, the following dependencies have to be present:

- MySQL server
- [golang-migrate](https://github.com/golang-migrate/migrate)
