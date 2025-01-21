package controllers

import "net/http"

func (db *Data) TakeId(secion string) (int, string, string) {
	query := `
	SELECT 
    	session.user_id, 
    	user.username, 
    	user.uid 
	FROM 
	    session 
	INNER JOIN 
    	user 
	ON 
    	session.user_id = user.id 
	WHERE 
    	session.token = ?;

	`
	id := 0
	name := ""
	uid := ""
	err := db.Db.QueryRow(query, secion).Scan(&id, &name, &uid)
	if err != nil {
		return -1, "", ""
	}
	return id, name, uid
}

func (db *Data) GetIdFromReq(r *http.Request) (int, string, string) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return -1, "", ""
	}
	userId, name, uid := db.TakeId(cookie.Value)
	return userId, name, uid
}
