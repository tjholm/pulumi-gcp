// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudfunctions;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudfunctions.FunctionArgs;
import com.pulumi.gcp.cloudfunctions.inputs.FunctionState;
import com.pulumi.gcp.cloudfunctions.outputs.FunctionEventTrigger;
import com.pulumi.gcp.cloudfunctions.outputs.FunctionSecretEnvironmentVariable;
import com.pulumi.gcp.cloudfunctions.outputs.FunctionSecretVolume;
import com.pulumi.gcp.cloudfunctions.outputs.FunctionSourceRepository;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a new Cloud Function. For more information see:
 * 
 * * [API documentation](https://cloud.google.com/functions/docs/reference/rest/v1/projects.locations.functions)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/functions/docs)
 * 
 * &gt; **Warning:** As of November 1, 2019, newly created Functions are
 * private-by-default and will require [appropriate IAM permissions](https://cloud.google.com/functions/docs/reference/iam/roles)
 * to be invoked. See below examples for how to set up the appropriate permissions,
 * or view the [Cloud Functions IAM resources](https://www.terraform.io/docs/providers/google/r/cloudfunctions_cloud_function_iam.html)
 * for Cloud Functions.
 * 
 * ## Example Usage
 * 
 * ### Public Function
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
 * import com.pulumi.gcp.cloudfunctions.Function;
 * import com.pulumi.gcp.cloudfunctions.FunctionArgs;
 * import com.pulumi.gcp.cloudfunctions.FunctionIamMember;
 * import com.pulumi.gcp.cloudfunctions.FunctionIamMemberArgs;
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
 *         var bucket = new Bucket("bucket", BucketArgs.builder()
 *             .name("test-bucket")
 *             .location("US")
 *             .build());
 * 
 *         var archive = new BucketObject("archive", BucketObjectArgs.builder()
 *             .name("index.zip")
 *             .bucket(bucket.name())
 *             .source(new FileAsset("./path/to/zip/file/which/contains/code"))
 *             .build());
 * 
 *         var function = new Function("function", FunctionArgs.builder()
 *             .name("function-test")
 *             .description("My function")
 *             .runtime("nodejs16")
 *             .availableMemoryMb(128)
 *             .sourceArchiveBucket(bucket.name())
 *             .sourceArchiveObject(archive.name())
 *             .triggerHttp(true)
 *             .entryPoint("helloGET")
 *             .build());
 * 
 *         // IAM entry for all users to invoke the function
 *         var invoker = new FunctionIamMember("invoker", FunctionIamMemberArgs.builder()
 *             .project(function.project())
 *             .region(function.region())
 *             .cloudFunction(function.name())
 *             .role("roles/cloudfunctions.invoker")
 *             .member("allUsers")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Single User
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
 * import com.pulumi.gcp.cloudfunctions.Function;
 * import com.pulumi.gcp.cloudfunctions.FunctionArgs;
 * import com.pulumi.gcp.cloudfunctions.FunctionIamMember;
 * import com.pulumi.gcp.cloudfunctions.FunctionIamMemberArgs;
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
 *         var bucket = new Bucket("bucket", BucketArgs.builder()
 *             .name("test-bucket")
 *             .location("US")
 *             .build());
 * 
 *         var archive = new BucketObject("archive", BucketObjectArgs.builder()
 *             .name("index.zip")
 *             .bucket(bucket.name())
 *             .source(new FileAsset("./path/to/zip/file/which/contains/code"))
 *             .build());
 * 
 *         var function = new Function("function", FunctionArgs.builder()
 *             .name("function-test")
 *             .description("My function")
 *             .runtime("nodejs16")
 *             .availableMemoryMb(128)
 *             .sourceArchiveBucket(bucket.name())
 *             .sourceArchiveObject(archive.name())
 *             .triggerHttp(true)
 *             .httpsTriggerSecurityLevel("SECURE_ALWAYS")
 *             .timeout(60)
 *             .entryPoint("helloGET")
 *             .labels(Map.of("my-label", "my-label-value"))
 *             .environmentVariables(Map.of("MY_ENV_VAR", "my-env-var-value"))
 *             .build());
 * 
 *         // IAM entry for a single user to invoke the function
 *         var invoker = new FunctionIamMember("invoker", FunctionIamMemberArgs.builder()
 *             .project(function.project())
 *             .region(function.region())
 *             .cloudFunction(function.name())
 *             .role("roles/cloudfunctions.invoker")
 *             .member("user:myFunctionInvoker{@literal @}example.com")
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
 * Functions can be imported using the `name` or `{{project}}/{{region}}/name`, e.g.
 * 
 * * `{{project}}/{{region}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, Functions can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:cloudfunctions/function:Function default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:cloudfunctions/function:Function default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:cloudfunctions/function:Function")
public class Function extends com.pulumi.resources.CustomResource {
    /**
     * Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
     * 
     */
    @Export(name="availableMemoryMb", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> availableMemoryMb;

    /**
     * @return Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
     * 
     */
    public Output<Optional<Integer>> availableMemoryMb() {
        return Codegen.optional(this.availableMemoryMb);
    }
    /**
     * A set of key/value environment variable pairs available during build time.
     * 
     */
    @Export(name="buildEnvironmentVariables", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> buildEnvironmentVariables;

    /**
     * @return A set of key/value environment variable pairs available during build time.
     * 
     */
    public Output<Optional<Map<String,Object>>> buildEnvironmentVariables() {
        return Codegen.optional(this.buildEnvironmentVariables);
    }
    /**
     * Name of the Cloud Build Custom Worker Pool that should be used to build the function.
     * 
     */
    @Export(name="buildWorkerPool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> buildWorkerPool;

    /**
     * @return Name of the Cloud Build Custom Worker Pool that should be used to build the function.
     * 
     */
    public Output<Optional<String>> buildWorkerPool() {
        return Codegen.optional(this.buildWorkerPool);
    }
    /**
     * Description of the function.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the function.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Docker Registry to use for storing the function&#39;s Docker images. Allowed values are ARTIFACT_REGISTRY (default) and CONTAINER_REGISTRY.
     * 
     */
    @Export(name="dockerRegistry", refs={String.class}, tree="[0]")
    private Output<String> dockerRegistry;

    /**
     * @return Docker Registry to use for storing the function&#39;s Docker images. Allowed values are ARTIFACT_REGISTRY (default) and CONTAINER_REGISTRY.
     * 
     */
    public Output<String> dockerRegistry() {
        return this.dockerRegistry;
    }
    /**
     * User-managed repository created in Artifact Registry to which the function&#39;s Docker image will be pushed after it is built by Cloud Build. May optionally be encrypted with a customer-managed encryption key (CMEK). If unspecified and `docker_registry` is not explicitly set to `CONTAINER_REGISTRY`, GCF will create and use a default Artifact Registry repository named &#39;gcf-artifacts&#39; in the region.
     * 
     */
    @Export(name="dockerRepository", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dockerRepository;

    /**
     * @return User-managed repository created in Artifact Registry to which the function&#39;s Docker image will be pushed after it is built by Cloud Build. May optionally be encrypted with a customer-managed encryption key (CMEK). If unspecified and `docker_registry` is not explicitly set to `CONTAINER_REGISTRY`, GCF will create and use a default Artifact Registry repository named &#39;gcf-artifacts&#39; in the region.
     * 
     */
    public Output<Optional<String>> dockerRepository() {
        return Codegen.optional(this.dockerRepository);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Name of the function that will be executed when the Google Cloud Function is triggered.
     * 
     */
    @Export(name="entryPoint", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> entryPoint;

    /**
     * @return Name of the function that will be executed when the Google Cloud Function is triggered.
     * 
     */
    public Output<Optional<String>> entryPoint() {
        return Codegen.optional(this.entryPoint);
    }
    /**
     * A set of key/value environment variable pairs to assign to the function.
     * 
     */
    @Export(name="environmentVariables", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> environmentVariables;

    /**
     * @return A set of key/value environment variable pairs to assign to the function.
     * 
     */
    public Output<Optional<Map<String,Object>>> environmentVariables() {
        return Codegen.optional(this.environmentVariables);
    }
    /**
     * A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `trigger_http`.
     * 
     */
    @Export(name="eventTrigger", refs={FunctionEventTrigger.class}, tree="[0]")
    private Output<FunctionEventTrigger> eventTrigger;

    /**
     * @return A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `trigger_http`.
     * 
     */
    public Output<FunctionEventTrigger> eventTrigger() {
        return this.eventTrigger;
    }
    /**
     * The security level for the function. The following options are available:
     * 
     * * `SECURE_ALWAYS` Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path. Query parameters are reserved for the redirect.
     * * `SECURE_OPTIONAL` Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly.
     * 
     */
    @Export(name="httpsTriggerSecurityLevel", refs={String.class}, tree="[0]")
    private Output<String> httpsTriggerSecurityLevel;

    /**
     * @return The security level for the function. The following options are available:
     * 
     * * `SECURE_ALWAYS` Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path. Query parameters are reserved for the redirect.
     * * `SECURE_OPTIONAL` Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly.
     * 
     */
    public Output<String> httpsTriggerSecurityLevel() {
        return this.httpsTriggerSecurityLevel;
    }
    /**
     * URL which triggers function execution. Returned only if `trigger_http` is used.
     * 
     */
    @Export(name="httpsTriggerUrl", refs={String.class}, tree="[0]")
    private Output<String> httpsTriggerUrl;

    /**
     * @return URL which triggers function execution. Returned only if `trigger_http` is used.
     * 
     */
    public Output<String> httpsTriggerUrl() {
        return this.httpsTriggerUrl;
    }
    /**
     * String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
     * 
     */
    @Export(name="ingressSettings", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ingressSettings;

    /**
     * @return String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
     * 
     */
    public Output<Optional<String>> ingressSettings() {
        return Codegen.optional(this.ingressSettings);
    }
    /**
     * Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
     * If specified, you must also provide an artifact registry repository using the `docker_repository` field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
     * 
     */
    @Export(name="kmsKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyName;

    /**
     * @return Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
     * If specified, you must also provide an artifact registry repository using the `docker_repository` field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
     * 
     */
    public Output<Optional<String>> kmsKeyName() {
        return Codegen.optional(this.kmsKeyName);
    }
    /**
     * A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field &#39;effective_labels&#39; for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> labels;

    /**
     * @return A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field &#39;effective_labels&#39; for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     * 
     */
    @Export(name="maxInstances", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxInstances;

    /**
     * @return The limit on the maximum number of function instances that may coexist at a given time.
     * 
     */
    public Output<Integer> maxInstances() {
        return this.maxInstances;
    }
    /**
     * The limit on the minimum number of function instances that may coexist at a given time.
     * 
     */
    @Export(name="minInstances", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minInstances;

    /**
     * @return The limit on the minimum number of function instances that may coexist at a given time.
     * 
     */
    public Output<Optional<Integer>> minInstances() {
        return Codegen.optional(this.minInstances);
    }
    /**
     * A user-defined name of the function. Function names must be unique globally.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A user-defined name of the function. Function names must be unique globally.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Project of the function. If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return Project of the function. If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Region of function. If it is not provided, the provider region is used.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Region of function. If it is not provided, the provider region is used.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The runtime in which the function is going to run.
     * Eg. `&#34;nodejs16&#34;`, `&#34;python39&#34;`, `&#34;dotnet3&#34;`, `&#34;go116&#34;`, `&#34;java11&#34;`, `&#34;ruby30&#34;`, `&#34;php74&#34;`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
     * 
     * ***
     * 
     */
    @Export(name="runtime", refs={String.class}, tree="[0]")
    private Output<String> runtime;

    /**
     * @return The runtime in which the function is going to run.
     * Eg. `&#34;nodejs16&#34;`, `&#34;python39&#34;`, `&#34;dotnet3&#34;`, `&#34;go116&#34;`, `&#34;java11&#34;`, `&#34;ruby30&#34;`, `&#34;php74&#34;`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
     * 
     * ***
     * 
     */
    public Output<String> runtime() {
        return this.runtime;
    }
    /**
     * Secret environment variables configuration. Structure is documented below.
     * 
     */
    @Export(name="secretEnvironmentVariables", refs={List.class,FunctionSecretEnvironmentVariable.class}, tree="[0,1]")
    private Output</* @Nullable */ List<FunctionSecretEnvironmentVariable>> secretEnvironmentVariables;

    /**
     * @return Secret environment variables configuration. Structure is documented below.
     * 
     */
    public Output<Optional<List<FunctionSecretEnvironmentVariable>>> secretEnvironmentVariables() {
        return Codegen.optional(this.secretEnvironmentVariables);
    }
    /**
     * Secret volumes configuration. Structure is documented below.
     * 
     */
    @Export(name="secretVolumes", refs={List.class,FunctionSecretVolume.class}, tree="[0,1]")
    private Output</* @Nullable */ List<FunctionSecretVolume>> secretVolumes;

    /**
     * @return Secret volumes configuration. Structure is documented below.
     * 
     */
    public Output<Optional<List<FunctionSecretVolume>>> secretVolumes() {
        return Codegen.optional(this.secretVolumes);
    }
    /**
     * If provided, the self-provided service account to run the function with.
     * 
     */
    @Export(name="serviceAccountEmail", refs={String.class}, tree="[0]")
    private Output<String> serviceAccountEmail;

    /**
     * @return If provided, the self-provided service account to run the function with.
     * 
     */
    public Output<String> serviceAccountEmail() {
        return this.serviceAccountEmail;
    }
    /**
     * The GCS bucket containing the zip archive which contains the function.
     * 
     */
    @Export(name="sourceArchiveBucket", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceArchiveBucket;

    /**
     * @return The GCS bucket containing the zip archive which contains the function.
     * 
     */
    public Output<Optional<String>> sourceArchiveBucket() {
        return Codegen.optional(this.sourceArchiveBucket);
    }
    /**
     * The source archive object (file) in archive bucket.
     * 
     */
    @Export(name="sourceArchiveObject", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceArchiveObject;

    /**
     * @return The source archive object (file) in archive bucket.
     * 
     */
    public Output<Optional<String>> sourceArchiveObject() {
        return Codegen.optional(this.sourceArchiveObject);
    }
    /**
     * Represents parameters related to source repository where a function is hosted.
     * Cannot be set alongside `source_archive_bucket` or `source_archive_object`. Structure is documented below. It must match the pattern `projects/{project}/locations/{location}/repositories/{repository}`.*
     * 
     */
    @Export(name="sourceRepository", refs={FunctionSourceRepository.class}, tree="[0]")
    private Output</* @Nullable */ FunctionSourceRepository> sourceRepository;

    /**
     * @return Represents parameters related to source repository where a function is hosted.
     * Cannot be set alongside `source_archive_bucket` or `source_archive_object`. Structure is documented below. It must match the pattern `projects/{project}/locations/{location}/repositories/{repository}`.*
     * 
     */
    public Output<Optional<FunctionSourceRepository>> sourceRepository() {
        return Codegen.optional(this.sourceRepository);
    }
    /**
     * Describes the current stage of a deployment.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Describes the current stage of a deployment.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeout;

    /**
     * @return Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
     * 
     */
    public Output<Optional<Integer>> timeout() {
        return Codegen.optional(this.timeout);
    }
    /**
     * Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `https_trigger_url`. Cannot be used with `event_trigger`.
     * 
     */
    @Export(name="triggerHttp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> triggerHttp;

    /**
     * @return Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `https_trigger_url`. Cannot be used with `event_trigger`.
     * 
     */
    public Output<Optional<Boolean>> triggerHttp() {
        return Codegen.optional(this.triggerHttp);
    }
    /**
     * The version identifier of the Cloud Function. Each deployment attempt results in a new version of a function being
     * created.
     * 
     */
    @Export(name="versionId", refs={String.class}, tree="[0]")
    private Output<String> versionId;

    /**
     * @return The version identifier of the Cloud Function. Each deployment attempt results in a new version of a function being
     * created.
     * 
     */
    public Output<String> versionId() {
        return this.versionId;
    }
    /**
     * The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*{@literal /}locations/*{@literal /}connectors/*`.
     * 
     */
    @Export(name="vpcConnector", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpcConnector;

    /**
     * @return The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*{@literal /}locations/*{@literal /}connectors/*`.
     * 
     */
    public Output<Optional<String>> vpcConnector() {
        return Codegen.optional(this.vpcConnector);
    }
    /**
     * The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
     * 
     */
    @Export(name="vpcConnectorEgressSettings", refs={String.class}, tree="[0]")
    private Output<String> vpcConnectorEgressSettings;

    /**
     * @return The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
     * 
     */
    public Output<String> vpcConnectorEgressSettings() {
        return this.vpcConnectorEgressSettings;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Function(String name) {
        this(name, FunctionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Function(String name, FunctionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Function(String name, FunctionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudfunctions/function:Function", name, args == null ? FunctionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Function(String name, Output<String> id, @Nullable FunctionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudfunctions/function:Function", name, state, makeResourceOptions(options, id));
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
    public static Function get(String name, Output<String> id, @Nullable FunctionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Function(name, id, state, options);
    }
}
