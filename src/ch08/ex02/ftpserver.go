// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type status struct {
	user     string
	host     string
	port     int
	dataType string
	granted  bool
}
type command func([]string, *status) (string, int, error)

var commands = map[string]command{
	"USER": userCommand,
	"PASS": passCommand,
	"QUIT": quitCommand,
	"PORT": portCommand,
	"TYPE": typeCommand,
	"MODE": nil,
	"STRU": nil,
	"RETR": nil,
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
	input := bufio.NewScanner(conn)
	var s status
	for input.Scan() {
		line := input.Text()
		message, code, err := handleCommand(line, &s)
		if err != nil {
			fmt.Fprintf(conn, "%d %s\r\n", code, err.Error())
		}
		fmt.Fprintf(conn, "%d %s\r\n", code, message)
		if code == 221 {
			return
		}
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
	if len(op) < 1 {
		return "PASS command is required a parammeter.", 500, nil
	}
	if s.granted {
		return "You are already authed.", 503, nil
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

	if len(strings.Split(op[0], ",")) != 6 {
		return fmt.Sprintf("invalid format: %s", op[0]), 500, nil
	}

	for _, o := range op {
		io, err := strconv.Atoi(o)
		if err != nil {
			return "", 500, err
		}
		if io > 255 || io < 0 {
			return "", 500, nil
		}
	}
	p1, _ := strconv.Atoi(op[4])
	p2, _ := strconv.Atoi(op[5])
	s.host = fmt.Sprintf("%s.%s.%s.%s", op[1], op[2], op[3], op[4])
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
