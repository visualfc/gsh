// Code generated by gop (Go+); DO NOT EDIT.

package gsh

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

const GopPackage = true
const _ = true
// App is project class of this classfile.
type App struct {
	fout io.Writer
	ferr io.Writer
	fin  io.Reader
	cout string
	err  error
}
//line app.gop:19:1
func (p *App) initApp() {
//line app.gop:20:1
	p.fin = os.Stdin
//line app.gop:21:1
	p.fout = os.Stdout
//line app.gop:22:1
	p.ferr = os.Stderr
}
//line app.gop:25:1
// Gop_Exec executes a shell command.
func (p *App) Gop_Exec(name string, args ...string) error {
//line app.gop:27:1
	cmd := exec.Command(name, args...) 
//line app.gop:28:1
	cmd.Stdin = p.fin
//line app.gop:29:1
	cmd.Stdout = p.fout
//line app.gop:30:1
	cmd.Stderr = p.ferr
//line app.gop:31:1
	p.err = cmd.Run()
//line app.gop:32:1
	return p.err
}
//line app.gop:35:1
// LastErr returns error of last command execution.
func (p *App) LastErr() error {
//line app.gop:37:1
	return p.err
}
//line app.gop:40:1
// Capout captures stdout of doSth() execution and save it to output.
func (p *App) Capout(doSth func()) (string, error) {
//line app.gop:42:1
	var out bytes.Buffer
//line app.gop:43:1
	old := p.fout
//line app.gop:44:1
	p.fout = &out
//line app.gop:45:1
	defer func() {
//line app.gop:46:1
		p.fout = old
	}()
//line app.gop:48:1
	doSth()
//line app.gop:49:1
	p.cout = out.String()
//line app.gop:50:1
	return p.cout, p.err
}
//line app.gop:53:1
// Output returns result of last capout.
func (p *App) Output() string {
//line app.gop:55:1
	return p.cout
}
//line app.gop:58:1
// Gopt_App_Main is main entry of this classfile.
func Gopt_App_Main(a interface {
	initApp()
}) {
//line app.gop:60:1
	a.initApp()
//line app.gop:61:1
	if
//line app.gop:61:1
	app, ok := a.(interface {
		MainEntry()
	}); ok {
//line app.gop:62:1
		app.MainEntry()
	}
}
