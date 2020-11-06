// +build !linux,!freebsd
// +build !confonly

package tcp

import (
	"github.com/xuanlove/v2ray-core/common/net"
	"github.com/xuanlove/v2ray-core/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
