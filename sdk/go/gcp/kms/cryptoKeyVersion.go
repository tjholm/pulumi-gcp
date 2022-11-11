// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A `CryptoKeyVersion` represents an individual cryptographic key, and the associated key material.
//
// Destroying a cryptoKeyVersion will not delete the resource from the project.
//
// To get more information about CryptoKeyVersion, see:
//
// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions)
// * How-to Guides
//   - [Creating a key Version](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions/create)
//
// ## Example Usage
// ### Kms Crypto Key Version Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			keyring, err := kms.NewKeyRing(ctx, "keyring", &kms.KeyRingArgs{
//				Location: pulumi.String("global"),
//			})
//			if err != nil {
//				return err
//			}
//			cryptokey, err := kms.NewCryptoKey(ctx, "cryptokey", &kms.CryptoKeyArgs{
//				KeyRing:        keyring.ID(),
//				RotationPeriod: pulumi.String("100000s"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kms.NewCryptoKeyVersion(ctx, "example-key", &kms.CryptoKeyVersionArgs{
//				CryptoKey: cryptokey.ID(),
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
// # CryptoKeyVersion can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:kms/cryptoKeyVersion:CryptoKeyVersion default {{name}}
//
// ```
type CryptoKeyVersion struct {
	pulumi.CustomResourceState

	// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the
	// key as stored on the HSM, independently of Google. Only provided for key versions with protectionLevel HSM.
	Attestations CryptoKeyVersionAttestationArrayOutput `pulumi:"attestations"`
	// The name of the cryptoKey associated with the CryptoKeyVersions.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringOutput `pulumi:"cryptoKey"`
	// The time this CryptoKeyVersion key material was generated
	GenerateTime pulumi.StringOutput `pulumi:"generateTime"`
	// The resource name for this CryptoKeyVersion.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
	ProtectionLevel pulumi.StringOutput `pulumi:"protectionLevel"`
	// The current state of the CryptoKeyVersion.
	// Possible values are `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, and `IMPORT_FAILED`.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewCryptoKeyVersion registers a new resource with the given unique name, arguments, and options.
func NewCryptoKeyVersion(ctx *pulumi.Context,
	name string, args *CryptoKeyVersionArgs, opts ...pulumi.ResourceOption) (*CryptoKeyVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CryptoKey == nil {
		return nil, errors.New("invalid value for required argument 'CryptoKey'")
	}
	var resource CryptoKeyVersion
	err := ctx.RegisterResource("gcp:kms/cryptoKeyVersion:CryptoKeyVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCryptoKeyVersion gets an existing CryptoKeyVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCryptoKeyVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CryptoKeyVersionState, opts ...pulumi.ResourceOption) (*CryptoKeyVersion, error) {
	var resource CryptoKeyVersion
	err := ctx.ReadResource("gcp:kms/cryptoKeyVersion:CryptoKeyVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CryptoKeyVersion resources.
type cryptoKeyVersionState struct {
	// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
	Algorithm *string `pulumi:"algorithm"`
	// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the
	// key as stored on the HSM, independently of Google. Only provided for key versions with protectionLevel HSM.
	Attestations []CryptoKeyVersionAttestation `pulumi:"attestations"`
	// The name of the cryptoKey associated with the CryptoKeyVersions.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey *string `pulumi:"cryptoKey"`
	// The time this CryptoKeyVersion key material was generated
	GenerateTime *string `pulumi:"generateTime"`
	// The resource name for this CryptoKeyVersion.
	Name *string `pulumi:"name"`
	// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
	ProtectionLevel *string `pulumi:"protectionLevel"`
	// The current state of the CryptoKeyVersion.
	// Possible values are `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, and `IMPORT_FAILED`.
	State *string `pulumi:"state"`
}

type CryptoKeyVersionState struct {
	// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
	Algorithm pulumi.StringPtrInput
	// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the
	// key as stored on the HSM, independently of Google. Only provided for key versions with protectionLevel HSM.
	Attestations CryptoKeyVersionAttestationArrayInput
	// The name of the cryptoKey associated with the CryptoKeyVersions.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringPtrInput
	// The time this CryptoKeyVersion key material was generated
	GenerateTime pulumi.StringPtrInput
	// The resource name for this CryptoKeyVersion.
	Name pulumi.StringPtrInput
	// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
	ProtectionLevel pulumi.StringPtrInput
	// The current state of the CryptoKeyVersion.
	// Possible values are `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, and `IMPORT_FAILED`.
	State pulumi.StringPtrInput
}

func (CryptoKeyVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyVersionState)(nil)).Elem()
}

type cryptoKeyVersionArgs struct {
	// The name of the cryptoKey associated with the CryptoKeyVersions.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey string `pulumi:"cryptoKey"`
	// The current state of the CryptoKeyVersion.
	// Possible values are `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, and `IMPORT_FAILED`.
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a CryptoKeyVersion resource.
type CryptoKeyVersionArgs struct {
	// The name of the cryptoKey associated with the CryptoKeyVersions.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringInput
	// The current state of the CryptoKeyVersion.
	// Possible values are `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, and `IMPORT_FAILED`.
	State pulumi.StringPtrInput
}

func (CryptoKeyVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyVersionArgs)(nil)).Elem()
}

type CryptoKeyVersionInput interface {
	pulumi.Input

	ToCryptoKeyVersionOutput() CryptoKeyVersionOutput
	ToCryptoKeyVersionOutputWithContext(ctx context.Context) CryptoKeyVersionOutput
}

func (*CryptoKeyVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**CryptoKeyVersion)(nil)).Elem()
}

