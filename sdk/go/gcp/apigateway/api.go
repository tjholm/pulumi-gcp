// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Apigateway Api Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigateway.NewApi(ctx, "api", &apigateway.ApiArgs{
//				ApiId: pulumi.String("api"),
//			}, pulumi.Provider(google_beta))
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
// # Api can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:apigateway/api:Api default projects/{{project}}/locations/global/apis/{{api_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigateway/api:Api default {{project}}/{{api_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigateway/api:Api default {{api_id}}
//
// ```
type Api struct {
	pulumi.CustomResourceState

	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A user-visible name for the API.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService pulumi.StringOutput `pulumi:"managedService"`
	// The resource name of the API. Format `projects/{{project}}/locations/global/apis/{{apiId}}`
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	var resource Api
	err := ctx.RegisterResource("gcp:apigateway/api:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("gcp:apigateway/api:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId *string `pulumi:"apiId"`
	// Creation timestamp in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService *string `pulumi:"managedService"`
	// The resource name of the API. Format `projects/{{project}}/locations/global/apis/{{apiId}}`
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ApiState struct {
	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService pulumi.StringPtrInput
	// The resource name of the API. Format `projects/{{project}}/locations/global/apis/{{apiId}}`
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId string `pulumi:"apiId"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService *string `pulumi:"managedService"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
	ApiId pulumi.StringInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
	// If not specified, a new Service will automatically be created in the same project as this API.
	ManagedService pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

// ApiArrayInput is an input type that accepts ApiArray and ApiArrayOutput values.
// You can construct a concrete instance of `ApiArrayInput` via:
//
//	ApiArray{ ApiArgs{...} }
type ApiArrayInput interface {
	pulumi.Input

	ToApiArrayOutput() ApiArrayOutput
	ToApiArrayOutputWithContext(context.Context) ApiArrayOutput
}

type ApiArray []ApiInput

func (ApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (i ApiArray) ToApiArrayOutput() ApiArrayOutput {
	return i.ToApiArrayOutputWithContext(context.Background())
}

func (i ApiArray) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiArrayOutput)
}

// ApiMapInput is an input type that accepts ApiMap and ApiMapOutput values.
// You can construct a concrete instance of `ApiMapInput` via:
//
//	ApiMap{ "key": ApiArgs{...} }
type ApiMapInput interface {
	pulumi.Input

	ToApiMapOutput() ApiMapOutput
	ToApiMapOutputWithContext(context.Context) ApiMapOutput
}

type ApiMap map[string]ApiInput

func (ApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (i ApiMap) ToApiMapOutput() ApiMapOutput {
	return i.ToApiMapOutputWithContext(context.Background())
}

func (i ApiMap) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMapOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

// Identifier to assign to the API. Must be unique within scope of the parent resource(project)
func (o ApiOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ApiOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A user-visible name for the API.
func (o ApiOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource labels to represent user-provided metadata.
func (o ApiOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Api) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
// If not specified, a new Service will automatically be created in the same project as this API.
func (o ApiOutput) ManagedService() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ManagedService }).(pulumi.StringOutput)
}

// The resource name of the API. Format `projects/{{project}}/locations/global/apis/{{apiId}}`
func (o ApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o ApiOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ApiArrayOutput struct{ *pulumi.OutputState }

func (ApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Api)(nil)).Elem()
}

func (o ApiArrayOutput) ToApiArrayOutput() ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) ToApiArrayOutputWithContext(ctx context.Context) ApiArrayOutput {
	return o
}

func (o ApiArrayOutput) Index(i pulumi.IntInput) ApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Api {
		return vs[0].([]*Api)[vs[1].(int)]
	}).(ApiOutput)
}

type ApiMapOutput struct{ *pulumi.OutputState }

func (ApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Api)(nil)).Elem()
}

func (o ApiMapOutput) ToApiMapOutput() ApiMapOutput {
	return o
}

func (o ApiMapOutput) ToApiMapOutputWithContext(ctx context.Context) ApiMapOutput {
	return o
}

func (o ApiMapOutput) MapIndex(k pulumi.StringInput) ApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Api {
		return vs[0].(map[string]*Api)[vs[1].(string)]
	}).(ApiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiArrayInput)(nil)).Elem(), ApiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiMapInput)(nil)).Elem(), ApiMap{})
	pulumi.RegisterOutputType(ApiOutput{})
	pulumi.RegisterOutputType(ApiArrayOutput{})
	pulumi.RegisterOutputType(ApiMapOutput{})
}
