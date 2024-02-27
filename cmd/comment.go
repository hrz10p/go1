package main

import (
	"main/pkg/models"
	"main/pkg/services"
	"main/pkg/utils/validators"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	Service *services.Service
}

func NewCommentHandler(Service *services.Service) *CommentHandler {
	return &CommentHandler{
		Service: Service,
	}
}

func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUserFromContext(r)
	if (user == models.User{}) {
		http.Error(w, "can't find a user :(", http.StatusBadRequest)
		return
	}

	postID := r.FormValue("postID")

	commentID := r.FormValue("commentID")
	commentIDInt, err := strconv.Atoi(commentID)
	if err != nil {
		http.Error(w, "commentID is not correct", http.StatusBadRequest)
		return
	}

	err = h.Service.CommentService.DeleteComment(commentIDInt)
	if err != nil {
		http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/post/"+postID, http.StatusSeeOther)
}

func (h *CommentHandler) SubmitComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUserFromContext(r)
	if (user == models.User{}) {
		http.Error(w, "cant find a user :(", http.StatusBadRequest)
		return
	}
	postID := r.FormValue("postID")
	postint, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, "postID not correct", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	cnt := 0
	for i := 0; i < len(content); i++ {
		if (content[i] >= '0' && content[i] <= '9') || (content[i] >= 'A' && content[i] <= 'Z') || (content[i] >= 'a' && content[i] <= 'z') {
			cnt++
		}
	}
	if cnt == 0 {
		content = ""
	}
	if validators.NonBlankValidate(content) != nil || validators.LengthRangeValidate(content, 1, 256) != nil {
		http.Error(w, "Comment content is too long or too short", http.StatusBadRequest)
		return
	}

	_, err = h.Service.PostService.GetPostByID(postint)
	if err != nil {
		http.Error(w, "Post not found", http.StatusBadRequest)
		return
	}

	err = h.Service.CommentService.SubmitCommentForPost(models.Comment{UID: user.ID, PostID: postint, Content: content})

	if err != nil {
		http.Error(w, "Comment creation error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/post/"+postID, http.StatusSeeOther)
	return
}
