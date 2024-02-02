// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// FeatureView is representation of values that the FeatureOnlineStore will serve based on its syncConfig.
//
// To get more information about FeatureOnlineStoreFeatureview, see:
//
// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores.featureViews)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/vertex-ai/docs)
//
// ## Example Usage
// ### Vertex Ai Featureonlinestore Featureview
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vertex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			featureonlinestore, err := vertex.NewAiFeatureOnlineStore(ctx, "featureonlinestore", &vertex.AiFeatureOnlineStoreArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Region: pulumi.String("us-central1"),
//				Bigtable: &vertex.AiFeatureOnlineStoreBigtableArgs{
//					AutoScaling: &vertex.AiFeatureOnlineStoreBigtableAutoScalingArgs{
//						MinNodeCount:         pulumi.Int(1),
//						MaxNodeCount:         pulumi.Int(2),
//						CpuUtilizationTarget: pulumi.Int(80),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewDataset(ctx, "tf-test-dataset", &bigquery.DatasetArgs{
//				DatasetId:    pulumi.String("example_feature_view"),
//				FriendlyName: pulumi.String("test"),
//				Description:  pulumi.String("This is a test description"),
//				Location:     pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewTable(ctx, "tf-test-table", &bigquery.TableArgs{
//				DeletionProtection: pulumi.Bool(false),
//				DatasetId:          tf_test_dataset.DatasetId,
//				TableId:            pulumi.String("example_feature_view"),
//				Schema: pulumi.String(`  [
//	  {
//	    "name": "entity_id",
//	    "mode": "NULLABLE",
//	    "type": "STRING",
//	    "description": "Test default entity_id"
//	  },
//	    {
//	    "name": "test_entity_column",
//	    "mode": "NULLABLE",
//	    "type": "STRING",
//	    "description": "test secondary entity column"
//	  },
//	  {
//	    "name": "feature_timestamp",
//	    "mode": "NULLABLE",
//	    "type": "TIMESTAMP",
//	    "description": "Default timestamp value"
//	  }
//
// ]
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vertex.NewAiFeatureOnlineStoreFeatureview(ctx, "featureview", &vertex.AiFeatureOnlineStoreFeatureviewArgs{
//				Region:             pulumi.String("us-central1"),
//				FeatureOnlineStore: featureonlinestore.Name,
//				SyncConfig: &vertex.AiFeatureOnlineStoreFeatureviewSyncConfigArgs{
//					Cron: pulumi.String("0 0 * * *"),
//				},
//				BigQuerySource: &vertex.AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs{
//					Uri: pulumi.All(tf_test_table.Project, tf_test_table.DatasetId, tf_test_table.TableId).ApplyT(func(_args []interface{}) (string, error) {
//						project := _args[0].(string)
//						datasetId := _args[1].(string)
//						tableId := _args[2].(string)
//						return fmt.Sprintf("bq://%v.%v.%v", project, datasetId, tableId), nil
//					}).(pulumi.StringOutput),
//					EntityIdColumns: pulumi.StringArray{
//						pulumi.String("test_entity_column"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Vertex Ai Featureonlinestore Featureview With Vector Search
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vertex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			featureonlinestore, err := vertex.NewAiFeatureOnlineStore(ctx, "featureonlinestore", &vertex.AiFeatureOnlineStoreArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Region: pulumi.String("us-central1"),
//				Bigtable: &vertex.AiFeatureOnlineStoreBigtableArgs{
//					AutoScaling: &vertex.AiFeatureOnlineStoreBigtableAutoScalingArgs{
//						MinNodeCount:         pulumi.Int(1),
//						MaxNodeCount:         pulumi.Int(2),
//						CpuUtilizationTarget: pulumi.Int(80),
//					},
//				},
//				EmbeddingManagement: &vertex.AiFeatureOnlineStoreEmbeddingManagementArgs{
//					Enabled: pulumi.Bool(true),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewDataset(ctx, "tf-test-dataset", &bigquery.DatasetArgs{
//				DatasetId:    pulumi.String("example_feature_view_vector_search"),
//				FriendlyName: pulumi.String("test"),
//				Description:  pulumi.String("This is a test description"),
//				Location:     pulumi.String("US"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewTable(ctx, "tf-test-table", &bigquery.TableArgs{
//				DeletionProtection: pulumi.Bool(false),
//				DatasetId:          tf_test_dataset.DatasetId,
//				TableId:            pulumi.String("example_feature_view_vector_search"),
//				Schema: pulumi.String(`[
//
//	{
//	  "name": "test_primary_id",
//	  "mode": "NULLABLE",
//	  "type": "STRING",
//	  "description": "primary test id"
//	},
//
//	{
//	  "name": "embedding",
//	  "mode": "REPEATED",
//	  "type": "FLOAT",
//	  "description": "embedding column for primary_id column"
//	},
//
//	{
//	  "name": "country",
//	  "mode": "NULLABLE",
//	  "type": "STRING",
//	  "description": "country"
//	},
//
//	{
//	  "name": "test_crowding_column",
//	  "mode": "NULLABLE",
//	  "type": "INTEGER",
//	  "description": "test crowding column"
//	},
//
//	{
//	  "name": "entity_id",
//	  "mode": "NULLABLE",
//	  "type": "STRING",
//	  "description": "Test default entity_id"
//	},
//
//	{
//	  "name": "test_entity_column",
//	  "mode": "NULLABLE",
//	  "type": "STRING",
//	  "description": "test secondary entity column"
//	},
//
//	{
//	  "name": "feature_timestamp",
//	  "mode": "NULLABLE",
//	  "type": "TIMESTAMP",
//	  "description": "Default timestamp value"
//	}
//
// ]
// `),
//
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = vertex.NewAiFeatureOnlineStoreFeatureview(ctx, "featureviewVectorSearch", &vertex.AiFeatureOnlineStoreFeatureviewArgs{
//				Region:             pulumi.String("us-central1"),
//				FeatureOnlineStore: featureonlinestore.Name,
//				SyncConfig: &vertex.AiFeatureOnlineStoreFeatureviewSyncConfigArgs{
//					Cron: pulumi.String("0 0 * * *"),
//				},
//				BigQuerySource: &vertex.AiFeatureOnlineStoreFeatureviewBigQuerySourceArgs{
//					Uri: pulumi.All(tf_test_table.Project, tf_test_table.DatasetId, tf_test_table.TableId).ApplyT(func(_args []interface{}) (string, error) {
//						project := _args[0].(string)
//						datasetId := _args[1].(string)
//						tableId := _args[2].(string)
//						return fmt.Sprintf("bq://%v.%v.%v", project, datasetId, tableId), nil
//					}).(pulumi.StringOutput),
//					EntityIdColumns: pulumi.StringArray{
//						pulumi.String("test_entity_column"),
//					},
//				},
//				VectorSearchConfig: &vertex.AiFeatureOnlineStoreFeatureviewVectorSearchConfigArgs{
//					EmbeddingColumn: pulumi.String("embedding"),
//					FilterColumns: pulumi.StringArray{
//						pulumi.String("country"),
//					},
//					CrowdingColumn:      pulumi.String("test_crowding_column"),
//					DistanceMeasureType: pulumi.String("DOT_PRODUCT_DISTANCE"),
//					TreeAhConfig: &vertex.AiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfigArgs{
//						LeafNodeEmbeddingCount: pulumi.String("1000"),
//					},
//					EmbeddingDimension: pulumi.Int(2),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = organizations.LookupProject(ctx, nil, nil)
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
// FeatureOnlineStoreFeatureview can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/featureOnlineStores/{{feature_online_store}}/featureViews/{{name}}` * `{{project}}/{{region}}/{{feature_online_store}}/{{name}}` * `{{region}}/{{feature_online_store}}/{{name}}` * `{{feature_online_store}}/{{name}}` When using the `pulumi import` command, FeatureOnlineStoreFeatureview can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview default projects/{{project}}/locations/{{region}}/featureOnlineStores/{{feature_online_store}}/featureViews/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview default {{project}}/{{region}}/{{feature_online_store}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview default {{region}}/{{feature_online_store}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview default {{feature_online_store}}/{{name}}
//
// ```
type AiFeatureOnlineStoreFeatureview struct {
	pulumi.CustomResourceState

	// Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
	// Structure is documented below.
	BigQuerySource AiFeatureOnlineStoreFeatureviewBigQuerySourcePtrOutput `pulumi:"bigQuerySource"`
	// The timestamp of when the featureOnlinestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The name of the FeatureOnlineStore to use for the featureview.
	FeatureOnlineStore pulumi.StringOutput `pulumi:"featureOnlineStore"`
	// A set of key/value label pairs to assign to this FeatureView.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// The region for the resource. It should be the same as the featureonlinestore region.
	//
	// ***
	Region pulumi.StringOutput `pulumi:"region"`
	// Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
	// Structure is documented below.
	SyncConfig AiFeatureOnlineStoreFeatureviewSyncConfigPtrOutput `pulumi:"syncConfig"`
	// The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
	// Structure is documented below.
	VectorSearchConfig AiFeatureOnlineStoreFeatureviewVectorSearchConfigPtrOutput `pulumi:"vectorSearchConfig"`
}

// NewAiFeatureOnlineStoreFeatureview registers a new resource with the given unique name, arguments, and options.
func NewAiFeatureOnlineStoreFeatureview(ctx *pulumi.Context,
	name string, args *AiFeatureOnlineStoreFeatureviewArgs, opts ...pulumi.ResourceOption) (*AiFeatureOnlineStoreFeatureview, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeatureOnlineStore == nil {
		return nil, errors.New("invalid value for required argument 'FeatureOnlineStore'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiFeatureOnlineStoreFeatureview
	err := ctx.RegisterResource("gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiFeatureOnlineStoreFeatureview gets an existing AiFeatureOnlineStoreFeatureview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiFeatureOnlineStoreFeatureview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiFeatureOnlineStoreFeatureviewState, opts ...pulumi.ResourceOption) (*AiFeatureOnlineStoreFeatureview, error) {
	var resource AiFeatureOnlineStoreFeatureview
	err := ctx.ReadResource("gcp:vertex/aiFeatureOnlineStoreFeatureview:AiFeatureOnlineStoreFeatureview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiFeatureOnlineStoreFeatureview resources.
type aiFeatureOnlineStoreFeatureviewState struct {
	// Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
	// Structure is documented below.
	BigQuerySource *AiFeatureOnlineStoreFeatureviewBigQuerySource `pulumi:"bigQuerySource"`
	// The timestamp of when the featureOnlinestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime *string `pulumi:"createTime"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The name of the FeatureOnlineStore to use for the featureview.
	FeatureOnlineStore *string `pulumi:"featureOnlineStore"`
	// A set of key/value label pairs to assign to this FeatureView.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// The region for the resource. It should be the same as the featureonlinestore region.
	//
	// ***
	Region *string `pulumi:"region"`
	// Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
	// Structure is documented below.
	SyncConfig *AiFeatureOnlineStoreFeatureviewSyncConfig `pulumi:"syncConfig"`
	// The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime *string `pulumi:"updateTime"`
	// Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
	// Structure is documented below.
	VectorSearchConfig *AiFeatureOnlineStoreFeatureviewVectorSearchConfig `pulumi:"vectorSearchConfig"`
}

type AiFeatureOnlineStoreFeatureviewState struct {
	// Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
	// Structure is documented below.
	BigQuerySource AiFeatureOnlineStoreFeatureviewBigQuerySourcePtrInput
	// The timestamp of when the featureOnlinestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapInput
	// The name of the FeatureOnlineStore to use for the featureview.
	FeatureOnlineStore pulumi.StringPtrInput
	// A set of key/value label pairs to assign to this FeatureView.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// The region for the resource. It should be the same as the featureonlinestore region.
	//
	// ***
	Region pulumi.StringPtrInput
	// Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
	// Structure is documented below.
	SyncConfig AiFeatureOnlineStoreFeatureviewSyncConfigPtrInput
	// The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime pulumi.StringPtrInput
	// Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
	// Structure is documented below.
	VectorSearchConfig AiFeatureOnlineStoreFeatureviewVectorSearchConfigPtrInput
}

func (AiFeatureOnlineStoreFeatureviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureOnlineStoreFeatureviewState)(nil)).Elem()
}

type aiFeatureOnlineStoreFeatureviewArgs struct {
	// Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
	// Structure is documented below.
	BigQuerySource *AiFeatureOnlineStoreFeatureviewBigQuerySource `pulumi:"bigQuerySource"`
	// The name of the FeatureOnlineStore to use for the featureview.
	FeatureOnlineStore string `pulumi:"featureOnlineStore"`
	// A set of key/value label pairs to assign to this FeatureView.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region for the resource. It should be the same as the featureonlinestore region.
	//
	// ***
	Region string `pulumi:"region"`
	// Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
	// Structure is documented below.
	SyncConfig *AiFeatureOnlineStoreFeatureviewSyncConfig `pulumi:"syncConfig"`
	// Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
	// Structure is documented below.
	VectorSearchConfig *AiFeatureOnlineStoreFeatureviewVectorSearchConfig `pulumi:"vectorSearchConfig"`
}

// The set of arguments for constructing a AiFeatureOnlineStoreFeatureview resource.
type AiFeatureOnlineStoreFeatureviewArgs struct {
	// Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
	// Structure is documented below.
	BigQuerySource AiFeatureOnlineStoreFeatureviewBigQuerySourcePtrInput
	// The name of the FeatureOnlineStore to use for the featureview.
	FeatureOnlineStore pulumi.StringInput
	// A set of key/value label pairs to assign to this FeatureView.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region for the resource. It should be the same as the featureonlinestore region.
	//
	// ***
	Region pulumi.StringInput
	// Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
	// Structure is documented below.
	SyncConfig AiFeatureOnlineStoreFeatureviewSyncConfigPtrInput
	// Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
	// Structure is documented below.
	VectorSearchConfig AiFeatureOnlineStoreFeatureviewVectorSearchConfigPtrInput
}

func (AiFeatureOnlineStoreFeatureviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureOnlineStoreFeatureviewArgs)(nil)).Elem()
}

type AiFeatureOnlineStoreFeatureviewInput interface {
	pulumi.Input

	ToAiFeatureOnlineStoreFeatureviewOutput() AiFeatureOnlineStoreFeatureviewOutput
	ToAiFeatureOnlineStoreFeatureviewOutputWithContext(ctx context.Context) AiFeatureOnlineStoreFeatureviewOutput
}

func (*AiFeatureOnlineStoreFeatureview) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureOnlineStoreFeatureview)(nil)).Elem()
}

func (i *AiFeatureOnlineStoreFeatureview) ToAiFeatureOnlineStoreFeatureviewOutput() AiFeatureOnlineStoreFeatureviewOutput {
	return i.ToAiFeatureOnlineStoreFeatureviewOutputWithContext(context.Background())
}

func (i *AiFeatureOnlineStoreFeatureview) ToAiFeatureOnlineStoreFeatureviewOutputWithContext(ctx context.Context) AiFeatureOnlineStoreFeatureviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureOnlineStoreFeatureviewOutput)
}

// AiFeatureOnlineStoreFeatureviewArrayInput is an input type that accepts AiFeatureOnlineStoreFeatureviewArray and AiFeatureOnlineStoreFeatureviewArrayOutput values.
// You can construct a concrete instance of `AiFeatureOnlineStoreFeatureviewArrayInput` via:
//
//	AiFeatureOnlineStoreFeatureviewArray{ AiFeatureOnlineStoreFeatureviewArgs{...} }
type AiFeatureOnlineStoreFeatureviewArrayInput interface {
	pulumi.Input

	ToAiFeatureOnlineStoreFeatureviewArrayOutput() AiFeatureOnlineStoreFeatureviewArrayOutput
	ToAiFeatureOnlineStoreFeatureviewArrayOutputWithContext(context.Context) AiFeatureOnlineStoreFeatureviewArrayOutput
}

type AiFeatureOnlineStoreFeatureviewArray []AiFeatureOnlineStoreFeatureviewInput

func (AiFeatureOnlineStoreFeatureviewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureOnlineStoreFeatureview)(nil)).Elem()
}

