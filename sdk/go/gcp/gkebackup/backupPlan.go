// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkebackup

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a Backup Plan instance.
//
// To get more information about BackupPlan, see:
//
// * [API documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/projects.locations.backupPlans)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke)
//
// ## Example Usage
// ### Gkebackup Backupplan Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
//				Location:         pulumi.String("us-central1"),
//				InitialNodeCount: pulumi.Int(1),
//				WorkloadIdentityConfig: &container.ClusterWorkloadIdentityConfigArgs{
//					WorkloadPool: pulumi.String("my-project-name.svc.id.goog"),
//				},
//				AddonsConfig: &container.ClusterAddonsConfigArgs{
//					GkeBackupAgentConfig: &container.ClusterAddonsConfigGkeBackupAgentConfigArgs{
//						Enabled: pulumi.Bool(true),
//					},
//				},
//				DeletionProtection: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gkebackup.NewBackupPlan(ctx, "basic", &gkebackup.BackupPlanArgs{
//				Cluster:  primary.ID(),
//				Location: pulumi.String("us-central1"),
//				BackupConfig: &gkebackup.BackupPlanBackupConfigArgs{
//					IncludeVolumeData: pulumi.Bool(true),
//					IncludeSecrets:    pulumi.Bool(true),
//					AllNamespaces:     pulumi.Bool(true),
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
// ### Gkebackup Backupplan Autopilot
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
//				Location:           pulumi.String("us-central1"),
//				EnableAutopilot:    pulumi.Bool(true),
//				IpAllocationPolicy: nil,
//				ReleaseChannel: &container.ClusterReleaseChannelArgs{
//					Channel: pulumi.String("RAPID"),
//				},
//				AddonsConfig: &container.ClusterAddonsConfigArgs{
//					GkeBackupAgentConfig: &container.ClusterAddonsConfigGkeBackupAgentConfigArgs{
//						Enabled: pulumi.Bool(true),
//					},
//				},
//				DeletionProtection: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gkebackup.NewBackupPlan(ctx, "autopilot", &gkebackup.BackupPlanArgs{
//				Cluster:  primary.ID(),
//				Location: pulumi.String("us-central1"),
//				BackupConfig: &gkebackup.BackupPlanBackupConfigArgs{
//					IncludeVolumeData: pulumi.Bool(true),
//					IncludeSecrets:    pulumi.Bool(true),
//					AllNamespaces:     pulumi.Bool(true),
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
// ### Gkebackup Backupplan Cmek
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/kms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
//				Location:         pulumi.String("us-central1"),
//				InitialNodeCount: pulumi.Int(1),
//				WorkloadIdentityConfig: &container.ClusterWorkloadIdentityConfigArgs{
//					WorkloadPool: pulumi.String("my-project-name.svc.id.goog"),
//				},
//				AddonsConfig: &container.ClusterAddonsConfigArgs{
//					GkeBackupAgentConfig: &container.ClusterAddonsConfigGkeBackupAgentConfigArgs{
//						Enabled: pulumi.Bool(true),
//					},
//				},
//				DeletionProtection: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			keyRing, err := kms.NewKeyRing(ctx, "keyRing", &kms.KeyRingArgs{
//				Location: pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			cryptoKey, err := kms.NewCryptoKey(ctx, "cryptoKey", &kms.CryptoKeyArgs{
//				KeyRing: keyRing.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gkebackup.NewBackupPlan(ctx, "cmek", &gkebackup.BackupPlanArgs{
//				Cluster:  primary.ID(),
//				Location: pulumi.String("us-central1"),
//				BackupConfig: &gkebackup.BackupPlanBackupConfigArgs{
//					IncludeVolumeData: pulumi.Bool(true),
//					IncludeSecrets:    pulumi.Bool(true),
//					SelectedNamespaces: &gkebackup.BackupPlanBackupConfigSelectedNamespacesArgs{
//						Namespaces: pulumi.StringArray{
//							pulumi.String("default"),
//							pulumi.String("test"),
//						},
//					},
//					EncryptionKey: &gkebackup.BackupPlanBackupConfigEncryptionKeyArgs{
//						GcpKmsEncryptionKey: cryptoKey.ID(),
//					},
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
// ### Gkebackup Backupplan Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkebackup"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
//				Location:         pulumi.String("us-central1"),
//				InitialNodeCount: pulumi.Int(1),
//				WorkloadIdentityConfig: &container.ClusterWorkloadIdentityConfigArgs{
//					WorkloadPool: pulumi.String("my-project-name.svc.id.goog"),
//				},
//				AddonsConfig: &container.ClusterAddonsConfigArgs{
//					GkeBackupAgentConfig: &container.ClusterAddonsConfigGkeBackupAgentConfigArgs{
//						Enabled: pulumi.Bool(true),
//					},
//				},
//				DeletionProtection: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = gkebackup.NewBackupPlan(ctx, "full", &gkebackup.BackupPlanArgs{
//				Cluster:  primary.ID(),
//				Location: pulumi.String("us-central1"),
//				RetentionPolicy: &gkebackup.BackupPlanRetentionPolicyArgs{
//					BackupDeleteLockDays: pulumi.Int(30),
//					BackupRetainDays:     pulumi.Int(180),
//				},
//				BackupSchedule: &gkebackup.BackupPlanBackupScheduleArgs{
//					CronSchedule: pulumi.String("0 9 * * 1"),
//				},
//				BackupConfig: &gkebackup.BackupPlanBackupConfigArgs{
//					IncludeVolumeData: pulumi.Bool(true),
//					IncludeSecrets:    pulumi.Bool(true),
//					SelectedApplications: &gkebackup.BackupPlanBackupConfigSelectedApplicationsArgs{
//						NamespacedNames: gkebackup.BackupPlanBackupConfigSelectedApplicationsNamespacedNameArray{
//							&gkebackup.BackupPlanBackupConfigSelectedApplicationsNamespacedNameArgs{
//								Name:      pulumi.String("app1"),
//								Namespace: pulumi.String("ns1"),
//							},
//							&gkebackup.BackupPlanBackupConfigSelectedApplicationsNamespacedNameArgs{
//								Name:      pulumi.String("app2"),
//								Namespace: pulumi.String("ns2"),
//							},
//						},
//					},
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
//
// ## Import
//
// # BackupPlan can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:gkebackup/backupPlan:BackupPlan default projects/{{project}}/locations/{{location}}/backupPlans/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkebackup/backupPlan:BackupPlan default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkebackup/backupPlan:BackupPlan default {{location}}/{{name}}
//
// ```
type BackupPlan struct {
	pulumi.CustomResourceState

	// Defines the configuration of Backups created via this BackupPlan.
	// Structure is documented below.
	BackupConfig BackupPlanBackupConfigPtrOutput `pulumi:"backupConfig"`
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	// Structure is documented below.
	BackupSchedule BackupPlanBackupSchedulePtrOutput `pulumi:"backupSchedule"`
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	Deactivated pulumi.BoolOutput `pulumi:"deactivated"`
	// User specified descriptive string for this BackupPlan.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous
	// updates of a backup plan from overwriting each other. It is strongly suggested that
	// systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates
	// in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
	// and systems are expected to put that etag in the request to backupPlans.patch or
	// backupPlans.delete to ensure that their change will be applied to the same version of the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The region of the Backup Plan.
	//
	// ***
	Location pulumi.StringOutput `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
	ProtectedPodCount pulumi.IntOutput `pulumi:"protectedPodCount"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	// Structure is documented below.
	RetentionPolicy BackupPlanRetentionPolicyPtrOutput `pulumi:"retentionPolicy"`
	// The State of the BackupPlan.
	State pulumi.StringOutput `pulumi:"state"`
	// Detailed description of why BackupPlan is in its current state.
	StateReason pulumi.StringOutput `pulumi:"stateReason"`
	// Server generated, unique identifier of UUID format.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewBackupPlan registers a new resource with the given unique name, arguments, and options.
func NewBackupPlan(ctx *pulumi.Context,
	name string, args *BackupPlanArgs, opts ...pulumi.ResourceOption) (*BackupPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
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
	var resource BackupPlan
	err := ctx.RegisterResource("gcp:gkebackup/backupPlan:BackupPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPlan gets an existing BackupPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPlanState, opts ...pulumi.ResourceOption) (*BackupPlan, error) {
	var resource BackupPlan
	err := ctx.ReadResource("gcp:gkebackup/backupPlan:BackupPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPlan resources.
type backupPlanState struct {
	// Defines the configuration of Backups created via this BackupPlan.
	// Structure is documented below.
	BackupConfig *BackupPlanBackupConfig `pulumi:"backupConfig"`
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	// Structure is documented below.
	BackupSchedule *BackupPlanBackupSchedule `pulumi:"backupSchedule"`
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster *string `pulumi:"cluster"`
	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	Deactivated *bool `pulumi:"deactivated"`
	// User specified descriptive string for this BackupPlan.
	Description *string `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous
	// updates of a backup plan from overwriting each other. It is strongly suggested that
	// systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates
	// in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
	// and systems are expected to put that etag in the request to backupPlans.patch or
	// backupPlans.delete to ensure that their change will be applied to the same version of the resource.
	Etag *string `pulumi:"etag"`
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The region of the Backup Plan.
	//
	// ***
	Location *string `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
	ProtectedPodCount *int `pulumi:"protectedPodCount"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	// Structure is documented below.
	RetentionPolicy *BackupPlanRetentionPolicy `pulumi:"retentionPolicy"`
	// The State of the BackupPlan.
	State *string `pulumi:"state"`
	// Detailed description of why BackupPlan is in its current state.
	StateReason *string `pulumi:"stateReason"`
	// Server generated, unique identifier of UUID format.
	Uid *string `pulumi:"uid"`
}

type BackupPlanState struct {
	// Defines the configuration of Backups created via this BackupPlan.
	// Structure is documented below.
	BackupConfig BackupPlanBackupConfigPtrInput
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	// Structure is documented below.
	BackupSchedule BackupPlanBackupSchedulePtrInput
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster pulumi.StringPtrInput
	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	Deactivated pulumi.BoolPtrInput
	// User specified descriptive string for this BackupPlan.
	Description pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapInput
	// etag is used for optimistic concurrency control as a way to help prevent simultaneous
	// updates of a backup plan from overwriting each other. It is strongly suggested that
	// systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates
	// in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
	// and systems are expected to put that etag in the request to backupPlans.patch or
	// backupPlans.delete to ensure that their change will be applied to the same version of the resource.
	Etag pulumi.StringPtrInput
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The region of the Backup Plan.
	//
	// ***
	Location pulumi.StringPtrInput
	// The full name of the BackupPlan Resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
	ProtectedPodCount pulumi.IntPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	// Structure is documented below.
	RetentionPolicy BackupPlanRetentionPolicyPtrInput
	// The State of the BackupPlan.
	State pulumi.StringPtrInput
	// Detailed description of why BackupPlan is in its current state.
	StateReason pulumi.StringPtrInput
	// Server generated, unique identifier of UUID format.
	Uid pulumi.StringPtrInput
}

func (BackupPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPlanState)(nil)).Elem()
}

type backupPlanArgs struct {
	// Defines the configuration of Backups created via this BackupPlan.
	// Structure is documented below.
	BackupConfig *BackupPlanBackupConfig `pulumi:"backupConfig"`
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	// Structure is documented below.
	BackupSchedule *BackupPlanBackupSchedule `pulumi:"backupSchedule"`
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster string `pulumi:"cluster"`
	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	Deactivated *bool `pulumi:"deactivated"`
	// User specified descriptive string for this BackupPlan.
	Description *string `pulumi:"description"`
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The region of the Backup Plan.
	//
	// ***
	Location string `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	// Structure is documented below.
	RetentionPolicy *BackupPlanRetentionPolicy `pulumi:"retentionPolicy"`
}

// The set of arguments for constructing a BackupPlan resource.
type BackupPlanArgs struct {
	// Defines the configuration of Backups created via this BackupPlan.
	// Structure is documented below.
	BackupConfig BackupPlanBackupConfigPtrInput
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	// Structure is documented below.
	BackupSchedule BackupPlanBackupSchedulePtrInput
	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster pulumi.StringInput
	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	Deactivated pulumi.BoolPtrInput
	// User specified descriptive string for this BackupPlan.
	Description pulumi.StringPtrInput
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The region of the Backup Plan.
	//
	// ***
	Location pulumi.StringInput
	// The full name of the BackupPlan Resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	// Structure is documented below.
	RetentionPolicy BackupPlanRetentionPolicyPtrInput
}

func (BackupPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPlanArgs)(nil)).Elem()
}

type BackupPlanInput interface {
	pulumi.Input

	ToBackupPlanOutput() BackupPlanOutput
	ToBackupPlanOutputWithContext(ctx context.Context) BackupPlanOutput
}

func (*BackupPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPlan)(nil)).Elem()
}

