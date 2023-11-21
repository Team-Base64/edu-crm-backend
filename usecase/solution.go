package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
	"strconv"
)

func (uc *Usecase) GetSolutionByID(id int) (*model.SolutionByID, error) {
	sol, err := uc.store.GetSolutionByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sol, nil
}

func (uc *Usecase) GetSolutionsByClassID(classID int) ([]model.SolutionFromClass, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	sols, err := uc.store.GetSolutionsByClassID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sols, nil
}

func (uc *Usecase) GetSolutionsByHomeworkID(homeworkID int) ([]model.SolutionForHw, error) {
	if err := uc.store.CheckHomeworkExistence(homeworkID); err != nil {
		return nil, e.StacktraceError(err)
	}

	sols, err := uc.store.GetSolutionsByHomeworkID(homeworkID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sols, nil
}

func (uc *Usecase) EvaluateSolutionbyID(solutionID int, evaluation *model.SolutionEvaluation) error {
	msg, err := uc.genEvaluationMsg(solutionID, evaluation)
	if err != nil {
		return e.StacktraceError(err)
	}

	if err := uc.store.AddEvaluationForSolution(solutionID, evaluation.IsApproved, msg); err != nil {
		return e.StacktraceError(err)
	}

	chatID, err := uc.store.GetChatIDBySolutionID(solutionID)
	if err != nil {
		return e.StacktraceError(err) // TODO возрат состояния или оповещение о проблеме с доставкой
	}
	if err := uc.chatService.SendMsg(&model.SingleMessage{
		ChatID:   chatID,
		Text:     msg,
		Attaches: []string{},
	}); err != nil {
		return e.StacktraceError(err) // TODO возрат состояния или оповещение о проблеме с доставкой
	}

	return nil
}

func (uc *Usecase) genEvaluationMsg(solutionID int, evaluation *model.SolutionEvaluation) (string, error) {
	info, err := uc.store.GetInfoForEvaluationMsgBySolutionID(solutionID)
	if err != nil {
		return "", e.StacktraceError(err)
	}

	var msg string
	msg += "Преподаватель проверил ваше решение от " +
		info.SolutionCreateTime.Format("15:4 02.01.2006") +
		" для домашнего задания: " + info.HomeworkTitle + "\n"

	for id, taskEval := range evaluation.Tasks {
		msg += "Задание №" + strconv.Itoa(id+1) + ":\n" +
			taskEval.Evaluation + "\n"
	}

	msg += "Результат: "
	if evaluation.IsApproved {
		msg += "Зачтено!"
	} else {
		msg += "Не зачетно!"
	}

	return msg, nil
}
