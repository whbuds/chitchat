package handlers

import (
	"chitchat/models"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

// POST /thread/post
// 在指定群组下创建新主题
func PostThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "thread_not_found",
			})
			errorMessage(w, r, msg)
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(w, r, url, 302)
	}
}
