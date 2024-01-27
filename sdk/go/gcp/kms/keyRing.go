// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A `KeyRing` is a toplevel logical grouping of `CryptoKeys`.
//
// > **Note:** KeyRings cannot be deleted from Google Cloud Platform.
// Destroying a provider-managed KeyRing will remove it from state but
// *will not delete the resource from the project.*
//
// To get more information about KeyRing, see:
//
// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings)
// * How-to Guides
//   - [Creating a key ring](https://cloud.google.com/kms/docs/creating-keys#create_a_key_ring)
//
// ## Example Usage
// ### Kms Key Ring Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kms.NewKeyRing(ctx, "example-keyring", &kms.KeyRingArgs{
//				Location: pulumi.String("global"),
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
// KeyRing can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/keyRings/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, KeyRing can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:kms/keyRing:KeyRing default projects/{{project}}/locations/{{location}}/keyRings/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:kms/keyRing:KeyRing default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:kms/keyRing:KeyRing default {{location}}/{{name}}
//
// ```
type KeyRing struct {
	pulumi.CustomResourceState

	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	//
	// ***
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for the KeyRing.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewKeyRing registers a new resource with the given unique name, arguments, and options.
func NewKeyRing(ctx *pulumi.Context,
	name string, args *KeyRingArgs, opts ...pulumi.ResourceOption) (*KeyRing, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KeyRing
	err := ctx.RegisterResource("gcp:kms/keyRing:KeyRing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyRing gets an existing KeyRing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyRingState, opts ...pulumi.ResourceOption) (*KeyRing, error) {
	var resource KeyRing
	err := ctx.ReadResource("gcp:kms/keyRing:KeyRing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyRing resources.
type keyRingState struct {
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	//
	// ***
	Location *string `pulumi:"location"`
	// The resource name for the KeyRing.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type KeyRingState struct {
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	//
	// ***
	Location pulumi.StringPtrInput
	// The resource name for the KeyRing.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (KeyRingState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingState)(nil)).Elem()
}

type keyRingArgs struct {
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	//
	// ***
	Location string `pulumi:"location"`
	// The resource name for the KeyRing.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a KeyRing resource.
type KeyRingArgs struct {
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	//
	// ***
	Location pulumi.StringInput
	// The resource name for the KeyRing.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (KeyRingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingArgs)(nil)).Elem()
}

type KeyRingInput interface {
	pulumi.Input

	ToKeyRingOutput() KeyRingOutput
	ToKeyRingOutputWithContext(ctx context.Context) KeyRingOutput
}

func (*KeyRing) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRing)(nil)).Elem()
}

func (i *KeyRing) ToKeyRingOutput() KeyRingOutput {
	return i.ToKeyRingOutputWithContext(context.Background())
}

func (i *KeyRing) ToKeyRingOutputWithContext(ctx context.Context) KeyRingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingOutput)
}

// KeyRingArrayInput is an input type that accepts KeyRingArray and KeyRingArrayOutput values.
// You can construct a concrete instance of `KeyRingArrayInput` via:
//
//	KeyRingArray{ KeyRingArgs{...} }
type KeyRingArrayInput interface {
	pulumi.Input

	ToKeyRingArrayOutput() KeyRingArrayOutput
	ToKeyRingArrayOutputWithContext(context.Context) KeyRingArrayOutput
}

type KeyRingArray []KeyRingInput

func (KeyRingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyRing)(nil)).Elem()
}

func (i KeyRingArray) ToKeyRingArrayOutput() KeyRingArrayOutput {
	return i.ToKeyRingArrayOutputWithContext(context.Background())
}

func (i KeyRingArray) ToKeyRingArrayOutputWithContext(ctx context.Context) KeyRingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingArrayOutput)
}

// KeyRingMapInput is an input type that accepts KeyRingMap and KeyRingMapOutput values.
// You can construct a concrete instance of `KeyRingMapInput` via:
//
//	KeyRingMap{ "key": KeyRingArgs{...} }
type KeyRingMapInput interface {
	pulumi.Input

	ToKeyRingMapOutput() KeyRingMapOutput
	ToKeyRingMapOutputWithContext(context.Context) KeyRingMapOutput
}

type KeyRingMap map[string]KeyRingInput

func (KeyRingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyRing)(nil)).Elem()
}

func (i KeyRingMap) ToKeyRingMapOutput() KeyRingMapOutput {
	return i.ToKeyRingMapOutputWithContext(context.Background())
}

func (i KeyRingMap) ToKeyRingMapOutputWithContext(ctx context.Context) KeyRingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingMapOutput)
}

type KeyRingOutput struct{ *pulumi.OutputState }

func (KeyRingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRing)(nil)).Elem()
}

func (o KeyRingOutput) ToKeyRingOutput() KeyRingOutput {
	return o
}

func (o KeyRingOutput) ToKeyRingOutputWithContext(ctx context.Context) KeyRingOutput {
	return o
}

// The location for the KeyRing.
// A full list of valid locations can be found by running `gcloud kms locations list`.
//
// ***
func (o KeyRingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRing) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name for the KeyRing.
func (o KeyRingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRing) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o KeyRingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyRing) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type KeyRingArrayOutput struct{ *pulumi.OutputState }

func (KeyRingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyRing)(nil)).Elem()
}

func (o KeyRingArrayOutput) ToKeyRingArrayOutput() KeyRingArrayOutput {
	return o
}

func (o KeyRingArrayOutput) ToKeyRingArrayOutputWithContext(ctx context.Context) KeyRingArrayOutput {
	return o
}

func (o KeyRingArrayOutput) Index(i pulumi.IntInput) KeyRingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KeyRing {
		return vs[0].([]*KeyRing)[vs[1].(int)]
	}).(KeyRingOutput)
}

type KeyRingMapOutput struct{ *pulumi.OutputState }

func (KeyRingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyRing)(nil)).Elem()
}

func (o KeyRingMapOutput) ToKeyRingMapOutput() KeyRingMapOutput {
	return o
}

func (o KeyRingMapOutput) ToKeyRingMapOutputWithContext(ctx context.Context) KeyRingMapOutput {
	return o
}

func (o KeyRingMapOutput) MapIndex(k pulumi.StringInput) KeyRingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KeyRing {
		return vs[0].(map[string]*KeyRing)[vs[1].(string)]
	}).(KeyRingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRingInput)(nil)).Elem(), &KeyRing{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRingArrayInput)(nil)).Elem(), KeyRingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRingMapInput)(nil)).Elem(), KeyRingMap{})
	pulumi.RegisterOutputType(KeyRingOutput{})
	pulumi.RegisterOutputType(KeyRingArrayOutput{})
	pulumi.RegisterOutputType(KeyRingMapOutput{})
}
