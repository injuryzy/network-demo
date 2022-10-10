package cerr

import "github.com/sirupsen/logrus"

func Check(err error) {
	if err != nil {
		logrus.Error(err)
	}
}
