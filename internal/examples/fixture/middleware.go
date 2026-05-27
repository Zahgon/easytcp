package fixture

import (
	"github.com/DarthPestilane/easytcp"
	"github.com/sirupsen/logrus"
)

func RecoverMiddleware(log *logrus.Logger) easytcp.MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(easytcp.MiddlewareFunc)
}
