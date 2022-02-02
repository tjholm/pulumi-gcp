// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// DNS record sets can be imported using either of these accepted formats
//
// ```sh
//  $ pulumi import gcp:dns/recordSet:RecordSet frontend projects/{{project}}/managedZones/{{zone}}/rrsets/{{name}}/{{type}}
// ```
//
// ```sh
//  $ pulumi import gcp:dns/recordSet:RecordSet frontend {{project}}/{{zone}}/{{name}}/{{type}}
// ```
//
// ```sh
//  $ pulumi import gcp:dns/recordSet:RecordSet frontend {{zone}}/{{name}}/{{type}}
// ```
//
//  NoteThe record name must include the trailing dot at the end.
type RecordSet struct {
	pulumi.CustomResourceState

	// The name of the zone in which this record set will
	// reside.
	ManagedZone pulumi.StringOutput `pulumi:"managedZone"`
	// The DNS name this record set will apply to.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
	// data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
	// record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
	// string (e.g. "first255characters\"\"morecharacters").
	Rrdatas pulumi.StringArrayOutput `pulumi:"rrdatas"`
	// The time-to-live of this record set (seconds).
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The DNS record set type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRecordSet registers a new resource with the given unique name, arguments, and options.
func NewRecordSet(ctx *pulumi.Context,
	name string, args *RecordSetArgs, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedZone == nil {
		return nil, errors.New("invalid value for required argument 'ManagedZone'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Rrdatas == nil {
		return nil, errors.New("invalid value for required argument 'Rrdatas'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource RecordSet
	err := ctx.RegisterResource("gcp:dns/recordSet:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordSet gets an existing RecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("gcp:dns/recordSet:RecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordSet resources.
type recordSetState struct {
	// The name of the zone in which this record set will
	// reside.
	ManagedZone *string `pulumi:"managedZone"`
	// The DNS name this record set will apply to.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
	// data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
	// record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
	// string (e.g. "first255characters\"\"morecharacters").
	Rrdatas []string `pulumi:"rrdatas"`
	// The time-to-live of this record set (seconds).
	Ttl *int `pulumi:"ttl"`
	// The DNS record set type.
	Type *string `pulumi:"type"`
}

type RecordSetState struct {
	// The name of the zone in which this record set will
	// reside.
	ManagedZone pulumi.StringPtrInput
	// The DNS name this record set will apply to.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
	// data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
	// record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
	// string (e.g. "first255characters\"\"morecharacters").
	Rrdatas pulumi.StringArrayInput
	// The time-to-live of this record set (seconds).
	Ttl pulumi.IntPtrInput
	// The DNS record set type.
	Type pulumi.StringPtrInput
}

func (RecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetState)(nil)).Elem()
}

type recordSetArgs struct {
	// The name of the zone in which this record set will
	// reside.
	ManagedZone string `pulumi:"managedZone"`
	// The DNS name this record set will apply to.
	Name string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
	// data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
	// record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
	// string (e.g. "first255characters\"\"morecharacters").
	Rrdatas []string `pulumi:"rrdatas"`
	// The time-to-live of this record set (seconds).
	Ttl *int `pulumi:"ttl"`
	// The DNS record set type.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a RecordSet resource.
type RecordSetArgs struct {
	// The name of the zone in which this record set will
	// reside.
	ManagedZone pulumi.StringInput
	// The DNS name this record set will apply to.
	Name pulumi.StringInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string
	// data contains spaces, add surrounding \" if you don't want your string to get split on spaces. To specify a single
	// record value longer than 255 characters such as a TXT record for DKIM, add \"\" inside the Terraform configuration
	// string (e.g. "first255characters\"\"morecharacters").
	Rrdatas pulumi.StringArrayInput
	// The time-to-live of this record set (seconds).
	Ttl pulumi.IntPtrInput
	// The DNS record set type.
	Type pulumi.StringInput
}

func (RecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetArgs)(nil)).Elem()
}

type RecordSetInput interface {
	pulumi.Input

	ToRecordSetOutput() RecordSetOutput
	ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput
}

func (*RecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil)).Elem()
}

func (i *RecordSet) ToRecordSetOutput() RecordSetOutput {
	return i.ToRecordSetOutputWithContext(context.Background())
}

func (i *RecordSet) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetOutput)
}

