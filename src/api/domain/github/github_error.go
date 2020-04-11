package github

type GitHubErrorReponse struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Errors []GithubError `json:"errors"`
	DocumentationUrl string `json:"documentation_url"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code string `json:"code"`
	Field string `json:"field"`
	Message string `json:"message"`
}


