// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vertex.AiEndpointIamPolicyArgs;
import com.pulumi.gcp.vertex.inputs.AiEndpointIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

@ResourceType(type="gcp:vertex/aiEndpointIamPolicy:AiEndpointIamPolicy")
public class AiEndpointIamPolicy extends com.pulumi.resources.CustomResource {
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    public Output<String> endpoint() {
        return this.endpoint;
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    public Output<String> location() {
        return this.location;
    }
    @Export(name="policyData", refs={String.class}, tree="[0]")
    private Output<String> policyData;

    public Output<String> policyData() {
        return this.policyData;
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AiEndpointIamPolicy(java.lang.String name) {
        this(name, AiEndpointIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AiEndpointIamPolicy(java.lang.String name, AiEndpointIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AiEndpointIamPolicy(java.lang.String name, AiEndpointIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiEndpointIamPolicy:AiEndpointIamPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AiEndpointIamPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable AiEndpointIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiEndpointIamPolicy:AiEndpointIamPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static AiEndpointIamPolicyArgs makeArgs(AiEndpointIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AiEndpointIamPolicyArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static AiEndpointIamPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable AiEndpointIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AiEndpointIamPolicy(name, id, state, options);
    }
}
