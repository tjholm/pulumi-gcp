// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.healthcare;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.healthcare.FhirStoreArgs;
import com.pulumi.gcp.healthcare.inputs.FhirStoreState;
import com.pulumi.gcp.healthcare.outputs.FhirStoreNotificationConfig;
import com.pulumi.gcp.healthcare.outputs.FhirStoreStreamConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A FhirStore is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/fhir/STU3/)
 * standard for Healthcare information exchange
 * 
 * To get more information about FhirStore, see:
 * 
 * * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.fhirStores)
 * * How-to Guides
 *     * [Creating a FHIR store](https://cloud.google.com/healthcare/docs/how-tos/fhir)
 * 
 * ## Example Usage
 * ### Healthcare Fhir Store Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.healthcare.Dataset;
 * import com.pulumi.gcp.healthcare.DatasetArgs;
 * import com.pulumi.gcp.healthcare.FhirStore;
 * import com.pulumi.gcp.healthcare.FhirStoreArgs;
 * import com.pulumi.gcp.healthcare.inputs.FhirStoreNotificationConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var topic = new Topic(&#34;topic&#34;);
 * 
 *         var dataset = new Dataset(&#34;dataset&#34;, DatasetArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .build());
 * 
 *         var default_ = new FhirStore(&#34;default&#34;, FhirStoreArgs.builder()        
 *             .dataset(dataset.id())
 *             .version(&#34;R4&#34;)
 *             .complexDataTypeReferenceParsing(&#34;DISABLED&#34;)
 *             .enableUpdateCreate(false)
 *             .disableReferentialIntegrity(false)
 *             .disableResourceVersioning(false)
 *             .enableHistoryImport(false)
 *             .defaultSearchHandlingStrict(false)
 *             .notificationConfig(FhirStoreNotificationConfigArgs.builder()
 *                 .pubsubTopic(topic.id())
 *                 .build())
 *             .labels(Map.of(&#34;label1&#34;, &#34;labelvalue1&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Healthcare Fhir Store Streaming Config
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.healthcare.Dataset;
 * import com.pulumi.gcp.healthcare.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.healthcare.FhirStore;
 * import com.pulumi.gcp.healthcare.FhirStoreArgs;
 * import com.pulumi.gcp.healthcare.inputs.FhirStoreStreamConfigArgs;
 * import com.pulumi.gcp.healthcare.inputs.FhirStoreStreamConfigBigqueryDestinationArgs;
 * import com.pulumi.gcp.healthcare.inputs.FhirStoreStreamConfigBigqueryDestinationSchemaConfigArgs;
 * import com.pulumi.gcp.healthcare.inputs.FhirStoreStreamConfigBigqueryDestinationSchemaConfigLastUpdatedPartitionConfigArgs;
 * import com.pulumi.gcp.pubsub.Topic;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var dataset = new Dataset(&#34;dataset&#34;, DatasetArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .build());
 * 
 *         var bqDataset = new Dataset(&#34;bqDataset&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;bq_example_dataset&#34;)
 *             .friendlyName(&#34;test&#34;)
 *             .description(&#34;This is a test description&#34;)
 *             .location(&#34;US&#34;)
 *             .deleteContentsOnDestroy(true)
 *             .build());
 * 
 *         var default_ = new FhirStore(&#34;default&#34;, FhirStoreArgs.builder()        
 *             .dataset(dataset.id())
 *             .version(&#34;R4&#34;)
 *             .enableUpdateCreate(false)
 *             .disableReferentialIntegrity(false)
 *             .disableResourceVersioning(false)
 *             .enableHistoryImport(false)
 *             .labels(Map.of(&#34;label1&#34;, &#34;labelvalue1&#34;))
 *             .streamConfigs(FhirStoreStreamConfigArgs.builder()
 *                 .resourceTypes(&#34;Observation&#34;)
 *                 .bigqueryDestination(FhirStoreStreamConfigBigqueryDestinationArgs.builder()
 *                     .datasetUri(Output.tuple(bqDataset.project(), bqDataset.datasetId()).applyValue(values -&gt; {
 *                         var project = values.t1;
 *                         var datasetId = values.t2;
 *                         return String.format(&#34;bq://%s.%s&#34;, project,datasetId);
 *                     }))
 *                     .schemaConfig(FhirStoreStreamConfigBigqueryDestinationSchemaConfigArgs.builder()
 *                         .recursiveStructureDepth(3)
 *                         .lastUpdatedPartitionConfig(FhirStoreStreamConfigBigqueryDestinationSchemaConfigLastUpdatedPartitionConfigArgs.builder()
 *                             .type(&#34;HOUR&#34;)
 *                             .expirationMs(1000000)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var topic = new Topic(&#34;topic&#34;);
 * 
 *     }
 * }
 * ```
 * ### Healthcare Fhir Store Notification Config
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.healthcare.Dataset;
 * import com.pulumi.gcp.healthcare.DatasetArgs;
 * import com.pulumi.gcp.healthcare.FhirStore;
 * import com.pulumi.gcp.healthcare.FhirStoreArgs;
 * import com.pulumi.gcp.healthcare.inputs.FhirStoreNotificationConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var topic = new Topic(&#34;topic&#34;);
 * 
 *         var dataset = new Dataset(&#34;dataset&#34;, DatasetArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .build());
 * 
 *         var default_ = new FhirStore(&#34;default&#34;, FhirStoreArgs.builder()        
 *             .dataset(dataset.id())
 *             .version(&#34;R4&#34;)
 *             .enableUpdateCreate(false)
 *             .disableReferentialIntegrity(false)
 *             .disableResourceVersioning(false)
 *             .enableHistoryImport(false)
 *             .labels(Map.of(&#34;label1&#34;, &#34;labelvalue1&#34;))
 *             .notificationConfig(FhirStoreNotificationConfigArgs.builder()
 *                 .pubsubTopic(topic.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Healthcare Fhir Store Notification Configs
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.pubsub.TopicArgs;
 * import com.pulumi.gcp.healthcare.Dataset;
 * import com.pulumi.gcp.healthcare.DatasetArgs;
 * import com.pulumi.gcp.healthcare.FhirStore;
 * import com.pulumi.gcp.healthcare.FhirStoreArgs;
 * import com.pulumi.gcp.healthcare.inputs.FhirStoreNotificationConfigArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var topic = new Topic(&#34;topic&#34;, TopicArgs.Empty, CustomResourceOptions.builder()
 *             .provider(google_beta)
 *             .build());
 * 
 *         var dataset = new Dataset(&#34;dataset&#34;, DatasetArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var default_ = new FhirStore(&#34;default&#34;, FhirStoreArgs.builder()        
 *             .dataset(dataset.id())
 *             .version(&#34;R4&#34;)
 *             .enableUpdateCreate(false)
 *             .disableReferentialIntegrity(false)
 *             .disableResourceVersioning(false)
 *             .enableHistoryImport(false)
 *             .labels(Map.of(&#34;label1&#34;, &#34;labelvalue1&#34;))
 *             .notificationConfigs(FhirStoreNotificationConfigArgs.builder()
 *                 .pubsubTopic(topic.id())
 *                 .sendFullResource(true)
 *                 .sendPreviousResourceOnDelete(true)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * FhirStore can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:healthcare/fhirStore:FhirStore default {{dataset}}/fhirStores/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:healthcare/fhirStore:FhirStore default {{dataset}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:healthcare/fhirStore:FhirStore")
public class FhirStore extends com.pulumi.resources.CustomResource {
    /**
     * Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
     * Possible values are: `COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED`, `DISABLED`, `ENABLED`.
     * 
     */
    @Export(name="complexDataTypeReferenceParsing", refs={String.class}, tree="[0]")
    private Output<String> complexDataTypeReferenceParsing;

    /**
     * @return Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
     * Possible values are: `COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED`, `DISABLED`, `ENABLED`.
     * 
     */
    public Output<String> complexDataTypeReferenceParsing() {
        return this.complexDataTypeReferenceParsing;
    }
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * &#39;projects/{project}/locations/{location}/datasets/{dataset}&#39;
     * 
     * ***
     * 
     */
    @Export(name="dataset", refs={String.class}, tree="[0]")
    private Output<String> dataset;

    /**
     * @return Identifies the dataset addressed by this request. Must be in the format
     * &#39;projects/{project}/locations/{location}/datasets/{dataset}&#39;
     * 
     * ***
     * 
     */
    public Output<String> dataset() {
        return this.dataset;
    }
    /**
     * If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.
     * If false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.
     * The handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.
     * 
     */
    @Export(name="defaultSearchHandlingStrict", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> defaultSearchHandlingStrict;

    /**
     * @return If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.
     * If false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.
     * The handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.
     * 
     */
    public Output<Optional<Boolean>> defaultSearchHandlingStrict() {
        return Codegen.optional(this.defaultSearchHandlingStrict);
    }
    /**
     * Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
     * creation. The default value is false, meaning that the API will enforce referential integrity and fail the
     * requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
     * will skip referential integrity check. Consequently, operations that rely on references, such as
     * Patient.get$everything, will not return all the results if broken references exist.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * 
     */
    @Export(name="disableReferentialIntegrity", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableReferentialIntegrity;

    /**
     * @return Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
     * creation. The default value is false, meaning that the API will enforce referential integrity and fail the
     * requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
     * will skip referential integrity check. Consequently, operations that rely on references, such as
     * Patient.get$everything, will not return all the results if broken references exist.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * 
     */
    public Output<Optional<Boolean>> disableReferentialIntegrity() {
        return Codegen.optional(this.disableReferentialIntegrity);
    }
    /**
     * Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
     * of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
     * versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
     * cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
     * attempts to read the historical versions.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * 
     */
    @Export(name="disableResourceVersioning", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableResourceVersioning;

    /**
     * @return Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
     * of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
     * versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
     * cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
     * attempts to read the historical versions.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * 
     */
    public Output<Optional<Boolean>> disableResourceVersioning() {
        return Codegen.optional(this.disableResourceVersioning);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Whether to allow the bulk import API to accept history bundles and directly insert historical resource
     * versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
     * occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
     * will fail with an error.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
     * 
     */
    @Export(name="enableHistoryImport", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableHistoryImport;

    /**
     * @return Whether to allow the bulk import API to accept history bundles and directly insert historical resource
     * versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
     * occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
     * will fail with an error.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
     * 
     */
    public Output<Optional<Boolean>> enableHistoryImport() {
        return Codegen.optional(this.enableHistoryImport);
    }
    /**
     * Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
     * operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
     * the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
     * logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
     * identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
     * notifications.
     * 
     */
    @Export(name="enableUpdateCreate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableUpdateCreate;

    /**
     * @return Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
     * operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
     * the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
     * logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
     * identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
     * notifications.
     * 
     */
    public Output<Optional<Boolean>> enableUpdateCreate() {
        return Codegen.optional(this.enableUpdateCreate);
    }
    /**
     * User-supplied key-value pairs used to organize FHIR stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of &#34;key&#34;: value pairs.
     * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return User-supplied key-value pairs used to organize FHIR stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of &#34;key&#34;: value pairs.
     * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The resource name for the FhirStore.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name for the FhirStore.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A nested object resource
     * Structure is documented below.
     * 
     */
    @Export(name="notificationConfig", refs={FhirStoreNotificationConfig.class}, tree="[0]")
    private Output</* @Nullable */ FhirStoreNotificationConfig> notificationConfig;

    /**
     * @return A nested object resource
     * Structure is documented below.
     * 
     */
    public Output<Optional<FhirStoreNotificationConfig>> notificationConfig() {
        return Codegen.optional(this.notificationConfig);
    }
    /**
     * A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.
     * 
     */
    @Export(name="notificationConfigs", refs={List.class,FhirStoreNotificationConfig.class}, tree="[0,1]")
    private Output</* @Nullable */ List<FhirStoreNotificationConfig>> notificationConfigs;

    /**
     * @return A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.
     * 
     */
    public Output<Optional<List<FhirStoreNotificationConfig>>> notificationConfigs() {
        return Codegen.optional(this.notificationConfigs);
    }
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * The fully qualified name of this dataset
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
    private Output<String> selfLink;

    /**
     * @return The fully qualified name of this dataset
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * A list of streaming configs that configure the destinations of streaming export for every resource mutation in
     * this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
     * resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
     * from the list, the server stops streaming to that location. Before adding a new config, you must add the required
     * bigquery.dataEditor role to your project&#39;s Cloud Healthcare Service Agent service account. Some lag (typically on
     * the order of dozens of seconds) is expected before the results show up in the streaming destination.
     * Structure is documented below.
     * 
     */
    @Export(name="streamConfigs", refs={List.class,FhirStoreStreamConfig.class}, tree="[0,1]")
    private Output</* @Nullable */ List<FhirStoreStreamConfig>> streamConfigs;

    /**
     * @return A list of streaming configs that configure the destinations of streaming export for every resource mutation in
     * this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
     * resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
     * from the list, the server stops streaming to that location. Before adding a new config, you must add the required
     * bigquery.dataEditor role to your project&#39;s Cloud Healthcare Service Agent service account. Some lag (typically on
     * the order of dozens of seconds) is expected before the results show up in the streaming destination.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<FhirStoreStreamConfig>>> streamConfigs() {
        return Codegen.optional(this.streamConfigs);
    }
    /**
     * The FHIR specification version.
     * Default value is `STU3`.
     * Possible values are: `DSTU2`, `STU3`, `R4`.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> version;

    /**
     * @return The FHIR specification version.
     * Default value is `STU3`.
     * Possible values are: `DSTU2`, `STU3`, `R4`.
     * 
     */
    public Output<Optional<String>> version() {
        return Codegen.optional(this.version);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FhirStore(String name) {
        this(name, FhirStoreArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FhirStore(String name, FhirStoreArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FhirStore(String name, FhirStoreArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:healthcare/fhirStore:FhirStore", name, args == null ? FhirStoreArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FhirStore(String name, Output<String> id, @Nullable FhirStoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:healthcare/fhirStore:FhirStore", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static FhirStore get(String name, Output<String> id, @Nullable FhirStoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FhirStore(name, id, state, options);
    }
}
