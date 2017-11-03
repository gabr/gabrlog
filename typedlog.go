package main

import (
	"fmt"
	"log"
	"os"
)

type TypedLogType string

const (
	LLog  TypedLogType = "LOG"
	LWarn TypedLogType = "WARNING"
	LErr  TypedLogType = "ERROR"
)

type TypedLog struct {
	logFile     *os.File
	LogFileInfo os.FileInfo
	detailed    bool
}

func (gl *TypedLog) SetFile(path string) error {
	var err error
	gl.logFile, err = os.OpenFile(
		path,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		os.ModeAppend)
	if err != nil {
		return err
	}

	gl.LogFileInfo, err = gl.logFile.Stat()
	if err != nil {
		return err
	}

	log.SetOutput(gl.logFile)
	return nil
}

func (gl *TypedLog) CloseFile() error {
	if gl.logFile == nil {
		return fmt.Errorf("Log file was not set")
	}

	return gl.logFile.Close()
}

func (gl *TypedLog) Printf(logType TypedLogType, msg string, a ...interface{}) {
	if gl.detailed == false {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		gl.detailed = true
	}

	if len(a) == 0 {
		log.Printf(fmt.Sprintf("%s %s", string(logType), msg))
	} else {
		log.Printf(fmt.Sprintf("%s %s", string(logType), msg), a)
	}
}

func main() {
	fmt.Println("Logging test")

	log.Println("First log entry")
	//log.Fatal("First log entry")
	//log.Panicln("First log entry")

	log.Printf("Used flags: %[1]v (%08[1]bb)\n", log.Flags())
	log.Printf("Current prefix: %q\n", log.Prefix())

	lf := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	log.SetFlags(lf)
	log.Printf("Used flags: %[1]v (%08[1]bb)\n", log.Flags())

	log.SetPrefix("> ")
	log.Printf("Current prefix: %q\n", log.Prefix())

	lfile, err := os.OpenFile(
		"logs.txt",
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer lfile.Close()

	lfileInfo, err := lfile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("This file will become logs output: %q", lfileInfo.Name())
	log.SetOutput(lfile)
	log.Printf("First log in file")

	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Executable path: %q", execPath)

	glog := TypedLog{}
	err = glog.SetFile("glog.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer (func() {
		err := glog.CloseFile()
		if err != nil {
			fmt.Println("Error while closing file:", err)
		}
	})()

	glog.Printf(LErr, "This is an error\n")
	glog.Printf(LErr, "This is an error with arguments %d\n", 13)
}
