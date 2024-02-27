package views

type CommentView struct {
	ID        int
	Author    string
	Content   string
	CanDelete bool
}
