

# 🔮 ANOMA INTENT: REDUX 🔮

## 🌟 Overview
This project is a **Frontend + Backend (mock API)** demo showcasing the concept of **Intent-Centric Applications** in the **Anoma ecosystem**.  
Users can enter an **Intent** (e.g., *"Buy 1 ETH on Polygon"*) and the system will:

- Parse and process the intent using a **Golang backend**.  
- Display the parsed result on the **frontend UI**.  
- Store the transaction in a **multi-chain transaction history** (Ethereum, BSC, Polygon, Arbitrum, Optimism).  

### Objectives
- Demonstrate **Intent-Centric Architecture**.  
- Provide a smooth **multi-chain user experience**.  
- Optimize **UI/UX** with animations, toast notifications, and a transaction explorer.  

---

## 📂 Project Structure


anoma-intent-redux/
├── handlers/ # API handlers (intent, history)
│ └── intent.go
├── services/ # Business logic
│ └── parser.go
├── utils/ # Storage, helpers
│ └── storage.go
├── static/ # Frontend (UI demo)
│ ├── index.html
│ ├── explorer.html
│ ├── explorer.js
│ └── logo.jpg
├── history.json # Stores transaction history
├── main.go # Backend entry point
├── go.mod # Go module
├── go.sum
└── README.md


---

## ⚙️ Install Go
If you don’t have Go installed yet:

1. Download from [https://go.dev/dl](https://go.dev/dl)  
2. Install according to your OS instructions.  
3. Verify installation:
   ```bash
   go version


Example output:

go version go1.22.3 linux/amd64

📦 Install Dependencies

Clone the project and fetch required dependencies:

git clone https://github.com/<your-username>/anoma-intent-redux.git
cd anoma-intent-redux
go mod tidy

🚀 Run Backend

Start the backend mock API server:

go run main.go


By default, the server runs on:
👉 http://localhost:8080

Available endpoints:

POST /intent → Parse and process user intent.

GET /history → Return stored transaction history.

💻 Run Frontend

The frontend files are in the static/ folder.

Option 1: Open the file directly:

static/index.html


Option 2 (recommended): Access it via the backend server:

http://localhost:8080

🎮 Demo Flow

Enter an intent, for example:

Buy 1 ETH on Ethereum


Or use quick chain buttons (Ethereum, BSC, Polygon, Arbitrum, Optimism).

Click Send → Backend processes → Frontend displays results.

Transactions are saved to History, filterable by chain.

Click Tx Hash to view details in explorer.html.

📌 Why It Fits the Theme

Intent-Centric: The entire architecture is built around parsing and handling user intents.

Multi-Chain: Supports multiple blockchain networks in the demo.

Enhanced UX: Beautiful UI with animations, toast notifications, and explorer integration.

Easy to Deploy: Simple Go backend + HTML/CSS/JS frontend.

📜 License

MIT License – Free to use for learning and demo purposes.


---

