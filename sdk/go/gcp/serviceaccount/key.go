// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ### Creating A New Key
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myaccount, err := serviceaccount.NewAccount(ctx, "myaccount", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("myaccount"),
//				DisplayName: pulumi.String("My Service Account"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewKey(ctx, "mykey", &serviceaccount.KeyArgs{
//				ServiceAccountId: myaccount.Name,
//				PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creating And Regularly Rotating A Key
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-time/sdk/go/time"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myaccount, err := serviceaccount.NewAccount(ctx, "myaccount", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("myaccount"),
//				DisplayName: pulumi.String("My Service Account"),
//			})
//			if err != nil {
//				return err
//			}
//			mykeyRotation, err := time.NewRotating(ctx, "mykeyRotation", &time.RotatingArgs{
//				RotationDays: pulumi.Int(30),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewKey(ctx, "mykey", &serviceaccount.KeyArgs{
//				ServiceAccountId: myaccount.Name,
//				Keepers: pulumi.Map{
//					"rotation_time": mykeyRotation.RotationRfc3339,
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
type Key struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will trigger a new key to be generated.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm pulumi.StringPtrOutput `pulumi:"keyAlgorithm"`
	// The name used for this key pair
	Name pulumi.StringOutput `pulumi:"name"`
	// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
	// service account keys through the CLI or web console. This is only populated when creating a new key.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType pulumi.StringPtrOutput `pulumi:"privateKeyType"`
	// The public key, base64 encoded
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `publicKeyType` and `privateKeyType`.
	PublicKeyData pulumi.StringPtrOutput `pulumi:"publicKeyData"`
	// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
	PublicKeyType pulumi.StringPtrOutput `pulumi:"publicKeyType"`
	// The Service account id of the Key. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
	// the **full** email address of the service account or its name can be specified as a value, in which case the project will
	// automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
	// syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
	// unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidAfter pulumi.StringOutput `pulumi:"validAfter"`
	// The key can be used before this timestamp.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidBefore pulumi.StringOutput `pulumi:"validBefore"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("gcp:serviceAccount/key:Key"),
		},
	})
	opts = append(opts, aliases)
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Key
	err := ctx.RegisterResource("gcp:serviceaccount/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("gcp:serviceaccount/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// Arbitrary map of values that, when changed, will trigger a new key to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// The name used for this key pair
	Name *string `pulumi:"name"`
	// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
	// service account keys through the CLI or web console. This is only populated when creating a new key.
	PrivateKey *string `pulumi:"privateKey"`
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType *string `pulumi:"privateKeyType"`
	// The public key, base64 encoded
	PublicKey *string `pulumi:"publicKey"`
	// Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `publicKeyType` and `privateKeyType`.
	PublicKeyData *string `pulumi:"publicKeyData"`
	// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
	PublicKeyType *string `pulumi:"publicKeyType"`
	// The Service account id of the Key. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
	// the **full** email address of the service account or its name can be specified as a value, in which case the project will
	// automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
	// syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
	// unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidAfter *string `pulumi:"validAfter"`
	// The key can be used before this timestamp.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidBefore *string `pulumi:"validBefore"`
}

type KeyState struct {
	// Arbitrary map of values that, when changed, will trigger a new key to be generated.
	Keepers pulumi.MapInput
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm pulumi.StringPtrInput
	// The name used for this key pair
	Name pulumi.StringPtrInput
	// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
	// service account keys through the CLI or web console. This is only populated when creating a new key.
	PrivateKey pulumi.StringPtrInput
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType pulumi.StringPtrInput
	// The public key, base64 encoded
	PublicKey pulumi.StringPtrInput
	// Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `publicKeyType` and `privateKeyType`.
	PublicKeyData pulumi.StringPtrInput
	// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
	PublicKeyType pulumi.StringPtrInput
	// The Service account id of the Key. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
	// the **full** email address of the service account or its name can be specified as a value, in which case the project will
	// automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
	// syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
	// unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
	ServiceAccountId pulumi.StringPtrInput
	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidAfter pulumi.StringPtrInput
	// The key can be used before this timestamp.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidBefore pulumi.StringPtrInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Arbitrary map of values that, when changed, will trigger a new key to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType *string `pulumi:"privateKeyType"`
	// Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `publicKeyType` and `privateKeyType`.
	PublicKeyData *string `pulumi:"publicKeyData"`
	// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
	PublicKeyType *string `pulumi:"publicKeyType"`
	// The Service account id of the Key. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
	// the **full** email address of the service account or its name can be specified as a value, in which case the project will
	// automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
	// syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
	// unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Arbitrary map of values that, when changed, will trigger a new key to be generated.
	Keepers pulumi.MapInput
	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	// (only used on create)
	KeyAlgorithm pulumi.StringPtrInput
	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType pulumi.StringPtrInput
	// Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `publicKeyType` and `privateKeyType`.
	PublicKeyData pulumi.StringPtrInput
	// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
	PublicKeyType pulumi.StringPtrInput
	// The Service account id of the Key. This can be a string in the format
	// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
	// the **full** email address of the service account or its name can be specified as a value, in which case the project will
	// automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
	// syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
	// unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
	ServiceAccountId pulumi.StringInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

func (i *Key) ToOutput(ctx context.Context) pulumix.Output[*Key] {
	return pulumix.Output[*Key]{
		OutputState: i.ToKeyOutputWithContext(ctx).OutputState,
	}
}

// KeyArrayInput is an input type that accepts KeyArray and KeyArrayOutput values.
// You can construct a concrete instance of `KeyArrayInput` via:
//
//	KeyArray{ KeyArgs{...} }
type KeyArrayInput interface {
	pulumi.Input

	ToKeyArrayOutput() KeyArrayOutput
	ToKeyArrayOutputWithContext(context.Context) KeyArrayOutput
}

type KeyArray []KeyInput

func (KeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (i KeyArray) ToKeyArrayOutput() KeyArrayOutput {
	return i.ToKeyArrayOutputWithContext(context.Background())
}

func (i KeyArray) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyArrayOutput)
}

