// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Models are deployed into it, and afterwards Endpoint is called to obtain predictions and explanations.
//
// To get more information about Endpoint, see:
//
// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.endpoints)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/vertex-ai/docs)
//
// ## Example Usage
// ### Vertex Ai Endpoint Network
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/kms"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/servicenetworking"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vertex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vertexNetwork, err := compute.NewNetwork(ctx, "vertexNetwork", nil)
//			if err != nil {
//				return err
//			}
//			vertexRange, err := compute.NewGlobalAddress(ctx, "vertexRange", &compute.GlobalAddressArgs{
//				Purpose:      pulumi.String("VPC_PEERING"),
//				AddressType:  pulumi.String("INTERNAL"),
//				PrefixLength: pulumi.Int(24),
//				Network:      vertexNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			vertexVpcConnection, err := servicenetworking.NewConnection(ctx, "vertexVpcConnection", &servicenetworking.ConnectionArgs{
//				Network: vertexNetwork.ID(),
//				Service: pulumi.String("servicenetworking.googleapis.com"),
//				ReservedPeeringRanges: pulumi.StringArray{
//					vertexRange.Name,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vertex.NewAiEndpoint(ctx, "endpoint", &vertex.AiEndpointArgs{
//				DisplayName: pulumi.String("sample-endpoint"),
//				Description: pulumi.String("A sample vertex endpoint"),
//				Location:    pulumi.String("us-central1"),
//				Region:      pulumi.String("us-central1"),
//				Labels: pulumi.StringMap{
//					"label-one": pulumi.String("value-one"),
//				},
//				Network: vertexNetwork.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("projects/%v/global/networks/%v", project.Number, name), nil
//				}).(pulumi.StringOutput),
//				EncryptionSpec: &vertex.AiEndpointEncryptionSpecArgs{
//					KmsKeyName: pulumi.String("kms-name"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				vertexVpcConnection,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
//				CryptoKeyId: pulumi.String("kms-name"),
//				Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
//				Member:      pulumi.String(fmt.Sprintf("serviceAccount:service-%v@gcp-sa-aiplatform.iam.gserviceaccount.com", project.Number)),
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
// # Endpoint can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiEndpoint:AiEndpoint default projects/{{project}}/locations/{{location}}/endpoints/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiEndpoint:AiEndpoint default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiEndpoint:AiEndpoint default {{location}}/{{name}}
//
// ```
type AiEndpoint struct {
	pulumi.CustomResourceState

	// (Output)
	// Output only. Timestamp when the DeployedModel was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
	// Structure is documented below.
	DeployedModels AiEndpointDeployedModelArrayOutput `pulumi:"deployedModels"`
	// The description of the Endpoint.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
	// Structure is documented below.
	EncryptionSpec AiEndpointEncryptionSpecPtrOutput `pulumi:"encryptionSpec"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	//
	// ***
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
	ModelDeploymentMonitoringJob pulumi.StringOutput `pulumi:"modelDeploymentMonitoringJob"`
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name pulumi.StringOutput `pulumi:"name"`
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// The region for the resource
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// Output only. Timestamp when this Endpoint was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAiEndpoint registers a new resource with the given unique name, arguments, and options.
func NewAiEndpoint(ctx *pulumi.Context,
	name string, args *AiEndpointArgs, opts ...pulumi.ResourceOption) (*AiEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiEndpoint
	err := ctx.RegisterResource("gcp:vertex/aiEndpoint:AiEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiEndpoint gets an existing AiEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiEndpointState, opts ...pulumi.ResourceOption) (*AiEndpoint, error) {
	var resource AiEndpoint
	err := ctx.ReadResource("gcp:vertex/aiEndpoint:AiEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiEndpoint resources.
type aiEndpointState struct {
	// (Output)
	// Output only. Timestamp when the DeployedModel was created.
	CreateTime *string `pulumi:"createTime"`
	// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
	// Structure is documented below.
	DeployedModels []AiEndpointDeployedModel `pulumi:"deployedModels"`
	// The description of the Endpoint.
	Description *string `pulumi:"description"`
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName *string `pulumi:"displayName"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
	// Structure is documented below.
	EncryptionSpec *AiEndpointEncryptionSpec `pulumi:"encryptionSpec"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag *string `pulumi:"etag"`
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	//
	// ***
	Location *string `pulumi:"location"`
	// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
	ModelDeploymentMonitoringJob *string `pulumi:"modelDeploymentMonitoringJob"`
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name *string `pulumi:"name"`
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// The region for the resource
	Region *string `pulumi:"region"`
	// Output only. Timestamp when this Endpoint was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type AiEndpointState struct {
	// (Output)
	// Output only. Timestamp when the DeployedModel was created.
	CreateTime pulumi.StringPtrInput
	// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
	// Structure is documented below.
	DeployedModels AiEndpointDeployedModelArrayInput
	// The description of the Endpoint.
	Description pulumi.StringPtrInput
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapInput
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
	// Structure is documented below.
	EncryptionSpec AiEndpointEncryptionSpecPtrInput
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag pulumi.StringPtrInput
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	//
	// ***
	Location pulumi.StringPtrInput
	// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
	ModelDeploymentMonitoringJob pulumi.StringPtrInput
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name pulumi.StringPtrInput
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// The region for the resource
	Region pulumi.StringPtrInput
	// Output only. Timestamp when this Endpoint was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (AiEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiEndpointState)(nil)).Elem()
}

type aiEndpointArgs struct {
	// The description of the Endpoint.
	Description *string `pulumi:"description"`
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `pulumi:"displayName"`
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
	// Structure is documented below.
	EncryptionSpec *AiEndpointEncryptionSpec `pulumi:"encryptionSpec"`
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	//
	// ***
	Location string `pulumi:"location"`
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name *string `pulumi:"name"`
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region for the resource
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a AiEndpoint resource.
type AiEndpointArgs struct {
	// The description of the Endpoint.
	Description pulumi.StringPtrInput
	// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName pulumi.StringInput
	// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
	// Structure is documented below.
	EncryptionSpec AiEndpointEncryptionSpecPtrInput
	// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	//
	// ***
	Location pulumi.StringInput
	// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
	Name pulumi.StringPtrInput
	// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region for the resource
	Region pulumi.StringPtrInput
}

func (AiEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiEndpointArgs)(nil)).Elem()
}

type AiEndpointInput interface {
	pulumi.Input

	ToAiEndpointOutput() AiEndpointOutput
	ToAiEndpointOutputWithContext(ctx context.Context) AiEndpointOutput
}

func (*AiEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**AiEndpoint)(nil)).Elem()
}

func (i *AiEndpoint) ToAiEndpointOutput() AiEndpointOutput {
	return i.ToAiEndpointOutputWithContext(context.Background())
}

func (i *AiEndpoint) ToAiEndpointOutputWithContext(ctx context.Context) AiEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiEndpointOutput)
}

func (i *AiEndpoint) ToOutput(ctx context.Context) pulumix.Output[*AiEndpoint] {
	return pulumix.Output[*AiEndpoint]{
		OutputState: i.ToAiEndpointOutputWithContext(ctx).OutputState,
	}
}

// AiEndpointArrayInput is an input type that accepts AiEndpointArray and AiEndpointArrayOutput values.
// You can construct a concrete instance of `AiEndpointArrayInput` via:
//
//	AiEndpointArray{ AiEndpointArgs{...} }
type AiEndpointArrayInput interface {
	pulumi.Input

	ToAiEndpointArrayOutput() AiEndpointArrayOutput
	ToAiEndpointArrayOutputWithContext(context.Context) AiEndpointArrayOutput
}

type AiEndpointArray []AiEndpointInput

func (AiEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiEndpoint)(nil)).Elem()
}

func (i AiEndpointArray) ToAiEndpointArrayOutput() AiEndpointArrayOutput {
	return i.ToAiEndpointArrayOutputWithContext(context.Background())
}

func (i AiEndpointArray) ToAiEndpointArrayOutputWithContext(ctx context.Context) AiEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiEndpointArrayOutput)
}

func (i AiEndpointArray) ToOutput(ctx context.Context) pulumix.Output[[]*AiEndpoint] {
	return pulumix.Output[[]*AiEndpoint]{
		OutputState: i.ToAiEndpointArrayOutputWithContext(ctx).OutputState,
	}
}

// AiEndpointMapInput is an input type that accepts AiEndpointMap and AiEndpointMapOutput values.
// You can construct a concrete instance of `AiEndpointMapInput` via:
//
//	AiEndpointMap{ "key": AiEndpointArgs{...} }
type AiEndpointMapInput interface {
	pulumi.Input

	ToAiEndpointMapOutput() AiEndpointMapOutput
	ToAiEndpointMapOutputWithContext(context.Context) AiEndpointMapOutput
}

type AiEndpointMap map[string]AiEndpointInput

func (AiEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiEndpoint)(nil)).Elem()
}

func (i AiEndpointMap) ToAiEndpointMapOutput() AiEndpointMapOutput {
	return i.ToAiEndpointMapOutputWithContext(context.Background())
}

func (i AiEndpointMap) ToAiEndpointMapOutputWithContext(ctx context.Context) AiEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiEndpointMapOutput)
}

func (i AiEndpointMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AiEndpoint] {
	return pulumix.Output[map[string]*AiEndpoint]{
		OutputState: i.ToAiEndpointMapOutputWithContext(ctx).OutputState,
	}
}

type AiEndpointOutput struct{ *pulumi.OutputState }

func (AiEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiEndpoint)(nil)).Elem()
}

func (o AiEndpointOutput) ToAiEndpointOutput() AiEndpointOutput {
	return o
}

func (o AiEndpointOutput) ToAiEndpointOutputWithContext(ctx context.Context) AiEndpointOutput {
	return o
}

func (o AiEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*AiEndpoint] {
	return pulumix.Output[*AiEndpoint]{
		OutputState: o.OutputState,
	}
}

// (Output)
// Output only. Timestamp when the DeployedModel was created.
func (o AiEndpointOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/).
// Structure is documented below.
func (o AiEndpointOutput) DeployedModels() AiEndpointDeployedModelArrayOutput {
	return o.ApplyT(func(v *AiEndpoint) AiEndpointDeployedModelArrayOutput { return v.DeployedModels }).(AiEndpointDeployedModelArrayOutput)
}

// The description of the Endpoint.
func (o AiEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters.
func (o AiEndpointOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o AiEndpointOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key.
// Structure is documented below.
func (o AiEndpointOutput) EncryptionSpec() AiEndpointEncryptionSpecPtrOutput {
	return o.ApplyT(func(v *AiEndpoint) AiEndpointEncryptionSpecPtrOutput { return v.EncryptionSpec }).(AiEndpointEncryptionSpecPtrOutput)
}

// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
func (o AiEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The labels with user-defined metadata to organize your Endpoints. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o AiEndpointOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the resource
//
// ***
func (o AiEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: `projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}`
func (o AiEndpointOutput) ModelDeploymentMonitoringJob() pulumi.StringOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringOutput { return v.ModelDeploymentMonitoringJob }).(pulumi.StringOutput)
}

// The resource name of the Endpoint. The name must be numeric with no leading zeros and can be at most 10 digits.
func (o AiEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`. Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
func (o AiEndpointOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringPtrOutput { return v.Network }).(pulumi.StringPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o AiEndpointOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o AiEndpointOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// The region for the resource
func (o AiEndpointOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// Output only. Timestamp when this Endpoint was last updated.
func (o AiEndpointOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AiEndpoint) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type AiEndpointArrayOutput struct{ *pulumi.OutputState }

func (AiEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiEndpoint)(nil)).Elem()
}

