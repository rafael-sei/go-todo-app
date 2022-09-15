package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		//[app/views/templates/layout.html app/views/templates/top.html]
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	//テンプレートのキャッシュを作成
	//template.Must()は独自にエラーチェックを行うため、errorを返り値には持たず、ハンドリングす必要がありません。
	//つまりParseFilesがエラーの場合、panicになる。
	//template.ParseFiles()は可変長引数をとり、その引数としてキャッシュさせたいファイルの名前を指定します。
	templates := template.Must(template.ParseFiles(files...))
	//difineでテンプレートを定義した場合、ExecuteTemplateでlayoutを明示的に指定する必要がある
	templates.ExecuteTemplate(w, "layout", data)
}

// アクセス制限
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	//クッキから値を受け取る
	cookie, err := r.Cookie("_cookie")
	// error無ければ
	if err == nil {
		//cookieで取得したValueを入れる
		sess = models.Session{UUID: cookie.Value}
		//sessionsにUUIDがあるかどうか、無ければエラーを返す
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StartMainServer() error {

	//静的ファイルを読み込み
	files := http.FileServer(http.Dir(config.Config.Static))

	// static フォルダーがない時 SripPrefixで取り除く(パスを書くときはstaticは必要)
	// http.Handle("/static/",http.StripPrefix("/static/",files))
	//static フォルダーがある時
	http.Handle("/static/", files)

	//第１引数URLの登録
	//第２引数はハンドラー
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	// 第１引数 = Port指定
	// 第２引数 nil = 存在しないURLがあれば(404 not found)を返す
	return http.ListenAndServe("127.0.0.1:"+config.Config.Port, nil)
}
