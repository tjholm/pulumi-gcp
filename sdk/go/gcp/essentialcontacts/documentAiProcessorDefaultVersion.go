// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package essentialcontacts

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The default version for the processor. Deleting this resource is a no-op, and does not unset the default version.
//
// ## Example Usage
// ### Documentai Default Version
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/essentialcontacts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			processorDocumentAiProcessor, err := essentialcontacts.NewDocumentAiProcessor(ctx, "processorDocumentAiProcessor", &essentialcontacts.DocumentAiProcessorArgs{
//				Location:    pulumi.String("us"),
//				DisplayName: pulumi.String("test-processor"),
//				Type:        pulumi.String("OCR_PROCESSOR"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = essentialcontacts.NewDocumentAiProcessorDefaultVersion(ctx, "processorDocumentAiProcessorDefaultVersion", &essentialcontacts.DocumentAiProcessorDefaultVersionArgs{
//				Processor: processorDocumentAiProcessor.ID(),
//				Version: processorDocumentAiProcessor.ID().ApplyT(func(id string) (string, error) {
//					return fmt.Sprintf("%v/processorVersions/stable", id), nil
//				}).(pulumi.StringOutput),
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
// ProcessorDefaultVersion can be imported using any of these accepted formats* `{{processor}}` When using the `pulumi import` command, ProcessorDefaultVersion can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:essentialcontacts/documentAiProcessorDefaultVersion:DocumentAiProcessorDefaultVersion default {{processor}}
//
// ```
type DocumentAiProcessorDefaultVersion struct {
	pulumi.CustomResourceState

	// The processor to set the version on.
	//
	// ***
	Processor pulumi.StringOutput `pulumi:"processor"`
	// The version to set. Using `stable` or `rc` will cause the API to return the latest version in that release channel.
	// Apply `lifecycle.ignore_changes` to the `version` field to suppress this diff.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewDocumentAiProcessorDefaultVersion registers a new resource with the given unique name, arguments, and options.
func NewDocumentAiProcessorDefaultVersion(ctx *pulumi.Context,
	name string, args *DocumentAiProcessorDefaultVersionArgs, opts ...pulumi.ResourceOption) (*DocumentAiProcessorDefaultVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Processor == nil {
		return nil, errors.New("invalid value for required argument 'Processor'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentAiProcessorDefaultVersion
	err := ctx.RegisterResource("gcp:essentialcontacts/documentAiProcessorDefaultVersion:DocumentAiProcessorDefaultVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentAiProcessorDefaultVersion gets an existing DocumentAiProcessorDefaultVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentAiProcessorDefaultVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentAiProcessorDefaultVersionState, opts ...pulumi.ResourceOption) (*DocumentAiProcessorDefaultVersion, error) {
	var resource DocumentAiProcessorDefaultVersion
	err := ctx.ReadResource("gcp:essentialcontacts/documentAiProcessorDefaultVersion:DocumentAiProcessorDefaultVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentAiProcessorDefaultVersion resources.
type documentAiProcessorDefaultVersionState struct {
	// The processor to set the version on.
	//
	// ***
	Processor *string `pulumi:"processor"`
	// The version to set. Using `stable` or `rc` will cause the API to return the latest version in that release channel.
	// Apply `lifecycle.ignore_changes` to the `version` field to suppress this diff.
	Version *string `pulumi:"version"`
}

type DocumentAiProcessorDefaultVersionState struct {
	// The processor to set the version on.
	//
	// ***
	Processor pulumi.StringPtrInput
	// The version to set. Using `stable` or `rc` will cause the API to return the latest version in that release channel.
	// Apply `lifecycle.ignore_changes` to the `version` field to suppress this diff.
	Version pulumi.StringPtrInput
}

func (DocumentAiProcessorDefaultVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentAiProcessorDefaultVersionState)(nil)).Elem()
}

type documentAiProcessorDefaultVersionArgs struct {
	// The processor to set the version on.
	//
	// ***
	Processor string `pulumi:"processor"`
	// The version to set. Using `stable` or `rc` will cause the API to return the latest version in that release channel.
	// Apply `lifecycle.ignore_changes` to the `version` field to suppress this diff.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a DocumentAiProcessorDefaultVersion resource.
type DocumentAiProcessorDefaultVersionArgs struct {
	// The processor to set the version on.
	//
	// ***
	Processor pulumi.StringInput
	// The version to set. Using `stable` or `rc` will cause the API to return the latest version in that release channel.
	// Apply `lifecycle.ignore_changes` to the `version` field to suppress this diff.
	Version pulumi.StringInput
}

func (DocumentAiProcessorDefaultVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentAiProcessorDefaultVersionArgs)(nil)).Elem()
}

type DocumentAiProcessorDefaultVersionInput interface {
	pulumi.Input

	ToDocumentAiProcessorDefaultVersionOutput() DocumentAiProcessorDefaultVersionOutput
	ToDocumentAiProcessorDefaultVersionOutputWithContext(ctx context.Context) DocumentAiProcessorDefaultVersionOutput
}

func (*DocumentAiProcessorDefaultVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentAiProcessorDefaultVersion)(nil)).Elem()
}

func (i *DocumentAiProcessorDefaultVersion) ToDocumentAiProcessorDefaultVersionOutput() DocumentAiProcessorDefaultVersionOutput {
	return i.ToDocumentAiProcessorDefaultVersionOutputWithContext(context.Background())
}

func (i *DocumentAiProcessorDefaultVersion) ToDocumentAiProcessorDefaultVersionOutputWithContext(ctx context.Context) DocumentAiProcessorDefaultVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentAiProcessorDefaultVersionOutput)
}

// DocumentAiProcessorDefaultVersionArrayInput is an input type that accepts DocumentAiProcessorDefaultVersionArray and DocumentAiProcessorDefaultVersionArrayOutput values.
// You can construct a concrete instance of `DocumentAiProcessorDefaultVersionArrayInput` via:
//
//	DocumentAiProcessorDefaultVersionArray{ DocumentAiProcessorDefaultVersionArgs{...} }
type DocumentAiProcessorDefaultVersionArrayInput interface {
	pulumi.Input

	ToDocumentAiProcessorDefaultVersionArrayOutput() DocumentAiProcessorDefaultVersionArrayOutput
	ToDocumentAiProcessorDefaultVersionArrayOutputWithContext(context.Context) DocumentAiProcessorDefaultVersionArrayOutput
}

type DocumentAiProcessorDefaultVersionArray []DocumentAiProcessorDefaultVersionInput

func (DocumentAiProcessorDefaultVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentAiProcessorDefaultVersion)(nil)).Elem()
}

