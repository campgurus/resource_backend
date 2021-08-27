package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"resource-api/Config"
)

//GetAllClients Fetch all client data
func GetAllClients(client *[]Client) (err error) {
	if err = Config.DB.Find(client).Error; err != nil {
		return err
	}
	return nil
}

//CreateClient ... Insert New data
func CreateClient(client *Client) (err error) {
	if err = Config.DB.Create(client).Error; err != nil {
		return err
	}
	return nil
}

//GetClientByID ... Fetch only one client by Id
func GetClientByID(client *Client, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(client).Error; err != nil {
		return err
	}
	return nil
}

//UpdateClient ... Update client
func UpdateClient(client *Client, id string) (err error) {
	fmt.Println(client)
	Config.DB.Save(client)
	return nil
}

//DeleteClient ... Delete client
func DeleteClient(client *Client, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(client)
	return nil
}