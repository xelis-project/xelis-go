package wallet

import (
	"log"
	"sync"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
	d "github.com/xelis-project/xelis-go-sdk/data"
)

func prepareWS(t *testing.T) (wallet *WebSocket) {
	wallet, err := NewWebSocket(config.LOCAL_WALLET_WS, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestWSGetVersion(t *testing.T) {
	wallet := prepareWS(t)
	version, err := wallet.GetVersion()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%s", version)
	wallet.Close()
}

func TestWSGetNetwork(t *testing.T) {
	wallet := prepareWS(t)
	network, err := wallet.GetNetwork()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("%s", network)
	wallet.Close()
}

func TestWSNewTopoheight(t *testing.T) {
	wallet := prepareWS(t)

	var wg sync.WaitGroup
	wg.Add(1)
	err := wallet.NewTopoheightFunc(func(newTopoheight uint64, err error) {
		if err != nil {
			log.Fatal(err)
		}

		t.Log(newTopoheight)
		wg.Done()
	})
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	wallet.Close()
}

func TestWSNewTopoheightChannel(t *testing.T) {
	wallet := prepareWS(t)

	newTopoheight, newTopoheightErr, err := wallet.NewTopoheightChannel()
	if err != nil {
		t.Fatal(err)
	}

	select {
	case topoheight := <-newTopoheight:
		t.Log(topoheight)
	case err := <-newTopoheightErr:
		t.Fatal(err)
	}

	close(newTopoheight)
	close(newTopoheightErr)
	wallet.Close()
}

func TestWSOnlineOffline(t *testing.T) {
	wallet := prepareWS(t)

	var wg sync.WaitGroup
	wg.Add(2)

	err := wallet.OnlineFunc(func() {
		t.Log("Online")
		wg.Done()
	})
	if err != nil {
		log.Fatal(err)
	}

	err = wallet.OfflineFunc(func() {
		t.Log("Offline")
		wg.Done()
	})
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	wallet.Close()
}

func TestConnectionErr(t *testing.T) {
	wallet := prepareWS(t)

	err := <-wallet.ConnectionErr() // Close the wallet connection to test
	t.Log(err)
}

func TestSignData(t *testing.T) {
	wallet := prepareWS(t)

	element := d.Element{
		Value: 3456349494,
	}

	data, err := wallet.SignData(element)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}
