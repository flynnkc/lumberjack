package lumberjack

import "log"

// A Woodpile is a collection of log(ger)s
type Woodpile struct {
	Fatal *log.Logger
	Error *log.Logger
	Info  *log.Logger
	Debug *log.Logger
	level int8
}

// FATAL calls the fatal event logger which logs the message and calls os.Exit(1)
func (w *Woodpile) FATAL(s string) {
	w.Fatal.Fatalln(s)
}

// ERROR calls the error logger
func (w *Woodpile) ERROR(s string) {
	if w.level > 0 {
		w.Error.Println(s)
	}
}

// INFO calls the info logger
func (w *Woodpile) INFO(s string) {
	if w.level > 1 {
		w.Info.Println(s)
	}
}

// DEBUG calls the debug logger
func (w *Woodpile) DEBUG(s string) {
	if w.level > 2 {
		w.Debug.Println(s)
	}
}
