package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Gracefully() {
	quit := make(chan os.Signal, 1)
	defer close(quit)
	fmt.Println("in gracefilly shutting down")
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