func (i AiFeatureOnlineStoreFeatureviewArray) ToAiFeatureOnlineStoreFeatureviewArrayOutput() AiFeatureOnlineStoreFeatureviewArrayOutput {
	return i.ToAiFeatureOnlineStoreFeatureviewArrayOutputWithContext(context.Background())
}

func (i AiFeatureOnlineStoreFeatureviewArray) ToAiFeatureOnlineStoreFeatureviewArrayOutputWithContext(ctx context.Context) AiFeatureOnlineStoreFeatureviewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureOnlineStoreFeatureviewArrayOutput)
}

// AiFeatureOnlineStoreFeatureviewMapInput is an input type that accepts AiFeatureOnlineStoreFeatureviewMap and AiFeatureOnlineStoreFeatureviewMapOutput values.
// You can construct a concrete instance of `AiFeatureOnlineStoreFeatureviewMapInput` via:
//
//	AiFeatureOnlineStoreFeatureviewMap{ "key": AiFeatureOnlineStoreFeatureviewArgs{...} }
type AiFeatureOnlineStoreFeatureviewMapInput interface {
	pulumi.Input

	ToAiFeatureOnlineStoreFeatureviewMapOutput() AiFeatureOnlineStoreFeatureviewMapOutput
	ToAiFeatureOnlineStoreFeatureviewMapOutputWithContext(context.Context) AiFeatureOnlineStoreFeatureviewMapOutput
}