func (i *BackupPlan) ToBackupPlanOutput() BackupPlanOutput {
	return i.ToBackupPlanOutputWithContext(context.Background())
}

func (i *BackupPlan) ToBackupPlanOutputWithContext(ctx context.Context) BackupPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanOutput)
}

func (i *BackupPlan) ToOutput(ctx context.Context) pulumix.Output[*BackupPlan] {
	return pulumix.Output[*BackupPlan]{
		OutputState: i.ToBackupPlanOutputWithContext(ctx).OutputState,
	}
}

// BackupPlanArrayInput is an input type that accepts BackupPlanArray and BackupPlanArrayOutput values.
// You can construct a concrete instance of `BackupPlanArrayInput` via:
//
//	BackupPlanArray{ BackupPlanArgs{...} }
type BackupPlanArrayInput interface {
	pulumi.Input

	ToBackupPlanArrayOutput() BackupPlanArrayOutput
	ToBackupPlanArrayOutputWithContext(context.Context) BackupPlanArrayOutput
}

type BackupPlanArray []BackupPlanInput

func (BackupPlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPlan)(nil)).Elem()
}

func (i BackupPlanArray) ToBackupPlanArrayOutput() BackupPlanArrayOutput {
	return i.ToBackupPlanArrayOutputWithContext(context.Background())
}

