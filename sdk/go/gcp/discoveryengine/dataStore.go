// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package discoveryengine

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data store is a collection of websites and documents used to find answers for
// end-user's questions in Discovery Engine (a.k.a. Vertex AI Search and
// Conversation).
//
// To get more information about DataStore, see:
//
// * [API documentation](https://cloud.google.com/generative-ai-app-builder/docs/reference/rest/v1/projects.locations.collections.dataStores)
// * How-to Guides
//   - [Create a search data store](https://cloud.google.com/generative-ai-app-builder/docs/create-data-store-es)
//
// ## Example Usage
// ### Discoveryengine Datastore Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/discoveryengine"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := discoveryengine.NewDataStore(ctx, "basic", &discoveryengine.DataStoreArgs{
//				ContentConfig:            pulumi.String("NO_CONTENT"),
//				CreateAdvancedSiteSearch: pulumi.Bool(false),
//				DataStoreId:              pulumi.String("data-store-id"),
//				DisplayName:              pulumi.String("tf-test-structured-datastore"),
//				IndustryVertical:         pulumi.String("GENERIC"),
//				Location:                 pulumi.String("global"),
//				SolutionTypes: pulumi.StringArray{
//					pulumi.String("SOLUTION_TYPE_SEARCH"),
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
// DataStore can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}` * `{{project}}/{{location}}/{{data_store_id}}` * `{{location}}/{{data_store_id}}` When using the `pulumi import` command, DataStore can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:discoveryengine/dataStore:DataStore default projects/{{project}}/locations/{{location}}/collections/default_collection/dataStores/{{data_store_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:discoveryengine/dataStore:DataStore default {{project}}/{{location}}/{{data_store_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:discoveryengine/dataStore:DataStore default {{location}}/{{data_store_id}}
//
// ```
type DataStore struct {
	pulumi.CustomResourceState

	// The content config of the data store.
	// Possible values are: `NO_CONTENT`, `CONTENT_REQUIRED`, `PUBLIC_WEBSITE`.
	ContentConfig pulumi.StringOutput `pulumi:"contentConfig"`
	// If true, an advanced data store for site search will be created. If the
	// data store is not configured as site search (GENERIC vertical and
	// PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch pulumi.BoolPtrOutput `pulumi:"createAdvancedSiteSearch"`
	// Timestamp when the DataStore was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The unique id of the data store.
	//
	// ***
	DataStoreId pulumi.StringOutput `pulumi:"dataStoreId"`
	// The id of the default Schema associated with this data store.
	DefaultSchemaId pulumi.StringOutput `pulumi:"defaultSchemaId"`
	// The display name of the data store. This field must be a UTF-8 encoded
	// string with a length limit of 128 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The industry vertical that the data store registers.
	// Possible values are: `GENERIC`, `MEDIA`.
	IndustryVertical pulumi.StringOutput `pulumi:"industryVertical"`
	// The geographic location where the data store should reside. The value can
	// only be one of "global", "us" and "eu".
	Location pulumi.StringOutput `pulumi:"location"`
	// The unique full resource name of the data store. Values are of the format
	// `projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}`.
	// This field must be a UTF-8 encoded string with a length limit of 1024
	// characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The solutions that the data store enrolls.
	// Each value may be one of: `SOLUTION_TYPE_RECOMMENDATION`, `SOLUTION_TYPE_SEARCH`, `SOLUTION_TYPE_CHAT`.
	SolutionTypes pulumi.StringArrayOutput `pulumi:"solutionTypes"`
}

