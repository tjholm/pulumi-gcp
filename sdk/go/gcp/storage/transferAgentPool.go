// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an On-Premises Agent pool.
//
// To get more information about AgentPool, see:
//
// * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/projects.agentPools)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/storage-transfer/docs/on-prem-agent-pools)
//
// ## Example Usage
//
// ### Agent Pool Basic
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := storage.GetTransferProjectServiceAccount(ctx, &storage.GetTransferProjectServiceAccountArgs{
//				Project: pulumi.StringRef("my-project-name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			pubsubEditorRole, err := projects.NewIAMMember(ctx, "pubsub_editor_role", &projects.IAMMemberArgs{
//				Project: pulumi.String("my-project-name"),
//				Role:    pulumi.String("roles/pubsub.editor"),
//				Member:  pulumi.Sprintf("serviceAccount:%v", _default.Email),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewTransferAgentPool(ctx, "example", &storage.TransferAgentPoolArgs{
//				Name:        pulumi.String("agent-pool-example"),
//				DisplayName: pulumi.String("Source A to destination Z"),
//				BandwidthLimit: &storage.TransferAgentPoolBandwidthLimitArgs{
//					LimitMbps: pulumi.String("120"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				pubsubEditorRole,
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
// AgentPool can be imported using any of these accepted formats:
//
// * `projects/{{project}}/agentPools/{{name}}`
//
// * `{{project}}/{{name}}`
//
// * `{{name}}`
//
// When using the `pulumi import` command, AgentPool can be imported using one of the formats above. For example:
//
// ```sh
// $ pulumi import gcp:storage/transferAgentPool:TransferAgentPool default projects/{{project}}/agentPools/{{name}}
// ```
//
// ```sh
// $ pulumi import gcp:storage/transferAgentPool:TransferAgentPool default {{project}}/{{name}}
// ```
//
// ```sh
// $ pulumi import gcp:storage/transferAgentPool:TransferAgentPool default {{name}}
// ```
type TransferAgentPool struct {
	pulumi.CustomResourceState

	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	BandwidthLimit TransferAgentPoolBandwidthLimitPtrOutput `pulumi:"bandwidthLimit"`
	// Specifies the client-specified AgentPool description.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The ID of the agent pool to create.
	// The agentPoolId must meet the following requirements:
	// * Length of 128 characters or less.
	// * Not start with the string goog.
	// * Start with a lowercase ASCII character, followed by:
	// * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
	// * One or more numerals or lowercase ASCII characters.
	//   As expressed by the regular expression: ^(?!goog)a-z?$.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies the state of the AgentPool.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewTransferAgentPool registers a new resource with the given unique name, arguments, and options.
func NewTransferAgentPool(ctx *pulumi.Context,
	name string, args *TransferAgentPoolArgs, opts ...pulumi.ResourceOption) (*TransferAgentPool, error) {
	if args == nil {
		args = &TransferAgentPoolArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransferAgentPool
	err := ctx.RegisterResource("gcp:storage/transferAgentPool:TransferAgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransferAgentPool gets an existing TransferAgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransferAgentPoolState, opts ...pulumi.ResourceOption) (*TransferAgentPool, error) {
	var resource TransferAgentPool
	err := ctx.ReadResource("gcp:storage/transferAgentPool:TransferAgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransferAgentPool resources.
type transferAgentPoolState struct {
	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	BandwidthLimit *TransferAgentPoolBandwidthLimit `pulumi:"bandwidthLimit"`
	// Specifies the client-specified AgentPool description.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the agent pool to create.
	// The agentPoolId must meet the following requirements:
	// * Length of 128 characters or less.
	// * Not start with the string goog.
	// * Start with a lowercase ASCII character, followed by:
	// * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
	// * One or more numerals or lowercase ASCII characters.
	//   As expressed by the regular expression: ^(?!goog)a-z?$.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Specifies the state of the AgentPool.
	State *string `pulumi:"state"`
}

type TransferAgentPoolState struct {
	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	BandwidthLimit TransferAgentPoolBandwidthLimitPtrInput
	// Specifies the client-specified AgentPool description.
	DisplayName pulumi.StringPtrInput
	// The ID of the agent pool to create.
	// The agentPoolId must meet the following requirements:
	// * Length of 128 characters or less.
	// * Not start with the string goog.
	// * Start with a lowercase ASCII character, followed by:
	// * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
	// * One or more numerals or lowercase ASCII characters.
	//   As expressed by the regular expression: ^(?!goog)a-z?$.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Specifies the state of the AgentPool.
	State pulumi.StringPtrInput
}

func (TransferAgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*transferAgentPoolState)(nil)).Elem()
}

type transferAgentPoolArgs struct {
	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	BandwidthLimit *TransferAgentPoolBandwidthLimit `pulumi:"bandwidthLimit"`
	// Specifies the client-specified AgentPool description.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the agent pool to create.
	// The agentPoolId must meet the following requirements:
	// * Length of 128 characters or less.
	// * Not start with the string goog.
	// * Start with a lowercase ASCII character, followed by:
	// * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
	// * One or more numerals or lowercase ASCII characters.
	//   As expressed by the regular expression: ^(?!goog)a-z?$.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a TransferAgentPool resource.
type TransferAgentPoolArgs struct {
	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	// Structure is documented below.
	BandwidthLimit TransferAgentPoolBandwidthLimitPtrInput
	// Specifies the client-specified AgentPool description.
	DisplayName pulumi.StringPtrInput
	// The ID of the agent pool to create.
	// The agentPoolId must meet the following requirements:
	// * Length of 128 characters or less.
	// * Not start with the string goog.
	// * Start with a lowercase ASCII character, followed by:
	// * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
	// * One or more numerals or lowercase ASCII characters.
	//   As expressed by the regular expression: ^(?!goog)a-z?$.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TransferAgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transferAgentPoolArgs)(nil)).Elem()
}

type TransferAgentPoolInput interface {
	pulumi.Input

	ToTransferAgentPoolOutput() TransferAgentPoolOutput
	ToTransferAgentPoolOutputWithContext(ctx context.Context) TransferAgentPoolOutput
}

func (*TransferAgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferAgentPool)(nil)).Elem()
}

func (i *TransferAgentPool) ToTransferAgentPoolOutput() TransferAgentPoolOutput {
	return i.ToTransferAgentPoolOutputWithContext(context.Background())
}

func (i *TransferAgentPool) ToTransferAgentPoolOutputWithContext(ctx context.Context) TransferAgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAgentPoolOutput)
}

// TransferAgentPoolArrayInput is an input type that accepts TransferAgentPoolArray and TransferAgentPoolArrayOutput values.
// You can construct a concrete instance of `TransferAgentPoolArrayInput` via:
//
//	TransferAgentPoolArray{ TransferAgentPoolArgs{...} }
type TransferAgentPoolArrayInput interface {
	pulumi.Input

	ToTransferAgentPoolArrayOutput() TransferAgentPoolArrayOutput
	ToTransferAgentPoolArrayOutputWithContext(context.Context) TransferAgentPoolArrayOutput
}

type TransferAgentPoolArray []TransferAgentPoolInput

func (TransferAgentPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransferAgentPool)(nil)).Elem()
}

