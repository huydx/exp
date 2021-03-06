// +build ignore

// This program generates tables.go
// go run gen.go -output tables.go
//
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"math"
)

const notFound = 0xFF

func scanRight(x uint64) uint64 {
	for k := 63; k >= 0; k -= 1 {
		if x&(1<<byte(k)) != 0 {
			return uint64(k)
		}
	}
	return notFound
}

func scanLeft(x uint64) uint64 {
	for k := byte(0); k < 64; k += 1 {
		if x&(1<<k) != 0 {
			return uint64(k)
		}
	}
	return notFound
}

func reverse(x, width uint64) (rx uint64) {
	for i := uint64(0); i < width; i += 1 {
		rx <<= 1
		rx |= x & 1
		x >>= 1
	}
	return
}

var filename = flag.String("output", "tables.go", "output file name")

func table(out io.Writer, name string, count int, fn func(v uint64) uint64) {
	line := int(math.Sqrt(float64(count)))

	fmt.Fprintf(out, "var %s = [%d]byte{\n", name, count)
	for i := 0; i < count; i += 1 {
		if i&(line-1) == 0 {
			fmt.Fprint(out, "\t")
		} else {
			fmt.Fprint(out, " ")
		}
		fmt.Fprintf(out, "0x%02x, ", fn(uint64(i)))
		if i&(line-1) == line-1 {
			fmt.Fprintln(out)
		}
	}
	fmt.Fprintln(out, "}")
	fmt.Fprintln(out)
}

func main() {
	flag.Parse()

	var buf bytes.Buffer

	fmt.Fprintln(&buf, "package bit")
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "// autogenerated by go run gen.go -output tables.go")
	fmt.Fprintln(&buf)

	table(&buf, "reverseByte", 256, func(v uint64) uint64 { return reverse(v, 8) })

	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, "const scanTableBits = 16")
	fmt.Fprintln(&buf, "const scanTableMask = (1<<scanTableBits)-1")

	table(&buf, "scanLeftTable", 256*256, scanLeft)

	table(&buf, "scanRightTable", 256*256, scanRight)

	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(*filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
