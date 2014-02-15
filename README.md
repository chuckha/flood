# Flood

A command line interface to the Digital Ocean API.

## Installation

Ensure `$GOPATH/bin` is in your PATH environment variable.

Run `$ go get github.com/chuckha/flood`

## Usage

You need to set two environment variables for this to work.

1. `DIGITAL_OCEAN_CLIENT_ID`
2. `DIGITAL_OCEAN_API_KEY`

Then you can run the `flood` command!

See `flood` for usage.

## Examples

* See all of your droplets:

        flood droplet list

* Show all available droplet sizes:

        flood size list

## License

MIT.


