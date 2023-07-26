package db

type TableRows struct {
	Schemaname  string `json:"schemaname" xml:"schemaname"`
	Tablename   string `json:"tablename" xml:"tablename"`
	Tableowner  string `json:"tableowner" xml:"tableowner"`
	Tablespace  string `json:"tablespace" xml:"tablespace"`
	Hasindexes  string `json:"hasindexes" xml:"hasindexes"`
	Hasrules    string `json:"hasrules" xml:"hasrules"`
	Hastriggers string `json:"hastriggers" xml:"hastriggers"`
	Rowsecurity string `json:"rowsecurity" xml:"rowsecurity"`
}
