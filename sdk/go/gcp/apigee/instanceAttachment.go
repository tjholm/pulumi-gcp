// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigee

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An `Instance attachment` in Apigee.
//
// To get more information about InstanceAttachment, see:
//
// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances.attachments/create)
// * How-to Guides
//   - [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
//
// ## Example Usage
// ### Apigee Instance Attachment Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/apigee"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/projects"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/servicenetworking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.NewProject(ctx, "project", &organizations.ProjectArgs{
//				OrgId:          pulumi.String(""),
//				BillingAccount: pulumi.String(""),
//			})
//			if err != nil {
//				return err
//			}
//			apigee, err := projects.NewService(ctx, "apigee", &projects.ServiceArgs{
//				Project: project.ProjectId,
//				Service: pulumi.String("apigee.googleapis.com"),
//			})
//			if err != nil {
//				return err
//			}
//			compute, err := projects.NewService(ctx, "compute", &projects.ServiceArgs{
//				Project: project.ProjectId,
//				Service: pulumi.String("compute.googleapis.com"),
//			})
//			if err != nil {
//				return err
//			}
//			servicenetworking, err := projects.NewService(ctx, "servicenetworking", &projects.ServiceArgs{
//				Project: project.ProjectId,
//				Service: pulumi.String("servicenetworking.googleapis.com"),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeNetwork, err := compute.NewNetwork(ctx, "apigeeNetwork", &compute.NetworkArgs{
//				Project: project.ProjectId,
//			}, pulumi.DependsOn([]pulumi.Resource{
//				compute,
//			}))
//			if err != nil {
//				return err
//			}
//			apigeeRange, err := compute.NewGlobalAddress(ctx, "apigeeRange", &compute.GlobalAddressArgs{
//				Purpose:      pulumi.String("VPC_PEERING"),
//				AddressType:  pulumi.String("INTERNAL"),
//				PrefixLength: pulumi.Int(16),
//				Network:      apigeeNetwork.ID(),
//				Project:      project.ProjectId,
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
//			}, pulumi.DependsOn([]pulumi.Resource{
//				servicenetworking,
//			}))
//			if err != nil {
//				return err
//			}
//			apigeeOrg, err := apigee.NewOrganization(ctx, "apigeeOrg", &apigee.OrganizationArgs{
//				AnalyticsRegion:   pulumi.String("us-central1"),
//				ProjectId:         project.ProjectId,
//				AuthorizedNetwork: apigeeNetwork.ID(),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				apigeeVpcConnection,
//				apigee,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewInstance(ctx, "apigeeIns", &apigee.InstanceArgs{
//				Location: pulumi.String("us-central1"),
//				OrgId:    apigeeOrg.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			apigeeEnv, err := apigee.NewEnvironment(ctx, "apigeeEnv", &apigee.EnvironmentArgs{
//				OrgId:       apigeeOrg.ID(),
//				Description: pulumi.String("Apigee Environment"),
//				DisplayName: pulumi.String("environment-1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = apigee.NewInstanceAttachment(ctx, "instanceAttachment", &apigee.InstanceAttachmentArgs{
//				InstanceId:  pulumi.Any(google_apigee_instance.Apigee_instance.Id),
//				Environment: apigeeEnv.Name,
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
// InstanceAttachment can be imported using any of these accepted formats* `{{instance_id}}/attachments/{{name}}` * `{{instance_id}}/{{name}}` When using the `pulumi import` command, InstanceAttachment can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:apigee/instanceAttachment:InstanceAttachment default {{instance_id}}/attachments/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:apigee/instanceAttachment:InstanceAttachment default {{instance_id}}/{{name}}
//
// ```
type InstanceAttachment struct {
	pulumi.CustomResourceState

	// The resource ID of the environment.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	//
	// ***
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of the newly created  attachment (output parameter).
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewInstanceAttachment registers a new resource with the given unique name, arguments, and options.
func NewInstanceAttachment(ctx *pulumi.Context,
	name string, args *InstanceAttachmentArgs, opts ...pulumi.ResourceOption) (*InstanceAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceAttachment
	err := ctx.RegisterResource("gcp:apigee/instanceAttachment:InstanceAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceAttachment gets an existing InstanceAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceAttachmentState, opts ...pulumi.ResourceOption) (*InstanceAttachment, error) {
	var resource InstanceAttachment
	err := ctx.ReadResource("gcp:apigee/instanceAttachment:InstanceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceAttachment resources.
type instanceAttachmentState struct {
	// The resource ID of the environment.
	Environment *string `pulumi:"environment"`
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	//
	// ***
	InstanceId *string `pulumi:"instanceId"`
	// The name of the newly created  attachment (output parameter).
	Name *string `pulumi:"name"`
}

type InstanceAttachmentState struct {
	// The resource ID of the environment.
	Environment pulumi.StringPtrInput
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	//
	// ***
	InstanceId pulumi.StringPtrInput
	// The name of the newly created  attachment (output parameter).
	Name pulumi.StringPtrInput
}

func (InstanceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceAttachmentState)(nil)).Elem()
}

type instanceAttachmentArgs struct {
	// The resource ID of the environment.
	Environment string `pulumi:"environment"`
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	//
	// ***
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a InstanceAttachment resource.
type InstanceAttachmentArgs struct {
	// The resource ID of the environment.
	Environment pulumi.StringInput
	// The Apigee instance associated with the Apigee environment,
	// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
	//
	// ***
	InstanceId pulumi.StringInput
}

func (InstanceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceAttachmentArgs)(nil)).Elem()
}

type InstanceAttachmentInput interface {
	pulumi.Input

	ToInstanceAttachmentOutput() InstanceAttachmentOutput
	ToInstanceAttachmentOutputWithContext(ctx context.Context) InstanceAttachmentOutput
}

func (*InstanceAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceAttachment)(nil)).Elem()
}

