package main
import "testing"

func TestHttpStatus(t *testing.T){
	LoadConfig()
	for i:=int64(1); i<=Config.Level; i++{
		for j:=int64(1); j<=Config.Level; j++{
			HttpStatus(i, j)
		}
	}
}
