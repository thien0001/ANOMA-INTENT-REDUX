package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type IntentResponse struct {
	Action  string `json:"Action"`
	Amount  string `json:"Amount"`
	Token   string `json:"Token"`
	Chain   string `json:"Chain"`
	GasFee  string `json:"GasFee"`
	TxHash  string `json:"TxHash"`
	Message string `json:"Message"`
}

var (
	historyFile = "history.json"
	logFile     = "transactions.txt"
	mu          sync.Mutex
)

func LoadHistory() []IntentResponse {
	mu.Lock()
	defer mu.Unlock()

	if _, err := os.Stat(historyFile); os.IsNotExist(err) {
		return []IntentResponse{}
	}

	data, err := ioutil.ReadFile(historyFile)
	if err != nil {
		fmt.Println("Không đọc được file history:", err)
		return []IntentResponse{}
	}

	var history []IntentResponse
	if err := json.Unmarshal(data, &history); err != nil {
		fmt.Println("Lỗi parse JSON history:", err)
		return []IntentResponse{}
	}
	return history
}

func SaveHistory(history []IntentResponse) {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		fmt.Println("Lỗi JSON Marshal history:", err)
		return
	}

	if err := ioutil.WriteFile(historyFile, data, 0644); err != nil {
		fmt.Println("Lỗi ghi history file:", err)
	}
}

func AppendLog(entry string) {
	mu.Lock()
	defer mu.Unlock()
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Lỗi mở log file:", err)
		return
	}
	defer f.Close()
	f.WriteString(entry)
}

func FakeTxHash() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "0xDEADBEEF"
	}
	return "0x" + hex.EncodeToString(b)
}

func FakeGasFee(chain string) string {
	switch strings.ToLower(chain) {
	case "polygon", "matic":
		return "0.0020 MATIC"
	case "bsc", "binance", "bnb":
		return "0.0010 BNB"
	case "arbitrum":
		return "0.0005 ETH"
	case "optimism":
		return "0.0006 ETH"
	default:
		return "0.00042 ETH"
	}
}
