package pubsub

import (
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch-go/codec"
)

type BrokerOptions struct {
	coreAPI iface.CoreAPI
	codec   codec.Codec
}

type BrokerOption func(o *BrokerOptions)

func BrokerCoreAPI(coreAPI iface.CoreAPI) BrokerOption {
	return func(o *BrokerOptions) {
		o.coreAPI = coreAPI
	}
}

func BrokerCodec(codec codec.Codec) BrokerOption {
	return func(o *BrokerOptions) {
		o.codec = codec
	}
}

type SubOptions struct {
	Topic   string
	codec   codec.Codec
	coreAPI iface.CoreAPI
	Handler Handler
}

func NewSubOptions() SubOptions {
	return SubOptions{}
}

type SubOption func(o *SubOptions)

func SubCoreAPI(coreAPI iface.CoreAPI) SubOption {
	return func(o *SubOptions) {
		o.coreAPI = coreAPI
	}
}

func SubTopic(topic string) SubOption {
	return func(o *SubOptions) {
		o.Topic = topic
	}
}

func SubHandler(handler Handler) SubOption {
	return func(o *SubOptions) {
		o.Handler = handler
	}
}

func SubCodec(codec codec.Codec) SubOption {
	return func(o *SubOptions) {
		o.codec = codec
	}
}
