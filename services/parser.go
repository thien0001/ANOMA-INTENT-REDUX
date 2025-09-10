package services

import (
	"anoma_intent_app/utils"
	"regexp"
	"strings"
)

// ParseIntent nhận input người dùng và trả utils.IntentResponse
func ParseIntent(input string) utils.IntentResponse {
	text := strings.ToLower(strings.TrimSpace(input))

	resp := utils.IntentResponse{
		Action:  "Hollow",
		Amount:  "1",
		Token:   "ETH",
		Chain:   "Ethereum",
		TxHash:  utils.FakeTxHash(),
		GasFee:  utils.FakeGasFee("ethereum"),
		Message: "Processed mock intent",
	}

	// Get action
	if strings.Contains(text, "buy") || strings.Contains(text, "buy") {
		resp.Action = "buy"
	} else if strings.Contains(text, "sell") || strings.Contains(text, "sell") {
		resp.Action = "sell"
	} else if strings.Contains(text, "transfer") || strings.Contains(text, "transfer") {
		resp.Action = "transfer"
	}

	// Get token
	tokens := []string{"eth", "btc", "usdt", "matic", "bnb", "usdc"}
	for _, t := range tokens {
		if strings.Contains(text, t) {
			resp.Token = strings.ToUpper(t)
			break
		}
	}

	// Get chain
	switch {
	case strings.Contains(text, "polygon") || strings.Contains(text, "matic"):
		resp.Chain = "Polygon"
	case strings.Contains(text, "bsc") || strings.Contains(text, "binance"):
		resp.Chain = "BSC"
	case strings.Contains(text, "arbitrum"):
		resp.Chain = "Arbitrum"
	case strings.Contains(text, "optimism"):
		resp.Chain = "Optimism"
	default:
		resp.Chain = "Ethereum"
	}

	// Get quantity (get first number that appears)
	reAmount := regexp.MustCompile(`(\d+(\.\d+)?)`)
	if m := reAmount.FindStringSubmatch(text); len(m) > 0 {
		resp.Amount = m[1]
	}

	// Fake gas fee & tx hash theo chain
	resp.GasFee = utils.FakeGasFee(resp.Chain)
	resp.TxHash = utils.FakeTxHash()

	return resp
}
