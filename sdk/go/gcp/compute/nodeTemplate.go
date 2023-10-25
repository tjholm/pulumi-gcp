// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a NodeTemplate resource. Node templates specify properties
// for creating sole-tenant nodes, such as node type, vCPU and memory
// requirements, node affinity labels, and region.
//
// To get more information about NodeTemplate, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTemplates)
// * How-to Guides
//   - [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
//
// ## Example Usage
// ### Node Template Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewNodeTemplate(ctx, "template", &compute.NodeTemplateArgs{
//				NodeType: pulumi.String("n1-node-96-624"),
//				Region:   pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Node Template Server Binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.GetNodeTypes(ctx, &compute.GetNodeTypesArgs{
//				Zone: pulumi.StringRef("us-central1-a"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNodeTemplate(ctx, "template", &compute.NodeTemplateArgs{
//				NodeAffinityLabels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				NodeType: pulumi.String("n1-node-96-624"),
//				Region:   pulumi.String("us-central1"),
//				ServerBinding: &compute.NodeTemplateServerBindingArgs{
//					Type: pulumi.String("RESTART_NODE_ON_MINIMAL_SERVERS"),
//				},
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
// # NodeTemplate can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/nodeTemplate:NodeTemplate default projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{project}}/{{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{name}}
//
// ```
type NodeTemplate struct {
	pulumi.CustomResourceState

	// CPU overcommit.
	// Default value is `NONE`.
	// Possible values are: `ENABLED`, `NONE`.
	CpuOvercommitType pulumi.StringPtrOutput `pulumi:"cpuOvercommitType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional textual description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Labels to use for node affinity, which will be used in
	// instance scheduling.
	NodeAffinityLabels pulumi.StringMapOutput `pulumi:"nodeAffinityLabels"`
	// Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeType pulumi.StringPtrOutput `pulumi:"nodeType"`
	// Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	// Structure is documented below.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityPtrOutput `pulumi:"nodeTypeFlexibility"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where nodes using the node template will be created.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	// Structure is documented below.
	ServerBinding NodeTemplateServerBindingOutput `pulumi:"serverBinding"`
}

