package cache

import "fmt"

type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting using LFU strategy")

}
