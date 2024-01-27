// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.storage.TransferAgentPoolArgs;
import com.pulumi.gcp.storage.inputs.TransferAgentPoolState;
import com.pulumi.gcp.storage.outputs.TransferAgentPoolBandwidthLimit;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents an On-Premises Agent pool.
 * 
 * To get more information about AgentPool, see:
 * 
 * * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/projects.agentPools)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/storage-transfer/docs/on-prem-agent-pools)
 * 
 * ## Example Usage
 * ### Agent Pool Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.storage.StorageFunctions;
 * import com.pulumi.gcp.storage.inputs.GetTransferProjectServiceAccountArgs;
 * import com.pulumi.gcp.projects.IAMMember;
 * import com.pulumi.gcp.projects.IAMMemberArgs;
 * import com.pulumi.gcp.storage.TransferAgentPool;
 * import com.pulumi.gcp.storage.TransferAgentPoolArgs;
 * import com.pulumi.gcp.storage.inputs.TransferAgentPoolBandwidthLimitArgs;
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
 *         final var default = StorageFunctions.getTransferProjectServiceAccount(GetTransferProjectServiceAccountArgs.builder()
 *             .project(&#34;my-project-name&#34;)
 *             .build());
 * 
 *         var pubsubEditorRole = new IAMMember(&#34;pubsubEditorRole&#34;, IAMMemberArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .role(&#34;roles/pubsub.editor&#34;)
 *             .member(String.format(&#34;serviceAccount:%s&#34;, default_.email()))
 *             .build());
 * 
 *         var example = new TransferAgentPool(&#34;example&#34;, TransferAgentPoolArgs.builder()        
 *             .displayName(&#34;Source A to destination Z&#34;)
 *             .bandwidthLimit(TransferAgentPoolBandwidthLimitArgs.builder()
 *                 .limitMbps(&#34;120&#34;)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(pubsubEditorRole)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AgentPool can be imported using any of these accepted formats* `projects/{{project}}/agentPools/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, AgentPool can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:storage/transferAgentPool:TransferAgentPool default projects/{{project}}/agentPools/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:storage/transferAgentPool:TransferAgentPool default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:storage/transferAgentPool:TransferAgentPool default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:storage/transferAgentPool:TransferAgentPool")
public class TransferAgentPool extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the bandwidth limit details. If this field is unspecified, the default value is set as &#39;No Limit&#39;.
     * Structure is documented below.
     * 
     */
    @Export(name="bandwidthLimit", refs={TransferAgentPoolBandwidthLimit.class}, tree="[0]")
    private Output</* @Nullable */ TransferAgentPoolBandwidthLimit> bandwidthLimit;

    /**
     * @return Specifies the bandwidth limit details. If this field is unspecified, the default value is set as &#39;No Limit&#39;.
     * Structure is documented below.
     * 
     */
    public Output<Optional<TransferAgentPoolBandwidthLimit>> bandwidthLimit() {
        return Codegen.optional(this.bandwidthLimit);
    }
    /**
     * Specifies the client-specified AgentPool description.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Specifies the client-specified AgentPool description.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The ID of the agent pool to create.
     * The agentPoolId must meet the following requirements:
     * * Length of 128 characters or less.
     * * Not start with the string goog.
     * * Start with a lowercase ASCII character, followed by:
     * * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
     * * One or more numerals or lowercase ASCII characters.
     *   As expressed by the regular expression: ^(?!goog)a-z?$.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The ID of the agent pool to create.
     * The agentPoolId must meet the following requirements:
     * * Length of 128 characters or less.
     * * Not start with the string goog.
     * * Start with a lowercase ASCII character, followed by:
     * * Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
     * * One or more numerals or lowercase ASCII characters.
     *   As expressed by the regular expression: ^(?!goog)a-z?$.
     * 
     * ***
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
     * Specifies the state of the AgentPool.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Specifies the state of the AgentPool.
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TransferAgentPool(String name) {
        this(name, TransferAgentPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TransferAgentPool(String name, @Nullable TransferAgentPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TransferAgentPool(String name, @Nullable TransferAgentPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:storage/transferAgentPool:TransferAgentPool", name, args == null ? TransferAgentPoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TransferAgentPool(String name, Output<String> id, @Nullable TransferAgentPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:storage/transferAgentPool:TransferAgentPool", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static TransferAgentPool get(String name, Output<String> id, @Nullable TransferAgentPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TransferAgentPool(name, id, state, options);
    }
}
