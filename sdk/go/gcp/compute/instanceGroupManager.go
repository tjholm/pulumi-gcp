// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Google Compute Engine Instance Group Manager API creates and manages pools
// of homogeneous Compute Engine virtual machine instances from a common instance
// template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/manager)
// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroupManagers)
//
// > **Note:** Use [compute.RegionInstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager.html) to create a regional (multi-zone) instance group manager.
//
// ## Example Usage
// ### With Top Level Instance Template (`Google` Provider)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			autohealing, err := compute.NewHealthCheck(ctx, "autohealing", &compute.HealthCheckArgs{
//				CheckIntervalSec:   pulumi.Int(5),
//				TimeoutSec:         pulumi.Int(5),
//				HealthyThreshold:   pulumi.Int(2),
//				UnhealthyThreshold: pulumi.Int(10),
//				HttpHealthCheck: &compute.HealthCheckHttpHealthCheckArgs{
//					RequestPath: pulumi.String("/healthz"),
//					Port:        pulumi.Int(8080),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstanceGroupManager(ctx, "appserver", &compute.InstanceGroupManagerArgs{
//				BaseInstanceName: pulumi.String("app"),
//				Zone:             pulumi.String("us-central1-a"),
//				Versions: compute.InstanceGroupManagerVersionArray{
//					&compute.InstanceGroupManagerVersionArgs{
//						InstanceTemplate: pulumi.Any(google_compute_instance_template.Appserver.Id),
//					},
//				},
//				AllInstancesConfig: &compute.InstanceGroupManagerAllInstancesConfigArgs{
//					Metadata: pulumi.StringMap{
//						"metadata_key": pulumi.String("metadata_value"),
//					},
//					Labels: pulumi.StringMap{
//						"label_key": pulumi.String("label_value"),
//					},
//				},
//				TargetPools: pulumi.StringArray{
//					pulumi.Any(google_compute_target_pool.Appserver.Id),
//				},
//				TargetSize: pulumi.Int(2),
//				NamedPorts: compute.InstanceGroupManagerNamedPortArray{
//					&compute.InstanceGroupManagerNamedPortArgs{
//						Name: pulumi.String("customhttp"),
//						Port: pulumi.Int(8888),
//					},
//				},
//				AutoHealingPolicies: &compute.InstanceGroupManagerAutoHealingPoliciesArgs{
//					HealthCheck:     autohealing.ID(),
//					InitialDelaySec: pulumi.Int(300),
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
// ### With Multiple Versions (`Google-Beta` Provider)
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewInstanceGroupManager(ctx, "appserver", &compute.InstanceGroupManagerArgs{
//				BaseInstanceName: pulumi.String("app"),
//				Zone:             pulumi.String("us-central1-a"),
//				TargetSize:       pulumi.Int(5),
//				Versions: compute.InstanceGroupManagerVersionArray{
//					&compute.InstanceGroupManagerVersionArgs{
//						Name:             pulumi.String("appserver"),
//						InstanceTemplate: pulumi.Any(google_compute_instance_template.Appserver.Id),
//					},
//					&compute.InstanceGroupManagerVersionArgs{
//						Name:             pulumi.String("appserver-canary"),
//						InstanceTemplate: pulumi.Any(google_compute_instance_template.Appserver - canary.Id),
//						TargetSize: &compute.InstanceGroupManagerVersionTargetSizeArgs{
//							Fixed: pulumi.Int(1),
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
// # Instance group managers can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{project}}/{{zone}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceGroupManager:InstanceGroupManager appserver {{name}}
//
// ```
type InstanceGroupManager struct {
	pulumi.CustomResourceState

	// )
	// Properties to set on all instances in the group. After setting
	// allInstancesConfig on the group, you must update the group's instances to
	// apply the configuration.
	AllInstancesConfig InstanceGroupManagerAllInstancesConfigPtrOutput `pulumi:"allInstancesConfig"`
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrOutput `pulumi:"autoHealingPolicies"`
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
	// The fingerprint of the instance group manager.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The full URL of the instance group created by the manager.
	InstanceGroup pulumi.StringOutput `pulumi:"instanceGroup"`
	// - Version name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayOutput `pulumi:"namedPorts"`
	Operation  pulumi.StringOutput                      `pulumi:"operation"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URL of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayOutput `pulumi:"statefulDisks"`
	// The status of this managed instance group.
	Statuses InstanceGroupManagerStatusArrayOutput `pulumi:"statuses"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayOutput `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntOutput `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyOutput `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayOutput `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrOutput `pulumi:"waitForInstances"`
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus pulumi.StringPtrOutput `pulumi:"waitForInstancesStatus"`
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroupManager registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroupManager(ctx *pulumi.Context,
	name string, args *InstanceGroupManagerArgs, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'BaseInstanceName'")
	}
	if args.Versions == nil {
		return nil, errors.New("invalid value for required argument 'Versions'")
	}
	var resource InstanceGroupManager
	err := ctx.RegisterResource("gcp:compute/instanceGroupManager:InstanceGroupManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroupManager gets an existing InstanceGroupManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroupManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupManagerState, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	var resource InstanceGroupManager
	err := ctx.ReadResource("gcp:compute/instanceGroupManager:InstanceGroupManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroupManager resources.
type instanceGroupManagerState struct {
	// )
	// Properties to set on all instances in the group. After setting
	// allInstancesConfig on the group, you must update the group's instances to
	// apply the configuration.
	AllInstancesConfig *InstanceGroupManagerAllInstancesConfig `pulumi:"allInstancesConfig"`
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies *InstanceGroupManagerAutoHealingPolicies `pulumi:"autoHealingPolicies"`
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
	// The fingerprint of the instance group manager.
	Fingerprint *string `pulumi:"fingerprint"`
	// The full URL of the instance group created by the manager.
	InstanceGroup *string `pulumi:"instanceGroup"`
	// - Version name.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupManagerNamedPort `pulumi:"namedPorts"`
	Operation  *string                         `pulumi:"operation"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URL of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks []InstanceGroupManagerStatefulDisk `pulumi:"statefulDisks"`
	// The status of this managed instance group.
	Statuses []InstanceGroupManagerStatus `pulumi:"statuses"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools []string `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions []InstanceGroupManagerVersion `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances *bool `pulumi:"waitForInstances"`
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus *string `pulumi:"waitForInstancesStatus"`
	// The zone that instances in this group should be created
	// in.
	Zone *string `pulumi:"zone"`
}

type InstanceGroupManagerState struct {
	// )
	// Properties to set on all instances in the group. After setting
	// allInstancesConfig on the group, you must update the group's instances to
	// apply the configuration.
	AllInstancesConfig InstanceGroupManagerAllInstancesConfigPtrInput
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrInput
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
	// The fingerprint of the instance group manager.
	Fingerprint pulumi.StringPtrInput
	// The full URL of the instance group created by the manager.
	InstanceGroup pulumi.StringPtrInput
	// - Version name.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayInput
	Operation  pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URL of the created resource.
	SelfLink pulumi.StringPtrInput
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayInput
	// The status of this managed instance group.
	Statuses InstanceGroupManagerStatusArrayInput
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayInput
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyPtrInput
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayInput
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrInput
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus pulumi.StringPtrInput
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupManagerState)(nil)).Elem()
}

type instanceGroupManagerArgs struct {
	// )
	// Properties to set on all instances in the group. After setting
	// allInstancesConfig on the group, you must update the group's instances to
	// apply the configuration.
	AllInstancesConfig *InstanceGroupManagerAllInstancesConfig `pulumi:"allInstancesConfig"`
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies *InstanceGroupManagerAutoHealingPolicies `pulumi:"autoHealingPolicies"`
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
	// - Version name.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupManagerNamedPort `pulumi:"namedPorts"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks []InstanceGroupManagerStatefulDisk `pulumi:"statefulDisks"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools []string `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions []InstanceGroupManagerVersion `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances *bool `pulumi:"waitForInstances"`
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus *string `pulumi:"waitForInstancesStatus"`
	// The zone that instances in this group should be created
	// in.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroupManager resource.
type InstanceGroupManagerArgs struct {
	// )
	// Properties to set on all instances in the group. After setting
	// allInstancesConfig on the group, you must update the group's instances to
	// apply the configuration.
	AllInstancesConfig InstanceGroupManagerAllInstancesConfigPtrInput
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrInput
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
	// - Version name.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayInput
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayInput
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyPtrInput
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayInput
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrInput
	// When used with `waitForInstances` it specifies the status to wait for.
	// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	WaitForInstancesStatus pulumi.StringPtrInput
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupManagerArgs)(nil)).Elem()
}

type InstanceGroupManagerInput interface {
	pulumi.Input

	ToInstanceGroupManagerOutput() InstanceGroupManagerOutput
	ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput
}

func (*InstanceGroupManager) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupManager)(nil)).Elem()
}

