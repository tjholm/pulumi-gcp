// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AzureClient resources hold client authentication information needed by the Anthos Multi-Cloud API to manage Azure resources on your Azure subscription.When an AzureCluster is created, an AzureClient resource needs to be provided and all operations on Azure resources associated to that cluster will authenticate to Azure services using the given client.AzureClient resources are immutable and cannot be modified upon creation.Each AzureClient resource is bound to a single Azure Active Directory Application and tenant.
//
// For more information, see:
// * [Multicloud overview](https://cloud.google.com/anthos/clusters/docs/multi-cloud)
// ## Example Usage
// ### Basic_azure_client
// A basic example of a containerazure azure client
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := container.NewAzureClient(ctx, "primary", &container.AzureClientArgs{
//				ApplicationId: pulumi.String("12345678-1234-1234-1234-123456789111"),
//				Location:      pulumi.String("us-west1"),
//				Project:       pulumi.String("my-project-name"),
//				TenantId:      pulumi.String("12345678-1234-1234-1234-123456789111"),
//			})
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
// Client can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/azureClients/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Client can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:container/azureClient:AzureClient default projects/{{project}}/locations/{{location}}/azureClients/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:container/azureClient:AzureClient default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:container/azureClient:AzureClient default {{location}}/{{name}}
//
// ```
type AzureClient struct {
	pulumi.CustomResourceState

	// The Azure Active Directory Application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Output only. The PEM encoded x509 certificate.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Output only. The time at which this resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The Azure Active Directory Tenant ID.
	//
	// ***
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Output only. A globally unique identifier for the client.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewAzureClient registers a new resource with the given unique name, arguments, and options.
func NewAzureClient(ctx *pulumi.Context,
	name string, args *AzureClientArgs, opts ...pulumi.ResourceOption) (*AzureClient, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AzureClient
	err := ctx.RegisterResource("gcp:container/azureClient:AzureClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureClient gets an existing AzureClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureClientState, opts ...pulumi.ResourceOption) (*AzureClient, error) {
	var resource AzureClient
	err := ctx.ReadResource("gcp:container/azureClient:AzureClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureClient resources.
type azureClientState struct {
	// The Azure Active Directory Application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Output only. The PEM encoded x509 certificate.
	Certificate *string `pulumi:"certificate"`
	// Output only. The time at which this resource was created.
	CreateTime *string `pulumi:"createTime"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The Azure Active Directory Tenant ID.
	//
	// ***
	TenantId *string `pulumi:"tenantId"`
	// Output only. A globally unique identifier for the client.
	Uid *string `pulumi:"uid"`
}

type AzureClientState struct {
	// The Azure Active Directory Application ID.
	ApplicationId pulumi.StringPtrInput
	// Output only. The PEM encoded x509 certificate.
	Certificate pulumi.StringPtrInput
	// Output only. The time at which this resource was created.
	CreateTime pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The Azure Active Directory Tenant ID.
	//
	// ***
	TenantId pulumi.StringPtrInput
	// Output only. A globally unique identifier for the client.
	Uid pulumi.StringPtrInput
}

func (AzureClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureClientState)(nil)).Elem()
}

type azureClientArgs struct {
	// The Azure Active Directory Application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The Azure Active Directory Tenant ID.
	//
	// ***
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a AzureClient resource.
type AzureClientArgs struct {
	// The Azure Active Directory Application ID.
	ApplicationId pulumi.StringInput
	// The location for the resource
	Location pulumi.StringInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The Azure Active Directory Tenant ID.
	//
	// ***
	TenantId pulumi.StringInput
}

func (AzureClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureClientArgs)(nil)).Elem()
}

type AzureClientInput interface {
	pulumi.Input

	ToAzureClientOutput() AzureClientOutput
	ToAzureClientOutputWithContext(ctx context.Context) AzureClientOutput
}

func (*AzureClient) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureClient)(nil)).Elem()
}

func (i *AzureClient) ToAzureClientOutput() AzureClientOutput {
	return i.ToAzureClientOutputWithContext(context.Background())
}

func (i *AzureClient) ToAzureClientOutputWithContext(ctx context.Context) AzureClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureClientOutput)
}

// AzureClientArrayInput is an input type that accepts AzureClientArray and AzureClientArrayOutput values.
// You can construct a concrete instance of `AzureClientArrayInput` via:
//
//	AzureClientArray{ AzureClientArgs{...} }
type AzureClientArrayInput interface {
	pulumi.Input

	ToAzureClientArrayOutput() AzureClientArrayOutput
	ToAzureClientArrayOutputWithContext(context.Context) AzureClientArrayOutput
}

type AzureClientArray []AzureClientInput

func (AzureClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureClient)(nil)).Elem()
}

func (i AzureClientArray) ToAzureClientArrayOutput() AzureClientArrayOutput {
	return i.ToAzureClientArrayOutputWithContext(context.Background())
}

func (i AzureClientArray) ToAzureClientArrayOutputWithContext(ctx context.Context) AzureClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureClientArrayOutput)
}

// AzureClientMapInput is an input type that accepts AzureClientMap and AzureClientMapOutput values.
// You can construct a concrete instance of `AzureClientMapInput` via:
//
//	AzureClientMap{ "key": AzureClientArgs{...} }
type AzureClientMapInput interface {
	pulumi.Input

	ToAzureClientMapOutput() AzureClientMapOutput
	ToAzureClientMapOutputWithContext(context.Context) AzureClientMapOutput
}

type AzureClientMap map[string]AzureClientInput

func (AzureClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureClient)(nil)).Elem()
}

func (i AzureClientMap) ToAzureClientMapOutput() AzureClientMapOutput {
	return i.ToAzureClientMapOutputWithContext(context.Background())
}

func (i AzureClientMap) ToAzureClientMapOutputWithContext(ctx context.Context) AzureClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureClientMapOutput)
}

type AzureClientOutput struct{ *pulumi.OutputState }

func (AzureClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureClient)(nil)).Elem()
}

func (o AzureClientOutput) ToAzureClientOutput() AzureClientOutput {
	return o
}

func (o AzureClientOutput) ToAzureClientOutputWithContext(ctx context.Context) AzureClientOutput {
	return o
}

// The Azure Active Directory Application ID.
func (o AzureClientOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureClient) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// Output only. The PEM encoded x509 certificate.
func (o AzureClientOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureClient) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Output only. The time at which this resource was created.
func (o AzureClientOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureClient) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The location for the resource
func (o AzureClientOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureClient) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of this resource.
func (o AzureClientOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureClient) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o AzureClientOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureClient) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The Azure Active Directory Tenant ID.
//
// ***
func (o AzureClientOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureClient) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Output only. A globally unique identifier for the client.
func (o AzureClientOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureClient) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

type AzureClientArrayOutput struct{ *pulumi.OutputState }

func (AzureClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureClient)(nil)).Elem()
}

func (o AzureClientArrayOutput) ToAzureClientArrayOutput() AzureClientArrayOutput {
	return o
}

func (o AzureClientArrayOutput) ToAzureClientArrayOutputWithContext(ctx context.Context) AzureClientArrayOutput {
	return o
}

func (o AzureClientArrayOutput) Index(i pulumi.IntInput) AzureClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureClient {
		return vs[0].([]*AzureClient)[vs[1].(int)]
	}).(AzureClientOutput)
}

type AzureClientMapOutput struct{ *pulumi.OutputState }

func (AzureClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureClient)(nil)).Elem()
}

func (o AzureClientMapOutput) ToAzureClientMapOutput() AzureClientMapOutput {
	return o
}

func (o AzureClientMapOutput) ToAzureClientMapOutputWithContext(ctx context.Context) AzureClientMapOutput {
	return o
}

func (o AzureClientMapOutput) MapIndex(k pulumi.StringInput) AzureClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureClient {
		return vs[0].(map[string]*AzureClient)[vs[1].(string)]
	}).(AzureClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureClientInput)(nil)).Elem(), &AzureClient{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureClientArrayInput)(nil)).Elem(), AzureClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureClientMapInput)(nil)).Elem(), AzureClientMap{})
	pulumi.RegisterOutputType(AzureClientOutput{})
	pulumi.RegisterOutputType(AzureClientArrayOutput{})
	pulumi.RegisterOutputType(AzureClientMapOutput{})
}
