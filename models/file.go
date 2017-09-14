package models

//go:generate reform

//reform:datafiles
type DataFiles struct {
	id   int64  `reform:id`
	UUID string `reform:"uuid,pk"`
	Name string `reform:"name"`
}
