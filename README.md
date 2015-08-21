# goprop
Golang read the properties file

# use

install
```
go get github.com/tjz101/goprop
```
prop.properties
```
port = 2770
driverName = mysql
dataSourceName = root:mysql99@tcp(127.0.0.1:3306)/capture?charset=utf8
maxOpenConns = 2000
maxIdleConns = 1000
```
main.go
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
	fmt.Println(prop.Get("maxIdleConns"))
}
```
