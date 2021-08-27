package Models

type Client struct {
	Id           uint      `json:"id"`
	Name        string    `json:"name"`
	URL       string    `json:"url"`
	Description  string    `json:"description"`
}
func (c *Client) TableName() string {
	return "client"
}