func (i *InstanceGroupManager) ToInstanceGroupManagerOutput() InstanceGroupManagerOutput {
	return i.ToInstanceGroupManagerOutputWithContext(context.Background())
}

func (i *InstanceGroupManager) ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerOutput)
}

// InstanceGroupManagerArrayInput is an input type that accepts InstanceGroupManagerArray and InstanceGroupManagerArrayOutput values.
// You can construct a concrete instance of `InstanceGroupManagerArrayInput` via:
//
//	InstanceGroupManagerArray{ InstanceGroupManagerArgs{...} }
type InstanceGroupManagerArrayInput interface {
	pulumi.Input

	ToInstanceGroupManagerArrayOutput() InstanceGroupManagerArrayOutput
	ToInstanceGroupManagerArrayOutputWithContext(context.Context) InstanceGroupManagerArrayOutput
}

type InstanceGroupManagerArray []InstanceGroupManagerInput

func (InstanceGroupManagerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceGroupManager)(nil)).Elem()
}

func (i InstanceGroupManagerArray) ToInstanceGroupManagerArrayOutput() InstanceGroupManagerArrayOutput {
	return i.ToInstanceGroupManagerArrayOutputWithContext(context.Background())
}

func (i InstanceGroupManagerArray) ToInstanceGroupManagerArrayOutputWithContext(ctx context.Context) InstanceGroupManagerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerArrayOutput)
}

