package backend

import (
	"strconv"

	e "main/domain/errors"
	m "main/domain/model"
	u "main/domain/utils"
)

func (uc *BackendUsecase) GetSolutionByID(id int) (*m.SolutionByID, error) {
	sol, err := uc.dataStore.GetSolutionByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sol, nil
}

func (uc *BackendUsecase) GetSolutionsByClassID(classID int) ([]m.SolutionFromClass, error) {
	if err := uc.dataStore.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	sols, err := uc.dataStore.GetSolutionsByClassID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sols, nil
}

func (uc *BackendUsecase) GetSolutionsByHomeworkID(homeworkID int) ([]m.SolutionForHw, error) {
	if err := uc.dataStore.CheckHomeworkExistence(homeworkID); err != nil {
		return nil, e.StacktraceError(err)
	}

	sols, err := uc.dataStore.GetSolutionsByHomeworkID(homeworkID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sols, nil
}

func (uc *BackendUsecase) EvaluateSolutionbyID(solutionID int, evaluation *m.SolutionEvaluation) error {
	msg, err := uc.genEvaluationMsg(solutionID, evaluation)
	if err != nil {
		return e.StacktraceError(err)
	}

	if err := uc.dataStore.AddEvaluationForSolution(solutionID, evaluation.IsApproved, msg); err != nil {
		return e.StacktraceError(err)
	}

	chatID, err := uc.dataStore.GetChatIDBySolutionID(solutionID)
	if err != nil {
		return e.StacktraceError(err) // TODO возрат состояния или оповещение о проблеме с доставкой
	}
	if err := uc.chat.SendNotification(&m.SingleMessage{
		ChatID:   chatID,
		Text:     msg,
		Attaches: []string{},
	}); err != nil {
		return e.StacktraceError(err) // TODO возрат состояния или оповещение о проблеме с доставкой
	}

	return nil
}

func (uc *BackendUsecase) genEvaluationMsg(solutionID int, evaluation *m.SolutionEvaluation) (string, error) {
	info, err := uc.dataStore.GetInfoForEvaluationMsgBySolutionID(solutionID)
	if err != nil {
		return "", e.StacktraceError(err)
	}

	var msg string
	msg += "Преподаватель проверил ваше решение от " +
		u.TimeToString(info.SolutionCreateTime) +
		" для домашнего задания: " + info.HomeworkTitle + "\n"

	for id, taskEval := range evaluation.Tasks {
		msg += "Задание №" + strconv.Itoa(id+1) + ":\n" +
			taskEval.Evaluation + "\n"
	}

	msg += "Результат: "
	if evaluation.IsApproved {
		msg += "Зачтено!"
	} else {
		msg += "Не зачтено!"
	}

	return msg, nil
}
