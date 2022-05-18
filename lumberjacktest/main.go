package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	/*"github.com/natefinch/lumberjack"*/

	"runtime/debug"
)


func main(){
	/*log.SetOutput(&lumberjack.Logger{
		Filename:   "log.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Compress:   false, // disabled by default
	})*/

	//log.SetFormatter(&log.JSONFormatter{})


	log.Println("asdasd")

	//log.Fatalln("eish")

	log.WithFields(log.Fields{
		"animal": "walrus",
	  }).Info("A walrus appears")

	log.Debug("asdasd")

	log.WithFields(log.Fields{
		"f1": "val1",
		"f2": "val2",
	}).Error(fmt.Errorf("some error occured"))

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
	
	log.Info(string(debug.Stack()))

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}

