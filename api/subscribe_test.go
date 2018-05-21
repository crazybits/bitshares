package api

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/denkhaus/bitshares/tests"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
	"github.com/stretchr/testify/suite"
)

const (
	SetSubscribeCallbackSubscriberID = 5
	SetSubscribeCallbackMsgs         = 8
	SetSubscribeCallbackDuration     = 30 * time.Second

	SubscribeToMarketSubscriberID = 4
	SubscribeToMarketMsgs         = 3
	SubscribeToMarketDuration     = 90 * time.Second
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type subscribeTest struct {
	suite.Suite
	TestAPI BitsharesAPI
}

func (suite *subscribeTest) SetupTest() {
	api := New(tests.WsFullApiUrl, tests.RpcApiUrl)
	if err := api.Connect(); err != nil {
		suite.FailNow(err.Error(), "Connect")
	}

	api.OnError(func(err error) {
		suite.FailNow(err.Error(), "OnError")
	})

	suite.TestAPI = api
}

func (suite *subscribeTest) TearDown() {
	if err := suite.TestAPI.Close(); err != nil {
		suite.FailNow(err.Error(), "Close")
	}
}

func (suite *subscribeTest) Test_SubscribeToMarket() {
	fmt.Printf("SubscribeToMarket: wait %s for %d incoming notifications\n",
		SubscribeToMarketDuration, SubscribeToMarketMsgs)

	if err := suite.TestAPI.SubscribeToMarket(SubscribeToMarketSubscriberID,
		tests.AssetBTS, tests.AssetCNY); err != nil {
		suite.FailNow(err.Error(), "SubscribeToMarket")
	}

	bar := pb.StartNew(SubscribeToMarketMsgs).Prefix("wait for data")
	err := suite.TestAPI.OnNotify(SubscribeToMarketSubscriberID, func(msg interface{}) error {
		bar.Increment()
		return nil
	})

	if err != nil {
		suite.FailNow(err.Error(), "OnNotify")
	}

	suite.Condition(func() bool {
		return util.WaitForCondition(SubscribeToMarketDuration, func() bool {
			return int(bar.Get()) >= SubscribeToMarketMsgs
		})
	}, "not enough incomming notifications during test window")

	bar.Finish()

	fmt.Println("Call UnsubscribeFromMarket")
	if err := suite.TestAPI.UnsubscribeFromMarket(tests.AssetBTS, tests.AssetCNY); err != nil {
		suite.FailNow(err.Error(), "UnsubscribeFromMarket")
	}
}

func (suite *subscribeTest) Test_SetSubscribeCallback() {

	fmt.Printf("SetSubscribeCallback: wait %s for %d incoming notifications\n",
		SetSubscribeCallbackDuration, SetSubscribeCallbackMsgs)
	if err := suite.TestAPI.SetSubscribeCallback(SetSubscribeCallbackSubscriberID, false); err != nil {
		suite.FailNow(err.Error(), "SetSubscribeCallback")
	}

	_, err := suite.TestAPI.CallWsAPI(suite.TestAPI.DatabaseAPIID(), "get_objects", []interface{}{"2.1.0"})
	if err != nil {
		suite.FailNow(err.Error(), "CallAPI->get_objects")
	}

	bar := pb.StartNew(SetSubscribeCallbackMsgs).Prefix("wait for data")
	err = suite.TestAPI.OnNotify(SetSubscribeCallbackSubscriberID, func(msg interface{}) error {
		bar.Increment()
		return nil
	})

	if err != nil {
		log.Fatal(errors.Annotate(err, "OnNotify"))
	}

	suite.Condition(func() bool {
		return util.WaitForCondition(SetSubscribeCallbackDuration, func() bool {
			return int(bar.Get()) >= SetSubscribeCallbackMsgs
		})
	}, "not enough incomming notifications during test window")

	bar.Finish()

	fmt.Println("Call CancelAllSubscriptions")
	if err := suite.TestAPI.CancelAllSubscriptions(); err != nil {
		suite.FailNow(err.Error(), "CancelAllSubscriptions")
	}
}

func TestSubscribe(t *testing.T) {
	testSuite := new(subscribeTest)
	suite.Run(t, testSuite)
	testSuite.TearDown()
}
