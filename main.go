package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"modlishka-render/config"
	"modlishka-render/core"
	"modlishka-render/log"
	"modlishka-render/plugin"
	"modlishka-render/runtime"
)

type Configuration struct{ config.Options }

// Initializes the logging object
func (c *Configuration) initLogging() {
	//
	// Logger
	//
	log.WithColors = true

	if *c.Debug == true {
		log.MinLevel = log.DEBUG
	} else {
		log.MinLevel = log.INFO
	}

	logGET := true
	if *c.LogPostOnly {
		logGET = false
	}

	log.Options = log.LoggingOptions{
		GET:            logGET,
		POST:           *c.LogPostOnly,
		LogRequestPath: *c.LogRequestFile,
	}

	log.Init(*c.LogRequestFile)
}

func (c *Configuration) initPlugins() {
	//
	// Initialize plugins
	//
	plugin.Init(c.Options)
}

func (c *Configuration) initRuntime() {
	//
	// Initialize runtime
	//
	runtime.Init(c.Options)
}

func (c *Configuration) initCore() {
	//
	// Initialize core
	//
	core.Init(c.Options)
}

func main() {
	//
	// Parse command line arguments
	//
	configFile := flag.String("config", "config.json", "JSON configuration file")
	flag.Parse()

	//
	// Load configuration
	//
	conf := &Configuration{}
	conf.Options = config.LoadConfiguration(*configFile)

	//
	// Initialize components
	//
	conf.initLogging()
	conf.initPlugins()
	conf.initRuntime()
	conf.initCore()

	//
	// Start the proxy
	//
	fmt.Println("Starting Modlishka reverse proxy...")
	fmt.Printf("Target: %s\n", *conf.Target)
	fmt.Printf("Proxy Domain: %s\n", *conf.ProxyDomain)
	fmt.Printf("Listening on: %s\n", *conf.ListeningAddress)

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nShutting down Modlishka...")
		os.Exit(0)
	}()

	// Start the server
	core.Start()
}