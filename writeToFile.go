package main

import (
	"fmt"
	"os"
)

func write_to_file(d []string, website string) {
	f, err := os.Create("urls_" + website + ".csv")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
