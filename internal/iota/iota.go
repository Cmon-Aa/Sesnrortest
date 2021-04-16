package iota

import (
	. "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/converter"
)

type Iota struct {
	seed         string
	minWeightMag uint64
	depth        uint64
	api          *API
}

func WithSeed(seed string) (*Iota, error) {

	api, err := ComposeAPI(HTTPClientSettings{URI: "https://nodes.devnet.iota.org"})
	if err != nil {
		return nil, err
	}

	iota := &Iota{
		seed:         seed,
		depth:        3,
		minWeightMag: 9,
		api:          api,
	}

	return iota, nil
}

func (iota *Iota) SendToTangle(message string) (string, error) {

	// generate a new address:
	addresses, err := iota.api.GetNewAddress(iota.seed, GetNewAddressOptions{Security: 2})
	if err != nil {
		panic(err)
	}

	// convert the plaintext message to IOTA trytes
	message, err = converter.ASCIIToTrytes(message)
	if err != nil {
		return "", err
	}

	transfers := bundle.Transfers{
		{
			Address: addresses[0],
			Value:   0,
			Message: message,
		},
	}

	trytes, err := iota.api.PrepareTransfers(iota.seed, transfers, PrepareTransfersOptions{})
	if err != nil {
		return "", err
	}

	myBundle, err := iota.api.SendTrytes(trytes, iota.depth, iota.minWeightMag)
	if err != nil {
		return "", err
	}

	return bundle.TailTransactionHash(myBundle), nil
}
