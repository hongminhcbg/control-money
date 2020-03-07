package services

import "github.com/hongminhcbg/control-money/dtos"

type AnalysisService interface {
	ByTag(request dtos.AnalysisRequest) (*dtos.AnalysisTagResponse, error)
	ByDay(request dtos.AnalysisRequest) (*dtos.AnalysisDayResponse, error)
}

type analysisImpl struct {

}

func NewAnalysisService() AnalysisService {
	return &analysisImpl{}
}

func (service *analysisImpl) ByTag(request dtos.AnalysisRequest) (*dtos.AnalysisTagResponse, error)  {
	return nil, nil
}

func (service *analysisImpl) ByDay(request dtos.AnalysisRequest) (*dtos.AnalysisDayResponse, error)  {
	return nil, nil
}