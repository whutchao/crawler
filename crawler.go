package main

import (
	"github.com/whutchao/crawler/engine"
	"github.com/whutchao/crawler/samecity/parser"
	"github.com/whutchao/crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:100,
	}
	e.Run(engine.Request{
		Url:"https://bj.58.com/chuzu/?PGTID=0d100000-0000-101e-c465-298d85d88a10&ClickID=8",
		ParserFunc:parser.ParseCity,
	})
}

