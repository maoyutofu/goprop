# goprop
Golang read the properties file

# use
 
```
package main

import (
	"fmt"
	"github.com/tjz101/goprop"
)

func main() {
	prop := goprop.NewProp()

	prop.Read("prop.properties")

	fmt.Println(prop.Get("port"))
	fmt.Println(prop.Get("driverName"))
	fmt.Println(prop.Get("dataSourceName"))
	fmt.Println(prop.Get("maxOpenConns"))
	fmt.Println(prop.Get("mMaxIdleConns"))
}
```
