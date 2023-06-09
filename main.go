package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

var filePath = os.Getenv("AUDIT_FILE_PATH")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type AuditData struct {
	auditContent []byte
}

func (ad *AuditData) auditFileCheck() []byte {
	ad.auditContent, _ = ioutil.ReadFile(filePath)
	return ad.auditContent
}

func logHandler() {
	audit := &AuditData{}
	audit.auditFileCheck()

	// Log the Audit Data as STDOUT
	log.Println(string(audit.auditContent))

	// Truncate the file content after it is being alerted
	if err := os.Truncate(filePath, 0); err != nil {
		check(err)
	}
}

// watchFile is a function to work on runtime to monitor any change
// on the filePath and Log it as STDOUT then truncate the file content
func watchFile() {
	watcher, err := fsnotify.NewWatcher()
	check(err)
	defer watcher.Close()
	// starting to listen on events
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					logHandler()
				}
			case err := <-watcher.Errors:
				check(err)
			}
		}
	}()
	err = watcher.Add(filePath)
	check(err)
	<-make(chan struct{})
}

func main() {
	watchFile()
}
