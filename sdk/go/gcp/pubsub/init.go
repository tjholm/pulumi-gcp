// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "gcp:pubsub/liteReservation:LiteReservation":
		r = &LiteReservation{}
	case "gcp:pubsub/liteSubscription:LiteSubscription":
		r = &LiteSubscription{}
	case "gcp:pubsub/liteTopic:LiteTopic":
		r = &LiteTopic{}
	case "gcp:pubsub/schema:Schema":
		r = &Schema{}
	case "gcp:pubsub/subscription:Subscription":
		r = &Subscription{}
	case "gcp:pubsub/subscriptionIAMBinding:SubscriptionIAMBinding":
		r = &SubscriptionIAMBinding{}
	case "gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember":
		r = &SubscriptionIAMMember{}
	case "gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy":
		r = &SubscriptionIAMPolicy{}
	case "gcp:pubsub/topic:Topic":
		r = &Topic{}
	case "gcp:pubsub/topicIAMBinding:TopicIAMBinding":
		r = &TopicIAMBinding{}
	case "gcp:pubsub/topicIAMMember:TopicIAMMember":
		r = &TopicIAMMember{}
	case "gcp:pubsub/topicIAMPolicy:TopicIAMPolicy":
		r = &TopicIAMPolicy{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/liteReservation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/liteSubscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/liteTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/schema",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/subscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/subscriptionIAMBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/subscriptionIAMMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/subscriptionIAMPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/topic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/topicIAMBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/topicIAMMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"pubsub/topicIAMPolicy",
		&module{version},
	)
}
