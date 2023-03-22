// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrunv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Cloud Run Job resource that references a container image which is run to completion.
//
// To get more information about Job, see:
//
// * [API documentation](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.jobs)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/run/docs/)
//
// ## Example Usage
// ### Cloudrunv2 Job Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrunv2.NewJob(ctx, "default", &cloudrunv2.JobArgs{
//				LaunchStage: pulumi.String("BETA"),
//				Location:    pulumi.String("us-central1"),
//				Template: &cloudrunv2.JobTemplateArgs{
//					Template: &cloudrunv2.JobTemplateTemplateArgs{
//						Containers: cloudrunv2.JobTemplateTemplateContainerArray{
//							&cloudrunv2.JobTemplateTemplateContainerArgs{
//								Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
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
// ### Cloudrunv2 Job Sql
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/secretmanager"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			secret, err := secretmanager.NewSecret(ctx, "secret", &secretmanager.SecretArgs{
//				SecretId: pulumi.String("secret"),
//				Replication: &secretmanager.SecretReplicationArgs{
//					Automatic: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := sql.NewDatabaseInstance(ctx, "instance", &sql.DatabaseInstanceArgs{
//				Region:          pulumi.String("us-central1"),
//				DatabaseVersion: pulumi.String("MYSQL_5_7"),
//				Settings: &sql.DatabaseInstanceSettingsArgs{
//					Tier: pulumi.String("db-f1-micro"),
//				},
//				DeletionProtection: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudrunv2.NewJob(ctx, "default", &cloudrunv2.JobArgs{
//				Location:    pulumi.String("us-central1"),
//				LaunchStage: pulumi.String("BETA"),
//				Template: &cloudrunv2.JobTemplateArgs{
//					Template: &cloudrunv2.JobTemplateTemplateArgs{
//						Volumes: cloudrunv2.JobTemplateTemplateVolumeArray{
//							&cloudrunv2.JobTemplateTemplateVolumeArgs{
//								Name: pulumi.String("cloudsql"),
//								CloudSqlInstance: &cloudrunv2.JobTemplateTemplateVolumeCloudSqlInstanceArgs{
//									Instances: pulumi.StringArray{
//										instance.ConnectionName,
//									},
//								},
//							},
//						},
//						Containers: cloudrunv2.JobTemplateTemplateContainerArray{
//							&cloudrunv2.JobTemplateTemplateContainerArgs{
//								Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
//								Envs: cloudrunv2.JobTemplateTemplateContainerEnvArray{
//									&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
//										Name:  pulumi.String("FOO"),
//										Value: pulumi.String("bar"),
//									},
//									&cloudrunv2.JobTemplateTemplateContainerEnvArgs{
//										Name: pulumi.String("latestdclsecret"),
//										ValueSource: &cloudrunv2.JobTemplateTemplateContainerEnvValueSourceArgs{
//											SecretKeyRef: &cloudrunv2.JobTemplateTemplateContainerEnvValueSourceSecretKeyRefArgs{
//												Secret:  secret.SecretId,
//												Version: pulumi.String("1"),
//											},
//										},
//									},
//								},
//								VolumeMounts: cloudrunv2.JobTemplateTemplateContainerVolumeMountArray{
//									&cloudrunv2.JobTemplateTemplateContainerVolumeMountArgs{
//										Name:      pulumi.String("cloudsql"),
//										MountPath: pulumi.String("/cloudsql"),
//									},
//								},
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = secretmanager.NewSecretVersion(ctx, "secret-version-data", &secretmanager.SecretVersionArgs{
//				Secret:     secret.Name,
//				SecretData: pulumi.String("secret-data"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = secretmanager.NewSecretIamMember(ctx, "secret-access", &secretmanager.SecretIamMemberArgs{
//				SecretId: secret.ID(),
//				Role:     pulumi.String("roles/secretmanager.secretAccessor"),
//				Member:   pulumi.String(fmt.Sprintf("serviceAccount:%v-compute@developer.gserviceaccount.com", project.Number)),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				secret,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Cloudrunv2 Job Vpcaccess
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/vpcaccess"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			customTestNetwork, err := compute.NewNetwork(ctx, "customTestNetwork", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			customTestSubnetwork, err := compute.NewSubnetwork(ctx, "customTestSubnetwork", &compute.SubnetworkArgs{
//				IpCidrRange: pulumi.String("10.2.0.0/28"),
//				Region:      pulumi.String("us-central1"),
//				Network:     customTestNetwork.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			connector, err := vpcaccess.NewConnector(ctx, "connector", &vpcaccess.ConnectorArgs{
//				Subnet: &vpcaccess.ConnectorSubnetArgs{
//					Name: customTestSubnetwork.Name,
//				},
//				MachineType:  pulumi.String("e2-standard-4"),
//				MinInstances: pulumi.Int(2),
//				MaxInstances: pulumi.Int(3),
//				Region:       pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudrunv2.NewJob(ctx, "default", &cloudrunv2.JobArgs{
//				Location:    pulumi.String("us-central1"),
//				LaunchStage: pulumi.String("BETA"),
//				Template: &cloudrunv2.JobTemplateArgs{
//					Template: &cloudrunv2.JobTemplateTemplateArgs{
//						Containers: cloudrunv2.JobTemplateTemplateContainerArray{
//							&cloudrunv2.JobTemplateTemplateContainerArgs{
//								Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
//							},
//						},
//						VpcAccess: &cloudrunv2.JobTemplateTemplateVpcAccessArgs{
//							Connector: connector.ID(),
//							Egress:    pulumi.String("ALL_TRAFFIC"),
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
// ### Cloudrunv2 Job Secret
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrunv2"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/secretmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			secret, err := secretmanager.NewSecret(ctx, "secret", &secretmanager.SecretArgs{
//				SecretId: pulumi.String("secret"),
//				Replication: &secretmanager.SecretReplicationArgs{
//					Automatic: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cloudrunv2.NewJob(ctx, "default", &cloudrunv2.JobArgs{
//				Location:    pulumi.String("us-central1"),
//				LaunchStage: pulumi.String("BETA"),
//				Template: &cloudrunv2.JobTemplateArgs{
//					Template: &cloudrunv2.JobTemplateTemplateArgs{
//						Volumes: cloudrunv2.JobTemplateTemplateVolumeArray{
//							&cloudrunv2.JobTemplateTemplateVolumeArgs{
//								Name: pulumi.String("a-volume"),
//								Secret: &cloudrunv2.JobTemplateTemplateVolumeSecretArgs{
//									Secret:      secret.SecretId,
//									DefaultMode: pulumi.Int(292),
//									Items: cloudrunv2.JobTemplateTemplateVolumeSecretItemArray{
//										&cloudrunv2.JobTemplateTemplateVolumeSecretItemArgs{
//											Version: pulumi.String("1"),
//											Path:    pulumi.String("my-secret"),
//											Mode:    pulumi.Int(256),
//										},
//									},
//								},
//							},
//						},
//						Containers: cloudrunv2.JobTemplateTemplateContainerArray{
//							&cloudrunv2.JobTemplateTemplateContainerArgs{
//								Image: pulumi.String("us-docker.pkg.dev/cloudrun/container/hello"),
//								VolumeMounts: cloudrunv2.JobTemplateTemplateContainerVolumeMountArray{
//									&cloudrunv2.JobTemplateTemplateContainerVolumeMountArgs{
//										Name:      pulumi.String("a-volume"),
//										MountPath: pulumi.String("/secrets"),
//									},
//								},
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = secretmanager.NewSecretVersion(ctx, "secret-version-data", &secretmanager.SecretVersionArgs{
//				Secret:     secret.Name,
//				SecretData: pulumi.String("secret-data"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = secretmanager.NewSecretIamMember(ctx, "secret-access", &secretmanager.SecretIamMemberArgs{
//				SecretId: secret.ID(),
//				Role:     pulumi.String("roles/secretmanager.secretAccessor"),
//				Member:   pulumi.String(fmt.Sprintf("serviceAccount:%v-compute@developer.gserviceaccount.com", project.Number)),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				secret,
//			}))
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
// # Job can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/job:Job default projects/{{project}}/locations/{{location}}/jobs/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/job:Job default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:cloudrunv2/job:Job default {{location}}/{{name}}
//
// ```
type Job struct {
	pulumi.CustomResourceState

	// Settings for the Binary Authorization feature.
	// Structure is documented below.
	BinaryAuthorization JobBinaryAuthorizationPtrOutput `pulumi:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	Client pulumi.StringPtrOutput `pulumi:"client"`
	// Arbitrary version identifier for the API client.
	ClientVersion pulumi.StringPtrOutput `pulumi:"clientVersion"`
	// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in reconciling for additional information on `reconciliation` process in Cloud Run.
	// Structure is documented below.
	Conditions JobConditionArrayOutput `pulumi:"conditions"`
	// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Number of executions created for this job.
	ExecutionCount pulumi.IntOutput `pulumi:"executionCount"`
	// A number that monotonically increases every time the user modifies the desired state.
	Generation pulumi.StringOutput `pulumi:"generation"`
	// KRM-style labels for the resource.
	// (Optional)
	// KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the last created execution.
	// Structure is documented below.
	LatestCreatedExecutions JobLatestCreatedExecutionArrayOutput `pulumi:"latestCreatedExecutions"`
	// The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
	// Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage pulumi.StringOutput `pulumi:"launchStage"`
	// The location of the cloud run job
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the Job.
	Name pulumi.StringOutput `pulumi:"name"`
	// The generation of this Job. See comments in reconciling for additional information on reconciliation process in Cloud Run.
	ObservedGeneration pulumi.StringOutput `pulumi:"observedGeneration"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Returns true if the Job is currently being acted upon by the system to bring it into the desired state.
	// When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.
	// If reconciliation succeeded, the following fields will match: observedGeneration and generation, latestSucceededExecution and latestCreatedExecution.
	// If reconciliation failed, observedGeneration and latestSucceededExecution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// The template used to create executions for this Job.
	// Structure is documented below.
	Template JobTemplateOutput `pulumi:"template"`
	// The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state
	// Structure is documented below.
	TerminalConditions JobTerminalConditionArrayOutput `pulumi:"terminalConditions"`
	// Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	var resource Job
	err := ctx.RegisterResource("gcp:cloudrunv2/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("gcp:cloudrunv2/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// Settings for the Binary Authorization feature.
	// Structure is documented below.
	BinaryAuthorization *JobBinaryAuthorization `pulumi:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	Client *string `pulumi:"client"`
	// Arbitrary version identifier for the API client.
	ClientVersion *string `pulumi:"clientVersion"`
	// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in reconciling for additional information on `reconciliation` process in Cloud Run.
	// Structure is documented below.
	Conditions []JobCondition `pulumi:"conditions"`
	// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
	Etag *string `pulumi:"etag"`
	// Number of executions created for this job.
	ExecutionCount *int `pulumi:"executionCount"`
	// A number that monotonically increases every time the user modifies the desired state.
	Generation *string `pulumi:"generation"`
	// KRM-style labels for the resource.
	// (Optional)
	// KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
	Labels map[string]string `pulumi:"labels"`
	// Name of the last created execution.
	// Structure is documented below.
	LatestCreatedExecutions []JobLatestCreatedExecution `pulumi:"latestCreatedExecutions"`
	// The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
	// Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage *string `pulumi:"launchStage"`
	// The location of the cloud run job
	Location *string `pulumi:"location"`
	// Name of the Job.
	Name *string `pulumi:"name"`
	// The generation of this Job. See comments in reconciling for additional information on reconciliation process in Cloud Run.
	ObservedGeneration *string `pulumi:"observedGeneration"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Returns true if the Job is currently being acted upon by the system to bring it into the desired state.
	// When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.
	// If reconciliation succeeded, the following fields will match: observedGeneration and generation, latestSucceededExecution and latestCreatedExecution.
	// If reconciliation failed, observedGeneration and latestSucceededExecution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions
	Reconciling *bool `pulumi:"reconciling"`
	// The template used to create executions for this Job.
	// Structure is documented below.
	Template *JobTemplate `pulumi:"template"`
	// The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state
	// Structure is documented below.
	TerminalConditions []JobTerminalCondition `pulumi:"terminalConditions"`
	// Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid *string `pulumi:"uid"`
}

type JobState struct {
	// Settings for the Binary Authorization feature.
	// Structure is documented below.
	BinaryAuthorization JobBinaryAuthorizationPtrInput
	// Arbitrary identifier for the API client.
	Client pulumi.StringPtrInput
	// Arbitrary version identifier for the API client.
	ClientVersion pulumi.StringPtrInput
	// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in reconciling for additional information on `reconciliation` process in Cloud Run.
	// Structure is documented below.
	Conditions JobConditionArrayInput
	// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
	Etag pulumi.StringPtrInput
	// Number of executions created for this job.
	ExecutionCount pulumi.IntPtrInput
	// A number that monotonically increases every time the user modifies the desired state.
	Generation pulumi.StringPtrInput
	// KRM-style labels for the resource.
	// (Optional)
	// KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
	Labels pulumi.StringMapInput
	// Name of the last created execution.
	// Structure is documented below.
	LatestCreatedExecutions JobLatestCreatedExecutionArrayInput
	// The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
	// Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage pulumi.StringPtrInput
	// The location of the cloud run job
	Location pulumi.StringPtrInput
	// Name of the Job.
	Name pulumi.StringPtrInput
	// The generation of this Job. See comments in reconciling for additional information on reconciliation process in Cloud Run.
	ObservedGeneration pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Returns true if the Job is currently being acted upon by the system to bring it into the desired state.
	// When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.
	// If reconciliation succeeded, the following fields will match: observedGeneration and generation, latestSucceededExecution and latestCreatedExecution.
	// If reconciliation failed, observedGeneration and latestSucceededExecution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions
	Reconciling pulumi.BoolPtrInput
	// The template used to create executions for this Job.
	// Structure is documented below.
	Template JobTemplatePtrInput
	// The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state
	// Structure is documented below.
	TerminalConditions JobTerminalConditionArrayInput
	// Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// Settings for the Binary Authorization feature.
	// Structure is documented below.
	BinaryAuthorization *JobBinaryAuthorization `pulumi:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	Client *string `pulumi:"client"`
	// Arbitrary version identifier for the API client.
	ClientVersion *string `pulumi:"clientVersion"`
	// KRM-style labels for the resource.
	// (Optional)
	// KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
	Labels map[string]string `pulumi:"labels"`
	// The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
	// Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage *string `pulumi:"launchStage"`
	// The location of the cloud run job
	Location *string `pulumi:"location"`
	// Name of the Job.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The template used to create executions for this Job.
	// Structure is documented below.
	Template JobTemplate `pulumi:"template"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// Settings for the Binary Authorization feature.
	// Structure is documented below.
	BinaryAuthorization JobBinaryAuthorizationPtrInput
	// Arbitrary identifier for the API client.
	Client pulumi.StringPtrInput
	// Arbitrary version identifier for the API client.
	ClientVersion pulumi.StringPtrInput
	// KRM-style labels for the resource.
	// (Optional)
	// KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
	Labels pulumi.StringMapInput
	// The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
	// Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
	LaunchStage pulumi.StringPtrInput
	// The location of the cloud run job
	Location pulumi.StringPtrInput
	// Name of the Job.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The template used to create executions for this Job.
	// Structure is documented below.
	Template JobTemplateInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

// JobArrayInput is an input type that accepts JobArray and JobArrayOutput values.
// You can construct a concrete instance of `JobArrayInput` via:
//
//	JobArray{ JobArgs{...} }
type JobArrayInput interface {
	pulumi.Input

	ToJobArrayOutput() JobArrayOutput
	ToJobArrayOutputWithContext(context.Context) JobArrayOutput
}

type JobArray []JobInput

func (JobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (i JobArray) ToJobArrayOutput() JobArrayOutput {
	return i.ToJobArrayOutputWithContext(context.Background())
}

func (i JobArray) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobArrayOutput)
}

// JobMapInput is an input type that accepts JobMap and JobMapOutput values.
// You can construct a concrete instance of `JobMapInput` via:
//
//	JobMap{ "key": JobArgs{...} }
type JobMapInput interface {
	pulumi.Input

	ToJobMapOutput() JobMapOutput
	ToJobMapOutputWithContext(context.Context) JobMapOutput
}

type JobMap map[string]JobInput

func (JobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (i JobMap) ToJobMapOutput() JobMapOutput {
	return i.ToJobMapOutputWithContext(context.Background())
}

func (i JobMap) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMapOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

// Settings for the Binary Authorization feature.
// Structure is documented below.
func (o JobOutput) BinaryAuthorization() JobBinaryAuthorizationPtrOutput {
	return o.ApplyT(func(v *Job) JobBinaryAuthorizationPtrOutput { return v.BinaryAuthorization }).(JobBinaryAuthorizationPtrOutput)
}

// Arbitrary identifier for the API client.
func (o JobOutput) Client() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Client }).(pulumi.StringPtrOutput)
}

// Arbitrary version identifier for the API client.
func (o JobOutput) ClientVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.ClientVersion }).(pulumi.StringPtrOutput)
}

