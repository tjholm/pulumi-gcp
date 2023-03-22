// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Encrypts secret data with Google Cloud KMS and provides access to the ciphertext.
//
// > **NOTE:** Using this resource will allow you to conceal secret data within your
// resource definitions, but it does not take care of protecting that data in the
// logging output, plan output, or state output.  Please take care to secure your secret
// data outside of resource definitions.
//
// To get more information about SecretCiphertext, see:
//
// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys/encrypt)
// * How-to Guides
//   - [Encrypting and decrypting data with a symmetric key](https://cloud.google.com/kms/docs/encrypt-decrypt)
//
// > **Warning:** All arguments including `plaintext` and `additionalAuthenticatedData` will be stored in the raw state as plain-text.
//
// ## Example Usage
// ### Kms Secret Ciphertext Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
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
//			myPassword, err := kms.NewSecretCiphertext(ctx, "myPassword", &kms.SecretCiphertextArgs{
//				CryptoKey: cryptokey.ID(),
//				Plaintext: pulumi.String("my-secret-password"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstance(ctx, "instance", &compute.InstanceArgs{
//				MachineType: pulumi.String("e2-medium"),
//				Zone:        pulumi.String("us-central1-a"),
//				BootDisk: &compute.InstanceBootDiskArgs{
//					InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
//						Image: pulumi.String("debian-cloud/debian-11"),
//					},
//				},
//				NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
//					&compute.InstanceNetworkInterfaceArgs{
//						Network: pulumi.String("default"),
//						AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
//							nil,
//						},
//					},
//				},
//				Metadata: pulumi.StringMap{
//					"password": myPassword.Ciphertext,
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
// This resource does not support import.
type SecretCiphertext struct {
	pulumi.CustomResourceState

	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData pulumi.StringPtrOutput `pulumi:"additionalAuthenticatedData"`
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext pulumi.StringOutput `pulumi:"ciphertext"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringOutput `pulumi:"cryptoKey"`
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext pulumi.StringOutput `pulumi:"plaintext"`
}

// NewSecretCiphertext registers a new resource with the given unique name, arguments, and options.
func NewSecretCiphertext(ctx *pulumi.Context,
	name string, args *SecretCiphertextArgs, opts ...pulumi.ResourceOption) (*SecretCiphertext, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CryptoKey == nil {
		return nil, errors.New("invalid value for required argument 'CryptoKey'")
	}
	if args.Plaintext == nil {
		return nil, errors.New("invalid value for required argument 'Plaintext'")
	}
	if args.AdditionalAuthenticatedData != nil {
		args.AdditionalAuthenticatedData = pulumi.ToSecret(args.AdditionalAuthenticatedData).(pulumi.StringPtrInput)
	}
	if args.Plaintext != nil {
		args.Plaintext = pulumi.ToSecret(args.Plaintext).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"additionalAuthenticatedData",
		"plaintext",
	})
	opts = append(opts, secrets)
	var resource SecretCiphertext
	err := ctx.RegisterResource("gcp:kms/secretCiphertext:SecretCiphertext", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretCiphertext gets an existing SecretCiphertext resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretCiphertext(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretCiphertextState, opts ...pulumi.ResourceOption) (*SecretCiphertext, error) {
	var resource SecretCiphertext
	err := ctx.ReadResource("gcp:kms/secretCiphertext:SecretCiphertext", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretCiphertext resources.
type secretCiphertextState struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData *string `pulumi:"additionalAuthenticatedData"`
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext *string `pulumi:"ciphertext"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey *string `pulumi:"cryptoKey"`
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext *string `pulumi:"plaintext"`
}

type SecretCiphertextState struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData pulumi.StringPtrInput
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext pulumi.StringPtrInput
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringPtrInput
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext pulumi.StringPtrInput
}

func (SecretCiphertextState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCiphertextState)(nil)).Elem()
}

