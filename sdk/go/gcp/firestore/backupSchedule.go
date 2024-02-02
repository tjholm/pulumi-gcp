// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firestore

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A backup schedule for a Cloud Firestore Database.
// This resource is owned by the database it is backing up, and is deleted along with the database.
// The actual backups are not though.
//
// To get more information about BackupSchedule, see:
//
// * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.backupSchedules)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/firestore/docs/backups)
//
// > **Warning:** This resource creates a Firestore Backup Schedule on a project that already has
// a Firestore database.
// This resource is owned by the database it is backing up, and is deleted along
// with the database. The actual backups are not though.
//
// ## Example Usage
// ### Firestore Backup Schedule Daily
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/firestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			database, err := firestore.NewDatabase(ctx, "database", &firestore.DatabaseArgs{
//				Project:               pulumi.String("my-project-name"),
//				LocationId:            pulumi.String("nam5"),
//				Type:                  pulumi.String("FIRESTORE_NATIVE"),
//				DeleteProtectionState: pulumi.String("DELETE_PROTECTION_ENABLED"),
//				DeletionPolicy:        pulumi.String("DELETE"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = firestore.NewBackupSchedule(ctx, "daily-backup", &firestore.BackupScheduleArgs{
//				Project:         pulumi.String("my-project-name"),
//				Database:        database.Name,
//				Retention:       pulumi.String("604800s"),
//				DailyRecurrence: nil,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Firestore Backup Schedule Weekly
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/firestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			database, err := firestore.NewDatabase(ctx, "database", &firestore.DatabaseArgs{
//				Project:               pulumi.String("my-project-name"),
//				LocationId:            pulumi.String("nam5"),
//				Type:                  pulumi.String("FIRESTORE_NATIVE"),
//				DeleteProtectionState: pulumi.String("DELETE_PROTECTION_ENABLED"),
//				DeletionPolicy:        pulumi.String("DELETE"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = firestore.NewBackupSchedule(ctx, "weekly-backup", &firestore.BackupScheduleArgs{
//				Project:   pulumi.String("my-project-name"),
//				Database:  database.Name,
//				Retention: pulumi.String("8467200s"),
//				WeeklyRecurrence: &firestore.BackupScheduleWeeklyRecurrenceArgs{
//					Day: pulumi.String("SUNDAY"),
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
// BackupSchedule can be imported using any of these accepted formats* `projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}` * `{{project}}/{{database}}/{{name}}` * `{{database}}/{{name}}` When using the `pulumi import` command, BackupSchedule can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:firestore/backupSchedule:BackupSchedule default projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firestore/backupSchedule:BackupSchedule default {{project}}/{{database}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firestore/backupSchedule:BackupSchedule default {{database}}/{{name}}
//
// ```
type BackupSchedule struct {
	pulumi.CustomResourceState

	// For a schedule that runs daily at a specified time.
	DailyRecurrence BackupScheduleDailyRecurrencePtrOutput `pulumi:"dailyRecurrence"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// The unique backup schedule identifier across all locations and databases for the given project. Format:
	// `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
	//
	// ***
	Retention pulumi.StringOutput `pulumi:"retention"`
	// For a schedule that runs weekly on a specific day and time.
	// Structure is documented below.
	WeeklyRecurrence BackupScheduleWeeklyRecurrencePtrOutput `pulumi:"weeklyRecurrence"`
}

// NewBackupSchedule registers a new resource with the given unique name, arguments, and options.
func NewBackupSchedule(ctx *pulumi.Context,
	name string, args *BackupScheduleArgs, opts ...pulumi.ResourceOption) (*BackupSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Retention == nil {
		return nil, errors.New("invalid value for required argument 'Retention'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupSchedule
	err := ctx.RegisterResource("gcp:firestore/backupSchedule:BackupSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupSchedule gets an existing BackupSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupScheduleState, opts ...pulumi.ResourceOption) (*BackupSchedule, error) {
	var resource BackupSchedule
	err := ctx.ReadResource("gcp:firestore/backupSchedule:BackupSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupSchedule resources.
type backupScheduleState struct {
	// For a schedule that runs daily at a specified time.
	DailyRecurrence *BackupScheduleDailyRecurrence `pulumi:"dailyRecurrence"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database *string `pulumi:"database"`
	// The unique backup schedule identifier across all locations and databases for the given project. Format:
	// `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
	//
	// ***
	Retention *string `pulumi:"retention"`
	// For a schedule that runs weekly on a specific day and time.
	// Structure is documented below.
	WeeklyRecurrence *BackupScheduleWeeklyRecurrence `pulumi:"weeklyRecurrence"`
}

type BackupScheduleState struct {
	// For a schedule that runs daily at a specified time.
	DailyRecurrence BackupScheduleDailyRecurrencePtrInput
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrInput
	// The unique backup schedule identifier across all locations and databases for the given project. Format:
	// `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
	//
	// ***
	Retention pulumi.StringPtrInput
	// For a schedule that runs weekly on a specific day and time.
	// Structure is documented below.
	WeeklyRecurrence BackupScheduleWeeklyRecurrencePtrInput
}

func (BackupScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupScheduleState)(nil)).Elem()
}

type backupScheduleArgs struct {
	// For a schedule that runs daily at a specified time.
	DailyRecurrence *BackupScheduleDailyRecurrence `pulumi:"dailyRecurrence"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database *string `pulumi:"database"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
	//
	// ***
	Retention string `pulumi:"retention"`
	// For a schedule that runs weekly on a specific day and time.
	// Structure is documented below.
	WeeklyRecurrence *BackupScheduleWeeklyRecurrence `pulumi:"weeklyRecurrence"`
}

// The set of arguments for constructing a BackupSchedule resource.
type BackupScheduleArgs struct {
	// For a schedule that runs daily at a specified time.
	DailyRecurrence BackupScheduleDailyRecurrencePtrInput
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
	//
	// ***
	Retention pulumi.StringInput
	// For a schedule that runs weekly on a specific day and time.
	// Structure is documented below.
	WeeklyRecurrence BackupScheduleWeeklyRecurrencePtrInput
}

func (BackupScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupScheduleArgs)(nil)).Elem()
}

type BackupScheduleInput interface {
	pulumi.Input

	ToBackupScheduleOutput() BackupScheduleOutput
	ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput
}

func (*BackupSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (i *BackupSchedule) ToBackupScheduleOutput() BackupScheduleOutput {
	return i.ToBackupScheduleOutputWithContext(context.Background())
}

func (i *BackupSchedule) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleOutput)
}

// BackupScheduleArrayInput is an input type that accepts BackupScheduleArray and BackupScheduleArrayOutput values.
// You can construct a concrete instance of `BackupScheduleArrayInput` via:
//
//	BackupScheduleArray{ BackupScheduleArgs{...} }
type BackupScheduleArrayInput interface {
	pulumi.Input

	ToBackupScheduleArrayOutput() BackupScheduleArrayOutput
	ToBackupScheduleArrayOutputWithContext(context.Context) BackupScheduleArrayOutput
}

type BackupScheduleArray []BackupScheduleInput

func (BackupScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupSchedule)(nil)).Elem()
}

func (i BackupScheduleArray) ToBackupScheduleArrayOutput() BackupScheduleArrayOutput {
	return i.ToBackupScheduleArrayOutputWithContext(context.Background())
}

func (i BackupScheduleArray) ToBackupScheduleArrayOutputWithContext(ctx context.Context) BackupScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleArrayOutput)
}