// NewDataStore registers a new resource with the given unique name, arguments, and options.
func NewDataStore(ctx *pulumi.Context,
	name string, args *DataStoreArgs, opts ...pulumi.ResourceOption) (*DataStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentConfig == nil {
		return nil, errors.New("invalid value for required argument 'ContentConfig'")
	}
	if args.DataStoreId == nil {
		return nil, errors.New("invalid value for required argument 'DataStoreId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.IndustryVertical == nil {
		return nil, errors.New("invalid value for required argument 'IndustryVertical'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataStore
	err := ctx.RegisterResource("gcp:discoveryengine/dataStore:DataStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataStore gets an existing DataStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataStoreState, opts ...pulumi.ResourceOption) (*DataStore, error) {
	var resource DataStore
	err := ctx.ReadResource("gcp:discoveryengine/dataStore:DataStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataStore resources.
type dataStoreState struct {
	// The content config of the data store.
	// Possible values are: `NO_CONTENT`, `CONTENT_REQUIRED`, `PUBLIC_WEBSITE`.
	ContentConfig *string `pulumi:"contentConfig"`
	// If true, an advanced data store for site search will be created. If the
	// data store is not configured as site search (GENERIC vertical and
	// PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch *bool `pulumi:"createAdvancedSiteSearch"`
	// Timestamp when the DataStore was created.
	CreateTime *string `pulumi:"createTime"`
	// The unique id of the data store.
	//
	// ***
	DataStoreId *string `pulumi:"dataStoreId"`
	// The id of the default Schema associated with this data store.
	DefaultSchemaId *string `pulumi:"defaultSchemaId"`
	// The display name of the data store. This field must be a UTF-8 encoded
	// string with a length limit of 128 characters.
	DisplayName *string `pulumi:"displayName"`
	// The industry vertical that the data store registers.
	// Possible values are: `GENERIC`, `MEDIA`.
	IndustryVertical *string `pulumi:"industryVertical"`
	// The geographic location where the data store should reside. The value can
	// only be one of "global", "us" and "eu".
	Location *string `pulumi:"location"`
	// The unique full resource name of the data store. Values are of the format
	// `projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}`.
	// This field must be a UTF-8 encoded string with a length limit of 1024
	// characters.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The solutions that the data store enrolls.
	// Each value may be one of: `SOLUTION_TYPE_RECOMMENDATION`, `SOLUTION_TYPE_SEARCH`, `SOLUTION_TYPE_CHAT`.
	SolutionTypes []string `pulumi:"solutionTypes"`
}

type DataStoreState struct {
	// The content config of the data store.
	// Possible values are: `NO_CONTENT`, `CONTENT_REQUIRED`, `PUBLIC_WEBSITE`.
	ContentConfig pulumi.StringPtrInput
	// If true, an advanced data store for site search will be created. If the
	// data store is not configured as site search (GENERIC vertical and
	// PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch pulumi.BoolPtrInput
	// Timestamp when the DataStore was created.
	CreateTime pulumi.StringPtrInput
	// The unique id of the data store.
	//
	// ***
	DataStoreId pulumi.StringPtrInput
	// The id of the default Schema associated with this data store.
	DefaultSchemaId pulumi.StringPtrInput
	// The display name of the data store. This field must be a UTF-8 encoded
	// string with a length limit of 128 characters.
	DisplayName pulumi.StringPtrInput
	// The industry vertical that the data store registers.
	// Possible values are: `GENERIC`, `MEDIA`.
	IndustryVertical pulumi.StringPtrInput
	// The geographic location where the data store should reside. The value can
	// only be one of "global", "us" and "eu".
	Location pulumi.StringPtrInput
	// The unique full resource name of the data store. Values are of the format
	// `projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}`.
	// This field must be a UTF-8 encoded string with a length limit of 1024
	// characters.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The solutions that the data store enrolls.
	// Each value may be one of: `SOLUTION_TYPE_RECOMMENDATION`, `SOLUTION_TYPE_SEARCH`, `SOLUTION_TYPE_CHAT`.
	SolutionTypes pulumi.StringArrayInput
}

func (DataStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataStoreState)(nil)).Elem()
}

type dataStoreArgs struct {
	// The content config of the data store.
	// Possible values are: `NO_CONTENT`, `CONTENT_REQUIRED`, `PUBLIC_WEBSITE`.
	ContentConfig string `pulumi:"contentConfig"`
	// If true, an advanced data store for site search will be created. If the
	// data store is not configured as site search (GENERIC vertical and
	// PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch *bool `pulumi:"createAdvancedSiteSearch"`
	// The unique id of the data store.
	//
	// ***
	DataStoreId string `pulumi:"dataStoreId"`
	// The display name of the data store. This field must be a UTF-8 encoded
	// string with a length limit of 128 characters.
	DisplayName string `pulumi:"displayName"`
	// The industry vertical that the data store registers.
	// Possible values are: `GENERIC`, `MEDIA`.
	IndustryVertical string `pulumi:"industryVertical"`
	// The geographic location where the data store should reside. The value can
	// only be one of "global", "us" and "eu".
	Location string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The solutions that the data store enrolls.
	// Each value may be one of: `SOLUTION_TYPE_RECOMMENDATION`, `SOLUTION_TYPE_SEARCH`, `SOLUTION_TYPE_CHAT`.
	SolutionTypes []string `pulumi:"solutionTypes"`
}

// The set of arguments for constructing a DataStore resource.
type DataStoreArgs struct {
	// The content config of the data store.
	// Possible values are: `NO_CONTENT`, `CONTENT_REQUIRED`, `PUBLIC_WEBSITE`.
	ContentConfig pulumi.StringInput
	// If true, an advanced data store for site search will be created. If the
	// data store is not configured as site search (GENERIC vertical and
	// PUBLIC_WEBSITE contentConfig), this flag will be ignored.
	CreateAdvancedSiteSearch pulumi.BoolPtrInput
	// The unique id of the data store.
	//
	// ***
	DataStoreId pulumi.StringInput
	// The display name of the data store. This field must be a UTF-8 encoded
	// string with a length limit of 128 characters.
	DisplayName pulumi.StringInput
	// The industry vertical that the data store registers.
	// Possible values are: `GENERIC`, `MEDIA`.
	IndustryVertical pulumi.StringInput
	// The geographic location where the data store should reside. The value can
	// only be one of "global", "us" and "eu".
	Location pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The solutions that the data store enrolls.
	// Each value may be one of: `SOLUTION_TYPE_RECOMMENDATION`, `SOLUTION_TYPE_SEARCH`, `SOLUTION_TYPE_CHAT`.
	SolutionTypes pulumi.StringArrayInput
}

func (DataStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataStoreArgs)(nil)).Elem()
}

type DataStoreInput interface {
	pulumi.Input

	ToDataStoreOutput() DataStoreOutput
	ToDataStoreOutputWithContext(ctx context.Context) DataStoreOutput
}

func (*DataStore) ElementType() reflect.Type {
	return reflect.TypeOf((**DataStore)(nil)).Elem()
}

func (i *DataStore) ToDataStoreOutput() DataStoreOutput {
	return i.ToDataStoreOutputWithContext(context.Background())
}

func (i *DataStore) ToDataStoreOutputWithContext(ctx context.Context) DataStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreOutput)
}

