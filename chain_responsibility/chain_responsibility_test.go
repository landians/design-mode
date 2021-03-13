package chain_responsibility

import "testing"

func Test_ChainResponsibility(t *testing.T) {
	logger := newSimpleLogger()
	logger.Debug("debug")
	logger.Info("info")
	logger.Error("error")
}