type AiFeatureOnlineStoreFeatureviewMap map[string]AiFeatureOnlineStoreFeatureviewInput

func (AiFeatureOnlineStoreFeatureviewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureOnlineStoreFeatureview)(nil)).Elem()
}

func (i AiFeatureOnlineStoreFeatureviewMap) ToAiFeatureOnlineStoreFeatureviewMapOutput() AiFeatureOnlineStoreFeatureviewMapOutput {
	return i.ToAiFeatureOnlineStoreFeatureviewMapOutputWithContext(context.Background())
}

func (i AiFeatureOnlineStoreFeatureviewMap) ToAiFeatureOnlineStoreFeatureviewMapOutputWithContext(ctx context.Context) AiFeatureOnlineStoreFeatureviewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureOnlineStoreFeatureviewMapOutput)
}

type AiFeatureOnlineStoreFeatureviewOutput struct{ *pulumi.OutputState }

func (AiFeatureOnlineStoreFeatureviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureOnlineStoreFeatureview)(nil)).Elem()
}

func (o AiFeatureOnlineStoreFeatureviewOutput) ToAiFeatureOnlineStoreFeatureviewOutput() AiFeatureOnlineStoreFeatureviewOutput {
	return o
}

func (o AiFeatureOnlineStoreFeatureviewOutput) ToAiFeatureOnlineStoreFeatureviewOutputWithContext(ctx context.Context) AiFeatureOnlineStoreFeatureviewOutput {
	return o
}

// Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
// Structure is documented below.
func (o AiFeatureOnlineStoreFeatureviewOutput) BigQuerySource() AiFeatureOnlineStoreFeatureviewBigQuerySourcePtrOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) AiFeatureOnlineStoreFeatureviewBigQuerySourcePtrOutput {
		return v.BigQuerySource
	}).(AiFeatureOnlineStoreFeatureviewBigQuerySourcePtrOutput)
}

// The timestamp of when the featureOnlinestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
func (o AiFeatureOnlineStoreFeatureviewOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o AiFeatureOnlineStoreFeatureviewOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The name of the FeatureOnlineStore to use for the featureview.
func (o AiFeatureOnlineStoreFeatureviewOutput) FeatureOnlineStore() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringOutput { return v.FeatureOnlineStore }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to this FeatureView.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o AiFeatureOnlineStoreFeatureviewOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the FeatureView. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
func (o AiFeatureOnlineStoreFeatureviewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o AiFeatureOnlineStoreFeatureviewOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o AiFeatureOnlineStoreFeatureviewOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// The region for the resource. It should be the same as the featureonlinestore region.
//
// ***
func (o AiFeatureOnlineStoreFeatureviewOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
// Structure is documented below.
func (o AiFeatureOnlineStoreFeatureviewOutput) SyncConfig() AiFeatureOnlineStoreFeatureviewSyncConfigPtrOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) AiFeatureOnlineStoreFeatureviewSyncConfigPtrOutput {
		return v.SyncConfig
	}).(AiFeatureOnlineStoreFeatureviewSyncConfigPtrOutput)
}

