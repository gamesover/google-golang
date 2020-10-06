package main

import (
	"fmt"
	"sync"
	"time"
)

var eatGroup sync.WaitGroup

type chopS struct{ sync.Mutex }
type philo struct {
	id              int
	leftCS, rightCS *chopS
}

func (p philo) eat() {
	defer eatGroup.Done()

	for j := 0; j < 3; j++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Second)
		fmt.Printf("finishing eating %d\n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	csticks := make([]*chopS, 5)
	for i := 0; i < 5; i++ {
		csticks[i] = new(chopS)
	}
	philos := make([]*philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &philo{i, csticks[i], csticks[(i+1)%5]}
		eatGroup.Add(1)
		go philos[i].eat()
	}

	eatGroup.Wait()
}
