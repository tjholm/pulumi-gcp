// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Service-Level Objective (SLO) describes the level of desired good
// service. It consists of a service-level indicator (SLI), a performance
// goal, and a period over which the objective is to be evaluated against
// that goal. The SLO can use SLIs defined in a number of different manners.
// Typical SLOs might include "99% of requests in each rolling week have
// latency below 200 milliseconds" or "99.5% of requests in each calendar
// month return successfully."
//
// To get more information about Slo, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services.serviceLevelObjectives)
// * How-to Guides
//   - [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
//   - [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
//
// ## Example Usage
// ### Monitoring Slo Appengine
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/monitoring"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := monitoring.GetAppEngineService(ctx, &monitoring.GetAppEngineServiceArgs{
//				ModuleId: "default",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = monitoring.NewSlo(ctx, "appengSlo", &monitoring.SloArgs{
//				Service:        *pulumi.String(_default.ServiceId),
//				SloId:          pulumi.String("ae-slo"),
//				DisplayName:    pulumi.String("Test SLO for App Engine"),
//				Goal:           pulumi.Float64(0.9),
//				CalendarPeriod: pulumi.String("DAY"),
//				BasicSli: &monitoring.SloBasicSliArgs{
//					Latency: &monitoring.SloBasicSliLatencyArgs{
//						Threshold: pulumi.String("1s"),
//					},
//				},
//				UserLabels: pulumi.StringMap{
//					"my_key":       pulumi.String("my_value"),
//					"my_other_key": pulumi.String("my_other_value"),
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
// ### Monitoring Slo Request Based
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/monitoring"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			customsrv, err := monitoring.NewCustomService(ctx, "customsrv", &monitoring.CustomServiceArgs{
//				ServiceId:   pulumi.String("custom-srv-request-slos"),
//				DisplayName: pulumi.String("My Custom Service"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = monitoring.NewSlo(ctx, "requestBasedSlo", &monitoring.SloArgs{
//				Service:           customsrv.ServiceId,
//				SloId:             pulumi.String("consumed-api-slo"),
//				DisplayName:       pulumi.String("Test SLO with request based SLI (good total ratio)"),
//				Goal:              pulumi.Float64(0.9),
//				RollingPeriodDays: pulumi.Int(30),
//				RequestBasedSli: &monitoring.SloRequestBasedSliArgs{
//					DistributionCut: &monitoring.SloRequestBasedSliDistributionCutArgs{
//						DistributionFilter: pulumi.String("metric.type=\"serviceruntime.googleapis.com/api/request_latencies\" resource.type=\"api\"  "),
//						Range: &monitoring.SloRequestBasedSliDistributionCutRangeArgs{
//							Max: pulumi.Float64(0.5),
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
// Slo can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Slo can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:monitoring/slo:Slo default {{name}}
//
// ```
type Slo struct {
	pulumi.CustomResourceState

	// Basic Service-Level Indicator (SLI) on a well-known service type.
	// Performance will be computed on the basis of pre-defined metrics.
	// SLIs are used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	BasicSli SloBasicSliPtrOutput `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current
	// <calendarPeriod>".
	// Possible values are: `DAY`, `WEEK`, `FORTNIGHT`, `MONTH`.
	CalendarPeriod pulumi.StringPtrOutput `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective
	// to be met. 0 < goal <= 0.999
	Goal pulumi.Float64Output `pulumi:"goal"`
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A request-based SLI defines a SLI for which atomic units of
	// service are counted directly.
	// A SLI describes a good service.
	// It is used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	RequestBasedSli SloRequestBasedSliPtrOutput `pulumi:"requestBasedSli"`
	// A rolling time period, semantically "in the past X days".
	// Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.IntPtrOutput `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	//
	// ***
	Service pulumi.StringOutput `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId pulumi.StringOutput `pulumi:"sloId"`
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels pulumi.StringMapOutput `pulumi:"userLabels"`
	// A windows-based SLI defines the criteria for time windows.
	// goodService is defined based off the count of these time windows
	// for which the provided service was of good quality.
	// A SLI describes a good service. It is used to measure and calculate
	// the quality of the Service's performance with respect to a single
	// aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	WindowsBasedSli SloWindowsBasedSliPtrOutput `pulumi:"windowsBasedSli"`
}

// NewSlo registers a new resource with the given unique name, arguments, and options.
func NewSlo(ctx *pulumi.Context,
	name string, args *SloArgs, opts ...pulumi.ResourceOption) (*Slo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Goal == nil {
		return nil, errors.New("invalid value for required argument 'Goal'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Slo
	err := ctx.RegisterResource("gcp:monitoring/slo:Slo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlo gets an existing Slo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SloState, opts ...pulumi.ResourceOption) (*Slo, error) {
	var resource Slo
	err := ctx.ReadResource("gcp:monitoring/slo:Slo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Slo resources.
type sloState struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type.
	// Performance will be computed on the basis of pre-defined metrics.
	// SLIs are used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	BasicSli *SloBasicSli `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current
	// <calendarPeriod>".
	// Possible values are: `DAY`, `WEEK`, `FORTNIGHT`, `MONTH`.
	CalendarPeriod *string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName *string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective
	// to be met. 0 < goal <= 0.999
	Goal *float64 `pulumi:"goal"`
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A request-based SLI defines a SLI for which atomic units of
	// service are counted directly.
	// A SLI describes a good service.
	// It is used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	RequestBasedSli *SloRequestBasedSli `pulumi:"requestBasedSli"`
	// A rolling time period, semantically "in the past X days".
	// Must be between 1 to 30 days, inclusive.
	RollingPeriodDays *int `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	//
	// ***
	Service *string `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId *string `pulumi:"sloId"`
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
	// A windows-based SLI defines the criteria for time windows.
	// goodService is defined based off the count of these time windows
	// for which the provided service was of good quality.
	// A SLI describes a good service. It is used to measure and calculate
	// the quality of the Service's performance with respect to a single
	// aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	WindowsBasedSli *SloWindowsBasedSli `pulumi:"windowsBasedSli"`
}

