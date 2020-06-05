package models

/*Relacion modelo para grabar la relación de un usuario con otro*/
type Relacion struct {
	UsuarioID        string `bson:"usuarioid" json:"usuarioId"`
	UsurioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