// InstanceGroupManagerMapInput is an input type that accepts InstanceGroupManagerMap and InstanceGroupManagerMapOutput values.
// You can construct a concrete instance of `InstanceGroupManagerMapInput` via:
//
//	InstanceGroupManagerMap{ "key": InstanceGroupManagerArgs{...} }
type InstanceGroupManagerMapInput interface {
	pulumi.Input

	ToInstanceGroupManagerMapOutput() InstanceGroupManagerMapOutput
	ToInstanceGroupManagerMapOutputWithContext(context.Context) InstanceGroupManagerMapOutput
}

type InstanceGroupManagerMap map[string]InstanceGroupManagerInput

func (InstanceGroupManagerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceGroupManager)(nil)).Elem()
}

func (i InstanceGroupManagerMap) ToInstanceGroupManagerMapOutput() InstanceGroupManagerMapOutput {
	return i.ToInstanceGroupManagerMapOutputWithContext(context.Background())
}

func (i InstanceGroupManagerMap) ToInstanceGroupManagerMapOutputWithContext(ctx context.Context) InstanceGroupManagerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupManagerMapOutput)
}

type InstanceGroupManagerOutput struct{ *pulumi.OutputState }

func (InstanceGroupManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupManager)(nil)).Elem()
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerOutput() InstanceGroupManagerOutput {
	return o
}

func (o InstanceGroupManagerOutput) ToInstanceGroupManagerOutputWithContext(ctx context.Context) InstanceGroupManagerOutput {
	return o
}

// )
// Properties to set on all instances in the group. After setting
// allInstancesConfig on the group, you must update the group's instances to
// apply the configuration.
func (o InstanceGroupManagerOutput) AllInstancesConfig() InstanceGroupManagerAllInstancesConfigPtrOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerAllInstancesConfigPtrOutput {
		return v.AllInstancesConfig
	}).(InstanceGroupManagerAllInstancesConfigPtrOutput)
}

// The autohealing policies for this managed instance
// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
func (o InstanceGroupManagerOutput) AutoHealingPolicies() InstanceGroupManagerAutoHealingPoliciesPtrOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerAutoHealingPoliciesPtrOutput {
		return v.AutoHealingPolicies
	}).(InstanceGroupManagerAutoHealingPoliciesPtrOutput)
}

// The base instance name to use for
// instances in this group. The value must be a valid
// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
// are lowercase letters, numbers, and hyphens (-). Instances are named by
// appending a hyphen and a random four-character string to the base instance
// name.
func (o InstanceGroupManagerOutput) BaseInstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.BaseInstanceName }).(pulumi.StringOutput)
}

// An optional textual description of the instance
// group manager.
func (o InstanceGroupManagerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The fingerprint of the instance group manager.
func (o InstanceGroupManagerOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The full URL of the instance group created by the manager.
func (o InstanceGroupManagerOutput) InstanceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.InstanceGroup }).(pulumi.StringOutput)
}

// - Version name.
func (o InstanceGroupManagerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The named port configuration. See the section below
// for details on configuration.
func (o InstanceGroupManagerOutput) NamedPorts() InstanceGroupManagerNamedPortArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerNamedPortArrayOutput { return v.NamedPorts }).(InstanceGroupManagerNamedPortArrayOutput)
}

