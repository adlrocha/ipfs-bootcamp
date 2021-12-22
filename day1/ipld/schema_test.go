package schema

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ipfs/go-datastore"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/stretchr/testify/require"
)

func TestNFT(t *testing.T) {
	dstore := datastore.NewMapDatastore()
	lsys := mkLinkSystem(context.TODO(), dstore)

	// Create metadata
	metadata := &Metadata{
		Owner:   "myAddress",
		Network: "ethereum",
		Item:    0,
	}
	mNode := bindnode.Wrap(metadata, MetadataSchema)
	mlnk, err := lsys.Store(ipld.LinkContext{},
		Linkproto,
		mNode.Representation())
	require.NoError(t, err)

	// Create NFT
	nft := &NFT{
		Blob:     []byte("myNFT"),
		Metadata: mlnk.(cidlink.Link).Cid,
	}
	nftNode := bindnode.Wrap(nft, NFTSchema)
	nftlnk, err := lsys.Store(ipld.LinkContext{},
		Linkproto,
		nftNode.Representation())
	require.NoError(t, err)
	t.Log("Created nft with link:", nftlnk.(cidlink.Link).Cid)

}

func TestMarshalNFT(t *testing.T) {
	dstore := datastore.NewMapDatastore()
	lsys := mkLinkSystem(context.TODO(), dstore)

	// Create metadata
	metadata := &Metadata{
		Owner:   "myAddress",
		Network: "ethereum",
		Item:    0,
	}
	mNode := bindnode.Wrap(metadata, MetadataSchema)
	mlnk, err := lsys.Store(ipld.LinkContext{},
		Linkproto,
		mNode.Representation())
	require.NoError(t, err)

	// Create NFT
	nft := &NFT{
		Blob:     []byte("myNFT"),
		Metadata: mlnk.(cidlink.Link).Cid,
	}
	nftNode := bindnode.Wrap(nft, NFTSchema)
	nftlnk, err := lsys.Store(ipld.LinkContext{},
		Linkproto,
		nftNode.Representation())
	require.NoError(t, err)
	t.Log("Created nft with link:", nftlnk.(cidlink.Link).Cid)
	t.Log("Both NFTs are equal")
	t.Log(nft)

	// Marshal
	// Marshal
	bs, err := nft.MarshalBinary()
	require.NoError(t, err)

	// Unmarshal and check equal
	nft2 := &NFT{}
	err = nft2.UnmarshalBinary(bs)
	require.NoError(t, err)
	require.Equal(t, nft2, nft)

}

func TestWithoutSchemaMarshal(t *testing.T) {
	np := basicnode.Prototype.Any // Pick a prototype: this is how we decide what implementation will store the in-memory data.
	nb := np.NewBuilder()         // Create a builder.
	ma, _ := nb.BeginMap(2)       // Begin assembling a map.
	ma.AssembleKey().AssignString("hey")
	ma.AssembleValue().AssignString("it works!")
	ma.AssembleKey().AssignString("yes")
	ma.AssembleValue().AssignBool(true)
	ma.Finish()     // Call 'Finish' on the map assembly to let it know no more data is coming.
	n := nb.Build() // Call 'Build' to get the resulting Node.  (It's immutable!)

	dagjson.Encode(n, os.Stdout)

	// Output:
	// {"hey":"it works!","yes":true}
}

// Example_unmarshalData shows how you can use a Decoder
// and a NodeBuilder (or NodePrototype) together to do unmarshalling.
//
// Often you'll do this implicitly through a LinkSystem.Load call instead,
// but you can do it directly, too.
func TestWithoutSchemaUnmarshal(t *testing.T) {
	serial := strings.NewReader(`{"hey":"it works!","yes": true}`)

	np := basicnode.Prototype.Any // Pick a stle for the in-memory data.
	nb := np.NewBuilder()         // Create a builder.
	dagjson.Decode(nb, serial)    // Hand the builder to decoding -- decoding will fill it in!
	n := nb.Build()               // Call 'Build' to get the resulting Node.  (It's immutable!)

	fmt.Printf("the data decoded was a %s kind\n", n.Kind())
	fmt.Printf("the length of the node is %d\n", n.Length())

	// Output:
	// the data decoded was a map kind
	// the length of the node is 2
}
