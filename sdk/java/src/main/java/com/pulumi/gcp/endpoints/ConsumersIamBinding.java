// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.endpoints;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.endpoints.ConsumersIamBindingArgs;
import com.pulumi.gcp.endpoints.inputs.ConsumersIamBindingState;
import com.pulumi.gcp.endpoints.outputs.ConsumersIamBindingCondition;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Cloud Endpoints ServiceConsumers. Each of these resources serves a different use case:
 * 
 * * `gcp.endpoints.ConsumersIamPolicy`: Authoritative. Sets the IAM policy for the serviceconsumers and replaces any existing policy already attached.
 * * `gcp.endpoints.ConsumersIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the serviceconsumers are preserved.
 * * `gcp.endpoints.ConsumersIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the serviceconsumers are preserved.
 * 
 * &gt; **Note:** `gcp.endpoints.ConsumersIamPolicy` **cannot** be used in conjunction with `gcp.endpoints.ConsumersIamBinding` and `gcp.endpoints.ConsumersIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.endpoints.ConsumersIamBinding` resources **can be** used in conjunction with `gcp.endpoints.ConsumersIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* services/{{service_name}}/consumers/{{consumer_project}} * {{service_name}}/{{consumer_project}} * {{consumer_project}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Endpoints serviceconsumers IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:endpoints/consumersIamBinding:ConsumersIamBinding editor &#34;services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:endpoints/consumersIamBinding:ConsumersIamBinding editor &#34;services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:endpoints/consumersIamBinding:ConsumersIamBinding editor services/{{service_name}}/consumers/{{consumer_project}}
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:endpoints/consumersIamBinding:ConsumersIamBinding")
public class ConsumersIamBinding extends com.pulumi.resources.CustomResource {
    @Export(name="condition", type=ConsumersIamBindingCondition.class, parameters={})
    private Output</* @Nullable */ ConsumersIamBindingCondition> condition;

    public Output<Optional<ConsumersIamBindingCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    @Export(name="consumerProject", type=String.class, parameters={})
    private Output<String> consumerProject;

    public Output<String> consumerProject() {
        return this.consumerProject;
    }
    /**
     * (Computed) The etag of the IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="members", type=List.class, parameters={String.class})
    private Output<List<String>> members;

    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.endpoints.ConsumersIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.endpoints.ConsumersIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    @Export(name="serviceName", type=String.class, parameters={})
    private Output<String> serviceName;

    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ConsumersIamBinding(String name) {
        this(name, ConsumersIamBindingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConsumersIamBinding(String name, ConsumersIamBindingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConsumersIamBinding(String name, ConsumersIamBindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:endpoints/consumersIamBinding:ConsumersIamBinding", name, args == null ? ConsumersIamBindingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ConsumersIamBinding(String name, Output<String> id, @Nullable ConsumersIamBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:endpoints/consumersIamBinding:ConsumersIamBinding", name, state, makeResourceOptions(options, id));
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
    public static ConsumersIamBinding get(String name, Output<String> id, @Nullable ConsumersIamBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConsumersIamBinding(name, id, state, options);
    }
}
