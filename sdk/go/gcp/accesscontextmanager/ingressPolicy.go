// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource has been deprecated, please refer to ServicePerimeterIngressPolicy.
//
// To get more information about IngressPolicy, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters#ingresspolicy)
//
// ## Import
//
// IngressPolicy can be imported using any of these accepted formats* `{{ingress_policy_name}}/{{resource}}` When using the `pulumi import` command, IngressPolicy can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:accesscontextmanager/ingressPolicy:IngressPolicy default {{ingress_policy_name}}/{{resource}}
//
// ```
type IngressPolicy struct {
	pulumi.CustomResourceState

	// The name of the Service Perimeter to add this resource to.
	//
	// ***
	IngressPolicyName pulumi.StringOutput `pulumi:"ingressPolicyName"`
	// A GCP resource that is inside of the service perimeter.
	Resource pulumi.StringOutput `pulumi:"resource"`
}

// NewIngressPolicy registers a new resource with the given unique name, arguments, and options.
func NewIngressPolicy(ctx *pulumi.Context,
	name string, args *IngressPolicyArgs, opts ...pulumi.ResourceOption) (*IngressPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IngressPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'IngressPolicyName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IngressPolicy
	err := ctx.RegisterResource("gcp:accesscontextmanager/ingressPolicy:IngressPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIngressPolicy gets an existing IngressPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIngressPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IngressPolicyState, opts ...pulumi.ResourceOption) (*IngressPolicy, error) {
	var resource IngressPolicy
	err := ctx.ReadResource("gcp:accesscontextmanager/ingressPolicy:IngressPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IngressPolicy resources.
type ingressPolicyState struct {
	// The name of the Service Perimeter to add this resource to.
	//
	// ***
	IngressPolicyName *string `pulumi:"ingressPolicyName"`
	// A GCP resource that is inside of the service perimeter.
	Resource *string `pulumi:"resource"`
}

type IngressPolicyState struct {
	// The name of the Service Perimeter to add this resource to.
	//
	// ***
	IngressPolicyName pulumi.StringPtrInput
	// A GCP resource that is inside of the service perimeter.
	Resource pulumi.StringPtrInput
}

func (IngressPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressPolicyState)(nil)).Elem()
}

type ingressPolicyArgs struct {
	// The name of the Service Perimeter to add this resource to.
	//
	// ***
	IngressPolicyName string `pulumi:"ingressPolicyName"`
	// A GCP resource that is inside of the service perimeter.
	Resource string `pulumi:"resource"`
}

// The set of arguments for constructing a IngressPolicy resource.
type IngressPolicyArgs struct {
	// The name of the Service Perimeter to add this resource to.
	//
	// ***
	IngressPolicyName pulumi.StringInput
	// A GCP resource that is inside of the service perimeter.
	Resource pulumi.StringInput
}

func (IngressPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ingressPolicyArgs)(nil)).Elem()
}

type IngressPolicyInput interface {
	pulumi.Input

	ToIngressPolicyOutput() IngressPolicyOutput
	ToIngressPolicyOutputWithContext(ctx context.Context) IngressPolicyOutput
}

func (*IngressPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressPolicy)(nil)).Elem()
}

func (i *IngressPolicy) ToIngressPolicyOutput() IngressPolicyOutput {
	return i.ToIngressPolicyOutputWithContext(context.Background())
}

func (i *IngressPolicy) ToIngressPolicyOutputWithContext(ctx context.Context) IngressPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressPolicyOutput)
}

// IngressPolicyArrayInput is an input type that accepts IngressPolicyArray and IngressPolicyArrayOutput values.
// You can construct a concrete instance of `IngressPolicyArrayInput` via:
//
//	IngressPolicyArray{ IngressPolicyArgs{...} }
type IngressPolicyArrayInput interface {
	pulumi.Input

	ToIngressPolicyArrayOutput() IngressPolicyArrayOutput
	ToIngressPolicyArrayOutputWithContext(context.Context) IngressPolicyArrayOutput
}

