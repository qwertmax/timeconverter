// Copyright © 2014, 2015 Maxim Tishchenko
// All Rights Reserved.

package model

type Error struct {
	Type    string `json:"type"`    // The type of error returned. Can be … TODO.
	Message string `json:"message"` // optional	A human-readable message giving more details about the error. If something needs to be shown to end-users, this should be it.
	Param   string `json:"param"`   // The parameter the error relates to if the error is parameter-specific.
}