func (o InstanceGroupManagerOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Operation }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (o InstanceGroupManagerOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URL of the created resource.
func (o InstanceGroupManagerOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
func (o InstanceGroupManagerOutput) StatefulDisks() InstanceGroupManagerStatefulDiskArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerStatefulDiskArrayOutput { return v.StatefulDisks }).(InstanceGroupManagerStatefulDiskArrayOutput)
}

// The status of this managed instance group.
func (o InstanceGroupManagerOutput) Statuses() InstanceGroupManagerStatusArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerStatusArrayOutput { return v.Statuses }).(InstanceGroupManagerStatusArrayOutput)
}

// The full URL of all target pools to which new
// instances in the group are added. Updating the target pools attribute does
// not affect existing instances.
func (o InstanceGroupManagerOutput) TargetPools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringArrayOutput { return v.TargetPools }).(pulumi.StringArrayOutput)
}

// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
func (o InstanceGroupManagerOutput) TargetSize() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.IntOutput { return v.TargetSize }).(pulumi.IntOutput)
}

// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
func (o InstanceGroupManagerOutput) UpdatePolicy() InstanceGroupManagerUpdatePolicyOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerUpdatePolicyOutput { return v.UpdatePolicy }).(InstanceGroupManagerUpdatePolicyOutput)
}

// Application versions managed by this instance group. Each
// version deals with a specific instance template, allowing canary release scenarios.
// Structure is documented below.
func (o InstanceGroupManagerOutput) Versions() InstanceGroupManagerVersionArrayOutput {
	return o.ApplyT(func(v *InstanceGroupManager) InstanceGroupManagerVersionArrayOutput { return v.Versions }).(InstanceGroupManagerVersionArrayOutput)
}

// Whether to wait for all instances to be created/updated before
// returning. Note that if this is set to true and the operation does not succeed, this provider will
// continue trying until it times out.
func (o InstanceGroupManagerOutput) WaitForInstances() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.BoolPtrOutput { return v.WaitForInstances }).(pulumi.BoolPtrOutput)
}

// When used with `waitForInstances` it specifies the status to wait for.
// When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
// set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
// instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
func (o InstanceGroupManagerOutput) WaitForInstancesStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringPtrOutput { return v.WaitForInstancesStatus }).(pulumi.StringPtrOutput)
}

// The zone that instances in this group should be created
// in.
func (o InstanceGroupManagerOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupManager) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceGroupManagerArrayOutput struct{ *pulumi.OutputState }

func (InstanceGroupManagerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceGroupManager)(nil)).Elem()
}

func (o InstanceGroupManagerArrayOutput) ToInstanceGroupManagerArrayOutput() InstanceGroupManagerArrayOutput {
	return o
}

func (o InstanceGroupManagerArrayOutput) ToInstanceGroupManagerArrayOutputWithContext(ctx context.Context) InstanceGroupManagerArrayOutput {
	return o
}

func (o InstanceGroupManagerArrayOutput) Index(i pulumi.IntInput) InstanceGroupManagerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceGroupManager {
		return vs[0].([]*InstanceGroupManager)[vs[1].(int)]
	}).(InstanceGroupManagerOutput)
}

type InstanceGroupManagerMapOutput struct{ *pulumi.OutputState }

func (InstanceGroupManagerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceGroupManager)(nil)).Elem()
}

func (o InstanceGroupManagerMapOutput) ToInstanceGroupManagerMapOutput() InstanceGroupManagerMapOutput {
	return o
}

func (o InstanceGroupManagerMapOutput) ToInstanceGroupManagerMapOutputWithContext(ctx context.Context) InstanceGroupManagerMapOutput {
	return o
}

func (o InstanceGroupManagerMapOutput) MapIndex(k pulumi.StringInput) InstanceGroupManagerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceGroupManager {
		return vs[0].(map[string]*InstanceGroupManager)[vs[1].(string)]
	}).(InstanceGroupManagerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupManagerInput)(nil)).Elem(), &InstanceGroupManager{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupManagerArrayInput)(nil)).Elem(), InstanceGroupManagerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupManagerMapInput)(nil)).Elem(), InstanceGroupManagerMap{})
	pulumi.RegisterOutputType(InstanceGroupManagerOutput{})
	pulumi.RegisterOutputType(InstanceGroupManagerArrayOutput{})
	pulumi.RegisterOutputType(InstanceGroupManagerMapOutput{})
}
