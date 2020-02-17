# go-haversine
The Haversine formula determines the great-circle distance between two points on a sphere given their latitudes and longitudes. Important in navigation, it is a special case of a more general formula in spherical trigonometry, the law of Haversines, that relates the sides and angles of spherical triangles.

## Installation

```
go get https://github.com/kwahome/go-haversine
```

## Usage

Below is an example showing how to calculate the shortest path between two coordinates on the surface of the Earth.

    package main
    
    import (
    	"fmt"
    	"github.com/kwahome/go-haversine/pkg/haversine"
    )
    
    func main() {
    	nairobi := haversine.Coordinate{
    		Latitude: 1.2921,
    		Longitude: 36.8219,
    	}
    	mombasa := haversine.Coordinate{
    		Latitude: 4.0435,
    		Longitude: 39.6682,
    	}
    
    	units := haversine.M
    	distance := nairobi.DistanceTo(mombasa, units)
    	fmt.Println("Distance from Nairobi =", nairobi, "to Mombasa =", mombasa, "in", units, "is", distance)
    
    	units = haversine.KM
    	distance = nairobi.DistanceTo(mombasa, units)
    	fmt.Println("Distance from Nairobi =", nairobi, "to Mombasa =", mombasa, "in", units, "is", distance)
    
    	units = haversine.MI
    	distance = nairobi.DistanceTo(mombasa, units)
    	fmt.Println("Distance from Nairobi =", nairobi, "to Mombasa =", mombasa, "in", units, "is", distance)
    }

## Development

Find the entire code in the repository:

```
https://github.com/kwahome/go-haversine
```

Use convenience scripts in `./bin` to run tests and format code:

Tests:

```bash
./bin/unit_test.sh
```

Code formatting:

```bash
./bin/code_format.sh
```

## Docs
https://godoc.org/github.com/kwahome/go-haversine

## License

See the [LICENSE](LICENSE) file for license rights and limitations (MIT).
