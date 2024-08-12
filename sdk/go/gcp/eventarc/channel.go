// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventarc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Eventarc Channel resource
//
// ## Example Usage
//
// ### Basic
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/eventarc"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/kms"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testProject, err := organizations.LookupProject(ctx, &organizations.LookupProjectArgs{
//				ProjectId: pulumi.StringRef("my-project-name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			testKeyRing, err := kms.GetKMSKeyRing(ctx, &kms.GetKMSKeyRingArgs{
//				Name:     "keyring",
//				Location: "us-west1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = kms.GetKMSCryptoKey(ctx, &kms.GetKMSCryptoKeyArgs{
//				Name:    "key",
//				KeyRing: testKeyRing.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			key1Member, err := kms.NewCryptoKeyIAMMember(ctx, "key1_member", &kms.CryptoKeyIAMMemberArgs{
//				CryptoKeyId: pulumi.Any(key1.Id),
//				Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
//				Member:      pulumi.Sprintf("serviceAccount:service-%v@gcp-sa-eventarc.iam.gserviceaccount.com", testProject.Number),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eventarc.NewChannel(ctx, "primary", &eventarc.ChannelArgs{
//				Location:           pulumi.String("us-west1"),
//				Name:               pulumi.String("channel"),
//				Project:            pulumi.String(testProject.ProjectId),
//				CryptoKeyName:      pulumi.Any(key1.Id),
//				ThirdPartyProvider: pulumi.Sprintf("projects/%v/locations/us-west1/providers/datadog", testProject.ProjectId),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				key1Member,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Channel can be imported using any of these accepted formats:
//
// * `projects/{{project}}/locations/{{location}}/channels/{{name}}`
//
// * `{{project}}/{{location}}/{{name}}`
//
// * `{{location}}/{{name}}`
//
// When using the `pulumi import` command, Channel can be imported using one of the formats above. For example:
//
// ```sh
// $ pulumi import gcp:eventarc/channel:Channel default projects/{{project}}/locations/{{location}}/channels/{{name}}
// ```
//
// ```sh
// $ pulumi import gcp:eventarc/channel:Channel default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
// $ pulumi import gcp:eventarc/channel:Channel default {{location}}/{{name}}
// ```
type Channel struct {
	pulumi.CustomResourceState

	// Output only. The activation token for the channel. The token must be used by the provider to register the channel for publishing.
	ActivationToken pulumi.StringOutput `pulumi:"activationToken"`
	// Output only. The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	CryptoKeyName pulumi.StringPtrOutput `pulumi:"cryptoKeyName"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Required. The resource name of the channel. Must be unique within the location on the project.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{project}/topics/{topic_id}`.
	PubsubTopic pulumi.StringOutput `pulumi:"pubsubTopic"`
	// Output only. The state of a Channel. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, INACTIVE
	State pulumi.StringOutput `pulumi:"state"`
	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
	ThirdPartyProvider pulumi.StringPtrOutput `pulumi:"thirdPartyProvider"`
	// Output only. Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The last-modified time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Channel
	err := ctx.RegisterResource("gcp:eventarc/channel:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("gcp:eventarc/channel:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
	// Output only. The activation token for the channel. The token must be used by the provider to register the channel for publishing.
	ActivationToken *string `pulumi:"activationToken"`
	// Output only. The creation time.
	CreateTime *string `pulumi:"createTime"`
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	CryptoKeyName *string `pulumi:"cryptoKeyName"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Required. The resource name of the channel. Must be unique within the location on the project.
	//
	// ***
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{project}/topics/{topic_id}`.
	PubsubTopic *string `pulumi:"pubsubTopic"`
	// Output only. The state of a Channel. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, INACTIVE
	State *string `pulumi:"state"`
	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
	ThirdPartyProvider *string `pulumi:"thirdPartyProvider"`
	// Output only. Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid *string `pulumi:"uid"`
	// Output only. The last-modified time.
	UpdateTime *string `pulumi:"updateTime"`
}

type ChannelState struct {
	// Output only. The activation token for the channel. The token must be used by the provider to register the channel for publishing.
	ActivationToken pulumi.StringPtrInput
	// Output only. The creation time.
	CreateTime pulumi.StringPtrInput
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	CryptoKeyName pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Required. The resource name of the channel. Must be unique within the location on the project.
	//
	// ***
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{project}/topics/{topic_id}`.
	PubsubTopic pulumi.StringPtrInput
	// Output only. The state of a Channel. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, INACTIVE
	State pulumi.StringPtrInput
	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
	ThirdPartyProvider pulumi.StringPtrInput
	// Output only. Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid pulumi.StringPtrInput
	// Output only. The last-modified time.
	UpdateTime pulumi.StringPtrInput
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	CryptoKeyName *string `pulumi:"cryptoKeyName"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Required. The resource name of the channel. Must be unique within the location on the project.
	//
	// ***
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
	ThirdPartyProvider *string `pulumi:"thirdPartyProvider"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
	CryptoKeyName pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringInput
	// Required. The resource name of the channel. Must be unique within the location on the project.
	//
	// ***
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
	ThirdPartyProvider pulumi.StringPtrInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

// ChannelArrayInput is an input type that accepts ChannelArray and ChannelArrayOutput values.
// You can construct a concrete instance of `ChannelArrayInput` via:
//
//	ChannelArray{ ChannelArgs{...} }
type ChannelArrayInput interface {
	pulumi.Input

	ToChannelArrayOutput() ChannelArrayOutput
	ToChannelArrayOutputWithContext(context.Context) ChannelArrayOutput
}

type ChannelArray []ChannelInput

func (ChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Channel)(nil)).Elem()
}

func (i ChannelArray) ToChannelArrayOutput() ChannelArrayOutput {
	return i.ToChannelArrayOutputWithContext(context.Background())
}

func (i ChannelArray) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelArrayOutput)
}

// ChannelMapInput is an input type that accepts ChannelMap and ChannelMapOutput values.
// You can construct a concrete instance of `ChannelMapInput` via:
//
//	ChannelMap{ "key": ChannelArgs{...} }
type ChannelMapInput interface {
	pulumi.Input

	ToChannelMapOutput() ChannelMapOutput
	ToChannelMapOutputWithContext(context.Context) ChannelMapOutput
}

type ChannelMap map[string]ChannelInput

func (ChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Channel)(nil)).Elem()
}

func (i ChannelMap) ToChannelMapOutput() ChannelMapOutput {
	return i.ToChannelMapOutputWithContext(context.Background())
}

func (i ChannelMap) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelMapOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

// Output only. The activation token for the channel. The token must be used by the provider to register the channel for publishing.
func (o ChannelOutput) ActivationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.ActivationToken }).(pulumi.StringOutput)
}

// Output only. The creation time.
func (o ChannelOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
func (o ChannelOutput) CryptoKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.CryptoKeyName }).(pulumi.StringPtrOutput)
}

// The location for the resource
func (o ChannelOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Required. The resource name of the channel. Must be unique within the location on the project.
//
// ***
func (o ChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o ChannelOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{project}/topics/{topic_id}`.
func (o ChannelOutput) PubsubTopic() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.PubsubTopic }).(pulumi.StringOutput)
}

// Output only. The state of a Channel. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE, INACTIVE
func (o ChannelOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
func (o ChannelOutput) ThirdPartyProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.ThirdPartyProvider }).(pulumi.StringPtrOutput)
}

// Output only. Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
func (o ChannelOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The last-modified time.
func (o ChannelOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type ChannelArrayOutput struct{ *pulumi.OutputState }

func (ChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Channel)(nil)).Elem()
}

func (o ChannelArrayOutput) ToChannelArrayOutput() ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) Index(i pulumi.IntInput) ChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Channel {
		return vs[0].([]*Channel)[vs[1].(int)]
	}).(ChannelOutput)
}

type ChannelMapOutput struct{ *pulumi.OutputState }

func (ChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Channel)(nil)).Elem()
}

func (o ChannelMapOutput) ToChannelMapOutput() ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) MapIndex(k pulumi.StringInput) ChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Channel {
		return vs[0].(map[string]*Channel)[vs[1].(string)]
	}).(ChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelInput)(nil)).Elem(), &Channel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelArrayInput)(nil)).Elem(), ChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelMapInput)(nil)).Elem(), ChannelMap{})
	pulumi.RegisterOutputType(ChannelOutput{})
	pulumi.RegisterOutputType(ChannelArrayOutput{})
	pulumi.RegisterOutputType(ChannelMapOutput{})
}