func (i BackupPlanArray) ToBackupPlanArrayOutputWithContext(ctx context.Context) BackupPlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanArrayOutput)
}

func (i BackupPlanArray) ToOutput(ctx context.Context) pulumix.Output[[]*BackupPlan] {
	return pulumix.Output[[]*BackupPlan]{
		OutputState: i.ToBackupPlanArrayOutputWithContext(ctx).OutputState,
	}
}

// BackupPlanMapInput is an input type that accepts BackupPlanMap and BackupPlanMapOutput values.
// You can construct a concrete instance of `BackupPlanMapInput` via:
//
//	BackupPlanMap{ "key": BackupPlanArgs{...} }
type BackupPlanMapInput interface {
	pulumi.Input

	ToBackupPlanMapOutput() BackupPlanMapOutput
	ToBackupPlanMapOutputWithContext(context.Context) BackupPlanMapOutput
}

type BackupPlanMap map[string]BackupPlanInput

func (BackupPlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPlan)(nil)).Elem()
}

func (i BackupPlanMap) ToBackupPlanMapOutput() BackupPlanMapOutput {
	return i.ToBackupPlanMapOutputWithContext(context.Background())
}

func (i BackupPlanMap) ToBackupPlanMapOutputWithContext(ctx context.Context) BackupPlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPlanMapOutput)
}