// The timestamp of when the featureOnlinestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
func (o AiFeatureOnlineStoreFeatureviewOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
// Structure is documented below.
func (o AiFeatureOnlineStoreFeatureviewOutput) VectorSearchConfig() AiFeatureOnlineStoreFeatureviewVectorSearchConfigPtrOutput {
	return o.ApplyT(func(v *AiFeatureOnlineStoreFeatureview) AiFeatureOnlineStoreFeatureviewVectorSearchConfigPtrOutput {
		return v.VectorSearchConfig
	}).(AiFeatureOnlineStoreFeatureviewVectorSearchConfigPtrOutput)
}

type AiFeatureOnlineStoreFeatureviewArrayOutput struct{ *pulumi.OutputState }

func (AiFeatureOnlineStoreFeatureviewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureOnlineStoreFeatureview)(nil)).Elem()
}

func (o AiFeatureOnlineStoreFeatureviewArrayOutput) ToAiFeatureOnlineStoreFeatureviewArrayOutput() AiFeatureOnlineStoreFeatureviewArrayOutput {
	return o
}

func (o AiFeatureOnlineStoreFeatureviewArrayOutput) ToAiFeatureOnlineStoreFeatureviewArrayOutputWithContext(ctx context.Context) AiFeatureOnlineStoreFeatureviewArrayOutput {
	return o
}

