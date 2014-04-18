# BringGo v0.1

> Author: Brian Tomlinson <brian.tomlinson@linux.com>


## Description

> Description: A Go CLI client using [libbring](http://github.com/darthlukan/libbring) to allow users
> to retrieve data provided by Bring's API.  Features include tracking packages by
> tracking number and finding the nearest pickup point based on location criteria
> such as Zip/Postal Code, Latitude and Longitude, and location ID.


## Usage

```
    $ go get github.com/darthlukan/bringgo

    // Tracking number lookup
    $ bringgo <tracking number>

    // Pickup Location Search by latitude and longitude
    $ bringgo <lat> <lon>
```

## TODO

* Add more of the features available in [libbring](https://github.com/darthlukan/libbring)
* Tests


## NOTE
> The only thing that works right now is location search via latitude and longitude.


## License

> GPLv2, see LICENSE file.
