// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package essentialcontacts

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The first-class citizen for Document AI. Each processor defines how to extract structural information from a document.
//
// To get more information about Processor, see:
//
// * [API documentation](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations.processors)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/document-ai/docs/overview)
//
// ## Example Usage
// ### Documentai Processor
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/essentialcontacts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := essentialcontacts.NewDocumentAiProcessor(ctx, "processor", &essentialcontacts.DocumentAiProcessorArgs{
//				DisplayName: pulumi.String("test-processor"),
//				Location:    pulumi.String("us"),
//				Type:        pulumi.String("OCR_PROCESSOR"),
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
// # Processor can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor default projects/{{project}}/locations/{{location}}/processors/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor default {{location}}/{{name}}
//
// ```
type DocumentAiProcessor struct {
	pulumi.CustomResourceState

	// The display name. Must be unique.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName pulumi.StringPtrOutput `pulumi:"kmsKeyName"`
	// The location of the resource.
	//
	// ***
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the processor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDocumentAiProcessor registers a new resource with the given unique name, arguments, and options.
func NewDocumentAiProcessor(ctx *pulumi.Context,
	name string, args *DocumentAiProcessorArgs, opts ...pulumi.ResourceOption) (*DocumentAiProcessor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentAiProcessor
	err := ctx.RegisterResource("gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentAiProcessor gets an existing DocumentAiProcessor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentAiProcessor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentAiProcessorState, opts ...pulumi.ResourceOption) (*DocumentAiProcessor, error) {
	var resource DocumentAiProcessor
	err := ctx.ReadResource("gcp:essentialcontacts/documentAiProcessor:DocumentAiProcessor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentAiProcessor resources.
type documentAiProcessorState struct {
	// The display name. Must be unique.
	DisplayName *string `pulumi:"displayName"`
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// The location of the resource.
	//
	// ***
	Location *string `pulumi:"location"`
	// The resource name of the processor.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type *string `pulumi:"type"`
}

type DocumentAiProcessorState struct {
	// The display name. Must be unique.
	DisplayName pulumi.StringPtrInput
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName pulumi.StringPtrInput
	// The location of the resource.
	//
	// ***
	Location pulumi.StringPtrInput
	// The resource name of the processor.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type pulumi.StringPtrInput
}

func (DocumentAiProcessorState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentAiProcessorState)(nil)).Elem()
}

type documentAiProcessorArgs struct {
	// The display name. Must be unique.
	DisplayName string `pulumi:"displayName"`
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// The location of the resource.
	//
	// ***
	Location string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DocumentAiProcessor resource.
type DocumentAiProcessorArgs struct {
	// The display name. Must be unique.
	DisplayName pulumi.StringInput
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName pulumi.StringPtrInput
	// The location of the resource.
	//
	// ***
	Location pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type pulumi.StringInput
}

func (DocumentAiProcessorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentAiProcessorArgs)(nil)).Elem()
}

type DocumentAiProcessorInput interface {
	pulumi.Input

	ToDocumentAiProcessorOutput() DocumentAiProcessorOutput
	ToDocumentAiProcessorOutputWithContext(ctx context.Context) DocumentAiProcessorOutput
}

func (*DocumentAiProcessor) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentAiProcessor)(nil)).Elem()
}

func (i *DocumentAiProcessor) ToDocumentAiProcessorOutput() DocumentAiProcessorOutput {
	return i.ToDocumentAiProcessorOutputWithContext(context.Background())
}

func (i *DocumentAiProcessor) ToDocumentAiProcessorOutputWithContext(ctx context.Context) DocumentAiProcessorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentAiProcessorOutput)
}

func (i *DocumentAiProcessor) ToOutput(ctx context.Context) pulumix.Output[*DocumentAiProcessor] {
	return pulumix.Output[*DocumentAiProcessor]{
		OutputState: i.ToDocumentAiProcessorOutputWithContext(ctx).OutputState,
	}
}

// DocumentAiProcessorArrayInput is an input type that accepts DocumentAiProcessorArray and DocumentAiProcessorArrayOutput values.
// You can construct a concrete instance of `DocumentAiProcessorArrayInput` via:
//
//	DocumentAiProcessorArray{ DocumentAiProcessorArgs{...} }
type DocumentAiProcessorArrayInput interface {
	pulumi.Input

	ToDocumentAiProcessorArrayOutput() DocumentAiProcessorArrayOutput
	ToDocumentAiProcessorArrayOutputWithContext(context.Context) DocumentAiProcessorArrayOutput
}

type DocumentAiProcessorArray []DocumentAiProcessorInput

func (DocumentAiProcessorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentAiProcessor)(nil)).Elem()
}

func (i DocumentAiProcessorArray) ToDocumentAiProcessorArrayOutput() DocumentAiProcessorArrayOutput {
	return i.ToDocumentAiProcessorArrayOutputWithContext(context.Background())
}

func (i DocumentAiProcessorArray) ToDocumentAiProcessorArrayOutputWithContext(ctx context.Context) DocumentAiProcessorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentAiProcessorArrayOutput)
}

