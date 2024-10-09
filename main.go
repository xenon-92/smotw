package main

import (
	"github.com/xenon-92/smotw/candidates"
	"github.com/xenon-92/smotw/filewriter"
	"github.com/xenon-92/smotw/shuffle"
	"github.com/xenon-92/smotw/workweek"
)

func main() {
	names := candidates.GetCandidateList()
	shuffle.ShuffleNames(names)
	list := workweek.GetAssignmentList(names)
	filewriter.WriteCandidateOutput(list, names)
}