func (i KeyArray) ToOutput(ctx context.Context) pulumix.Output[[]*Key] {
	return pulumix.Output[[]*Key]{
		OutputState: i.ToKeyArrayOutputWithContext(ctx).OutputState,
	}
}

// KeyMapInput is an input type that accepts KeyMap and KeyMapOutput values.
// You can construct a concrete instance of `KeyMapInput` via:
//
//	KeyMap{ "key": KeyArgs{...} }
type KeyMapInput interface {
	pulumi.Input

	ToKeyMapOutput() KeyMapOutput
	ToKeyMapOutputWithContext(context.Context) KeyMapOutput
}

type KeyMap map[string]KeyInput

func (KeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (i KeyMap) ToKeyMapOutput() KeyMapOutput {
	return i.ToKeyMapOutputWithContext(context.Background())
}

func (i KeyMap) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyMapOutput)
}

func (i KeyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Key] {
	return pulumix.Output[map[string]*Key]{
		OutputState: i.ToKeyMapOutputWithContext(ctx).OutputState,
	}
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func (o KeyOutput) ToOutput(ctx context.Context) pulumix.Output[*Key] {
	return pulumix.Output[*Key]{
		OutputState: o.OutputState,
	}
}

// Arbitrary map of values that, when changed, will trigger a new key to be generated.
func (o KeyOutput) Keepers() pulumi.MapOutput {
	return o.ApplyT(func(v *Key) pulumi.MapOutput { return v.Keepers }).(pulumi.MapOutput)
}

// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
// Valid values are listed at
// [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
// (only used on create)
func (o KeyOutput) KeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.KeyAlgorithm }).(pulumi.StringPtrOutput)
}

// The name used for this key pair
func (o KeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
// service account keys through the CLI or web console. This is only populated when creating a new key.
func (o KeyOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
func (o KeyOutput) PrivateKeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.PrivateKeyType }).(pulumi.StringPtrOutput)
}

// The public key, base64 encoded
func (o KeyOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `publicKeyType` and `privateKeyType`.
func (o KeyOutput) PublicKeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.PublicKeyData }).(pulumi.StringPtrOutput)
}

// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
func (o KeyOutput) PublicKeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.PublicKeyType }).(pulumi.StringPtrOutput)
}

// The Service account id of the Key. This can be a string in the format
// `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
// the **full** email address of the service account or its name can be specified as a value, in which case the project will
// automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
// syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
// unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
func (o KeyOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.ServiceAccountId }).(pulumi.StringOutput)
}

// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (o KeyOutput) ValidAfter() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.ValidAfter }).(pulumi.StringOutput)
}

// The key can be used before this timestamp.
// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (o KeyOutput) ValidBefore() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.ValidBefore }).(pulumi.StringOutput)
}

type KeyArrayOutput struct{ *pulumi.OutputState }

func (KeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (o KeyArrayOutput) ToKeyArrayOutput() KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Key] {
	return pulumix.Output[[]*Key]{
		OutputState: o.OutputState,
	}
}

func (o KeyArrayOutput) Index(i pulumi.IntInput) KeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Key {
		return vs[0].([]*Key)[vs[1].(int)]
	}).(KeyOutput)
}

type KeyMapOutput struct{ *pulumi.OutputState }

func (KeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (o KeyMapOutput) ToKeyMapOutput() KeyMapOutput {
	return o
}

func (o KeyMapOutput) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return o
}

func (o KeyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Key] {
	return pulumix.Output[map[string]*Key]{
		OutputState: o.OutputState,
	}
}

func (o KeyMapOutput) MapIndex(k pulumi.StringInput) KeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Key {
		return vs[0].(map[string]*Key)[vs[1].(string)]
	}).(KeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyArrayInput)(nil)).Elem(), KeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyMapInput)(nil)).Elem(), KeyMap{})
	pulumi.RegisterOutputType(KeyOutput{})
	pulumi.RegisterOutputType(KeyArrayOutput{})
	pulumi.RegisterOutputType(KeyMapOutput{})
}
