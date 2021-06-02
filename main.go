package main

import (
	"fmt"
	"main/app/models"
)


func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)
	// user, _ := models.GetUser(2)
	// user.CreateTodo("First todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	user, _ := models.GetUser(3)
	user.CreateTodo("Third todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(1)
	t.Content = "Update"
	t.UpdateTodo()
}