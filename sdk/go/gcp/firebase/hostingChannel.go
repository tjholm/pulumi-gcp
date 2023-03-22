// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firebase

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Firebasehosting Channel Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultHostingSite, err := firebase.NewHostingSite(ctx, "defaultHostingSite", &firebase.HostingSiteArgs{
//				Project: pulumi.String("my-project-name"),
//				SiteId:  pulumi.String("site-with-channel"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = firebase.NewHostingChannel(ctx, "defaultHostingChannel", &firebase.HostingChannelArgs{
//				SiteId:    defaultHostingSite.SiteId,
//				ChannelId: pulumi.String("channel-basic"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Firebasehosting Channel Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firebase.NewHostingSite(ctx, "default", &firebase.HostingSiteArgs{
//				Project: pulumi.String("my-project-name"),
//				SiteId:  pulumi.String("site-with-channel"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = firebase.NewHostingChannel(ctx, "full", &firebase.HostingChannelArgs{
//				SiteId:               _default.SiteId,
//				ChannelId:            pulumi.String("channel-full"),
//				Ttl:                  pulumi.String("86400s"),
//				RetainedReleaseCount: pulumi.Int(20),
//				Labels: pulumi.StringMap{
//					"some-key": pulumi.String("some-value"),
//				},
//			}, pulumi.Provider(google_beta))
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
// # Channel can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:firebase/hostingChannel:HostingChannel default sites/{{site_id}}/channels/{{channel_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firebase/hostingChannel:HostingChannel default {{site_id}}/{{channel_id}}
//
// ```
type HostingChannel struct {
	pulumi.CustomResourceState

	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId pulumi.StringOutput `pulumi:"channelId"`
	// The time at which the channel will be automatically deleted. If null, the channel
	// will not be automatically deleted. This field is present in the output whether it's
	// set directly or via the `ttl` field.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// Text labels used for extra metadata and/or filtering
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format:
	// sites/SITE_ID/channels/CHANNEL_ID
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of previous releases to retain on the channel for rollback or other
	// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.IntOutput `pulumi:"retainedReleaseCount"`
	// Required. The ID of the site in which to create this channel.
	SiteId pulumi.StringOutput `pulumi:"siteId"`
	// Input only. A time-to-live for this channel. Sets `expireTime` to the provided
	// duration past the time of the request. A duration in seconds with up to nine fractional
	// digits, terminated by 's'. Example: "86400s" (one day).
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
}

// NewHostingChannel registers a new resource with the given unique name, arguments, and options.
func NewHostingChannel(ctx *pulumi.Context,
	name string, args *HostingChannelArgs, opts ...pulumi.ResourceOption) (*HostingChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChannelId == nil {
		return nil, errors.New("invalid value for required argument 'ChannelId'")
	}
	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	var resource HostingChannel
	err := ctx.RegisterResource("gcp:firebase/hostingChannel:HostingChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostingChannel gets an existing HostingChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostingChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostingChannelState, opts ...pulumi.ResourceOption) (*HostingChannel, error) {
	var resource HostingChannel
	err := ctx.ReadResource("gcp:firebase/hostingChannel:HostingChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostingChannel resources.
type hostingChannelState struct {
	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId *string `pulumi:"channelId"`
	// The time at which the channel will be automatically deleted. If null, the channel
	// will not be automatically deleted. This field is present in the output whether it's
	// set directly or via the `ttl` field.
	ExpireTime *string `pulumi:"expireTime"`
	// Text labels used for extra metadata and/or filtering
	Labels map[string]string `pulumi:"labels"`
	// The fully-qualified resource name for the channel, in the format:
	// sites/SITE_ID/channels/CHANNEL_ID
	Name *string `pulumi:"name"`
	// The number of previous releases to retain on the channel for rollback or other
	// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount *int `pulumi:"retainedReleaseCount"`
	// Required. The ID of the site in which to create this channel.
	SiteId *string `pulumi:"siteId"`
	// Input only. A time-to-live for this channel. Sets `expireTime` to the provided
	// duration past the time of the request. A duration in seconds with up to nine fractional
	// digits, terminated by 's'. Example: "86400s" (one day).
	Ttl *string `pulumi:"ttl"`
}

type HostingChannelState struct {
	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId pulumi.StringPtrInput
	// The time at which the channel will be automatically deleted. If null, the channel
	// will not be automatically deleted. This field is present in the output whether it's
	// set directly or via the `ttl` field.
	ExpireTime pulumi.StringPtrInput
	// Text labels used for extra metadata and/or filtering
	Labels pulumi.StringMapInput
	// The fully-qualified resource name for the channel, in the format:
	// sites/SITE_ID/channels/CHANNEL_ID
	Name pulumi.StringPtrInput
	// The number of previous releases to retain on the channel for rollback or other
	// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.IntPtrInput
	// Required. The ID of the site in which to create this channel.
	SiteId pulumi.StringPtrInput
	// Input only. A time-to-live for this channel. Sets `expireTime` to the provided
	// duration past the time of the request. A duration in seconds with up to nine fractional
	// digits, terminated by 's'. Example: "86400s" (one day).
	Ttl pulumi.StringPtrInput
}

func (HostingChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostingChannelState)(nil)).Elem()
}

type hostingChannelArgs struct {
	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId string `pulumi:"channelId"`
	// The time at which the channel will be automatically deleted. If null, the channel
	// will not be automatically deleted. This field is present in the output whether it's
	// set directly or via the `ttl` field.
	ExpireTime *string `pulumi:"expireTime"`
	// Text labels used for extra metadata and/or filtering
	Labels map[string]string `pulumi:"labels"`
	// The number of previous releases to retain on the channel for rollback or other
	// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount *int `pulumi:"retainedReleaseCount"`
	// Required. The ID of the site in which to create this channel.
	SiteId string `pulumi:"siteId"`
	// Input only. A time-to-live for this channel. Sets `expireTime` to the provided
	// duration past the time of the request. A duration in seconds with up to nine fractional
	// digits, terminated by 's'. Example: "86400s" (one day).
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a HostingChannel resource.
type HostingChannelArgs struct {
	// Required. Immutable. A unique ID within the site that identifies the channel.
	ChannelId pulumi.StringInput
	// The time at which the channel will be automatically deleted. If null, the channel
	// will not be automatically deleted. This field is present in the output whether it's
	// set directly or via the `ttl` field.
	ExpireTime pulumi.StringPtrInput
	// Text labels used for extra metadata and/or filtering
	Labels pulumi.StringMapInput
	// The number of previous releases to retain on the channel for rollback or other
	// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	RetainedReleaseCount pulumi.IntPtrInput
	// Required. The ID of the site in which to create this channel.
	SiteId pulumi.StringInput
	// Input only. A time-to-live for this channel. Sets `expireTime` to the provided
	// duration past the time of the request. A duration in seconds with up to nine fractional
	// digits, terminated by 's'. Example: "86400s" (one day).
	Ttl pulumi.StringPtrInput
}

func (HostingChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostingChannelArgs)(nil)).Elem()
}

type HostingChannelInput interface {
	pulumi.Input

	ToHostingChannelOutput() HostingChannelOutput
	ToHostingChannelOutputWithContext(ctx context.Context) HostingChannelOutput
}

func (*HostingChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingChannel)(nil)).Elem()
}

