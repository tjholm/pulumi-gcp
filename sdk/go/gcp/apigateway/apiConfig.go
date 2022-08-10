// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An API Configuration is an association of an API Controller Config and a Gateway Config
//
// To get more information about ApiConfig, see:
//
// * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis.configs)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/api-gateway/docs/creating-api-config)
//
// ## Example Usage
// ### Apigateway Api Config Basic
//
// ```go
// package main
//
// import (
//
//	"encoding/base64"
//	"io/ioutil"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func filebase64OrPanic(path string) pulumi.StringPtrInput {
//		if fileData, err := ioutil.ReadFile(path); err == nil {
//			return pulumi.String(base64.StdEncoding.EncodeToString(fileData[:]))
//		} else {
//			panic(err.Error())
//		}
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			apiCfgApi, err := apigateway.NewApi(ctx, "apiCfgApi", &apigateway.ApiArgs{
//				ApiId: pulumi.String("api-cfg"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = apigateway.NewApiConfig(ctx, "apiCfgApiConfig", &apigateway.ApiConfigArgs{
//				Api:         apiCfgApi.ApiId,
//				ApiConfigId: pulumi.String("cfg"),
//				OpenapiDocuments: apigateway.ApiConfigOpenapiDocumentArray{
//					&apigateway.ApiConfigOpenapiDocumentArgs{
//						Document: &apigateway.ApiConfigOpenapiDocumentDocumentArgs{
//							Path:     pulumi.String("spec.yaml"),
//							Contents: filebase64OrPanic("test-fixtures/apigateway/openapi.yaml"),
//						},
//					},
//				},
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
// # ApiConfig can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:apigateway/apiConfig:ApiConfig default projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{project}}/{{api}}/{{api_config_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{api}}/{{api_config_id}}
//
// ```
type ApiConfig struct {
	pulumi.CustomResourceState

	// The API to attach the config to.
	Api pulumi.StringOutput `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringOutput `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix pulumi.StringOutput `pulumi:"apiConfigIdPrefix"`
	// A user-visible name for the API.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig ApiConfigGatewayConfigPtrOutput `pulumi:"gatewayConfig"`
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	// Structure is documented below.
	GrpcServices ApiConfigGrpcServiceArrayOutput `pulumi:"grpcServices"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
	// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	// Structure is documented below.
	ManagedServiceConfigs ApiConfigManagedServiceConfigArrayOutput `pulumi:"managedServiceConfigs"`
	// The resource name of the API Config.
	Name pulumi.StringOutput `pulumi:"name"`
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	// Structure is documented below.
	OpenapiDocuments ApiConfigOpenapiDocumentArrayOutput `pulumi:"openapiDocuments"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId pulumi.StringOutput `pulumi:"serviceConfigId"`
}

// NewApiConfig registers a new resource with the given unique name, arguments, and options.
func NewApiConfig(ctx *pulumi.Context,
	name string, args *ApiConfigArgs, opts ...pulumi.ResourceOption) (*ApiConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	var resource ApiConfig
	err := ctx.RegisterResource("gcp:apigateway/apiConfig:ApiConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiConfig gets an existing ApiConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiConfigState, opts ...pulumi.ResourceOption) (*ApiConfig, error) {
	var resource ApiConfig
	err := ctx.ReadResource("gcp:apigateway/apiConfig:ApiConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiConfig resources.
type apiConfigState struct {
	// The API to attach the config to.
	Api *string `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId *string `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix *string `pulumi:"apiConfigIdPrefix"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig *ApiConfigGatewayConfig `pulumi:"gatewayConfig"`
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	// Structure is documented below.
	GrpcServices []ApiConfigGrpcService `pulumi:"grpcServices"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
	// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	// Structure is documented below.
	ManagedServiceConfigs []ApiConfigManagedServiceConfig `pulumi:"managedServiceConfigs"`
	// The resource name of the API Config.
	Name *string `pulumi:"name"`
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	// Structure is documented below.
	OpenapiDocuments []ApiConfigOpenapiDocument `pulumi:"openapiDocuments"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId *string `pulumi:"serviceConfigId"`
}

type ApiConfigState struct {
	// The API to attach the config to.
	Api pulumi.StringPtrInput
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix pulumi.StringPtrInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig ApiConfigGatewayConfigPtrInput
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	// Structure is documented below.
	GrpcServices ApiConfigGrpcServiceArrayInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
	// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	// Structure is documented below.
	ManagedServiceConfigs ApiConfigManagedServiceConfigArrayInput
	// The resource name of the API Config.
	Name pulumi.StringPtrInput
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	// Structure is documented below.
	OpenapiDocuments ApiConfigOpenapiDocumentArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId pulumi.StringPtrInput
}

func (ApiConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigState)(nil)).Elem()
}

type apiConfigArgs struct {
	// The API to attach the config to.
	Api string `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId *string `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix *string `pulumi:"apiConfigIdPrefix"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig *ApiConfigGatewayConfig `pulumi:"gatewayConfig"`
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	// Structure is documented below.
	GrpcServices []ApiConfigGrpcService `pulumi:"grpcServices"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
	// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	// Structure is documented below.
	ManagedServiceConfigs []ApiConfigManagedServiceConfig `pulumi:"managedServiceConfigs"`
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	// Structure is documented below.
	OpenapiDocuments []ApiConfigOpenapiDocument `pulumi:"openapiDocuments"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ApiConfig resource.
type ApiConfigArgs struct {
	// The API to attach the config to.
	Api pulumi.StringInput
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
	ApiConfigIdPrefix pulumi.StringPtrInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Immutable. Gateway specific configuration.
	// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
	// Structure is documented below.
	GatewayConfig ApiConfigGatewayConfigPtrInput
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	// Structure is documented below.
	GrpcServices ApiConfigGrpcServiceArrayInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
	// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
	// Structure is documented below.
	ManagedServiceConfigs ApiConfigManagedServiceConfigArrayInput
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	// Structure is documented below.
	OpenapiDocuments ApiConfigOpenapiDocumentArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigArgs)(nil)).Elem()
}

type ApiConfigInput interface {
	pulumi.Input

	ToApiConfigOutput() ApiConfigOutput
	ToApiConfigOutputWithContext(ctx context.Context) ApiConfigOutput
}

func (*ApiConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfig)(nil)).Elem()
}

func (i *ApiConfig) ToApiConfigOutput() ApiConfigOutput {
	return i.ToApiConfigOutputWithContext(context.Background())
}

func (i *ApiConfig) ToApiConfigOutputWithContext(ctx context.Context) ApiConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigOutput)
}