type IngressPolicyArray []IngressPolicyInput

func (IngressPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IngressPolicy)(nil)).Elem()
}

func (i IngressPolicyArray) ToIngressPolicyArrayOutput() IngressPolicyArrayOutput {
	return i.ToIngressPolicyArrayOutputWithContext(context.Background())
}

func (i IngressPolicyArray) ToIngressPolicyArrayOutputWithContext(ctx context.Context) IngressPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressPolicyArrayOutput)
}

// IngressPolicyMapInput is an input type that accepts IngressPolicyMap and IngressPolicyMapOutput values.
// You can construct a concrete instance of `IngressPolicyMapInput` via:
//
//	IngressPolicyMap{ "key": IngressPolicyArgs{...} }
type IngressPolicyMapInput interface {
	pulumi.Input

	ToIngressPolicyMapOutput() IngressPolicyMapOutput
	ToIngressPolicyMapOutputWithContext(context.Context) IngressPolicyMapOutput
}

type IngressPolicyMap map[string]IngressPolicyInput

func (IngressPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IngressPolicy)(nil)).Elem()
}

func (i IngressPolicyMap) ToIngressPolicyMapOutput() IngressPolicyMapOutput {
	return i.ToIngressPolicyMapOutputWithContext(context.Background())
}

func (i IngressPolicyMap) ToIngressPolicyMapOutputWithContext(ctx context.Context) IngressPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressPolicyMapOutput)
}

type IngressPolicyOutput struct{ *pulumi.OutputState }

func (IngressPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressPolicy)(nil)).Elem()
}

func (o IngressPolicyOutput) ToIngressPolicyOutput() IngressPolicyOutput {
	return o
}

func (o IngressPolicyOutput) ToIngressPolicyOutputWithContext(ctx context.Context) IngressPolicyOutput {
	return o
}

// The name of the Service Perimeter to add this resource to.
//
// ***
func (o IngressPolicyOutput) IngressPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *IngressPolicy) pulumi.StringOutput { return v.IngressPolicyName }).(pulumi.StringOutput)
}

// A GCP resource that is inside of the service perimeter.
func (o IngressPolicyOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v *IngressPolicy) pulumi.StringOutput { return v.Resource }).(pulumi.StringOutput)
}

type IngressPolicyArrayOutput struct{ *pulumi.OutputState }

func (IngressPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IngressPolicy)(nil)).Elem()
}

func (o IngressPolicyArrayOutput) ToIngressPolicyArrayOutput() IngressPolicyArrayOutput {
	return o
}

func (o IngressPolicyArrayOutput) ToIngressPolicyArrayOutputWithContext(ctx context.Context) IngressPolicyArrayOutput {
	return o
}

func (o IngressPolicyArrayOutput) Index(i pulumi.IntInput) IngressPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IngressPolicy {
		return vs[0].([]*IngressPolicy)[vs[1].(int)]
	}).(IngressPolicyOutput)
}

type IngressPolicyMapOutput struct{ *pulumi.OutputState }

func (IngressPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IngressPolicy)(nil)).Elem()
}

func (o IngressPolicyMapOutput) ToIngressPolicyMapOutput() IngressPolicyMapOutput {
	return o
}

func (o IngressPolicyMapOutput) ToIngressPolicyMapOutputWithContext(ctx context.Context) IngressPolicyMapOutput {
	return o
}

func (o IngressPolicyMapOutput) MapIndex(k pulumi.StringInput) IngressPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IngressPolicy {
		return vs[0].(map[string]*IngressPolicy)[vs[1].(string)]
	}).(IngressPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IngressPolicyInput)(nil)).Elem(), &IngressPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressPolicyArrayInput)(nil)).Elem(), IngressPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IngressPolicyMapInput)(nil)).Elem(), IngressPolicyMap{})
	pulumi.RegisterOutputType(IngressPolicyOutput{})
	pulumi.RegisterOutputType(IngressPolicyArrayOutput{})
	pulumi.RegisterOutputType(IngressPolicyMapOutput{})
}
