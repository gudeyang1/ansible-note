package cmd

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"

)

type External_Ladp_Body struct {
	Ldap_server_ip 				string  `json:"ldap_server_ip"`
	Ldap_server_port			string  `json:"ldap_server_port"`
	Ldap_use_ssl 				string  `json:"ldap_use_ssl"`
	Ldap_search_sub 			string  `json:"ldap_search_sub"`
	Ldap_base_dn 				string  `json:"ldap_base_dn"`
	Ldap_bind_user 				string  `json:"ldap_bind_user"`
	Ldap_bind_user_password		string  `json:"ldap_bind_user_password"`
	Ldap_user_base_dn 			string  `json:"ldap_user_base_dn"`
	Ldap_user_class 			string  `json:"ldap_user_class"`
	Ldap_user_filter 			string  `json:"ldap_user_filter"`
	Ldap_user_firstname 		string  `json:"ldap_user_firstname"`
	Ldap_user_lastname 			string  `json:"ldap_user_lastname"`
	Ldap_user_disp_name_attr 	string  `json:"ldap_user_disp_name_attr"`
	Ldap_user_name_attr 		string  `json:"ldap_user_name_attr"`
	Ldap_user_email				string	`json:"ldap_user_email"`
	Ldap_user_avatar			string	`json:"ldap_user_avatar"`
	Ldap_user_manager_id		string	`json:"ldap_user_manager_id"`
	Ldap_user_manager_id_value	string	`json:"ldap_user_manager_id_value"`
	Ldap_user_phone				string	`json:"ldap_user_phone"`
	Ldap_user_lastmodified		string	`json:"ldap_user_lastmodified"`
	Ldap_group_base_dn			string	`json:"ldap_group_base_dn"`
	Ldap_group_class			string	`json:"ldap_group_class"`
	Ldap_group_base_filter		string	`json:"ldap_group_base_filter"`
	Ldap_group_name_attr		string	`json:"ldap_group_name_attr"`
	Ldap_group_member_attr		string	`json:"ldap_group_member_attr"`
	Ldap_user_group				string	`json:"ldap_user_group"`
	Ldap_admin_group			string	`json:"ldap_admin_group"`


}


func JsonFileCheck(filename string) []byte  {
	buf ,err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Fprintf(os.Stderr,"Read file error:%s",err)
		os.Exit(1)
	}
	return buf

}

func JsonFormatCheck(jsonContent []byte,jsonFormat interface{},JsonFileName string)  {

	err := json.Unmarshal(jsonContent,&jsonFormat)
	if err != nil{
		fmt.Printf("%s Json format error:%s\n",JsonFileName,err)
		os.Exit(1)
	}
	fmt.Printf("===CheckingParameters: Check %s Json format Success.\n",JsonFileName)
	//fmt.Println(string(jsonContent))

}




func External_Ldap_Main() {
	// main entrance of external ldap and DB

	var s External_Ladp_Body

		JsonContent :=JsonFileCheck(ExternalLdapJsonPath)
		JsonFormatCheck(JsonContent,&s,ExternalLdapJsonPath)

}




