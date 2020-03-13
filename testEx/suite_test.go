package testex

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

// 此檔案所有測試開始之前
func (s *Suite) SetupSuite() {
	s.VariableThatShouldStartAtFive = 5
	log.Println("SetupSuite")
}

// 此檔案每個測試之前
func (s *Suite) SetupTest() {
	log.Println("SetupTest")
}

func (s *Suite) TestSuite1() {
	s.Equal(s.VariableThatShouldStartAtFive, 5)
	log.Println("TestSuite1")
}

func (s *Suite) TestSuite2() {
	assert.Equal(s.T(), 5, s.VariableThatShouldStartAtFive)
	log.Println("TestSuite2")
}

// 此檔案每個測試之後
func (s *Suite) TearDownTest() {
	log.Println("TearDownTest")
}

// 此檔案所有測試完成之後
func (s *Suite) TearDownSuite() {
	log.Println("TearDownSuite")
}