func (o AiFeatureOnlineStoreFeatureviewArrayOutput) Index(i pulumi.IntInput) AiFeatureOnlineStoreFeatureviewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiFeatureOnlineStoreFeatureview {
		return vs[0].([]*AiFeatureOnlineStoreFeatureview)[vs[1].(int)]
	}).(AiFeatureOnlineStoreFeatureviewOutput)
}

type AiFeatureOnlineStoreFeatureviewMapOutput struct{ *pulumi.OutputState }

func (AiFeatureOnlineStoreFeatureviewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureOnlineStoreFeatureview)(nil)).Elem()
}

func (o AiFeatureOnlineStoreFeatureviewMapOutput) ToAiFeatureOnlineStoreFeatureviewMapOutput() AiFeatureOnlineStoreFeatureviewMapOutput {
	return o
}

func (o AiFeatureOnlineStoreFeatureviewMapOutput) ToAiFeatureOnlineStoreFeatureviewMapOutputWithContext(ctx context.Context) AiFeatureOnlineStoreFeatureviewMapOutput {
	return o
}

func (o AiFeatureOnlineStoreFeatureviewMapOutput) MapIndex(k pulumi.StringInput) AiFeatureOnlineStoreFeatureviewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiFeatureOnlineStoreFeatureview {
		return vs[0].(map[string]*AiFeatureOnlineStoreFeatureview)[vs[1].(string)]
	}).(AiFeatureOnlineStoreFeatureviewOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureOnlineStoreFeatureviewInput)(nil)).Elem(), &AiFeatureOnlineStoreFeatureview{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureOnlineStoreFeatureviewArrayInput)(nil)).Elem(), AiFeatureOnlineStoreFeatureviewArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureOnlineStoreFeatureviewMapInput)(nil)).Elem(), AiFeatureOnlineStoreFeatureviewMap{})
	pulumi.RegisterOutputType(AiFeatureOnlineStoreFeatureviewOutput{})
	pulumi.RegisterOutputType(AiFeatureOnlineStoreFeatureviewArrayOutput{})
	pulumi.RegisterOutputType(AiFeatureOnlineStoreFeatureviewMapOutput{})
}
