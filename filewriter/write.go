package filewriter

import (
	"fmt"
	"os"

	"github.com/xenon-92/smotw/candidates"
)

func getColumnNames() []string {
	return []string{"Month", "Name", "WorkWeek"}
}

func WriteCandidateOutput(data [][]string, names []string) {
	const fileName = "smotw.txt"
	//delete the already existing file
	if _, err := os.Stat(fileName); err == nil {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("Error deleting the file:", err)
		}
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	writeOriginalList(file)
	writeShuffledList(file, names)
	for i, header := range getColumnNames() {
		if i != 0 {
			fmt.Fprint(file, "\t\t\t") // Use tab as separator
		}
		fmt.Fprint(file, header)
	}
	fmt.Fprintln(file) // Newline after headers
	for _, row := range data {
		for i, col := range row {
			if i != 0 {
				fmt.Fprint(file, "\t") // Use tab as separator
			}
			fmt.Fprint(file, col)
		}
		fmt.Fprintln(file) // Newline after each row
	}
}

func writeOriginalList(file *os.File) {
	fmt.Fprint(file, "Original candidateList:")
	fmt.Fprintln(file)
	for i, header := range candidates.GetCandidateList() {
		if i != 0 {
			fmt.Fprint(file, "\t") // Use tab as separator
		}
		fmt.Fprint(file, header)
	}
	fmt.Fprintln(file)
	fmt.Fprintln(file)
}
func writeShuffledList(file *os.File, names []string) {
	fmt.Fprint(file, "Shuffled candidateList:")
	fmt.Fprintln(file)
	for i, header := range names {
		if i != 0 {
			fmt.Fprint(file, "\t") // Use tab as separator
		}
		fmt.Fprint(file, header)
	}
	fmt.Fprintln(file)
	fmt.Fprintln(file)
}
