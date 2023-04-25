package intern

import (
	"context"
	"flag"
	"fmt"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/jhyehuang/fil-cmd/config"
	"github.com/jhyehuang/fil-cmd/intern/log"
	"github.com/rickiey/loggo"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
)

func NewFullNodeAPI(ApiURL string, token string) (api.FullNode, error) {

	var ctx = context.Background()

	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+token)

	lv1api, _, err := client.NewFullNodeRPCV1(ctx, ApiURL, headers)
	if err != nil {
		log.Errorf("connecting with lotus failed: %s", err)
		return nil, err
	}

	v, err := lv1api.Version(ctx)
	if err != nil {
		loggo.Errorf("Version: %v", err)
		return nil, err
	}
	fmt.Printf("Lotus Version: %s", v.Version)

	return lv1api, nil
}

func NewFullNodeAPIV2(ApiURL string, token string) (api.FullNode, error) {

	headers := http.Header{}

	headers.Add("Authorization", "Bearer "+token)

	var lapi api.FullNodeStruct
	if ApiURL == "" {
		ApiURL = config.LotusNodeAddr
	}

	_, err := jsonrpc.NewMergeClient(context.Background(), ApiURL, "Filecoin", []interface{}{&lapi.Internal, &lapi.CommonStruct.Internal}, headers, jsonrpc.WithErrors(api.RPCErrors))
	if err != nil {
		loggo.Panicf("connecting with lotus failed: %s", err)
	}
	return &lapi, nil
}

func NewGetFullNodeAPIV1(envValue string) (api.FullNode, error) {

	app := &cli.App{}
	set := flag.NewFlagSet("test", 0)
	set.Set("api", "")
	var ctx = cli.NewContext(app, set, nil)
	fmt.Printf("FULLNODE_API_INFO: %s \n", os.Getenv("FULLNODE_API_INFO"))
	os.Setenv("FULLNODE_API_INFO", envValue)

	lv1api, _, err := cliutil.GetFullNodeAPIV1(ctx)
	if err != nil {
		loggo.Panicf("connecting with lotus failed: %s", err)
	}

	v, err := lv1api.Version(context.TODO())
	if err != nil {
		loggo.Errorf("Version: %v", err)
		return nil, err
	}
	fmt.Printf("Lotus Version: %s", v.Version)
	return lv1api, nil
	//defer closerv1()
}
