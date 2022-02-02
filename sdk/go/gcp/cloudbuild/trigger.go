// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudbuild

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration for an automated build in response to source repository changes.
//
// To get more information about Trigger, see:
//
// * [API documentation](https://cloud.google.com/cloud-build/docs/api/reference/rest/v1/projects.triggers)
// * How-to Guides
//     * [Automating builds using build triggers](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds)
//
// > **Note:** You can retrieve the email of the Cloud Build Service Account used in jobs by using the `projects.ServiceIdentity` resource.
//
// ## Example Usage
// ### Cloudbuild Trigger Filename
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudbuild"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudbuild.NewTrigger(ctx, "filename-trigger", &cloudbuild.TriggerArgs{
// 			Filename: pulumi.String("cloudbuild.yaml"),
// 			Substitutions: pulumi.StringMap{
// 				"_BAZ": pulumi.String("qux"),
// 				"_FOO": pulumi.String("bar"),
// 			},
// 			TriggerTemplate: &cloudbuild.TriggerTriggerTemplateArgs{
// 				BranchName: pulumi.String("master"),
// 				RepoName:   pulumi.String("my-repo"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloudbuild Trigger Build
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudbuild"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudbuild.NewTrigger(ctx, "build-trigger", &cloudbuild.TriggerArgs{
// 			Build: &cloudbuild.TriggerBuildArgs{
// 				Artifacts: &cloudbuild.TriggerBuildArtifactsArgs{
// 					Images: pulumi.StringArray{
// 						pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v", "gcr.io/", "$", "PROJECT_ID/", "$", "REPO_NAME:", "$", "COMMIT_SHA")),
// 					},
// 					Objects: &cloudbuild.TriggerBuildArtifactsObjectsArgs{
// 						Location: pulumi.String("gs://bucket/path/to/somewhere/"),
// 						Paths: pulumi.StringArray{
// 							pulumi.String("path"),
// 						},
// 					},
// 				},
// 				AvailableSecrets: &cloudbuild.TriggerBuildAvailableSecretsArgs{
// 					SecretManager: []map[string]interface{}{
// 						map[string]interface{}{
// 							"env":         "MY_SECRET",
// 							"versionName": "projects/myProject/secrets/mySecret/versions/latest",
// 						},
// 					},
// 				},
// 				LogsBucket: pulumi.String("gs://mybucket/logs"),
// 				Options: &cloudbuild.TriggerBuildOptionsArgs{
// 					DiskSizeGb:           pulumi.Int(100),
// 					DynamicSubstitutions: pulumi.Bool(true),
// 					Env: []string{
// 						"ekey = evalue",
// 					},
// 					LogStreamingOption:    pulumi.String("STREAM_OFF"),
// 					Logging:               pulumi.String("LEGACY"),
// 					MachineType:           pulumi.String("N1_HIGHCPU_8"),
// 					RequestedVerifyOption: pulumi.String("VERIFIED"),
// 					SecretEnv: []string{
// 						"secretenv = svalue",
// 					},
// 					SourceProvenanceHash: []string{
// 						"MD5",
// 					},
// 					SubstitutionOption: pulumi.String("ALLOW_LOOSE"),
// 					Volumes: cloudbuild.TriggerBuildOptionsVolumeArray{
// 						&cloudbuild.TriggerBuildOptionsVolumeArgs{
// 							Name: pulumi.String("v1"),
// 							Path: pulumi.String("v1"),
// 						},
// 					},
// 					WorkerPool: pulumi.String("pool"),
// 				},
// 				QueueTtl: pulumi.String("20s"),
// 				Secrets: cloudbuild.TriggerBuildSecretArray{
// 					&cloudbuild.TriggerBuildSecretArgs{
// 						KmsKeyName: pulumi.String("projects/myProject/locations/global/keyRings/keyring-name/cryptoKeys/key-name"),
// 						SecretEnv: pulumi.StringMap{
// 							"PASSWORD": pulumi.String("ZW5jcnlwdGVkLXBhc3N3b3JkCg=="),
// 						},
// 					},
// 				},
// 				Source: &cloudbuild.TriggerBuildSourceArgs{
// 					StorageSource: &cloudbuild.TriggerBuildSourceStorageSourceArgs{
// 						Bucket: pulumi.String("mybucket"),
// 						Object: pulumi.String("source_code.tar.gz"),
// 					},
// 				},
// 				Steps: cloudbuild.TriggerBuildStepArray{
// 					&cloudbuild.TriggerBuildStepArgs{
// 						Args: pulumi.StringArray{
// 							pulumi.String("cp"),
// 							pulumi.String("gs://mybucket/remotefile.zip"),
// 							pulumi.String("localfile.zip"),
// 						},
// 						Name: pulumi.String("gcr.io/cloud-builders/gsutil"),
// 						SecretEnv: []string{
// 							"MY_SECRET",
// 						},
// 						Timeout: pulumi.String("120s"),
// 					},
// 				},
// 				Substitutions: pulumi.StringMap{
// 					"_BAZ": pulumi.String("qux"),
// 					"_FOO": pulumi.String("bar"),
// 				},
// 				Tags: pulumi.StringArray{
// 					pulumi.String("build"),
// 					pulumi.String("newFeature"),
// 				},
// 			},
// 			TriggerTemplate: &cloudbuild.TriggerTriggerTemplateArgs{
// 				BranchName: pulumi.String("master"),
// 				RepoName:   pulumi.String("my-repo"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloudbuild Trigger Service Account
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudbuild"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project, err := organizations.LookupProject(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		cloudbuildServiceAccount, err := serviceAccount.NewAccount(ctx, "cloudbuildServiceAccount", &serviceAccount.AccountArgs{
// 			AccountId: pulumi.String("my-service-account"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		actAs, err := projects.NewIAMMember(ctx, "actAs", &projects.IAMMemberArgs{
// 			Project: pulumi.String(project.ProjectId),
// 			Role:    pulumi.String("roles/iam.serviceAccountUser"),
// 			Member: cloudbuildServiceAccount.Email.ApplyT(func(email string) (string, error) {
// 				return fmt.Sprintf("%v%v", "serviceAccount:", email), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		logsWriter, err := projects.NewIAMMember(ctx, "logsWriter", &projects.IAMMemberArgs{
// 			Project: pulumi.String(project.ProjectId),
// 			Role:    pulumi.String("roles/logging.logWriter"),
// 			Member: cloudbuildServiceAccount.Email.ApplyT(func(email string) (string, error) {
// 				return fmt.Sprintf("%v%v", "serviceAccount:", email), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudbuild.NewTrigger(ctx, "service-account-trigger", &cloudbuild.TriggerArgs{
// 			TriggerTemplate: &cloudbuild.TriggerTriggerTemplateArgs{
// 				BranchName: pulumi.String("master"),
// 				RepoName:   pulumi.String("my-repo"),
// 			},
// 			ServiceAccount: cloudbuildServiceAccount.ID(),
// 			Filename:       pulumi.String("cloudbuild.yaml"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			actAs,
// 			logsWriter,
// 		}))
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
// Trigger can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudbuild/trigger:Trigger default projects/{{project}}/triggers/{{trigger_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudbuild/trigger:Trigger default {{project}}/{{trigger_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudbuild/trigger:Trigger default {{trigger_id}}
// ```
type Trigger struct {
	pulumi.CustomResourceState

	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build TriggerBuildPtrOutput `pulumi:"build"`
	// Time when the trigger was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Human-readable description of the trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename pulumi.StringPtrOutput `pulumi:"filename"`
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	Github TriggerGithubPtrOutput `pulumi:"github"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayOutput `pulumi:"ignoredFiles"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles pulumi.StringArrayOutput `pulumi:"includedFiles"`
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// PubsubConfig describes the configuration of a trigger that creates
	// a build whenever a Pub/Sub message is published.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	PubsubConfig TriggerPubsubConfigPtrOutput `pulumi:"pubsubConfig"`
	// The service account used for all user-controlled operations including
	// triggers.patch, triggers.run, builds.create, and builds.cancel.
	// If no service account is set, then the standard Cloud Build service account
	// ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
	// Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
	ServiceAccount pulumi.StringPtrOutput `pulumi:"serviceAccount"`
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions pulumi.StringMapOutput `pulumi:"substitutions"`
	// Tags for annotation of a Build. These are not docker tags.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The unique identifier for the trigger.
	TriggerId pulumi.StringOutput `pulumi:"triggerId"`
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	TriggerTemplate TriggerTriggerTemplatePtrOutput `pulumi:"triggerTemplate"`
	// WebhookConfig describes the configuration of a trigger that creates
	// a build whenever a webhook is sent to a trigger's webhook URL.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	WebhookConfig TriggerWebhookConfigPtrOutput `pulumi:"webhookConfig"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		args = &TriggerArgs{}
	}

	var resource Trigger
	err := ctx.RegisterResource("gcp:cloudbuild/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("gcp:cloudbuild/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build *TriggerBuild `pulumi:"build"`
	// Time when the trigger was created.
	CreateTime *string `pulumi:"createTime"`
	// Human-readable description of the trigger.
	Description *string `pulumi:"description"`
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled *bool `pulumi:"disabled"`
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename *string `pulumi:"filename"`
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	Github *TriggerGithub `pulumi:"github"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles []string `pulumi:"ignoredFiles"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles []string `pulumi:"includedFiles"`
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// PubsubConfig describes the configuration of a trigger that creates
	// a build whenever a Pub/Sub message is published.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	PubsubConfig *TriggerPubsubConfig `pulumi:"pubsubConfig"`
	// The service account used for all user-controlled operations including
	// triggers.patch, triggers.run, builds.create, and builds.cancel.
	// If no service account is set, then the standard Cloud Build service account
	// ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
	// Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions map[string]string `pulumi:"substitutions"`
	// Tags for annotation of a Build. These are not docker tags.
	Tags []string `pulumi:"tags"`
	// The unique identifier for the trigger.
	TriggerId *string `pulumi:"triggerId"`
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	TriggerTemplate *TriggerTriggerTemplate `pulumi:"triggerTemplate"`
	// WebhookConfig describes the configuration of a trigger that creates
	// a build whenever a webhook is sent to a trigger's webhook URL.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	WebhookConfig *TriggerWebhookConfig `pulumi:"webhookConfig"`
}

