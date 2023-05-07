package sounds

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bypepe77/Rajoy-Hits-API/rest"
	"github.com/gin-gonic/gin"
)

type SoundsInterface interface {
	GetSounds(c *gin.Context)
}

type sounds struct {
	httpClient rest.HTTPClient
}

func NewSoundController() SoundsInterface {
	return &sounds{
		httpClient: &http.Client{Timeout: time.Duration(30) * time.Second},
	}
}

func (s *sounds) GetSounds(c *gin.Context) {
	response, err := http.Get("https://tuna-api.voicemod.net/v1/sounds/search?search=rajoy&page=1&size=20&tagsSize=50&noDmca=false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al realizar la solicitud GET"})
		return
	}
	defer response.Body.Close()

	// Leer la respuesta JSON en una estructura de datos
	sounds := &SoundsResponse{}
	err = json.NewDecoder(response.Body).Decode(sounds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar la respuesta JSON"})
		return
	}

	// Devolver la respuesta JSON al cliente
	c.JSON(http.StatusOK, gin.H{"data": sounds})
}