// RecordSetArrayInput is an input type that accepts RecordSetArray and RecordSetArrayOutput values.
// You can construct a concrete instance of `RecordSetArrayInput` via:
//
//          RecordSetArray{ RecordSetArgs{...} }
type RecordSetArrayInput interface {
	pulumi.Input

	ToRecordSetArrayOutput() RecordSetArrayOutput
	ToRecordSetArrayOutputWithContext(context.Context) RecordSetArrayOutput
}

type RecordSetArray []RecordSetInput

func (RecordSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RecordSet)(nil)).Elem()
}

func (i RecordSetArray) ToRecordSetArrayOutput() RecordSetArrayOutput {
	return i.ToRecordSetArrayOutputWithContext(context.Background())
}

func (i RecordSetArray) ToRecordSetArrayOutputWithContext(ctx context.Context) RecordSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetArrayOutput)
}

// RecordSetMapInput is an input type that accepts RecordSetMap and RecordSetMapOutput values.
// You can construct a concrete instance of `RecordSetMapInput` via:
//
//          RecordSetMap{ "key": RecordSetArgs{...} }
type RecordSetMapInput interface {
	pulumi.Input

	ToRecordSetMapOutput() RecordSetMapOutput
	ToRecordSetMapOutputWithContext(context.Context) RecordSetMapOutput
}

type RecordSetMap map[string]RecordSetInput

func (RecordSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RecordSet)(nil)).Elem()
}

func (i RecordSetMap) ToRecordSetMapOutput() RecordSetMapOutput {
	return i.ToRecordSetMapOutputWithContext(context.Background())
}

func (i RecordSetMap) ToRecordSetMapOutputWithContext(ctx context.Context) RecordSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetMapOutput)
}

type RecordSetOutput struct{ *pulumi.OutputState }

func (RecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil)).Elem()
}

func (o RecordSetOutput) ToRecordSetOutput() RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return o
}

type RecordSetArrayOutput struct{ *pulumi.OutputState }

func (RecordSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RecordSet)(nil)).Elem()
}

func (o RecordSetArrayOutput) ToRecordSetArrayOutput() RecordSetArrayOutput {
	return o
}

func (o RecordSetArrayOutput) ToRecordSetArrayOutputWithContext(ctx context.Context) RecordSetArrayOutput {
	return o
}

func (o RecordSetArrayOutput) Index(i pulumi.IntInput) RecordSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RecordSet {
		return vs[0].([]*RecordSet)[vs[1].(int)]
	}).(RecordSetOutput)
}

type RecordSetMapOutput struct{ *pulumi.OutputState }

func (RecordSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RecordSet)(nil)).Elem()
}

func (o RecordSetMapOutput) ToRecordSetMapOutput() RecordSetMapOutput {
	return o
}

func (o RecordSetMapOutput) ToRecordSetMapOutputWithContext(ctx context.Context) RecordSetMapOutput {
	return o
}

func (o RecordSetMapOutput) MapIndex(k pulumi.StringInput) RecordSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RecordSet {
		return vs[0].(map[string]*RecordSet)[vs[1].(string)]
	}).(RecordSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecordSetInput)(nil)).Elem(), &RecordSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecordSetArrayInput)(nil)).Elem(), RecordSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecordSetMapInput)(nil)).Elem(), RecordSetMap{})
	pulumi.RegisterOutputType(RecordSetOutput{})
	pulumi.RegisterOutputType(RecordSetArrayOutput{})
	pulumi.RegisterOutputType(RecordSetMapOutput{})
}
