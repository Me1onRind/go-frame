package routers

import (
	"go-frame/internal/controller/http/audio"
	"go-frame/internal/core/gateway"
	"go-frame/protocol/audio_proto"

	"github.com/gin-gonic/gin"
)

func registerAudioApi(router *gin.RouterGroup) {
	audioController := audio.NewAudioController()

	router = router.Group("/audio")
	router.POST("upload_file", gateway.Json(audioController.UploadAudio, &audio_proto.CreateAudioReq{}))
	router.GET("list", gateway.Json(audioController.List, &audio_proto.ListAudioReq{}))
}
