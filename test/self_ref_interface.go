package main

import (
	"log"
	"strings"
)

type Logger interface {
	WithField(string, string) Logger
	Info(string)
}

func DoStuff(t Logger) {
	t.WithField("go", "1.18").Info("is awesome")
}

type MyLogger struct {
	fields []string
}

func (m *MyLogger) WithField(k string, v string) Logger {
	m.fields = append(m.fields, k+"="+v)
	return m
}

func (m *MyLogger) Info(msg string) {
	log.Printf("%s : %s", strings.Join(m.fields, ","), msg)
}

func main() {
	DoStuff(&MyLogger{fields: make([]string, 0)})
}
