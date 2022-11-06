package tcp

import (
	"context"

	"../v2fly_core/common"
	"../v2fly_core/common/net"
	"../v2fly_core/common/serial"
	"../v2fly_core/common/session"
	"../v2fly_core/transport/internet"
	"../v2fly_core/transport/internet/tls"
)

// Dial dials a new TCP connection to the given destination.
func Dial(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (internet.Connection, error) {
	newError("dialing TCP to ", dest).WriteToLog(session.ExportIDToError(ctx))
	conn, err := internet.DialSystem(ctx, dest, streamSettings.SocketSettings)
	if err != nil {
		return nil, err
	}

	if config := tls.ConfigFromStreamSettings(streamSettings); config != nil {
		tlsConfig := config.GetTLSConfig(tls.WithDestination(dest))
		/*
			if config.IsExperiment8357() {
				conn = tls.UClient(conn, tlsConfig)
			} else {
				conn = tls.Client(conn, tlsConfig)
			}
		*/
		conn = tls.Client(conn, tlsConfig)
	}

	tcpSettings := streamSettings.ProtocolSettings.(*Config)
	if tcpSettings.HeaderSettings != nil {
		headerConfig, err := serial.GetInstanceOf(tcpSettings.HeaderSettings)
		if err != nil {
			return nil, newError("failed to get header settings").Base(err).AtError()
		}
		auth, err := internet.CreateConnectionAuthenticator(headerConfig)
		if err != nil {
			return nil, newError("failed to create header authenticator").Base(err).AtError()
		}
		conn = auth.Client(conn)
	}
	return internet.Connection(conn), nil
}

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName, Dial))
}
