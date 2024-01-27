// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Tensorboard is a physical database that stores users' training metrics. A default Tensorboard is provided in each region of a GCP project. If needed users can also create extra Tensorboards in their projects.
//
// To get more information about Tensorboard, see:
//
// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.tensorboards)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/vertex-ai/docs)
//
// ## Example Usage
// ### Vertex Ai Tensorboard
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vertex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vertex.NewAiTensorboard(ctx, "tensorboard", &vertex.AiTensorboardArgs{
//				DisplayName: pulumi.String("terraform"),
//				Description: pulumi.String("sample description"),
//				Labels: pulumi.StringMap{
//					"key1": pulumi.String("value1"),
//					"key2": pulumi.String("value2"),
//				},
//				Region: pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Vertex Ai Tensorboard Full
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/kms"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vertex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			cryptoKey, err := kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
//				CryptoKeyId: pulumi.String("kms-name"),
//				Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
//				Member:      pulumi.String(fmt.Sprintf("serviceAccount:service-%v@gcp-sa-aiplatform.iam.gserviceaccount.com", project.Number)),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vertex.NewAiTensorboard(ctx, "tensorboard", &vertex.AiTensorboardArgs{
//				DisplayName: pulumi.String("terraform"),
//				Description: pulumi.String("sample description"),
//				Labels: pulumi.StringMap{
//					"key1": pulumi.String("value1"),
//					"key2": pulumi.String("value2"),
//				},
//				Region: pulumi.String("us-central1"),
//				EncryptionSpec: &vertex.AiTensorboardEncryptionSpecArgs{
//					KmsKeyName: pulumi.String("kms-name"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				cryptoKey,
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
// Tensorboard can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/tensorboards/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Tensorboard can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiTensorboard:AiTensorboard default projects/{{project}}/locations/{{region}}/tensorboards/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiTensorboard:AiTensorboard default {{project}}/{{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiTensorboard:AiTensorboard default {{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiTensorboard:AiTensorboard default {{name}}
//
// ```
type AiTensorboard struct {
	pulumi.CustomResourceState

	// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not end with a '/'.
	BlobStoragePathPrefix pulumi.StringOutput `pulumi:"blobStoragePathPrefix"`
	// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of this Tensorboard.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User provided name of this Tensorboard.
	//
	// ***
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	// Structure is documented below.
	EncryptionSpec AiTensorboardEncryptionSpecPtrOutput `pulumi:"encryptionSpec"`
	// The labels with user-defined metadata to organize your Tensorboards.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the Tensorboard.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// The region of the tensorboard. eg us-central1
	Region pulumi.StringOutput `pulumi:"region"`
	// The number of Runs stored in this Tensorboard.
	RunCount pulumi.StringOutput `pulumi:"runCount"`
	// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAiTensorboard registers a new resource with the given unique name, arguments, and options.
func NewAiTensorboard(ctx *pulumi.Context,
	name string, args *AiTensorboardArgs, opts ...pulumi.ResourceOption) (*AiTensorboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiTensorboard
	err := ctx.RegisterResource("gcp:vertex/aiTensorboard:AiTensorboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiTensorboard gets an existing AiTensorboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiTensorboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiTensorboardState, opts ...pulumi.ResourceOption) (*AiTensorboard, error) {
	var resource AiTensorboard
	err := ctx.ReadResource("gcp:vertex/aiTensorboard:AiTensorboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiTensorboard resources.
type aiTensorboardState struct {
	// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not end with a '/'.
	BlobStoragePathPrefix *string `pulumi:"blobStoragePathPrefix"`
	// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime *string `pulumi:"createTime"`
	// Description of this Tensorboard.
	Description *string `pulumi:"description"`
	// User provided name of this Tensorboard.
	//
	// ***
	DisplayName *string `pulumi:"displayName"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	// Structure is documented below.
	EncryptionSpec *AiTensorboardEncryptionSpec `pulumi:"encryptionSpec"`
	// The labels with user-defined metadata to organize your Tensorboards.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the Tensorboard.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// The region of the tensorboard. eg us-central1
	Region *string `pulumi:"region"`
	// The number of Runs stored in this Tensorboard.
	RunCount *string `pulumi:"runCount"`
	// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime *string `pulumi:"updateTime"`
}

type AiTensorboardState struct {
	// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not end with a '/'.
	BlobStoragePathPrefix pulumi.StringPtrInput
	// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime pulumi.StringPtrInput
	// Description of this Tensorboard.
	Description pulumi.StringPtrInput
	// User provided name of this Tensorboard.
	//
	// ***
	DisplayName pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapInput
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	// Structure is documented below.
	EncryptionSpec AiTensorboardEncryptionSpecPtrInput
	// The labels with user-defined metadata to organize your Tensorboards.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Name of the Tensorboard.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// The region of the tensorboard. eg us-central1
	Region pulumi.StringPtrInput
	// The number of Runs stored in this Tensorboard.
	RunCount pulumi.StringPtrInput
	// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime pulumi.StringPtrInput
}

func (AiTensorboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiTensorboardState)(nil)).Elem()
}

type aiTensorboardArgs struct {
	// Description of this Tensorboard.
	Description *string `pulumi:"description"`
	// User provided name of this Tensorboard.
	//
	// ***
	DisplayName string `pulumi:"displayName"`
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	// Structure is documented below.
	EncryptionSpec *AiTensorboardEncryptionSpec `pulumi:"encryptionSpec"`
	// The labels with user-defined metadata to organize your Tensorboards.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the tensorboard. eg us-central1
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a AiTensorboard resource.
type AiTensorboardArgs struct {
	// Description of this Tensorboard.
	Description pulumi.StringPtrInput
	// User provided name of this Tensorboard.
	//
	// ***
	DisplayName pulumi.StringInput
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	// Structure is documented below.
	EncryptionSpec AiTensorboardEncryptionSpecPtrInput
	// The labels with user-defined metadata to organize your Tensorboards.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the tensorboard. eg us-central1
	Region pulumi.StringPtrInput
}

func (AiTensorboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiTensorboardArgs)(nil)).Elem()
}

type AiTensorboardInput interface {
	pulumi.Input

	ToAiTensorboardOutput() AiTensorboardOutput
	ToAiTensorboardOutputWithContext(ctx context.Context) AiTensorboardOutput
}

func (*AiTensorboard) ElementType() reflect.Type {
	return reflect.TypeOf((**AiTensorboard)(nil)).Elem()
}

func (i *AiTensorboard) ToAiTensorboardOutput() AiTensorboardOutput {
	return i.ToAiTensorboardOutputWithContext(context.Background())
}

func (i *AiTensorboard) ToAiTensorboardOutputWithContext(ctx context.Context) AiTensorboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiTensorboardOutput)
}

// AiTensorboardArrayInput is an input type that accepts AiTensorboardArray and AiTensorboardArrayOutput values.
// You can construct a concrete instance of `AiTensorboardArrayInput` via:
//
//	AiTensorboardArray{ AiTensorboardArgs{...} }
type AiTensorboardArrayInput interface {
	pulumi.Input

	ToAiTensorboardArrayOutput() AiTensorboardArrayOutput
	ToAiTensorboardArrayOutputWithContext(context.Context) AiTensorboardArrayOutput
}

type AiTensorboardArray []AiTensorboardInput

func (AiTensorboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiTensorboard)(nil)).Elem()
}

func (i AiTensorboardArray) ToAiTensorboardArrayOutput() AiTensorboardArrayOutput {
	return i.ToAiTensorboardArrayOutputWithContext(context.Background())
}

func (i AiTensorboardArray) ToAiTensorboardArrayOutputWithContext(ctx context.Context) AiTensorboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiTensorboardArrayOutput)
}

// AiTensorboardMapInput is an input type that accepts AiTensorboardMap and AiTensorboardMapOutput values.
// You can construct a concrete instance of `AiTensorboardMapInput` via:
//
//	AiTensorboardMap{ "key": AiTensorboardArgs{...} }
type AiTensorboardMapInput interface {
	pulumi.Input

	ToAiTensorboardMapOutput() AiTensorboardMapOutput
	ToAiTensorboardMapOutputWithContext(context.Context) AiTensorboardMapOutput
}

type AiTensorboardMap map[string]AiTensorboardInput

func (AiTensorboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiTensorboard)(nil)).Elem()
}

func (i AiTensorboardMap) ToAiTensorboardMapOutput() AiTensorboardMapOutput {
	return i.ToAiTensorboardMapOutputWithContext(context.Background())
}

func (i AiTensorboardMap) ToAiTensorboardMapOutputWithContext(ctx context.Context) AiTensorboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiTensorboardMapOutput)
}

type AiTensorboardOutput struct{ *pulumi.OutputState }

func (AiTensorboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiTensorboard)(nil)).Elem()
}

func (o AiTensorboardOutput) ToAiTensorboardOutput() AiTensorboardOutput {
	return o
}

func (o AiTensorboardOutput) ToAiTensorboardOutputWithContext(ctx context.Context) AiTensorboardOutput {
	return o
}

// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not end with a '/'.
func (o AiTensorboardOutput) BlobStoragePathPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringOutput { return v.BlobStoragePathPrefix }).(pulumi.StringOutput)
}

// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
func (o AiTensorboardOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of this Tensorboard.
func (o AiTensorboardOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User provided name of this Tensorboard.
//
// ***
func (o AiTensorboardOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o AiTensorboardOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
// Structure is documented below.
func (o AiTensorboardOutput) EncryptionSpec() AiTensorboardEncryptionSpecPtrOutput {
	return o.ApplyT(func(v *AiTensorboard) AiTensorboardEncryptionSpecPtrOutput { return v.EncryptionSpec }).(AiTensorboardEncryptionSpecPtrOutput)
}

// The labels with user-defined metadata to organize your Tensorboards.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o AiTensorboardOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the Tensorboard.
func (o AiTensorboardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o AiTensorboardOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o AiTensorboardOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// The region of the tensorboard. eg us-central1
func (o AiTensorboardOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The number of Runs stored in this Tensorboard.
func (o AiTensorboardOutput) RunCount() pulumi.StringOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringOutput { return v.RunCount }).(pulumi.StringOutput)
}

// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
func (o AiTensorboardOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AiTensorboard) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type AiTensorboardArrayOutput struct{ *pulumi.OutputState }

func (AiTensorboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiTensorboard)(nil)).Elem()
}

func (o AiTensorboardArrayOutput) ToAiTensorboardArrayOutput() AiTensorboardArrayOutput {
	return o
}

func (o AiTensorboardArrayOutput) ToAiTensorboardArrayOutputWithContext(ctx context.Context) AiTensorboardArrayOutput {
	return o
}

func (o AiTensorboardArrayOutput) Index(i pulumi.IntInput) AiTensorboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiTensorboard {
		return vs[0].([]*AiTensorboard)[vs[1].(int)]
	}).(AiTensorboardOutput)
}

type AiTensorboardMapOutput struct{ *pulumi.OutputState }

func (AiTensorboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiTensorboard)(nil)).Elem()
}

func (o AiTensorboardMapOutput) ToAiTensorboardMapOutput() AiTensorboardMapOutput {
	return o
}

func (o AiTensorboardMapOutput) ToAiTensorboardMapOutputWithContext(ctx context.Context) AiTensorboardMapOutput {
	return o
}

func (o AiTensorboardMapOutput) MapIndex(k pulumi.StringInput) AiTensorboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiTensorboard {
		return vs[0].(map[string]*AiTensorboard)[vs[1].(string)]
	}).(AiTensorboardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiTensorboardInput)(nil)).Elem(), &AiTensorboard{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiTensorboardArrayInput)(nil)).Elem(), AiTensorboardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiTensorboardMapInput)(nil)).Elem(), AiTensorboardMap{})
	pulumi.RegisterOutputType(AiTensorboardOutput{})
	pulumi.RegisterOutputType(AiTensorboardArrayOutput{})
	pulumi.RegisterOutputType(AiTensorboardMapOutput{})
}