func (i *CryptoKeyVersion) ToCryptoKeyVersionOutput() CryptoKeyVersionOutput {
	return i.ToCryptoKeyVersionOutputWithContext(context.Background())
}

func (i *CryptoKeyVersion) ToCryptoKeyVersionOutputWithContext(ctx context.Context) CryptoKeyVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyVersionOutput)
}

// CryptoKeyVersionArrayInput is an input type that accepts CryptoKeyVersionArray and CryptoKeyVersionArrayOutput values.
// You can construct a concrete instance of `CryptoKeyVersionArrayInput` via:
//
//	CryptoKeyVersionArray{ CryptoKeyVersionArgs{...} }
type CryptoKeyVersionArrayInput interface {
	pulumi.Input

	ToCryptoKeyVersionArrayOutput() CryptoKeyVersionArrayOutput
	ToCryptoKeyVersionArrayOutputWithContext(context.Context) CryptoKeyVersionArrayOutput
}

type CryptoKeyVersionArray []CryptoKeyVersionInput

func (CryptoKeyVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CryptoKeyVersion)(nil)).Elem()
}

func (i CryptoKeyVersionArray) ToCryptoKeyVersionArrayOutput() CryptoKeyVersionArrayOutput {
	return i.ToCryptoKeyVersionArrayOutputWithContext(context.Background())
}

func (i CryptoKeyVersionArray) ToCryptoKeyVersionArrayOutputWithContext(ctx context.Context) CryptoKeyVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyVersionArrayOutput)
}

// CryptoKeyVersionMapInput is an input type that accepts CryptoKeyVersionMap and CryptoKeyVersionMapOutput values.
// You can construct a concrete instance of `CryptoKeyVersionMapInput` via:
//
//	CryptoKeyVersionMap{ "key": CryptoKeyVersionArgs{...} }
type CryptoKeyVersionMapInput interface {
	pulumi.Input

	ToCryptoKeyVersionMapOutput() CryptoKeyVersionMapOutput
	ToCryptoKeyVersionMapOutputWithContext(context.Context) CryptoKeyVersionMapOutput
}

type CryptoKeyVersionMap map[string]CryptoKeyVersionInput

func (CryptoKeyVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CryptoKeyVersion)(nil)).Elem()
}

func (i CryptoKeyVersionMap) ToCryptoKeyVersionMapOutput() CryptoKeyVersionMapOutput {
	return i.ToCryptoKeyVersionMapOutputWithContext(context.Background())
}

