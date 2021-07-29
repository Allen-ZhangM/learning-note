package mocks

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
)

func generateRegexpChecker(re string) func([]byte) error {
	return func(val []byte) error {
		matched, err := regexp.MatchString(re, string(val))
		if err != nil {
			return errors.New("Error while trying to match the input message with the expected pattern: " + err.Error())
		}
		if !matched {
			return fmt.Errorf("No match between input value \"%s\" and expected pattern \"%s\"", val, re)
		}
		return nil
	}
}

type testReporterMock struct {
	errors []string
}

func newTestReporterMock() *testReporterMock {
	return &testReporterMock{errors: make([]string, 0)}
}

func (trm *testReporterMock) Errorf(format string, args ...interface{}) {
	trm.errors = append(trm.errors, fmt.Sprintf(format, args...))
}

func TestMockAsyncProducerImplementsAsyncProducerInterface(t *testing.T) {
	var mp interface{} = &mocks.AsyncProducer{}
	if _, ok := mp.(sarama.AsyncProducer); !ok {
		t.Error("The mock producer should implement the sarama.Producer interface.")
	}
}

func TestProducerReturnsExpectationsToChannels(t *testing.T) {
	config := mocks.NewTestConfig()
	config.Producer.Return.Successes = true
	mp := mocks.NewAsyncProducer(t, config)

	mp.ExpectInputAndSucceed()
	mp.ExpectInputAndSucceed()
	mp.ExpectInputAndFail(sarama.ErrOutOfBrokers)

	mp.Input() <- &sarama.ProducerMessage{Topic: "test 1"}
	mp.Input() <- &sarama.ProducerMessage{Topic: "test 2"}
	mp.Input() <- &sarama.ProducerMessage{Topic: "test 3"}

	msg1 := <-mp.Successes()
	msg2 := <-mp.Successes()
	err1 := <-mp.Errors()

	if msg1.Topic != "test 1" {
		t.Error("Expected message 1 to be returned first")
	}

	if msg2.Topic != "test 2" {
		t.Error("Expected message 2 to be returned second")
	}

	if err1.Msg.Topic != "test 3" || err1.Err != sarama.ErrOutOfBrokers {
		t.Error("Expected message 3 to be returned as error")
	}

	if err := mp.Close(); err != nil {
		t.Error(err)
	}
}

func TestProducerWithTooFewExpectations(t *testing.T) {
	trm := newTestReporterMock()
	mp := mocks.NewAsyncProducer(trm, nil)
	mp.ExpectInputAndSucceed()

	mp.Input() <- &sarama.ProducerMessage{Topic: "test"}
	mp.Input() <- &sarama.ProducerMessage{Topic: "test"}

	if err := mp.Close(); err != nil {
		t.Error(err)
	}

	if len(trm.errors) != 1 {
		t.Error("Expected to report an error")
	}
}

func TestProducerWithTooManyExpectations(t *testing.T) {
	trm := newTestReporterMock()
	mp := mocks.NewAsyncProducer(trm, nil)
	mp.ExpectInputAndSucceed()
	mp.ExpectInputAndFail(sarama.ErrOutOfBrokers)

	mp.Input() <- &sarama.ProducerMessage{Topic: "test"}
	if err := mp.Close(); err != nil {
		t.Error(err)
	}

	if len(trm.errors) != 1 {
		t.Error("Expected to report an error")
	}
}

func TestProducerWithCheckerFunction(t *testing.T) {
	trm := newTestReporterMock()
	mp := mocks.NewAsyncProducer(trm, nil)
	mp.ExpectInputWithCheckerFunctionAndSucceed(generateRegexpChecker("^tes"))
	mp.ExpectInputWithCheckerFunctionAndSucceed(generateRegexpChecker("^tes$"))

	mp.Input() <- &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("test")}
	mp.Input() <- &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("test")}
	if err := mp.Close(); err != nil {
		t.Error(err)
	}

	if len(mp.Errors()) != 1 {
		t.Error("Expected to report an error")
	}

	err1 := <-mp.Errors()
	if !strings.HasPrefix(err1.Err.Error(), "No match") {
		t.Error("Expected to report a value check error, found: ", err1.Err)
	}
}

func TestProducerWithBrokenPartitioner(t *testing.T) {
	trm := newTestReporterMock()
	config := sarama.NewConfig()
	config.Producer.Partitioner = func(string) sarama.Partitioner {
		return brokePartitioner{}
	}
	mp := mocks.NewAsyncProducer(trm, config)
	mp.ExpectInputWithMessageCheckerFunctionAndSucceed(func(msg *sarama.ProducerMessage) error {
		if msg.Partition != 15 {
			t.Error("Expected partition 15, found: ", msg.Partition)
		}
		if msg.Topic != "test" {
			t.Errorf(`Expected topic "test", found: %q`, msg.Topic)
		}
		return nil
	})
	mp.ExpectInputAndSucceed() // should actually fail in partitioning

	mp.Input() <- &sarama.ProducerMessage{Topic: "test"}
	mp.Input() <- &sarama.ProducerMessage{Topic: "not-test"}
	if err := mp.Close(); err != nil {
		t.Error(err)
	}

	if len(trm.errors) != 1 || !strings.Contains(trm.errors[0], "partitioning unavailable") {
		t.Error("Expected to report partitioning unavailable, found", trm.errors)
	}
}

// brokeProducer refuses to partition anything not on the “test” topic, and sends everything on
// that topic to partition 15.
type brokePartitioner struct{}

func (brokePartitioner) Partition(msg *sarama.ProducerMessage, n int32) (int32, error) {
	if msg.Topic == "test" {
		return 15, nil
	}
	return 0, errors.New("partitioning unavailable")
}

func (brokePartitioner) RequiresConsistency() bool { return false }
