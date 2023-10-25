// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkebackup

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a Restore Plan instance.
//
// To get more information about RestorePlan, see:
//
// * [API documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/projects.locations.restorePlans)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke)
//
// ## Example Usage
// ### Gkebackup Restoreplan All Namespaces
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkebackup"
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
//			})
//			if err != nil {
//				return err
//			}
//			basic, err := gkebackup.NewBackupPlan(ctx, "basic", &gkebackup.BackupPlanArgs{
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
//			_, err = gkebackup.NewRestorePlan(ctx, "allNs", &gkebackup.RestorePlanArgs{
//				Location:   pulumi.String("us-central1"),
//				BackupPlan: basic.ID(),
//				Cluster:    primary.ID(),
//				RestoreConfig: &gkebackup.RestorePlanRestoreConfigArgs{
//					AllNamespaces:                 pulumi.Bool(true),
//					NamespacedResourceRestoreMode: pulumi.String("FAIL_ON_CONFLICT"),
//					VolumeDataRestorePolicy:       pulumi.String("RESTORE_VOLUME_DATA_FROM_BACKUP"),
//					ClusterResourceRestoreScope: &gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs{
//						AllGroupKinds: pulumi.Bool(true),
//					},
//					ClusterResourceConflictPolicy: pulumi.String("USE_EXISTING_VERSION"),
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
// ### Gkebackup Restoreplan Rollback Namespace
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkebackup"
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
//			})
//			if err != nil {
//				return err
//			}
//			basic, err := gkebackup.NewBackupPlan(ctx, "basic", &gkebackup.BackupPlanArgs{
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
//			_, err = gkebackup.NewRestorePlan(ctx, "rollbackNs", &gkebackup.RestorePlanArgs{
//				Location:   pulumi.String("us-central1"),
//				BackupPlan: basic.ID(),
//				Cluster:    primary.ID(),
//				RestoreConfig: &gkebackup.RestorePlanRestoreConfigArgs{
//					SelectedNamespaces: &gkebackup.RestorePlanRestoreConfigSelectedNamespacesArgs{
//						Namespaces: pulumi.StringArray{
//							pulumi.String("my-ns"),
//						},
//					},
//					NamespacedResourceRestoreMode: pulumi.String("DELETE_AND_RESTORE"),
//					VolumeDataRestorePolicy:       pulumi.String("RESTORE_VOLUME_DATA_FROM_BACKUP"),
//					ClusterResourceRestoreScope: &gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs{
//						SelectedGroupKinds: gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKindArray{
//							&gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKindArgs{
//								ResourceGroup: pulumi.String("apiextension.k8s.io"),
//								ResourceKind:  pulumi.String("CustomResourceDefinition"),
//							},
//							&gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeSelectedGroupKindArgs{
//								ResourceGroup: pulumi.String("storage.k8s.io"),
//								ResourceKind:  pulumi.String("StorageClass"),
//							},
//						},
//					},
//					ClusterResourceConflictPolicy: pulumi.String("USE_EXISTING_VERSION"),
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
// ### Gkebackup Restoreplan Protected Application
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkebackup"
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
//			})
//			if err != nil {
//				return err
//			}
//			basic, err := gkebackup.NewBackupPlan(ctx, "basic", &gkebackup.BackupPlanArgs{
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
//			_, err = gkebackup.NewRestorePlan(ctx, "rollbackApp", &gkebackup.RestorePlanArgs{
//				Location:   pulumi.String("us-central1"),
//				BackupPlan: basic.ID(),
//				Cluster:    primary.ID(),
//				RestoreConfig: &gkebackup.RestorePlanRestoreConfigArgs{
//					SelectedApplications: &gkebackup.RestorePlanRestoreConfigSelectedApplicationsArgs{
//						NamespacedNames: gkebackup.RestorePlanRestoreConfigSelectedApplicationsNamespacedNameArray{
//							&gkebackup.RestorePlanRestoreConfigSelectedApplicationsNamespacedNameArgs{
//								Name:      pulumi.String("my-app"),
//								Namespace: pulumi.String("my-ns"),
//							},
//						},
//					},
//					NamespacedResourceRestoreMode: pulumi.String("DELETE_AND_RESTORE"),
//					VolumeDataRestorePolicy:       pulumi.String("REUSE_VOLUME_HANDLE_FROM_BACKUP"),
//					ClusterResourceRestoreScope: &gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs{
//						NoGroupKinds: pulumi.Bool(true),
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
// ### Gkebackup Restoreplan All Cluster Resources
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkebackup"
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
//			})
//			if err != nil {
//				return err
//			}
//			basic, err := gkebackup.NewBackupPlan(ctx, "basic", &gkebackup.BackupPlanArgs{
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
//			_, err = gkebackup.NewRestorePlan(ctx, "allClusterResources", &gkebackup.RestorePlanArgs{
//				Location:   pulumi.String("us-central1"),
//				BackupPlan: basic.ID(),
//				Cluster:    primary.ID(),
//				RestoreConfig: &gkebackup.RestorePlanRestoreConfigArgs{
//					NoNamespaces:                  pulumi.Bool(true),
//					NamespacedResourceRestoreMode: pulumi.String("FAIL_ON_CONFLICT"),
//					ClusterResourceRestoreScope: &gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs{
//						AllGroupKinds: pulumi.Bool(true),
//					},
//					ClusterResourceConflictPolicy: pulumi.String("USE_EXISTING_VERSION"),
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
// ### Gkebackup Restoreplan Rename Namespace
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkebackup"
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
//			})
//			if err != nil {
//				return err
//			}
//			basic, err := gkebackup.NewBackupPlan(ctx, "basic", &gkebackup.BackupPlanArgs{
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
//			_, err = gkebackup.NewRestorePlan(ctx, "renameNs", &gkebackup.RestorePlanArgs{
//				Location:   pulumi.String("us-central1"),
//				BackupPlan: basic.ID(),
//				Cluster:    primary.ID(),
//				RestoreConfig: &gkebackup.RestorePlanRestoreConfigArgs{
//					SelectedNamespaces: &gkebackup.RestorePlanRestoreConfigSelectedNamespacesArgs{
//						Namespaces: pulumi.StringArray{
//							pulumi.String("ns1"),
//						},
//					},
//					NamespacedResourceRestoreMode: pulumi.String("FAIL_ON_CONFLICT"),
//					VolumeDataRestorePolicy:       pulumi.String("REUSE_VOLUME_HANDLE_FROM_BACKUP"),
//					ClusterResourceRestoreScope: &gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs{
//						NoGroupKinds: pulumi.Bool(true),
//					},
//					TransformationRules: gkebackup.RestorePlanRestoreConfigTransformationRuleArray{
//						&gkebackup.RestorePlanRestoreConfigTransformationRuleArgs{
//							Description: pulumi.String("rename namespace from ns1 to ns2"),
//							ResourceFilter: &gkebackup.RestorePlanRestoreConfigTransformationRuleResourceFilterArgs{
//								GroupKinds: gkebackup.RestorePlanRestoreConfigTransformationRuleResourceFilterGroupKindArray{
//									&gkebackup.RestorePlanRestoreConfigTransformationRuleResourceFilterGroupKindArgs{
//										ResourceKind: pulumi.String("Namespace"),
//									},
//								},
//								JsonPath: pulumi.String(".metadata[?(@.name == 'ns1')]"),
//							},
//							FieldActions: gkebackup.RestorePlanRestoreConfigTransformationRuleFieldActionArray{
//								&gkebackup.RestorePlanRestoreConfigTransformationRuleFieldActionArgs{
//									Op:    pulumi.String("REPLACE"),
//									Path:  pulumi.String("/metadata/name"),
//									Value: pulumi.String("ns2"),
//								},
//							},
//						},
//						&gkebackup.RestorePlanRestoreConfigTransformationRuleArgs{
//							Description: pulumi.String("move all resources from ns1 to ns2"),
//							ResourceFilter: &gkebackup.RestorePlanRestoreConfigTransformationRuleResourceFilterArgs{
//								Namespaces: pulumi.StringArray{
//									pulumi.String("ns1"),
//								},
//							},
//							FieldActions: gkebackup.RestorePlanRestoreConfigTransformationRuleFieldActionArray{
//								&gkebackup.RestorePlanRestoreConfigTransformationRuleFieldActionArgs{
//									Op:    pulumi.String("REPLACE"),
//									Path:  pulumi.String("/metadata/namespace"),
//									Value: pulumi.String("ns2"),
//								},
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
// ### Gkebackup Restoreplan Second Transformation
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkebackup"
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
//			})
//			if err != nil {
//				return err
//			}
//			basic, err := gkebackup.NewBackupPlan(ctx, "basic", &gkebackup.BackupPlanArgs{
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
//			_, err = gkebackup.NewRestorePlan(ctx, "transformRule", &gkebackup.RestorePlanArgs{
//				Description: pulumi.String("copy nginx env variables"),
//				Labels: pulumi.StringMap{
//					"app": pulumi.String("nginx"),
//				},
//				Location:   pulumi.String("us-central1"),
//				BackupPlan: basic.ID(),
//				Cluster:    primary.ID(),
//				RestoreConfig: &gkebackup.RestorePlanRestoreConfigArgs{
//					ExcludedNamespaces: &gkebackup.RestorePlanRestoreConfigExcludedNamespacesArgs{
//						Namespaces: pulumi.StringArray{
//							pulumi.String("my-ns"),
//						},
//					},
//					NamespacedResourceRestoreMode: pulumi.String("DELETE_AND_RESTORE"),
//					VolumeDataRestorePolicy:       pulumi.String("RESTORE_VOLUME_DATA_FROM_BACKUP"),
//					ClusterResourceRestoreScope: &gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeArgs{
//						ExcludedGroupKinds: gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeExcludedGroupKindArray{
//							&gkebackup.RestorePlanRestoreConfigClusterResourceRestoreScopeExcludedGroupKindArgs{
//								ResourceGroup: pulumi.String("apiextension.k8s.io"),
//								ResourceKind:  pulumi.String("CustomResourceDefinition"),
//							},
//						},
//					},
//					ClusterResourceConflictPolicy: pulumi.String("USE_EXISTING_VERSION"),
//					TransformationRules: gkebackup.RestorePlanRestoreConfigTransformationRuleArray{
//						&gkebackup.RestorePlanRestoreConfigTransformationRuleArgs{
//							Description: pulumi.String("Copy environment variables from the nginx container to the install init container."),
//							ResourceFilter: &gkebackup.RestorePlanRestoreConfigTransformationRuleResourceFilterArgs{
//								GroupKinds: gkebackup.RestorePlanRestoreConfigTransformationRuleResourceFilterGroupKindArray{
//									&gkebackup.RestorePlanRestoreConfigTransformationRuleResourceFilterGroupKindArgs{
//										ResourceKind:  pulumi.String("Pod"),
//										ResourceGroup: pulumi.String(""),
//									},
//								},
//								JsonPath: pulumi.String(".metadata[?(@.name == 'nginx')]"),
//							},
//							FieldActions: gkebackup.RestorePlanRestoreConfigTransformationRuleFieldActionArray{
//								&gkebackup.RestorePlanRestoreConfigTransformationRuleFieldActionArgs{
//									Op:       pulumi.String("COPY"),
//									Path:     pulumi.String("/spec/initContainers/0/env"),
//									FromPath: pulumi.String("/spec/containers/0/env"),
//								},
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
// # RestorePlan can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:gkebackup/restorePlan:RestorePlan default projects/{{project}}/locations/{{location}}/restorePlans/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkebackup/restorePlan:RestorePlan default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkebackup/restorePlan:RestorePlan default {{location}}/{{name}}
//
// ```
type RestorePlan struct {
	pulumi.CustomResourceState

	// A reference to the BackupPlan from which Backups may be used
	// as the source for Restores created via this RestorePlan.
	BackupPlan pulumi.StringOutput `pulumi:"backupPlan"`
	// The source cluster from which Restores will be created via this RestorePlan.
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// User specified descriptive string for this RestorePlan.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The region of the Restore Plan.
	Location pulumi.StringOutput `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Defines the configuration of Restores created via this RestorePlan.
	// Structure is documented below.
	RestoreConfig RestorePlanRestoreConfigOutput `pulumi:"restoreConfig"`
	// The State of the RestorePlan.
	State pulumi.StringOutput `pulumi:"state"`
	// Detailed description of why RestorePlan is in its current state.
	StateReason pulumi.StringOutput `pulumi:"stateReason"`
	// Server generated, unique identifier of UUID format.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewRestorePlan registers a new resource with the given unique name, arguments, and options.
func NewRestorePlan(ctx *pulumi.Context,
	name string, args *RestorePlanArgs, opts ...pulumi.ResourceOption) (*RestorePlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupPlan == nil {
		return nil, errors.New("invalid value for required argument 'BackupPlan'")
	}
	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.RestoreConfig == nil {
		return nil, errors.New("invalid value for required argument 'RestoreConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RestorePlan
	err := ctx.RegisterResource("gcp:gkebackup/restorePlan:RestorePlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestorePlan gets an existing RestorePlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestorePlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestorePlanState, opts ...pulumi.ResourceOption) (*RestorePlan, error) {
	var resource RestorePlan
	err := ctx.ReadResource("gcp:gkebackup/restorePlan:RestorePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestorePlan resources.
type restorePlanState struct {
	// A reference to the BackupPlan from which Backups may be used
	// as the source for Restores created via this RestorePlan.
	BackupPlan *string `pulumi:"backupPlan"`
	// The source cluster from which Restores will be created via this RestorePlan.
	Cluster *string `pulumi:"cluster"`
	// User specified descriptive string for this RestorePlan.
	Description *string `pulumi:"description"`
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The region of the Restore Plan.
	Location *string `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Defines the configuration of Restores created via this RestorePlan.
	// Structure is documented below.
	RestoreConfig *RestorePlanRestoreConfig `pulumi:"restoreConfig"`
	// The State of the RestorePlan.
	State *string `pulumi:"state"`
	// Detailed description of why RestorePlan is in its current state.
	StateReason *string `pulumi:"stateReason"`
	// Server generated, unique identifier of UUID format.
	Uid *string `pulumi:"uid"`
}

