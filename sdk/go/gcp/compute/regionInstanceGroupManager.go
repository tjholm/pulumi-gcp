// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Google Compute Engine Regional Instance Group Manager API creates and manages pools
// of homogeneous Compute Engine virtual machine instances from a common instance
// template.
//
// To get more information about regionInstanceGroupManagers, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroupManagers)
// * How-to Guides
//     * [Regional Instance Groups Guide](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups)
//
// > **Note:** Use [compute.InstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html) to create a zonal instance group manager.
//
// ## Example Usage
// ### With Multiple Versions
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewRegionInstanceGroupManager(ctx, "appserver", &compute.RegionInstanceGroupManagerArgs{
// 			BaseInstanceName: pulumi.String("app"),
// 			Region:           pulumi.String("us-central1"),
// 			TargetSize:       pulumi.Int(5),
// 			Versions: compute.RegionInstanceGroupManagerVersionArray{
// 				&compute.RegionInstanceGroupManagerVersionArgs{
// 					InstanceTemplate: pulumi.Any(google_compute_instance_template.Appserver.Id),
// 				},
// 				&compute.RegionInstanceGroupManagerVersionArgs{
// 					InstanceTemplate: pulumi.Any(google_compute_instance_template.Appserver - canary.Id),
// 					TargetSize: &compute.RegionInstanceGroupManagerVersionTargetSizeArgs{
// 						Fixed: pulumi.Int(1),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Instance group managers can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager appserver appserver-igm
// ```
type RegionInstanceGroupManager struct {
	pulumi.CustomResourceState

	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies RegionInstanceGroupManagerAutoHealingPoliciesPtrOutput `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringOutput `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The shape to which the group converges either proactively or on resize events (depending on the value set in update_policy.0.instance_redistribution_type). For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/regional-mig-distribution-shape).
	DistributionPolicyTargetShape pulumi.StringOutput `pulumi:"distributionPolicyTargetShape"`
	// The distribution policy for this managed instance
	// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
	DistributionPolicyZones pulumi.StringArrayOutput `pulumi:"distributionPolicyZones"`
	// The fingerprint of the instance group manager.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The full URL of the instance group created by the manager.
	InstanceGroup pulumi.StringOutput `pulumi:"instanceGroup"`
	// - Version name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts RegionInstanceGroupManagerNamedPortArrayOutput `pulumi:"namedPorts"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region where the managed instance group resides. If not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URL of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs). Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the `updatePolicy`.
	StatefulDisks RegionInstanceGroupManagerStatefulDiskArrayOutput `pulumi:"statefulDisks"`
	// The status of this managed instance group.
	Statuses RegionInstanceGroupManagerStatusArrayOutput `pulumi:"statuses"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayOutput `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntOutput `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
	UpdatePolicy RegionInstanceGroupManagerUpdatePolicyOutput `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions RegionInstanceGroupManagerVersionArrayOutput `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, the provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrOutput `pulumi:"waitForInstances"`
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus pulumi.StringPtrOutput `pulumi:"waitForInstancesStatus"`
}

// NewRegionInstanceGroupManager registers a new resource with the given unique name, arguments, and options.
func NewRegionInstanceGroupManager(ctx *pulumi.Context,
	name string, args *RegionInstanceGroupManagerArgs, opts ...pulumi.ResourceOption) (*RegionInstanceGroupManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'BaseInstanceName'")
	}
	if args.Versions == nil {
		return nil, errors.New("invalid value for required argument 'Versions'")
	}
	var resource RegionInstanceGroupManager
	err := ctx.RegisterResource("gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionInstanceGroupManager gets an existing RegionInstanceGroupManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionInstanceGroupManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionInstanceGroupManagerState, opts ...pulumi.ResourceOption) (*RegionInstanceGroupManager, error) {
	var resource RegionInstanceGroupManager
	err := ctx.ReadResource("gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionInstanceGroupManager resources.
type regionInstanceGroupManagerState struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies *RegionInstanceGroupManagerAutoHealingPolicies `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName *string `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description *string `pulumi:"description"`
	// The shape to which the group converges either proactively or on resize events (depending on the value set in update_policy.0.instance_redistribution_type). For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/regional-mig-distribution-shape).
	DistributionPolicyTargetShape *string `pulumi:"distributionPolicyTargetShape"`
	// The distribution policy for this managed instance
	// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
	DistributionPolicyZones []string `pulumi:"distributionPolicyZones"`
	// The fingerprint of the instance group manager.
	Fingerprint *string `pulumi:"fingerprint"`
	// The full URL of the instance group created by the manager.
	InstanceGroup *string `pulumi:"instanceGroup"`
	// - Version name.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []RegionInstanceGroupManagerNamedPort `pulumi:"namedPorts"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region where the managed instance group resides. If not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URL of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs). Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the `updatePolicy`.
	StatefulDisks []RegionInstanceGroupManagerStatefulDisk `pulumi:"statefulDisks"`
	// The status of this managed instance group.
	Statuses []RegionInstanceGroupManagerStatus `pulumi:"statuses"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools []string `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
	UpdatePolicy *RegionInstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions []RegionInstanceGroupManagerVersion `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, the provider will
	// continue trying until it times out.
	WaitForInstances *bool `pulumi:"waitForInstances"`
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus *string `pulumi:"waitForInstancesStatus"`
}

type RegionInstanceGroupManagerState struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies RegionInstanceGroupManagerAutoHealingPoliciesPtrInput
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringPtrInput
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrInput
	// The shape to which the group converges either proactively or on resize events (depending on the value set in update_policy.0.instance_redistribution_type). For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/regional-mig-distribution-shape).
	DistributionPolicyTargetShape pulumi.StringPtrInput
	// The distribution policy for this managed instance
	// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
	DistributionPolicyZones pulumi.StringArrayInput
	// The fingerprint of the instance group manager.
	Fingerprint pulumi.StringPtrInput
	// The full URL of the instance group created by the manager.
	InstanceGroup pulumi.StringPtrInput
	// - Version name.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts RegionInstanceGroupManagerNamedPortArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region where the managed instance group resides. If not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URL of the created resource.
	SelfLink pulumi.StringPtrInput
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs). Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the `updatePolicy`.
	StatefulDisks RegionInstanceGroupManagerStatefulDiskArrayInput
	// The status of this managed instance group.
	Statuses RegionInstanceGroupManagerStatusArrayInput
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayInput
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
	UpdatePolicy RegionInstanceGroupManagerUpdatePolicyPtrInput
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions RegionInstanceGroupManagerVersionArrayInput
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, the provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrInput
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus pulumi.StringPtrInput
}

func (RegionInstanceGroupManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionInstanceGroupManagerState)(nil)).Elem()
}

type regionInstanceGroupManagerArgs struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies *RegionInstanceGroupManagerAutoHealingPolicies `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName string `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description *string `pulumi:"description"`
	// The shape to which the group converges either proactively or on resize events (depending on the value set in update_policy.0.instance_redistribution_type). For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/regional-mig-distribution-shape).
	DistributionPolicyTargetShape *string `pulumi:"distributionPolicyTargetShape"`
	// The distribution policy for this managed instance
	// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
	DistributionPolicyZones []string `pulumi:"distributionPolicyZones"`
	// - Version name.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []RegionInstanceGroupManagerNamedPort `pulumi:"namedPorts"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region where the managed instance group resides. If not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs). Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the `updatePolicy`.
	StatefulDisks []RegionInstanceGroupManagerStatefulDisk `pulumi:"statefulDisks"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools []string `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
	UpdatePolicy *RegionInstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions []RegionInstanceGroupManagerVersion `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, the provider will
	// continue trying until it times out.
	WaitForInstances *bool `pulumi:"waitForInstances"`
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus *string `pulumi:"waitForInstancesStatus"`
}

// The set of arguments for constructing a RegionInstanceGroupManager resource.
type RegionInstanceGroupManagerArgs struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies RegionInstanceGroupManagerAutoHealingPoliciesPtrInput
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringInput
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrInput
	// The shape to which the group converges either proactively or on resize events (depending on the value set in update_policy.0.instance_redistribution_type). For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/regional-mig-distribution-shape).
	DistributionPolicyTargetShape pulumi.StringPtrInput
	// The distribution policy for this managed instance
	// group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
	DistributionPolicyZones pulumi.StringArrayInput
	// - Version name.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts RegionInstanceGroupManagerNamedPortArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region where the managed instance group resides. If not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs). Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the `updatePolicy`.
	StatefulDisks RegionInstanceGroupManagerStatefulDiskArrayInput
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayInput
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)
	UpdatePolicy RegionInstanceGroupManagerUpdatePolicyPtrInput
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions RegionInstanceGroupManagerVersionArrayInput
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, the provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrInput
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus pulumi.StringPtrInput
}

func (RegionInstanceGroupManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionInstanceGroupManagerArgs)(nil)).Elem()
}

type RegionInstanceGroupManagerInput interface {
	pulumi.Input

	ToRegionInstanceGroupManagerOutput() RegionInstanceGroupManagerOutput
	ToRegionInstanceGroupManagerOutputWithContext(ctx context.Context) RegionInstanceGroupManagerOutput
}

func (*RegionInstanceGroupManager) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionInstanceGroupManager)(nil))
}

func (i *RegionInstanceGroupManager) ToRegionInstanceGroupManagerOutput() RegionInstanceGroupManagerOutput {
	return i.ToRegionInstanceGroupManagerOutputWithContext(context.Background())
}

func (i *RegionInstanceGroupManager) ToRegionInstanceGroupManagerOutputWithContext(ctx context.Context) RegionInstanceGroupManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceGroupManagerOutput)
}

func (i *RegionInstanceGroupManager) ToRegionInstanceGroupManagerPtrOutput() RegionInstanceGroupManagerPtrOutput {
	return i.ToRegionInstanceGroupManagerPtrOutputWithContext(context.Background())
}

func (i *RegionInstanceGroupManager) ToRegionInstanceGroupManagerPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceGroupManagerPtrOutput)
}

type RegionInstanceGroupManagerPtrInput interface {
	pulumi.Input

	ToRegionInstanceGroupManagerPtrOutput() RegionInstanceGroupManagerPtrOutput
	ToRegionInstanceGroupManagerPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerPtrOutput
}

type regionInstanceGroupManagerPtrType RegionInstanceGroupManagerArgs

func (*regionInstanceGroupManagerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionInstanceGroupManager)(nil))
}

func (i *regionInstanceGroupManagerPtrType) ToRegionInstanceGroupManagerPtrOutput() RegionInstanceGroupManagerPtrOutput {
	return i.ToRegionInstanceGroupManagerPtrOutputWithContext(context.Background())
}

func (i *regionInstanceGroupManagerPtrType) ToRegionInstanceGroupManagerPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceGroupManagerPtrOutput)
}

// RegionInstanceGroupManagerArrayInput is an input type that accepts RegionInstanceGroupManagerArray and RegionInstanceGroupManagerArrayOutput values.
// You can construct a concrete instance of `RegionInstanceGroupManagerArrayInput` via:
//
//          RegionInstanceGroupManagerArray{ RegionInstanceGroupManagerArgs{...} }
type RegionInstanceGroupManagerArrayInput interface {
	pulumi.Input

	ToRegionInstanceGroupManagerArrayOutput() RegionInstanceGroupManagerArrayOutput
	ToRegionInstanceGroupManagerArrayOutputWithContext(context.Context) RegionInstanceGroupManagerArrayOutput
}

type RegionInstanceGroupManagerArray []RegionInstanceGroupManagerInput

func (RegionInstanceGroupManagerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RegionInstanceGroupManager)(nil))
}

func (i RegionInstanceGroupManagerArray) ToRegionInstanceGroupManagerArrayOutput() RegionInstanceGroupManagerArrayOutput {
	return i.ToRegionInstanceGroupManagerArrayOutputWithContext(context.Background())
}

func (i RegionInstanceGroupManagerArray) ToRegionInstanceGroupManagerArrayOutputWithContext(ctx context.Context) RegionInstanceGroupManagerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceGroupManagerArrayOutput)
}

// RegionInstanceGroupManagerMapInput is an input type that accepts RegionInstanceGroupManagerMap and RegionInstanceGroupManagerMapOutput values.
// You can construct a concrete instance of `RegionInstanceGroupManagerMapInput` via:
//
//          RegionInstanceGroupManagerMap{ "key": RegionInstanceGroupManagerArgs{...} }
type RegionInstanceGroupManagerMapInput interface {
	pulumi.Input

	ToRegionInstanceGroupManagerMapOutput() RegionInstanceGroupManagerMapOutput
	ToRegionInstanceGroupManagerMapOutputWithContext(context.Context) RegionInstanceGroupManagerMapOutput
}

type RegionInstanceGroupManagerMap map[string]RegionInstanceGroupManagerInput

func (RegionInstanceGroupManagerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RegionInstanceGroupManager)(nil))
}

func (i RegionInstanceGroupManagerMap) ToRegionInstanceGroupManagerMapOutput() RegionInstanceGroupManagerMapOutput {
	return i.ToRegionInstanceGroupManagerMapOutputWithContext(context.Background())
}

func (i RegionInstanceGroupManagerMap) ToRegionInstanceGroupManagerMapOutputWithContext(ctx context.Context) RegionInstanceGroupManagerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionInstanceGroupManagerMapOutput)
}

type RegionInstanceGroupManagerOutput struct {
	*pulumi.OutputState
}

func (RegionInstanceGroupManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionInstanceGroupManager)(nil))
}

func (o RegionInstanceGroupManagerOutput) ToRegionInstanceGroupManagerOutput() RegionInstanceGroupManagerOutput {
	return o
}

func (o RegionInstanceGroupManagerOutput) ToRegionInstanceGroupManagerOutputWithContext(ctx context.Context) RegionInstanceGroupManagerOutput {
	return o
}

func (o RegionInstanceGroupManagerOutput) ToRegionInstanceGroupManagerPtrOutput() RegionInstanceGroupManagerPtrOutput {
	return o.ToRegionInstanceGroupManagerPtrOutputWithContext(context.Background())
}

func (o RegionInstanceGroupManagerOutput) ToRegionInstanceGroupManagerPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerPtrOutput {
	return o.ApplyT(func(v RegionInstanceGroupManager) *RegionInstanceGroupManager {
		return &v
	}).(RegionInstanceGroupManagerPtrOutput)
}

type RegionInstanceGroupManagerPtrOutput struct {
	*pulumi.OutputState
}

func (RegionInstanceGroupManagerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionInstanceGroupManager)(nil))
}

func (o RegionInstanceGroupManagerPtrOutput) ToRegionInstanceGroupManagerPtrOutput() RegionInstanceGroupManagerPtrOutput {
	return o
}

func (o RegionInstanceGroupManagerPtrOutput) ToRegionInstanceGroupManagerPtrOutputWithContext(ctx context.Context) RegionInstanceGroupManagerPtrOutput {
	return o
}

type RegionInstanceGroupManagerArrayOutput struct{ *pulumi.OutputState }

func (RegionInstanceGroupManagerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegionInstanceGroupManager)(nil))
}

func (o RegionInstanceGroupManagerArrayOutput) ToRegionInstanceGroupManagerArrayOutput() RegionInstanceGroupManagerArrayOutput {
	return o
}

func (o RegionInstanceGroupManagerArrayOutput) ToRegionInstanceGroupManagerArrayOutputWithContext(ctx context.Context) RegionInstanceGroupManagerArrayOutput {
	return o
}

func (o RegionInstanceGroupManagerArrayOutput) Index(i pulumi.IntInput) RegionInstanceGroupManagerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegionInstanceGroupManager {
		return vs[0].([]RegionInstanceGroupManager)[vs[1].(int)]
	}).(RegionInstanceGroupManagerOutput)
}

type RegionInstanceGroupManagerMapOutput struct{ *pulumi.OutputState }

func (RegionInstanceGroupManagerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RegionInstanceGroupManager)(nil))
}

func (o RegionInstanceGroupManagerMapOutput) ToRegionInstanceGroupManagerMapOutput() RegionInstanceGroupManagerMapOutput {
	return o
}

func (o RegionInstanceGroupManagerMapOutput) ToRegionInstanceGroupManagerMapOutputWithContext(ctx context.Context) RegionInstanceGroupManagerMapOutput {
	return o
}

func (o RegionInstanceGroupManagerMapOutput) MapIndex(k pulumi.StringInput) RegionInstanceGroupManagerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RegionInstanceGroupManager {
		return vs[0].(map[string]RegionInstanceGroupManager)[vs[1].(string)]
	}).(RegionInstanceGroupManagerOutput)
}

func init() {
	pulumi.RegisterOutputType(RegionInstanceGroupManagerOutput{})
	pulumi.RegisterOutputType(RegionInstanceGroupManagerPtrOutput{})
	pulumi.RegisterOutputType(RegionInstanceGroupManagerArrayOutput{})
	pulumi.RegisterOutputType(RegionInstanceGroupManagerMapOutput{})
}
