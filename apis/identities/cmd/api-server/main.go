package main

import (
	"context"
	"log"
	"net"

	identities "github.com/go-cqrses/teddyshop/spec/identities/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	srv := grpc.NewServer()
	impl := &identitiesImpl{}
	identities.RegisterIdentitiesServer(srv, impl)

	nl, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("unable to `", err)
	}
	log.Fatal(srv.Serve(nl))
}

type (
	identitiesImpl struct {
	}
)


func (i *identitiesImpl) Login(context.Context, *identities.LoginRequest) (*identities.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

func (i *identitiesImpl) TokenRevoke(context.Context, *identities.TokenRevokeRequest) (*identities.TokenRevokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

func (i *identitiesImpl) TokenCheck(context.Context, *identities.TokenCheckRequest) (*identities.TokenCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

func (i *identitiesImpl) CreateAccount(context.Context, *identities.CreateAccountRequest) (*identities.CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

func (i *identitiesImpl) QueryAccounts(context.Context, *identities.QueryAccountsRequest) (*identities.QueryAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

func (i *identitiesImpl) GetAccount(context.Context, *identities.GetAccountRequest) (*identities.GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

func (i *identitiesImpl) UpdateAccountPassword(context.Context, *identities.UpdateAccountPasswordRequest) (*identities.UpdateAccountPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

func (i *identitiesImpl) UpdateAccountDetails(context.Context, *identities.UpdateAccountDetailsRequest) (*identities.UpdateAccountDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

func (i *identitiesImpl) GetAccountWithCredentials(context.Context, *identities.GetAccountWithCredentialsRequest) (*identities.GetAccountWithCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "coming soon to a teddyshop near you")
}

