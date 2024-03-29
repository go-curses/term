//go:build !windows && !solaris
// +build !windows,!solaris

package term

import (
	"os"

	"golang.org/x/sys/unix"

	"github.com/go-curses/term/termios"
)

// Open opens an asynchronous communications port.
func Open(name string, options ...func(*Term) error) (*Term, error) {
	fd, e := unix.Open(name, unix.O_NOCTTY|unix.O_CLOEXEC|unix.O_NDELAY|unix.O_RDWR, 0666)
	if e != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  e,
		}
	}

	orig, err := termios.Tcgetattr(uintptr(fd))
	if err != nil {
		unix.Close(fd)
		return nil, err
	}
	t := Term{name: name, fd: fd, orig: *orig}
	if err := t.SetOption(options...); err != nil {
		unix.Close(fd)
		return nil, err
	}

	if err := unix.SetNonblock(t.fd, false); err != nil {
		unix.Close(fd)
		return nil, err
	}

	return &t, nil
}

// Restore restores the state of the terminal captured at the point that
// the terminal was originally opened.
func (t *Term) Restore() error {
	return termios.Tcsetattr(uintptr(t.fd), termios.TCIOFLUSH, &t.orig)
}
