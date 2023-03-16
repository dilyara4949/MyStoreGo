package main

import (
	"github.com/dilyara4949/MyStoreGo/pkg"
)


func main() {
	pkg.Routes()
}




// lsof -i :8000
// kill -9  <PID>

// % docker build .  
// docker run -d -p 8000:8080 288ca3be3974   (image id)  
// docker run --name assigment2 -p 8000:8000 -d 1feb2c724dab