func (i DocumentAiProcessorArray) ToOutput(ctx context.Context) pulumix.Output[[]*DocumentAiProcessor] {
	return pulumix.Output[[]*DocumentAiProcessor]{
		OutputState: i.ToDocumentAiProcessorArrayOutputWithContext(ctx).OutputState,
	}
}

// DocumentAiProcessorMapInput is an input type that accepts DocumentAiProcessorMap and DocumentAiProcessorMapOutput values.
// You can construct a concrete instance of `DocumentAiProcessorMapInput` via:
//
//	DocumentAiProcessorMap{ "key": DocumentAiProcessorArgs{...} }
type DocumentAiProcessorMapInput interface {
	pulumi.Input

	ToDocumentAiProcessorMapOutput() DocumentAiProcessorMapOutput
	ToDocumentAiProcessorMapOutputWithContext(context.Context) DocumentAiProcessorMapOutput
}

type DocumentAiProcessorMap map[string]DocumentAiProcessorInput

func (DocumentAiProcessorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentAiProcessor)(nil)).Elem()
}

func (i DocumentAiProcessorMap) ToDocumentAiProcessorMapOutput() DocumentAiProcessorMapOutput {
	return i.ToDocumentAiProcessorMapOutputWithContext(context.Background())
}

func (i DocumentAiProcessorMap) ToDocumentAiProcessorMapOutputWithContext(ctx context.Context) DocumentAiProcessorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentAiProcessorMapOutput)
}

func (i DocumentAiProcessorMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DocumentAiProcessor] {
	return pulumix.Output[map[string]*DocumentAiProcessor]{
		OutputState: i.ToDocumentAiProcessorMapOutputWithContext(ctx).OutputState,
	}
}

type DocumentAiProcessorOutput struct{ *pulumi.OutputState }

func (DocumentAiProcessorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentAiProcessor)(nil)).Elem()
}

func (o DocumentAiProcessorOutput) ToDocumentAiProcessorOutput() DocumentAiProcessorOutput {
	return o
}

func (o DocumentAiProcessorOutput) ToDocumentAiProcessorOutputWithContext(ctx context.Context) DocumentAiProcessorOutput {
	return o
}

func (o DocumentAiProcessorOutput) ToOutput(ctx context.Context) pulumix.Output[*DocumentAiProcessor] {
	return pulumix.Output[*DocumentAiProcessor]{
		OutputState: o.OutputState,
	}
}

// The display name. Must be unique.
func (o DocumentAiProcessorOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
func (o DocumentAiProcessorOutput) KmsKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringPtrOutput { return v.KmsKeyName }).(pulumi.StringPtrOutput)
}

// The location of the resource.
//
// ***
func (o DocumentAiProcessorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the processor.
func (o DocumentAiProcessorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DocumentAiProcessorOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
func (o DocumentAiProcessorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type DocumentAiProcessorArrayOutput struct{ *pulumi.OutputState }

func (DocumentAiProcessorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentAiProcessor)(nil)).Elem()
}

func (o DocumentAiProcessorArrayOutput) ToDocumentAiProcessorArrayOutput() DocumentAiProcessorArrayOutput {
	return o
}

func (o DocumentAiProcessorArrayOutput) ToDocumentAiProcessorArrayOutputWithContext(ctx context.Context) DocumentAiProcessorArrayOutput {
	return o
}

func (o DocumentAiProcessorArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DocumentAiProcessor] {
	return pulumix.Output[[]*DocumentAiProcessor]{
		OutputState: o.OutputState,
	}
}

func (o DocumentAiProcessorArrayOutput) Index(i pulumi.IntInput) DocumentAiProcessorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentAiProcessor {
		return vs[0].([]*DocumentAiProcessor)[vs[1].(int)]
	}).(DocumentAiProcessorOutput)
}

type DocumentAiProcessorMapOutput struct{ *pulumi.OutputState }

func (DocumentAiProcessorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentAiProcessor)(nil)).Elem()
}

func (o DocumentAiProcessorMapOutput) ToDocumentAiProcessorMapOutput() DocumentAiProcessorMapOutput {
	return o
}

func (o DocumentAiProcessorMapOutput) ToDocumentAiProcessorMapOutputWithContext(ctx context.Context) DocumentAiProcessorMapOutput {
	return o
}

func (o DocumentAiProcessorMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DocumentAiProcessor] {
	return pulumix.Output[map[string]*DocumentAiProcessor]{
		OutputState: o.OutputState,
	}
}

func (o DocumentAiProcessorMapOutput) MapIndex(k pulumi.StringInput) DocumentAiProcessorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentAiProcessor {
		return vs[0].(map[string]*DocumentAiProcessor)[vs[1].(string)]
	}).(DocumentAiProcessorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentAiProcessorInput)(nil)).Elem(), &DocumentAiProcessor{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentAiProcessorArrayInput)(nil)).Elem(), DocumentAiProcessorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentAiProcessorMapInput)(nil)).Elem(), DocumentAiProcessorMap{})
	pulumi.RegisterOutputType(DocumentAiProcessorOutput{})
	pulumi.RegisterOutputType(DocumentAiProcessorArrayOutput{})
	pulumi.RegisterOutputType(DocumentAiProcessorMapOutput{})
}
