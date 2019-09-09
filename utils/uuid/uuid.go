package uuid

import (
	"github.com/google/uuid"
	"github.com/sillyhatxu/retry-utils"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func GeneratorUUID() string {
	var resultId string
	err := retry.Do(func() error {
		id, err := uuid.NewUUID()
		if err != nil {
			return err
		}
		resultId = id.String()
		return nil
	}, retry.Attempts(10), retry.ErrorCallback(func(n uint, err error) {
		logrus.Errorf("retry [%d] generator id error : %v", n, err)
	}))
	if err != nil {
		logrus.Errorf("Generator UUID error. %v", err)
		return strconv.Itoa(time.Now().Nanosecond())
	}
	return strings.ToUpper(strings.ReplaceAll(resultId, "-", ""))
}
