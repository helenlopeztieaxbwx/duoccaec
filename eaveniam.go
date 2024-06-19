import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/wallet/apiv1"
	"cloud.google.com/go/wallet/apiv1/walletpb"
)

// createAddress creates a new address for a given Wallet.
func createAddress(w io.Writer, projectID string, location string, walletID string) (*walletpb.Address, error) {
	// projectID := "my-project-id"
	// location := "us-ca1"
	// walletID := "my-wallet"

	ctx := context.Background()
	client, err := wallet.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("wallet.NewClient: %v", err)
	}
	defer client.Close()

	req := &walletpb.CreateAddressRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s/wallets/%s", projectID, location, walletID),
	}

	address, err := client.CreateAddress(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("CreateAddress: %v", err)
	}

	fmt.Fprintf(w, "Successfully created address %s\n", address.Name)

	return address, nil
}
  