func (i BackupPlanMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BackupPlan] {
	return pulumix.Output[map[string]*BackupPlan]{
		OutputState: i.ToBackupPlanMapOutputWithContext(ctx).OutputState,
	}
}

type BackupPlanOutput struct{ *pulumi.OutputState }

func (BackupPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPlan)(nil)).Elem()
}

func (o BackupPlanOutput) ToBackupPlanOutput() BackupPlanOutput {
	return o
}

func (o BackupPlanOutput) ToBackupPlanOutputWithContext(ctx context.Context) BackupPlanOutput {
	return o
}

func (o BackupPlanOutput) ToOutput(ctx context.Context) pulumix.Output[*BackupPlan] {
	return pulumix.Output[*BackupPlan]{
		OutputState: o.OutputState,
	}
}

// Defines the configuration of Backups created via this BackupPlan.
// Structure is documented below.
func (o BackupPlanOutput) BackupConfig() BackupPlanBackupConfigPtrOutput {
	return o.ApplyT(func(v *BackupPlan) BackupPlanBackupConfigPtrOutput { return v.BackupConfig }).(BackupPlanBackupConfigPtrOutput)
}

// Defines a schedule for automatic Backup creation via this BackupPlan.
// Structure is documented below.
func (o BackupPlanOutput) BackupSchedule() BackupPlanBackupSchedulePtrOutput {
	return o.ApplyT(func(v *BackupPlan) BackupPlanBackupSchedulePtrOutput { return v.BackupSchedule }).(BackupPlanBackupSchedulePtrOutput)
}

