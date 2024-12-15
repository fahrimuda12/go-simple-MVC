package dto

type WalletDTO struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Balance         float64 `json:"balance"`
	Key_Phrase      string  `json:"key_phrase"`
	User_ID         uint    `json:"user_id"`
	Virtual_Account string  `json:"virtual_account"`
	Tag_Name        string  `json:"tag_name"`
}