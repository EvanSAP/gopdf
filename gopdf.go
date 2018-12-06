package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/gopdf/client"
	"github.com/gopdf/client/orders"
	"github.com/gopdf/commerce"
	"github.com/gopdf/models"
	"github.com/jung-kurt/gofpdf"
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	_ "github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zsais/go-gin-prometheus"
	_ "golang.org/x/net/context/ctxhttp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const STOREFRONT = "electronics"

var (
	client           *apiclient.Commercesuite
	occUrl           url.URL
	gatewayUrlString string
	whGatewayUrl     string
	opsProcessed     = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gopdf_events_processed",
		Help: "The total number of processed events",
	})
	pdfsGenerated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gopdf_pdfs_generated",
		Help: "The total number of processed events",
	})
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

	whGatewayUrl = os.Getenv("WH_GATEWAY_URL")
	if whGatewayUrl == "" {
		panic("Failed to set WH_GATEWAY_URL environment variable")
	}
	log.Printf("Warehouse Gateway URL: %s", whGatewayUrl)

	// Ensure the GATEWAY_URL environment variable is set
	// typically this comes in from a secret for a bound servicefactory instance
	gatewayUrl := os.Getenv("GATEWAY_URL")
	if gatewayUrl == "" {
		panic("Failed to set GATEWAY_URL environment variable")
	}
	log.Printf("Gateway URL: %s", gatewayUrl)
	gatewayUrlString = gatewayUrl

	occUrl, err := url.Parse(gatewayUrl)
	if err != nil {
		log.Panicf("Failed to set GATEWAY_URL environment variable, %s", err)
	}

	// Define custom datetime format
	formats := strfmt.NewFormats()
	formats.Add("datetime", &models.CustomDateTime{}, nil)

	// create the transport
	transport := httptransport.New(occUrl.Host, "", []string{occUrl.Scheme})

	// create the API client, with the transport
	client = apiclient.New(transport, formats)

	// Enable Prometheus metrics on port :8081/metrics
	// Kyma likes metrics on 8081
	promhttp.Handler()
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8081", nil)

	// Initialize the real API handlers
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/cardTypes", func(c *gin.Context) {
		cardTypesUrl := fmt.Sprintf("%s/%s/%s", occUrl, STOREFRONT, "cardtypes")
		resp, err := http.Get(cardTypesUrl)
		if err != nil {
			c.String(500, "Internal server error %s", err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.String(500, "Internal server error %s", err)
			return
		}

		c.Data(200, "application/json", body)
	})

	r.GET("/orders", getOrders)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "World!",
			"me": gin.H{
				"this":   "is",
				"things": []string{"test", "for", "me",},
			},
		})
	})

	r.POST("/event/:eventType", func(c *gin.Context) {

		// record that we received an event
		opsProcessed.Inc()

		eventType := c.Param("eventType")
		value, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(500, "Error %s", err)
		}

		log.Printf("Event Posted %s -- %s", eventType, string(value))
		//consigment.consignmentprocessing
		//v1
		//Consignment Processing Event v1
		switch eventType {
		case "order.created":
			orderEvent := commerce.UnmarshalOrderCreatedEvent(value)
			log.Printf("Handling order.created. order number %s", orderEvent.OrderCode)
			go handleOrderCreated(orderEvent.OrderCode)
		case "consigment.consignmentprocessing":
			orderEvent := commerce.UnmarshalOrderCreatedEvent(value)
			log.Printf("Handling consigment.consignmentprocessing. order number %s", orderEvent.OrderCode)
			go handleOrderCreated(orderEvent.OrderCode)
		default:
			log.Printf("Unrecognized Event received: %s", eventType)
			break
		}

		c.String(200, "Event Received OK")
	})

	r.GET("/pdf", func(c *gin.Context) {
		pdfsGenerated.Inc()

		pdfId := c.Query("pdfid")
		orderId := c.Query("orderid")
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "Order id: "+orderId+" PDF id: "+pdfId)
		c.Render(200, &PdfRenderer{pdf: pdf})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func getOrders(c *gin.Context) {
	ordersUrl := fmt.Sprintf("%s/%s/users/%s/orders", gatewayUrlString, STOREFRONT, "xuhao318@gmail.com")
	log.Printf("Orders url: %s", ordersUrl)
	resp, err := http.Get(ordersUrl)
	if err != nil {
		c.String(500, "Internal server error %s", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(500, "Internal server error %s", err)
		return
	}

	c.Data(200, "application/json", body)
}

func handleOrderCreated(orderCode string) {

	log.Printf("Waiting 80 seconds to process order %s", orderCode)
	time.Sleep(time.Second * 80)
	log.Printf("Continuing to process order %s", orderCode)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	params := orders.GetOrderUsingGETParams{
		BaseSiteID: STOREFRONT,
		Code:       orderCode,
		Context:    ctx,
	}
	defer cancel()

	result, err := client.Orders.GetOrderUsingGET(&params, nil)
	if err != nil {
		switch v := err.(type) {
		case *time.ParseError:
			log.Printf("Parse error %v", v)
		case *runtime.APIError:
			log.Printf("Runtime error %v", v)
			errResponse := v.Response.(runtime.ClientResponse)
			responseBody, _ := ioutil.ReadAll(errResponse.Body())

			log.Printf("Unable to run rest request:  %v, %v", responseBody, err)
		default:
			log.Printf("Unknown error %v", v)

		}
		return
	}

	value, _ := json.MarshalIndent(result, "", "  ")
	log.Printf("Received order %s", value)

	for _, consignment := range result.Payload.Consignments {
		sendMail(result.Payload.User.UID,
			result.Payload.User.Name,
			result.Payload.Code,
			consignment.Code)
		consignmentUrl := whGatewayUrl + "/consignments/" + consignment.Code + "/confirm-shipping"
		log.Printf("Consignment shipped  URL: %s", consignmentUrl)

		resp, err :=
			http.Post(consignmentUrl, "application/json", strings.NewReader(""))
		log.Printf("Wharehousing Consignment Shipped response: %v, %v", resp, err)

	}

}

func sendMail(email string, name string, orderId string, pdfId string) {

	pdfUrl := fmt.Sprintf("https://gopdf.sa-hackathon-09.cluster.extend.sap.cx/pdf?orderid=%s&pdfid=%s",
		orderId, pdfId)

	publicKey := "9054deeb766cac10b61b46d52b07b011"
	secretKey := "-----secret---key----here-------"

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	param := &mailjet.InfoSendMail{
		FromEmail: "evan@evanhegyi.com",
		FromName:  "Evan Hegyi",
		Recipients: []mailjet.Recipient{
			mailjet.Recipient{
				Email: email,
			},
		},
		Subject:  "The PDF your ordered, just for you " + name,
		TextPart: fmt.Sprintf("Hi there %s, Donwnload Here: \n%s", name, pdfUrl),
	}
	res, err := mj.SendMail(param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
}
