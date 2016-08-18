// Copyright (c) 2016 by akeboshi. All Rights Reserved.
// Usage: ftp -A ftp://localhost:8000/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"io/ioutil"
)

type response struct {
	message string
	code    int
	err     error
}
type status struct {
	user     string
	host     string
	path     string
	port     int
	dataType string
	granted  bool
	resp     chan response
}

type command func([]string, *status) (string, int, error)

var commands = map[string]command{
	"USER": userCommand,
	"PASS": passCommand,
	"QUIT": quitCommand,
	"PORT": portCommand,
	"TYPE": typeCommand,
	"MODE": modeCommand,
	"STRU": struCommand,
	"RETR": retrCommand,
	"CWD":  cwdCommand,
	"LIST": listCommand,
	"STOR": nil,
	"NOOP": nil,
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	var s status

	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(conn, "Fatal: %s", err.Error())
		return
	}
	s.path = dir

	input := bufio.NewScanner(conn)
	resp := make(chan response)
	s.resp = resp
	go writeResponse(conn, resp)

	fmt.Fprintf(conn, "220 Connected\n")

	for input.Scan() {
		line := input.Text()
		message, code, err := handleCommand(line, &s)
		if err != nil {
			fmt.Fprintf(conn, "%d %s\r\n", code, err.Error())
			continue
		}
		fmt.Fprintf(conn, "%d %s\r\n", code, message)
		if code == 221 {
			return
		}
	}
}

func writeResponse(conn net.Conn, resp <-chan response) {
	for msg := range resp {
		fmt.Fprintf(conn, "%d %s\r\n", msg.code, msg.message)
	}
}

func handleCommand(line string, s *status) (message string, code int, err error) {
	name, op := parseCommand(line)
	command := commands[name]
	if command == nil {
		return fmt.Sprintf("%s is unsupported.", name), 500, nil
	}
	return command(op, s)
}

func parseCommand(line string) (string, []string) {
	lines := strings.Fields(line)
	return strings.ToUpper(lines[0]), lines[1:]
}

// commands

func passCommand(op []string, s *status) (string, int, error) {
	if s.user == "anonymous" {
		return "anonymous don't require pass.", 230, nil
	}

	if s.granted {
		return "You are already authed.", 230, nil
	}

	if len(op) < 1 {
		return "PASS command is required a parammeter.", 500, nil
	}

	return "Login Failure", 530, nil
}

func userCommand(op []string, s *status) (string, int, error) {
	if len(op) < 1 {
		return "USER command is required a parammeter.", 500, nil
	}

	if s.granted {
		return "You are already authed.", 503, nil
	}

	if op[0] == "anonymous" {
		s.granted = true
		return "Logged in anonymous", 331, nil
	}

	return "Plz input " + s.user + " password.", 331, nil
}

func portCommand(op []string, s *status) (string, int, error) {
	if len(op) < 1 {
		return "PORT command is required a parammeter.", 500, nil
	}
	ops := strings.Split(op[0], ",")
	if len(ops) != 6 {
		return fmt.Sprintf("invalid format: %s", op[0]), 500, nil
	}

	for _, o := range ops[:4] {
		io, err := strconv.Atoi(o)
		if err != nil {
			return "", 500, err
		}
		if io > 255 || io < 0 {
			return "", 500, nil
		}
	}
	p1, _ := strconv.Atoi(ops[4])
	p2, _ := strconv.Atoi(ops[5])
	s.host = fmt.Sprintf("%s.%s.%s.%s", ops[0], ops[1], ops[2], ops[3])
	s.port = p1*256 + p2

	return fmt.Sprintf("Port set %s:%d", s.host, s.port), 200, nil
}

func quitCommand(op []string, s *status) (string, int, error) {
	return "Turn off the connection.", 221, nil
}

func typeCommand(op []string, s *status) (string, int, error) {
	if len(op) < 1 {
		return "TYPE command is required a parammeter.", 500, nil
	}

	if op[0] == "A" || op[0] == "I" {
		s.dataType = op[0]
		return fmt.Sprintf("Type set %s", op[0]), 200, nil
	}

	return "invalid value", 500, nil
}

func modeCommand(op []string, s *status) (string, int, error) {
	if len(op) < 1 {
		return "MODE command is required a parammeter.", 500, nil
	}

	if op[0] == "S" {
		return "Set stream mode", 200, nil
	}

	if op[0] == "B" || op[0] == "C" {
		return fmt.Sprintf("%s is unsupported mode.", op[0]), 500, nil
	}
	return fmt.Sprintf("invalid parameter %s.", op[0]), 500, nil
}

func struCommand(op []string, s *status) (string, int, error) {
	if len(op) < 1 {
		return "STRU command is required a parammeter.", 500, nil
	}

	if op[0] == "F" {
		return "Set file structure.", 200, nil
	}

	if op[0] == "R" || op[0] == "P" {
		return fmt.Sprintf("%s is unsupported mode.", op[0]), 500, nil
	}
	return fmt.Sprintf("invalid parameter %s.", op[0]), 500, nil
}

func retrCommand(op []string, s *status) (string, int, error) {
	if len(op) < 1 {
		return "RETR command is required a parammeter.", 500, nil
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	if addr == ":" {
		return "Plz set PORT.", 425, nil
	}

	path := filepath.Join(s.path, op[0])

	file, err := os.Open(path)

	defer file.Close()
	if err != nil {
		return "", 550, err
	}

	s.resp <- response{"File ok", 150, nil}

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		return "", 425, err
	}


	io.Copy(conn, file)
	return "Success receive file", 226, nil
}

func cwdCommand(op []string, s *status) (string, int, error) {
	if !s.granted {
		return "Plz login", 530, nil
	}

	path := filepath.Join(s.path, op[0])

	if f, err := os.Stat(path); err != nil || !f.IsDir() {
		return "Failed dir changed", 550, err
	}

	current, err := os.Getwd()
	if err != nil {
		return "", 500, err
	}
	if !filepath.HasPrefix(path, current) {
		s.path = path
		return "Dir is " + path, 250, nil
	}

	s.path = path
	return "Dir is " + path, 250, nil
}

func listCommand(op []string, s *status) (string, int, error) {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	if addr == ":" {
		return "Plz set PORT.", 425, nil
	}

	path := s.path
	if len(op) > 0{
		path = filepath.Join(path, op[0])
	}

	stat, err := os.Stat(path)
	if err != nil {
		return s.path+" cant open.",550, err
	}
	var str string = ""
	if stat.IsDir() {
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			return path + "cant read.", 550, err
		}

		for _, d := range dir {
			str += d.Name() + " "
		}
	} else {
		str = path
	}
	s.resp <- response{"Dir ok", 150, nil}

	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		return "", 425, err
	}

	_, err = fmt.Fprintf(conn, "%s\r\n", str)
	if err != nil {
		return path + "cant send.", 550, err
	}

	return "Success receive", 250, nil
}
