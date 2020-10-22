// MIT License
//
// Copyright (c) 2020 Tweag IO
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package transport

import (
	"context"
	"crypto/ed25519"
	"crypto/tls"
	"fmt"
	"github.com/tweag/trustix/config"
	pb "github.com/tweag/trustix/proto"
	"github.com/tweag/trustix/signer"
	"github.com/tweag/trustix/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

type grpcTransport struct {
	c pb.TrustixKVClient
}

type grpcTxn struct {
	c pb.TrustixKVClient
}

func (t *grpcTxn) Get(bucket []byte, key []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := t.c.Get(ctx, &pb.KVRequest{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return nil, err
	}

	return r.Value, nil
}

func (t *grpcTxn) Set(bucket []byte, key []byte, value []byte) error {
	return fmt.Errorf("Cannot Set on remote")
}

func (t *grpcTxn) Size(bucket []byte) (int, error) {
	return 0, fmt.Errorf("Cannot Size on remote")
}

func NewGRPCTransport(t *config.GRPCTransportConfig, s signer.TrustixSigner) (*grpcTransport, error) {

	config := &tls.Config{
		InsecureSkipVerify: true,
		VerifyConnection: func(state tls.ConnectionState) error {
			if len(state.PeerCertificates) != 1 {
				return fmt.Errorf("Dont know how to handle %d certs", len(state.PeerCertificates))
			}

			cert := state.PeerCertificates[0]

			edPub, ok := cert.PublicKey.(ed25519.PublicKey)
			if !ok {
				return fmt.Errorf("Key not ed25519")
			}

			if !edPub.Equal(s.Public()) {
				return fmt.Errorf("Expected key mismatch")
			}

			err := cert.CheckSignature(cert.SignatureAlgorithm, cert.RawTBSCertificate, cert.Signature)
			if err != nil {
				fmt.Errorf("Signature check failed %d certs", len(state.PeerCertificates))
				return err
			}

			return nil
		},
	}

	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(t.Remote, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}
	c := pb.NewTrustixKVClient(conn)

	return &grpcTransport{
		c: c,
	}, nil
}

func (g *grpcTransport) Close() {
}

func (s *grpcTransport) runTX(fn func(storage.Transaction) error) error {
	t := &grpcTxn{
		c: s.c,
	}

	err := fn(t)
	if err != nil {
		return err
	}

	return err
}

func (s *grpcTransport) View(fn func(storage.Transaction) error) error {
	return s.runTX(fn)
}

func (s *grpcTransport) Update(fn func(storage.Transaction) error) error {
	return s.runTX(fn)
}
