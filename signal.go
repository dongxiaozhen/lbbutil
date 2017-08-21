package lbbutil

import (
	"os"
	"os/signal"
)

func AsyncRegistSignal(handle func(), sig ...os.Signal) {
	c := make(chan os.Signal, len(sig))
	signal.Notify(c, sig...)
	go func() {
		<-c
		handle()
	}()
}

func MakeSignal(sig ...os.Signal) chan os.Signal {
	c := make(chan os.Signal, len(sig))
	signal.Notify(c, sig...)
	return c
}
func RegistSignal(handle func(), sig ...os.Signal) {
	c := make(chan os.Signal, len(sig))
	signal.Notify(c, sig...)
	<-c
	handle()
}