type SloState struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type.
	// Performance will be computed on the basis of pre-defined metrics.
	// SLIs are used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	BasicSli SloBasicSliPtrInput
	// A calendar period, semantically "since the start of the current
	// <calendarPeriod>".
	// Possible values are: `DAY`, `WEEK`, `FORTNIGHT`, `MONTH`.
	CalendarPeriod pulumi.StringPtrInput
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrInput
	// The fraction of service that must be good in order for this objective
	// to be met. 0 < goal <= 0.999
	Goal pulumi.Float64PtrInput
	// The full resource name for this service. The syntax is:
	// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A request-based SLI defines a SLI for which atomic units of
	// service are counted directly.
	// A SLI describes a good service.
	// It is used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	RequestBasedSli SloRequestBasedSliPtrInput
	// A rolling time period, semantically "in the past X days".
	// Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.IntPtrInput
	// ID of the service to which this SLO belongs.
	//
	// ***
	Service pulumi.StringPtrInput
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId pulumi.StringPtrInput
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels pulumi.StringMapInput
	// A windows-based SLI defines the criteria for time windows.
	// goodService is defined based off the count of these time windows
	// for which the provided service was of good quality.
	// A SLI describes a good service. It is used to measure and calculate
	// the quality of the Service's performance with respect to a single
	// aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	WindowsBasedSli SloWindowsBasedSliPtrInput
}

func (SloState) ElementType() reflect.Type {
	return reflect.TypeOf((*sloState)(nil)).Elem()
}