// ApiConfigArrayInput is an input type that accepts ApiConfigArray and ApiConfigArrayOutput values.
// You can construct a concrete instance of `ApiConfigArrayInput` via:
//
//	ApiConfigArray{ ApiConfigArgs{...} }
type ApiConfigArrayInput interface {
	pulumi.Input

	ToApiConfigArrayOutput() ApiConfigArrayOutput
	ToApiConfigArrayOutputWithContext(context.Context) ApiConfigArrayOutput
}

type ApiConfigArray []ApiConfigInput

func (ApiConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiConfig)(nil)).Elem()
}

func (i ApiConfigArray) ToApiConfigArrayOutput() ApiConfigArrayOutput {
	return i.ToApiConfigArrayOutputWithContext(context.Background())
}

func (i ApiConfigArray) ToApiConfigArrayOutputWithContext(ctx context.Context) ApiConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigArrayOutput)
}

// ApiConfigMapInput is an input type that accepts ApiConfigMap and ApiConfigMapOutput values.
// You can construct a concrete instance of `ApiConfigMapInput` via:
//
//	ApiConfigMap{ "key": ApiConfigArgs{...} }
type ApiConfigMapInput interface {
	pulumi.Input

	ToApiConfigMapOutput() ApiConfigMapOutput
	ToApiConfigMapOutputWithContext(context.Context) ApiConfigMapOutput
}

type ApiConfigMap map[string]ApiConfigInput

func (ApiConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiConfig)(nil)).Elem()
}

func (i ApiConfigMap) ToApiConfigMapOutput() ApiConfigMapOutput {
	return i.ToApiConfigMapOutputWithContext(context.Background())
}

func (i ApiConfigMap) ToApiConfigMapOutputWithContext(ctx context.Context) ApiConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigMapOutput)
}

type ApiConfigOutput struct{ *pulumi.OutputState }

