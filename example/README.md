# Minimal example
## Prerequisites
- Kong installed (tested with v0.8.0)
## Configuration
Please change variable `kong_server` in `config.yaml` with your Kong gateway IP
## Usage
The server runs on `localhost:8080`
##### `/login`
- Returns a token for the user with username `test` and custom id `123`, if the user `test` doesn't exist it creates a new user into Kong gateway
- Make a request to endpoint `/login`

#### `/endpoint_auth`
- Returns the user authenticated by the Kong gateway.
- To simulate the gateway, pass these 3 headers during request:
    * `X-Consumer-Username` and set an username;
    * `X-Consumer-Custom-ID` and set a custom ID;
    * `X-Consumer-ID` and set a random alphanumeric string;
