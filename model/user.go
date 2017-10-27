package model

type User struct {
	ID       int     `json:"id"`
	Nickname string  `json:"nickname"`
	WSConn   *WSConn `json:"-"`
}
