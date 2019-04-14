# Basic Example

A simple demo on using Viaduct locally. Viaduct acts as an API gateway that proxies requests to the specified APIs.

## Installation

Refer to the main [README](https://github.com/jace-ys/viaduct#basic-example).

## APIs
1. [Reqres](https://reqres.in) - Users
2. [JSONPlaceholder](https://jsonplaceholder.typicode.com) - To Do's
3. [PunkAPI](https://punkapi.com/documentation/v2) - Brewdog Beers
4. [An API of Ice And Fire](https://anapioficeandfire.com/) - Game of Thrones
5. [MetaWeather (London)](https://www.metaweather.com/api/) - London Weather
5. [MetaWeather (Singapore)](https://www.metaweather.com/api/) - Singapore Weather

## Endpoints

API endpoints can be found in their respective API documentations.

Requests to `http://localhost:5000/{service-prefix}/{request-uri}` will be proxied to:

```
{upstream-url}/{request-uri}
```

Refer to the [config file](https://github.com/jace-ys/viaduct/blob/master/examples/basic/config.yml) for the prefix and upstream URL of each API.