func (o AiEndpointArrayOutput) ToAiEndpointArrayOutput() AiEndpointArrayOutput {
	return o
}

func (o AiEndpointArrayOutput) ToAiEndpointArrayOutputWithContext(ctx context.Context) AiEndpointArrayOutput {
	return o
}

func (o AiEndpointArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AiEndpoint] {
	return pulumix.Output[[]*AiEndpoint]{
		OutputState: o.OutputState,
	}
}

func (o AiEndpointArrayOutput) Index(i pulumi.IntInput) AiEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiEndpoint {
		return vs[0].([]*AiEndpoint)[vs[1].(int)]
	}).(AiEndpointOutput)
}

type AiEndpointMapOutput struct{ *pulumi.OutputState }

func (AiEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiEndpoint)(nil)).Elem()
}

func (o AiEndpointMapOutput) ToAiEndpointMapOutput() AiEndpointMapOutput {
	return o
}

func (o AiEndpointMapOutput) ToAiEndpointMapOutputWithContext(ctx context.Context) AiEndpointMapOutput {
	return o
}

func (o AiEndpointMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AiEndpoint] {
	return pulumix.Output[map[string]*AiEndpoint]{
		OutputState: o.OutputState,
	}
}

func (o AiEndpointMapOutput) MapIndex(k pulumi.StringInput) AiEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiEndpoint {
		return vs[0].(map[string]*AiEndpoint)[vs[1].(string)]
	}).(AiEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiEndpointInput)(nil)).Elem(), &AiEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiEndpointArrayInput)(nil)).Elem(), AiEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiEndpointMapInput)(nil)).Elem(), AiEndpointMap{})
	pulumi.RegisterOutputType(AiEndpointOutput{})
	pulumi.RegisterOutputType(AiEndpointArrayOutput{})
	pulumi.RegisterOutputType(AiEndpointMapOutput{})
}
