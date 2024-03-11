package DB	
import (
	"fmt"
	"io/ioutil"
    "strings"
)
const (
	logFilePath = "JobSyncLogs.txt" 
)

func TopLogs() (string,error){
 file, err := ioutil.ReadFile(logFilePath)
    if err != nil {
	    return "", fmt.Errorf("unable to read file: %v", err)
	}

    // Split the file content into lines
    lines := strings.Split(string(file), "\n")

    // Get the first 50 lines
    var top50 []string
    for i := 0; i < 50 && i < len(lines); i++ {
        top50 = append(top50, lines[i])
    }

    // Combine the lines back into a single string
    response := strings.Join(top50, "\n")	
    return response , nil
}
