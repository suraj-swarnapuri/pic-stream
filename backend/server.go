package backend

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	path, _   = os.Getwd()
	imagePath = path + "/backend/images/"
)

func SetupServer() *gin.Engine {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/picture", picture)
	return router
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "hello-world!")
}

func picture(c *gin.Context) {
	fileName := "The_San_Antonio_Pass,_Peru_[OC]_[5464x6830].jpg"
	/*err := downloadImage("https://i.redd.it/o3v7dwlp4yn81.jpg", fileName)
	if err != nil {
		fmt.Println(err)
	}*/
	print(imagePath + fileName)
	c.FileAttachment(imagePath+fileName, fileName)
	
}
func downloadImage(url, fileName string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	print(os.Getwd())
	//Create a empty file
	file, err := os.Create(imagePath + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
