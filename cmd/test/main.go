package main

import (
	"fmt"
	"os"

	"gitlab.com/scabala/zelligo"
)

type MyPlugin struct {
	message string
}

func (p *MyPlugin) Load(configuration map[string]string) {
	fmt.Fprintf(os.Stderr, "configuration for plugin is: %v", configuration)
	zelligo.Subscribe([]zelligo.EventType{
		zelligo.EventType_CopyToClipboard,
	})
}

func (p *MyPlugin) Update(event zelligo.Event) bool {
	fmt.Fprintf(os.Stderr, "log message: %v", event)
	return true
}

func (p *MyPlugin) Render(cols uint32, rows uint32) {
	fmt.Println(p.message)
}

func init() {
	zelligo.RegisterPlugin(&MyPlugin{message: "Hello from Go!"})
}

func main() {}
