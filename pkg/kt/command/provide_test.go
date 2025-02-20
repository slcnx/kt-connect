package command

import (
	"errors"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/alibaba/kt-connect/pkg/kt/cluster"

	"github.com/alibaba/kt-connect/pkg/kt/connect"

	fakeKt "github.com/alibaba/kt-connect/pkg/kt"
	"github.com/alibaba/kt-connect/pkg/kt/options"
	"github.com/alibaba/kt-connect/pkg/kt/util"
	"github.com/golang/mock/gomock"
	"github.com/urfave/cli"
	coreV1 "k8s.io/api/core/v1"
)

func Test_runCommand(t *testing.T) {

	ctl := gomock.NewController(t)
	fakeKtCli := fakeKt.NewMockCliInterface(ctl)

	mockAction := NewMockActionInterface(ctl)
	mockAction.EXPECT().Provide(gomock.Eq("service"), fakeKtCli, gomock.Any()).Return(nil).AnyTimes()

	cases := []struct {
		testArgs               []string
		skipFlagParsing        bool
		useShortOptionHandling bool
		expectedErr            error
	}{
		{testArgs: []string{"provide", "service", "--expose", "8080", "--external"}, skipFlagParsing: false, useShortOptionHandling: false, expectedErr: nil},
		{testArgs: []string{"provide", "service"}, skipFlagParsing: false, useShortOptionHandling: false, expectedErr: errors.New("--expose is required")},
		{testArgs: []string{"provide"}, skipFlagParsing: false, useShortOptionHandling: false, expectedErr: errors.New("an service name must be specified")},
	}

	for _, c := range cases {

		app := &cli.App{Writer: ioutil.Discard}
		set := flag.NewFlagSet("test", 0)
		_ = set.Parse(c.testArgs)

		context := cli.NewContext(app, set, nil)

		opts := options.NewDaemonOptions("test")
		opts.Debug = true
		command := NewProvideCommand(fakeKtCli, opts, mockAction)
		err := command.Run(context)

		if c.expectedErr != nil {
			if err.Error() != c.expectedErr.Error() {
				t.Errorf("expected %v but is %v", c.expectedErr, err)
			}
		} else if err != c.expectedErr {
			t.Errorf("expected %v but is %v", c.expectedErr, err)
		}
	}
}

func testDaemonOptions(labels string, opt *options.ProvideOptions) *options.DaemonOptions {
	daemonOptions := options.NewDaemonOptions("test")
	daemonOptions.WithLabels = labels
	daemonOptions.ProvideOptions = opt
	return daemonOptions
}

func getHandlers(t *testing.T) (*fakeKt.MockCliInterface, *cluster.MockKubernetesInterface, *connect.MockShadowInterface) {
	ctl := gomock.NewController(t)
	fakeKtCli := fakeKt.NewMockCliInterface(ctl)
	kubernetes := cluster.NewMockKubernetesInterface(ctl)
	shadow := connect.NewMockShadowInterface(ctl)

	fakeKtCli.EXPECT().Kubernetes().AnyTimes().Return(kubernetes, nil)
	fakeKtCli.EXPECT().Shadow().AnyTimes().Return(shadow)
	return fakeKtCli, kubernetes, shadow
}

type args struct {
	service         string
	options         *options.DaemonOptions
	shadowResponse  createShadowResponse
	serviceResponse createServiceResponse
	inboundResponse inboundResponse
}

type inboundResponse struct {
	err error
}

type createServiceResponse struct {
	service *coreV1.Service
	err     error
}

type createShadowResponse struct {
	podIP      string
	podName    string
	sshcm      string
	credential *util.SSHCredential
	err        error
}
