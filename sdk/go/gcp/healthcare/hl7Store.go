// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Hl7V2Store is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/hl7V2/STU3/)
// standard for Healthcare information exchange
//
// To get more information about Hl7V2Store, see:
//
// * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.hl7V2Stores)
// * How-to Guides
//     * [Creating a HL7v2 Store](https://cloud.google.com/healthcare/docs/how-tos/hl7v2)
//
// ## Example Usage
// ### Healthcare Hl7 V2 Store Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		topic, err := pubsub.NewTopic(ctx, "topic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		dataset, err := healthcare.NewDataset(ctx, "dataset", &healthcare.DatasetArgs{
// 			Location: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewHl7Store(ctx, "store", &healthcare.Hl7StoreArgs{
// 			Dataset: dataset.ID(),
// 			NotificationConfigs: healthcare.Hl7StoreNotificationConfigsArray{
// 				&healthcare.Hl7StoreNotificationConfigsArgs{
// 					PubsubTopic: topic.ID(),
// 				},
// 			},
// 			Labels: pulumi.StringMap{
// 				"label1": pulumi.String("labelvalue1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Healthcare Hl7 V2 Store Parser Config
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dataset, err := healthcare.NewDataset(ctx, "dataset", &healthcare.DatasetArgs{
// 			Location: pulumi.String("us-central1"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewHl7Store(ctx, "store", &healthcare.Hl7StoreArgs{
// 			Dataset: dataset.ID(),
// 			ParserConfig: &healthcare.Hl7StoreParserConfigArgs{
// 				AllowNullHeader:   pulumi.Bool(false),
// 				SegmentTerminator: pulumi.String("Jw=="),
// 				Schema:            pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"schemas\": [{\n", "    \"messageSchemaConfigs\": {\n", "      \"ADT_A01\": {\n", "        \"name\": \"ADT_A01\",\n", "        \"minOccurs\": 1,\n", "        \"maxOccurs\": 1,\n", "        \"members\": [{\n", "            \"segment\": {\n", "              \"type\": \"MSH\",\n", "              \"minOccurs\": 1,\n", "              \"maxOccurs\": 1\n", "            }\n", "          },\n", "          {\n", "            \"segment\": {\n", "              \"type\": \"EVN\",\n", "              \"minOccurs\": 1,\n", "              \"maxOccurs\": 1\n", "            }\n", "          },\n", "          {\n", "            \"segment\": {\n", "              \"type\": \"PID\",\n", "              \"minOccurs\": 1,\n", "              \"maxOccurs\": 1\n", "            }\n", "          },\n", "          {\n", "            \"segment\": {\n", "              \"type\": \"ZPD\",\n", "              \"minOccurs\": 1,\n", "              \"maxOccurs\": 1\n", "            }\n", "          },\n", "          {\n", "            \"segment\": {\n", "              \"type\": \"OBX\"\n", "            }\n", "          },\n", "          {\n", "            \"group\": {\n", "              \"name\": \"PROCEDURE\",\n", "              \"members\": [{\n", "                  \"segment\": {\n", "                    \"type\": \"PR1\",\n", "                    \"minOccurs\": 1,\n", "                    \"maxOccurs\": 1\n", "                  }\n", "                },\n", "                {\n", "                  \"segment\": {\n", "                    \"type\": \"ROL\"\n", "                  }\n", "                }\n", "              ]\n", "            }\n", "          },\n", "          {\n", "            \"segment\": {\n", "              \"type\": \"PDA\",\n", "              \"maxOccurs\": 1\n", "            }\n", "          }\n", "        ]\n", "      }\n", "    }\n", "  }],\n", "  \"types\": [{\n", "    \"type\": [{\n", "        \"name\": \"ZPD\",\n", "        \"primitive\": \"VARIES\"\n", "      }\n", "\n", "    ]\n", "  }],\n", "  \"ignoreMinOccurs\": true\n", "}\n")),
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Healthcare Hl7 V2 Store Unschematized
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dataset, err := healthcare.NewDataset(ctx, "dataset", &healthcare.DatasetArgs{
// 			Location: pulumi.String("us-central1"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewHl7Store(ctx, "store", &healthcare.Hl7StoreArgs{
// 			Dataset: dataset.ID(),
// 			ParserConfig: &healthcare.Hl7StoreParserConfigArgs{
// 				AllowNullHeader:   pulumi.Bool(false),
// 				SegmentTerminator: pulumi.String("Jw=="),
// 				Version:           pulumi.String("V2"),
// 			},
// 		}, pulumi.Provider(google_beta))
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
// Hl7V2Store can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:healthcare/hl7Store:Hl7Store default {{dataset}}/hl7V2Stores/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:healthcare/hl7Store:Hl7Store default {{dataset}}/{{name}}
// ```
type Hl7Store struct {
	pulumi.CustomResourceState

	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize HL7v2 stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for the Hl7V2Store.
	// ** Changing this property may recreate the Hl7v2 store (removing all data) **
	Name pulumi.StringOutput `pulumi:"name"`
	// -
	// (Optional, Deprecated)
	// A nested object resource
	// Structure is documented below.
	//
	// Deprecated: This field has been replaced by notificationConfigs
	NotificationConfig Hl7StoreNotificationConfigPtrOutput `pulumi:"notificationConfig"`
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a
	// message (both Ingest & Create) on the corresponding notification destination. Only the message name
	// is sent as part of the notification. Supplied by the client.
	// Structure is documented below.
	NotificationConfigs Hl7StoreNotificationConfigsArrayOutput `pulumi:"notificationConfigs"`
	// A nested object resource
	// Structure is documented below.
	ParserConfig Hl7StoreParserConfigOutput `pulumi:"parserConfig"`
	// The fully qualified name of this dataset
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewHl7Store registers a new resource with the given unique name, arguments, and options.
func NewHl7Store(ctx *pulumi.Context,
	name string, args *Hl7StoreArgs, opts ...pulumi.ResourceOption) (*Hl7Store, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	var resource Hl7Store
	err := ctx.RegisterResource("gcp:healthcare/hl7Store:Hl7Store", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHl7Store gets an existing Hl7Store resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHl7Store(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Hl7StoreState, opts ...pulumi.ResourceOption) (*Hl7Store, error) {
	var resource Hl7Store
	err := ctx.ReadResource("gcp:healthcare/hl7Store:Hl7Store", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hl7Store resources.
type hl7StoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset *string `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize HL7v2 stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the Hl7V2Store.
	// ** Changing this property may recreate the Hl7v2 store (removing all data) **
	Name *string `pulumi:"name"`
	// -
	// (Optional, Deprecated)
	// A nested object resource
	// Structure is documented below.
	//
	// Deprecated: This field has been replaced by notificationConfigs
	NotificationConfig *Hl7StoreNotificationConfig `pulumi:"notificationConfig"`
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a
	// message (both Ingest & Create) on the corresponding notification destination. Only the message name
	// is sent as part of the notification. Supplied by the client.
	// Structure is documented below.
	NotificationConfigs []Hl7StoreNotificationConfigs `pulumi:"notificationConfigs"`
	// A nested object resource
	// Structure is documented below.
	ParserConfig *Hl7StoreParserConfig `pulumi:"parserConfig"`
	// The fully qualified name of this dataset
	SelfLink *string `pulumi:"selfLink"`
}

type Hl7StoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize HL7v2 stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The resource name for the Hl7V2Store.
	// ** Changing this property may recreate the Hl7v2 store (removing all data) **
	Name pulumi.StringPtrInput
	// -
	// (Optional, Deprecated)
	// A nested object resource
	// Structure is documented below.
	//
	// Deprecated: This field has been replaced by notificationConfigs
	NotificationConfig Hl7StoreNotificationConfigPtrInput
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a
	// message (both Ingest & Create) on the corresponding notification destination. Only the message name
	// is sent as part of the notification. Supplied by the client.
	// Structure is documented below.
	NotificationConfigs Hl7StoreNotificationConfigsArrayInput
	// A nested object resource
	// Structure is documented below.
	ParserConfig Hl7StoreParserConfigPtrInput
	// The fully qualified name of this dataset
	SelfLink pulumi.StringPtrInput
}

func (Hl7StoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7StoreState)(nil)).Elem()
}

type hl7StoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset string `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize HL7v2 stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the Hl7V2Store.
	// ** Changing this property may recreate the Hl7v2 store (removing all data) **
	Name *string `pulumi:"name"`
	// -
	// (Optional, Deprecated)
	// A nested object resource
	// Structure is documented below.
	//
	// Deprecated: This field has been replaced by notificationConfigs
	NotificationConfig *Hl7StoreNotificationConfig `pulumi:"notificationConfig"`
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a
	// message (both Ingest & Create) on the corresponding notification destination. Only the message name
	// is sent as part of the notification. Supplied by the client.
	// Structure is documented below.
	NotificationConfigs []Hl7StoreNotificationConfigs `pulumi:"notificationConfigs"`
	// A nested object resource
	// Structure is documented below.
	ParserConfig *Hl7StoreParserConfig `pulumi:"parserConfig"`
}

// The set of arguments for constructing a Hl7Store resource.
type Hl7StoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringInput
	// User-supplied key-value pairs used to organize HL7v2 stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The resource name for the Hl7V2Store.
	// ** Changing this property may recreate the Hl7v2 store (removing all data) **
	Name pulumi.StringPtrInput
	// -
	// (Optional, Deprecated)
	// A nested object resource
	// Structure is documented below.
	//
	// Deprecated: This field has been replaced by notificationConfigs
	NotificationConfig Hl7StoreNotificationConfigPtrInput
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a
	// message (both Ingest & Create) on the corresponding notification destination. Only the message name
	// is sent as part of the notification. Supplied by the client.
	// Structure is documented below.
	NotificationConfigs Hl7StoreNotificationConfigsArrayInput
	// A nested object resource
	// Structure is documented below.
	ParserConfig Hl7StoreParserConfigPtrInput
}

func (Hl7StoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7StoreArgs)(nil)).Elem()
}

type Hl7StoreInput interface {
	pulumi.Input

	ToHl7StoreOutput() Hl7StoreOutput
	ToHl7StoreOutputWithContext(ctx context.Context) Hl7StoreOutput
}

func (*Hl7Store) ElementType() reflect.Type {
	return reflect.TypeOf((*Hl7Store)(nil))
}

func (i *Hl7Store) ToHl7StoreOutput() Hl7StoreOutput {
	return i.ToHl7StoreOutputWithContext(context.Background())
}

func (i *Hl7Store) ToHl7StoreOutputWithContext(ctx context.Context) Hl7StoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Hl7StoreOutput)
}

func (i *Hl7Store) ToHl7StorePtrOutput() Hl7StorePtrOutput {
	return i.ToHl7StorePtrOutputWithContext(context.Background())
}

func (i *Hl7Store) ToHl7StorePtrOutputWithContext(ctx context.Context) Hl7StorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Hl7StorePtrOutput)
}

type Hl7StorePtrInput interface {
	pulumi.Input

	ToHl7StorePtrOutput() Hl7StorePtrOutput
	ToHl7StorePtrOutputWithContext(ctx context.Context) Hl7StorePtrOutput
}

type hl7StorePtrType Hl7StoreArgs

func (*hl7StorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Hl7Store)(nil))
}

