package main

import (
	"log-parser/engine"
	"log-parser/parser"
)

func main() {
	parser := parser.NewParser()
	engine := engine.NewEngine(parser)
	engine.Start()
}
