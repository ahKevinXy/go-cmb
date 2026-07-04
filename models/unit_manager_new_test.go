package models

import (
	"encoding/json"
	"testing"
)

func TestUnitManagerNewDCTASKIDBodyMarshal(t *testing.T) {
	body := UnitManagerNewDCTASKIDBody{
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

func TestUnitManagerNewNTDMAADDBodyMarshal(t *testing.T) {
	body := UnitManagerNewNTDMAADDBody{
		Ntbusmody: []*UnitManagerNewNtbusmody{{Busmod: "00001"}},
		Ntdmaaddx: []*UnitManagerNewNtdmaaddx{{
			Accnbr: "755900000000000",
			Dmanbr: "000001",
			Dmanam: "test",
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

	if decoded["ntbusmody"][0]["busmod"] != "00001" {
		t.Fatalf("missing ntbusmody.busmod in %s", got)
	}
	if decoded["ntdmaaddx"][0]["dmanbr"] != "000001" {
		t.Fatalf("missing ntdmaaddx.dmanbr in %s", got)
	}
}