func (i *hl7StorePtrType) ToHl7StorePtrOutput() Hl7StorePtrOutput {
	return i.ToHl7StorePtrOutputWithContext(context.Background())
}

func (i *hl7StorePtrType) ToHl7StorePtrOutputWithContext(ctx context.Context) Hl7StorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Hl7StorePtrOutput)
}

// Hl7StoreArrayInput is an input type that accepts Hl7StoreArray and Hl7StoreArrayOutput values.
// You can construct a concrete instance of `Hl7StoreArrayInput` via:
//
//          Hl7StoreArray{ Hl7StoreArgs{...} }
type Hl7StoreArrayInput interface {
	pulumi.Input

	ToHl7StoreArrayOutput() Hl7StoreArrayOutput
	ToHl7StoreArrayOutputWithContext(context.Context) Hl7StoreArrayOutput
}

type Hl7StoreArray []Hl7StoreInput

func (Hl7StoreArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Hl7Store)(nil))
}

func (i Hl7StoreArray) ToHl7StoreArrayOutput() Hl7StoreArrayOutput {
	return i.ToHl7StoreArrayOutputWithContext(context.Background())
}

func (i Hl7StoreArray) ToHl7StoreArrayOutputWithContext(ctx context.Context) Hl7StoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Hl7StoreArrayOutput)
}

