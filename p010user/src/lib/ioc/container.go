package ioc


import (
	"log"
	"sync"
	"go.uber.org/dig"
)

var instance *dig.Container
var once sync.Once

func GetContainerInstance() *dig.Container {

	once.Do(func() {

		c:= dig.New()

		log.Println("registering containers")

		/////// register provider here /////////
		if err := c.Provide(testProvider); err != nil {
			log.Panic(err.Error())
		}

		instance = c
	})


	return instance
}
