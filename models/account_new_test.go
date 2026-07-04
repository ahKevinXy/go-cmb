package models

import (
	"encoding/json"
	"testing"
)

func TestAccountNewDCTASKIDBodyMarshal(t *testing.T) {
	body := AccountNewDCTASKIDBody{
		Taskid: "task-id",
		Qwenab: "false",
	}

	got, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"taskid":"task-id","qwenab":"false"}`
	if string(got) != want {
		t.Fatalf("unexpected JSON\nwant: %s\n got: %s", want, got)
	}
}

func TestAccountNewNTQACINFBodyMarshal(t *testing.T) {
	body := AccountNewNTQACINFBody{
		Ntqacinfx: []*AccountNewNtqacinfx{{
			Bbknbr: "69",
			Accnbr: "755900000000000",
			Ccynbr: "10",
		}},
	}

	got, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	var decoded map[string][]map[string]string
	if err := json.Unmarshal(got, &decoded); err != nil {
		t.Fatal(err)
	}

	if decoded["ntqacinfx"][0]["accnbr"] != "755900000000000" {
		t.Fatalf("missing ntqacinfx.accnbr in %s", got)
	}
	if decoded["ntqacinfx"][0]["bbknbr"] != "69" {
		t.Fatalf("missing ntqacinfx.bbknbr in %s", got)
	}
}
