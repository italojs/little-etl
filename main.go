package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	//open the file train.csv
	file, err := os.Open("train.csv")
	if err != nil {
		log.Fatalln("Error opening the file train.csv: ", err)
		return
	}
	defer file.Close()

	//read the opened file train.csv
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error read file train.csv: ", err)
	}

	var wg sync.WaitGroup
	//wait all lines be executed
	wg.Add(len(record))

	// for each number in this csv, change the values to smaller values
	for line := range record {
		go func(line int) {
			defer wg.Done()
			for colum := range record[line] {
				//don't get the first colum because it's the classification value
				if colum > 0 {
					linei, err := strconv.Atoi(record[line][colum])
					if err != nil {
						log.Fatalf("Error to converter the value %s to integer: ", record[line][colum], err)
						os.Exit(2)
					}
					//change values to smaller values
					if linei > 0 {
						if linei < 100 {
							record[line][colum] = "1"
						} else if linei < 200 {
							record[line][colum] = "2"
						} else {
							record[line][colum] = "3"
						}
					}
				}
			}
		}(line)
	}

	wg.Wait()

	newFile, err := os.Create("data.csv")
	if err != nil {
		log.Fatalln("Error to create the file data.csv: ", err)
	}
	defer newFile.Close()

	w := csv.NewWriter(newFile)
	defer w.Flush()

	w.WriteAll(record)

	if w.Error() != nil {
		log.Fatalln("error writing data.csv:", w.Error())
	}

}
