package services

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// ApiResponse mapea la estructura del JSON con `items` y `next_page`
type ApiResponse struct {
    Items    []Item `json:"items"`
    NextPage string `json:"next_page"`
}

// Item representa cada objeto dentro de `items`
type Item struct {
    Ticker        string `json:"ticker"`
    ObjetivoDesde string `json:"target_from"`
    ObjetivoA     string `json:"target_to"`
    Empresa       string `json:"company"`
    Accion        string `json:"action"`
    Corretaje     string `json:"brokerage"`
    RatingFrom    string `json:"rating_from"`
    RatingTo      string `json:"rating_to"`
    Hora          string `json:"time"`
}

// FetchDataWithPagination obtiene múltiples páginas de la API
func FetchDataWithPagination(baseURL string, bearerToken string, maxPages int) ([]Item, error) {
    var allItems []Item
    nextPage := ""

    for i := 0; i < maxPages; i++ {
        // Construir URL con `next_page` si existe
        requestURL := baseURL
        if nextPage != "" {
            requestURL = fmt.Sprintf("%s?next_page=%s", baseURL, nextPage)
        }

        fmt.Println("Fetching:", requestURL) // Debug

        // Hacer la petición HTTP
        req, err := http.NewRequest("GET", requestURL, nil)
        if err != nil {
            return nil, err
        }

        // Agregar Bearer Token si es necesario
        if bearerToken != "" {
            req.Header.Set("Authorization", "Bearer "+bearerToken)
        }

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            return nil, fmt.Errorf("error HTTP %d al obtener datos", resp.StatusCode)
        }

        // Leer respuesta
        bodyBytes, err := io.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }

        var apiResp ApiResponse
        if err := json.Unmarshal(bodyBytes, &apiResp); err != nil {
            return nil, err
        }

        // Agregar items al resultado
        allItems = append(allItems, apiResp.Items...)

        // Verificar si hay `next_page`
        if apiResp.NextPage == "" {
            break // No hay más páginas, salimos del bucle
        }
        nextPage = apiResp.NextPage
    }

    return allItems, nil
}
