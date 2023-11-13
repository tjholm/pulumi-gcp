// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudfunctionsv2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudfunctionsv2.FunctionArgs;
import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionState;
import com.pulumi.gcp.cloudfunctionsv2.outputs.FunctionBuildConfig;
import com.pulumi.gcp.cloudfunctionsv2.outputs.FunctionEventTrigger;
import com.pulumi.gcp.cloudfunctionsv2.outputs.FunctionServiceConfig;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Cloud Function that contains user computation executed in response to an event.
 * 
 * To get more information about function, see:
 * 
 * * [API documentation](https://cloud.google.com/functions/docs/reference/rest/v2beta/projects.locations.functions)
 * 
 * ## Example Usage
 * ### Cloudfunctions2 Basic Gcs
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.storage.BucketObject;
 * import com.pulumi.gcp.storage.BucketObjectArgs;
 * import com.pulumi.gcp.storage.StorageFunctions;
 * import com.pulumi.gcp.storage.inputs.GetProjectServiceAccountArgs;
 * import com.pulumi.gcp.projects.IAMMember;
 * import com.pulumi.gcp.projects.IAMMemberArgs;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.Function;
 * import com.pulumi.gcp.cloudfunctionsv2.FunctionArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionBuildConfigArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionBuildConfigSourceArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionBuildConfigSourceStorageSourceArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionServiceConfigArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionEventTriggerArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var source_bucket = new Bucket(&#34;source-bucket&#34;, BucketArgs.builder()        
 *             .location(&#34;US&#34;)
 *             .uniformBucketLevelAccess(true)
 *             .build());
 * 
 *         var object = new BucketObject(&#34;object&#34;, BucketObjectArgs.builder()        
 *             .bucket(source_bucket.name())
 *             .source(new FileAsset(&#34;function-source.zip&#34;))
 *             .build());
 * 
 *         var trigger_bucket = new Bucket(&#34;trigger-bucket&#34;, BucketArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .uniformBucketLevelAccess(true)
 *             .build());
 * 
 *         final var gcsAccount = StorageFunctions.getProjectServiceAccount();
 * 
 *         var gcs_pubsub_publishing = new IAMMember(&#34;gcs-pubsub-publishing&#34;, IAMMemberArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .role(&#34;roles/pubsub.publisher&#34;)
 *             .member(String.format(&#34;serviceAccount:%s&#34;, gcsAccount.applyValue(getProjectServiceAccountResult -&gt; getProjectServiceAccountResult.emailAddress())))
 *             .build());
 * 
 *         var account = new Account(&#34;account&#34;, AccountArgs.builder()        
 *             .accountId(&#34;gcf-sa&#34;)
 *             .displayName(&#34;Test Service Account - used for both the cloud function and eventarc trigger in the test&#34;)
 *             .build());
 * 
 *         var invoking = new IAMMember(&#34;invoking&#34;, IAMMemberArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .role(&#34;roles/run.invoker&#34;)
 *             .member(account.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(gcs_pubsub_publishing)
 *                 .build());
 * 
 *         var event_receiving = new IAMMember(&#34;event-receiving&#34;, IAMMemberArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .role(&#34;roles/eventarc.eventReceiver&#34;)
 *             .member(account.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(invoking)
 *                 .build());
 * 
 *         var artifactregistry_reader = new IAMMember(&#34;artifactregistry-reader&#34;, IAMMemberArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .role(&#34;roles/artifactregistry.reader&#34;)
 *             .member(account.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(event_receiving)
 *                 .build());
 * 
 *         var function = new Function(&#34;function&#34;, FunctionArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .description(&#34;a new function&#34;)
 *             .buildConfig(FunctionBuildConfigArgs.builder()
 *                 .runtime(&#34;nodejs12&#34;)
 *                 .entryPoint(&#34;entryPoint&#34;)
 *                 .environmentVariables(Map.of(&#34;BUILD_CONFIG_TEST&#34;, &#34;build_test&#34;))
 *                 .source(FunctionBuildConfigSourceArgs.builder()
 *                     .storageSource(FunctionBuildConfigSourceStorageSourceArgs.builder()
 *                         .bucket(source_bucket.name())
 *                         .object(object.name())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .serviceConfig(FunctionServiceConfigArgs.builder()
 *                 .maxInstanceCount(3)
 *                 .minInstanceCount(1)
 *                 .availableMemory(&#34;256M&#34;)
 *                 .timeoutSeconds(60)
 *                 .environmentVariables(Map.of(&#34;SERVICE_CONFIG_TEST&#34;, &#34;config_test&#34;))
 *                 .ingressSettings(&#34;ALLOW_INTERNAL_ONLY&#34;)
 *                 .allTrafficOnLatestRevision(true)
 *                 .serviceAccountEmail(account.email())
 *                 .build())
 *             .eventTrigger(FunctionEventTriggerArgs.builder()
 *                 .triggerRegion(&#34;us-central1&#34;)
 *                 .eventType(&#34;google.cloud.storage.object.v1.finalized&#34;)
 *                 .retryPolicy(&#34;RETRY_POLICY_RETRY&#34;)
 *                 .serviceAccountEmail(account.email())
 *                 .eventFilters(FunctionEventTriggerEventFilterArgs.builder()
 *                     .attribute(&#34;bucket&#34;)
 *                     .value(trigger_bucket.name())
 *                     .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     event_receiving,
 *                     artifactregistry_reader)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Cloudfunctions2 Basic Auditlogs
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.storage.BucketObject;
 * import com.pulumi.gcp.storage.BucketObjectArgs;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.projects.IAMMember;
 * import com.pulumi.gcp.projects.IAMMemberArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.Function;
 * import com.pulumi.gcp.cloudfunctionsv2.FunctionArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionBuildConfigArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionBuildConfigSourceArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionBuildConfigSourceStorageSourceArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionServiceConfigArgs;
 * import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionEventTriggerArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var source_bucket = new Bucket(&#34;source-bucket&#34;, BucketArgs.builder()        
 *             .location(&#34;US&#34;)
 *             .uniformBucketLevelAccess(true)
 *             .build());
 * 
 *         var object = new BucketObject(&#34;object&#34;, BucketObjectArgs.builder()        
 *             .bucket(source_bucket.name())
 *             .source(new FileAsset(&#34;function-source.zip&#34;))
 *             .build());
 * 
 *         var account = new Account(&#34;account&#34;, AccountArgs.builder()        
 *             .accountId(&#34;gcf-sa&#34;)
 *             .displayName(&#34;Test Service Account - used for both the cloud function and eventarc trigger in the test&#34;)
 *             .build());
 * 
 *         var audit_log_bucket = new Bucket(&#34;audit-log-bucket&#34;, BucketArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .uniformBucketLevelAccess(true)
 *             .build());
 * 
 *         var invoking = new IAMMember(&#34;invoking&#34;, IAMMemberArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .role(&#34;roles/run.invoker&#34;)
 *             .member(account.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build());
 * 
 *         var event_receiving = new IAMMember(&#34;event-receiving&#34;, IAMMemberArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .role(&#34;roles/eventarc.eventReceiver&#34;)
 *             .member(account.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(invoking)
 *                 .build());
 * 
 *         var artifactregistry_reader = new IAMMember(&#34;artifactregistry-reader&#34;, IAMMemberArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .role(&#34;roles/artifactregistry.reader&#34;)
 *             .member(account.email().applyValue(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(event_receiving)
 *                 .build());
 * 
 *         var function = new Function(&#34;function&#34;, FunctionArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .description(&#34;a new function&#34;)
 *             .buildConfig(FunctionBuildConfigArgs.builder()
 *                 .runtime(&#34;nodejs12&#34;)
 *                 .entryPoint(&#34;entryPoint&#34;)
 *                 .environmentVariables(Map.of(&#34;BUILD_CONFIG_TEST&#34;, &#34;build_test&#34;))
 *                 .source(FunctionBuildConfigSourceArgs.builder()
 *                     .storageSource(FunctionBuildConfigSourceStorageSourceArgs.builder()
 *                         .bucket(source_bucket.name())
 *                         .object(object.name())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .serviceConfig(FunctionServiceConfigArgs.builder()
 *                 .maxInstanceCount(3)
 *                 .minInstanceCount(1)
 *                 .availableMemory(&#34;256M&#34;)
 *                 .timeoutSeconds(60)
 *                 .environmentVariables(Map.of(&#34;SERVICE_CONFIG_TEST&#34;, &#34;config_test&#34;))
 *                 .ingressSettings(&#34;ALLOW_INTERNAL_ONLY&#34;)
 *                 .allTrafficOnLatestRevision(true)
 *                 .serviceAccountEmail(account.email())
 *                 .build())
 *             .eventTrigger(FunctionEventTriggerArgs.builder()
 *                 .triggerRegion(&#34;us-central1&#34;)
 *                 .eventType(&#34;google.cloud.audit.log.v1.written&#34;)
 *                 .retryPolicy(&#34;RETRY_POLICY_RETRY&#34;)
 *                 .serviceAccountEmail(account.email())
 *                 .eventFilters(                
 *                     FunctionEventTriggerEventFilterArgs.builder()
 *                         .attribute(&#34;serviceName&#34;)
 *                         .value(&#34;storage.googleapis.com&#34;)
 *                         .build(),
 *                     FunctionEventTriggerEventFilterArgs.builder()
 *                         .attribute(&#34;methodName&#34;)
 *                         .value(&#34;storage.objects.create&#34;)
 *                         .build(),
 *                     FunctionEventTriggerEventFilterArgs.builder()
 *                         .attribute(&#34;resourceName&#34;)
 *                         .value(audit_log_bucket.name().applyValue(name -&gt; String.format(&#34;/projects/_/buckets/%s/objects/*.txt&#34;, name)))
 *                         .operator(&#34;match-path-pattern&#34;)
 *                         .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     event_receiving,
 *                     artifactregistry_reader)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * function can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/function:Function default projects/{{project}}/locations/{{location}}/functions/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/function:Function default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/function:Function default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:cloudfunctionsv2/function:Function")
public class Function extends com.pulumi.resources.CustomResource {
    /**
     * Describes the Build step of the function that builds a container
     * from the given source.
     * Structure is documented below.
     * 
     */
    @Export(name="buildConfig", refs={FunctionBuildConfig.class}, tree="[0]")
    private Output</* @Nullable */ FunctionBuildConfig> buildConfig;

    /**
     * @return Describes the Build step of the function that builds a container
     * from the given source.
     * Structure is documented below.
     * 
     */
    public Output<Optional<FunctionBuildConfig>> buildConfig() {
        return Codegen.optional(this.buildConfig);
    }
    /**
     * User-provided description of a function.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User-provided description of a function.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * The environment the function is hosted on.
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output<String> environment;

    /**
     * @return The environment the function is hosted on.
     * 
     */
    public Output<String> environment() {
        return this.environment;
    }
    /**
     * An Eventarc trigger managed by Google Cloud Functions that fires events in
     * response to a condition in another service.
     * Structure is documented below.
     * 
     */
    @Export(name="eventTrigger", refs={FunctionEventTrigger.class}, tree="[0]")
    private Output</* @Nullable */ FunctionEventTrigger> eventTrigger;

    /**
     * @return An Eventarc trigger managed by Google Cloud Functions that fires events in
     * response to a condition in another service.
     * Structure is documented below.
     * 
     */
    public Output<Optional<FunctionEventTrigger>> eventTrigger() {
        return Codegen.optional(this.eventTrigger);
    }
    /**
     * Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
     * It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
     * 
     */
    @Export(name="kmsKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> kmsKeyName;

    /**
     * @return Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
     * It must match the pattern projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}.
     * 
     */
    public Output<Optional<String>> kmsKeyName() {
        return Codegen.optional(this.kmsKeyName);
    }
    /**
     * A set of key/value label pairs associated with this Cloud Function.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs associated with this Cloud Function.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location of this cloud function.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location of this cloud function.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * A user-defined name of the function. Function names must
     * be unique globally and match pattern `projects/*{@literal /}locations/*{@literal /}functions/*`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A user-defined name of the function. Function names must
     * be unique globally and match pattern `projects/*{@literal /}locations/*{@literal /}functions/*`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
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
     * Describes the Service being deployed.
     * Structure is documented below.
     * 
     */
    @Export(name="serviceConfig", refs={FunctionServiceConfig.class}, tree="[0]")
    private Output</* @Nullable */ FunctionServiceConfig> serviceConfig;

    /**
     * @return Describes the Service being deployed.
     * Structure is documented below.
     * 
     */
    public Output<Optional<FunctionServiceConfig>> serviceConfig() {
        return Codegen.optional(this.serviceConfig);
    }
    /**
     * Describes the current state of the function.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Describes the current state of the function.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The last update timestamp of a Cloud Function.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The last update timestamp of a Cloud Function.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Output only. The deployed url for the function.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return Output only. The deployed url for the function.
     * 
     */
    public Output<String> url() {
        return this.url;
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
        super("gcp:cloudfunctionsv2/function:Function", name, args == null ? FunctionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Function(String name, Output<String> id, @Nullable FunctionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudfunctionsv2/function:Function", name, state, makeResourceOptions(options, id));
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
