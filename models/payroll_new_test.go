package models

import (
	"encoding/json"
	"testing"
)

func TestPayrollNewDCTASKIDBodyMarshal(t *testing.T) {
	body := PayrollNewDCTASKIDBody{
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

func TestPayrollNewBB6BTHHLBodyMarshal(t *testing.T) {
	body := PayrollNewBB6BTHHLBody{
		Bb6busmod: []*PayrollNewBb6busmod{{
			Buscod: "N03010",
			Busmod: "00001",
		}},
		Bb6cdcbhx1: []*PayrollNewBb6cdcbhx1{{
			Begtag: "Y",
			Endtag: "Y",
			Accnbr: "755900000000000",
			Ttlamt: "100.00",
			Ttlcnt: "1",
			Ttlnum: "1",
			Curamt: "100.00",
			Curcnt: "1",
			Ccynbr: "10",
			Nusage: "代发",
			Eptdat: "20260704",
			Yurref: "REF001",
			Chlflg: "Y",
		}},
		Bb6cdcdlx1: []*PayrollNewBb6cdcdlx1{{
			Trxseq: "00000001",
			Accnbr: "6214000000000000",
			Accnam: "张三",
			Trsamt: "100.00",
			Bnkflg: "Y",
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

	if decoded["bb6busmod"][0]["buscod"] != "N03010" {
		t.Fatalf("missing bb6busmod.buscod in %s", got)
	}
	if decoded["bb6cdcbhx1"][0]["yurref"] != "REF001" {
		t.Fatalf("missing bb6cdcbhx1.yurref in %s", got)
	}
	if decoded["bb6cdcdlx1"][0]["trxseq"] != "00000001" {
		t.Fatalf("missing bb6cdcdlx1.trxseq in %s", got)
	}
}
