package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func runandwait(a, b int) int {
	c := a + b
	log.WithFields(log.Fields{
		"a": a,
		"b": b,
	}).Infof("c=%d", c)
	time.Sleep(time.Millisecond * 1)
	return 10
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
}

func main() {

	for i := 0; i < 100; i++ {
		a := runandwait(i, 2*i)
		log.Info("runandwait")
		log.Infof("result= %d", a)

		if i > 90 {
			log.WithFields(log.Fields{
				"i": i,
			}).Warn("close to 100...")
		}
		if i == 99 {
			log.Fatal("Reached 100")
		}
	}
	fmt.Println("done")
}
