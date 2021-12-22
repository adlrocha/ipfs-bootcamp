package main

import (
	"bytes"
	"io"

	"bytes"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/multiformats/go-multicodec"

	"github.com/ipfs/go-datastore"
	ipld "github.com/ipld/go-ipld-prime"

	_ "github.com/ipld/go-ipld-prime/codec/dagjson"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
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

func mkLinkSystem(ds datastore.Batching) ipld.LinkSystem {
	lsys := cidlink.DefaultLinkSystem()
	lsys.StorageReadOpener = func(lctx ipld.LinkContext, lnk ipld.Link) (io.Reader, error) {
		c := lnk.(cidlink.Link).Cid
		val, err := ds.Get(datastore.NewKey(c.String()))
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(val), nil
	}
	lsys.StorageWriteOpener = func(lctx ipld.LinkContext) (io.Writer, ipld.BlockWriteCommitter, error) {
		buf := bytes.NewBuffer(nil)
		return buf, func(lnk ipld.Link) error {
			c := lnk.(cidlink.Link).Cid
			return ds.Put(datastore.NewKey(c.String()), buf.Bytes())
		}, nil
	}
	return lsys
}

var (
	CheckpointSchema schema.Type
	MsgMetaSchema    schema.Type

	// NoPreviousCheck is a work-around to avoid undefined CIDs,
	// that results in unexpected errors when marshalling.
	// This needs a fix in go-ipld-prime::bindnode
	NoPreviousCheck cid.Cid

	// EmptyCheckpoint is an empty checkpoint that can be Marshalled
	EmptyCheckpoint *Checkpoint
)

func init() {
	CheckpointSchema = initCheckpointSchema()
	MsgMetaSchema = initCrossMsgMetaSchema()
	var err error
	NoPreviousCheck, err = Linkproto.Sum([]byte("nil"))
	if err != nil {
		panic(err)
	}

	EmptyCheckpoint = &Checkpoint{
		Data: CheckData{
			Source:       "",
			Epoch:        0,
			PrevCheckCid: NoPreviousCheck.Bytes(),
		},
	}
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

// CrossMsgMeta includes information about the messages being propagated from and to
// a subnet.
//
// MsgsCid is the cid of the list of cids of the mesasges propagated
// for a specific subnet in that checkpoint
type CrossMsgMeta struct {
	From    string // Determines the source of the messages being propagated in MsgsCid
	To      string // Determines the destination of the messages included in MsgsCid
	MsgsCid []byte // cid.Cid
	Nonce   int    // Nonce of the msgMeta
}

// CheckData is the data included in a Checkpoint.
type CheckData struct {
	Source string
	TipSet []byte // NOTE: For simplicity we add TipSetKey. We could include full TipSet
	Epoch  int
	// NOTE: Under these bytes there's a cid.Cid. The reason for doing this is
	// to prevent the VM from interpreting it as a CID from the state
	// tree trying to fetch it and failing because it can't find anything, so we
	// are "hiding" them behing a byte type
	PrevCheckCid []byte
	Childs       []ChildCheck   // List of child checks
	CrossMsgs    []CrossMsgMeta // List with meta of msgs being propagated.
}

// Checkpoint data structure
//
// - Data includes all the data for the checkpoint. The Cid of Data
// is what identifies a checkpoint uniquely.
// - Signature adds the signature from a miner. According to the verifier
// used for checkpoint this may be different things.
type Checkpoint struct {
	Data      CheckData
	Signature []byte
}

// initCheckpointType initializes the Checkpoint schema
func initCrossMsgMetaSchema() schema.Type {
	ts := schema.TypeSystem{}
	ts.Init()
	ts.Accumulate(schema.SpawnString("String"))
	ts.Accumulate(schema.SpawnInt("Int"))
	ts.Accumulate(schema.SpawnLink("Link"))
	ts.Accumulate(schema.SpawnBytes("Bytes"))

	ts.Accumulate(schema.SpawnStruct("CrossMsgMeta",
		[]schema.StructField{
			schema.SpawnStructField("From", "String", false, false),
			schema.SpawnStructField("To", "String", false, false),
			schema.SpawnStructField("MsgsCid", "Bytes", false, false),
			schema.SpawnStructField("Nonce", "Int", false, false),
		},
		schema.SpawnStructRepresentationMap(map[string]string{}),
	))

	return ts.TypeByName("CrossMsgMeta")
}

// initCheckpointType initializes the Checkpoint schema
func initCheckpointSchema() schema.Type {
	ts := schema.TypeSystem{}
	ts.Init()
	ts.Accumulate(schema.SpawnString("String"))
	ts.Accumulate(schema.SpawnInt("Int"))
	ts.Accumulate(schema.SpawnLink("Link"))
	ts.Accumulate(schema.SpawnBytes("Bytes"))

	ts.Accumulate(schema.SpawnStruct("ChildCheck",
		[]schema.StructField{
			schema.SpawnStructField("Source", "String", false, false),
			schema.SpawnStructField("Checks", "List_Bytes", false, false),
		},
		schema.SpawnStructRepresentationMap(map[string]string{}),
	))
	ts.Accumulate(initCrossMsgMetaSchema())

	ts.Accumulate(schema.SpawnStruct("CheckData",
		[]schema.StructField{
			schema.SpawnStructField("Source", "String", false, false),
			schema.SpawnStructField("TipSet", "Bytes", false, false),
			schema.SpawnStructField("Epoch", "Int", false, false),
			schema.SpawnStructField("PrevCheckCid", "Bytes", false, false),
			schema.SpawnStructField("Childs", "List_ChildCheck", false, false),
			schema.SpawnStructField("CrossMsgs", "List_CrossMsgMeta", false, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	))
	ts.Accumulate(schema.SpawnStruct("Checkpoint",
		[]schema.StructField{
			schema.SpawnStructField("Data", "CheckData", false, false),
			schema.SpawnStructField("Signature", "Bytes", false, false),
		},
		schema.SpawnStructRepresentationMap(nil),
	))
	ts.Accumulate(schema.SpawnList("List_String", "String", false))
	ts.Accumulate(schema.SpawnList("List_Link", "Link", false))
	ts.Accumulate(schema.SpawnList("List_Bytes", "Bytes", false))
	ts.Accumulate(schema.SpawnList("List_ChildCheck", "ChildCheck", false))
	ts.Accumulate(schema.SpawnList("List_CrossMsgMeta", "CrossMsgMeta", false))

	return ts.TypeByName("Checkpoint")
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