type RestorePlanState struct {
	// A reference to the BackupPlan from which Backups may be used
	// as the source for Restores created via this RestorePlan.
	BackupPlan pulumi.StringPtrInput
	// The source cluster from which Restores will be created via this RestorePlan.
	Cluster pulumi.StringPtrInput
	// User specified descriptive string for this RestorePlan.
	Description pulumi.StringPtrInput
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The region of the Restore Plan.
	Location pulumi.StringPtrInput
	// The full name of the BackupPlan Resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Defines the configuration of Restores created via this RestorePlan.
	// Structure is documented below.
	RestoreConfig RestorePlanRestoreConfigPtrInput
	// The State of the RestorePlan.
	State pulumi.StringPtrInput
	// Detailed description of why RestorePlan is in its current state.
	StateReason pulumi.StringPtrInput
	// Server generated, unique identifier of UUID format.
	Uid pulumi.StringPtrInput
}

func (RestorePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePlanState)(nil)).Elem()
}

type restorePlanArgs struct {
	// A reference to the BackupPlan from which Backups may be used
	// as the source for Restores created via this RestorePlan.
	BackupPlan string `pulumi:"backupPlan"`
	// The source cluster from which Restores will be created via this RestorePlan.
	Cluster string `pulumi:"cluster"`
	// User specified descriptive string for this RestorePlan.
	Description *string `pulumi:"description"`
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The region of the Restore Plan.
	Location string `pulumi:"location"`
	// The full name of the BackupPlan Resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Defines the configuration of Restores created via this RestorePlan.
	// Structure is documented below.
	RestoreConfig RestorePlanRestoreConfig `pulumi:"restoreConfig"`
}

