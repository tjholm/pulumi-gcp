// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Dataform Repository Release Config
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/dataform"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/secretmanager"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/sourcerepo"
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
//			_, err = dataform.NewRepositoryReleaseConfig(ctx, "release", &dataform.RepositoryReleaseConfigArgs{
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
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// RepositoryReleaseConfig can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/repositories/{{repository}}/releaseConfigs/{{name}}` * `{{project}}/{{region}}/{{repository}}/{{name}}` * `{{region}}/{{repository}}/{{name}}` * `{{repository}}/{{name}}` When using the `pulumi import` command, RepositoryReleaseConfig can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig default projects/{{project}}/locations/{{region}}/repositories/{{repository}}/releaseConfigs/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig default {{project}}/{{region}}/{{repository}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig default {{region}}/{{repository}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig default {{repository}}/{{name}}
//
// ```
type RepositoryReleaseConfig struct {
	pulumi.CustomResourceState

	// Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
	// Structure is documented below.
	CodeCompilationConfig RepositoryReleaseConfigCodeCompilationConfigPtrOutput `pulumi:"codeCompilationConfig"`
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule pulumi.StringPtrOutput `pulumi:"cronSchedule"`
	// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
	//
	// ***
	GitCommitish pulumi.StringOutput `pulumi:"gitCommitish"`
	// The release's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Records of the 10 most recent scheduled release attempts, ordered in in descending order of releaseTime. Updated whenever automatic creation of a compilation result is triggered by cronSchedule.
	// Structure is documented below.
	RecentScheduledReleaseRecords RepositoryReleaseConfigRecentScheduledReleaseRecordArrayOutput `pulumi:"recentScheduledReleaseRecords"`
	// A reference to the region
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// A reference to the Dataform repository
	Repository pulumi.StringPtrOutput `pulumi:"repository"`
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone pulumi.StringPtrOutput `pulumi:"timeZone"`
}

// NewRepositoryReleaseConfig registers a new resource with the given unique name, arguments, and options.
func NewRepositoryReleaseConfig(ctx *pulumi.Context,
	name string, args *RepositoryReleaseConfigArgs, opts ...pulumi.ResourceOption) (*RepositoryReleaseConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GitCommitish == nil {
		return nil, errors.New("invalid value for required argument 'GitCommitish'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RepositoryReleaseConfig
	err := ctx.RegisterResource("gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryReleaseConfig gets an existing RepositoryReleaseConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryReleaseConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryReleaseConfigState, opts ...pulumi.ResourceOption) (*RepositoryReleaseConfig, error) {
	var resource RepositoryReleaseConfig
	err := ctx.ReadResource("gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryReleaseConfig resources.
type repositoryReleaseConfigState struct {
	// Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
	// Structure is documented below.
	CodeCompilationConfig *RepositoryReleaseConfigCodeCompilationConfig `pulumi:"codeCompilationConfig"`
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule *string `pulumi:"cronSchedule"`
	// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
	//
	// ***
	GitCommitish *string `pulumi:"gitCommitish"`
	// The release's name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Records of the 10 most recent scheduled release attempts, ordered in in descending order of releaseTime. Updated whenever automatic creation of a compilation result is triggered by cronSchedule.
	// Structure is documented below.
	RecentScheduledReleaseRecords []RepositoryReleaseConfigRecentScheduledReleaseRecord `pulumi:"recentScheduledReleaseRecords"`
	// A reference to the region
	Region *string `pulumi:"region"`
	// A reference to the Dataform repository
	Repository *string `pulumi:"repository"`
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone *string `pulumi:"timeZone"`
}

type RepositoryReleaseConfigState struct {
	// Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
	// Structure is documented below.
	CodeCompilationConfig RepositoryReleaseConfigCodeCompilationConfigPtrInput
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule pulumi.StringPtrInput
	// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
	//
	// ***
	GitCommitish pulumi.StringPtrInput
	// The release's name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Records of the 10 most recent scheduled release attempts, ordered in in descending order of releaseTime. Updated whenever automatic creation of a compilation result is triggered by cronSchedule.
	// Structure is documented below.
	RecentScheduledReleaseRecords RepositoryReleaseConfigRecentScheduledReleaseRecordArrayInput
	// A reference to the region
	Region pulumi.StringPtrInput
	// A reference to the Dataform repository
	Repository pulumi.StringPtrInput
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone pulumi.StringPtrInput
}

func (RepositoryReleaseConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryReleaseConfigState)(nil)).Elem()
}

type repositoryReleaseConfigArgs struct {
	// Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
	// Structure is documented below.
	CodeCompilationConfig *RepositoryReleaseConfigCodeCompilationConfig `pulumi:"codeCompilationConfig"`
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule *string `pulumi:"cronSchedule"`
	// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
	//
	// ***
	GitCommitish string `pulumi:"gitCommitish"`
	// The release's name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region
	Region *string `pulumi:"region"`
	// A reference to the Dataform repository
	Repository *string `pulumi:"repository"`
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a RepositoryReleaseConfig resource.
type RepositoryReleaseConfigArgs struct {
	// Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
	// Structure is documented below.
	CodeCompilationConfig RepositoryReleaseConfigCodeCompilationConfigPtrInput
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule pulumi.StringPtrInput
	// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
	//
	// ***
	GitCommitish pulumi.StringInput
	// The release's name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region
	Region pulumi.StringPtrInput
	// A reference to the Dataform repository
	Repository pulumi.StringPtrInput
	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone pulumi.StringPtrInput
}

func (RepositoryReleaseConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryReleaseConfigArgs)(nil)).Elem()
}

type RepositoryReleaseConfigInput interface {
	pulumi.Input

	ToRepositoryReleaseConfigOutput() RepositoryReleaseConfigOutput
	ToRepositoryReleaseConfigOutputWithContext(ctx context.Context) RepositoryReleaseConfigOutput
}

func (*RepositoryReleaseConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryReleaseConfig)(nil)).Elem()
}

