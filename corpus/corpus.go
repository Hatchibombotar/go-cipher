package main

import (
	_ "embed"
)

//go:embed corpus.txt
var corpus string

func GetRawCorpus() string {
	return corpus
}
