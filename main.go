package main

import "todo_app/app/controllers"

// "todo_app/config"

func main() {
	controllers.StartMainServer()

	// iniテスト
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)
	*/

	// loggingテスト
	// log.Println("test")

	// tabel作成テスト
	// fmt.Println(models.Db)

	// CreateUserテスト
	// u := &models.User{}
	// u.Name = "test5"
	// u.Email = "test5@example.com"
	// u.Password = "test5"
	// fmt.Println(u)
	// u.CreateUser()

	// GetUserテスト
	/*
		u, _ := models.GetUser(2)
		fmt.Println(u)
	*/

	// UpdateUserテスト
	/*
		u.Name = "UpdateUser"
		u.Email = "Update@exampl.com"
		u.UpdateUser()
		u, _ = models.GetUser(2)
		fmt.Println(u)
	*/

	// DeleteUserテスト
	/*
		u.DeleteUser()
		u, _ = models.GetUser(2)
		fmt.Println(u)
	*/

	// CreateTodoテスト
	// user, _ := models.GetUser(4)
	// user.CreateTodo("Second Todo")

	// GetTodoテスト
	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	// GetToodsテスト
	// todos, _ := models.GetTodos()
	// fmt.Println(todos)
	// for _, t := range todos {
	// 	fmt.Println(t)
	// }

	// GetTodoByUserテスト
	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodoByUser()
	// fmt.Println(todos)
	// for _, t := range todos {
	// 	fmt.Println(t)
	// }

	// UpdateTodoテスト
	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	// DeleteTodoテスト
	// t, _ := models.GetTodo(1)
	// t.DeleteTodo()

	// GetUserByEmailテスト
	// user, _ := models.GetUserByEmail("test@exmaple.com")
	// fmt.Println(user)

	// CreateSessionテスト
	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// CheckSessionテスト
	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
