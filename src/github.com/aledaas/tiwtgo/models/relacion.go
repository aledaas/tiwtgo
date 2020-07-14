package models

/*Relacion Model*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId,omitempty"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId,omitempty"`
}
