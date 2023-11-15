package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (uc *Usecase) GetSolutionByID(id int) (*model.SolutionByID, error) {
	sol, err := uc.store.GetSolutionByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sol, nil
}

func (uc *Usecase) GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	sols, err := uc.store.GetSolutionsByClassID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sols, nil
}

func (uc *Usecase) GetSolutionsByHomeworkID(homeworkID int) (*model.SolutionListForHw, error) {
	if err := uc.store.CheckHomeworkExistence(homeworkID); err != nil {
		return nil, e.StacktraceError(err)
	}

	sols, err := uc.store.GetSolutionsByHomeworkID(homeworkID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sols, nil
}
