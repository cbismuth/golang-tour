package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite

	gen *rand.Rand
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

func (suite *MainTestSuite) SetupSuite() {
	fmt.Println("Setup test suite")

	suite.gen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (suite *MainTestSuite) TearDownSuite() {
	fmt.Println("Tear down test suite")
}

func (suite *MainTestSuite) SetupTest() {
	fmt.Println("Setup test")
}

func (suite *MainTestSuite) TearDownTest() {
	fmt.Println("Tear down test")
}

func (suite *MainTestSuite) TestGen() {
	const n = 1<<63 - 1

	x := suite.gen.Intn(n)
	y := suite.gen.Intn(n)

	suite.NotEqual(x, y)
}
