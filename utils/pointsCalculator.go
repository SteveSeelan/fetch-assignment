package utils

import (
	"errors"
	"fetch-rewards/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type PointsCalculator struct {
	Receipt models.Receipt
	Points  int
}

func NewPointsCalculator(receipt models.Receipt) *PointsCalculator {
	return &PointsCalculator{
		Receipt: receipt,
		Points:  0,
	}
}

func (pc *PointsCalculator) CalculatePoints() error {
	calculations := []struct {
		name string
		fn   func() (int, error)
	}{
		{"calculateAlphaNumeric", func() (int, error) { return pc.calculateAlphaNumeric(pc.Receipt.Retailer) }},
		{"calculateRound", func() (int, error) { return pc.calculateRound(pc.Receipt.Total) }},
		{"calculateMultiple", func() (int, error) { return pc.calculateMultiple(pc.Receipt.Total) }},
		{"calculateDualItems", func() (int, error) { return pc.calculateDualItems(pc.Receipt.Items) }},
		{"calculateDescription", func() (int, error) { return pc.calculateDescription(pc.Receipt.Items) }},
		{"calculateOddDate", func() (int, error) { return pc.calculateOddDate(pc.Receipt.PurchaseDate) }},
		{"calculatePastDate", func() (int, error) { return pc.calculatePastDate(pc.Receipt.PurchaseTime) }},
	}
	var errorMessages []string

	for _, calc := range calculations {
		points, err := calc.fn()
		if err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("%s: %s", calc.name, err.Error()))
		}
		pc.Points += points
	}
	if len(errorMessages) > 0 {
		return errors.New(strings.Join(errorMessages, "\n"))
	}

	return nil
}

func (pc *PointsCalculator) calculateAlphaNumeric(retailer string) (int, error) {
	points := 0
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}
	return points, nil
}

func (pc *PointsCalculator) calculateRound(total string) (int, error) {
	dollarTotal, err := strconv.ParseFloat(total, 64)
	points := 0
	if err != nil {
		return points, err
	}
	if math.Mod(dollarTotal, 1) == 0 {
		points += 50
	}
	return points, nil
}

func (pc *PointsCalculator) calculateMultiple(total string) (int, error) {
	dollarTotal, err := strconv.ParseFloat(total, 64)
	points := 0
	if err != nil {
		return points, err
	}
	if math.Mod(dollarTotal, 0.25) == 0 {
		points += 25
	}
	return points, nil
}

func (pc *PointsCalculator) calculateDualItems(items []models.Item) (int, error) {
	pointMultiplier := len(items) / 2
	points := 5 * pointMultiplier
	return points, nil
}

func (pc *PointsCalculator) calculateDescription(items []models.Item) (int, error) {
	points := 0
	for _, item := range items {
		trimmedShortDesc := strings.TrimSpace(item.ShortDescription)
		trimmedLength := float64(len(trimmedShortDesc))

		if math.Mod(trimmedLength, 3) == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return points, err
			}

			points += int(math.Ceil(price * 0.2))
		}
	}
	return points, nil
}

func (pc *PointsCalculator) calculateOddDate(purchaseDate string) (int, error) {
	date, err := time.Parse(time.DateOnly, purchaseDate)
	points := 0
	if err != nil {
		return points, err
	}

	if math.Mod(float64(date.Day()), 2) != 0 {
		points += 6
	}
	return points, nil
}

func (pc *PointsCalculator) calculatePastDate(purchaseTime string) (int, error) {
	points := 0
	layout := "15:04"
	militaryTime, err := time.Parse(layout, purchaseTime)
	if err != nil {
		return points, err
	}

	startTime, _ := time.Parse(layout, "14:00")
	endTime, _ := time.Parse(layout, "16:00")
	if militaryTime.After(startTime) && militaryTime.Before(endTime) {
		points += 10
	}
	return points, nil
}
