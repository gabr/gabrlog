package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Logging test")

	log.Println("First log entry")
	//log.Fatal("First log entry")
	//log.Panicln("First log entry")

	/*
		About printing int as binary.

		The '%b' makes integer to be presented as binary string.
		But it will not show leading zeros.
		In order to pad resulted string, the number representing
		width has to be putted before format char: '%8b'.
		Unfortunately the default padding (even for binary) are
		spaces. To used '0' padding prefix the format with '0'.

		Result (eight zeros padded binary format): '%08b'

		The [1] is used to index the argument of Printf() function.
	*/
	log.Printf("Used flags: %[1]v (%08[1]bb)\n", log.Flags())
	log.Printf("Current prefix: %q\n", log.Prefix())

	lf := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	log.SetFlags(lf)
	log.Printf("Used flags: %[1]v (%08[1]bb)\n", log.Flags())

	log.SetPrefix("> ")
	log.Printf("Current prefix: %q\n", log.Prefix())

	lfile, err := os.OpenFile(
		"logs.txt",
		os.O_WRONLY | os.O_APPEND | os.O_CREATE,
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
	log.Printf("First log in file");

	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Executable path: %q", execPath)
}
