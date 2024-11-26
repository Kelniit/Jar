package utilities

import (
	"SimpleGolang/entities"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func OpenJSON(jsonroot string) ([]entities.ProductEntity, error) {
	// First Function - OpenFile
	jsonFile, errfile := OpenFile(jsonroot)
	if errfile != nil {
		log.Fatalf("Open JSON File Fail : %v", errfile)
		return nil, errfile
	}
	// Second Function - JSONFile
	result, refilerr := ReadFileJSON(jsonFile)
	if refilerr != nil {
		log.Fatalf("Read JSON File Fail : %v", refilerr)
		return nil, refilerr
	}
	// Final Function - Unmarshall JSON
	usertable, jsunerr := FinalTable(result)
	if jsunerr != nil {
		log.Fatalf("Unmarshall JSON File Fail : %v", jsunerr)
		return nil, jsunerr
	}
	return usertable, nil
}

func OpenFile(jsonroot string) (*os.File, error) {
	// Open File JSON
	jsonFile, err := os.Open(jsonroot)
	if err != nil {
		log.Fatalf("Open JSON File Fail : %v", err)
		return nil, err
	}
	return jsonFile, err
}

func ReadFileJSON(jsonFile *os.File) ([]byte, error) {
	// Read JSON File
	defer jsonFile.Close()
	result, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Read JSON File Fail : %v", err)
		return nil, err
	}
	return result, nil
}

func FinalTable(tabular []byte) ([]entities.ProductEntity, error) {
	// Final JSON Unmarshall
	var ProductTable []entities.ProductEntity
	err := json.Unmarshal(tabular, &ProductTable)
	if err != nil {
		log.Fatalf("Unmarshall JSON File Fail : %v", err)
		return nil, err
	}
	return ProductTable, nil
}
