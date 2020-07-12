package models

/*RespuestaLogin manda la respuesta del login */
type RespuestaLogin struct {
	Token string `json:"tocken,omitempty"`
}
