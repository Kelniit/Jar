package controllers

import (
	"SimpleGolang/entities"
	"SimpleGolang/utilities"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

func MainFile(C *gin.Context) {
	// Main File
	C.JSON(200, gin.H{"result": "Hallo ! Simple Application !"})
}

func GetAllProduct(C *gin.Context) {
	// Get All Product
	result, err := utilities.OpenJSON("product.json")
	if err != nil {
		log.Fatalf("Fail : %v", err)
	}
	C.JSON(200, result)
}

func GetProductID(C *gin.Context) {
	// Get Product ID
	PID := C.Param("ProductID")
	allresult, err := utilities.OpenJSON("product.json")
	if err != nil {
		log.Fatalf("Fail : %v", err)
	}
	for _, result := range allresult {
		if result.ProductID == PID {
			C.JSON(200, result)
			return
		}
	}
	C.JSON(404, gin.H{"result": "Product Unavailable !"})
}

func TambahProduct(C *gin.Context) {
	// Add Product
	var DataProduct entities.ProductEntity
	if input_err := C.BindJSON(&DataProduct); input_err != nil {
		log.Fatalf("Fail : %v", input_err)
	}

	resulta, open_error := utilities.OpenJSON("product.json")
	if open_error != nil {
		log.Fatalf("Fail : %v", open_error)
	}

	resulta = append(resulta, DataProduct)
	result_marshall, _ := json.Marshal(resulta)
	ioutil.WriteFile("product.json", result_marshall, os.ModePerm)
}
