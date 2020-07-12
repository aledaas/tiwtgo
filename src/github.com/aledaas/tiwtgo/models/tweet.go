package models

/*Tweet Model*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
