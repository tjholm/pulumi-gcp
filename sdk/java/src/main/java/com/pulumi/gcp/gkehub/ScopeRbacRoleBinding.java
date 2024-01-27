// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.gkehub.ScopeRbacRoleBindingArgs;
import com.pulumi.gcp.gkehub.outputs.ScopeRbacRoleBindingRole;
import com.pulumi.gcp.gkehub.outputs.ScopeRbacRoleBindingState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * RBACRoleBinding represents a rbacrolebinding across the Fleet.
 * 
 * To get more information about ScopeRBACRoleBinding, see:
 * 
 * * [API documentation](https://cloud.google.com/anthos/fleet-management/docs/reference/rest/v1/projects.locations.scopes.rbacrolebindings)
 * * How-to Guides
 *     * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * ScopeRBACRoleBinding can be imported using any of these accepted formats* `projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}` * `{{project}}/{{scope_id}}/{{scope_rbac_role_binding_id}}` * `{{scope_id}}/{{scope_rbac_role_binding_id}}` When using the `pulumi import` command, ScopeRBACRoleBinding can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:gkehub/scopeRbacRoleBinding:ScopeRbacRoleBinding default projects/{{project}}/locations/global/scopes/{{scope_id}}/rbacrolebindings/{{scope_rbac_role_binding_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:gkehub/scopeRbacRoleBinding:ScopeRbacRoleBinding default {{project}}/{{scope_id}}/{{scope_rbac_role_binding_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:gkehub/scopeRbacRoleBinding:ScopeRbacRoleBinding default {{scope_id}}/{{scope_rbac_role_binding_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:gkehub/scopeRbacRoleBinding:ScopeRbacRoleBinding")
public class ScopeRbacRoleBinding extends com.pulumi.resources.CustomResource {
    /**
     * Time the RBAC Role Binding was created in UTC.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Time the RBAC Role Binding was created in UTC.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Time the RBAC Role Binding was deleted in UTC.
     * 
     */
    @Export(name="deleteTime", refs={String.class}, tree="[0]")
    private Output<String> deleteTime;

    /**
     * @return Time the RBAC Role Binding was deleted in UTC.
     * 
     */
    public Output<String> deleteTime() {
        return this.deleteTime;
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
     * Principal that is be authorized in the cluster (at least of one the oneof
     * is required). Updating one will unset the other automatically.
     * group is the group, as seen by the kubernetes cluster.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> group;

    /**
     * @return Principal that is be authorized in the cluster (at least of one the oneof
     * is required). Updating one will unset the other automatically.
     * group is the group, as seen by the kubernetes cluster.
     * 
     */
    public Output<Optional<String>> group() {
        return Codegen.optional(this.group);
    }
    /**
     * Labels for this ScopeRBACRoleBinding.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels for this ScopeRBACRoleBinding.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The resource name for the RBAC Role Binding
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name for the RBAC Role Binding
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
     * Role to bind to the principal.
     * Structure is documented below.
     * 
     */
    @Export(name="role", refs={ScopeRbacRoleBindingRole.class}, tree="[0]")
    private Output<ScopeRbacRoleBindingRole> role;

    /**
     * @return Role to bind to the principal.
     * Structure is documented below.
     * 
     */
    public Output<ScopeRbacRoleBindingRole> role() {
        return this.role;
    }
    /**
     * Id of the scope
     * 
     */
    @Export(name="scopeId", refs={String.class}, tree="[0]")
    private Output<String> scopeId;

    /**
     * @return Id of the scope
     * 
     */
    public Output<String> scopeId() {
        return this.scopeId;
    }
    /**
     * The client-provided identifier of the RBAC Role Binding.
     * 
     */
    @Export(name="scopeRbacRoleBindingId", refs={String.class}, tree="[0]")
    private Output<String> scopeRbacRoleBindingId;

    /**
     * @return The client-provided identifier of the RBAC Role Binding.
     * 
     */
    public Output<String> scopeRbacRoleBindingId() {
        return this.scopeRbacRoleBindingId;
    }
    /**
     * State of the RBAC Role Binding resource.
     * Structure is documented below.
     * 
     */
    @Export(name="states", refs={List.class,ScopeRbacRoleBindingState.class}, tree="[0,1]")
    private Output<List<ScopeRbacRoleBindingState>> states;

    /**
     * @return State of the RBAC Role Binding resource.
     * Structure is documented below.
     * 
     */
    public Output<List<ScopeRbacRoleBindingState>> states() {
        return this.states;
    }
    /**
     * Google-generated UUID for this resource.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Google-generated UUID for this resource.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Time the RBAC Role Binding was updated in UTC.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Time the RBAC Role Binding was updated in UTC.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Principal that is be authorized in the cluster (at least of one the oneof
     * is required). Updating one will unset the other automatically.
     * user is the name of the user as seen by the kubernetes cluster, example
     * &#34;alice&#34; or &#34;alice@domain.tld&#34;
     * 
     */
    @Export(name="user", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> user;

    /**
     * @return Principal that is be authorized in the cluster (at least of one the oneof
     * is required). Updating one will unset the other automatically.
     * user is the name of the user as seen by the kubernetes cluster, example
     * &#34;alice&#34; or &#34;alice@domain.tld&#34;
     * 
     */
    public Output<Optional<String>> user() {
        return Codegen.optional(this.user);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ScopeRbacRoleBinding(String name) {
        this(name, ScopeRbacRoleBindingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ScopeRbacRoleBinding(String name, ScopeRbacRoleBindingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ScopeRbacRoleBinding(String name, ScopeRbacRoleBindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkehub/scopeRbacRoleBinding:ScopeRbacRoleBinding", name, args == null ? ScopeRbacRoleBindingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ScopeRbacRoleBinding(String name, Output<String> id, @Nullable com.pulumi.gcp.gkehub.inputs.ScopeRbacRoleBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkehub/scopeRbacRoleBinding:ScopeRbacRoleBinding", name, state, makeResourceOptions(options, id));
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
    public static ScopeRbacRoleBinding get(String name, Output<String> id, @Nullable com.pulumi.gcp.gkehub.inputs.ScopeRbacRoleBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ScopeRbacRoleBinding(name, id, state, options);
    }
}
