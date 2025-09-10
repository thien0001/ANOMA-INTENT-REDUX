package handlers

import (
	"anoma_intent_app/services"
	"anoma_intent_app/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type intentRequest struct {
	Text  string `json:"text"`
	Chain string `json:"chain"` // nhận chain từ frontend
}

func IntentHandler(c *fiber.Ctx) error {
	var req intentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	parsed := services.ParseIntent(req.Text)
	// Nếu frontend gửi chain, ghi đè
	if req.Chain != "" {
		parsed.Chain = req.Chain
		parsed.GasFee = utils.FakeGasFee(parsed.Chain)
		parsed.TxHash = utils.FakeTxHash()
	}

	// update history (save to history.json)
	history := utils.LoadHistory()
	history = append([]utils.IntentResponse{parsed}, history...)
	if len(history) > 50 {
		history = history[:50]
	}
	utils.SaveHistory(history)

	// write log in text format for faster (transactions.txt)
	logEntry := fmt.Sprintf("[%s] Chain: %s | Action: %s | Token: %s | Amount: %s | TxHash: %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		parsed.Chain, parsed.Action, parsed.Token, parsed.Amount, parsed.TxHash)
	utils.AppendLog(logEntry)

	// return JSON (fields are already Action/Amount/...)
	return c.JSON(parsed)
}

func HistoryHandler(c *fiber.Ctx) error {
	return c.JSON(utils.LoadHistory())
}