// The set of arguments for constructing a RestorePlan resource.
type RestorePlanArgs struct {
	// A reference to the BackupPlan from which Backups may be used
	// as the source for Restores created via this RestorePlan.
	BackupPlan pulumi.StringInput
	// The source cluster from which Restores will be created via this RestorePlan.
	Cluster pulumi.StringInput
	// User specified descriptive string for this RestorePlan.
	Description pulumi.StringPtrInput
	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The region of the Restore Plan.
	Location pulumi.StringInput
	// The full name of the BackupPlan Resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Defines the configuration of Restores created via this RestorePlan.
	// Structure is documented below.
	RestoreConfig RestorePlanRestoreConfigInput
}

func (RestorePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePlanArgs)(nil)).Elem()
}

type RestorePlanInput interface {
	pulumi.Input

	ToRestorePlanOutput() RestorePlanOutput
	ToRestorePlanOutputWithContext(ctx context.Context) RestorePlanOutput
}

func (*RestorePlan) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePlan)(nil)).Elem()
}

func (i *RestorePlan) ToRestorePlanOutput() RestorePlanOutput {
	return i.ToRestorePlanOutputWithContext(context.Background())
}

func (i *RestorePlan) ToRestorePlanOutputWithContext(ctx context.Context) RestorePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePlanOutput)
}

