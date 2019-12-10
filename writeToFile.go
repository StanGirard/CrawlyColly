package main

import (
	"fmt"
	"os"
	"strconv"
)

func write_to_file(d []string, website string, number int) string {

	message := ""
	if len(d) > 0 {
		f, err := os.Create("urls/urls_" + website + ".csv")
		if err != nil {
			fmt.Println(err)

			f.Close()
			return "Impossible to open the file: " + "urls/urls_" + website + ".csv"
		}

		for _, v := range d {
			fmt.Fprintln(f, v)
			if err != nil {
				fmt.Println(err)
				return "Impossible to write the file: " + "urls/urls_" + website + ".csv"
			}
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return "Impossible to write the file: " + "urls/urls_" + website + ".csv"
		}
		message += "--------------------------\n"

		message += "SUCCESS: " + strconv.Itoa(number) + " URLs found for: " + website + "\n"
	} else {
		message += "ERROR: No URLS found for " + website + "\n"

	}
	message += "--------------------------\n"
	return message

}
