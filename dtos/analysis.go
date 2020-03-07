package dtos

import "time"

type AnalysisRequest struct {
	UserID    int64
	BeginTime *time.Time
	EndTime   *time.Time
}

type AnalysisTagCommon struct {
	Tag   string `json:"tag"`
	Money int64  `json:"money"`
}

type AnalysisTagResponse struct {
	Data []AnalysisTagCommon
}

type AnalysisDayCommon struct {
	Day   string `json:"day"`
	Money int64  `json:"money"`
}

type AnalysisDayResponse struct {
	Data []AnalysisDayCommon
}
