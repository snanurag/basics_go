/*
	Try escape analysis using go build -gcflags '-m' ./main.go
	Output :
		$ go build -gcflags '-m' ./main.go
		# command-line-arguments
		./main.go:18:6: can inline CreateNewStateProcessor
		./main.go:22:6: can inline CreateNewStateProcessor2
		./main.go:19:26: NewStateProcessor literal escapes to heap
*/
package escape_analysis_on_interface

type WFStateProcessor interface {
}

type NewStateProcessor struct {
}

func CreateNewStateProcessor() WFStateProcessor {
	return NewStateProcessor{}
}

func CreateNewStateProcessor2() NewStateProcessor {
	return NewStateProcessor{}
}
