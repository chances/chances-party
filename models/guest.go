package models

import (
	"database/sql"
	"encoding/json"

	"github.com/vattle/sqlboiler/types"
)

type Guest struct {
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	CheckedIn bool   `json:"checked_in"`
}

func NewGuest(name string) Guest {
	return Guest{
		Name:      name,
		Alias:     "",
		CheckedIn: false,
	}
}

func (o *Party) Guests() ([]Guest, error) {
	guestList, err := o.GuestG().One()
	if err != nil {
		if err != sql.ErrNoRows {
			return make([]Guest, 0), err
		}
	}
	var guests []Guest
	err = guestList.Data.Unmarshal(&guests)
	if err != nil {
		return make([]Guest, 0), err
	}
	return guests, nil
}

func (o *Party) UpdateGuestList(guests []Guest) error {
	guestsJsonRaw, _ := json.Marshal(guests)
	guestsJson := types.JSON(guestsJsonRaw)
	guestList := GuestList{
		ID:   o.GuestsID,
		Data: guestsJson,
	}
	return guestList.UpdateG()
}
