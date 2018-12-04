package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jung-kurt/gofpdf"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

const STOREFRONT="electronics"
var occUrl string

type PdfRenderer struct {
    pdf *gofpdf.Fpdf
}

func (p PdfRenderer) Render(w http.ResponseWriter) error {
    return p.pdf.Output(w)
}
func (PdfRenderer) WriteContentType(w http.ResponseWriter) {
    header := w.Header()
    if val := header["Content-Type"]; len(val) == 0 {
        header["Content-Type"] = []string{"applicaton/pdf"}
    }
}


func main() {
    occUrl = os.Getenv("GATEWAY_URL")
    if occUrl == "" {
        panic("Failed to set GATEWAY_URL environment variable")
    }

    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.GET("/cardTypes", func(c *gin.Context) {
        cardTypesUrl := fmt.Sprintf("%s/%s/%s",occUrl,STOREFRONT,"cardtypes")
        resp, err := http.Get(cardTypesUrl)
        if err != nil {
            c.String(500,"Internal server error %s", err)
            return
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            c.String(500,"Internal server error %s", err)
            return
        }

        c.Data(200, "application/json", body)
    })

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "hello": "World!",
            "me":gin.H{
                "this":"is",
                "things": []string{"test","for","me",},
            },
        })
    })

    r.POST("/event/:eventType", func(c *gin.Context) {
    	eventType := c.Param("eventType")
        value, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            c.String(500, "Error %s", err)
        }

        //eventModel := OrderEvent{}
        //json.Unmarshal(value, &eventModel)

        log.Printf("Event Posted: %s",eventType)
        log.Printf(string(value))

        c.JSON(200, gin.H{
            "hello": "World!",
            "me":gin.H{
                "this":"is",
                "things": []string{"test","for","me",},
            },
        })
    })

    r.GET("/pdf", func(c *gin.Context) {
        pdfId := c.Query("pdfid")
        orderId := c.Query("orderid")
        pdf := gofpdf.New("P", "mm", "A4", "")
        pdf.AddPage()
        pdf.SetFont("Arial", "B", 16)
        pdf.Cell(40, 10, "Order id: " + orderId + " PDF id: " + pdfId)
        c.Render(200, &PdfRenderer{pdf:pdf})
    })

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
