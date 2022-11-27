# Simple MetaMask Auth Flow

- 安裝 live server 的 vscode plugin
- 在 client/index.html 點右鍵開啟於 5500.
- 開啟 5500 按下 login。可以在 devtool 看到 signData, Account。
- 複製 signData 和 Account。

- 在 main.go 的 13, 14 行 po 上 account 和 SignData。
- cli 執行 go main.go
- 可以看到解碼後的 address 跟真實的只差了大小寫。