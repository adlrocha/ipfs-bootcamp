package schema

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/multiformats/go-multicodec"
	"golang.org/x/xerrors"
)

// Linkproto is the ipld.LinkProtocol used for the ingestion protocol.
// Refer to it if you have encoding questions.
var Linkproto = cidlink.LinkPrototype{
	Prefix: cid.Prefix{
		Version:  1,
		Codec:    uint64(multicodec.DagJson),
		MhType:   uint64(multicodec.Sha2_256),
		MhLength: 16,
	},
}

var (
	NFTSchema      schema.Type
	MetadataSchema schema.Type
)

func init() {
	NFTSchema = initNFTSchema()
	MetadataSchema = initMetadataSchema()
}

// ChildCheck
type ChildCheck struct {
	Source string
	// NOTE: Same problem as below, checks is
	// []cid.Cid, but we are hiding it behind a bunch
	// of bytes to prevent the VM from trying to fetch the
	// cid from the state tree. We still want to use IPLD
	// for now. We may be able to remove this problem
	// if we use cbor-gen directly.
	Checks [][]byte //[]cid.Cid
}

type NFT struct {
	Blob     []byte
	Metadata cid.Cid
}

type Metadata struct {
	Owner   string
	Network string
	Item    int
}

// initCheckpointType initializes the Checkpoint schema
func initNFTSchema() schema.Type {
	ts := schema.TypeSystem{}
	ts.Init()
	// Preamble
	ts.Accumulate(schema.SpawnString("String"))
	ts.Accumulate(schema.SpawnInt("Int"))
	ts.Accumulate(schema.SpawnLink("Link"))
	ts.Accumulate(schema.SpawnBytes("Bytes"))

	ts.Accumulate(schema.SpawnStruct("NFT",
		[]schema.StructField{
			schema.SpawnStructField("Blob", "Bytes", false, false),
			schema.SpawnStructField("Metadata", "Link", false, false),
		},
		schema.SpawnStructRepresentationMap(map[string]string{}),
	))
	ts.Accumulate(schema.SpawnStruct("Metadata",
		[]schema.StructField{
			schema.SpawnStructField("Owner", "String", false, false),
			schema.SpawnStructField("Network", "String", false, false),
			schema.SpawnStructField("Item", "Int", false, false),
		},
		schema.SpawnStructRepresentationMap(map[string]string{}),
	))

	return ts.TypeByName("NFT")
}

func initMetadataSchema() schema.Type {
	ts := schema.TypeSystem{}
	ts.Init()
	// Preamble
	ts.Accumulate(schema.SpawnString("String"))
	ts.Accumulate(schema.SpawnInt("Int"))
	ts.Accumulate(schema.SpawnLink("Link"))
	ts.Accumulate(schema.SpawnBytes("Bytes"))

	ts.Accumulate(schema.SpawnStruct("Metadata",
		[]schema.StructField{
			schema.SpawnStructField("Owner", "String", false, false),
			schema.SpawnStructField("Network", "String", false, false),
			schema.SpawnStructField("Item", "Int", false, false),
		},
		schema.SpawnStructRepresentationMap(map[string]string{}),
	))

	return ts.TypeByName("Metadata")
}

func (nft *NFT) MarshalBinary() ([]byte, error) {
	node := bindnode.Wrap(nft, NFTSchema)
	nodeRepr := node.Representation()
	var buf bytes.Buffer
	err := dagjson.Encode(nodeRepr, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (nft *NFT) UnmarshalBinary(b []byte) error {
	nb := bindnode.Prototype(nft, NFTSchema).NewBuilder()
	err := dagjson.Decode(nb, bytes.NewReader(b))
	if err != nil {
		return err
	}
	n := bindnode.Unwrap(nb.Build())

	ch, ok := n.(*NFT)
	if !ok {
		return xerrors.Errorf("Unmarshalled node not of type Checkpoint")
	}
	*nft = *ch
	return nil
}

// Dumb linksystem used to generate links
//
// This linksystem doesn't store anything, just computes the Cid
// for a node.
func noStoreLinkSystem() ipld.LinkSystem {
	lsys := cidlink.DefaultLinkSystem()
	lsys.StorageWriteOpener = func(lctx ipld.LinkContext) (io.Writer, ipld.BlockWriteCommitter, error) {
		buf := bytes.NewBuffer(nil)
		return buf, func(lnk ipld.Link) error {
			return nil
		}, nil
	}
	return lsys
}

// LinkSystem determines what to do when coming across links
func mkLinkSystem(ctx context.Context, ds datastore.Batching) ipld.LinkSystem {
	lsys := cidlink.DefaultLinkSystem()
	lsys.StorageReadOpener = func(lctx ipld.LinkContext, lnk ipld.Link) (io.Reader, error) {
		fmt.Println(">> Came across a link while loading in lsys", lnk.(cidlink.Link).Cid)
		c := lnk.(cidlink.Link).Cid
		val, err := ds.Get(ctx, datastore.NewKey(c.String()))
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(val), nil
	}
	lsys.StorageWriteOpener = func(lctx ipld.LinkContext) (io.Writer, ipld.BlockWriteCommitter, error) {
		buf := bytes.NewBuffer(nil)
		return buf, func(lnk ipld.Link) error {
			c := lnk.(cidlink.Link).Cid
			return ds.Put(ctx, datastore.NewKey(c.String()), buf.Bytes())
		}, nil
	}
	return lsys
}
