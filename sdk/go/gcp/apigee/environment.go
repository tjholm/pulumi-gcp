// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigee

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An `Environment` in Apigee.
//
// To get more information about Environment, see:
//
// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments/create)
// * How-to Guides
//   - [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
//
// ## Example Usage
// ### Apigee Environment Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/apigee"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/servicenetworking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := organizations.GetClientConfig(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			apigeeNetwork, err := compute.NewNetwork(ctx, "apigeeNetwork", nil)
//			if err != nil {
//				return err
//			}
//			apigeeRange, err := compute.NewGlobalAddress(ctx, "apigeeRange", &compute.GlobalAddressArgs{
//				Purpose:      pulumi.String("VPC_PEERING"),
//				AddressType:  pulumi.String("INTERNAL"),
//				PrefixLength: pulumi.Int(16),
//				Network:      apigeeNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeVpcConnection, err := servicenetworking.NewConnection(ctx, "apigeeVpcConnection", &servicenetworking.ConnectionArgs{
//				Network: apigeeNetwork.ID(),
//				Service: pulumi.String("servicenetworking.googleapis.com"),
//				ReservedPeeringRanges: pulumi.StringArray{
//					apigeeRange.Name,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			apigeeOrg, err := apigee.NewOrganization(ctx, "apigeeOrg", &apigee.OrganizationArgs{
//				AnalyticsRegion:   pulumi.String("us-central1"),
//				ProjectId:         pulumi.String(current.Project),
//				AuthorizedNetwork: apigeeNetwork.ID(),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				apigeeVpcConnection,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewEnvironment(ctx, "env", &apigee.EnvironmentArgs{
//				Description: pulumi.String("Apigee Environment"),
//				DisplayName: pulumi.String("environment-1"),
//				OrgId:       apigeeOrg.ID(),
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
// # Environment can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:apigee/environment:Environment default {{org_id}}/environments/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigee/environment:Environment default {{org_id}}/{{name}}
//
// ```
type Environment struct {
	pulumi.CustomResourceState

	// Optional. API Proxy type supported by the environment. The type can be set when creating
	// the Environment and cannot be changed.
	// Possible values are `API_PROXY_TYPE_UNSPECIFIED`, `PROGRAMMABLE`, and `CONFIGURABLE`.
	ApiProxyType pulumi.StringOutput `pulumi:"apiProxyType"`
	// Optional. Deployment type supported by the environment. The deployment type can be
	// set when creating the environment and cannot be changed. When you enable archive
	// deployment, you will be prevented from performing a subset of actions within the
	// environment, including:
	// Managing the deployment of API proxy or shared flow revisions;
	// Creating, updating, or deleting resource files;
	// Creating, updating, or deleting target servers.
	// Possible values are `DEPLOYMENT_TYPE_UNSPECIFIED`, `PROXY`, and `ARCHIVE`.
	DeploymentType pulumi.StringOutput `pulumi:"deploymentType"`
	// Description of the environment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display name of the environment.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The resource ID of the environment.
	Name pulumi.StringOutput `pulumi:"name"`
	// NodeConfig for setting the min/max number of nodes associated with the environment.
	// Structure is documented below.
	NodeConfig EnvironmentNodeConfigOutput `pulumi:"nodeConfig"`
	// The Apigee Organization associated with the Apigee environment,
	// in the format `organizations/{{org_name}}`.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	var resource Environment
	err := ctx.RegisterResource("gcp:apigee/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("gcp:apigee/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// Optional. API Proxy type supported by the environment. The type can be set when creating
	// the Environment and cannot be changed.
	// Possible values are `API_PROXY_TYPE_UNSPECIFIED`, `PROGRAMMABLE`, and `CONFIGURABLE`.
	ApiProxyType *string `pulumi:"apiProxyType"`
	// Optional. Deployment type supported by the environment. The deployment type can be
	// set when creating the environment and cannot be changed. When you enable archive
	// deployment, you will be prevented from performing a subset of actions within the
	// environment, including:
	// Managing the deployment of API proxy or shared flow revisions;
	// Creating, updating, or deleting resource files;
	// Creating, updating, or deleting target servers.
	// Possible values are `DEPLOYMENT_TYPE_UNSPECIFIED`, `PROXY`, and `ARCHIVE`.
	DeploymentType *string `pulumi:"deploymentType"`
	// Description of the environment.
	Description *string `pulumi:"description"`
	// Display name of the environment.
	DisplayName *string `pulumi:"displayName"`
	// The resource ID of the environment.
	Name *string `pulumi:"name"`
	// NodeConfig for setting the min/max number of nodes associated with the environment.
	// Structure is documented below.
	NodeConfig *EnvironmentNodeConfig `pulumi:"nodeConfig"`
	// The Apigee Organization associated with the Apigee environment,
	// in the format `organizations/{{org_name}}`.
	OrgId *string `pulumi:"orgId"`
}

type EnvironmentState struct {
	// Optional. API Proxy type supported by the environment. The type can be set when creating
	// the Environment and cannot be changed.
	// Possible values are `API_PROXY_TYPE_UNSPECIFIED`, `PROGRAMMABLE`, and `CONFIGURABLE`.
	ApiProxyType pulumi.StringPtrInput
	// Optional. Deployment type supported by the environment. The deployment type can be
	// set when creating the environment and cannot be changed. When you enable archive
	// deployment, you will be prevented from performing a subset of actions within the
	// environment, including:
	// Managing the deployment of API proxy or shared flow revisions;
	// Creating, updating, or deleting resource files;
	// Creating, updating, or deleting target servers.
	// Possible values are `DEPLOYMENT_TYPE_UNSPECIFIED`, `PROXY`, and `ARCHIVE`.
	DeploymentType pulumi.StringPtrInput
	// Description of the environment.
	Description pulumi.StringPtrInput
	// Display name of the environment.
	DisplayName pulumi.StringPtrInput
	// The resource ID of the environment.
	Name pulumi.StringPtrInput
	// NodeConfig for setting the min/max number of nodes associated with the environment.
	// Structure is documented below.
	NodeConfig EnvironmentNodeConfigPtrInput
	// The Apigee Organization associated with the Apigee environment,
	// in the format `organizations/{{org_name}}`.
	OrgId pulumi.StringPtrInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// Optional. API Proxy type supported by the environment. The type can be set when creating
	// the Environment and cannot be changed.
	// Possible values are `API_PROXY_TYPE_UNSPECIFIED`, `PROGRAMMABLE`, and `CONFIGURABLE`.
	ApiProxyType *string `pulumi:"apiProxyType"`
	// Optional. Deployment type supported by the environment. The deployment type can be
	// set when creating the environment and cannot be changed. When you enable archive
	// deployment, you will be prevented from performing a subset of actions within the
	// environment, including:
	// Managing the deployment of API proxy or shared flow revisions;
	// Creating, updating, or deleting resource files;
	// Creating, updating, or deleting target servers.
	// Possible values are `DEPLOYMENT_TYPE_UNSPECIFIED`, `PROXY`, and `ARCHIVE`.
	DeploymentType *string `pulumi:"deploymentType"`
	// Description of the environment.
	Description *string `pulumi:"description"`
	// Display name of the environment.
	DisplayName *string `pulumi:"displayName"`
	// The resource ID of the environment.
	Name *string `pulumi:"name"`
	// NodeConfig for setting the min/max number of nodes associated with the environment.
	// Structure is documented below.
	NodeConfig *EnvironmentNodeConfig `pulumi:"nodeConfig"`
	// The Apigee Organization associated with the Apigee environment,
	// in the format `organizations/{{org_name}}`.
	OrgId string `pulumi:"orgId"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// Optional. API Proxy type supported by the environment. The type can be set when creating
	// the Environment and cannot be changed.
	// Possible values are `API_PROXY_TYPE_UNSPECIFIED`, `PROGRAMMABLE`, and `CONFIGURABLE`.
	ApiProxyType pulumi.StringPtrInput
	// Optional. Deployment type supported by the environment. The deployment type can be
	// set when creating the environment and cannot be changed. When you enable archive
	// deployment, you will be prevented from performing a subset of actions within the
	// environment, including:
	// Managing the deployment of API proxy or shared flow revisions;
	// Creating, updating, or deleting resource files;
	// Creating, updating, or deleting target servers.
	// Possible values are `DEPLOYMENT_TYPE_UNSPECIFIED`, `PROXY`, and `ARCHIVE`.
	DeploymentType pulumi.StringPtrInput
	// Description of the environment.
	Description pulumi.StringPtrInput
	// Display name of the environment.
	DisplayName pulumi.StringPtrInput
	// The resource ID of the environment.
	Name pulumi.StringPtrInput
	// NodeConfig for setting the min/max number of nodes associated with the environment.
	// Structure is documented below.
	NodeConfig EnvironmentNodeConfigPtrInput
	// The Apigee Organization associated with the Apigee environment,
	// in the format `organizations/{{org_name}}`.
	OrgId pulumi.StringInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

// EnvironmentArrayInput is an input type that accepts EnvironmentArray and EnvironmentArrayOutput values.
// You can construct a concrete instance of `EnvironmentArrayInput` via:
//
//	EnvironmentArray{ EnvironmentArgs{...} }
type EnvironmentArrayInput interface {
	pulumi.Input

	ToEnvironmentArrayOutput() EnvironmentArrayOutput
	ToEnvironmentArrayOutputWithContext(context.Context) EnvironmentArrayOutput
}

type EnvironmentArray []EnvironmentInput

func (EnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (i EnvironmentArray) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return i.ToEnvironmentArrayOutputWithContext(context.Background())
}

func (i EnvironmentArray) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentArrayOutput)
}

// EnvironmentMapInput is an input type that accepts EnvironmentMap and EnvironmentMapOutput values.
// You can construct a concrete instance of `EnvironmentMapInput` via:
//
//	EnvironmentMap{ "key": EnvironmentArgs{...} }
type EnvironmentMapInput interface {
	pulumi.Input

	ToEnvironmentMapOutput() EnvironmentMapOutput
	ToEnvironmentMapOutputWithContext(context.Context) EnvironmentMapOutput
}

type EnvironmentMap map[string]EnvironmentInput

func (EnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (i EnvironmentMap) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return i.ToEnvironmentMapOutputWithContext(context.Background())
}

func (i EnvironmentMap) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentMapOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

// Optional. API Proxy type supported by the environment. The type can be set when creating
// the Environment and cannot be changed.
// Possible values are `API_PROXY_TYPE_UNSPECIFIED`, `PROGRAMMABLE`, and `CONFIGURABLE`.
func (o EnvironmentOutput) ApiProxyType() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.ApiProxyType }).(pulumi.StringOutput)
}

// Optional. Deployment type supported by the environment. The deployment type can be
// set when creating the environment and cannot be changed. When you enable archive
// deployment, you will be prevented from performing a subset of actions within the
// environment, including:
// Managing the deployment of API proxy or shared flow revisions;
// Creating, updating, or deleting resource files;
// Creating, updating, or deleting target servers.
// Possible values are `DEPLOYMENT_TYPE_UNSPECIFIED`, `PROXY`, and `ARCHIVE`.
func (o EnvironmentOutput) DeploymentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.DeploymentType }).(pulumi.StringOutput)
}

// Description of the environment.
func (o EnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name of the environment.
func (o EnvironmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The resource ID of the environment.
func (o EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// NodeConfig for setting the min/max number of nodes associated with the environment.
// Structure is documented below.
func (o EnvironmentOutput) NodeConfig() EnvironmentNodeConfigOutput {
	return o.ApplyT(func(v *Environment) EnvironmentNodeConfigOutput { return v.NodeConfig }).(EnvironmentNodeConfigOutput)
}

// The Apigee Organization associated with the Apigee environment,
// in the format `organizations/{{org_name}}`.
func (o EnvironmentOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

type EnvironmentArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) Index(i pulumi.IntInput) EnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].([]*Environment)[vs[1].(int)]
	}).(EnvironmentOutput)
}

type EnvironmentMapOutput struct{ *pulumi.OutputState }

func (EnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) MapIndex(k pulumi.StringInput) EnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].(map[string]*Environment)[vs[1].(string)]
	}).(EnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentArrayInput)(nil)).Elem(), EnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentMapInput)(nil)).Elem(), EnvironmentMap{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentMapOutput{})
}