// DataStoreArrayInput is an input type that accepts DataStoreArray and DataStoreArrayOutput values.
// You can construct a concrete instance of `DataStoreArrayInput` via:
//
//	DataStoreArray{ DataStoreArgs{...} }
type DataStoreArrayInput interface {
	pulumi.Input

	ToDataStoreArrayOutput() DataStoreArrayOutput
	ToDataStoreArrayOutputWithContext(context.Context) DataStoreArrayOutput
}

type DataStoreArray []DataStoreInput

func (DataStoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataStore)(nil)).Elem()
}

func (i DataStoreArray) ToDataStoreArrayOutput() DataStoreArrayOutput {
	return i.ToDataStoreArrayOutputWithContext(context.Background())
}

func (i DataStoreArray) ToDataStoreArrayOutputWithContext(ctx context.Context) DataStoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreArrayOutput)
}

// DataStoreMapInput is an input type that accepts DataStoreMap and DataStoreMapOutput values.
// You can construct a concrete instance of `DataStoreMapInput` via:
//
//	DataStoreMap{ "key": DataStoreArgs{...} }
type DataStoreMapInput interface {
	pulumi.Input

	ToDataStoreMapOutput() DataStoreMapOutput
	ToDataStoreMapOutputWithContext(context.Context) DataStoreMapOutput
}

type DataStoreMap map[string]DataStoreInput

func (DataStoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataStore)(nil)).Elem()
}

func (i DataStoreMap) ToDataStoreMapOutput() DataStoreMapOutput {
	return i.ToDataStoreMapOutputWithContext(context.Background())
}