func (i CryptoKeyVersionMap) ToCryptoKeyVersionMapOutputWithContext(ctx context.Context) CryptoKeyVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyVersionMapOutput)
}

type CryptoKeyVersionOutput struct{ *pulumi.OutputState }

func (CryptoKeyVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CryptoKeyVersion)(nil)).Elem()
}

func (o CryptoKeyVersionOutput) ToCryptoKeyVersionOutput() CryptoKeyVersionOutput {
	return o
}

func (o CryptoKeyVersionOutput) ToCryptoKeyVersionOutputWithContext(ctx context.Context) CryptoKeyVersionOutput {
	return o
}

// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
func (o CryptoKeyVersionOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *CryptoKeyVersion) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the
// key as stored on the HSM, independently of Google. Only provided for key versions with protectionLevel HSM.
func (o CryptoKeyVersionOutput) Attestations() CryptoKeyVersionAttestationArrayOutput {
	return o.ApplyT(func(v *CryptoKeyVersion) CryptoKeyVersionAttestationArrayOutput { return v.Attestations }).(CryptoKeyVersionAttestationArrayOutput)
}

// The name of the cryptoKey associated with the CryptoKeyVersions.
// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
func (o CryptoKeyVersionOutput) CryptoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CryptoKeyVersion) pulumi.StringOutput { return v.CryptoKey }).(pulumi.StringOutput)
}

// The time this CryptoKeyVersion key material was generated
func (o CryptoKeyVersionOutput) GenerateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CryptoKeyVersion) pulumi.StringOutput { return v.GenerateTime }).(pulumi.StringOutput)
}

// The resource name for this CryptoKeyVersion.
func (o CryptoKeyVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CryptoKeyVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
func (o CryptoKeyVersionOutput) ProtectionLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *CryptoKeyVersion) pulumi.StringOutput { return v.ProtectionLevel }).(pulumi.StringOutput)
}

// The current state of the CryptoKeyVersion.
// Possible values are `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, and `IMPORT_FAILED`.
func (o CryptoKeyVersionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CryptoKeyVersion) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type CryptoKeyVersionArrayOutput struct{ *pulumi.OutputState }

func (CryptoKeyVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CryptoKeyVersion)(nil)).Elem()
}

func (o CryptoKeyVersionArrayOutput) ToCryptoKeyVersionArrayOutput() CryptoKeyVersionArrayOutput {
	return o
}

func (o CryptoKeyVersionArrayOutput) ToCryptoKeyVersionArrayOutputWithContext(ctx context.Context) CryptoKeyVersionArrayOutput {
	return o
}

func (o CryptoKeyVersionArrayOutput) Index(i pulumi.IntInput) CryptoKeyVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CryptoKeyVersion {
		return vs[0].([]*CryptoKeyVersion)[vs[1].(int)]
	}).(CryptoKeyVersionOutput)
}

type CryptoKeyVersionMapOutput struct{ *pulumi.OutputState }

func (CryptoKeyVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CryptoKeyVersion)(nil)).Elem()
}

func (o CryptoKeyVersionMapOutput) ToCryptoKeyVersionMapOutput() CryptoKeyVersionMapOutput {
	return o
}

func (o CryptoKeyVersionMapOutput) ToCryptoKeyVersionMapOutputWithContext(ctx context.Context) CryptoKeyVersionMapOutput {
	return o
}

func (o CryptoKeyVersionMapOutput) MapIndex(k pulumi.StringInput) CryptoKeyVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CryptoKeyVersion {
		return vs[0].(map[string]*CryptoKeyVersion)[vs[1].(string)]
	}).(CryptoKeyVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CryptoKeyVersionInput)(nil)).Elem(), &CryptoKeyVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*CryptoKeyVersionArrayInput)(nil)).Elem(), CryptoKeyVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CryptoKeyVersionMapInput)(nil)).Elem(), CryptoKeyVersionMap{})
	pulumi.RegisterOutputType(CryptoKeyVersionOutput{})
	pulumi.RegisterOutputType(CryptoKeyVersionArrayOutput{})
	pulumi.RegisterOutputType(CryptoKeyVersionMapOutput{})
}
