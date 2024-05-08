package suite

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type funSuite struct {
	suite.Suite
}

func (s *funSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("afterTest: suiteName=%s, testName=%s\n", suiteName, testName)
}

func (s *funSuite) BeforeTest(suiteName, testname string) {
	fmt.Printf("beforeTest: suiteName=%s, testName=%s\n", suiteName, testname)
}

func (s *funSuite) SetupSuite() {
	fmt.Printf("SetupSuite()...\n")
}

func (s *funSuite) TearDownSuite() {
	fmt.Printf("TearDownSuite()...\n")
}

func (s *funSuite) SetupTest() {
	fmt.Printf("SetupTest()...\n")
}

func (s *funSuite) TearDownTest() {
	fmt.Printf("TearDownTest()...\n")
}

func (s *funSuite) TestFoo() {
	foo()
}

func (s *funSuite) TestGoo() {
	goo()
}

func TestGooFoo(t *testing.T) {
	suite.Run(t, new(funSuite))
}

// super simple song
// 超级宝贝JOJO
// coco melon
//碰碰狐 pinkfong
//chuchu TV