func (i DocumentAiProcessorDefaultVersionArray) ToDocumentAiProcessorDefaultVersionArrayOutput() DocumentAiProcessorDefaultVersionArrayOutput {
	return i.ToDocumentAiProcessorDefaultVersionArrayOutputWithContext(context.Background())
}

func (i DocumentAiProcessorDefaultVersionArray) ToDocumentAiProcessorDefaultVersionArrayOutputWithContext(ctx context.Context) DocumentAiProcessorDefaultVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentAiProcessorDefaultVersionArrayOutput)
}

// DocumentAiProcessorDefaultVersionMapInput is an input type that accepts DocumentAiProcessorDefaultVersionMap and DocumentAiProcessorDefaultVersionMapOutput values.
// You can construct a concrete instance of `DocumentAiProcessorDefaultVersionMapInput` via:
//
//	DocumentAiProcessorDefaultVersionMap{ "key": DocumentAiProcessorDefaultVersionArgs{...} }
type DocumentAiProcessorDefaultVersionMapInput interface {
	pulumi.Input

	ToDocumentAiProcessorDefaultVersionMapOutput() DocumentAiProcessorDefaultVersionMapOutput
	ToDocumentAiProcessorDefaultVersionMapOutputWithContext(context.Context) DocumentAiProcessorDefaultVersionMapOutput
}

type DocumentAiProcessorDefaultVersionMap map[string]DocumentAiProcessorDefaultVersionInput

func (DocumentAiProcessorDefaultVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentAiProcessorDefaultVersion)(nil)).Elem()
}

