package question3
import "testing"
import "fmt"
import "time"

func TestTimeAfter(t *testing.T) {
	fmt.Println("Time before calling time.After: ", time.Now().Format(time.RFC850))
	timeAfterProg()
	fmt.Println("Time after calling time.After: ",time.Now().Format(time.RFC850))
}
