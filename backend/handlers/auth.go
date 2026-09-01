package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const capitalosSandboxURL = "https://api-sandbox.capitalos.com"

func GetCapitalosToken(c *gin.Context) {
	accountID := os.Getenv("CAPITALOS_ACCOUNT_ID")
	apiToken := os.Getenv("CAPITALOS_API_TOKEN")

	body, _ := json.Marshal(map[string]string{"clientType": "web"})

	req, err := http.NewRequest(
		http.MethodPost,
		capitalosSandboxURL+"/accounts/"+accountID+"/initiate-login",
		bytes.NewBuffer(body),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to build request"})
		return
	}

	req.Header.Set("Authorization", apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "capitalos unreachable"})
		return
	}
	defer resp.Body.Close()

	var result map[string]any
	respBody, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(respBody, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid response from capitalos"})
		return
	}

	// Extrai o token da resposta da Capitalos
	// ajuste o campo conforme o retorno real da API deles
	token, ok := result["token"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token not found in response", "raw": result})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
