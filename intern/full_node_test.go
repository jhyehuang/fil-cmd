package intern

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/jhyehuang/fil-cmd/config"
	"github.com/jhyehuang/fil-cmd/intern/rpc"
	"testing"
)

func TestNewFullNodeAPIV2(t *testing.T) {
	url := config.LotusNodeAddr
	token := config.LotusToken

	api, err := NewFullNodeAPIV2(url, token)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", api)

	head, err := api.ChainHead(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", head)
}

func TestNewFullNodeSign(t *testing.T) {
	url := config.LotusNodeAddr
	token := config.LotusToken

	api, err := NewFullNodeAPI(url, token)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", api)
	addr := "f3vu63vcqnnoil6snbrruqh3qyutczkabmr7jnh7hqlelccujcsi2rwgst75ykepy2pbkmgr7rqaojnr3vmj3q"
	mineraddr, err := address.NewFromString(addr)
	if err != nil {
		t.Error(err)
	}
	head, err := api.WalletSign(context.Background(), mineraddr, []byte("hello world"))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", head)
}

func TestNewFullNodeAPIV3(t *testing.T) {
	nodeAddress := config.LotusNodeAddr
	crawler := rpc.New(nodeAddress)
	addr := "f3vu63vcqnnoil6snbrruqh3qyutczkabmr7jnh7hqlelccujcsi2rwgst75ykepy2pbkmgr7rqaojnr3vmj3q"
	mineraddr, err := address.NewFromString(addr)
	if err != nil {
		t.Error(err)
	}
	err = crawler.FilWalletSign(mineraddr, []byte("hello world"))
	if err != nil {
		t.Error(err)
	}
}
func TestNewGetFullNodeAPIV1(t *testing.T) {

	api, err := NewGetFullNodeAPIV1("")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", api)
	addr := "f3qh3vw4emksrfjxlcaaokxv6kpiebzli7xhgwam7bv7gsgjtlhccu6utma3f25u2npjylj7i223y2imqjid2q"
	mineraddr, err := address.NewFromString(addr)
	if err != nil {
		t.Error(err)
	}
	head, err := api.WalletSign(context.Background(), mineraddr, []byte("hello world"))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", head)
}

func TestStateMinerProvingDeadline(t *testing.T) {
	ctx := context.Background()
	url := config.LotusNodeAddr
	token := config.LotusToken
	api, err := NewFullNodeAPI(url, token)
	if err != nil {
		t.Error(err)
	}
	//t.Logf("%+v", api)
	addr := "f01882234"
	mineraddr, err := address.NewFromString(addr)
	if err != nil {
		t.Error(err)
	}

	dl, err := api.StateMinerProvingDeadline(ctx, mineraddr, types.EmptyTSK)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", dl)
}