type TriggerState struct {
	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build TriggerBuildPtrInput
	// Time when the trigger was created.
	CreateTime pulumi.StringPtrInput
	// Human-readable description of the trigger.
	Description pulumi.StringPtrInput
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled pulumi.BoolPtrInput
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename pulumi.StringPtrInput
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	Github TriggerGithubPtrInput
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayInput
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles pulumi.StringArrayInput
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// PubsubConfig describes the configuration of a trigger that creates
	// a build whenever a Pub/Sub message is published.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	PubsubConfig TriggerPubsubConfigPtrInput
	// The service account used for all user-controlled operations including
	// triggers.patch, triggers.run, builds.create, and builds.cancel.
	// If no service account is set, then the standard Cloud Build service account
	// ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
	// Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
	ServiceAccount pulumi.StringPtrInput
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions pulumi.StringMapInput
	// Tags for annotation of a Build. These are not docker tags.
	Tags pulumi.StringArrayInput
	// The unique identifier for the trigger.
	TriggerId pulumi.StringPtrInput
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	TriggerTemplate TriggerTriggerTemplatePtrInput
	// WebhookConfig describes the configuration of a trigger that creates
	// a build whenever a webhook is sent to a trigger's webhook URL.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	WebhookConfig TriggerWebhookConfigPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build *TriggerBuild `pulumi:"build"`
	// Human-readable description of the trigger.
	Description *string `pulumi:"description"`
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled *bool `pulumi:"disabled"`
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename *string `pulumi:"filename"`
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	Github *TriggerGithub `pulumi:"github"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles []string `pulumi:"ignoredFiles"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles []string `pulumi:"includedFiles"`
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// PubsubConfig describes the configuration of a trigger that creates
	// a build whenever a Pub/Sub message is published.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	PubsubConfig *TriggerPubsubConfig `pulumi:"pubsubConfig"`
	// The service account used for all user-controlled operations including
	// triggers.patch, triggers.run, builds.create, and builds.cancel.
	// If no service account is set, then the standard Cloud Build service account
	// ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
	// Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions map[string]string `pulumi:"substitutions"`
	// Tags for annotation of a Build. These are not docker tags.
	Tags []string `pulumi:"tags"`
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	TriggerTemplate *TriggerTriggerTemplate `pulumi:"triggerTemplate"`
	// WebhookConfig describes the configuration of a trigger that creates
	// a build whenever a webhook is sent to a trigger's webhook URL.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	WebhookConfig *TriggerWebhookConfig `pulumi:"webhookConfig"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build TriggerBuildPtrInput
	// Human-readable description of the trigger.
	Description pulumi.StringPtrInput
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled pulumi.BoolPtrInput
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename pulumi.StringPtrInput
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	Github TriggerGithubPtrInput
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayInput
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles pulumi.StringArrayInput
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// PubsubConfig describes the configuration of a trigger that creates
	// a build whenever a Pub/Sub message is published.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	PubsubConfig TriggerPubsubConfigPtrInput
	// The service account used for all user-controlled operations including
	// triggers.patch, triggers.run, builds.create, and builds.cancel.
	// If no service account is set, then the standard Cloud Build service account
	// ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
	// Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
	ServiceAccount pulumi.StringPtrInput
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions pulumi.StringMapInput
	// Tags for annotation of a Build. These are not docker tags.
	Tags pulumi.StringArrayInput
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	TriggerTemplate TriggerTriggerTemplatePtrInput
	// WebhookConfig describes the configuration of a trigger that creates
	// a build whenever a webhook is sent to a trigger's webhook URL.
	// One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
	// Structure is documented below.
	WebhookConfig TriggerWebhookConfigPtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

// TriggerArrayInput is an input type that accepts TriggerArray and TriggerArrayOutput values.
// You can construct a concrete instance of `TriggerArrayInput` via:
//
//          TriggerArray{ TriggerArgs{...} }
type TriggerArrayInput interface {
	pulumi.Input

	ToTriggerArrayOutput() TriggerArrayOutput
	ToTriggerArrayOutputWithContext(context.Context) TriggerArrayOutput
}

type TriggerArray []TriggerInput

func (TriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (i TriggerArray) ToTriggerArrayOutput() TriggerArrayOutput {
	return i.ToTriggerArrayOutputWithContext(context.Background())
}

func (i TriggerArray) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerArrayOutput)
}

