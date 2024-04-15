// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vertex.AiFeatureStoreEntityTypeIamPolicyArgs;
import com.pulumi.gcp.vertex.inputs.AiFeatureStoreEntityTypeIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms:
 * 
 * * {{featurestore}}/entityTypes/{{name}}
 * 
 * * {{name}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * Vertex AI featurestoreentitytype IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy editor &#34;{{featurestore}}/entityTypes/{{featurestore_entitytype}} roles/viewer user:jane@example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy editor &#34;{{featurestore}}/entityTypes/{{featurestore_entitytype}} roles/viewer&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy editor {{featurestore}}/entityTypes/{{featurestore_entitytype}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy")
public class AiFeatureStoreEntityTypeIamPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="entitytype", refs={String.class}, tree="[0]")
    private Output<String> entitytype;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> entitytype() {
        return this.entitytype;
    }
    /**
     * (Computed) The etag of the IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="featurestore", refs={String.class}, tree="[0]")
    private Output<String> featurestore;

    /**
     * @return The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}. Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> featurestore() {
        return this.featurestore;
    }
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Export(name="policyData", refs={String.class}, tree="[0]")
    private Output<String> policyData;

    /**
     * @return The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    public Output<String> policyData() {
        return this.policyData;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AiFeatureStoreEntityTypeIamPolicy(String name) {
        this(name, AiFeatureStoreEntityTypeIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AiFeatureStoreEntityTypeIamPolicy(String name, AiFeatureStoreEntityTypeIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AiFeatureStoreEntityTypeIamPolicy(String name, AiFeatureStoreEntityTypeIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy", name, args == null ? AiFeatureStoreEntityTypeIamPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AiFeatureStoreEntityTypeIamPolicy(String name, Output<String> id, @Nullable AiFeatureStoreEntityTypeIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiFeatureStoreEntityTypeIamPolicy:AiFeatureStoreEntityTypeIamPolicy", name, state, makeResourceOptions(options, id));
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
    public static AiFeatureStoreEntityTypeIamPolicy get(String name, Output<String> id, @Nullable AiFeatureStoreEntityTypeIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AiFeatureStoreEntityTypeIamPolicy(name, id, state, options);
    }
}
