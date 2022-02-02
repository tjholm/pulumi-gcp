// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Restricts access to Cloud Console and Google Cloud APIs for a set of users using Context-Aware Access.
//
// To get more information about GcpUserAccessBinding, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/organizations.gcpUserAccessBindings)
//
// ## Example Usage
//
// ## Import
//
// GcpUserAccessBinding can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/gcpUserAccessBinding:GcpUserAccessBinding default {{name}}
// ```
type GcpUserAccessBinding struct {
	pulumi.CustomResourceState

	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels pulumi.StringOutput `pulumi:"accessLevels"`
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey pulumi.StringOutput `pulumi:"groupKey"`
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved
	// characters (as defined by RFC 3986 Section 2.3). Should not be specified by the client during creation. Example:
	// "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. ID of the parent organization.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
}

// NewGcpUserAccessBinding registers a new resource with the given unique name, arguments, and options.
func NewGcpUserAccessBinding(ctx *pulumi.Context,
	name string, args *GcpUserAccessBindingArgs, opts ...pulumi.ResourceOption) (*GcpUserAccessBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessLevels == nil {
		return nil, errors.New("invalid value for required argument 'AccessLevels'")
	}
	if args.GroupKey == nil {
		return nil, errors.New("invalid value for required argument 'GroupKey'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource GcpUserAccessBinding
	err := ctx.RegisterResource("gcp:accesscontextmanager/gcpUserAccessBinding:GcpUserAccessBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGcpUserAccessBinding gets an existing GcpUserAccessBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGcpUserAccessBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GcpUserAccessBindingState, opts ...pulumi.ResourceOption) (*GcpUserAccessBinding, error) {
	var resource GcpUserAccessBinding
	err := ctx.ReadResource("gcp:accesscontextmanager/gcpUserAccessBinding:GcpUserAccessBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GcpUserAccessBinding resources.
type gcpUserAccessBindingState struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels *string `pulumi:"accessLevels"`
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey *string `pulumi:"groupKey"`
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved
	// characters (as defined by RFC 3986 Section 2.3). Should not be specified by the client during creation. Example:
	// "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name *string `pulumi:"name"`
	// Required. ID of the parent organization.
	OrganizationId *string `pulumi:"organizationId"`
}

type GcpUserAccessBindingState struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels pulumi.StringPtrInput
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey pulumi.StringPtrInput
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved
	// characters (as defined by RFC 3986 Section 2.3). Should not be specified by the client during creation. Example:
	// "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name pulumi.StringPtrInput
	// Required. ID of the parent organization.
	OrganizationId pulumi.StringPtrInput
}

func (GcpUserAccessBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpUserAccessBindingState)(nil)).Elem()
}

type gcpUserAccessBindingArgs struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels string `pulumi:"accessLevels"`
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey string `pulumi:"groupKey"`
	// Required. ID of the parent organization.
	OrganizationId string `pulumi:"organizationId"`
}

// The set of arguments for constructing a GcpUserAccessBinding resource.
type GcpUserAccessBindingArgs struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels pulumi.StringInput
	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey pulumi.StringInput
	// Required. ID of the parent organization.
	OrganizationId pulumi.StringInput
}

func (GcpUserAccessBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpUserAccessBindingArgs)(nil)).Elem()
}

type GcpUserAccessBindingInput interface {
	pulumi.Input

	ToGcpUserAccessBindingOutput() GcpUserAccessBindingOutput
	ToGcpUserAccessBindingOutputWithContext(ctx context.Context) GcpUserAccessBindingOutput
}

func (*GcpUserAccessBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**GcpUserAccessBinding)(nil)).Elem()
}

func (i *GcpUserAccessBinding) ToGcpUserAccessBindingOutput() GcpUserAccessBindingOutput {
	return i.ToGcpUserAccessBindingOutputWithContext(context.Background())
}

func (i *GcpUserAccessBinding) ToGcpUserAccessBindingOutputWithContext(ctx context.Context) GcpUserAccessBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpUserAccessBindingOutput)
}

// GcpUserAccessBindingArrayInput is an input type that accepts GcpUserAccessBindingArray and GcpUserAccessBindingArrayOutput values.
// You can construct a concrete instance of `GcpUserAccessBindingArrayInput` via:
//
//          GcpUserAccessBindingArray{ GcpUserAccessBindingArgs{...} }
type GcpUserAccessBindingArrayInput interface {
	pulumi.Input

	ToGcpUserAccessBindingArrayOutput() GcpUserAccessBindingArrayOutput
	ToGcpUserAccessBindingArrayOutputWithContext(context.Context) GcpUserAccessBindingArrayOutput
}

