// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigquery.JobArgs;
import com.pulumi.gcp.bigquery.inputs.JobState;
import com.pulumi.gcp.bigquery.outputs.JobCopy;
import com.pulumi.gcp.bigquery.outputs.JobExtract;
import com.pulumi.gcp.bigquery.outputs.JobLoad;
import com.pulumi.gcp.bigquery.outputs.JobQuery;
import com.pulumi.gcp.bigquery.outputs.JobStatus;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Jobs are actions that BigQuery runs on your behalf to load data, export data, query data, or copy data.
 * Once a BigQuery job is created, it cannot be changed or deleted.
 * 
 * To get more information about Job, see:
 * 
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/jobs)
 * * How-to Guides
 *     * [BigQuery Jobs Intro](https://cloud.google.com/bigquery/docs/jobs-overview)
 * 
 * ## Example Usage
 * 
 * ### Bigquery Job Query
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.bigquery.Job;
 * import com.pulumi.gcp.bigquery.JobArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobQueryArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobQueryDestinationTableArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobQueryScriptOptionsArgs;
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
 *         var bar = new Dataset("bar", DatasetArgs.builder()
 *             .datasetId("job_query_dataset")
 *             .friendlyName("test")
 *             .description("This is a test description")
 *             .location("US")
 *             .build());
 * 
 *         var foo = new Table("foo", TableArgs.builder()
 *             .deletionProtection(false)
 *             .datasetId(bar.datasetId())
 *             .tableId("job_query_table")
 *             .build());
 * 
 *         var job = new Job("job", JobArgs.builder()
 *             .jobId("job_query")
 *             .labels(Map.of("example-label", "example-value"))
 *             .query(JobQueryArgs.builder()
 *                 .query("SELECT state FROM [lookerdata:cdc.project_tycho_reports]")
 *                 .destinationTable(JobQueryDestinationTableArgs.builder()
 *                     .projectId(foo.project())
 *                     .datasetId(foo.datasetId())
 *                     .tableId(foo.tableId())
 *                     .build())
 *                 .allowLargeResults(true)
 *                 .flattenResults(true)
 *                 .scriptOptions(JobQueryScriptOptionsArgs.builder()
 *                     .keyResultStatement("LAST")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Bigquery Job Query Table Reference
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.bigquery.Job;
 * import com.pulumi.gcp.bigquery.JobArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobQueryArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobQueryDestinationTableArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobQueryDefaultDatasetArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobQueryScriptOptionsArgs;
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
 *         var bar = new Dataset("bar", DatasetArgs.builder()
 *             .datasetId("job_query_dataset")
 *             .friendlyName("test")
 *             .description("This is a test description")
 *             .location("US")
 *             .build());
 * 
 *         var foo = new Table("foo", TableArgs.builder()
 *             .deletionProtection(false)
 *             .datasetId(bar.datasetId())
 *             .tableId("job_query_table")
 *             .build());
 * 
 *         var job = new Job("job", JobArgs.builder()
 *             .jobId("job_query")
 *             .labels(Map.of("example-label", "example-value"))
 *             .query(JobQueryArgs.builder()
 *                 .query("SELECT state FROM [lookerdata:cdc.project_tycho_reports]")
 *                 .destinationTable(JobQueryDestinationTableArgs.builder()
 *                     .tableId(foo.id())
 *                     .build())
 *                 .defaultDataset(JobQueryDefaultDatasetArgs.builder()
 *                     .datasetId(bar.id())
 *                     .build())
 *                 .allowLargeResults(true)
 *                 .flattenResults(true)
 *                 .scriptOptions(JobQueryScriptOptionsArgs.builder()
 *                     .keyResultStatement("LAST")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Bigquery Job Load
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.bigquery.Job;
 * import com.pulumi.gcp.bigquery.JobArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobLoadArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobLoadDestinationTableArgs;
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
 *         var bar = new Dataset("bar", DatasetArgs.builder()
 *             .datasetId("job_load_dataset")
 *             .friendlyName("test")
 *             .description("This is a test description")
 *             .location("US")
 *             .build());
 * 
 *         var foo = new Table("foo", TableArgs.builder()
 *             .deletionProtection(false)
 *             .datasetId(bar.datasetId())
 *             .tableId("job_load_table")
 *             .build());
 * 
 *         var job = new Job("job", JobArgs.builder()
 *             .jobId("job_load")
 *             .labels(Map.of("my_job", "load"))
 *             .load(JobLoadArgs.builder()
 *                 .sourceUris("gs://cloud-samples-data/bigquery/us-states/us-states-by-date.csv")
 *                 .destinationTable(JobLoadDestinationTableArgs.builder()
 *                     .projectId(foo.project())
 *                     .datasetId(foo.datasetId())
 *                     .tableId(foo.tableId())
 *                     .build())
 *                 .skipLeadingRows(1)
 *                 .schemaUpdateOptions(                
 *                     "ALLOW_FIELD_RELAXATION",
 *                     "ALLOW_FIELD_ADDITION")
 *                 .writeDisposition("WRITE_APPEND")
 *                 .autodetect(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Bigquery Job Load Geojson
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.storage.BucketObject;
 * import com.pulumi.gcp.storage.BucketObjectArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.bigquery.Job;
 * import com.pulumi.gcp.bigquery.JobArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobLoadArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobLoadDestinationTableArgs;
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
 *         final var project = "my-project-name";
 * 
 *         var bucket = new Bucket("bucket", BucketArgs.builder()
 *             .name(String.format("%s-bq-geojson", project))
 *             .location("US")
 *             .uniformBucketLevelAccess(true)
 *             .build());
 * 
 *         var object = new BucketObject("object", BucketObjectArgs.builder()
 *             .name("geojson-data.jsonl")
 *             .bucket(bucket.name())
 *             .content("""
 * {"type":"Feature","properties":{"continent":"Europe","region":"Scandinavia"},"geometry":{"type":"Polygon","coordinates":[[[-30.94,53.33],[33.05,53.33],[33.05,71.86],[-30.94,71.86],[-30.94,53.33]]]}}
 * {"type":"Feature","properties":{"continent":"Africa","region":"West Africa"},"geometry":{"type":"Polygon","coordinates":[[[-23.91,0],[11.95,0],[11.95,18.98],[-23.91,18.98],[-23.91,0]]]}}
 *             """)
 *             .build());
 * 
 *         var bar = new Dataset("bar", DatasetArgs.builder()
 *             .datasetId("job_load_dataset")
 *             .friendlyName("test")
 *             .description("This is a test description")
 *             .location("US")
 *             .build());
 * 
 *         var foo = new Table("foo", TableArgs.builder()
 *             .deletionProtection(false)
 *             .datasetId(bar.datasetId())
 *             .tableId("job_load_table")
 *             .build());
 * 
 *         var job = new Job("job", JobArgs.builder()
 *             .jobId("job_load")
 *             .labels(Map.of("my_job", "load"))
 *             .load(JobLoadArgs.builder()
 *                 .sourceUris(Output.tuple(object.bucket(), object.name()).applyValue(values -> {
 *                     var bucket = values.t1;
 *                     var name = values.t2;
 *                     return String.format("gs://%s/%s", bucket,name);
 *                 }))
 *                 .destinationTable(JobLoadDestinationTableArgs.builder()
 *                     .projectId(foo.project())
 *                     .datasetId(foo.datasetId())
 *                     .tableId(foo.tableId())
 *                     .build())
 *                 .writeDisposition("WRITE_TRUNCATE")
 *                 .autodetect(true)
 *                 .sourceFormat("NEWLINE_DELIMITED_JSON")
 *                 .jsonExtension("GEOJSON")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Bigquery Job Load Parquet
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.storage.BucketObject;
 * import com.pulumi.gcp.storage.BucketObjectArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.bigquery.Job;
 * import com.pulumi.gcp.bigquery.JobArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobLoadArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobLoadDestinationTableArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobLoadParquetOptionsArgs;
 * import com.pulumi.asset.FileAsset;
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
 *         var test = new Bucket("test", BucketArgs.builder()
 *             .name("job_load_bucket")
 *             .location("US")
 *             .uniformBucketLevelAccess(true)
 *             .build());
 * 
 *         var testBucketObject = new BucketObject("testBucketObject", BucketObjectArgs.builder()
 *             .name("job_load_bucket_object")
 *             .source(new FileAsset("./test-fixtures/test.parquet.gzip"))
 *             .bucket(test.name())
 *             .build());
 * 
 *         var testDataset = new Dataset("testDataset", DatasetArgs.builder()
 *             .datasetId("job_load_dataset")
 *             .friendlyName("test")
 *             .description("This is a test description")
 *             .location("US")
 *             .build());
 * 
 *         var testTable = new Table("testTable", TableArgs.builder()
 *             .deletionProtection(false)
 *             .tableId("job_load_table")
 *             .datasetId(testDataset.datasetId())
 *             .build());
 * 
 *         var job = new Job("job", JobArgs.builder()
 *             .jobId("job_load")
 *             .labels(Map.of("my_job", "load"))
 *             .load(JobLoadArgs.builder()
 *                 .sourceUris(Output.tuple(testBucketObject.bucket(), testBucketObject.name()).applyValue(values -> {
 *                     var bucket = values.t1;
 *                     var name = values.t2;
 *                     return String.format("gs://%s/%s", bucket,name);
 *                 }))
 *                 .destinationTable(JobLoadDestinationTableArgs.builder()
 *                     .projectId(testTable.project())
 *                     .datasetId(testTable.datasetId())
 *                     .tableId(testTable.tableId())
 *                     .build())
 *                 .schemaUpdateOptions(                
 *                     "ALLOW_FIELD_RELAXATION",
 *                     "ALLOW_FIELD_ADDITION")
 *                 .writeDisposition("WRITE_APPEND")
 *                 .sourceFormat("PARQUET")
 *                 .autodetect(true)
 *                 .parquetOptions(JobLoadParquetOptionsArgs.builder()
 *                     .enumAsString(true)
 *                     .enableListInference(true)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Bigquery Job Copy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.bigquery.inputs.TableEncryptionConfigurationArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.projects.IAMMember;
 * import com.pulumi.gcp.projects.IAMMemberArgs;
 * import com.pulumi.gcp.bigquery.Job;
 * import com.pulumi.gcp.bigquery.JobArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobCopyArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobCopyDestinationTableArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobCopyDestinationEncryptionConfigurationArgs;
 * import com.pulumi.codegen.internal.KeyedValue;
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
 *         final var count = 2;
 * 
 *         for (var i = 0; i < count; i++) {
 *             new Dataset("sourceDataset-" + i, DatasetArgs.builder()
 *                 .datasetId(String.format("job_copy_%s_dataset", range.value()))
 *                 .friendlyName("test")
 *                 .description("This is a test description")
 *                 .location("US")
 *                 .build());
 * 
 *         
 * }
 *         for (var i = 0; i < count; i++) {
 *             new Table("source-" + i, TableArgs.builder()
 *                 .deletionProtection(false)
 *                 .datasetId(sourceDataset[range.value()].datasetId())
 *                 .tableId(String.format("job_copy_%s_table", range.value()))
 *                 .schema("""
 * [
 *   {
 *     "name": "name",
 *     "type": "STRING",
 *     "mode": "NULLABLE"
 *   },
 *   {
 *     "name": "post_abbr",
 *     "type": "STRING",
 *     "mode": "NULLABLE"
 *   },
 *   {
 *     "name": "date",
 *     "type": "DATE",
 *     "mode": "NULLABLE"
 *   }
 * ]
 *                 """)
 *                 .build());
 * 
 *         
 * }
 *         var destDataset = new Dataset("destDataset", DatasetArgs.builder()
 *             .datasetId("job_copy_dest_dataset")
 *             .friendlyName("test")
 *             .description("This is a test description")
 *             .location("US")
 *             .build());
 * 
 *         var keyRing = new KeyRing("keyRing", KeyRingArgs.builder()
 *             .name("example-keyring")
 *             .location("global")
 *             .build());
 * 
 *         var cryptoKey = new CryptoKey("cryptoKey", CryptoKeyArgs.builder()
 *             .name("example-key")
 *             .keyRing(keyRing.id())
 *             .build());
 * 
 *         var dest = new Table("dest", TableArgs.builder()
 *             .deletionProtection(false)
 *             .datasetId(destDataset.datasetId())
 *             .tableId("job_copy_dest_table")
 *             .schema("""
 * [
 *   {
 *     "name": "name",
 *     "type": "STRING",
 *     "mode": "NULLABLE"
 *   },
 *   {
 *     "name": "post_abbr",
 *     "type": "STRING",
 *     "mode": "NULLABLE"
 *   },
 *   {
 *     "name": "date",
 *     "type": "DATE",
 *     "mode": "NULLABLE"
 *   }
 * ]
 *             """)
 *             .encryptionConfiguration(TableEncryptionConfigurationArgs.builder()
 *                 .kmsKeyName(cryptoKey.id())
 *                 .build())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject(GetProjectArgs.builder()
 *             .projectId("my-project-name")
 *             .build());
 * 
 *         var encryptRole = new IAMMember("encryptRole", IAMMemberArgs.builder()
 *             .project(project.applyValue(getProjectResult -> getProjectResult.projectId()))
 *             .role("roles/cloudkms.cryptoKeyEncrypterDecrypter")
 *             .member(String.format("serviceAccount:bq-%s{@literal @}bigquery-encryption.iam.gserviceaccount.com", project.applyValue(getProjectResult -> getProjectResult.number())))
 *             .build());
 * 
 *         var job = new Job("job", JobArgs.builder()
 *             .jobId("job_copy")
 *             .copy(JobCopyArgs.builder()
 *                 .sourceTables(                
 *                     JobCopySourceTableArgs.builder()
 *                         .projectId(source[0].project())
 *                         .datasetId(source[0].datasetId())
 *                         .tableId(source[0].tableId())
 *                         .build(),
 *                     JobCopySourceTableArgs.builder()
 *                         .projectId(source[1].project())
 *                         .datasetId(source[1].datasetId())
 *                         .tableId(source[1].tableId())
 *                         .build())
 *                 .destinationTable(JobCopyDestinationTableArgs.builder()
 *                     .projectId(dest.project())
 *                     .datasetId(dest.datasetId())
 *                     .tableId(dest.tableId())
 *                     .build())
 *                 .destinationEncryptionConfiguration(JobCopyDestinationEncryptionConfigurationArgs.builder()
 *                     .kmsKeyName(cryptoKey.id())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Bigquery Job Extract
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.bigquery.Job;
 * import com.pulumi.gcp.bigquery.JobArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobExtractArgs;
 * import com.pulumi.gcp.bigquery.inputs.JobExtractSourceTableArgs;
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
 *         var source_oneDataset = new Dataset("source-oneDataset", DatasetArgs.builder()
 *             .datasetId("job_extract_dataset")
 *             .friendlyName("test")
 *             .description("This is a test description")
 *             .location("US")
 *             .build());
 * 
 *         var source_one = new Table("source-one", TableArgs.builder()
 *             .deletionProtection(false)
 *             .datasetId(source_oneDataset.datasetId())
 *             .tableId("job_extract_table")
 *             .schema("""
 * [
 *   {
 *     "name": "name",
 *     "type": "STRING",
 *     "mode": "NULLABLE"
 *   },
 *   {
 *     "name": "post_abbr",
 *     "type": "STRING",
 *     "mode": "NULLABLE"
 *   },
 *   {
 *     "name": "date",
 *     "type": "DATE",
 *     "mode": "NULLABLE"
 *   }
 * ]
 *             """)
 *             .build());
 * 
 *         var dest = new Bucket("dest", BucketArgs.builder()
 *             .name("job_extract_bucket")
 *             .location("US")
 *             .forceDestroy(true)
 *             .build());
 * 
 *         var job = new Job("job", JobArgs.builder()
 *             .jobId("job_extract")
 *             .extract(JobExtractArgs.builder()
 *                 .destinationUris(dest.url().applyValue(url -> String.format("%s/extract", url)))
 *                 .sourceTable(JobExtractSourceTableArgs.builder()
 *                     .projectId(source_one.project())
 *                     .datasetId(source_one.datasetId())
 *                     .tableId(source_one.tableId())
 *                     .build())
 *                 .destinationFormat("NEWLINE_DELIMITED_JSON")
 *                 .compression("GZIP")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Job can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/jobs/{{job_id}}/location/{{location}}`
 * 
 * * `projects/{{project}}/jobs/{{job_id}}`
 * 
 * * `{{project}}/{{job_id}}/{{location}}`
 * 
 * * `{{job_id}}/{{location}}`
 * 
 * * `{{project}}/{{job_id}}`
 * 
 * * `{{job_id}}`
 * 
 * When using the `pulumi import` command, Job can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}/location/{{location}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}/{{location}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/job:Job default {{job_id}}/{{location}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/job:Job default {{job_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:bigquery/job:Job")
public class Job extends com.pulumi.resources.CustomResource {
    /**
     * Copies a table.
     * 
     */
    @Export(name="copy", refs={JobCopy.class}, tree="[0]")
    private Output</* @Nullable */ JobCopy> copy;

    /**
     * @return Copies a table.
     * 
     */
    public Output<Optional<JobCopy>> copy() {
        return Codegen.optional(this.copy);
    }
    /**
     * (Output)
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return (Output)
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Configures an extract job.
     * 
     */
    @Export(name="extract", refs={JobExtract.class}, tree="[0]")
    private Output</* @Nullable */ JobExtract> extract;

    /**
     * @return Configures an extract job.
     * 
     */
    public Output<Optional<JobExtract>> extract() {
        return Codegen.optional(this.extract);
    }
    /**
     * The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
     * 
     */
    @Export(name="jobId", refs={String.class}, tree="[0]")
    private Output<String> jobId;

    /**
     * @return The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
     * 
     */
    public Output<String> jobId() {
        return this.jobId;
    }
    /**
     * Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
     * 
     */
    @Export(name="jobTimeoutMs", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> jobTimeoutMs;

    /**
     * @return Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
     * 
     */
    public Output<Optional<String>> jobTimeoutMs() {
        return Codegen.optional(this.jobTimeoutMs);
    }
    /**
     * (Output)
     * The type of the job.
     * 
     */
    @Export(name="jobType", refs={String.class}, tree="[0]")
    private Output<String> jobType;

    /**
     * @return (Output)
     * The type of the job.
     * 
     */
    public Output<String> jobType() {
        return this.jobType;
    }
    /**
     * The labels associated with this job. You can use these to organize and group your jobs. **Note**: This field is
     * non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
     * &#39;effective_labels&#39; for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return The labels associated with this job. You can use these to organize and group your jobs. **Note**: This field is
     * non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
     * &#39;effective_labels&#39; for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Configures a load job.
     * 
     */
    @Export(name="load", refs={JobLoad.class}, tree="[0]")
    private Output</* @Nullable */ JobLoad> load;

    /**
     * @return Configures a load job.
     * 
     */
    public Output<Optional<JobLoad>> load() {
        return Codegen.optional(this.load);
    }
    /**
     * Specifies where the error occurred, if present.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> location;

    /**
     * @return Specifies where the error occurred, if present.
     * 
     */
    public Output<Optional<String>> location() {
        return Codegen.optional(this.location);
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }
    /**
     * (Output)
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return (Output)
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Configures a query job.
     * 
     */
    @Export(name="query", refs={JobQuery.class}, tree="[0]")
    private Output</* @Nullable */ JobQuery> query;

    /**
     * @return Configures a query job.
     * 
     */
    public Output<Optional<JobQuery>> query() {
        return Codegen.optional(this.query);
    }
    /**
     * The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
     * Structure is documented below.
     * 
     */
    @Export(name="statuses", refs={List.class,JobStatus.class}, tree="[0,1]")
    private Output<List<JobStatus>> statuses;

    /**
     * @return The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
     * Structure is documented below.
     * 
     */
    public Output<List<JobStatus>> statuses() {
        return this.statuses;
    }
    /**
     * Email address of the user who ran the job.
     * 
     */
    @Export(name="userEmail", refs={String.class}, tree="[0]")
    private Output<String> userEmail;

    /**
     * @return Email address of the user who ran the job.
     * 
     */
    public Output<String> userEmail() {
        return this.userEmail;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Job(String name) {
        this(name, JobArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Job(String name, JobArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Job(String name, JobArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/job:Job", name, args == null ? JobArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Job(String name, Output<String> id, @Nullable JobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/job:Job", name, state, makeResourceOptions(options, id));
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
    public static Job get(String name, Output<String> id, @Nullable JobState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Job(name, id, state, options);
    }
}
