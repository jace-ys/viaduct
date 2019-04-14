[![Build Status][ci-build]][ci-status]
[![Docker Pulls][docker-pulls]][docker-image]

[ci-build]:	https://img.shields.io/travis/jace-ys/viaduct/master.svg?style=for-the-badge&logo=travis
[ci-status]: https://travis-ci.com/jace-ys/viaduct
[docker-pulls]: https://img.shields.io/docker/pulls/jaceys/viaduct.svg?style=for-the-badge&logo=docker
[docker-image]: https://hub.docker.com/r/jaceys/viaduct

# Viaduct

<img src="https://www.getyourguide.com/magazine/wp-content/uploads/2018/07/Glenfinnan-Viaduct-GetYourGuide.jpg" height="200" align="right"/>

Viaduct is a lightweight and configurable API gateway written in Go, largely a fork of [jakewright/drawbridge](https://github.com/jakewright/drawbridge). It acts as a reverse proxy for one or more APIs.

Viaduct comes with a CLI for easy usage and configuration. See [usage](https://github.com/jace-ys/viaduct#running-locally) for list of commands and flags.

## Configuration

Viaduct requires no data store and is instead configured using a .yml file containing all the API definitions. File name is not important - you can specify the config file to be used via the CLI or environmental variables.

#### API definition:

1. Required

   * `name` - a friendly name for the API
   * `prefix` - the path prefix to be handled
   * `upstream_url` - the target URL to proxy to

2. Optional

   * `methods` - accepted HTTP methods

      Default: ["GET", "POST", "PUT", "PATCH", "DELETE"]

   * `allow_cross_origin` - enable CORS

      Default: false

   * `middlewares` - list of middlewares to be applied for the API

      Default: [] (none)

```
services:
  typicode:
    name: "JSONPlaceholder"
    prefix: "typicode"
    upstream_url: "https://jsonplaceholder.typicode.com"
    methods: ["GET", "POST"]
    allow_cross_origin: false
    middlewares:
      logging: true
```

If the Viaduct server was running at localhost:5000, http://localhost:5000/typicode/posts would resolve and proxy to https://jsonplaceholder.typicode.com/posts.

For more, see [sample config file](https://github.com/jace-ys/viaduct/blob/master/config/config.sample.yml) or [examples](https://github.com/jace-ys/viaduct/tree/master/examples).

## Usage

A Makefile with some useful targets has been included to aid in the setting up and usage of Viaduct.

#### Running Locally

1. Installation and Setup

   Clone the repository:

   ```
   git clone https://github.com/jace-ys/viaduct.git
   cd viaduct
   ```

   Build the Go binary:

   ```
   make build
   ```

2. Use the CLI

   Start the Viaduct server:

   ```
   make execute
   ```

   Or:

   ```
   ./viaduct start [FLAGS]
   ```

   #### Flags:

   _Note: Flags are overriden by environmental variables, making it easy to configure Viaduct for Docker via environment injection_

   1. Port `-p, --port`

      * Description: port to run Viaduct server on

      * Default: `80`

      * Env: `PORT`

   2. Config File `-f, --file`

      * Description: path to .yml configuration file

      * Default: `/config/config.yml`

      * Env: `CONFIG_FILE`

      * Recommendations:

         - Local development: use `config/config.sample.yml`

         - Docker: keep default

   3. Help `-h, --help`

#### With Docker / Docker Compose

1. Build Docker Image

```
make container
```

2. Run Docker Image (uses `config/config.sample.yml` as mounted volume)

```
make run
```

3. Run with Docker Compose (uses `config/config.sample.yml` as mounted volume)

```
make compose
```

## Examples

#### Basic Example

Run the following command to start Viaduct locally, using the specified config file. Refer [here](https://github.com/jace-ys/viaduct/tree/master/examples/basic) for more info on the API endpoints.

```
./viaduct start -p 5000 -f examples/basic/config.yml
```

#### Docker Compose Example

Run the following command to pull the required Docker images and start the containers. Refer [here](https://github.com/jace-ys/viaduct/tree/master/examples/docker-compose) for more info on the API endpoints.

```
docker-compose -f examples/docker-compose/docker-compose.yml up
```

## Contribute
If you are interested in contributing, make a pull request and/or email me at jaceys.tan@gmail.com.