func (i TransferAgentPoolArray) ToTransferAgentPoolArrayOutput() TransferAgentPoolArrayOutput {
	return i.ToTransferAgentPoolArrayOutputWithContext(context.Background())
}

func (i TransferAgentPoolArray) ToTransferAgentPoolArrayOutputWithContext(ctx context.Context) TransferAgentPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAgentPoolArrayOutput)
}

// TransferAgentPoolMapInput is an input type that accepts TransferAgentPoolMap and TransferAgentPoolMapOutput values.
// You can construct a concrete instance of `TransferAgentPoolMapInput` via:
//
//	TransferAgentPoolMap{ "key": TransferAgentPoolArgs{...} }
type TransferAgentPoolMapInput interface {
	pulumi.Input

	ToTransferAgentPoolMapOutput() TransferAgentPoolMapOutput
	ToTransferAgentPoolMapOutputWithContext(context.Context) TransferAgentPoolMapOutput
}

type TransferAgentPoolMap map[string]TransferAgentPoolInput

func (TransferAgentPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransferAgentPool)(nil)).Elem()
}

func (i TransferAgentPoolMap) ToTransferAgentPoolMapOutput() TransferAgentPoolMapOutput {
	return i.ToTransferAgentPoolMapOutputWithContext(context.Background())
}