func (i *RestorePlan) ToOutput(ctx context.Context) pulumix.Output[*RestorePlan] {
	return pulumix.Output[*RestorePlan]{
		OutputState: i.ToRestorePlanOutputWithContext(ctx).OutputState,
	}
}

// RestorePlanArrayInput is an input type that accepts RestorePlanArray and RestorePlanArrayOutput values.
// You can construct a concrete instance of `RestorePlanArrayInput` via:
//
//	RestorePlanArray{ RestorePlanArgs{...} }
type RestorePlanArrayInput interface {
	pulumi.Input

	ToRestorePlanArrayOutput() RestorePlanArrayOutput
	ToRestorePlanArrayOutputWithContext(context.Context) RestorePlanArrayOutput
}

type RestorePlanArray []RestorePlanInput

func (RestorePlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RestorePlan)(nil)).Elem()
}

func (i RestorePlanArray) ToRestorePlanArrayOutput() RestorePlanArrayOutput {
	return i.ToRestorePlanArrayOutputWithContext(context.Background())
}

func (i RestorePlanArray) ToRestorePlanArrayOutputWithContext(ctx context.Context) RestorePlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePlanArrayOutput)
}

func (i RestorePlanArray) ToOutput(ctx context.Context) pulumix.Output[[]*RestorePlan] {
	return pulumix.Output[[]*RestorePlan]{
		OutputState: i.ToRestorePlanArrayOutputWithContext(ctx).OutputState,
	}
}

