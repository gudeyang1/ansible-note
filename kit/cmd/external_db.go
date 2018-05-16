package cmd



type External_DB_Info struct {
	Product_name string `json:"product_name"`
	Internal	bool	`json:"internal"`
	DB_engine	string	`json:"db_engine"`
	DB_server	string	`json:"db_server"`
	DB_port		string	`json:"db_port"`
	DB_inst		string	`json:"db_inst"`
	DB_login	string	`json:"db_login"`
	DB_password	string	`json:"db_password"`

}

type External_DB_List struct {
	Database	[]External_DB_Info	`json:"database"`
}



func External_DB_Main() {
	// main entrance of external DB

	var s External_DB_Info

		JsonContent :=JsonFileCheck(ExternalDBJsonPath)
		JsonFormatCheck(JsonContent,&s,ExternalDBJsonPath)


}