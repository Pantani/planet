package cli_test

import (
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/suite"

	"planet/testutil/network"
	"planet/x/blog/types"
)

type IntegrationTestSuite struct {
	suite.Suite
}

func (s *IntegrationTestSuite) network(state proto.Message) *network.Network {
	s.T().Helper()
	cfg := network.DefaultConfig()
	if state != nil {
		buf, err := cfg.Codec.MarshalJSON(state)
		s.Require().NoError(err)
		cfg.GenesisState[types.ModuleName] = buf
	}
	return network.New(s.T(), cfg)
}

func (s *IntegrationTestSuite) SetupTest() {
	s.T().Log("setting up test")
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")
}

func (s *IntegrationTestSuite) TearDownTest() {
	s.T().Log("tearing down test")
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
