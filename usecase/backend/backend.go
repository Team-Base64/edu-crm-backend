package backend

import (
	"math/rand"

	d "main/delivery"
	rep "main/repository"
	uc "main/usecase"
)

type BackendUsecase struct {
	dataStore rep.DataStoreInterface
	letters   []rune
	tokenLen  int
	bufToken  []rune
	chat      d.ChatInterface
	calendar  d.CalendarInterface
	fileStore rep.FileStoreInterface
	urlDomain string
}

func NewBackendUsecase(
	ds rep.DataStoreInterface,
	lettes string,
	tokenLen int,
	chat d.ChatInterface,
	calendar d.CalendarInterface,
	fs rep.FileStoreInterface,
	ud string,
) uc.UsecaseInterface {
	return &BackendUsecase{
		dataStore: ds,
		letters:   []rune(lettes),
		tokenLen:  tokenLen,
		bufToken:  make([]rune, tokenLen),
		chat:      chat,
		calendar:  calendar,
		fileStore: fs,
		urlDomain: ud,
	}
}

func (uc BackendUsecase) genRandomToken() string {
	for i := range uc.bufToken {
		uc.bufToken[i] = uc.letters[rand.Intn(len(uc.letters))]
	}
	return string(uc.bufToken)
}