// Hl7StoreMapInput is an input type that accepts Hl7StoreMap and Hl7StoreMapOutput values.
// You can construct a concrete instance of `Hl7StoreMapInput` via:
//
//          Hl7StoreMap{ "key": Hl7StoreArgs{...} }
type Hl7StoreMapInput interface {
	pulumi.Input

	ToHl7StoreMapOutput() Hl7StoreMapOutput
	ToHl7StoreMapOutputWithContext(context.Context) Hl7StoreMapOutput
}

type Hl7StoreMap map[string]Hl7StoreInput

func (Hl7StoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Hl7Store)(nil))
}

func (i Hl7StoreMap) ToHl7StoreMapOutput() Hl7StoreMapOutput {
	return i.ToHl7StoreMapOutputWithContext(context.Background())
}

func (i Hl7StoreMap) ToHl7StoreMapOutputWithContext(ctx context.Context) Hl7StoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Hl7StoreMapOutput)
}

type Hl7StoreOutput struct {
	*pulumi.OutputState
}

func (Hl7StoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Hl7Store)(nil))
}

func (o Hl7StoreOutput) ToHl7StoreOutput() Hl7StoreOutput {
	return o
}

func (o Hl7StoreOutput) ToHl7StoreOutputWithContext(ctx context.Context) Hl7StoreOutput {
	return o
}

