package api


func (app *App) DoTts(text string, param Param) (audioBody []byte, err error) {
	paramMap := map[string]string{
		"aud": param.Aue,
		"auf": param.Auf,
		"voice_name": param.VoiceName,
	}
	return app.postReq(app.config.Api.Tts, text, paramMap)
}
