rabinkarp
=========

Rabin-Karp implementation for [Golang's hash package](http://golang.org/pkg/hash)

Installation
------------

    go get github.com/jbaikge/rabinkarp

Usage
-----

Most useful for generating numeric hashes for short []byte slices (words):

    package main

    import (
    	"fmt"
    	"github.com/jbaikge/rabinkarp"
    )

    func main() {
    	h := rabinkarp.New32()
    	words := [][]byte{
    		[]byte(`monkey`),
    		[]byte(`elephant`),
    	}
    	for _, word := range words {
    		h.Write(word)
    		fmt.Printf("%s: %d\n", word, h.Sum32())
    		h.Reset()
    	}
    }

Output:

    monkey: 663992827
    elephant: 1878464849
