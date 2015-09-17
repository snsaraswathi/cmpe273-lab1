package question3
import "time"

func timeAfterProg() {
	select{
		case <-time.After(30 * time.Second):
	}
}
