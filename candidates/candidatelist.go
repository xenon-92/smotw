package candidates

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetCandidateList() []string {
	const path string = "candidate_list.txt"
	exists := checkFileExists(path)
	if !exists {
		errorMsg := `Missing input to the Program....
			The list of candidates for Scrum Master Of The Week (SMOTW)
			should be present in a file named 'candidate_list.txt', to this root where you are running this exe
			The candidates entered in the file should be seperated by a delimiter '|' .Below is an example
			Joe|Joey|Joel|John|Jonathan|Jacky|Jim|Jack
		`
		fmt.Println(errorMsg)
		fmt.Println()
		fmt.Println()
		fmt.Println("Press enter to close the program")
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		os.Exit(1)
	}
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	names := strings.Split(string(contents), "|")
	return names
}

func checkFileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
