package main

import (
	"fmt"
	"os"

	"gitlab.com/scabala/zelligo"
)

type MyPlugin struct {
	message string
}

func (p *MyPlugin) Load(configuration map[string]string) error {
	fmt.Fprintf(os.Stderr, "configuration for plugin is: %v", configuration)
	zelligo.Subscribe([]zelligo.EventType{
		zelligo.EventType_CopyToClipboard,
	})
	return nil
}

func (p *MyPlugin) Update(event zelligo.Event) (bool, error) {
	fmt.Fprintf(os.Stderr, "log message: %v", event)
	return true, nil
}

func (p *MyPlugin) Render(cols uint32, rows uint32) error {
	fmt.Println(p.message)
	return nil
}

func init() {
	zelligo.RegisterPlugin(&MyPlugin{message: "Hello from Go!"})
}

func main() {}