func (ApiConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfig)(nil)).Elem()
}

func (o ApiConfigOutput) ToApiConfigOutput() ApiConfigOutput {
	return o
}

func (o ApiConfigOutput) ToApiConfigOutputWithContext(ctx context.Context) ApiConfigOutput {
	return o
}

// The API to attach the config to.
func (o ApiConfigOutput) Api() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfig) pulumi.StringOutput { return v.Api }).(pulumi.StringOutput)
}

// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
func (o ApiConfigOutput) ApiConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfig) pulumi.StringOutput { return v.ApiConfigId }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the
// specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
func (o ApiConfigOutput) ApiConfigIdPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfig) pulumi.StringOutput { return v.ApiConfigIdPrefix }).(pulumi.StringOutput)
}

// A user-visible name for the API.
func (o ApiConfigOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfig) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Immutable. Gateway specific configuration.
// If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
// Structure is documented below.
func (o ApiConfigOutput) GatewayConfig() ApiConfigGatewayConfigPtrOutput {
	return o.ApplyT(func(v *ApiConfig) ApiConfigGatewayConfigPtrOutput { return v.GatewayConfig }).(ApiConfigGatewayConfigPtrOutput)
}

// gRPC service definition files. If specified, openapiDocuments must not be included.
// Structure is documented below.
func (o ApiConfigOutput) GrpcServices() ApiConfigGrpcServiceArrayOutput {
	return o.ApplyT(func(v *ApiConfig) ApiConfigGrpcServiceArrayOutput { return v.GrpcServices }).(ApiConfigGrpcServiceArrayOutput)
}

// Resource labels to represent user-provided metadata.
func (o ApiConfigOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiConfig) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
// If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
// Structure is documented below.
func (o ApiConfigOutput) ManagedServiceConfigs() ApiConfigManagedServiceConfigArrayOutput {
	return o.ApplyT(func(v *ApiConfig) ApiConfigManagedServiceConfigArrayOutput { return v.ManagedServiceConfigs }).(ApiConfigManagedServiceConfigArrayOutput)
}

// The resource name of the API Config.
func (o ApiConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
// Structure is documented below.
func (o ApiConfigOutput) OpenapiDocuments() ApiConfigOpenapiDocumentArrayOutput {
	return o.ApplyT(func(v *ApiConfig) ApiConfigOpenapiDocumentArrayOutput { return v.OpenapiDocuments }).(ApiConfigOpenapiDocumentArrayOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o ApiConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
func (o ApiConfigOutput) ServiceConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfig) pulumi.StringOutput { return v.ServiceConfigId }).(pulumi.StringOutput)
}

type ApiConfigArrayOutput struct{ *pulumi.OutputState }

func (ApiConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiConfig)(nil)).Elem()
}

func (o ApiConfigArrayOutput) ToApiConfigArrayOutput() ApiConfigArrayOutput {
	return o
}

func (o ApiConfigArrayOutput) ToApiConfigArrayOutputWithContext(ctx context.Context) ApiConfigArrayOutput {
	return o
}

func (o ApiConfigArrayOutput) Index(i pulumi.IntInput) ApiConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiConfig {
		return vs[0].([]*ApiConfig)[vs[1].(int)]
	}).(ApiConfigOutput)
}

type ApiConfigMapOutput struct{ *pulumi.OutputState }

func (ApiConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiConfig)(nil)).Elem()
}

func (o ApiConfigMapOutput) ToApiConfigMapOutput() ApiConfigMapOutput {
	return o
}

func (o ApiConfigMapOutput) ToApiConfigMapOutputWithContext(ctx context.Context) ApiConfigMapOutput {
	return o
}

func (o ApiConfigMapOutput) MapIndex(k pulumi.StringInput) ApiConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiConfig {
		return vs[0].(map[string]*ApiConfig)[vs[1].(string)]
	}).(ApiConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigInput)(nil)).Elem(), &ApiConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigArrayInput)(nil)).Elem(), ApiConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigMapInput)(nil)).Elem(), ApiConfigMap{})
	pulumi.RegisterOutputType(ApiConfigOutput{})
	pulumi.RegisterOutputType(ApiConfigArrayOutput{})
	pulumi.RegisterOutputType(ApiConfigMapOutput{})
}
