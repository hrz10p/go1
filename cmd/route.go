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
	comm := NewCommentHandler(app.Service)
	admin := NewAdminandler(app.Service)
	fs := http.FileServer(http.Dir("./ui/static/"))
	app.Router.Handle("/static/", http.StripPrefix("/static", fs))
	app.Router.Handle("/login", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(auth.Login))))))
	app.Router.Handle("/register", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(auth.Registration))))))
	app.Router.Handle("/", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.secureBanned(http.HandlerFunc(post.Index)))))))
	app.Router.Handle("/post/", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(post.Post))))))
	app.Router.Handle("/createPost", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureTeacher(http.HandlerFunc(post.CreatePost))))))))
	app.Router.Handle("/filtered", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(http.HandlerFunc(post.Filtered)))))))
	app.Router.Handle("/logout", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(http.HandlerFunc(auth.Logout)))))))
	app.Router.Handle("/contacts", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(http.HandlerFunc(post.ContactsPage))))))

	app.Router.Handle("/admin", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureAdmin(http.HandlerFunc(admin.ViewUsers))))))))
	app.Router.Handle("/admin/promote", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureAdmin(http.HandlerFunc(admin.PromoteUser))))))))
	app.Router.Handle("/admin/demote", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureAdmin(http.HandlerFunc(admin.DemoteUser))))))))
	app.Router.Handle("/admin/ban", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureAdmin(http.HandlerFunc(admin.Ban))))))))
	app.Router.Handle("/admin/unban", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(middle.secureAdmin(http.HandlerFunc(admin.UNBan))))))))

	app.Router.Handle("/banned", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(http.HandlerFunc(post.BanPage)))))))

	app.Router.Handle("/submitComment", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(http.HandlerFunc(comm.SubmitComment)))))))
	app.Router.Handle("/deleteComment", middle.Authenticate(middle.LogRequest(middle.RecoverPanic(middle.SecureHeaders(middle.RequireAuthentication(http.HandlerFunc(comm.DeleteComment)))))))

	app.Router.HandleFunc("/deps", dep.GetAllDeps)
	app.Router.HandleFunc("/createDep", dep.CreateDep)
	app.Logger.Info("routs")
}