// NewNodeTemplate registers a new resource with the given unique name, arguments, and options.
func NewNodeTemplate(ctx *pulumi.Context,
	name string, args *NodeTemplateArgs, opts ...pulumi.ResourceOption) (*NodeTemplate, error) {
	if args == nil {
		args = &NodeTemplateArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NodeTemplate
	err := ctx.RegisterResource("gcp:compute/nodeTemplate:NodeTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeTemplate gets an existing NodeTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTemplateState, opts ...pulumi.ResourceOption) (*NodeTemplate, error) {
	var resource NodeTemplate
	err := ctx.ReadResource("gcp:compute/nodeTemplate:NodeTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeTemplate resources.
type nodeTemplateState struct {
	// CPU overcommit.
	// Default value is `NONE`.
	// Possible values are: `ENABLED`, `NONE`.
	CpuOvercommitType *string `pulumi:"cpuOvercommitType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional textual description of the resource.
	Description *string `pulumi:"description"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Labels to use for node affinity, which will be used in
	// instance scheduling.
	NodeAffinityLabels map[string]string `pulumi:"nodeAffinityLabels"`
	// Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeType *string `pulumi:"nodeType"`
	// Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	// Structure is documented below.
	NodeTypeFlexibility *NodeTemplateNodeTypeFlexibility `pulumi:"nodeTypeFlexibility"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where nodes using the node template will be created.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	// Structure is documented below.
	ServerBinding *NodeTemplateServerBinding `pulumi:"serverBinding"`
}

type NodeTemplateState struct {
	// CPU overcommit.
	// Default value is `NONE`.
	// Possible values are: `ENABLED`, `NONE`.
	CpuOvercommitType pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional textual description of the resource.
	Description pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Labels to use for node affinity, which will be used in
	// instance scheduling.
	NodeAffinityLabels pulumi.StringMapInput
	// Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeType pulumi.StringPtrInput
	// Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	// Structure is documented below.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where nodes using the node template will be created.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	// Structure is documented below.
	ServerBinding NodeTemplateServerBindingPtrInput
}

func (NodeTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateState)(nil)).Elem()
}

type nodeTemplateArgs struct {
	// CPU overcommit.
	// Default value is `NONE`.
	// Possible values are: `ENABLED`, `NONE`.
	CpuOvercommitType *string `pulumi:"cpuOvercommitType"`
	// An optional textual description of the resource.
	Description *string `pulumi:"description"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Labels to use for node affinity, which will be used in
	// instance scheduling.
	NodeAffinityLabels map[string]string `pulumi:"nodeAffinityLabels"`
	// Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeType *string `pulumi:"nodeType"`
	// Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	// Structure is documented below.
	NodeTypeFlexibility *NodeTemplateNodeTypeFlexibility `pulumi:"nodeTypeFlexibility"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where nodes using the node template will be created.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	// Structure is documented below.
	ServerBinding *NodeTemplateServerBinding `pulumi:"serverBinding"`
}

// The set of arguments for constructing a NodeTemplate resource.
type NodeTemplateArgs struct {
	// CPU overcommit.
	// Default value is `NONE`.
	// Possible values are: `ENABLED`, `NONE`.
	CpuOvercommitType pulumi.StringPtrInput
	// An optional textual description of the resource.
	Description pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Labels to use for node affinity, which will be used in
	// instance scheduling.
	NodeAffinityLabels pulumi.StringMapInput
	// Node type to use for nodes group that are created from this template.
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeType pulumi.StringPtrInput
	// Flexible properties for the desired node type. Node groups that
	// use this node template will create nodes of a type that matches
	// these properties. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	// Structure is documented below.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where nodes using the node template will be created.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The server binding policy for nodes using this template. Determines
	// where the nodes should restart following a maintenance event.
	// Structure is documented below.
	ServerBinding NodeTemplateServerBindingPtrInput
}

func (NodeTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateArgs)(nil)).Elem()
}

type NodeTemplateInput interface {
	pulumi.Input

	ToNodeTemplateOutput() NodeTemplateOutput
	ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput
}

func (*NodeTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeTemplate)(nil)).Elem()
}

func (i *NodeTemplate) ToNodeTemplateOutput() NodeTemplateOutput {
	return i.ToNodeTemplateOutputWithContext(context.Background())
}

func (i *NodeTemplate) ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTemplateOutput)
}

func (i *NodeTemplate) ToOutput(ctx context.Context) pulumix.Output[*NodeTemplate] {
	return pulumix.Output[*NodeTemplate]{
		OutputState: i.ToNodeTemplateOutputWithContext(ctx).OutputState,
	}
}

// NodeTemplateArrayInput is an input type that accepts NodeTemplateArray and NodeTemplateArrayOutput values.
// You can construct a concrete instance of `NodeTemplateArrayInput` via:
//
//	NodeTemplateArray{ NodeTemplateArgs{...} }
type NodeTemplateArrayInput interface {
	pulumi.Input

	ToNodeTemplateArrayOutput() NodeTemplateArrayOutput
	ToNodeTemplateArrayOutputWithContext(context.Context) NodeTemplateArrayOutput
}

type NodeTemplateArray []NodeTemplateInput

func (NodeTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeTemplate)(nil)).Elem()
}

func (i NodeTemplateArray) ToNodeTemplateArrayOutput() NodeTemplateArrayOutput {
	return i.ToNodeTemplateArrayOutputWithContext(context.Background())
}

func (i NodeTemplateArray) ToNodeTemplateArrayOutputWithContext(ctx context.Context) NodeTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTemplateArrayOutput)
}

func (i NodeTemplateArray) ToOutput(ctx context.Context) pulumix.Output[[]*NodeTemplate] {
	return pulumix.Output[[]*NodeTemplate]{
		OutputState: i.ToNodeTemplateArrayOutputWithContext(ctx).OutputState,
	}
}

// NodeTemplateMapInput is an input type that accepts NodeTemplateMap and NodeTemplateMapOutput values.
// You can construct a concrete instance of `NodeTemplateMapInput` via:
//
//	NodeTemplateMap{ "key": NodeTemplateArgs{...} }
type NodeTemplateMapInput interface {
	pulumi.Input

	ToNodeTemplateMapOutput() NodeTemplateMapOutput
	ToNodeTemplateMapOutputWithContext(context.Context) NodeTemplateMapOutput
}

type NodeTemplateMap map[string]NodeTemplateInput

func (NodeTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeTemplate)(nil)).Elem()
}

func (i NodeTemplateMap) ToNodeTemplateMapOutput() NodeTemplateMapOutput {
	return i.ToNodeTemplateMapOutputWithContext(context.Background())
}

func (i NodeTemplateMap) ToNodeTemplateMapOutputWithContext(ctx context.Context) NodeTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTemplateMapOutput)
}

func (i NodeTemplateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NodeTemplate] {
	return pulumix.Output[map[string]*NodeTemplate]{
		OutputState: i.ToNodeTemplateMapOutputWithContext(ctx).OutputState,
	}
}

type NodeTemplateOutput struct{ *pulumi.OutputState }

func (NodeTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeTemplate)(nil)).Elem()
}

func (o NodeTemplateOutput) ToNodeTemplateOutput() NodeTemplateOutput {
	return o
}

func (o NodeTemplateOutput) ToNodeTemplateOutputWithContext(ctx context.Context) NodeTemplateOutput {
	return o
}

func (o NodeTemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*NodeTemplate] {
	return pulumix.Output[*NodeTemplate]{
		OutputState: o.OutputState,
	}
}

// CPU overcommit.
// Default value is `NONE`.
// Possible values are: `ENABLED`, `NONE`.
func (o NodeTemplateOutput) CpuOvercommitType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringPtrOutput { return v.CpuOvercommitType }).(pulumi.StringPtrOutput)
}

// Creation timestamp in RFC3339 text format.
func (o NodeTemplateOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional textual description of the resource.
func (o NodeTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o NodeTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Labels to use for node affinity, which will be used in
// instance scheduling.
func (o NodeTemplateOutput) NodeAffinityLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringMapOutput { return v.NodeAffinityLabels }).(pulumi.StringMapOutput)
}

// Node type to use for nodes group that are created from this template.
// Only one of nodeTypeFlexibility and nodeType can be specified.
func (o NodeTemplateOutput) NodeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringPtrOutput { return v.NodeType }).(pulumi.StringPtrOutput)
}

// Flexible properties for the desired node type. Node groups that
// use this node template will create nodes of a type that matches
// these properties. Only one of nodeTypeFlexibility and nodeType can
// be specified.
// Structure is documented below.
func (o NodeTemplateOutput) NodeTypeFlexibility() NodeTemplateNodeTypeFlexibilityPtrOutput {
	return o.ApplyT(func(v *NodeTemplate) NodeTemplateNodeTypeFlexibilityPtrOutput { return v.NodeTypeFlexibility }).(NodeTemplateNodeTypeFlexibilityPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o NodeTemplateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Region where nodes using the node template will be created.
// If it is not provided, the provider region is used.
func (o NodeTemplateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o NodeTemplateOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeTemplate) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The server binding policy for nodes using this template. Determines
// where the nodes should restart following a maintenance event.
// Structure is documented below.
func (o NodeTemplateOutput) ServerBinding() NodeTemplateServerBindingOutput {
	return o.ApplyT(func(v *NodeTemplate) NodeTemplateServerBindingOutput { return v.ServerBinding }).(NodeTemplateServerBindingOutput)
}

type NodeTemplateArrayOutput struct{ *pulumi.OutputState }

func (NodeTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeTemplate)(nil)).Elem()
}

func (o NodeTemplateArrayOutput) ToNodeTemplateArrayOutput() NodeTemplateArrayOutput {
	return o
}

func (o NodeTemplateArrayOutput) ToNodeTemplateArrayOutputWithContext(ctx context.Context) NodeTemplateArrayOutput {
	return o
}

func (o NodeTemplateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NodeTemplate] {
	return pulumix.Output[[]*NodeTemplate]{
		OutputState: o.OutputState,
	}
}

func (o NodeTemplateArrayOutput) Index(i pulumi.IntInput) NodeTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NodeTemplate {
		return vs[0].([]*NodeTemplate)[vs[1].(int)]
	}).(NodeTemplateOutput)
}

type NodeTemplateMapOutput struct{ *pulumi.OutputState }

func (NodeTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeTemplate)(nil)).Elem()
}

func (o NodeTemplateMapOutput) ToNodeTemplateMapOutput() NodeTemplateMapOutput {
	return o
}

func (o NodeTemplateMapOutput) ToNodeTemplateMapOutputWithContext(ctx context.Context) NodeTemplateMapOutput {
	return o
}

func (o NodeTemplateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NodeTemplate] {
	return pulumix.Output[map[string]*NodeTemplate]{
		OutputState: o.OutputState,
	}
}

func (o NodeTemplateMapOutput) MapIndex(k pulumi.StringInput) NodeTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NodeTemplate {
		return vs[0].(map[string]*NodeTemplate)[vs[1].(string)]
	}).(NodeTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeTemplateInput)(nil)).Elem(), &NodeTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeTemplateArrayInput)(nil)).Elem(), NodeTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeTemplateMapInput)(nil)).Elem(), NodeTemplateMap{})
	pulumi.RegisterOutputType(NodeTemplateOutput{})
	pulumi.RegisterOutputType(NodeTemplateArrayOutput{})
	pulumi.RegisterOutputType(NodeTemplateMapOutput{})
}
