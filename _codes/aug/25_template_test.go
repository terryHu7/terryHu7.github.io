package aug

import (
	"fmt"
	"testing"
)

//Everybody is DA GONG REN
/////////////////////////////////////
type WorkerInterface interface {
	GetUp()
	Work()
	Sleep()
}
type Worker struct {
	WorkerInterface //extend
}

func NewWorker(w WorkerInterface) *Worker {
	return &Worker{WorkerInterface: w}
}

func (w *Worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

/////////////////////////////////////

type Coder struct {
}

func (c *Coder) GetUp() {
	fmt.Println("Coder GetUp ")
}

func (c *Coder) Work() {
	fmt.Println("Coder Work ")
}

func (c *Coder) Sleep() {
	fmt.Println("Coder Sleep ")
}

func TestWorker_Daily(t *testing.T) {
	worker := NewWorker(&Coder{})
	worker.Daily()
}
