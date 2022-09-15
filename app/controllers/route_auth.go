package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)
//サインアップ
func signup(w http.ResponseWriter, r *http.Request) {
	//GETの処理
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
		//POSTの処理
	} else if r.Method == "POST" {
		//入力処理の取得
		err := r.ParseForm() //form解析する
		if err != nil {
			log.Println(err)
		}

		//user struct生成
		user := models.User{
			//r.PostFormValueで属性を取得
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		// 成功したらトップにリダイレクトする
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
//ログイン
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
//ユーザの認証
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	//パスワード比較
	//userパスワードは暗号化されているので、formで送られて来たパスワードを暗号化して比較する。trueならsessionを作成する
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		//sessionを作成
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		//クッキを作成
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		//ブラウザのログイン情報を保存
		http.SetCookie(w, &cookie)
		// 成功したらトップにリダイレクトする
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		//passwordが一致しなければ
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
//ログアウト
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		//クッキのUUIDをセッションのUUIDに入れて削除
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
