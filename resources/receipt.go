package resources

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fetch-rewards/models"
	"fetch-rewards/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ReceiptResource struct {
	Points map[string]int
}

func NewReceiptResource() *ReceiptResource {
	return &ReceiptResource{
		Points: map[string]int{},
	}
}

func (rr *ReceiptResource) GetReceiptPointsHandler(c echo.Context) error {
	id := c.Param("id")
	points, receiptExists := rr.Points[id]

	if !receiptExists {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, map[string]int{"points": points})
}

func (rr *ReceiptResource) ProcessReceiptHandler(c echo.Context) error {
	fmt.Println("Processing receipt")
	var receipt models.Receipt

	fmt.Println("Extracting receipt json")
	if err := c.Bind(&receipt); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "receipt is invalid"})
	}

	jsonData, err := json.Marshal(receipt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "receipt is invalid"})
	}

	err = receipt.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "receipt is invalid"})
	}

	points, err := processReceipt(receipt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "receipt is invalid"})
	}

	fmt.Println("Hashing receipt json to generate id")
	receipt.ID = generateUUID(jsonData)
	fmt.Println("id generated: ", receipt.ID)

	rr.Points[receipt.ID] = points
	fmt.Println("Points after processing receipt: ", points)

	return c.JSON(http.StatusOK, map[string]string{"id": receipt.ID})
}

func processReceipt(receipt models.Receipt) (int, error) {
	pointCalculator := utils.NewPointsCalculator(receipt)
	err := pointCalculator.CalculatePoints()
	if err != nil {
		return 0, err
	}

	return pointCalculator.Points, nil
}

func generateUUID(jsonData []byte) string {
	hash := md5.Sum(jsonData)
	hashString := hex.EncodeToString(hash[:])
	receiptId := fmt.Sprintf(
		"%s-%s-%s-%s-%s",
		hashString[:8],
		hashString[8:12],
		hashString[12:16],
		hashString[16:20],
		hashString[20:])

	return receiptId
}

// SetupRoutes configures the routes for the TodoResource
func (rr *ReceiptResource) SetupRoutes(e *echo.Echo) {
	receiptGroup := e.Group("/receipts")
	{
		receiptGroup.GET("/:id/points", rr.GetReceiptPointsHandler)
		receiptGroup.POST("/process", rr.ProcessReceiptHandler)
	}
}
