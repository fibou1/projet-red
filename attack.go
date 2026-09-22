package main

import (
	"fmt"
	"time"
)

func poisonPot(monster *Monster) {
	for i := 0; i < 3; i++ {
		monster.CurrentHP -= 10
		fmt.Printf("%d/%d", monster.CurrentHP, monster.MaxHP)
		time.Sleep(1 * time.Second)
	}
}