func (i *HostingChannel) ToHostingChannelOutput() HostingChannelOutput {
	return i.ToHostingChannelOutputWithContext(context.Background())
}

func (i *HostingChannel) ToHostingChannelOutputWithContext(ctx context.Context) HostingChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingChannelOutput)
}

// HostingChannelArrayInput is an input type that accepts HostingChannelArray and HostingChannelArrayOutput values.
// You can construct a concrete instance of `HostingChannelArrayInput` via:
//
//	HostingChannelArray{ HostingChannelArgs{...} }
type HostingChannelArrayInput interface {
	pulumi.Input

	ToHostingChannelArrayOutput() HostingChannelArrayOutput
	ToHostingChannelArrayOutputWithContext(context.Context) HostingChannelArrayOutput
}

type HostingChannelArray []HostingChannelInput

func (HostingChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostingChannel)(nil)).Elem()
}

func (i HostingChannelArray) ToHostingChannelArrayOutput() HostingChannelArrayOutput {
	return i.ToHostingChannelArrayOutputWithContext(context.Background())
}

func (i HostingChannelArray) ToHostingChannelArrayOutputWithContext(ctx context.Context) HostingChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingChannelArrayOutput)
}

// HostingChannelMapInput is an input type that accepts HostingChannelMap and HostingChannelMapOutput values.
// You can construct a concrete instance of `HostingChannelMapInput` via:
//
//	HostingChannelMap{ "key": HostingChannelArgs{...} }
type HostingChannelMapInput interface {
	pulumi.Input

	ToHostingChannelMapOutput() HostingChannelMapOutput
	ToHostingChannelMapOutputWithContext(context.Context) HostingChannelMapOutput
}

