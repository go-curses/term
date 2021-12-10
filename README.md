
# term
    import "github.com/go-curses/term"

This package is a clone of "github.com/pkg/term" with only a few changes
in order to support [CDK](https://github.com/kckrinke/go-cdk).


Package term manages POSIX terminals. As POSIX terminals are connected to,
or emulate, a UART, this package also provides control over the various
UART and serial line parameters.






## func CBreakMode
``` go
func CBreakMode(t *Term) error
```
CBreakMode places the terminal into cbreak mode.


## func RawMode
``` go
func RawMode(t *Term) error
```
RawMode places the terminal into raw mode.


## func Speed
``` go
func Speed(baud int) func(*Term) error
```
Speed sets the baud rate option for the terminal.



## type Term
``` go
type Term struct {
    // contains filtered or unexported fields
}
```
Term represents an asynchronous communications port.









### func Open
``` go
func Open(name string, options ...func(*Term) error) (*Term, error)
```
Open opens an asynchronous communications port.




### func (\*Term) Available
``` go
func (t *Term) Available() (int, error)
```
Available returns how many bytes are unused in the buffer.



### func (\*Term) Buffered
``` go
func (t *Term) Buffered() (int, error)
```
Buffered returns the number of bytes that have been written into the current buffer.



### func (\*Term) Close
``` go
func (t *Term) Close() error
```
Close closes the device and releases any associated resources.



### func (\*Term) DTR
``` go
func (t *Term) DTR() (bool, error)
```
DTR returns the state of the DTR (data terminal ready) signal.



### func (\*Term) Flush
``` go
func (t *Term) Flush() error
```
Flush flushes both data received but not read, and data written but not transmitted.



### func (\*Term) RTS
``` go
func (t *Term) RTS() (bool, error)
```
RTS returns the state of the RTS (data terminal ready) signal.



### func (\*Term) Read
``` go
func (t *Term) Read(b []byte) (int, error)
```
Read reads up to len(b) bytes from the terminal. It returns the number of
bytes read and an error, if any. EOF is signaled by a zero count with
err set to io.EOF.



### func (\*Term) Restore
``` go
func (t *Term) Restore() error
```
Restore restores the state of the terminal captured at the point that
the terminal was originally opened.



### func (\*Term) SendBreak
``` go
func (t *Term) SendBreak() error
```
SendBreak sends a break signal.



### func (\*Term) SetCbreak
``` go
func (t *Term) SetCbreak() error
```
SetCbreak sets cbreak mode.



### func (\*Term) SetDTR
``` go
func (t *Term) SetDTR(v bool) error
```
SetDTR sets the DTR (data terminal ready) signal.



### func (\*Term) SetOption
``` go
func (t *Term) SetOption(options ...func(*Term) error) error
```
SetOption takes one or more option function and applies them in order to Term.



### func (\*Term) SetRTS
``` go
func (t *Term) SetRTS(v bool) error
```
SetRTS sets the RTS (data terminal ready) signal.



### func (\*Term) SetRaw
``` go
func (t *Term) SetRaw() error
```
SetRaw sets raw mode.



### func (\*Term) SetSpeed
``` go
func (t *Term) SetSpeed(baud int) error
```
SetSpeed sets the receive and transmit baud rates.



### func (\*Term) Write
``` go
func (t *Term) Write(b []byte) (int, error)
```
Write writes len(b) bytes to the terminal. It returns the number of bytes
written and an error, if any. Write returns a non-nil error when n !=
len(b).



### func (\*Term) Winsz
``` go
func (t *Term) Winsz() (w, h int, err error)
```
Get the width and height of the terminal.






- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
