// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A reservation is a mechanism used to guarantee BigQuery slots to users.
//
// To get more information about Reservation, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/reservations/rest/v1beta1/projects.locations.reservations/create)
// * How-to Guides
//     * [Introduction to Reservations](https://cloud.google.com/bigquery/docs/reservations-intro)
//
// ## Example Usage
// ### Bigquery Reservation Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigquery.NewReservation(ctx, "reservation", &bigquery.ReservationArgs{
// 			IgnoreIdleSlots: pulumi.Bool(false),
// 			Location:        pulumi.String("asia-northeast1"),
// 			SlotCapacity:    pulumi.Int(0),
// 		})
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
// Reservation can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:bigquery/reservation:Reservation default projects/{{project}}/locations/{{location}}/reservations/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/reservation:Reservation default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/reservation:Reservation default {{location}}/{{name}}
// ```
type Reservation struct {
	pulumi.CustomResourceState

	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrOutput `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.IntOutput `pulumi:"slotCapacity"`
}

// NewReservation registers a new resource with the given unique name, arguments, and options.
func NewReservation(ctx *pulumi.Context,
	name string, args *ReservationArgs, opts ...pulumi.ResourceOption) (*Reservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SlotCapacity == nil {
		return nil, errors.New("invalid value for required argument 'SlotCapacity'")
	}
	var resource Reservation
	err := ctx.RegisterResource("gcp:bigquery/reservation:Reservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReservation gets an existing Reservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReservationState, opts ...pulumi.ResourceOption) (*Reservation, error) {
	var resource Reservation
	err := ctx.ReadResource("gcp:bigquery/reservation:Reservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Reservation resources.
type reservationState struct {
	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	IgnoreIdleSlots *bool `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location *string `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity *int `pulumi:"slotCapacity"`
}

type ReservationState struct {
	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrInput
	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location pulumi.StringPtrInput
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.IntPtrInput
}

func (ReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationState)(nil)).Elem()
}

type reservationArgs struct {
	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	IgnoreIdleSlots *bool `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location *string `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity int `pulumi:"slotCapacity"`
}

// The set of arguments for constructing a Reservation resource.
type ReservationArgs struct {
	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrInput
	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location pulumi.StringPtrInput
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.IntInput
}

func (ReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationArgs)(nil)).Elem()
}

type ReservationInput interface {
	pulumi.Input

	ToReservationOutput() ReservationOutput
	ToReservationOutputWithContext(ctx context.Context) ReservationOutput
}

func (*Reservation) ElementType() reflect.Type {
	return reflect.TypeOf((*Reservation)(nil))
}

func (i *Reservation) ToReservationOutput() ReservationOutput {
	return i.ToReservationOutputWithContext(context.Background())
}

func (i *Reservation) ToReservationOutputWithContext(ctx context.Context) ReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationOutput)
}

func (i *Reservation) ToReservationPtrOutput() ReservationPtrOutput {
	return i.ToReservationPtrOutputWithContext(context.Background())
}

func (i *Reservation) ToReservationPtrOutputWithContext(ctx context.Context) ReservationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationPtrOutput)
}

type ReservationPtrInput interface {
	pulumi.Input

	ToReservationPtrOutput() ReservationPtrOutput
	ToReservationPtrOutputWithContext(ctx context.Context) ReservationPtrOutput
}

type reservationPtrType ReservationArgs

func (*reservationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Reservation)(nil))
}

func (i *reservationPtrType) ToReservationPtrOutput() ReservationPtrOutput {
	return i.ToReservationPtrOutputWithContext(context.Background())
}

func (i *reservationPtrType) ToReservationPtrOutputWithContext(ctx context.Context) ReservationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationPtrOutput)
}

// ReservationArrayInput is an input type that accepts ReservationArray and ReservationArrayOutput values.
// You can construct a concrete instance of `ReservationArrayInput` via:
//
//          ReservationArray{ ReservationArgs{...} }
type ReservationArrayInput interface {
	pulumi.Input

	ToReservationArrayOutput() ReservationArrayOutput
	ToReservationArrayOutputWithContext(context.Context) ReservationArrayOutput
}

type ReservationArray []ReservationInput

func (ReservationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Reservation)(nil)).Elem()
}

func (i ReservationArray) ToReservationArrayOutput() ReservationArrayOutput {
	return i.ToReservationArrayOutputWithContext(context.Background())
}

func (i ReservationArray) ToReservationArrayOutputWithContext(ctx context.Context) ReservationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationArrayOutput)
}

// ReservationMapInput is an input type that accepts ReservationMap and ReservationMapOutput values.
// You can construct a concrete instance of `ReservationMapInput` via:
//
//          ReservationMap{ "key": ReservationArgs{...} }
type ReservationMapInput interface {
	pulumi.Input

	ToReservationMapOutput() ReservationMapOutput
	ToReservationMapOutputWithContext(context.Context) ReservationMapOutput
}

type ReservationMap map[string]ReservationInput

func (ReservationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Reservation)(nil)).Elem()
}

func (i ReservationMap) ToReservationMapOutput() ReservationMapOutput {
	return i.ToReservationMapOutputWithContext(context.Background())
}

func (i ReservationMap) ToReservationMapOutputWithContext(ctx context.Context) ReservationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservationMapOutput)
}

type ReservationOutput struct{ *pulumi.OutputState }

func (ReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Reservation)(nil))
}

func (o ReservationOutput) ToReservationOutput() ReservationOutput {
	return o
}

func (o ReservationOutput) ToReservationOutputWithContext(ctx context.Context) ReservationOutput {
	return o
}

func (o ReservationOutput) ToReservationPtrOutput() ReservationPtrOutput {
	return o.ToReservationPtrOutputWithContext(context.Background())
}

func (o ReservationOutput) ToReservationPtrOutputWithContext(ctx context.Context) ReservationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Reservation) *Reservation {
		return &v
	}).(ReservationPtrOutput)
}

type ReservationPtrOutput struct{ *pulumi.OutputState }

func (ReservationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Reservation)(nil))
}

func (o ReservationPtrOutput) ToReservationPtrOutput() ReservationPtrOutput {
	return o
}

func (o ReservationPtrOutput) ToReservationPtrOutputWithContext(ctx context.Context) ReservationPtrOutput {
	return o
}

func (o ReservationPtrOutput) Elem() ReservationOutput {
	return o.ApplyT(func(v *Reservation) Reservation {
		if v != nil {
			return *v
		}
		var ret Reservation
		return ret
	}).(ReservationOutput)
}

type ReservationArrayOutput struct{ *pulumi.OutputState }

func (ReservationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Reservation)(nil))
}

func (o ReservationArrayOutput) ToReservationArrayOutput() ReservationArrayOutput {
	return o
}

func (o ReservationArrayOutput) ToReservationArrayOutputWithContext(ctx context.Context) ReservationArrayOutput {
	return o
}

func (o ReservationArrayOutput) Index(i pulumi.IntInput) ReservationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Reservation {
		return vs[0].([]Reservation)[vs[1].(int)]
	}).(ReservationOutput)
}

type ReservationMapOutput struct{ *pulumi.OutputState }

func (ReservationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Reservation)(nil))
}

func (o ReservationMapOutput) ToReservationMapOutput() ReservationMapOutput {
	return o
}

func (o ReservationMapOutput) ToReservationMapOutputWithContext(ctx context.Context) ReservationMapOutput {
	return o
}

func (o ReservationMapOutput) MapIndex(k pulumi.StringInput) ReservationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Reservation {
		return vs[0].(map[string]Reservation)[vs[1].(string)]
	}).(ReservationOutput)
}

func init() {
	pulumi.RegisterOutputType(ReservationOutput{})
	pulumi.RegisterOutputType(ReservationPtrOutput{})
	pulumi.RegisterOutputType(ReservationArrayOutput{})
	pulumi.RegisterOutputType(ReservationMapOutput{})
}
