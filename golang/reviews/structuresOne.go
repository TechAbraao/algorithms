package reviews
import (	
	"fmt"
	"encoding/json"
)


type DataJson struct {
	Name string 
	LastName string
	Age int
	
}

type Job struct {
	Title string `json:"title"`
	Enterprise string `json:"enterprise"`
}
	
func main() {
	var data DataJson = DataJson{
		Name: "John",
		LastName: "Doe",
		Age: 30,
	}

	var job Job = Job{
		Title: "Software Engineer",
		Enterprise: "Tech Corp",
	}

	fmt.Println(data)
	jobJson, jobJsonErr := json.Marshal(job)

	if jobJsonErr != nil {
		fmt.Println("Error: ", nil)
		return
	}

	JobJsonTwo, _ := json.MarshalIndent(job, "", " ")

	fmt.Println(string(jobJson))
	fmt.Println(string(JobJsonTwo))
}