func (i *RepositoryReleaseConfig) ToRepositoryReleaseConfigOutput() RepositoryReleaseConfigOutput {
	return i.ToRepositoryReleaseConfigOutputWithContext(context.Background())
}

func (i *RepositoryReleaseConfig) ToRepositoryReleaseConfigOutputWithContext(ctx context.Context) RepositoryReleaseConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryReleaseConfigOutput)
}

// RepositoryReleaseConfigArrayInput is an input type that accepts RepositoryReleaseConfigArray and RepositoryReleaseConfigArrayOutput values.
// You can construct a concrete instance of `RepositoryReleaseConfigArrayInput` via:
//
//	RepositoryReleaseConfigArray{ RepositoryReleaseConfigArgs{...} }
type RepositoryReleaseConfigArrayInput interface {
	pulumi.Input

	ToRepositoryReleaseConfigArrayOutput() RepositoryReleaseConfigArrayOutput
	ToRepositoryReleaseConfigArrayOutputWithContext(context.Context) RepositoryReleaseConfigArrayOutput
}

type RepositoryReleaseConfigArray []RepositoryReleaseConfigInput

func (RepositoryReleaseConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryReleaseConfig)(nil)).Elem()
}

func (i RepositoryReleaseConfigArray) ToRepositoryReleaseConfigArrayOutput() RepositoryReleaseConfigArrayOutput {
	return i.ToRepositoryReleaseConfigArrayOutputWithContext(context.Background())
}

func (i RepositoryReleaseConfigArray) ToRepositoryReleaseConfigArrayOutputWithContext(ctx context.Context) RepositoryReleaseConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryReleaseConfigArrayOutput)
}

// RepositoryReleaseConfigMapInput is an input type that accepts RepositoryReleaseConfigMap and RepositoryReleaseConfigMapOutput values.
// You can construct a concrete instance of `RepositoryReleaseConfigMapInput` via:
//
//	RepositoryReleaseConfigMap{ "key": RepositoryReleaseConfigArgs{...} }
type RepositoryReleaseConfigMapInput interface {
	pulumi.Input

	ToRepositoryReleaseConfigMapOutput() RepositoryReleaseConfigMapOutput
	ToRepositoryReleaseConfigMapOutputWithContext(context.Context) RepositoryReleaseConfigMapOutput
}

type RepositoryReleaseConfigMap map[string]RepositoryReleaseConfigInput

func (RepositoryReleaseConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryReleaseConfig)(nil)).Elem()
}

func (i RepositoryReleaseConfigMap) ToRepositoryReleaseConfigMapOutput() RepositoryReleaseConfigMapOutput {
	return i.ToRepositoryReleaseConfigMapOutputWithContext(context.Background())
}

func (i RepositoryReleaseConfigMap) ToRepositoryReleaseConfigMapOutputWithContext(ctx context.Context) RepositoryReleaseConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryReleaseConfigMapOutput)
}