// RestorePlanMapInput is an input type that accepts RestorePlanMap and RestorePlanMapOutput values.
// You can construct a concrete instance of `RestorePlanMapInput` via:
//
//	RestorePlanMap{ "key": RestorePlanArgs{...} }
type RestorePlanMapInput interface {
	pulumi.Input

	ToRestorePlanMapOutput() RestorePlanMapOutput
	ToRestorePlanMapOutputWithContext(context.Context) RestorePlanMapOutput
}

type RestorePlanMap map[string]RestorePlanInput

func (RestorePlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RestorePlan)(nil)).Elem()
}

func (i RestorePlanMap) ToRestorePlanMapOutput() RestorePlanMapOutput {
	return i.ToRestorePlanMapOutputWithContext(context.Background())
}

func (i RestorePlanMap) ToRestorePlanMapOutputWithContext(ctx context.Context) RestorePlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePlanMapOutput)
}

func (i RestorePlanMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RestorePlan] {
	return pulumix.Output[map[string]*RestorePlan]{
		OutputState: i.ToRestorePlanMapOutputWithContext(ctx).OutputState,
	}
}

type RestorePlanOutput struct{ *pulumi.OutputState }

func (RestorePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePlan)(nil)).Elem()
}

func (o RestorePlanOutput) ToRestorePlanOutput() RestorePlanOutput {
	return o
}

