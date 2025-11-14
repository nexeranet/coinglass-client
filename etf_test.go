package coinglassclient

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestEthereumETFList(t *testing.T) {
	secret := os.Getenv("KEY")
	if secret == "" {
		t.Fatal("Key is empty")
	}
	client := NewClient(secret, DefaultAPIUrl, time.Minute*5)
	ctx := context.Background()
	result, err := client.EthereumETFList(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range result {
		t.Log(item)
	}

}

func TestBitcoinETFList(t *testing.T) {
	secret := os.Getenv("KEY")
	if secret == "" {
		t.Fatal("Key is empty")
	}
	client := NewClient(secret, DefaultAPIUrl, time.Minute*5)
	ctx := context.Background()
	result, err := client.BitcoinETFList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range result {
		t.Log(item.AssetDetails)
	}

}