// BackupScheduleMapInput is an input type that accepts BackupScheduleMap and BackupScheduleMapOutput values.
// You can construct a concrete instance of `BackupScheduleMapInput` via:
//
//	BackupScheduleMap{ "key": BackupScheduleArgs{...} }
type BackupScheduleMapInput interface {
	pulumi.Input

	ToBackupScheduleMapOutput() BackupScheduleMapOutput
	ToBackupScheduleMapOutputWithContext(context.Context) BackupScheduleMapOutput
}

type BackupScheduleMap map[string]BackupScheduleInput

func (BackupScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupSchedule)(nil)).Elem()
}

func (i BackupScheduleMap) ToBackupScheduleMapOutput() BackupScheduleMapOutput {
	return i.ToBackupScheduleMapOutputWithContext(context.Background())
}

func (i BackupScheduleMap) ToBackupScheduleMapOutputWithContext(ctx context.Context) BackupScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleMapOutput)
}

type BackupScheduleOutput struct{ *pulumi.OutputState }

func (BackupScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (o BackupScheduleOutput) ToBackupScheduleOutput() BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return o
}

// For a schedule that runs daily at a specified time.
func (o BackupScheduleOutput) DailyRecurrence() BackupScheduleDailyRecurrencePtrOutput {
	return o.ApplyT(func(v *BackupSchedule) BackupScheduleDailyRecurrencePtrOutput { return v.DailyRecurrence }).(BackupScheduleDailyRecurrencePtrOutput)
}

