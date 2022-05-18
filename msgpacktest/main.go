package main

import (
	"github.com/ugorji/go/codec"
	"io"
	"net"
	"net/rpc"
	"reflect"
)

// create and configure Handle
var (
	mh codec.MsgpackHandle
)

func main(){
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))

	// configure extensions
	// e.g. for msgpack, define functions and enable Time support for tag 1
	// mh.SetExt(reflect.TypeOf(time.Time{}), 1, myExt)

	// create and use decoder/encoder
	var (
		r io.Reader
		w io.Writer
		b []byte
		h = &mh
	)

	dec := codec.NewDecoder(r, h)
	dec = codec.NewDecoderBytes(b, h)
	err := dec.Decode(&v)

	enc := codec.NewEncoder(w, h)
	enc = codec.NewEncoderBytes(&b, h)
	err = enc.Encode(v)

	// RPC Server
	go func() {
		for {
			conn, err := listener.Accept()
			rpcCodec := codec.GoRpc.ServerCodec(conn, h)
			//OR rpcCodec := codec.MsgpackSpecRpc.ServerCodec(conn, h)
			rpc.ServeCodec(rpcCodec)
		}
	}()

	// RPC Communication (client side)
	conn, err = net.Dial("tcp", "localhost:5555")
	rpcCodec := codec.GoRpc.ClientCodec(conn, h)
	//OR rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, h)
	client := rpc.NewClientWithCodec(rpcCodec)
}
