// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigee;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigee.EnvReferencesArgs;
import com.pulumi.gcp.apigee.inputs.EnvReferencesState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An `Environment Reference` in Apigee.
 * 
 * To get more information about EnvReferences, see:
 * 
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.references/create)
 * * How-to Guides
 *     * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
 * 
 * ## Import
 * 
 * EnvReferences can be imported using any of these accepted formats:
 * 
 * * `{{env_id}}/references/{{name}}`
 * 
 * * `{{env_id}}/{{name}}`
 * 
 * When using the `pulumi import` command, EnvReferences can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:apigee/envReferences:EnvReferences default {{env_id}}/references/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:apigee/envReferences:EnvReferences default {{env_id}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:apigee/envReferences:EnvReferences")
public class EnvReferences extends com.pulumi.resources.CustomResource {
    /**
     * Optional. A human-readable description of this reference.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional. A human-readable description of this reference.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Apigee environment group associated with the Apigee environment,
     * in the format `organizations/{{org_name}}/environments/{{env_name}}`.
     * 
     * ***
     * 
     */
    @Export(name="envId", refs={String.class}, tree="[0]")
    private Output<String> envId;

    /**
     * @return The Apigee environment group associated with the Apigee environment,
     * in the format `organizations/{{org_name}}/environments/{{env_name}}`.
     * 
     * ***
     * 
     */
    public Output<String> envId() {
        return this.envId;
    }
    /**
     * Required. The resource id of this reference. Values must match the regular expression [\w\s-.]+.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Required. The resource id of this reference. Values must match the regular expression [\w\s-.]+.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Required. The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resourceType.
     * 
     */
    @Export(name="refers", refs={String.class}, tree="[0]")
    private Output<String> refers;

    /**
     * @return Required. The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resourceType.
     * 
     */
    public Output<String> refers() {
        return this.refers;
    }
    /**
     * The type of resource referred to by this reference. Valid values are &#39;KeyStore&#39; or &#39;TrustStore&#39;.
     * 
     */
    @Export(name="resourceType", refs={String.class}, tree="[0]")
    private Output<String> resourceType;

    /**
     * @return The type of resource referred to by this reference. Valid values are &#39;KeyStore&#39; or &#39;TrustStore&#39;.
     * 
     */
    public Output<String> resourceType() {
        return this.resourceType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EnvReferences(java.lang.String name) {
        this(name, EnvReferencesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnvReferences(java.lang.String name, EnvReferencesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnvReferences(java.lang.String name, EnvReferencesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/envReferences:EnvReferences", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EnvReferences(java.lang.String name, Output<java.lang.String> id, @Nullable EnvReferencesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/envReferences:EnvReferences", name, state, makeResourceOptions(options, id), false);
    }

    private static EnvReferencesArgs makeArgs(EnvReferencesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EnvReferencesArgs.Empty : args;
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
    public static EnvReferences get(java.lang.String name, Output<java.lang.String> id, @Nullable EnvReferencesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnvReferences(name, id, state, options);
    }
}
