// // Lấy hash từ URL (?tx=hash123&chain=Ethereum)
// const params = new URLSearchParams(window.location.search);
// const txHash = params.get("tx") || "N/A";
// const chain = params.get("chain") || "Unknown";

// document.getElementById("txHash").innerText = txHash;
// document.getElementById("chain").innerText = chain;

async function loadTx() {
  const params = new URLSearchParams(window.location.search);
  const txHash = params.get("tx");
  if (!txHash) {
    document.getElementById("txHash").innerText = "No Tx Hash";
    return;
  }

  const res = await fetch("/history");
  const history = await res.json();
  const tx = history.find(h => h.TxHash === txHash);

  if (!tx) {
    document.getElementById("txHash").innerText = "No transaction found";
    return;
  }

  document.getElementById("txHash").innerText = tx.TxHash;
  document.getElementById("action").innerText = `${tx.Action}`;
  document.getElementById("value").innerText = `${tx.Amount} ${tx.Token}`;
  document.getElementById("gasFee").innerText = tx.GasFee;
  document.getElementById("chain").innerText = tx.Chain;
}

window.onload = loadTx;
