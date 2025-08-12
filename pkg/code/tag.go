package code

import (
	"git-auto-commit/infra/constants"
	"git-auto-commit/pkg/commit"
	"path/filepath"
)

func (c *Code) WithTag(files []string, formatted string, added, modified, deleted []string) string {
	// 1) ищем file-type теги по всем файлам (docs/style/test и т.п.)
	fileTagCount := map[string]int{}
	reprFileForTag := map[string]string{} // для выбранного тега — пример файла

	for _, f := range files {
		t := commit.DetectTagByFile(&f, "") // передаём пустой статус, чтобы DetectTagByFile смотрел только на имя/расширение
		if t == "" {
			continue
		}
		fileTagCount[t]++
		if _, ok := reprFileForTag[t]; !ok {
			reprFileForTag[t] = f
		}
	}

	if len(fileTagCount) > 0 {
		// выбираем самый "частый" file-type тег
		chosenTag := ""
		maxCnt := 0
		for tg, cnt := range fileTagCount {
			if cnt > maxCnt {
				maxCnt = cnt
				chosenTag = tg
			}
		}
		// используем representative file для сообщения
		fileForMsg := reprFileForTag[chosenTag]
		return commit.CreateAutoCommitMsg(&fileForMsg, &formatted, chosenTag)
	}

	// 2) если file-type тега нет — смотрим на статусы A/M/D/R
	statusCount := map[string]int{
		constants.NameStatus_Added:    len(added),
		constants.NameStatus_Modified: len(modified),
		constants.NameStatus_Deleted:  len(deleted),
	}

	maxStatus := ""
	maxCount := 0
	for st, cnt := range statusCount {
		if cnt > maxCount {
			maxCount = cnt
			maxStatus = st
		}
	}

	// маппим статус в тип коммита
	var tag string
	switch maxStatus {
	case constants.NameStatus_Added:
		tag = constants.Type_CommitFeat
	case constants.NameStatus_Modified:
		tag = constants.Type_CommitRefactor
	case constants.NameStatus_Deleted:
		tag = constants.Type_CommitFix
	case constants.NameStatus_Renamed:
		tag = constants.Type_CommitRefactor
	default:
		tag = constants.Type_CommitRefactor
	}

	// выбираем файл для CreateAutoCommitMsg по частоте расширений (как раньше)
	extCount := map[string]int{}
	maxFile := ""
	maxExtCount := 0
	for _, f := range files {
		ext := filepath.Ext(f)
		extCount[ext]++
		if extCount[ext] > maxExtCount {
			maxExtCount = extCount[ext]
			maxFile = f
		}
	}

	return commit.CreateAutoCommitMsg(&maxFile, &formatted, tag)
}