func (o Hl7StoreOutput) ToHl7StorePtrOutput() Hl7StorePtrOutput {
	return o.ToHl7StorePtrOutputWithContext(context.Background())
}

func (o Hl7StoreOutput) ToHl7StorePtrOutputWithContext(ctx context.Context) Hl7StorePtrOutput {
	return o.ApplyT(func(v Hl7Store) *Hl7Store {
		return &v
	}).(Hl7StorePtrOutput)
}

type Hl7StorePtrOutput struct {
	*pulumi.OutputState
}

func (Hl7StorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hl7Store)(nil))
}

func (o Hl7StorePtrOutput) ToHl7StorePtrOutput() Hl7StorePtrOutput {
	return o
}

func (o Hl7StorePtrOutput) ToHl7StorePtrOutputWithContext(ctx context.Context) Hl7StorePtrOutput {
	return o
}

type Hl7StoreArrayOutput struct{ *pulumi.OutputState }

func (Hl7StoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Hl7Store)(nil))
}

func (o Hl7StoreArrayOutput) ToHl7StoreArrayOutput() Hl7StoreArrayOutput {
	return o
}

func (o Hl7StoreArrayOutput) ToHl7StoreArrayOutputWithContext(ctx context.Context) Hl7StoreArrayOutput {
	return o
}

func (o Hl7StoreArrayOutput) Index(i pulumi.IntInput) Hl7StoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Hl7Store {
		return vs[0].([]Hl7Store)[vs[1].(int)]
	}).(Hl7StoreOutput)
}

type Hl7StoreMapOutput struct{ *pulumi.OutputState }

func (Hl7StoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Hl7Store)(nil))
}

func (o Hl7StoreMapOutput) ToHl7StoreMapOutput() Hl7StoreMapOutput {
	return o
}

func (o Hl7StoreMapOutput) ToHl7StoreMapOutputWithContext(ctx context.Context) Hl7StoreMapOutput {
	return o
}

func (o Hl7StoreMapOutput) MapIndex(k pulumi.StringInput) Hl7StoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Hl7Store {
		return vs[0].(map[string]Hl7Store)[vs[1].(string)]
	}).(Hl7StoreOutput)
}

func init() {
	pulumi.RegisterOutputType(Hl7StoreOutput{})
	pulumi.RegisterOutputType(Hl7StorePtrOutput{})
	pulumi.RegisterOutputType(Hl7StoreArrayOutput{})
	pulumi.RegisterOutputType(Hl7StoreMapOutput{})
}
