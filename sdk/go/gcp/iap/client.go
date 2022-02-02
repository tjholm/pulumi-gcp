// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains the data that describes an Identity Aware Proxy owned client.
//
// > **Note:** Only internal org clients can be created via declarative tools. External clients must be
// manually created via the GCP console. This restriction is due to the existing APIs and not lack of support
// in this tool.
//
// To get more information about Client, see:
//
// * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.brands.identityAwareProxyClients)
// * How-to Guides
//     * [Setting up IAP Client](https://cloud.google.com/iap/docs/authentication-howto)
//
// > **Warning:** All arguments including `secret` will be stored in the raw
// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
//
// ## Example Usage
//
// ## Import
//
// Client can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:iap/client:Client default {{brand}}/identityAwareProxyClients/{{client_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:iap/client:Client default {{brand}}/{{client_id}}
// ```
type Client struct {
	pulumi.CustomResourceState

	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand pulumi.StringOutput `pulumi:"brand"`
	// Output only. Unique identifier of the OAuth client.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Output only. Client secret of the OAuth client.
	Secret pulumi.StringOutput `pulumi:"secret"`
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOption) (*Client, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Brand == nil {
		return nil, errors.New("invalid value for required argument 'Brand'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Client
	err := ctx.RegisterResource("gcp:iap/client:Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientState, opts ...pulumi.ResourceOption) (*Client, error) {
	var resource Client
	err := ctx.ReadResource("gcp:iap/client:Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Client resources.
type clientState struct {
	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand *string `pulumi:"brand"`
	// Output only. Unique identifier of the OAuth client.
	ClientId *string `pulumi:"clientId"`
	// Human-friendly name given to the OAuth client.
	DisplayName *string `pulumi:"displayName"`
	// Output only. Client secret of the OAuth client.
	Secret *string `pulumi:"secret"`
}

type ClientState struct {
	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand pulumi.StringPtrInput
	// Output only. Unique identifier of the OAuth client.
	ClientId pulumi.StringPtrInput
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringPtrInput
	// Output only. Client secret of the OAuth client.
	Secret pulumi.StringPtrInput
}

func (ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientState)(nil)).Elem()
}

type clientArgs struct {
	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand string `pulumi:"brand"`
	// Human-friendly name given to the OAuth client.
	DisplayName string `pulumi:"displayName"`
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	// Identifier of the brand to which this client
	// is attached to. The format is
	// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
	Brand pulumi.StringInput
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringInput
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientArgs)(nil)).Elem()
}

type ClientInput interface {
	pulumi.Input

	ToClientOutput() ClientOutput
	ToClientOutputWithContext(ctx context.Context) ClientOutput
}

func (*Client) ElementType() reflect.Type {
	return reflect.TypeOf((**Client)(nil)).Elem()
}

func (i *Client) ToClientOutput() ClientOutput {
	return i.ToClientOutputWithContext(context.Background())
}

func (i *Client) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOutput)
}

// ClientArrayInput is an input type that accepts ClientArray and ClientArrayOutput values.
// You can construct a concrete instance of `ClientArrayInput` via:
//
//          ClientArray{ ClientArgs{...} }
type ClientArrayInput interface {
	pulumi.Input

	ToClientArrayOutput() ClientArrayOutput
	ToClientArrayOutputWithContext(context.Context) ClientArrayOutput
}

type ClientArray []ClientInput

func (ClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Client)(nil)).Elem()
}

func (i ClientArray) ToClientArrayOutput() ClientArrayOutput {
	return i.ToClientArrayOutputWithContext(context.Background())
}

func (i ClientArray) ToClientArrayOutputWithContext(ctx context.Context) ClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientArrayOutput)
}

// ClientMapInput is an input type that accepts ClientMap and ClientMapOutput values.
// You can construct a concrete instance of `ClientMapInput` via:
//
//          ClientMap{ "key": ClientArgs{...} }
type ClientMapInput interface {
	pulumi.Input

	ToClientMapOutput() ClientMapOutput
	ToClientMapOutputWithContext(context.Context) ClientMapOutput
}

type ClientMap map[string]ClientInput

func (ClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Client)(nil)).Elem()
}

func (i ClientMap) ToClientMapOutput() ClientMapOutput {
	return i.ToClientMapOutputWithContext(context.Background())
}

func (i ClientMap) ToClientMapOutputWithContext(ctx context.Context) ClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientMapOutput)
}

type ClientOutput struct{ *pulumi.OutputState }

func (ClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Client)(nil)).Elem()
}

func (o ClientOutput) ToClientOutput() ClientOutput {
	return o
}

func (o ClientOutput) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return o
}

type ClientArrayOutput struct{ *pulumi.OutputState }

func (ClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Client)(nil)).Elem()
}

func (o ClientArrayOutput) ToClientArrayOutput() ClientArrayOutput {
	return o
}

func (o ClientArrayOutput) ToClientArrayOutputWithContext(ctx context.Context) ClientArrayOutput {
	return o
}

func (o ClientArrayOutput) Index(i pulumi.IntInput) ClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Client {
		return vs[0].([]*Client)[vs[1].(int)]
	}).(ClientOutput)
}

type ClientMapOutput struct{ *pulumi.OutputState }

func (ClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Client)(nil)).Elem()
}

func (o ClientMapOutput) ToClientMapOutput() ClientMapOutput {
	return o
}

func (o ClientMapOutput) ToClientMapOutputWithContext(ctx context.Context) ClientMapOutput {
	return o
}

func (o ClientMapOutput) MapIndex(k pulumi.StringInput) ClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Client {
		return vs[0].(map[string]*Client)[vs[1].(string)]
	}).(ClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientInput)(nil)).Elem(), &Client{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientArrayInput)(nil)).Elem(), ClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientMapInput)(nil)).Elem(), ClientMap{})
	pulumi.RegisterOutputType(ClientOutput{})
	pulumi.RegisterOutputType(ClientArrayOutput{})
	pulumi.RegisterOutputType(ClientMapOutput{})
}