// The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in reconciling for additional information on `reconciliation` process in Cloud Run.
// Structure is documented below.
func (o JobOutput) Conditions() JobConditionArrayOutput {
	return o.ApplyT(func(v *Job) JobConditionArrayOutput { return v.Conditions }).(JobConditionArrayOutput)
}

// A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
func (o JobOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Number of executions created for this job.
func (o JobOutput) ExecutionCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Job) pulumi.IntOutput { return v.ExecutionCount }).(pulumi.IntOutput)
}

// A number that monotonically increases every time the user modifies the desired state.
func (o JobOutput) Generation() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Generation }).(pulumi.StringOutput)
}

// KRM-style labels for the resource.
// (Optional)
// KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
func (o JobOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the last created execution.
// Structure is documented below.
func (o JobOutput) LatestCreatedExecutions() JobLatestCreatedExecutionArrayOutput {
	return o.ApplyT(func(v *Job) JobLatestCreatedExecutionArrayOutput { return v.LatestCreatedExecutions }).(JobLatestCreatedExecutionArrayOutput)
}

// The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
// Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
func (o JobOutput) LaunchStage() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.LaunchStage }).(pulumi.StringOutput)
}

// The location of the cloud run job
func (o JobOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the Job.
func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The generation of this Job. See comments in reconciling for additional information on reconciliation process in Cloud Run.
func (o JobOutput) ObservedGeneration() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.ObservedGeneration }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o JobOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Returns true if the Job is currently being acted upon by the system to bring it into the desired state.
// When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.
// If reconciliation succeeded, the following fields will match: observedGeneration and generation, latestSucceededExecution and latestCreatedExecution.
// If reconciliation failed, observedGeneration and latestSucceededExecution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions
func (o JobOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// The template used to create executions for this Job.
// Structure is documented below.
func (o JobOutput) Template() JobTemplateOutput {
	return o.ApplyT(func(v *Job) JobTemplateOutput { return v.Template }).(JobTemplateOutput)
}

// The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state
// Structure is documented below.
func (o JobOutput) TerminalConditions() JobTerminalConditionArrayOutput {
	return o.ApplyT(func(v *Job) JobTerminalConditionArrayOutput { return v.TerminalConditions }).(JobTerminalConditionArrayOutput)
}

// Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
func (o JobOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

type JobArrayOutput struct{ *pulumi.OutputState }

func (JobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (o JobArrayOutput) ToJobArrayOutput() JobArrayOutput {
	return o
}

func (o JobArrayOutput) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return o
}

func (o JobArrayOutput) Index(i pulumi.IntInput) JobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Job {
		return vs[0].([]*Job)[vs[1].(int)]
	}).(JobOutput)
}

type JobMapOutput struct{ *pulumi.OutputState }

func (JobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (o JobMapOutput) ToJobMapOutput() JobMapOutput {
	return o
}

func (o JobMapOutput) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return o
}

func (o JobMapOutput) MapIndex(k pulumi.StringInput) JobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Job {
		return vs[0].(map[string]*Job)[vs[1].(string)]
	}).(JobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobArrayInput)(nil)).Elem(), JobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobMapInput)(nil)).Elem(), JobMap{})
	pulumi.RegisterOutputType(JobOutput{})
	pulumi.RegisterOutputType(JobArrayOutput{})
	pulumi.RegisterOutputType(JobMapOutput{})
}