type HostingChannelMap map[string]HostingChannelInput

func (HostingChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostingChannel)(nil)).Elem()
}

func (i HostingChannelMap) ToHostingChannelMapOutput() HostingChannelMapOutput {
	return i.ToHostingChannelMapOutputWithContext(context.Background())
}

func (i HostingChannelMap) ToHostingChannelMapOutputWithContext(ctx context.Context) HostingChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingChannelMapOutput)
}

type HostingChannelOutput struct{ *pulumi.OutputState }

func (HostingChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingChannel)(nil)).Elem()
}

func (o HostingChannelOutput) ToHostingChannelOutput() HostingChannelOutput {
	return o
}

func (o HostingChannelOutput) ToHostingChannelOutputWithContext(ctx context.Context) HostingChannelOutput {
	return o
}

// Required. Immutable. A unique ID within the site that identifies the channel.
func (o HostingChannelOutput) ChannelId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostingChannel) pulumi.StringOutput { return v.ChannelId }).(pulumi.StringOutput)
}

// The time at which the channel will be automatically deleted. If null, the channel
// will not be automatically deleted. This field is present in the output whether it's
// set directly or via the `ttl` field.
func (o HostingChannelOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *HostingChannel) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// Text labels used for extra metadata and/or filtering
func (o HostingChannelOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HostingChannel) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The fully-qualified resource name for the channel, in the format:
// sites/SITE_ID/channels/CHANNEL_ID
func (o HostingChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HostingChannel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of previous releases to retain on the channel for rollback or other
// purposes. Must be a number between 1-100. Defaults to 10 for new channels.
func (o HostingChannelOutput) RetainedReleaseCount() pulumi.IntOutput {
	return o.ApplyT(func(v *HostingChannel) pulumi.IntOutput { return v.RetainedReleaseCount }).(pulumi.IntOutput)
}

// Required. The ID of the site in which to create this channel.
func (o HostingChannelOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostingChannel) pulumi.StringOutput { return v.SiteId }).(pulumi.StringOutput)
}

// Input only. A time-to-live for this channel. Sets `expireTime` to the provided
// duration past the time of the request. A duration in seconds with up to nine fractional
// digits, terminated by 's'. Example: "86400s" (one day).
func (o HostingChannelOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingChannel) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

type HostingChannelArrayOutput struct{ *pulumi.OutputState }

func (HostingChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostingChannel)(nil)).Elem()
}

func (o HostingChannelArrayOutput) ToHostingChannelArrayOutput() HostingChannelArrayOutput {
	return o
}

func (o HostingChannelArrayOutput) ToHostingChannelArrayOutputWithContext(ctx context.Context) HostingChannelArrayOutput {
	return o
}

func (o HostingChannelArrayOutput) Index(i pulumi.IntInput) HostingChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostingChannel {
		return vs[0].([]*HostingChannel)[vs[1].(int)]
	}).(HostingChannelOutput)
}

type HostingChannelMapOutput struct{ *pulumi.OutputState }

func (HostingChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostingChannel)(nil)).Elem()
}

func (o HostingChannelMapOutput) ToHostingChannelMapOutput() HostingChannelMapOutput {
	return o
}

func (o HostingChannelMapOutput) ToHostingChannelMapOutputWithContext(ctx context.Context) HostingChannelMapOutput {
	return o
}

func (o HostingChannelMapOutput) MapIndex(k pulumi.StringInput) HostingChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostingChannel {
		return vs[0].(map[string]*HostingChannel)[vs[1].(string)]
	}).(HostingChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostingChannelInput)(nil)).Elem(), &HostingChannel{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostingChannelArrayInput)(nil)).Elem(), HostingChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostingChannelMapInput)(nil)).Elem(), HostingChannelMap{})
	pulumi.RegisterOutputType(HostingChannelOutput{})
	pulumi.RegisterOutputType(HostingChannelArrayOutput{})
	pulumi.RegisterOutputType(HostingChannelMapOutput{})
}