// The Firestore database id. Defaults to `"(default)"`.
func (o BackupScheduleOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringPtrOutput { return v.Database }).(pulumi.StringPtrOutput)
}

// The unique backup schedule identifier across all locations and databases for the given project. Format:
// `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
func (o BackupScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o BackupScheduleOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
// For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
//
// ***
func (o BackupScheduleOutput) Retention() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.Retention }).(pulumi.StringOutput)
}

// For a schedule that runs weekly on a specific day and time.
// Structure is documented below.
func (o BackupScheduleOutput) WeeklyRecurrence() BackupScheduleWeeklyRecurrencePtrOutput {
	return o.ApplyT(func(v *BackupSchedule) BackupScheduleWeeklyRecurrencePtrOutput { return v.WeeklyRecurrence }).(BackupScheduleWeeklyRecurrencePtrOutput)
}

type BackupScheduleArrayOutput struct{ *pulumi.OutputState }

func (BackupScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupSchedule)(nil)).Elem()
}

func (o BackupScheduleArrayOutput) ToBackupScheduleArrayOutput() BackupScheduleArrayOutput {
	return o
}

func (o BackupScheduleArrayOutput) ToBackupScheduleArrayOutputWithContext(ctx context.Context) BackupScheduleArrayOutput {
	return o
}

func (o BackupScheduleArrayOutput) Index(i pulumi.IntInput) BackupScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupSchedule {
		return vs[0].([]*BackupSchedule)[vs[1].(int)]
	}).(BackupScheduleOutput)
}

type BackupScheduleMapOutput struct{ *pulumi.OutputState }

func (BackupScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupSchedule)(nil)).Elem()
}

func (o BackupScheduleMapOutput) ToBackupScheduleMapOutput() BackupScheduleMapOutput {
	return o
}

func (o BackupScheduleMapOutput) ToBackupScheduleMapOutputWithContext(ctx context.Context) BackupScheduleMapOutput {
	return o
}

func (o BackupScheduleMapOutput) MapIndex(k pulumi.StringInput) BackupScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupSchedule {
		return vs[0].(map[string]*BackupSchedule)[vs[1].(string)]
	}).(BackupScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupScheduleInput)(nil)).Elem(), &BackupSchedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupScheduleArrayInput)(nil)).Elem(), BackupScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupScheduleMapInput)(nil)).Elem(), BackupScheduleMap{})
	pulumi.RegisterOutputType(BackupScheduleOutput{})
	pulumi.RegisterOutputType(BackupScheduleArrayOutput{})
	pulumi.RegisterOutputType(BackupScheduleMapOutput{})
}
