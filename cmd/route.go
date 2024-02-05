package main

import (
	"net/http"
)

// InitializeRoutes sets up the application routes
func (app *Application) InitializeRoutes() {
	auth := NewAuthHandler(app.Service)
	post := NewPostHandler(app.Service)
	middle := NewMiddle(app.Service)
	dep := NewDepHandler(app.Service)
	admin := NewAdminandler(app.Service)
	fs := http.FileServer(http.Dir("./ui/static/"))
	app.Router.Handle("/static/", http.StripPrefix("/static", fs))
	app.Router.Handle("/login", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(auth.Login))))))
	app.Router.Handle("/register", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(auth.Registration))))))
	app.Router.Handle("/", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(post.Index))))))
	app.Router.Handle("/post/", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(post.Post))))))
	app.Router.Handle("/createPost", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureTeacher(http.HandlerFunc(post.CreatePost))))))))
	app.Router.Handle("/filtered", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(http.HandlerFunc(post.Filtered)))))))
	app.Router.Handle("/logout", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(http.HandlerFunc(auth.Logout)))))))
	app.Router.Handle("/contacts", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(post.ContactsPage))))))

	app.Router.Handle("/admin", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureAdmin(http.HandlerFunc(admin.ViewUsers))))))))
	app.Router.Handle("/admin/promote", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureAdmin(http.HandlerFunc(admin.PromoteUser))))))))
	app.Router.Handle("/admin/demote", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureAdmin(http.HandlerFunc(admin.DemoteUser))))))))
	app.Router.HandleFunc("/deps", dep.GetAllDeps)
	app.Router.HandleFunc("/createDep", dep.CreateDep)
	app.Logger.Info("routs")
}
