package models

//go:generate reform

//reform:datafiles
type DataFiles struct {
	UUID string `reform:"uuid,pk"`
	Name string `reform:"name"`
}
