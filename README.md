# go-eero

This is a limited functionality SDK for interacting with the [Eero](https://eero.com) API. Currently, the following endpoints are supported:

- `/login` - Initiates login flow
- `/login/verify` - Validates login with code
- `/login/refresh` - Refreshes an existing user token
- `/account` - Retrieves details about the logged in account
- `/networks/<id>` - Retrieves information about the specified network ID
- `/networks/<id>/devices` - Retrieves a list of devices on the network
- `/data_usage/breakdown` - Retrieves data usage per device
