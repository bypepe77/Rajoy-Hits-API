package sounds

import "github.com/gin-gonic/gin"

type SoundsRouteController struct {
	r                gin.RouterGroup
	soundsController SoundsInterface
}

func NewSoundRoutes(r gin.RouterGroup) *SoundsRouteController {
	return &SoundsRouteController{
		r:                r,
		soundsController: NewSoundController(),
	}
}

func (s *SoundsRouteController) RegisterSoundsRoutes() {
	s.r.GET("/test", s.soundsController.GetSounds)
}
