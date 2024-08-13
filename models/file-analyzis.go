package models

type FileAnalysis struct {
	Lines      int `json:"lines"`
	Code       int `json:"code"`
	Comment    int `json:"comment"`
	Blank      int `json:"blank"`
	Complexity int `json:"complexity"`
}