type sloArgs struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type.
	// Performance will be computed on the basis of pre-defined metrics.
	// SLIs are used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	BasicSli *SloBasicSli `pulumi:"basicSli"`
	// A calendar period, semantically "since the start of the current
	// <calendarPeriod>".
	// Possible values are: `DAY`, `WEEK`, `FORTNIGHT`, `MONTH`.
	CalendarPeriod *string `pulumi:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	DisplayName *string `pulumi:"displayName"`
	// The fraction of service that must be good in order for this objective
	// to be met. 0 < goal <= 0.999
	Goal float64 `pulumi:"goal"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A request-based SLI defines a SLI for which atomic units of
	// service are counted directly.
	// A SLI describes a good service.
	// It is used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	RequestBasedSli *SloRequestBasedSli `pulumi:"requestBasedSli"`
	// A rolling time period, semantically "in the past X days".
	// Must be between 1 to 30 days, inclusive.
	RollingPeriodDays *int `pulumi:"rollingPeriodDays"`
	// ID of the service to which this SLO belongs.
	//
	// ***
	Service string `pulumi:"service"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId *string `pulumi:"sloId"`
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
	// A windows-based SLI defines the criteria for time windows.
	// goodService is defined based off the count of these time windows
	// for which the provided service was of good quality.
	// A SLI describes a good service. It is used to measure and calculate
	// the quality of the Service's performance with respect to a single
	// aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	WindowsBasedSli *SloWindowsBasedSli `pulumi:"windowsBasedSli"`
}

// The set of arguments for constructing a Slo resource.
type SloArgs struct {
	// Basic Service-Level Indicator (SLI) on a well-known service type.
	// Performance will be computed on the basis of pre-defined metrics.
	// SLIs are used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	BasicSli SloBasicSliPtrInput
	// A calendar period, semantically "since the start of the current
	// <calendarPeriod>".
	// Possible values are: `DAY`, `WEEK`, `FORTNIGHT`, `MONTH`.
	CalendarPeriod pulumi.StringPtrInput
	// Name used for UI elements listing this SLO.
	DisplayName pulumi.StringPtrInput
	// The fraction of service that must be good in order for this objective
	// to be met. 0 < goal <= 0.999
	Goal pulumi.Float64Input
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A request-based SLI defines a SLI for which atomic units of
	// service are counted directly.
	// A SLI describes a good service.
	// It is used to measure and calculate the quality of the Service's
	// performance with respect to a single aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	RequestBasedSli SloRequestBasedSliPtrInput
	// A rolling time period, semantically "in the past X days".
	// Must be between 1 to 30 days, inclusive.
	RollingPeriodDays pulumi.IntPtrInput
	// ID of the service to which this SLO belongs.
	//
	// ***
	Service pulumi.StringInput
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	SloId pulumi.StringPtrInput
	// This field is intended to be used for organizing and identifying the AlertPolicy
	// objects.The field can contain up to 64 entries. Each key and value is limited
	// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
	// can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels pulumi.StringMapInput
	// A windows-based SLI defines the criteria for time windows.
	// goodService is defined based off the count of these time windows
	// for which the provided service was of good quality.
	// A SLI describes a good service. It is used to measure and calculate
	// the quality of the Service's performance with respect to a single
	// aspect of service quality.
	// Exactly one of the following must be set:
	// `basicSli`, `requestBasedSli`, `windowsBasedSli`
	// Structure is documented below.
	WindowsBasedSli SloWindowsBasedSliPtrInput
}

func (SloArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sloArgs)(nil)).Elem()
}

type SloInput interface {
	pulumi.Input

	ToSloOutput() SloOutput
	ToSloOutputWithContext(ctx context.Context) SloOutput
}

func (*Slo) ElementType() reflect.Type {
	return reflect.TypeOf((**Slo)(nil)).Elem()
}

func (i *Slo) ToSloOutput() SloOutput {
	return i.ToSloOutputWithContext(context.Background())
}

func (i *Slo) ToSloOutputWithContext(ctx context.Context) SloOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SloOutput)
}

// SloArrayInput is an input type that accepts SloArray and SloArrayOutput values.
// You can construct a concrete instance of `SloArrayInput` via:
//
//	SloArray{ SloArgs{...} }
type SloArrayInput interface {
	pulumi.Input

	ToSloArrayOutput() SloArrayOutput
	ToSloArrayOutputWithContext(context.Context) SloArrayOutput
}

type SloArray []SloInput

func (SloArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Slo)(nil)).Elem()
}

func (i SloArray) ToSloArrayOutput() SloArrayOutput {
	return i.ToSloArrayOutputWithContext(context.Background())
}

func (i SloArray) ToSloArrayOutputWithContext(ctx context.Context) SloArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SloArrayOutput)
}

// SloMapInput is an input type that accepts SloMap and SloMapOutput values.
// You can construct a concrete instance of `SloMapInput` via:
//
//	SloMap{ "key": SloArgs{...} }
type SloMapInput interface {
	pulumi.Input

	ToSloMapOutput() SloMapOutput
	ToSloMapOutputWithContext(context.Context) SloMapOutput
}

type SloMap map[string]SloInput

func (SloMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Slo)(nil)).Elem()
}

func (i SloMap) ToSloMapOutput() SloMapOutput {
	return i.ToSloMapOutputWithContext(context.Background())
}

func (i SloMap) ToSloMapOutputWithContext(ctx context.Context) SloMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SloMapOutput)
}

type SloOutput struct{ *pulumi.OutputState }

func (SloOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Slo)(nil)).Elem()
}

func (o SloOutput) ToSloOutput() SloOutput {
	return o
}

func (o SloOutput) ToSloOutputWithContext(ctx context.Context) SloOutput {
	return o
}

// Basic Service-Level Indicator (SLI) on a well-known service type.
// Performance will be computed on the basis of pre-defined metrics.
// SLIs are used to measure and calculate the quality of the Service's
// performance with respect to a single aspect of service quality.
// Exactly one of the following must be set:
// `basicSli`, `requestBasedSli`, `windowsBasedSli`
// Structure is documented below.
func (o SloOutput) BasicSli() SloBasicSliPtrOutput {
	return o.ApplyT(func(v *Slo) SloBasicSliPtrOutput { return v.BasicSli }).(SloBasicSliPtrOutput)
}

// A calendar period, semantically "since the start of the current
// <calendarPeriod>".
// Possible values are: `DAY`, `WEEK`, `FORTNIGHT`, `MONTH`.
func (o SloOutput) CalendarPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringPtrOutput { return v.CalendarPeriod }).(pulumi.StringPtrOutput)
}

// Name used for UI elements listing this SLO.
func (o SloOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The fraction of service that must be good in order for this objective
// to be met. 0 < goal <= 0.999
func (o SloOutput) Goal() pulumi.Float64Output {
	return o.ApplyT(func(v *Slo) pulumi.Float64Output { return v.Goal }).(pulumi.Float64Output)
}

// The full resource name for this service. The syntax is:
// projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
func (o SloOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o SloOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A request-based SLI defines a SLI for which atomic units of
// service are counted directly.
// A SLI describes a good service.
// It is used to measure and calculate the quality of the Service's
// performance with respect to a single aspect of service quality.
// Exactly one of the following must be set:
// `basicSli`, `requestBasedSli`, `windowsBasedSli`
// Structure is documented below.
func (o SloOutput) RequestBasedSli() SloRequestBasedSliPtrOutput {
	return o.ApplyT(func(v *Slo) SloRequestBasedSliPtrOutput { return v.RequestBasedSli }).(SloRequestBasedSliPtrOutput)
}

// A rolling time period, semantically "in the past X days".
// Must be between 1 to 30 days, inclusive.
func (o SloOutput) RollingPeriodDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Slo) pulumi.IntPtrOutput { return v.RollingPeriodDays }).(pulumi.IntPtrOutput)
}

// ID of the service to which this SLO belongs.
//
// ***
func (o SloOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
func (o SloOutput) SloId() pulumi.StringOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringOutput { return v.SloId }).(pulumi.StringOutput)
}

// This field is intended to be used for organizing and identifying the AlertPolicy
// objects.The field can contain up to 64 entries. Each key and value is limited
// to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values
// can contain only lowercase letters, numerals, underscores, and dashes. Keys
// must begin with a letter.
func (o SloOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringMapOutput { return v.UserLabels }).(pulumi.StringMapOutput)
}

// A windows-based SLI defines the criteria for time windows.
// goodService is defined based off the count of these time windows
// for which the provided service was of good quality.
// A SLI describes a good service. It is used to measure and calculate
// the quality of the Service's performance with respect to a single
// aspect of service quality.
// Exactly one of the following must be set:
// `basicSli`, `requestBasedSli`, `windowsBasedSli`
// Structure is documented below.
func (o SloOutput) WindowsBasedSli() SloWindowsBasedSliPtrOutput {
	return o.ApplyT(func(v *Slo) SloWindowsBasedSliPtrOutput { return v.WindowsBasedSli }).(SloWindowsBasedSliPtrOutput)
}

type SloArrayOutput struct{ *pulumi.OutputState }

func (SloArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Slo)(nil)).Elem()
}

func (o SloArrayOutput) ToSloArrayOutput() SloArrayOutput {
	return o
}

func (o SloArrayOutput) ToSloArrayOutputWithContext(ctx context.Context) SloArrayOutput {
	return o
}

func (o SloArrayOutput) Index(i pulumi.IntInput) SloOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Slo {
		return vs[0].([]*Slo)[vs[1].(int)]
	}).(SloOutput)
}

type SloMapOutput struct{ *pulumi.OutputState }

func (SloMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Slo)(nil)).Elem()
}

func (o SloMapOutput) ToSloMapOutput() SloMapOutput {
	return o
}

func (o SloMapOutput) ToSloMapOutputWithContext(ctx context.Context) SloMapOutput {
	return o
}

func (o SloMapOutput) MapIndex(k pulumi.StringInput) SloOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Slo {
		return vs[0].(map[string]*Slo)[vs[1].(string)]
	}).(SloOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SloInput)(nil)).Elem(), &Slo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SloArrayInput)(nil)).Elem(), SloArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SloMapInput)(nil)).Elem(), SloMap{})
	pulumi.RegisterOutputType(SloOutput{})
	pulumi.RegisterOutputType(SloArrayOutput{})
	pulumi.RegisterOutputType(SloMapOutput{})
}
