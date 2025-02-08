package main

import "fmt"

/*
Interface Segregation Principle (ISP):
- Clients should not be forced to depend on interfaces they do not use.
- Instead of creating large, general-purpose interfaces, we should define smaller, specific interfaces.
*/

// Define small, focused interfaces.
type Reader interface {
	Read() string
}

type Writer interface {
	Write(content string)
}

type Printer interface {
	Print() string
}

// TextDocument implements all three interfaces.
type TextDocument struct {
	content string
}

func (d *TextDocument) Read() string {
	return d.content
}

func (d *TextDocument) Write(content string) {
	d.content = content
}

func (d *TextDocument) Print() string {
	return "Printing: " + d.content
}

// ReadOnlyDocument only implements Reader, avoiding unnecessary methods.
type ReadOnlyDocument struct {
	content string
}

func (d *ReadOnlyDocument) Read() string {
	return d.content
}

// Function that works with any Reader.
func DisplayContent(r Reader) {
	fmt.Println("Content:", r.Read())
}

// Function that works with any Printer.
func PrintDocument(p Printer) {
	fmt.Println(p.Print())
}

// func main() {
// 	textDoc := &TextDocument{}
// 	textDoc.Write("Hello, ISP!")

// 	readOnlyDoc := &ReadOnlyDocument{content: "Read-Only Content"}

// 	DisplayContent(textDoc)     // ✅ Works
// 	DisplayContent(readOnlyDoc) // ✅ Works (only requires Read)

// 	PrintDocument(textDoc) // ✅ Works (implements Printer)
// 	// PrintDocument(readOnlyDoc) // ❌ Compile error: ReadOnlyDocument doesn't implement Printer
// }
