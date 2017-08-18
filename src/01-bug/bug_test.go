package main
import "testing"

func TestBug(t *testing.T){
	if testing.Short(){
		Bug()
		t.Skip("Skipping Bug...")
	}

}