// TriggerMapInput is an input type that accepts TriggerMap and TriggerMapOutput values.
// You can construct a concrete instance of `TriggerMapInput` via:
//
//          TriggerMap{ "key": TriggerArgs{...} }
type TriggerMapInput interface {
	pulumi.Input

	ToTriggerMapOutput() TriggerMapOutput
	ToTriggerMapOutputWithContext(context.Context) TriggerMapOutput
}

type TriggerMap map[string]TriggerInput

func (TriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (i TriggerMap) ToTriggerMapOutput() TriggerMapOutput {
	return i.ToTriggerMapOutputWithContext(context.Background())
}

func (i TriggerMap) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMapOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

type TriggerArrayOutput struct{ *pulumi.OutputState }

func (TriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (o TriggerArrayOutput) ToTriggerArrayOutput() TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) Index(i pulumi.IntInput) TriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].([]*Trigger)[vs[1].(int)]
	}).(TriggerOutput)
}

type TriggerMapOutput struct{ *pulumi.OutputState }

func (TriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (o TriggerMapOutput) ToTriggerMapOutput() TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) MapIndex(k pulumi.StringInput) TriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].(map[string]*Trigger)[vs[1].(string)]
	}).(TriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerArrayInput)(nil)).Elem(), TriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerMapInput)(nil)).Elem(), TriggerMap{})
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerArrayOutput{})
	pulumi.RegisterOutputType(TriggerMapOutput{})
}
