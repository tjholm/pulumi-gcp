// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ### Dataform Repository Workflow Config
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataform"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/secretmanager"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sourcerepo"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			gitRepository, err := sourcerepo.NewRepository(ctx, "gitRepository", nil, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			secret, err := secretmanager.NewSecret(ctx, "secret", &secretmanager.SecretArgs{
//				SecretId: pulumi.String("my_secret"),
//				Replication: &secretmanager.SecretReplicationArgs{
//					Auto: nil,
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			secretVersion, err := secretmanager.NewSecretVersion(ctx, "secretVersion", &secretmanager.SecretVersionArgs{
//				Secret:     secret.ID(),
//				SecretData: pulumi.String("secret-data"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			repository, err := dataform.NewRepository(ctx, "repository", &dataform.RepositoryArgs{
//				Region: pulumi.String("us-central1"),
//				GitRemoteSettings: &dataform.RepositoryGitRemoteSettingsArgs{
//					Url:                              gitRepository.Url,
//					DefaultBranch:                    pulumi.String("main"),
//					AuthenticationTokenSecretVersion: secretVersion.ID(),
//				},
//				WorkspaceCompilationOverrides: &dataform.RepositoryWorkspaceCompilationOverridesArgs{
//					DefaultDatabase: pulumi.String("database"),
//					SchemaSuffix:    pulumi.String("_suffix"),
//					TablePrefix:     pulumi.String("prefix_"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			releaseConfig, err := dataform.NewRepositoryReleaseConfig(ctx, "releaseConfig", &dataform.RepositoryReleaseConfigArgs{
//				Project:      repository.Project,
//				Region:       repository.Region,
//				Repository:   repository.Name,
//				GitCommitish: pulumi.String("main"),
//				CronSchedule: pulumi.String("0 7 * * *"),
//				TimeZone:     pulumi.String("America/New_York"),
//				CodeCompilationConfig: &dataform.RepositoryReleaseConfigCodeCompilationConfigArgs{
//					DefaultDatabase: pulumi.String("gcp-example-project"),
//					DefaultSchema:   pulumi.String("example-dataset"),
//					DefaultLocation: pulumi.String("us-central1"),
//					AssertionSchema: pulumi.String("example-assertion-dataset"),
//					DatabaseSuffix:  pulumi.String(""),
//					SchemaSuffix:    pulumi.String(""),
//					TablePrefix:     pulumi.String(""),
//					Vars: pulumi.StringMap{
//						"var1": pulumi.String("value"),
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			dataformSa, err := serviceAccount.NewAccount(ctx, "dataformSa", &serviceAccount.AccountArgs{
//				AccountId:   pulumi.String("dataform-workflow-sa"),
//				DisplayName: pulumi.String("Dataform Service Account"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = dataform.NewRepositoryWorkflowConfig(ctx, "workflow", &dataform.RepositoryWorkflowConfigArgs{
//				Project:       repository.Project,
//				Region:        repository.Region,
//				Repository:    repository.Name,
//				ReleaseConfig: releaseConfig.ID(),
//				InvocationConfig: &dataform.RepositoryWorkflowConfigInvocationConfigArgs{
//					IncludedTargets: dataform.RepositoryWorkflowConfigInvocationConfigIncludedTargetArray{
//						&dataform.RepositoryWorkflowConfigInvocationConfigIncludedTargetArgs{
//							Database: pulumi.String("gcp-example-project"),
//							Schema:   pulumi.String("example-dataset"),
//							Name:     pulumi.String("target_1"),
//						},
//						&dataform.RepositoryWorkflowConfigInvocationConfigIncludedTargetArgs{
//							Database: pulumi.String("gcp-example-project"),
//							Schema:   pulumi.String("example-dataset"),
//							Name:     pulumi.String("target_2"),
//						},
//					},
//					IncludedTags: pulumi.StringArray{
//						pulumi.String("tag_1"),
//					},
//					TransitiveDependenciesIncluded:       pulumi.Bool(true),
//					TransitiveDependentsIncluded:         pulumi.Bool(true),
//					FullyRefreshIncrementalTablesEnabled: pulumi.Bool(false),
//					ServiceAccount:                       dataformSa.Email,
//				},
//				CronSchedule: pulumi.String("0 7 * * *"),
//				TimeZone:     pulumi.String("America/New_York"),
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
// # RepositoryWorkflowConfig can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig default projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig default {{project}}/{{region}}/{{repository}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig default {{region}}/{{repository}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig default {{repository}}/{{name}}
//
// ```
type RepositoryWorkflowConfig struct {
	pulumi.CustomResourceState

	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule pulumi.StringPtrOutput `pulumi:"cronSchedule"`
	// Optional. If left unset, a default InvocationConfig will be used.
	// Structure is documented below.
	InvocationConfig RepositoryWorkflowConfigInvocationConfigPtrOutput `pulumi:"invocationConfig"`
	// The workflow's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
	// Structure is documented below.
	RecentScheduledExecutionRecords RepositoryWorkflowConfigRecentScheduledExecutionRecordArrayOutput `pulumi:"recentScheduledExecutionRecords"`
	// A reference to the region
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*/locations/*/repositories/*/releaseConfigs/*.
	//
	// ***
	ReleaseConfig pulumi.StringOutput `pulumi:"releaseConfig"`
	// A reference to the Dataform repository
	Repository pulumi.StringPtrOutput `pulumi:"repository"`
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone pulumi.StringPtrOutput `pulumi:"timeZone"`
}

// NewRepositoryWorkflowConfig registers a new resource with the given unique name, arguments, and options.
func NewRepositoryWorkflowConfig(ctx *pulumi.Context,
	name string, args *RepositoryWorkflowConfigArgs, opts ...pulumi.ResourceOption) (*RepositoryWorkflowConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReleaseConfig == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryWorkflowConfig
	err := ctx.RegisterResource("gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryWorkflowConfig gets an existing RepositoryWorkflowConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryWorkflowConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryWorkflowConfigState, opts ...pulumi.ResourceOption) (*RepositoryWorkflowConfig, error) {
	var resource RepositoryWorkflowConfig
	err := ctx.ReadResource("gcp:dataform/repositoryWorkflowConfig:RepositoryWorkflowConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryWorkflowConfig resources.
type repositoryWorkflowConfigState struct {
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule *string `pulumi:"cronSchedule"`
	// Optional. If left unset, a default InvocationConfig will be used.
	// Structure is documented below.
	InvocationConfig *RepositoryWorkflowConfigInvocationConfig `pulumi:"invocationConfig"`
	// The workflow's name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
	// Structure is documented below.
	RecentScheduledExecutionRecords []RepositoryWorkflowConfigRecentScheduledExecutionRecord `pulumi:"recentScheduledExecutionRecords"`
	// A reference to the region
	Region *string `pulumi:"region"`
	// The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*/locations/*/repositories/*/releaseConfigs/*.
	//
	// ***
	ReleaseConfig *string `pulumi:"releaseConfig"`
	// A reference to the Dataform repository
	Repository *string `pulumi:"repository"`
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone *string `pulumi:"timeZone"`
}

type RepositoryWorkflowConfigState struct {
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule pulumi.StringPtrInput
	// Optional. If left unset, a default InvocationConfig will be used.
	// Structure is documented below.
	InvocationConfig RepositoryWorkflowConfigInvocationConfigPtrInput
	// The workflow's name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
	// Structure is documented below.
	RecentScheduledExecutionRecords RepositoryWorkflowConfigRecentScheduledExecutionRecordArrayInput
	// A reference to the region
	Region pulumi.StringPtrInput
	// The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*/locations/*/repositories/*/releaseConfigs/*.
	//
	// ***
	ReleaseConfig pulumi.StringPtrInput
	// A reference to the Dataform repository
	Repository pulumi.StringPtrInput
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone pulumi.StringPtrInput
}

func (RepositoryWorkflowConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryWorkflowConfigState)(nil)).Elem()
}

type repositoryWorkflowConfigArgs struct {
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule *string `pulumi:"cronSchedule"`
	// Optional. If left unset, a default InvocationConfig will be used.
	// Structure is documented below.
	InvocationConfig *RepositoryWorkflowConfigInvocationConfig `pulumi:"invocationConfig"`
	// The workflow's name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region
	Region *string `pulumi:"region"`
	// The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*/locations/*/repositories/*/releaseConfigs/*.
	//
	// ***
	ReleaseConfig string `pulumi:"releaseConfig"`
	// A reference to the Dataform repository
	Repository *string `pulumi:"repository"`
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a RepositoryWorkflowConfig resource.
type RepositoryWorkflowConfigArgs struct {
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule pulumi.StringPtrInput
	// Optional. If left unset, a default InvocationConfig will be used.
	// Structure is documented below.
	InvocationConfig RepositoryWorkflowConfigInvocationConfigPtrInput
	// The workflow's name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region
	Region pulumi.StringPtrInput
	// The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*/locations/*/repositories/*/releaseConfigs/*.
	//
	// ***
	ReleaseConfig pulumi.StringInput
	// A reference to the Dataform repository
	Repository pulumi.StringPtrInput
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone pulumi.StringPtrInput
}

func (RepositoryWorkflowConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryWorkflowConfigArgs)(nil)).Elem()
}

type RepositoryWorkflowConfigInput interface {
	pulumi.Input

	ToRepositoryWorkflowConfigOutput() RepositoryWorkflowConfigOutput
	ToRepositoryWorkflowConfigOutputWithContext(ctx context.Context) RepositoryWorkflowConfigOutput
}

func (*RepositoryWorkflowConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryWorkflowConfig)(nil)).Elem()
}

func (i *RepositoryWorkflowConfig) ToRepositoryWorkflowConfigOutput() RepositoryWorkflowConfigOutput {
	return i.ToRepositoryWorkflowConfigOutputWithContext(context.Background())
}

func (i *RepositoryWorkflowConfig) ToRepositoryWorkflowConfigOutputWithContext(ctx context.Context) RepositoryWorkflowConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWorkflowConfigOutput)
}

func (i *RepositoryWorkflowConfig) ToOutput(ctx context.Context) pulumix.Output[*RepositoryWorkflowConfig] {
	return pulumix.Output[*RepositoryWorkflowConfig]{
		OutputState: i.ToRepositoryWorkflowConfigOutputWithContext(ctx).OutputState,
	}
}

// RepositoryWorkflowConfigArrayInput is an input type that accepts RepositoryWorkflowConfigArray and RepositoryWorkflowConfigArrayOutput values.
// You can construct a concrete instance of `RepositoryWorkflowConfigArrayInput` via:
//
//	RepositoryWorkflowConfigArray{ RepositoryWorkflowConfigArgs{...} }
type RepositoryWorkflowConfigArrayInput interface {
	pulumi.Input

	ToRepositoryWorkflowConfigArrayOutput() RepositoryWorkflowConfigArrayOutput
	ToRepositoryWorkflowConfigArrayOutputWithContext(context.Context) RepositoryWorkflowConfigArrayOutput
}

type RepositoryWorkflowConfigArray []RepositoryWorkflowConfigInput

func (RepositoryWorkflowConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryWorkflowConfig)(nil)).Elem()
}

func (i RepositoryWorkflowConfigArray) ToRepositoryWorkflowConfigArrayOutput() RepositoryWorkflowConfigArrayOutput {
	return i.ToRepositoryWorkflowConfigArrayOutputWithContext(context.Background())
}

func (i RepositoryWorkflowConfigArray) ToRepositoryWorkflowConfigArrayOutputWithContext(ctx context.Context) RepositoryWorkflowConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWorkflowConfigArrayOutput)
}

func (i RepositoryWorkflowConfigArray) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryWorkflowConfig] {
	return pulumix.Output[[]*RepositoryWorkflowConfig]{
		OutputState: i.ToRepositoryWorkflowConfigArrayOutputWithContext(ctx).OutputState,
	}
}

// RepositoryWorkflowConfigMapInput is an input type that accepts RepositoryWorkflowConfigMap and RepositoryWorkflowConfigMapOutput values.
// You can construct a concrete instance of `RepositoryWorkflowConfigMapInput` via:
//
//	RepositoryWorkflowConfigMap{ "key": RepositoryWorkflowConfigArgs{...} }
type RepositoryWorkflowConfigMapInput interface {
	pulumi.Input

	ToRepositoryWorkflowConfigMapOutput() RepositoryWorkflowConfigMapOutput
	ToRepositoryWorkflowConfigMapOutputWithContext(context.Context) RepositoryWorkflowConfigMapOutput
}

type RepositoryWorkflowConfigMap map[string]RepositoryWorkflowConfigInput

func (RepositoryWorkflowConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryWorkflowConfig)(nil)).Elem()
}

func (i RepositoryWorkflowConfigMap) ToRepositoryWorkflowConfigMapOutput() RepositoryWorkflowConfigMapOutput {
	return i.ToRepositoryWorkflowConfigMapOutputWithContext(context.Background())
}

func (i RepositoryWorkflowConfigMap) ToRepositoryWorkflowConfigMapOutputWithContext(ctx context.Context) RepositoryWorkflowConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryWorkflowConfigMapOutput)
}

func (i RepositoryWorkflowConfigMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryWorkflowConfig] {
	return pulumix.Output[map[string]*RepositoryWorkflowConfig]{
		OutputState: i.ToRepositoryWorkflowConfigMapOutputWithContext(ctx).OutputState,
	}
}

type RepositoryWorkflowConfigOutput struct{ *pulumi.OutputState }

func (RepositoryWorkflowConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryWorkflowConfig)(nil)).Elem()
}

func (o RepositoryWorkflowConfigOutput) ToRepositoryWorkflowConfigOutput() RepositoryWorkflowConfigOutput {
	return o
}

func (o RepositoryWorkflowConfigOutput) ToRepositoryWorkflowConfigOutputWithContext(ctx context.Context) RepositoryWorkflowConfigOutput {
	return o
}

func (o RepositoryWorkflowConfigOutput) ToOutput(ctx context.Context) pulumix.Output[*RepositoryWorkflowConfig] {
	return pulumix.Output[*RepositoryWorkflowConfig]{
		OutputState: o.OutputState,
	}
}

// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
func (o RepositoryWorkflowConfigOutput) CronSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) pulumi.StringPtrOutput { return v.CronSchedule }).(pulumi.StringPtrOutput)
}

// Optional. If left unset, a default InvocationConfig will be used.
// Structure is documented below.
func (o RepositoryWorkflowConfigOutput) InvocationConfig() RepositoryWorkflowConfigInvocationConfigPtrOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) RepositoryWorkflowConfigInvocationConfigPtrOutput {
		return v.InvocationConfig
	}).(RepositoryWorkflowConfigInvocationConfigPtrOutput)
}

// The workflow's name.
func (o RepositoryWorkflowConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RepositoryWorkflowConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
// Structure is documented below.
func (o RepositoryWorkflowConfigOutput) RecentScheduledExecutionRecords() RepositoryWorkflowConfigRecentScheduledExecutionRecordArrayOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) RepositoryWorkflowConfigRecentScheduledExecutionRecordArrayOutput {
		return v.RecentScheduledExecutionRecords
	}).(RepositoryWorkflowConfigRecentScheduledExecutionRecordArrayOutput)
}

// A reference to the region
func (o RepositoryWorkflowConfigOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*/locations/*/repositories/*/releaseConfigs/*.
//
// ***
func (o RepositoryWorkflowConfigOutput) ReleaseConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) pulumi.StringOutput { return v.ReleaseConfig }).(pulumi.StringOutput)
}

// A reference to the Dataform repository
func (o RepositoryWorkflowConfigOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) pulumi.StringPtrOutput { return v.Repository }).(pulumi.StringPtrOutput)
}

// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
func (o RepositoryWorkflowConfigOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryWorkflowConfig) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type RepositoryWorkflowConfigArrayOutput struct{ *pulumi.OutputState }

func (RepositoryWorkflowConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryWorkflowConfig)(nil)).Elem()
}

func (o RepositoryWorkflowConfigArrayOutput) ToRepositoryWorkflowConfigArrayOutput() RepositoryWorkflowConfigArrayOutput {
	return o
}

func (o RepositoryWorkflowConfigArrayOutput) ToRepositoryWorkflowConfigArrayOutputWithContext(ctx context.Context) RepositoryWorkflowConfigArrayOutput {
	return o
}

func (o RepositoryWorkflowConfigArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RepositoryWorkflowConfig] {
	return pulumix.Output[[]*RepositoryWorkflowConfig]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryWorkflowConfigArrayOutput) Index(i pulumi.IntInput) RepositoryWorkflowConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryWorkflowConfig {
		return vs[0].([]*RepositoryWorkflowConfig)[vs[1].(int)]
	}).(RepositoryWorkflowConfigOutput)
}

type RepositoryWorkflowConfigMapOutput struct{ *pulumi.OutputState }

func (RepositoryWorkflowConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryWorkflowConfig)(nil)).Elem()
}

func (o RepositoryWorkflowConfigMapOutput) ToRepositoryWorkflowConfigMapOutput() RepositoryWorkflowConfigMapOutput {
	return o
}

func (o RepositoryWorkflowConfigMapOutput) ToRepositoryWorkflowConfigMapOutputWithContext(ctx context.Context) RepositoryWorkflowConfigMapOutput {
	return o
}

func (o RepositoryWorkflowConfigMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RepositoryWorkflowConfig] {
	return pulumix.Output[map[string]*RepositoryWorkflowConfig]{
		OutputState: o.OutputState,
	}
}

func (o RepositoryWorkflowConfigMapOutput) MapIndex(k pulumi.StringInput) RepositoryWorkflowConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryWorkflowConfig {
		return vs[0].(map[string]*RepositoryWorkflowConfig)[vs[1].(string)]
	}).(RepositoryWorkflowConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWorkflowConfigInput)(nil)).Elem(), &RepositoryWorkflowConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWorkflowConfigArrayInput)(nil)).Elem(), RepositoryWorkflowConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryWorkflowConfigMapInput)(nil)).Elem(), RepositoryWorkflowConfigMap{})
	pulumi.RegisterOutputType(RepositoryWorkflowConfigOutput{})
	pulumi.RegisterOutputType(RepositoryWorkflowConfigArrayOutput{})
	pulumi.RegisterOutputType(RepositoryWorkflowConfigMapOutput{})
}
