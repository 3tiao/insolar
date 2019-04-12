//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package transport

import (
	"context"
	"io"
	"net"
	"strings"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/network/hostnetwork/resolver"
)

type tcpTransport struct {
	listener  net.Listener
	address   string
	processor StreamHandler
	cancel    context.CancelFunc
}

func newTCPTransport(listenAddress, fixedPublicAddress string) (*tcpTransport, string, error) {

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to listen UDP")
	}
	publicAddress, err := resolver.Resolve(fixedPublicAddress, listener.Addr().String())
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to resolve public address")
	}

	transport := &tcpTransport{
		listener: listener,
	}

	return transport, publicAddress, nil
}

func (t *tcpTransport) SetStreamHandler(processor StreamHandler) {
	t.processor = processor
}

func (t *tcpTransport) Dial(ctx context.Context, address string) (io.ReadWriteCloser, error) {
	logger := inslogger.FromContext(ctx)
	tcpAddress, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, errors.New("[ Dial ] Failed to get tcp address")
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		logger.Errorf("[ Dial ] Failed to open connection to %s: %s", address, err.Error())
		return nil, errors.Wrap(err, "[ Dial ] Failed to open connection")
	}

	err = conn.SetKeepAlive(true)
	if err != nil {
		logger.Error("[ Dial ] Failed to set keep alive")
	}

	err = conn.SetNoDelay(true)
	if err != nil {
		logger.Error("[ Dial ] Failed to set connection no delay: ", err.Error())
	}

	return conn, nil
}

func (t *tcpTransport) prepareListen() (net.Listener, error) {
	if t.listener != nil {
		t.address = t.listener.Addr().String()
	} else {
		var err error
		t.listener, err = net.Listen("tcp", t.address)
		if err != nil {
			return nil, errors.Wrap(err, "failed to listen TCP")
		}
	}

	return t.listener, nil
}

// Start starts networking.
func (t *tcpTransport) Start(ctx context.Context) error {
	logger := inslogger.FromContext(ctx)
	logger.Info("[ Start ] Start TCP transport")
	ctx, t.cancel = context.WithCancel(ctx)

	go t.listen(ctx)
	return nil
}

func (t *tcpTransport) listen(ctx context.Context) {
	logger := inslogger.FromContext(ctx)

	listener, err := t.prepareListen()
	if err != nil {
		logger.Info("[ Start ] Failed to prepare TCP transport: ", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "use of closed network connection") {
				logger.Info("Connection closed, quiting accept loop")
				return
			}

			logger.Error("[ listen ] Failed to accept connection: ", err.Error())
			return
		}

		logger.Infof("[ listen ] Accepted new connection from %s", conn.RemoteAddr())

		// t.pool.AddConnection(ctx, conn)
		go t.handleAcceptedConnection(conn)

		// select {
		// case <-ctx.Done():
		// 	return
		// default:
		// }
	}
}

// Stop stops networking.
func (t *tcpTransport) Stop(ctx context.Context) error {
	log.Info("[ Stop ] Stop TCP transport")
	err := t.listener.Close()
	//	t.cancel()
	return err
}

func (t *tcpTransport) handleAcceptedConnection(conn net.Conn) {

	err := t.processor.HandleStream(conn.RemoteAddr().String(), conn)
	if err != nil {
		inslogger.FromContext(context.Background()).Error("failed to process TCP stream: ", err.Error())
	}
}