func (i DocumentAiProcessorDefaultVersionMap) ToDocumentAiProcessorDefaultVersionMapOutput() DocumentAiProcessorDefaultVersionMapOutput {
	return i.ToDocumentAiProcessorDefaultVersionMapOutputWithContext(context.Background())
}

func (i DocumentAiProcessorDefaultVersionMap) ToDocumentAiProcessorDefaultVersionMapOutputWithContext(ctx context.Context) DocumentAiProcessorDefaultVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentAiProcessorDefaultVersionMapOutput)
}

type DocumentAiProcessorDefaultVersionOutput struct{ *pulumi.OutputState }

func (DocumentAiProcessorDefaultVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentAiProcessorDefaultVersion)(nil)).Elem()
}

func (o DocumentAiProcessorDefaultVersionOutput) ToDocumentAiProcessorDefaultVersionOutput() DocumentAiProcessorDefaultVersionOutput {
	return o
}

func (o DocumentAiProcessorDefaultVersionOutput) ToDocumentAiProcessorDefaultVersionOutputWithContext(ctx context.Context) DocumentAiProcessorDefaultVersionOutput {
	return o
}

// The processor to set the version on.
//
// ***
func (o DocumentAiProcessorDefaultVersionOutput) Processor() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessorDefaultVersion) pulumi.StringOutput { return v.Processor }).(pulumi.StringOutput)
}

// The version to set. Using `stable` or `rc` will cause the API to return the latest version in that release channel.
// Apply `lifecycle.ignore_changes` to the `version` field to suppress this diff.
func (o DocumentAiProcessorDefaultVersionOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessorDefaultVersion) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type DocumentAiProcessorDefaultVersionArrayOutput struct{ *pulumi.OutputState }

func (DocumentAiProcessorDefaultVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentAiProcessorDefaultVersion)(nil)).Elem()
}

func (o DocumentAiProcessorDefaultVersionArrayOutput) ToDocumentAiProcessorDefaultVersionArrayOutput() DocumentAiProcessorDefaultVersionArrayOutput {
	return o
}

func (o DocumentAiProcessorDefaultVersionArrayOutput) ToDocumentAiProcessorDefaultVersionArrayOutputWithContext(ctx context.Context) DocumentAiProcessorDefaultVersionArrayOutput {
	return o
}

func (o DocumentAiProcessorDefaultVersionArrayOutput) Index(i pulumi.IntInput) DocumentAiProcessorDefaultVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentAiProcessorDefaultVersion {
		return vs[0].([]*DocumentAiProcessorDefaultVersion)[vs[1].(int)]
	}).(DocumentAiProcessorDefaultVersionOutput)
}

type DocumentAiProcessorDefaultVersionMapOutput struct{ *pulumi.OutputState }

func (DocumentAiProcessorDefaultVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentAiProcessorDefaultVersion)(nil)).Elem()
}

func (o DocumentAiProcessorDefaultVersionMapOutput) ToDocumentAiProcessorDefaultVersionMapOutput() DocumentAiProcessorDefaultVersionMapOutput {
	return o
}

func (o DocumentAiProcessorDefaultVersionMapOutput) ToDocumentAiProcessorDefaultVersionMapOutputWithContext(ctx context.Context) DocumentAiProcessorDefaultVersionMapOutput {
	return o
}

func (o DocumentAiProcessorDefaultVersionMapOutput) MapIndex(k pulumi.StringInput) DocumentAiProcessorDefaultVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentAiProcessorDefaultVersion {
		return vs[0].(map[string]*DocumentAiProcessorDefaultVersion)[vs[1].(string)]
	}).(DocumentAiProcessorDefaultVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentAiProcessorDefaultVersionInput)(nil)).Elem(), &DocumentAiProcessorDefaultVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentAiProcessorDefaultVersionArrayInput)(nil)).Elem(), DocumentAiProcessorDefaultVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentAiProcessorDefaultVersionMapInput)(nil)).Elem(), DocumentAiProcessorDefaultVersionMap{})
	pulumi.RegisterOutputType(DocumentAiProcessorDefaultVersionOutput{})
	pulumi.RegisterOutputType(DocumentAiProcessorDefaultVersionArrayOutput{})
	pulumi.RegisterOutputType(DocumentAiProcessorDefaultVersionMapOutput{})
}
