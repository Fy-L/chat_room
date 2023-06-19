package bytes

import (
	"encoding/json"
	"log"
	"testing"
)

type Msg struct {
	Id  int    `json:"id"`
	Mes string `json:"mes"`
}

func TestWriter(t *testing.T) {

	b := NewWriterSize(4096)
	m1 := Msg{1, "11"}
	m1_b, _ := json.Marshal(m1)
	m2 := Msg{2, "22"}
	m2_b, _ := json.Marshal(m2)

	b.Write(m1_b)
	b.Write(m2_b)

	buf := b.Buffer()

	log.Println(string(buf))
	var tmp interface{}

	json.Unmarshal(buf, &tmp)
	log.Printf("%v\n", tmp)
}
