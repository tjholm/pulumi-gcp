// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package osconfig

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch deployments are configurations that individual patch jobs use to complete a patch.
// These configurations include instance filter, package repository settings, and a schedule.
//
// To get more information about PatchDeployment, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/osconfig/rest)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/os-patch-management)
//
// ## Example Usage
// ### Os Config Patch Deployment Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/osconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := osconfig.NewPatchDeployment(ctx, "patch", &osconfig.PatchDeploymentArgs{
//				InstanceFilter: &osconfig.PatchDeploymentInstanceFilterArgs{
//					All: pulumi.Bool(true),
//				},
//				OneTimeSchedule: &osconfig.PatchDeploymentOneTimeScheduleArgs{
//					ExecuteTime: pulumi.String("2999-10-10T10:10:10.045123456Z"),
//				},
//				PatchDeploymentId: pulumi.String("patch-deploy"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Os Config Patch Deployment Daily
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/osconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := osconfig.NewPatchDeployment(ctx, "patch", &osconfig.PatchDeploymentArgs{
//				InstanceFilter: &osconfig.PatchDeploymentInstanceFilterArgs{
//					All: pulumi.Bool(true),
//				},
//				PatchDeploymentId: pulumi.String("patch-deploy"),
//				RecurringSchedule: &osconfig.PatchDeploymentRecurringScheduleArgs{
//					TimeOfDay: &osconfig.PatchDeploymentRecurringScheduleTimeOfDayArgs{
//						Hours:   pulumi.Int(0),
//						Minutes: pulumi.Int(30),
//						Nanos:   pulumi.Int(20),
//						Seconds: pulumi.Int(30),
//					},
//					TimeZone: &osconfig.PatchDeploymentRecurringScheduleTimeZoneArgs{
//						Id: pulumi.String("America/New_York"),
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
// ### Os Config Patch Deployment Daily Midnight
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/osconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := osconfig.NewPatchDeployment(ctx, "patch", &osconfig.PatchDeploymentArgs{
//				InstanceFilter: &osconfig.PatchDeploymentInstanceFilterArgs{
//					All: pulumi.Bool(true),
//				},
//				PatchDeploymentId: pulumi.String("patch-deploy"),
//				RecurringSchedule: &osconfig.PatchDeploymentRecurringScheduleArgs{
//					TimeOfDay: &osconfig.PatchDeploymentRecurringScheduleTimeOfDayArgs{
//						Hours:   pulumi.Int(0),
//						Minutes: pulumi.Int(0),
//						Nanos:   pulumi.Int(0),
//						Seconds: pulumi.Int(0),
//					},
//					TimeZone: &osconfig.PatchDeploymentRecurringScheduleTimeZoneArgs{
//						Id: pulumi.String("America/New_York"),
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
// ### Os Config Patch Deployment Instance
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/osconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myImage, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
//				Family:  pulumi.StringRef("debian-11"),
//				Project: pulumi.StringRef("debian-cloud"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			foobar, err := compute.NewInstance(ctx, "foobar", &compute.InstanceArgs{
//				MachineType:  pulumi.String("e2-medium"),
//				Zone:         pulumi.String("us-central1-a"),
//				CanIpForward: pulumi.Bool(false),
//				Tags: pulumi.StringArray{
//					pulumi.String("foo"),
//					pulumi.String("bar"),
//				},
//				BootDisk: &compute.InstanceBootDiskArgs{
//					InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
//						Image: *pulumi.String(myImage.SelfLink),
//					},
//				},
//				NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
//					&compute.InstanceNetworkInterfaceArgs{
//						Network: pulumi.String("default"),
//					},
//				},
//				Metadata: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = osconfig.NewPatchDeployment(ctx, "patch", &osconfig.PatchDeploymentArgs{
//				PatchDeploymentId: pulumi.String("patch-deploy"),
//				InstanceFilter: &osconfig.PatchDeploymentInstanceFilterArgs{
//					Instances: pulumi.StringArray{
//						foobar.ID(),
//					},
//				},
//				PatchConfig: &osconfig.PatchDeploymentPatchConfigArgs{
//					Yum: &osconfig.PatchDeploymentPatchConfigYumArgs{
//						Security: pulumi.Bool(true),
//						Minimal:  pulumi.Bool(true),
//						Excludes: pulumi.StringArray{
//							pulumi.String("bash"),
//						},
//					},
//				},
//				RecurringSchedule: &osconfig.PatchDeploymentRecurringScheduleArgs{
//					TimeZone: &osconfig.PatchDeploymentRecurringScheduleTimeZoneArgs{
//						Id: pulumi.String("America/New_York"),
//					},
//					TimeOfDay: &osconfig.PatchDeploymentRecurringScheduleTimeOfDayArgs{
//						Hours:   pulumi.Int(0),
//						Minutes: pulumi.Int(30),
//						Seconds: pulumi.Int(30),
//						Nanos:   pulumi.Int(20),
//					},
//					Monthly: &osconfig.PatchDeploymentRecurringScheduleMonthlyArgs{
//						MonthDay: pulumi.Int(1),
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
// ### Os Config Patch Deployment Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/osconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := osconfig.NewPatchDeployment(ctx, "patch", &osconfig.PatchDeploymentArgs{
//				Duration: pulumi.String("10s"),
//				InstanceFilter: &osconfig.PatchDeploymentInstanceFilterArgs{
//					GroupLabels: osconfig.PatchDeploymentInstanceFilterGroupLabelArray{
//						&osconfig.PatchDeploymentInstanceFilterGroupLabelArgs{
//							Labels: pulumi.StringMap{
//								"app": pulumi.String("web"),
//								"env": pulumi.String("dev"),
//							},
//						},
//					},
//					InstanceNamePrefixes: pulumi.StringArray{
//						pulumi.String("test-"),
//					},
//					Zones: pulumi.StringArray{
//						pulumi.String("us-central1-a"),
//						pulumi.String("us-central-1c"),
//					},
//				},
//				PatchConfig: &osconfig.PatchDeploymentPatchConfigArgs{
//					Apt: &osconfig.PatchDeploymentPatchConfigAptArgs{
//						Excludes: pulumi.StringArray{
//							pulumi.String("python"),
//						},
//						Type: pulumi.String("DIST"),
//					},
//					Goo: &osconfig.PatchDeploymentPatchConfigGooArgs{
//						Enabled: pulumi.Bool(true),
//					},
//					MigInstancesAllowed: pulumi.Bool(true),
//					PostStep: &osconfig.PatchDeploymentPatchConfigPostStepArgs{
//						LinuxExecStepConfig: &osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigArgs{
//							GcsObject: &osconfig.PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObjectArgs{
//								Bucket:           pulumi.String("my-patch-scripts"),
//								GenerationNumber: pulumi.String("1523477886880"),
//								Object:           pulumi.String("linux/post_patch_script"),
//							},
//						},
//						WindowsExecStepConfig: &osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigArgs{
//							GcsObject: &osconfig.PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObjectArgs{
//								Bucket:           pulumi.String("my-patch-scripts"),
//								GenerationNumber: pulumi.String("135920493447"),
//								Object:           pulumi.String("windows/post_patch_script.ps1"),
//							},
//							Interpreter: pulumi.String("POWERSHELL"),
//						},
//					},
//					PreStep: &osconfig.PatchDeploymentPatchConfigPreStepArgs{
//						LinuxExecStepConfig: &osconfig.PatchDeploymentPatchConfigPreStepLinuxExecStepConfigArgs{
//							AllowedSuccessCodes: pulumi.IntArray{
//								pulumi.Int(0),
//								pulumi.Int(3),
//							},
//							LocalPath: pulumi.String("/tmp/pre_patch_script.sh"),
//						},
//						WindowsExecStepConfig: &osconfig.PatchDeploymentPatchConfigPreStepWindowsExecStepConfigArgs{
//							AllowedSuccessCodes: pulumi.IntArray{
//								pulumi.Int(0),
//								pulumi.Int(2),
//							},
//							Interpreter: pulumi.String("SHELL"),
//							LocalPath:   pulumi.String("C:\\Users\\user\\pre-patch-script.cmd"),
//						},
//					},
//					RebootConfig: pulumi.String("ALWAYS"),
//					WindowsUpdate: &osconfig.PatchDeploymentPatchConfigWindowsUpdateArgs{
//						Classifications: pulumi.StringArray{
//							pulumi.String("CRITICAL"),
//							pulumi.String("SECURITY"),
//							pulumi.String("UPDATE"),
//						},
//					},
//					Yum: &osconfig.PatchDeploymentPatchConfigYumArgs{
//						Excludes: pulumi.StringArray{
//							pulumi.String("bash"),
//						},
//						Minimal:  pulumi.Bool(true),
//						Security: pulumi.Bool(true),
//					},
//					Zypper: &osconfig.PatchDeploymentPatchConfigZypperArgs{
//						Categories: pulumi.StringArray{
//							pulumi.String("security"),
//						},
//					},
//				},
//				PatchDeploymentId: pulumi.String("patch-deploy"),
//				RecurringSchedule: &osconfig.PatchDeploymentRecurringScheduleArgs{
//					Monthly: &osconfig.PatchDeploymentRecurringScheduleMonthlyArgs{
//						WeekDayOfMonth: &osconfig.PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthArgs{
//							DayOfWeek:   pulumi.String("TUESDAY"),
//							WeekOrdinal: -1,
//						},
//					},
//					TimeOfDay: &osconfig.PatchDeploymentRecurringScheduleTimeOfDayArgs{
//						Hours:   pulumi.Int(0),
//						Minutes: pulumi.Int(30),
//						Nanos:   pulumi.Int(20),
//						Seconds: pulumi.Int(30),
//					},
//					TimeZone: &osconfig.PatchDeploymentRecurringScheduleTimeZoneArgs{
//						Id: pulumi.String("America/New_York"),
//					},
//				},
//				Rollout: &osconfig.PatchDeploymentRolloutArgs{
//					DisruptionBudget: &osconfig.PatchDeploymentRolloutDisruptionBudgetArgs{
//						Fixed: pulumi.Int(1),
//					},
//					Mode: pulumi.String("ZONE_BY_ZONE"),
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
// # PatchDeployment can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:osconfig/patchDeployment:PatchDeployment default projects/{{project}}/patchDeployments/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:osconfig/patchDeployment:PatchDeployment default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:osconfig/patchDeployment:PatchDeployment default {{name}}
//
// ```
type PatchDeployment struct {
	pulumi.CustomResourceState

	// Time the patch deployment was created. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration pulumi.StringPtrOutput `pulumi:"duration"`
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter PatchDeploymentInstanceFilterOutput `pulumi:"instanceFilter"`
	// The time the last patch job ran successfully.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	LastExecuteTime pulumi.StringOutput `pulumi:"lastExecuteTime"`
	// Unique name for the patch deployment resource in a project.
	// The patch deployment name is in the form: projects/{project_id}/patchDeployments/{patchDeploymentId}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule PatchDeploymentOneTimeSchedulePtrOutput `pulumi:"oneTimeSchedule"`
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig PatchDeploymentPatchConfigPtrOutput `pulumi:"patchConfig"`
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId pulumi.StringOutput `pulumi:"patchDeploymentId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule PatchDeploymentRecurringSchedulePtrOutput `pulumi:"recurringSchedule"`
	// Rollout strategy of the patch job.
	// Structure is documented below.
	Rollout PatchDeploymentRolloutPtrOutput `pulumi:"rollout"`
	// Time the patch deployment was last updated. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewPatchDeployment registers a new resource with the given unique name, arguments, and options.
func NewPatchDeployment(ctx *pulumi.Context,
	name string, args *PatchDeploymentArgs, opts ...pulumi.ResourceOption) (*PatchDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceFilter == nil {
		return nil, errors.New("invalid value for required argument 'InstanceFilter'")
	}
	if args.PatchDeploymentId == nil {
		return nil, errors.New("invalid value for required argument 'PatchDeploymentId'")
	}
	var resource PatchDeployment
	err := ctx.RegisterResource("gcp:osconfig/patchDeployment:PatchDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPatchDeployment gets an existing PatchDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PatchDeploymentState, opts ...pulumi.ResourceOption) (*PatchDeployment, error) {
	var resource PatchDeployment
	err := ctx.ReadResource("gcp:osconfig/patchDeployment:PatchDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PatchDeployment resources.
type patchDeploymentState struct {
	// Time the patch deployment was created. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration *string `pulumi:"duration"`
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter *PatchDeploymentInstanceFilter `pulumi:"instanceFilter"`
	// The time the last patch job ran successfully.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	LastExecuteTime *string `pulumi:"lastExecuteTime"`
	// Unique name for the patch deployment resource in a project.
	// The patch deployment name is in the form: projects/{project_id}/patchDeployments/{patchDeploymentId}.
	Name *string `pulumi:"name"`
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule *PatchDeploymentOneTimeSchedule `pulumi:"oneTimeSchedule"`
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig *PatchDeploymentPatchConfig `pulumi:"patchConfig"`
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId *string `pulumi:"patchDeploymentId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule *PatchDeploymentRecurringSchedule `pulumi:"recurringSchedule"`
	// Rollout strategy of the patch job.
	// Structure is documented below.
	Rollout *PatchDeploymentRollout `pulumi:"rollout"`
	// Time the patch deployment was last updated. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type PatchDeploymentState struct {
	// Time the patch deployment was created. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration pulumi.StringPtrInput
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter PatchDeploymentInstanceFilterPtrInput
	// The time the last patch job ran successfully.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	LastExecuteTime pulumi.StringPtrInput
	// Unique name for the patch deployment resource in a project.
	// The patch deployment name is in the form: projects/{project_id}/patchDeployments/{patchDeploymentId}.
	Name pulumi.StringPtrInput
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule PatchDeploymentOneTimeSchedulePtrInput
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig PatchDeploymentPatchConfigPtrInput
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule PatchDeploymentRecurringSchedulePtrInput
	// Rollout strategy of the patch job.
	// Structure is documented below.
	Rollout PatchDeploymentRolloutPtrInput
	// Time the patch deployment was last updated. Timestamp is in RFC3339 text format.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (PatchDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*patchDeploymentState)(nil)).Elem()
}

type patchDeploymentArgs struct {
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description *string `pulumi:"description"`
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration *string `pulumi:"duration"`
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter PatchDeploymentInstanceFilter `pulumi:"instanceFilter"`
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule *PatchDeploymentOneTimeSchedule `pulumi:"oneTimeSchedule"`
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig *PatchDeploymentPatchConfig `pulumi:"patchConfig"`
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId string `pulumi:"patchDeploymentId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule *PatchDeploymentRecurringSchedule `pulumi:"recurringSchedule"`
	// Rollout strategy of the patch job.
	// Structure is documented below.
	Rollout *PatchDeploymentRollout `pulumi:"rollout"`
}

// The set of arguments for constructing a PatchDeployment resource.
type PatchDeploymentArgs struct {
	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	Description pulumi.StringPtrInput
	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	Duration pulumi.StringPtrInput
	// VM instances to patch.
	// Structure is documented below.
	InstanceFilter PatchDeploymentInstanceFilterInput
	// Schedule a one-time execution.
	// Structure is documented below.
	OneTimeSchedule PatchDeploymentOneTimeSchedulePtrInput
	// Patch configuration that is applied.
	// Structure is documented below.
	PatchConfig PatchDeploymentPatchConfigPtrInput
	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	PatchDeploymentId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule recurring executions.
	// Structure is documented below.
	RecurringSchedule PatchDeploymentRecurringSchedulePtrInput
	// Rollout strategy of the patch job.
	// Structure is documented below.
	Rollout PatchDeploymentRolloutPtrInput
}

func (PatchDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*patchDeploymentArgs)(nil)).Elem()
}

type PatchDeploymentInput interface {
	pulumi.Input

	ToPatchDeploymentOutput() PatchDeploymentOutput
	ToPatchDeploymentOutputWithContext(ctx context.Context) PatchDeploymentOutput
}

func (*PatchDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchDeployment)(nil)).Elem()
}

func (i *PatchDeployment) ToPatchDeploymentOutput() PatchDeploymentOutput {
	return i.ToPatchDeploymentOutputWithContext(context.Background())
}

func (i *PatchDeployment) ToPatchDeploymentOutputWithContext(ctx context.Context) PatchDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchDeploymentOutput)
}

// PatchDeploymentArrayInput is an input type that accepts PatchDeploymentArray and PatchDeploymentArrayOutput values.
// You can construct a concrete instance of `PatchDeploymentArrayInput` via:
//
//	PatchDeploymentArray{ PatchDeploymentArgs{...} }
type PatchDeploymentArrayInput interface {
	pulumi.Input

	ToPatchDeploymentArrayOutput() PatchDeploymentArrayOutput
	ToPatchDeploymentArrayOutputWithContext(context.Context) PatchDeploymentArrayOutput
}

type PatchDeploymentArray []PatchDeploymentInput

func (PatchDeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PatchDeployment)(nil)).Elem()
}

func (i PatchDeploymentArray) ToPatchDeploymentArrayOutput() PatchDeploymentArrayOutput {
	return i.ToPatchDeploymentArrayOutputWithContext(context.Background())
}

func (i PatchDeploymentArray) ToPatchDeploymentArrayOutputWithContext(ctx context.Context) PatchDeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchDeploymentArrayOutput)
}

// PatchDeploymentMapInput is an input type that accepts PatchDeploymentMap and PatchDeploymentMapOutput values.
// You can construct a concrete instance of `PatchDeploymentMapInput` via:
//
//	PatchDeploymentMap{ "key": PatchDeploymentArgs{...} }
type PatchDeploymentMapInput interface {
	pulumi.Input

	ToPatchDeploymentMapOutput() PatchDeploymentMapOutput
	ToPatchDeploymentMapOutputWithContext(context.Context) PatchDeploymentMapOutput
}

type PatchDeploymentMap map[string]PatchDeploymentInput

func (PatchDeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PatchDeployment)(nil)).Elem()
}

func (i PatchDeploymentMap) ToPatchDeploymentMapOutput() PatchDeploymentMapOutput {
	return i.ToPatchDeploymentMapOutputWithContext(context.Background())
}

func (i PatchDeploymentMap) ToPatchDeploymentMapOutputWithContext(ctx context.Context) PatchDeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchDeploymentMapOutput)
}

type PatchDeploymentOutput struct{ *pulumi.OutputState }

func (PatchDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchDeployment)(nil)).Elem()
}

func (o PatchDeploymentOutput) ToPatchDeploymentOutput() PatchDeploymentOutput {
	return o
}

func (o PatchDeploymentOutput) ToPatchDeploymentOutputWithContext(ctx context.Context) PatchDeploymentOutput {
	return o
}

// Time the patch deployment was created. Timestamp is in RFC3339 text format.
// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (o PatchDeploymentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchDeployment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the patch deployment. Length of the description is limited to 1024 characters.
func (o PatchDeploymentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchDeployment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Duration of the patch. After the duration ends, the patch times out.
// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
func (o PatchDeploymentOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchDeployment) pulumi.StringPtrOutput { return v.Duration }).(pulumi.StringPtrOutput)
}

// VM instances to patch.
// Structure is documented below.
func (o PatchDeploymentOutput) InstanceFilter() PatchDeploymentInstanceFilterOutput {
	return o.ApplyT(func(v *PatchDeployment) PatchDeploymentInstanceFilterOutput { return v.InstanceFilter }).(PatchDeploymentInstanceFilterOutput)
}

// The time the last patch job ran successfully.
// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (o PatchDeploymentOutput) LastExecuteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchDeployment) pulumi.StringOutput { return v.LastExecuteTime }).(pulumi.StringOutput)
}

// Unique name for the patch deployment resource in a project.
// The patch deployment name is in the form: projects/{project_id}/patchDeployments/{patchDeploymentId}.
func (o PatchDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Schedule a one-time execution.
// Structure is documented below.
func (o PatchDeploymentOutput) OneTimeSchedule() PatchDeploymentOneTimeSchedulePtrOutput {
	return o.ApplyT(func(v *PatchDeployment) PatchDeploymentOneTimeSchedulePtrOutput { return v.OneTimeSchedule }).(PatchDeploymentOneTimeSchedulePtrOutput)
}

// Patch configuration that is applied.
// Structure is documented below.
func (o PatchDeploymentOutput) PatchConfig() PatchDeploymentPatchConfigPtrOutput {
	return o.ApplyT(func(v *PatchDeployment) PatchDeploymentPatchConfigPtrOutput { return v.PatchConfig }).(PatchDeploymentPatchConfigPtrOutput)
}

// A name for the patch deployment in the project. When creating a name the following rules apply:
// * Must contain only lowercase letters, numbers, and hyphens.
// * Must start with a letter.
// * Must be between 1-63 characters.
// * Must end with a number or a letter.
// * Must be unique within the project.
func (o PatchDeploymentOutput) PatchDeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchDeployment) pulumi.StringOutput { return v.PatchDeploymentId }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o PatchDeploymentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchDeployment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Schedule recurring executions.
// Structure is documented below.
func (o PatchDeploymentOutput) RecurringSchedule() PatchDeploymentRecurringSchedulePtrOutput {
	return o.ApplyT(func(v *PatchDeployment) PatchDeploymentRecurringSchedulePtrOutput { return v.RecurringSchedule }).(PatchDeploymentRecurringSchedulePtrOutput)
}

// Rollout strategy of the patch job.
// Structure is documented below.
func (o PatchDeploymentOutput) Rollout() PatchDeploymentRolloutPtrOutput {
	return o.ApplyT(func(v *PatchDeployment) PatchDeploymentRolloutPtrOutput { return v.Rollout }).(PatchDeploymentRolloutPtrOutput)
}

// Time the patch deployment was last updated. Timestamp is in RFC3339 text format.
// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
func (o PatchDeploymentOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PatchDeployment) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type PatchDeploymentArrayOutput struct{ *pulumi.OutputState }

func (PatchDeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PatchDeployment)(nil)).Elem()
}

func (o PatchDeploymentArrayOutput) ToPatchDeploymentArrayOutput() PatchDeploymentArrayOutput {
	return o
}

func (o PatchDeploymentArrayOutput) ToPatchDeploymentArrayOutputWithContext(ctx context.Context) PatchDeploymentArrayOutput {
	return o
}

func (o PatchDeploymentArrayOutput) Index(i pulumi.IntInput) PatchDeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PatchDeployment {
		return vs[0].([]*PatchDeployment)[vs[1].(int)]
	}).(PatchDeploymentOutput)
}

type PatchDeploymentMapOutput struct{ *pulumi.OutputState }

func (PatchDeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PatchDeployment)(nil)).Elem()
}

func (o PatchDeploymentMapOutput) ToPatchDeploymentMapOutput() PatchDeploymentMapOutput {
	return o
}

func (o PatchDeploymentMapOutput) ToPatchDeploymentMapOutputWithContext(ctx context.Context) PatchDeploymentMapOutput {
	return o
}

func (o PatchDeploymentMapOutput) MapIndex(k pulumi.StringInput) PatchDeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PatchDeployment {
		return vs[0].(map[string]*PatchDeployment)[vs[1].(string)]
	}).(PatchDeploymentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PatchDeploymentInput)(nil)).Elem(), &PatchDeployment{})
	pulumi.RegisterInputType(reflect.TypeOf((*PatchDeploymentArrayInput)(nil)).Elem(), PatchDeploymentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PatchDeploymentMapInput)(nil)).Elem(), PatchDeploymentMap{})
	pulumi.RegisterOutputType(PatchDeploymentOutput{})
	pulumi.RegisterOutputType(PatchDeploymentArrayOutput{})
	pulumi.RegisterOutputType(PatchDeploymentMapOutput{})
}
