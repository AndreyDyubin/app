package models

//go:generate reform

//reform:datafiles
type DataFiles struct {
	ID   int64  `reform:"id,pk"`
	Name string `reform:"name"`
}