type GcpUserAccessBindingArray []GcpUserAccessBindingInput

func (GcpUserAccessBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GcpUserAccessBinding)(nil)).Elem()
}

func (i GcpUserAccessBindingArray) ToGcpUserAccessBindingArrayOutput() GcpUserAccessBindingArrayOutput {
	return i.ToGcpUserAccessBindingArrayOutputWithContext(context.Background())
}

func (i GcpUserAccessBindingArray) ToGcpUserAccessBindingArrayOutputWithContext(ctx context.Context) GcpUserAccessBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpUserAccessBindingArrayOutput)
}

// GcpUserAccessBindingMapInput is an input type that accepts GcpUserAccessBindingMap and GcpUserAccessBindingMapOutput values.
// You can construct a concrete instance of `GcpUserAccessBindingMapInput` via:
//
//          GcpUserAccessBindingMap{ "key": GcpUserAccessBindingArgs{...} }
type GcpUserAccessBindingMapInput interface {
	pulumi.Input

	ToGcpUserAccessBindingMapOutput() GcpUserAccessBindingMapOutput
	ToGcpUserAccessBindingMapOutputWithContext(context.Context) GcpUserAccessBindingMapOutput
}

type GcpUserAccessBindingMap map[string]GcpUserAccessBindingInput

func (GcpUserAccessBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GcpUserAccessBinding)(nil)).Elem()
}

func (i GcpUserAccessBindingMap) ToGcpUserAccessBindingMapOutput() GcpUserAccessBindingMapOutput {
	return i.ToGcpUserAccessBindingMapOutputWithContext(context.Background())
}

func (i GcpUserAccessBindingMap) ToGcpUserAccessBindingMapOutputWithContext(ctx context.Context) GcpUserAccessBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpUserAccessBindingMapOutput)
}

type GcpUserAccessBindingOutput struct{ *pulumi.OutputState }

func (GcpUserAccessBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcpUserAccessBinding)(nil)).Elem()
}

func (o GcpUserAccessBindingOutput) ToGcpUserAccessBindingOutput() GcpUserAccessBindingOutput {
	return o
}

func (o GcpUserAccessBindingOutput) ToGcpUserAccessBindingOutputWithContext(ctx context.Context) GcpUserAccessBindingOutput {
	return o
}

type GcpUserAccessBindingArrayOutput struct{ *pulumi.OutputState }

func (GcpUserAccessBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GcpUserAccessBinding)(nil)).Elem()
}

func (o GcpUserAccessBindingArrayOutput) ToGcpUserAccessBindingArrayOutput() GcpUserAccessBindingArrayOutput {
	return o
}

func (o GcpUserAccessBindingArrayOutput) ToGcpUserAccessBindingArrayOutputWithContext(ctx context.Context) GcpUserAccessBindingArrayOutput {
	return o
}

func (o GcpUserAccessBindingArrayOutput) Index(i pulumi.IntInput) GcpUserAccessBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GcpUserAccessBinding {
		return vs[0].([]*GcpUserAccessBinding)[vs[1].(int)]
	}).(GcpUserAccessBindingOutput)
}

type GcpUserAccessBindingMapOutput struct{ *pulumi.OutputState }

func (GcpUserAccessBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GcpUserAccessBinding)(nil)).Elem()
}

func (o GcpUserAccessBindingMapOutput) ToGcpUserAccessBindingMapOutput() GcpUserAccessBindingMapOutput {
	return o
}

func (o GcpUserAccessBindingMapOutput) ToGcpUserAccessBindingMapOutputWithContext(ctx context.Context) GcpUserAccessBindingMapOutput {
	return o
}

func (o GcpUserAccessBindingMapOutput) MapIndex(k pulumi.StringInput) GcpUserAccessBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GcpUserAccessBinding {
		return vs[0].(map[string]*GcpUserAccessBinding)[vs[1].(string)]
	}).(GcpUserAccessBindingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GcpUserAccessBindingInput)(nil)).Elem(), &GcpUserAccessBinding{})
	pulumi.RegisterInputType(reflect.TypeOf((*GcpUserAccessBindingArrayInput)(nil)).Elem(), GcpUserAccessBindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GcpUserAccessBindingMapInput)(nil)).Elem(), GcpUserAccessBindingMap{})
	pulumi.RegisterOutputType(GcpUserAccessBindingOutput{})
	pulumi.RegisterOutputType(GcpUserAccessBindingArrayOutput{})
	pulumi.RegisterOutputType(GcpUserAccessBindingMapOutput{})
}
