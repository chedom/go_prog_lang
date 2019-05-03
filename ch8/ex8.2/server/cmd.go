package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

var errNotFoundCmd = errors.New("no such command")

var BufferSize = 1000

type cmd struct {
	currDir string
	conn    net.Conn
}

func NewCmd(conn net.Conn) *cmd {
	dir, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	return &cmd{currDir: dir, conn: conn}
}

func (cmd *cmd) Listen() {
	scanner := bufio.NewScanner(cmd.conn)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Printf("scanner error: %v", err)
		}
		str := scanner.Text()
		log.Printf("Receive message: %s", str)
		words := strings.Split(str, " ")
		cmd.execute(words[0], words[1:])
	}
}

func (cmd *cmd) execute(command string, args []string) error {
	switch command {
	case "ls":
		return cmd.ls()
	case "cd":
		return cmd.cd(args[0])
	case "get":
		return cmd.get(args[0])
	case "close":
		return cmd.close()
	default:
		return errNotFoundCmd
	}
}

func (cmd *cmd) ls() error {
	files, err := ioutil.ReadDir(cmd.currDir)
	if err != nil {
		log.Printf("ls dir: %v\n", err)
		return err
	}

	for _, f := range files {
		_, err := io.WriteString(cmd.conn, f.Name()+"\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd *cmd) cd(p string) error {
	newDir := cmd.currDir + "/" + p
	err := os.Chdir(newDir)

	if err == nil {
		cmd.currDir = newDir
		fmt.Fprint(cmd.conn, newDir+"\n")
		return nil
	}

	log.Printf("cd error: %v\n", err)

	return err
}

func (cmd *cmd) get(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("cant open file: %s, error: %v\n", filename, err)
		return err
	}
	defer f.Close()

	buf := make([]byte, BufferSize)

	for {
		count, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("error while reading file: %v\n", err)
				return err
			}
		}

		if count == 0 {
			break
		}

		_, err = cmd.conn.Write(buf[:count])
		if err != nil {
			log.Printf("cant send file: %v\n", err)
			return err
		}
	}

	return nil
}

func (cmd *cmd) close() error {
	if err := cmd.conn.Close(); err != nil {
		if err != nil {
			log.Printf("close cmd: %v\n", err)
			return err
		}
	}

	return nil
}
