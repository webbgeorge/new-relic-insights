package cli_logger

import "github.com/sirupsen/logrus"

type Formatter struct {}

func (f Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}