type secretCiphertextArgs struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData *string `pulumi:"additionalAuthenticatedData"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey string `pulumi:"cryptoKey"`
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext string `pulumi:"plaintext"`
}

// The set of arguments for constructing a SecretCiphertext resource.
type SecretCiphertextArgs struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData pulumi.StringPtrInput
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringInput
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext pulumi.StringInput
}

func (SecretCiphertextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCiphertextArgs)(nil)).Elem()
}

type SecretCiphertextInput interface {
	pulumi.Input

	ToSecretCiphertextOutput() SecretCiphertextOutput
	ToSecretCiphertextOutputWithContext(ctx context.Context) SecretCiphertextOutput
}

func (*SecretCiphertext) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretCiphertext)(nil)).Elem()
}

func (i *SecretCiphertext) ToSecretCiphertextOutput() SecretCiphertextOutput {
	return i.ToSecretCiphertextOutputWithContext(context.Background())
}

func (i *SecretCiphertext) ToSecretCiphertextOutputWithContext(ctx context.Context) SecretCiphertextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCiphertextOutput)
}

// SecretCiphertextArrayInput is an input type that accepts SecretCiphertextArray and SecretCiphertextArrayOutput values.
// You can construct a concrete instance of `SecretCiphertextArrayInput` via:
//
//	SecretCiphertextArray{ SecretCiphertextArgs{...} }
type SecretCiphertextArrayInput interface {
	pulumi.Input

	ToSecretCiphertextArrayOutput() SecretCiphertextArrayOutput
	ToSecretCiphertextArrayOutputWithContext(context.Context) SecretCiphertextArrayOutput
}

type SecretCiphertextArray []SecretCiphertextInput

func (SecretCiphertextArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretCiphertext)(nil)).Elem()
}

func (i SecretCiphertextArray) ToSecretCiphertextArrayOutput() SecretCiphertextArrayOutput {
	return i.ToSecretCiphertextArrayOutputWithContext(context.Background())
}

func (i SecretCiphertextArray) ToSecretCiphertextArrayOutputWithContext(ctx context.Context) SecretCiphertextArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCiphertextArrayOutput)
}

// SecretCiphertextMapInput is an input type that accepts SecretCiphertextMap and SecretCiphertextMapOutput values.
// You can construct a concrete instance of `SecretCiphertextMapInput` via:
//
//	SecretCiphertextMap{ "key": SecretCiphertextArgs{...} }
type SecretCiphertextMapInput interface {
	pulumi.Input

	ToSecretCiphertextMapOutput() SecretCiphertextMapOutput
	ToSecretCiphertextMapOutputWithContext(context.Context) SecretCiphertextMapOutput
}

type SecretCiphertextMap map[string]SecretCiphertextInput

func (SecretCiphertextMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretCiphertext)(nil)).Elem()
}

func (i SecretCiphertextMap) ToSecretCiphertextMapOutput() SecretCiphertextMapOutput {
	return i.ToSecretCiphertextMapOutputWithContext(context.Background())
}

func (i SecretCiphertextMap) ToSecretCiphertextMapOutputWithContext(ctx context.Context) SecretCiphertextMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCiphertextMapOutput)
}

type SecretCiphertextOutput struct{ *pulumi.OutputState }

func (SecretCiphertextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretCiphertext)(nil)).Elem()
}

func (o SecretCiphertextOutput) ToSecretCiphertextOutput() SecretCiphertextOutput {
	return o
}

func (o SecretCiphertextOutput) ToSecretCiphertextOutputWithContext(ctx context.Context) SecretCiphertextOutput {
	return o
}

// The additional authenticated data used for integrity checks during encryption and decryption.
// **Note**: This property is sensitive and will not be displayed in the plan.
func (o SecretCiphertextOutput) AdditionalAuthenticatedData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretCiphertext) pulumi.StringPtrOutput { return v.AdditionalAuthenticatedData }).(pulumi.StringPtrOutput)
}

// Contains the result of encrypting the provided plaintext, encoded in base64.
func (o SecretCiphertextOutput) Ciphertext() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCiphertext) pulumi.StringOutput { return v.Ciphertext }).(pulumi.StringOutput)
}

// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
func (o SecretCiphertextOutput) CryptoKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCiphertext) pulumi.StringOutput { return v.CryptoKey }).(pulumi.StringOutput)
}

// The plaintext to be encrypted.
// **Note**: This property is sensitive and will not be displayed in the plan.
func (o SecretCiphertextOutput) Plaintext() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretCiphertext) pulumi.StringOutput { return v.Plaintext }).(pulumi.StringOutput)
}

type SecretCiphertextArrayOutput struct{ *pulumi.OutputState }

func (SecretCiphertextArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretCiphertext)(nil)).Elem()
}

func (o SecretCiphertextArrayOutput) ToSecretCiphertextArrayOutput() SecretCiphertextArrayOutput {
	return o
}

func (o SecretCiphertextArrayOutput) ToSecretCiphertextArrayOutputWithContext(ctx context.Context) SecretCiphertextArrayOutput {
	return o
}

func (o SecretCiphertextArrayOutput) Index(i pulumi.IntInput) SecretCiphertextOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretCiphertext {
		return vs[0].([]*SecretCiphertext)[vs[1].(int)]
	}).(SecretCiphertextOutput)
}

type SecretCiphertextMapOutput struct{ *pulumi.OutputState }

func (SecretCiphertextMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretCiphertext)(nil)).Elem()
}

func (o SecretCiphertextMapOutput) ToSecretCiphertextMapOutput() SecretCiphertextMapOutput {
	return o
}

func (o SecretCiphertextMapOutput) ToSecretCiphertextMapOutputWithContext(ctx context.Context) SecretCiphertextMapOutput {
	return o
}

func (o SecretCiphertextMapOutput) MapIndex(k pulumi.StringInput) SecretCiphertextOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretCiphertext {
		return vs[0].(map[string]*SecretCiphertext)[vs[1].(string)]
	}).(SecretCiphertextOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCiphertextInput)(nil)).Elem(), &SecretCiphertext{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCiphertextArrayInput)(nil)).Elem(), SecretCiphertextArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretCiphertextMapInput)(nil)).Elem(), SecretCiphertextMap{})
	pulumi.RegisterOutputType(SecretCiphertextOutput{})
	pulumi.RegisterOutputType(SecretCiphertextArrayOutput{})
	pulumi.RegisterOutputType(SecretCiphertextMapOutput{})
}