type RepositoryReleaseConfigOutput struct{ *pulumi.OutputState }

func (RepositoryReleaseConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryReleaseConfig)(nil)).Elem()
}

func (o RepositoryReleaseConfigOutput) ToRepositoryReleaseConfigOutput() RepositoryReleaseConfigOutput {
	return o
}

func (o RepositoryReleaseConfigOutput) ToRepositoryReleaseConfigOutputWithContext(ctx context.Context) RepositoryReleaseConfigOutput {
	return o
}

// Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
// Structure is documented below.
func (o RepositoryReleaseConfigOutput) CodeCompilationConfig() RepositoryReleaseConfigCodeCompilationConfigPtrOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) RepositoryReleaseConfigCodeCompilationConfigPtrOutput {
		return v.CodeCompilationConfig
	}).(RepositoryReleaseConfigCodeCompilationConfigPtrOutput)
}

// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
func (o RepositoryReleaseConfigOutput) CronSchedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) pulumi.StringPtrOutput { return v.CronSchedule }).(pulumi.StringPtrOutput)
}

// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
//
// ***
func (o RepositoryReleaseConfigOutput) GitCommitish() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) pulumi.StringOutput { return v.GitCommitish }).(pulumi.StringOutput)
}

// The release's name.
func (o RepositoryReleaseConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RepositoryReleaseConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Records of the 10 most recent scheduled release attempts, ordered in in descending order of releaseTime. Updated whenever automatic creation of a compilation result is triggered by cronSchedule.
// Structure is documented below.
func (o RepositoryReleaseConfigOutput) RecentScheduledReleaseRecords() RepositoryReleaseConfigRecentScheduledReleaseRecordArrayOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) RepositoryReleaseConfigRecentScheduledReleaseRecordArrayOutput {
		return v.RecentScheduledReleaseRecords
	}).(RepositoryReleaseConfigRecentScheduledReleaseRecordArrayOutput)
}

// A reference to the region
func (o RepositoryReleaseConfigOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// A reference to the Dataform repository
func (o RepositoryReleaseConfigOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) pulumi.StringPtrOutput { return v.Repository }).(pulumi.StringPtrOutput)
}

// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
func (o RepositoryReleaseConfigOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryReleaseConfig) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type RepositoryReleaseConfigArrayOutput struct{ *pulumi.OutputState }

func (RepositoryReleaseConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryReleaseConfig)(nil)).Elem()
}

func (o RepositoryReleaseConfigArrayOutput) ToRepositoryReleaseConfigArrayOutput() RepositoryReleaseConfigArrayOutput {
	return o
}

func (o RepositoryReleaseConfigArrayOutput) ToRepositoryReleaseConfigArrayOutputWithContext(ctx context.Context) RepositoryReleaseConfigArrayOutput {
	return o
}

func (o RepositoryReleaseConfigArrayOutput) Index(i pulumi.IntInput) RepositoryReleaseConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryReleaseConfig {
		return vs[0].([]*RepositoryReleaseConfig)[vs[1].(int)]
	}).(RepositoryReleaseConfigOutput)
}

type RepositoryReleaseConfigMapOutput struct{ *pulumi.OutputState }

func (RepositoryReleaseConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryReleaseConfig)(nil)).Elem()
}

func (o RepositoryReleaseConfigMapOutput) ToRepositoryReleaseConfigMapOutput() RepositoryReleaseConfigMapOutput {
	return o
}

func (o RepositoryReleaseConfigMapOutput) ToRepositoryReleaseConfigMapOutputWithContext(ctx context.Context) RepositoryReleaseConfigMapOutput {
	return o
}

func (o RepositoryReleaseConfigMapOutput) MapIndex(k pulumi.StringInput) RepositoryReleaseConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryReleaseConfig {
		return vs[0].(map[string]*RepositoryReleaseConfig)[vs[1].(string)]
	}).(RepositoryReleaseConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryReleaseConfigInput)(nil)).Elem(), &RepositoryReleaseConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryReleaseConfigArrayInput)(nil)).Elem(), RepositoryReleaseConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryReleaseConfigMapInput)(nil)).Elem(), RepositoryReleaseConfigMap{})
	pulumi.RegisterOutputType(RepositoryReleaseConfigOutput{})
	pulumi.RegisterOutputType(RepositoryReleaseConfigArrayOutput{})
	pulumi.RegisterOutputType(RepositoryReleaseConfigMapOutput{})
}
