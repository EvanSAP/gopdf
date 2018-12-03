package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jung-kurt/gofpdf"
    "net/http"
)

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
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
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

    r.GET("/pdf", func(c *gin.Context) {
        pdf := gofpdf.New("P", "mm", "A4", "")
        pdf.AddPage()
        pdf.SetFont("Arial", "B", 16)
        pdf.Cell(40, 10, "Hello, world")
        c.Render(200, &PdfRenderer{pdf:pdf})
    })

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
