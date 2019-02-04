package huobi

import (
	"github.com/jason-wj/GoEx"
	"github.com/jason-wj/bitesla/service/service-exchange/exchange"
	"github.com/jason-wj/bitesla/service/service-exchange/proto"
	"github.com/stretchr/testify/assert"
	"log"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return &url.URL{
				Scheme: "socks5",
				Host:   "127.0.0.1:1080"}, nil
		},
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	},
	Timeout: 10 * time.Second,
}

var (
	apikey    = ""
	secretkey = ""
)

//
var hbpro = NewHuoBiProSpot(httpProxyClient, apikey, secretkey)

func TestHuobiPro_GetTicker(t *testing.T) {
	return
	ticker, err := hbpro.GetTicker(bitesla_srv_trader.XrpBtc)
	assert.Nil(t, err)
	t.Log(ticker)
}

func TestHuobiPro_GetDepth(t *testing.T) {
	return
	dep, err := hbpro.GetDepth(2, bitesla_srv_trader.LtcUsdt)
	assert.Nil(t, err)
	t.Log(dep.AskList)
	t.Log(dep.BidList)
}

func TestHuobiPro_GetAccountInfo(t *testing.T) {
	return
	info, err := hbpro.GetAccountInfo("point")
	assert.Nil(t, err)
	t.Log(info)
}

//获取点卡剩余
func TestHuoBiPro_GetPoint(t *testing.T) {
	return
	point := NewHuoBiProPoint(httpProxyClient, apikey, secretkey)
	acc, _ := point.GetAccount()
	t.Log(acc.SubAccounts[HBPOINT])
}

//获取现货资产信息
func TestHuobiPro_GetAccount(t *testing.T) {
	return
	acc, err := hbpro.GetAccount()
	assert.Nil(t, err)
	t.Log(acc.SubAccounts)
}

func TestHuobiPro_LimitBuy(t *testing.T) {
	return
	ord, err := hbpro.LimitBuy("", "0.09122", bitesla_srv_trader.BccBtc)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_LimitSell(t *testing.T) {
	return
	ord, err := hbpro.LimitSell("1", "0.212", bitesla_srv_trader.BccBtc)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_MarketSell(t *testing.T) {
	return
	ord, err := hbpro.MarketSell("0.1738", "0.212", bitesla_srv_trader.BccBtc)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_MarketBuy(t *testing.T) {
	return
	ord, err := hbpro.MarketBuy("0.02", "", bitesla_srv_trader.BccBtc)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_GetUnfinishOrders(t *testing.T) {
	return
	ords, err := hbpro.GetUnfinishOrders(bitesla_srv_trader.EtcUsdt)
	assert.Nil(t, err)
	t.Log(ords)
}

func TestHuobiPro_CancelOrder(t *testing.T) {
	return
	r, err := hbpro.CancelOrder("600329873", bitesla_srv_trader.EthUsdt)
	assert.Nil(t, err)
	t.Log(r)
	t.Log(err)
}

func TestHuobiPro_GetOneOrder(t *testing.T) {
	return
	ord, err := hbpro.GetOneOrder("1116237737", bitesla_srv_trader.LtcBtc)
	assert.Nil(t, err)
	t.Log(ord)
}

func TestHuobiPro_GetOrderHistorys(t *testing.T) {
	return
	ords, err := hbpro.GetOrderHistorys(bitesla_srv_trader.NewCurrencyPair2("HT_USDT"), 1, 3)
	t.Log(err)
	t.Log(ords)
}

func TestHuobiPro_GetDepthWithWs(t *testing.T) {
	return
	hbpro.GetDepthWithWs(bitesla_srv_trader.BtcUsdt, func(dep *exchange.Depth) {
		log.Println("%+v", *dep)
	})
	time.Sleep(time.Minute)
}

func TestHuobiPro_GetTickerWithWs(t *testing.T) {
	return
	hbpro.GetTickerWithWs(bitesla_srv_trader.BtcUsdt, func(ticker *exchange.Ticker) {
		log.Println("%+v", *ticker)
	})
	time.Sleep(time.Minute)
}

func TestHuobiPro_GetKLineWithWs(t *testing.T) {
	return
	hbpro.GetKLineWithWs(bitesla_srv_trader.BtcUsdt, goex.KLINE_PERIOD_60MIN, func(kline *exchange.Kline) {
		log.Println("%+v", *kline)
	})
	time.Sleep(time.Minute)
}

func TestHuobiPro_GetCurrenciesList(t *testing.T) {
	return
	hbpro.GetCurrenciesList()
}

func TestHuobiPro_GetCurrenciesPrecision(t *testing.T) {
	//return
	t.Log(hbpro.GetCurrenciesPrecision())
}