func (i TransferAgentPoolMap) ToTransferAgentPoolMapOutputWithContext(ctx context.Context) TransferAgentPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAgentPoolMapOutput)
}

type TransferAgentPoolOutput struct{ *pulumi.OutputState }

func (TransferAgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferAgentPool)(nil)).Elem()
}

func (o TransferAgentPoolOutput) ToTransferAgentPoolOutput() TransferAgentPoolOutput {
	return o
}

func (o TransferAgentPoolOutput) ToTransferAgentPoolOutputWithContext(ctx context.Context) TransferAgentPoolOutput {
	return o
}

// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
// Structure is documented below.
func (o TransferAgentPoolOutput) BandwidthLimit() TransferAgentPoolBandwidthLimitPtrOutput {
	return o.ApplyT(func(v *TransferAgentPool) TransferAgentPoolBandwidthLimitPtrOutput { return v.BandwidthLimit }).(TransferAgentPoolBandwidthLimitPtrOutput)
}

// Specifies the client-specified AgentPool description.
func (o TransferAgentPoolOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferAgentPool) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of the agent pool to create.
// The agentPoolId must meet the following requirements:
//   - Length of 128 characters or less.
//   - Not start with the string goog.
//   - Start with a lowercase ASCII character, followed by:
//   - Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
//   - One or more numerals or lowercase ASCII characters.
//     As expressed by the regular expression: ^(?!goog)a-z?$.
//
// ***
func (o TransferAgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferAgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o TransferAgentPoolOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferAgentPool) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the state of the AgentPool.
func (o TransferAgentPoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferAgentPool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type TransferAgentPoolArrayOutput struct{ *pulumi.OutputState }

func (TransferAgentPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransferAgentPool)(nil)).Elem()
}

func (o TransferAgentPoolArrayOutput) ToTransferAgentPoolArrayOutput() TransferAgentPoolArrayOutput {
	return o
}

func (o TransferAgentPoolArrayOutput) ToTransferAgentPoolArrayOutputWithContext(ctx context.Context) TransferAgentPoolArrayOutput {
	return o
}

func (o TransferAgentPoolArrayOutput) Index(i pulumi.IntInput) TransferAgentPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransferAgentPool {
		return vs[0].([]*TransferAgentPool)[vs[1].(int)]
	}).(TransferAgentPoolOutput)
}

type TransferAgentPoolMapOutput struct{ *pulumi.OutputState }

func (TransferAgentPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransferAgentPool)(nil)).Elem()
}

func (o TransferAgentPoolMapOutput) ToTransferAgentPoolMapOutput() TransferAgentPoolMapOutput {
	return o
}

func (o TransferAgentPoolMapOutput) ToTransferAgentPoolMapOutputWithContext(ctx context.Context) TransferAgentPoolMapOutput {
	return o
}

func (o TransferAgentPoolMapOutput) MapIndex(k pulumi.StringInput) TransferAgentPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransferAgentPool {
		return vs[0].(map[string]*TransferAgentPool)[vs[1].(string)]
	}).(TransferAgentPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransferAgentPoolInput)(nil)).Elem(), &TransferAgentPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransferAgentPoolArrayInput)(nil)).Elem(), TransferAgentPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransferAgentPoolMapInput)(nil)).Elem(), TransferAgentPoolMap{})
	pulumi.RegisterOutputType(TransferAgentPoolOutput{})
	pulumi.RegisterOutputType(TransferAgentPoolArrayOutput{})
	pulumi.RegisterOutputType(TransferAgentPoolMapOutput{})
}