func (i DataStoreMap) ToDataStoreMapOutputWithContext(ctx context.Context) DataStoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreMapOutput)
}

type DataStoreOutput struct{ *pulumi.OutputState }

func (DataStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataStore)(nil)).Elem()
}

func (o DataStoreOutput) ToDataStoreOutput() DataStoreOutput {
	return o
}

func (o DataStoreOutput) ToDataStoreOutputWithContext(ctx context.Context) DataStoreOutput {
	return o
}

// The content config of the data store.
// Possible values are: `NO_CONTENT`, `CONTENT_REQUIRED`, `PUBLIC_WEBSITE`.
func (o DataStoreOutput) ContentConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.ContentConfig }).(pulumi.StringOutput)
}

// If true, an advanced data store for site search will be created. If the
// data store is not configured as site search (GENERIC vertical and
// PUBLIC_WEBSITE contentConfig), this flag will be ignored.
func (o DataStoreOutput) CreateAdvancedSiteSearch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataStore) pulumi.BoolPtrOutput { return v.CreateAdvancedSiteSearch }).(pulumi.BoolPtrOutput)
}

// Timestamp when the DataStore was created.
func (o DataStoreOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The unique id of the data store.
//
// ***
func (o DataStoreOutput) DataStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.DataStoreId }).(pulumi.StringOutput)
}

// The id of the default Schema associated with this data store.
func (o DataStoreOutput) DefaultSchemaId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.DefaultSchemaId }).(pulumi.StringOutput)
}

// The display name of the data store. This field must be a UTF-8 encoded
// string with a length limit of 128 characters.
func (o DataStoreOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The industry vertical that the data store registers.
// Possible values are: `GENERIC`, `MEDIA`.
func (o DataStoreOutput) IndustryVertical() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.IndustryVertical }).(pulumi.StringOutput)
}

// The geographic location where the data store should reside. The value can
// only be one of "global", "us" and "eu".
func (o DataStoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The unique full resource name of the data store. Values are of the format
// `projects/{project}/locations/{location}/collections/{collection_id}/dataStores/{data_store_id}`.
// This field must be a UTF-8 encoded string with a length limit of 1024
// characters.
func (o DataStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DataStoreOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The solutions that the data store enrolls.
// Each value may be one of: `SOLUTION_TYPE_RECOMMENDATION`, `SOLUTION_TYPE_SEARCH`, `SOLUTION_TYPE_CHAT`.
func (o DataStoreOutput) SolutionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataStore) pulumi.StringArrayOutput { return v.SolutionTypes }).(pulumi.StringArrayOutput)
}

type DataStoreArrayOutput struct{ *pulumi.OutputState }

func (DataStoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataStore)(nil)).Elem()
}

func (o DataStoreArrayOutput) ToDataStoreArrayOutput() DataStoreArrayOutput {
	return o
}

func (o DataStoreArrayOutput) ToDataStoreArrayOutputWithContext(ctx context.Context) DataStoreArrayOutput {
	return o
}

func (o DataStoreArrayOutput) Index(i pulumi.IntInput) DataStoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataStore {
		return vs[0].([]*DataStore)[vs[1].(int)]
	}).(DataStoreOutput)
}

type DataStoreMapOutput struct{ *pulumi.OutputState }

func (DataStoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataStore)(nil)).Elem()
}

func (o DataStoreMapOutput) ToDataStoreMapOutput() DataStoreMapOutput {
	return o
}

func (o DataStoreMapOutput) ToDataStoreMapOutputWithContext(ctx context.Context) DataStoreMapOutput {
	return o
}

func (o DataStoreMapOutput) MapIndex(k pulumi.StringInput) DataStoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataStore {
		return vs[0].(map[string]*DataStore)[vs[1].(string)]
	}).(DataStoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataStoreInput)(nil)).Elem(), &DataStore{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataStoreArrayInput)(nil)).Elem(), DataStoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataStoreMapInput)(nil)).Elem(), DataStoreMap{})
	pulumi.RegisterOutputType(DataStoreOutput{})
	pulumi.RegisterOutputType(DataStoreArrayOutput{})
	pulumi.RegisterOutputType(DataStoreMapOutput{})
}
