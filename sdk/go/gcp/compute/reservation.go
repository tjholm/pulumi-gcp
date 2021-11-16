// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a reservation resource. A reservation ensures that capacity is
// held in a specific zone even if the reserved VMs are not running.
//
// Reservations apply only to Compute Engine, Cloud Dataproc, and Google
// Kubernetes Engine VM usage.Reservations do not apply to `f1-micro` or
// `g1-small` machine types, preemptible VMs, sole tenant nodes, or other
// services not listed above
// like Cloud SQL and Dataflow.
//
// To get more information about Reservation, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/reservations)
// * How-to Guides
//     * [Reserving zonal resources](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources)
//
// ## Example Usage
// ### Reservation Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewReservation(ctx, "gceReservation", &compute.ReservationArgs{
// 			SpecificReservation: &compute.ReservationSpecificReservationArgs{
// 				Count: pulumi.Int(1),
// 				InstanceProperties: &compute.ReservationSpecificReservationInstancePropertiesArgs{
// 					MachineType:    pulumi.String("n2-standard-2"),
// 					MinCpuPlatform: pulumi.String("Intel Cascade Lake"),
// 				},
// 			},
// 			Zone: pulumi.String("us-central1-a"),
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
//  $ pulumi import gcp:compute/reservation:Reservation default projects/{{project}}/zones/{{zone}}/reservations/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/reservation:Reservation default {{project}}/{{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/reservation:Reservation default {{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/reservation:Reservation default {{name}}
// ```
type Reservation struct {
	pulumi.CustomResourceState

	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment pulumi.StringOutput `pulumi:"commitment"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Reservation for instances with specific machine shapes.
	// Structure is documented below.
	SpecificReservation ReservationSpecificReservationOutput `pulumi:"specificReservation"`
	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	SpecificReservationRequired pulumi.BoolPtrOutput `pulumi:"specificReservationRequired"`
	// The status of the reservation.
	Status pulumi.StringOutput `pulumi:"status"`
	// The zone where the reservation is made.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewReservation registers a new resource with the given unique name, arguments, and options.
func NewReservation(ctx *pulumi.Context,
	name string, args *ReservationArgs, opts ...pulumi.ResourceOption) (*Reservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SpecificReservation == nil {
		return nil, errors.New("invalid value for required argument 'SpecificReservation'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource Reservation
	err := ctx.RegisterResource("gcp:compute/reservation:Reservation", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:compute/reservation:Reservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Reservation resources.
type reservationState struct {
	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment *string `pulumi:"commitment"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Reservation for instances with specific machine shapes.
	// Structure is documented below.
	SpecificReservation *ReservationSpecificReservation `pulumi:"specificReservation"`
	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	SpecificReservationRequired *bool `pulumi:"specificReservationRequired"`
	// The status of the reservation.
	Status *string `pulumi:"status"`
	// The zone where the reservation is made.
	Zone *string `pulumi:"zone"`
}

type ReservationState struct {
	// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
	Commitment pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Reservation for instances with specific machine shapes.
	// Structure is documented below.
	SpecificReservation ReservationSpecificReservationPtrInput
	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	SpecificReservationRequired pulumi.BoolPtrInput
	// The status of the reservation.
	Status pulumi.StringPtrInput
	// The zone where the reservation is made.
	Zone pulumi.StringPtrInput
}

func (ReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationState)(nil)).Elem()
}

type reservationArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Reservation for instances with specific machine shapes.
	// Structure is documented below.
	SpecificReservation ReservationSpecificReservation `pulumi:"specificReservation"`
	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	SpecificReservationRequired *bool `pulumi:"specificReservationRequired"`
	// The zone where the reservation is made.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a Reservation resource.
type ReservationArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Reservation for instances with specific machine shapes.
	// Structure is documented below.
	SpecificReservation ReservationSpecificReservationInput
	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	SpecificReservationRequired pulumi.BoolPtrInput
	// The zone where the reservation is made.
	Zone pulumi.StringInput
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
