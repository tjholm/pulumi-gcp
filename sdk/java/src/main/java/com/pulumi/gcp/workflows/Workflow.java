// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.workflows;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.workflows.WorkflowArgs;
import com.pulumi.gcp.workflows.inputs.WorkflowState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Workflow program to be executed by Workflows.
 * 
 * To get more information about Workflow, see:
 * 
 * * [API documentation](https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows)
 * * How-to Guides
 *     * [Managing Workflows](https://cloud.google.com/workflows/docs/creating-updating-workflow)
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:workflows/workflow:Workflow")
public class Workflow extends com.pulumi.resources.CustomResource {
    /**
     * The timestamp of when the workflow was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The timestamp of when the workflow was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The KMS key used to encrypt workflow and execution data.
     * Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
     * 
     */
    @Export(name="cryptoKeyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cryptoKeyName;

    /**
     * @return The KMS key used to encrypt workflow and execution data.
     * Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
     * 
     */
    public Output<Optional<String>> cryptoKeyName() {
        return Codegen.optional(this.cryptoKeyName);
    }
    /**
     * Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
     * 
     */
    public Output<String> description() {
        return this.description;
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
     * A set of key/value label pairs to assign to this Workflow.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return A set of key/value label pairs to assign to this Workflow.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Name of the Workflow.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the Workflow.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Creates a unique name beginning with the
     * specified prefix. If this and name are unspecified, a random value is chosen for the name.
     * 
     */
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the
     * specified prefix. If this and name are unspecified, a random value is chosen for the name.
     * 
     */
    public Output<String> namePrefix() {
        return this.namePrefix;
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
     * The region of the workflow.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return The region of the workflow.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * The revision of the workflow. A new one is generated if the service account or source contents is changed.
     * 
     */
    @Export(name="revisionId", refs={String.class}, tree="[0]")
    private Output<String> revisionId;

    /**
     * @return The revision of the workflow. A new one is generated if the service account or source contents is changed.
     * 
     */
    public Output<String> revisionId() {
        return this.revisionId;
    }
    /**
     * Name of the service account associated with the latest workflow version. This service
     * account represents the identity of the workflow and determines what permissions the workflow has.
     * Format: projects/{project}/serviceAccounts/{account} or {account}.
     * Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
     * The {account} value can be the email address or the unique_id of the service account.
     * If not provided, workflow will use the project&#39;s default service account.
     * Modifying this field for an existing workflow results in a new workflow revision.
     * 
     */
    @Export(name="serviceAccount", refs={String.class}, tree="[0]")
    private Output<String> serviceAccount;

    /**
     * @return Name of the service account associated with the latest workflow version. This service
     * account represents the identity of the workflow and determines what permissions the workflow has.
     * Format: projects/{project}/serviceAccounts/{account} or {account}.
     * Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
     * The {account} value can be the email address or the unique_id of the service account.
     * If not provided, workflow will use the project&#39;s default service account.
     * Modifying this field for an existing workflow results in a new workflow revision.
     * 
     */
    public Output<String> serviceAccount() {
        return this.serviceAccount;
    }
    /**
     * Workflow code to be executed. The size limit is 32KB.
     * 
     */
    @Export(name="sourceContents", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceContents;

    /**
     * @return Workflow code to be executed. The size limit is 32KB.
     * 
     */
    public Output<Optional<String>> sourceContents() {
        return Codegen.optional(this.sourceContents);
    }
    /**
     * State of the workflow deployment.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the workflow deployment.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The timestamp of when the workflow was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The timestamp of when the workflow was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * User-defined environment variables associated with this workflow revision. This map has a maximum length of 20. Each string can take up to 4KiB. Keys cannot be empty strings and cannot start with “GOOGLE” or “WORKFLOWS&#34;.
     * 
     */
    @Export(name="userEnvVars", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> userEnvVars;

    /**
     * @return User-defined environment variables associated with this workflow revision. This map has a maximum length of 20. Each string can take up to 4KiB. Keys cannot be empty strings and cannot start with “GOOGLE” or “WORKFLOWS&#34;.
     * 
     */
    public Output<Optional<Map<String,String>>> userEnvVars() {
        return Codegen.optional(this.userEnvVars);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Workflow(String name) {
        this(name, WorkflowArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Workflow(String name, @Nullable WorkflowArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Workflow(String name, @Nullable WorkflowArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:workflows/workflow:Workflow", name, args == null ? WorkflowArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Workflow(String name, Output<String> id, @Nullable WorkflowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:workflows/workflow:Workflow", name, state, makeResourceOptions(options, id));
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
    public static Workflow get(String name, Output<String> id, @Nullable WorkflowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Workflow(name, id, state, options);
    }
}