func (i *InstanceAttachment) ToInstanceAttachmentOutput() InstanceAttachmentOutput {
	return i.ToInstanceAttachmentOutputWithContext(context.Background())
}

func (i *InstanceAttachment) ToInstanceAttachmentOutputWithContext(ctx context.Context) InstanceAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceAttachmentOutput)
}

// InstanceAttachmentArrayInput is an input type that accepts InstanceAttachmentArray and InstanceAttachmentArrayOutput values.
// You can construct a concrete instance of `InstanceAttachmentArrayInput` via:
//
//	InstanceAttachmentArray{ InstanceAttachmentArgs{...} }
type InstanceAttachmentArrayInput interface {
	pulumi.Input

	ToInstanceAttachmentArrayOutput() InstanceAttachmentArrayOutput
	ToInstanceAttachmentArrayOutputWithContext(context.Context) InstanceAttachmentArrayOutput
}

type InstanceAttachmentArray []InstanceAttachmentInput

func (InstanceAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceAttachment)(nil)).Elem()
}

func (i InstanceAttachmentArray) ToInstanceAttachmentArrayOutput() InstanceAttachmentArrayOutput {
	return i.ToInstanceAttachmentArrayOutputWithContext(context.Background())
}

func (i InstanceAttachmentArray) ToInstanceAttachmentArrayOutputWithContext(ctx context.Context) InstanceAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceAttachmentArrayOutput)
}

// InstanceAttachmentMapInput is an input type that accepts InstanceAttachmentMap and InstanceAttachmentMapOutput values.
// You can construct a concrete instance of `InstanceAttachmentMapInput` via:
//
//	InstanceAttachmentMap{ "key": InstanceAttachmentArgs{...} }
type InstanceAttachmentMapInput interface {
	pulumi.Input

	ToInstanceAttachmentMapOutput() InstanceAttachmentMapOutput
	ToInstanceAttachmentMapOutputWithContext(context.Context) InstanceAttachmentMapOutput
}

type InstanceAttachmentMap map[string]InstanceAttachmentInput

func (InstanceAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceAttachment)(nil)).Elem()
}

func (i InstanceAttachmentMap) ToInstanceAttachmentMapOutput() InstanceAttachmentMapOutput {
	return i.ToInstanceAttachmentMapOutputWithContext(context.Background())
}

func (i InstanceAttachmentMap) ToInstanceAttachmentMapOutputWithContext(ctx context.Context) InstanceAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceAttachmentMapOutput)
}

type InstanceAttachmentOutput struct{ *pulumi.OutputState }

func (InstanceAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceAttachment)(nil)).Elem()
}

func (o InstanceAttachmentOutput) ToInstanceAttachmentOutput() InstanceAttachmentOutput {
	return o
}

func (o InstanceAttachmentOutput) ToInstanceAttachmentOutputWithContext(ctx context.Context) InstanceAttachmentOutput {
	return o
}

// The resource ID of the environment.
func (o InstanceAttachmentOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceAttachment) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The Apigee instance associated with the Apigee environment,
// in the format `organizations/{{org_name}}/instances/{{instance_name}}`.
//
// ***
func (o InstanceAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The name of the newly created  attachment (output parameter).
func (o InstanceAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type InstanceAttachmentArrayOutput struct{ *pulumi.OutputState }

func (InstanceAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceAttachment)(nil)).Elem()
}

func (o InstanceAttachmentArrayOutput) ToInstanceAttachmentArrayOutput() InstanceAttachmentArrayOutput {
	return o
}

func (o InstanceAttachmentArrayOutput) ToInstanceAttachmentArrayOutputWithContext(ctx context.Context) InstanceAttachmentArrayOutput {
	return o
}

func (o InstanceAttachmentArrayOutput) Index(i pulumi.IntInput) InstanceAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceAttachment {
		return vs[0].([]*InstanceAttachment)[vs[1].(int)]
	}).(InstanceAttachmentOutput)
}

type InstanceAttachmentMapOutput struct{ *pulumi.OutputState }

func (InstanceAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceAttachment)(nil)).Elem()
}

func (o InstanceAttachmentMapOutput) ToInstanceAttachmentMapOutput() InstanceAttachmentMapOutput {
	return o
}

func (o InstanceAttachmentMapOutput) ToInstanceAttachmentMapOutputWithContext(ctx context.Context) InstanceAttachmentMapOutput {
	return o
}

func (o InstanceAttachmentMapOutput) MapIndex(k pulumi.StringInput) InstanceAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceAttachment {
		return vs[0].(map[string]*InstanceAttachment)[vs[1].(string)]
	}).(InstanceAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceAttachmentInput)(nil)).Elem(), &InstanceAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceAttachmentArrayInput)(nil)).Elem(), InstanceAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceAttachmentMapInput)(nil)).Elem(), InstanceAttachmentMap{})
	pulumi.RegisterOutputType(InstanceAttachmentOutput{})
	pulumi.RegisterOutputType(InstanceAttachmentArrayOutput{})
	pulumi.RegisterOutputType(InstanceAttachmentMapOutput{})
}
