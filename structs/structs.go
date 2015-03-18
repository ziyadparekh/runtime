package structs

type Event struct {
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Url       string `json:"url"`
	State     string `json:"state"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
	Head      Commit `json:"head"`
}

type User struct {
	Login  string `json:"login"`
	Avatar string `json:"avatar_url"`
	Url    string `json:"url"`
}

type Commit struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
}

type Ref struct {
	Ref        string     `json:"ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Created    bool       `json:"created"`
	Deleted    bool       `json:"deleted"`
	HeadCommit HeadCommit `json:"head_commit"`
	User       User       `json:"sender"`
}

type HeadCommit struct {
	Id        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp string    `json:"timestamp"`
	Url       string    `json:"url"`
	Author    Committer `json:"author"`
	Committer Committer `json:"committer"`
}

type Committer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Branch struct {
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	MasterBranch string `json:"master_branch,omitempty"`
	Event        string
}
