package matrix

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type MatrixResponse struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}

type StatsResponse struct {
    Max         float64 `json:"max"`
    Min         float64 `json:"min"`
    Average     float64 `json:"average"`
    Sum         float64 `json:"sum"`
    IsQDiagonal bool    `json:"isQDiagonal"`
    IsRDiagonal bool    `json:"isRDiagonal"`
}

func QRHandler(c *fiber.Ctx) error {
	var req MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	// Descomponemos en q y r
	q, r, err := QRDecomposition(req.Matrix)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "QR decomposition failed")
	}

	// Enviar a API Node
	stats, err := sendToNodeAPI(q, r)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Node API failed"})
	}

	return c.JSON(fiber.Map{
		"Q":      q,
		"R":      r,
		"stats":  stats,
	})

}

func sendToNodeAPI(q, r [][]float64) (StatsResponse, error) {
    payload := map[string][][]float64{"Q": q, "R": r}
    jsonData, _ := json.Marshal(payload)

    resp, err := http.Post("http://localhost:4000/api/v1/matrix/process", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
		log.Println("❌ Failed post api:", err)
        return StatsResponse{}, err
    }
    defer resp.Body.Close()

    var stats StatsResponse
    if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		log.Println("❌ Failed to decode:", err)
        return StatsResponse{}, err
    }

    return stats, nil
}

