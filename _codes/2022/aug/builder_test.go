package aug

import (
	"fmt"
	"testing"
)

type Example struct {
	Value int
}

type ExampleBuilder struct {
	*Example
}

// user-friendly
func (builder *ExampleBuilder) SetExampleValue() *ExampleBuilder {
	builder.Example.Value = 728
	return builder
}

func (builder ExampleBuilder) GetExample() *Example {
	return builder.Example
}

func TestBuilder(t *testing.T) {

	newbuilder := &ExampleBuilder{&Example{}}

	newexample := newbuilder.SetExampleValue().GetExample()

	fmt.Printf("newexample: %+v\n", newexample)

}
