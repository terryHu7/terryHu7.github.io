package aug

import (
	"fmt"
	"testing"
	"time"
)

type ts struct {
	Id   int
	Name string
}

func TestUnmarshal(t *testing.T) {
	// var tc []ts
	// tb := []byte(`{"Id":1,"Name":"A"}`)
	// if err := json.Unmarshal(tb, &tc); err != nil {
	// 	json.Unmarshal([]byte(`[{"Id":1,"Name":"A"}]`), &tc)
	// }
	// fmt.Printf("%+v", tc)
	tpin := time.Now()
	costtime := time.Since(tpin)
	fmt.Printf("suc costtime: %v", costtime)
}
