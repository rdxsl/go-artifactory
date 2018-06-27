package artifactory

type ErrorsJson struct {
	Errors []ErrorJson `json:"errors,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type ErrorJson struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
