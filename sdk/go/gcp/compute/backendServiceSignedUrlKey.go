// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ## Import
//
// This resource does not support import.
type BackendServiceSignedUrlKey struct {
	pulumi.CustomResourceState

	// The backend service this signed URL key belongs.
	BackendService pulumi.StringOutput `pulumi:"backendService"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringOutput `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewBackendServiceSignedUrlKey registers a new resource with the given unique name, arguments, and options.
func NewBackendServiceSignedUrlKey(ctx *pulumi.Context,
	name string, args *BackendServiceSignedUrlKeyArgs, opts ...pulumi.ResourceOption) (*BackendServiceSignedUrlKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendService == nil {
		return nil, errors.New("invalid value for required argument 'BackendService'")
	}
	if args.KeyValue == nil {
		return nil, errors.New("invalid value for required argument 'KeyValue'")
	}
	if args.KeyValue != nil {
		args.KeyValue = pulumi.ToSecret(args.KeyValue).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"keyValue",
	})
	opts = append(opts, secrets)
	var resource BackendServiceSignedUrlKey
	err := ctx.RegisterResource("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendServiceSignedUrlKey gets an existing BackendServiceSignedUrlKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendServiceSignedUrlKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendServiceSignedUrlKeyState, opts ...pulumi.ResourceOption) (*BackendServiceSignedUrlKey, error) {
	var resource BackendServiceSignedUrlKey
	err := ctx.ReadResource("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendServiceSignedUrlKey resources.
type backendServiceSignedUrlKeyState struct {
	// The backend service this signed URL key belongs.
	BackendService *string `pulumi:"backendService"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue *string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type BackendServiceSignedUrlKeyState struct {
	// The backend service this signed URL key belongs.
	BackendService pulumi.StringPtrInput
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringPtrInput
	// Name of the signed URL key.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendServiceSignedUrlKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceSignedUrlKeyState)(nil)).Elem()
}

type backendServiceSignedUrlKeyArgs struct {
	// The backend service this signed URL key belongs.
	BackendService string `pulumi:"backendService"`
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a BackendServiceSignedUrlKey resource.
type BackendServiceSignedUrlKeyArgs struct {
	// The backend service this signed URL key belongs.
	BackendService pulumi.StringInput
	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	KeyValue pulumi.StringInput
	// Name of the signed URL key.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (BackendServiceSignedUrlKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendServiceSignedUrlKeyArgs)(nil)).Elem()
}

type BackendServiceSignedUrlKeyInput interface {
	pulumi.Input

	ToBackendServiceSignedUrlKeyOutput() BackendServiceSignedUrlKeyOutput
	ToBackendServiceSignedUrlKeyOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyOutput
}

func (*BackendServiceSignedUrlKey) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceSignedUrlKey)(nil)).Elem()
}

func (i *BackendServiceSignedUrlKey) ToBackendServiceSignedUrlKeyOutput() BackendServiceSignedUrlKeyOutput {
	return i.ToBackendServiceSignedUrlKeyOutputWithContext(context.Background())
}

func (i *BackendServiceSignedUrlKey) ToBackendServiceSignedUrlKeyOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceSignedUrlKeyOutput)
}

// BackendServiceSignedUrlKeyArrayInput is an input type that accepts BackendServiceSignedUrlKeyArray and BackendServiceSignedUrlKeyArrayOutput values.
// You can construct a concrete instance of `BackendServiceSignedUrlKeyArrayInput` via:
//
//	BackendServiceSignedUrlKeyArray{ BackendServiceSignedUrlKeyArgs{...} }
type BackendServiceSignedUrlKeyArrayInput interface {
	pulumi.Input

	ToBackendServiceSignedUrlKeyArrayOutput() BackendServiceSignedUrlKeyArrayOutput
	ToBackendServiceSignedUrlKeyArrayOutputWithContext(context.Context) BackendServiceSignedUrlKeyArrayOutput
}

type BackendServiceSignedUrlKeyArray []BackendServiceSignedUrlKeyInput

func (BackendServiceSignedUrlKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServiceSignedUrlKey)(nil)).Elem()
}

func (i BackendServiceSignedUrlKeyArray) ToBackendServiceSignedUrlKeyArrayOutput() BackendServiceSignedUrlKeyArrayOutput {
	return i.ToBackendServiceSignedUrlKeyArrayOutputWithContext(context.Background())
}

func (i BackendServiceSignedUrlKeyArray) ToBackendServiceSignedUrlKeyArrayOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceSignedUrlKeyArrayOutput)
}

// BackendServiceSignedUrlKeyMapInput is an input type that accepts BackendServiceSignedUrlKeyMap and BackendServiceSignedUrlKeyMapOutput values.
// You can construct a concrete instance of `BackendServiceSignedUrlKeyMapInput` via:
//
//	BackendServiceSignedUrlKeyMap{ "key": BackendServiceSignedUrlKeyArgs{...} }
type BackendServiceSignedUrlKeyMapInput interface {
	pulumi.Input

	ToBackendServiceSignedUrlKeyMapOutput() BackendServiceSignedUrlKeyMapOutput
	ToBackendServiceSignedUrlKeyMapOutputWithContext(context.Context) BackendServiceSignedUrlKeyMapOutput
}

type BackendServiceSignedUrlKeyMap map[string]BackendServiceSignedUrlKeyInput

func (BackendServiceSignedUrlKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServiceSignedUrlKey)(nil)).Elem()
}

func (i BackendServiceSignedUrlKeyMap) ToBackendServiceSignedUrlKeyMapOutput() BackendServiceSignedUrlKeyMapOutput {
	return i.ToBackendServiceSignedUrlKeyMapOutputWithContext(context.Background())
}

func (i BackendServiceSignedUrlKeyMap) ToBackendServiceSignedUrlKeyMapOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceSignedUrlKeyMapOutput)
}

type BackendServiceSignedUrlKeyOutput struct{ *pulumi.OutputState }

func (BackendServiceSignedUrlKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceSignedUrlKey)(nil)).Elem()
}

func (o BackendServiceSignedUrlKeyOutput) ToBackendServiceSignedUrlKeyOutput() BackendServiceSignedUrlKeyOutput {
	return o
}

func (o BackendServiceSignedUrlKeyOutput) ToBackendServiceSignedUrlKeyOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyOutput {
	return o
}

// The backend service this signed URL key belongs.
func (o BackendServiceSignedUrlKeyOutput) BackendService() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceSignedUrlKey) pulumi.StringOutput { return v.BackendService }).(pulumi.StringOutput)
}

// 128-bit key value used for signing the URL. The key value must be a
// valid RFC 4648 Section 5 base64url encoded string.
// **Note**: This property is sensitive and will not be displayed in the plan.
func (o BackendServiceSignedUrlKeyOutput) KeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceSignedUrlKey) pulumi.StringOutput { return v.KeyValue }).(pulumi.StringOutput)
}

// Name of the signed URL key.
func (o BackendServiceSignedUrlKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceSignedUrlKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o BackendServiceSignedUrlKeyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackendServiceSignedUrlKey) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type BackendServiceSignedUrlKeyArrayOutput struct{ *pulumi.OutputState }

func (BackendServiceSignedUrlKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendServiceSignedUrlKey)(nil)).Elem()
}

func (o BackendServiceSignedUrlKeyArrayOutput) ToBackendServiceSignedUrlKeyArrayOutput() BackendServiceSignedUrlKeyArrayOutput {
	return o
}

func (o BackendServiceSignedUrlKeyArrayOutput) ToBackendServiceSignedUrlKeyArrayOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyArrayOutput {
	return o
}

func (o BackendServiceSignedUrlKeyArrayOutput) Index(i pulumi.IntInput) BackendServiceSignedUrlKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackendServiceSignedUrlKey {
		return vs[0].([]*BackendServiceSignedUrlKey)[vs[1].(int)]
	}).(BackendServiceSignedUrlKeyOutput)
}

type BackendServiceSignedUrlKeyMapOutput struct{ *pulumi.OutputState }

func (BackendServiceSignedUrlKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendServiceSignedUrlKey)(nil)).Elem()
}

func (o BackendServiceSignedUrlKeyMapOutput) ToBackendServiceSignedUrlKeyMapOutput() BackendServiceSignedUrlKeyMapOutput {
	return o
}

func (o BackendServiceSignedUrlKeyMapOutput) ToBackendServiceSignedUrlKeyMapOutputWithContext(ctx context.Context) BackendServiceSignedUrlKeyMapOutput {
	return o
}

func (o BackendServiceSignedUrlKeyMapOutput) MapIndex(k pulumi.StringInput) BackendServiceSignedUrlKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackendServiceSignedUrlKey {
		return vs[0].(map[string]*BackendServiceSignedUrlKey)[vs[1].(string)]
	}).(BackendServiceSignedUrlKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceSignedUrlKeyInput)(nil)).Elem(), &BackendServiceSignedUrlKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceSignedUrlKeyArrayInput)(nil)).Elem(), BackendServiceSignedUrlKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackendServiceSignedUrlKeyMapInput)(nil)).Elem(), BackendServiceSignedUrlKeyMap{})
	pulumi.RegisterOutputType(BackendServiceSignedUrlKeyOutput{})
	pulumi.RegisterOutputType(BackendServiceSignedUrlKeyArrayOutput{})
	pulumi.RegisterOutputType(BackendServiceSignedUrlKeyMapOutput{})
}
