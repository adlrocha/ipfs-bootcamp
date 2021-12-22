package main

import (
	"bufio"
	"context"
	"fmt"
	"time"

	"github.com/adlrocha/libp2p-boilerplate/cbor-example/msg"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	mplex "github.com/libp2p/go-libp2p-mplex"
	routing "github.com/libp2p/go-libp2p-routing"
	tls "github.com/libp2p/go-libp2p-tls"
	yamux "github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
)

type h struct {
	ch chan bool
}

var pid protocol.ID = "/test/v0"

func main() {
	ctx := context.Background()
	// Instantiating hosts
	h1, err := startHostWithOptions()
	if err != nil {
		panic(err)
	}
	h2, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	defer h1.Close()
	defer h2.Close()

	fmt.Println("[*] Connecting peers")
	// Connect h1-h2
	err = DialOtherPeer(ctx, h1, *host.InfoFromHost(h2))
	if err != nil {
		panic(err)
	}

	h := h{ch: make(chan bool)}

	// Set handler
	h1.SetStreamHandler(pid, h.handleNewStream)
	time.Sleep(time.Second)
	fmt.Println("Starting to send messages")

	// TODO: Try sending a marshalled IPLD node using libp2p
	send(ctx, h2, h1.ID(), msg.Msg{Data: []byte("test"), Err: 1})

	<-h.ch
}

func (h *h) handleNewStream(s network.Stream) {
	obj := msg.Msg{}
	err := cborutil.ReadCborRPC(bufio.NewReader(s), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message Received Successfully with DATA: ", string(obj.Data))
	close(h.ch)

}

func send(ctx context.Context, h host.Host, p peer.ID, obj msg.Msg) error {

	s, err := h.NewStream(ctx, p, []protocol.ID{pid}...)
	//defer s.Close()
	if err != nil {
		return err
	}
	buffered := bufio.NewWriter(s)
	err = cborutil.WriteCborRPC(buffered, &obj)
	if err != nil {
		panic(err)
	}
	err = buffered.Flush()
	fmt.Println("Message sent successfully")
	return nil
}

// DialOtherPeers connects to a set of peers in the experiment.
func DialOtherPeer(ctx context.Context, self host.Host, ai peer.AddrInfo) error {
	if err := self.Connect(ctx, ai); err != nil {
		return fmt.Errorf("Error while dialing peer %v: %w", ai.Addrs, err)
	}
	return nil
}

func startHostWithOptions() (host.Host, error) {
	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(ws.New),
	)

	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport),
		libp2p.Muxer("/mplex/6.7.0", mplex.DefaultTransport),
	)

	security := libp2p.Security(tls.ID, tls.New)

	listenAddrs := libp2p.ListenAddrStrings(
		"/ip4/0.0.0.0/tcp/0",
		"/ip4/0.0.0.0/tcp/0/ws",
	)

	var dht *kaddht.IpfsDHT
	newDHT := func(h host.Host) (routing.PeerRouting, error) {
		var err error
		dht, err = kaddht.New(context.TODO(), h)
		return dht, err
	}
	routing := libp2p.Routing(newDHT)
	return libp2p.New(
		transports,
		listenAddrs,
		muxers,
		security,
		routing)
}