// The source cluster from which Backups will be created via this BackupPlan.
func (o BackupPlanOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

// This flag indicates whether this BackupPlan has been deactivated.
// Setting this field to True locks the BackupPlan such that no further updates will be allowed
// (except deletes), including the deactivated field itself. It also prevents any new Backups
// from being created via this BackupPlan (including scheduled Backups).
func (o BackupPlanOutput) Deactivated() pulumi.BoolOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.BoolOutput { return v.Deactivated }).(pulumi.BoolOutput)
}

// User specified descriptive string for this BackupPlan.
func (o BackupPlanOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o BackupPlanOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// etag is used for optimistic concurrency control as a way to help prevent simultaneous
// updates of a backup plan from overwriting each other. It is strongly suggested that
// systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates
// in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
// and systems are expected to put that etag in the request to backupPlans.patch or
// backupPlans.delete to ensure that their change will be applied to the same version of the resource.
func (o BackupPlanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Description: A set of custom labels supplied by the user.
// A list of key->value pairs.
// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o BackupPlanOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The region of the Backup Plan.
//
// ***
func (o BackupPlanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The full name of the BackupPlan Resource.
func (o BackupPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o BackupPlanOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
func (o BackupPlanOutput) ProtectedPodCount() pulumi.IntOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.IntOutput { return v.ProtectedPodCount }).(pulumi.IntOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o BackupPlanOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// RetentionPolicy governs lifecycle of Backups created under this plan.
// Structure is documented below.
func (o BackupPlanOutput) RetentionPolicy() BackupPlanRetentionPolicyPtrOutput {
	return o.ApplyT(func(v *BackupPlan) BackupPlanRetentionPolicyPtrOutput { return v.RetentionPolicy }).(BackupPlanRetentionPolicyPtrOutput)
}

// The State of the BackupPlan.
func (o BackupPlanOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Detailed description of why BackupPlan is in its current state.
func (o BackupPlanOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringOutput { return v.StateReason }).(pulumi.StringOutput)
}

// Server generated, unique identifier of UUID format.
func (o BackupPlanOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupPlan) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

type BackupPlanArrayOutput struct{ *pulumi.OutputState }

func (BackupPlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupPlan)(nil)).Elem()
}

func (o BackupPlanArrayOutput) ToBackupPlanArrayOutput() BackupPlanArrayOutput {
	return o
}

func (o BackupPlanArrayOutput) ToBackupPlanArrayOutputWithContext(ctx context.Context) BackupPlanArrayOutput {
	return o
}

func (o BackupPlanArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BackupPlan] {
	return pulumix.Output[[]*BackupPlan]{
		OutputState: o.OutputState,
	}
}

func (o BackupPlanArrayOutput) Index(i pulumi.IntInput) BackupPlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupPlan {
		return vs[0].([]*BackupPlan)[vs[1].(int)]
	}).(BackupPlanOutput)
}

type BackupPlanMapOutput struct{ *pulumi.OutputState }

func (BackupPlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupPlan)(nil)).Elem()
}

func (o BackupPlanMapOutput) ToBackupPlanMapOutput() BackupPlanMapOutput {
	return o
}

func (o BackupPlanMapOutput) ToBackupPlanMapOutputWithContext(ctx context.Context) BackupPlanMapOutput {
	return o
}

func (o BackupPlanMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BackupPlan] {
	return pulumix.Output[map[string]*BackupPlan]{
		OutputState: o.OutputState,
	}
}

func (o BackupPlanMapOutput) MapIndex(k pulumi.StringInput) BackupPlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupPlan {
		return vs[0].(map[string]*BackupPlan)[vs[1].(string)]
	}).(BackupPlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanInput)(nil)).Elem(), &BackupPlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanArrayInput)(nil)).Elem(), BackupPlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPlanMapInput)(nil)).Elem(), BackupPlanMap{})
	pulumi.RegisterOutputType(BackupPlanOutput{})
	pulumi.RegisterOutputType(BackupPlanArrayOutput{})
	pulumi.RegisterOutputType(BackupPlanMapOutput{})
}