func (o RestorePlanOutput) ToRestorePlanOutputWithContext(ctx context.Context) RestorePlanOutput {
	return o
}

func (o RestorePlanOutput) ToOutput(ctx context.Context) pulumix.Output[*RestorePlan] {
	return pulumix.Output[*RestorePlan]{
		OutputState: o.OutputState,
	}
}

// A reference to the BackupPlan from which Backups may be used
// as the source for Restores created via this RestorePlan.
func (o RestorePlanOutput) BackupPlan() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringOutput { return v.BackupPlan }).(pulumi.StringOutput)
}

// The source cluster from which Restores will be created via this RestorePlan.
func (o RestorePlanOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

// User specified descriptive string for this RestorePlan.
func (o RestorePlanOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Description: A set of custom labels supplied by the user.
// A list of key->value pairs.
// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
func (o RestorePlanOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The region of the Restore Plan.
func (o RestorePlanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The full name of the BackupPlan Resource.
func (o RestorePlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RestorePlanOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Defines the configuration of Restores created via this RestorePlan.
// Structure is documented below.
func (o RestorePlanOutput) RestoreConfig() RestorePlanRestoreConfigOutput {
	return o.ApplyT(func(v *RestorePlan) RestorePlanRestoreConfigOutput { return v.RestoreConfig }).(RestorePlanRestoreConfigOutput)
}

// The State of the RestorePlan.
func (o RestorePlanOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Detailed description of why RestorePlan is in its current state.
func (o RestorePlanOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringOutput { return v.StateReason }).(pulumi.StringOutput)
}

// Server generated, unique identifier of UUID format.
func (o RestorePlanOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePlan) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

type RestorePlanArrayOutput struct{ *pulumi.OutputState }

func (RestorePlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RestorePlan)(nil)).Elem()
}

func (o RestorePlanArrayOutput) ToRestorePlanArrayOutput() RestorePlanArrayOutput {
	return o
}

func (o RestorePlanArrayOutput) ToRestorePlanArrayOutputWithContext(ctx context.Context) RestorePlanArrayOutput {
	return o
}

func (o RestorePlanArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RestorePlan] {
	return pulumix.Output[[]*RestorePlan]{
		OutputState: o.OutputState,
	}
}

func (o RestorePlanArrayOutput) Index(i pulumi.IntInput) RestorePlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RestorePlan {
		return vs[0].([]*RestorePlan)[vs[1].(int)]
	}).(RestorePlanOutput)
}

type RestorePlanMapOutput struct{ *pulumi.OutputState }

func (RestorePlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RestorePlan)(nil)).Elem()
}

func (o RestorePlanMapOutput) ToRestorePlanMapOutput() RestorePlanMapOutput {
	return o
}

func (o RestorePlanMapOutput) ToRestorePlanMapOutputWithContext(ctx context.Context) RestorePlanMapOutput {
	return o
}

func (o RestorePlanMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RestorePlan] {
	return pulumix.Output[map[string]*RestorePlan]{
		OutputState: o.OutputState,
	}
}

func (o RestorePlanMapOutput) MapIndex(k pulumi.StringInput) RestorePlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RestorePlan {
		return vs[0].(map[string]*RestorePlan)[vs[1].(string)]
	}).(RestorePlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RestorePlanInput)(nil)).Elem(), &RestorePlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestorePlanArrayInput)(nil)).Elem(), RestorePlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestorePlanMapInput)(nil)).Elem(), RestorePlanMap{})
	pulumi.RegisterOutputType(RestorePlanOutput{})
	pulumi.RegisterOutputType(RestorePlanArrayOutput{})
	pulumi.RegisterOutputType(RestorePlanMapOutput{})
}
