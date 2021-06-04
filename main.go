package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"

	"./config"
	"github.com/BurntSushi/toml"
	"gopkg.in/robfig/cron.v2"
)

var (
	sha1ver    string // sha1 revision used to build the program
	buildTime  string // when the executable was built
	flgVersion bool
)

func main() {
	parseCmdLineFlags()

	cf := ReadConfig()

	c := cron.New()

	// Run mon-fri
	scheduledTime := fmt.Sprintf(" %d %d * * 1-5", cf.Time.Minute, cf.Time.Hour)

	c.AddFunc(scheduledTime, func() { sourceBash(cf.Input.SourceScript) })
	c.Start()
	log.Println("goScheduler has started!")

	// Channel that listens for an incoming signal to stop the program
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	log.Println("Program Stopped")
}

// Function to read config file
func ReadConfig() config.Details {
	var cf config.Details

	if _, err := toml.DecodeFile("config.toml", &cf); err != nil {
		log.Fatalln("Could not find config.toml at root, please copy from config/template.toml and populate.\nError:", err)
	}

	return cf
}

// Function to read and execute bash script
func sourceBash(scriptName string) {
	cmd := exec.Command("/bin/bash", scriptName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	log.Println("Script has run")
	return
}

// Function to parse command line flags
func parseCmdLineFlags() {
	flag.BoolVar(&flgVersion, "version", false, "if true, print version and exit")
	flag.Parse()
	if flgVersion {
		fmt.Printf("Build on %s from sha1 %s\n", buildTime, sha1ver)
		os.Exit(0)
	}
}
