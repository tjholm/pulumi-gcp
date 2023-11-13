// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.gkehub.inputs.ScopeRbacRoleBindingRoleArgs;
import com.pulumi.gcp.gkehub.inputs.ScopeRbacRoleBindingStateArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ScopeRbacRoleBindingState extends com.pulumi.resources.ResourceArgs {

    public static final ScopeRbacRoleBindingState Empty = new ScopeRbacRoleBindingState();

    /**
     * Time the RBAC Role Binding was created in UTC.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Time the RBAC Role Binding was created in UTC.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Time the RBAC Role Binding was deleted in UTC.
     * 
     */
    @Import(name="deleteTime")
    private @Nullable Output<String> deleteTime;

    /**
     * @return Time the RBAC Role Binding was deleted in UTC.
     * 
     */
    public Optional<Output<String>> deleteTime() {
        return Optional.ofNullable(this.deleteTime);
    }

    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Import(name="effectiveLabels")
    private @Nullable Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Optional<Output<Map<String,String>>> effectiveLabels() {
        return Optional.ofNullable(this.effectiveLabels);
    }

    /**
     * Principal that is be authorized in the cluster (at least of one the oneof
     * is required). Updating one will unset the other automatically.
     * group is the group, as seen by the kubernetes cluster.
     * 
     */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return Principal that is be authorized in the cluster (at least of one the oneof
     * is required). Updating one will unset the other automatically.
     * group is the group, as seen by the kubernetes cluster.
     * 
     */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
    }

    /**
     * Labels for this ScopeRBACRoleBinding.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Labels for this ScopeRBACRoleBinding.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The resource name for the RBAC Role Binding
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The resource name for the RBAC Role Binding
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Import(name="pulumiLabels")
    private @Nullable Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Optional<Output<Map<String,String>>> pulumiLabels() {
        return Optional.ofNullable(this.pulumiLabels);
    }

    /**
     * Role to bind to the principal.
     * Structure is documented below.
     * 
     */
    @Import(name="role")
    private @Nullable Output<ScopeRbacRoleBindingRoleArgs> role;

    /**
     * @return Role to bind to the principal.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ScopeRbacRoleBindingRoleArgs>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * Id of the scope
     * 
     */
    @Import(name="scopeId")
    private @Nullable Output<String> scopeId;

    /**
     * @return Id of the scope
     * 
     */
    public Optional<Output<String>> scopeId() {
        return Optional.ofNullable(this.scopeId);
    }

    /**
     * The client-provided identifier of the RBAC Role Binding.
     * 
     */
    @Import(name="scopeRbacRoleBindingId")
    private @Nullable Output<String> scopeRbacRoleBindingId;

    /**
     * @return The client-provided identifier of the RBAC Role Binding.
     * 
     */
    public Optional<Output<String>> scopeRbacRoleBindingId() {
        return Optional.ofNullable(this.scopeRbacRoleBindingId);
    }

    /**
     * State of the RBAC Role Binding resource.
     * Structure is documented below.
     * 
     */
    @Import(name="states")
    private @Nullable Output<List<ScopeRbacRoleBindingStateArgs>> states;

    /**
     * @return State of the RBAC Role Binding resource.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<ScopeRbacRoleBindingStateArgs>>> states() {
        return Optional.ofNullable(this.states);
    }

    /**
     * Google-generated UUID for this resource.
     * 
     */
    @Import(name="uid")
    private @Nullable Output<String> uid;

    /**
     * @return Google-generated UUID for this resource.
     * 
     */
    public Optional<Output<String>> uid() {
        return Optional.ofNullable(this.uid);
    }

    /**
     * Time the RBAC Role Binding was updated in UTC.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return Time the RBAC Role Binding was updated in UTC.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    /**
     * Principal that is be authorized in the cluster (at least of one the oneof
     * is required). Updating one will unset the other automatically.
     * user is the name of the user as seen by the kubernetes cluster, example
     * &#34;alice&#34; or &#34;alice@domain.tld&#34;
     * 
     */
    @Import(name="user")
    private @Nullable Output<String> user;

    /**
     * @return Principal that is be authorized in the cluster (at least of one the oneof
     * is required). Updating one will unset the other automatically.
     * user is the name of the user as seen by the kubernetes cluster, example
     * &#34;alice&#34; or &#34;alice@domain.tld&#34;
     * 
     */
    public Optional<Output<String>> user() {
        return Optional.ofNullable(this.user);
    }

    private ScopeRbacRoleBindingState() {}

    private ScopeRbacRoleBindingState(ScopeRbacRoleBindingState $) {
        this.createTime = $.createTime;
        this.deleteTime = $.deleteTime;
        this.effectiveLabels = $.effectiveLabels;
        this.group = $.group;
        this.labels = $.labels;
        this.name = $.name;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.role = $.role;
        this.scopeId = $.scopeId;
        this.scopeRbacRoleBindingId = $.scopeRbacRoleBindingId;
        this.states = $.states;
        this.uid = $.uid;
        this.updateTime = $.updateTime;
        this.user = $.user;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ScopeRbacRoleBindingState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ScopeRbacRoleBindingState $;

        public Builder() {
            $ = new ScopeRbacRoleBindingState();
        }

        public Builder(ScopeRbacRoleBindingState defaults) {
            $ = new ScopeRbacRoleBindingState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime Time the RBAC Role Binding was created in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Time the RBAC Role Binding was created in UTC.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param deleteTime Time the RBAC Role Binding was deleted in UTC.
         * 
         * @return builder
         * 
         */
        public Builder deleteTime(@Nullable Output<String> deleteTime) {
            $.deleteTime = deleteTime;
            return this;
        }

        /**
         * @param deleteTime Time the RBAC Role Binding was deleted in UTC.
         * 
         * @return builder
         * 
         */
        public Builder deleteTime(String deleteTime) {
            return deleteTime(Output.of(deleteTime));
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(@Nullable Output<Map<String,String>> effectiveLabels) {
            $.effectiveLabels = effectiveLabels;
            return this;
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(Map<String,String> effectiveLabels) {
            return effectiveLabels(Output.of(effectiveLabels));
        }

        /**
         * @param group Principal that is be authorized in the cluster (at least of one the oneof
         * is required). Updating one will unset the other automatically.
         * group is the group, as seen by the kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group Principal that is be authorized in the cluster (at least of one the oneof
         * is required). Updating one will unset the other automatically.
         * group is the group, as seen by the kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param labels Labels for this ScopeRBACRoleBinding.
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Labels for this ScopeRBACRoleBinding.
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param name The resource name for the RBAC Role Binding
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The resource name for the RBAC Role Binding
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(@Nullable Output<Map<String,String>> pulumiLabels) {
            $.pulumiLabels = pulumiLabels;
            return this;
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(Map<String,String> pulumiLabels) {
            return pulumiLabels(Output.of(pulumiLabels));
        }

        /**
         * @param role Role to bind to the principal.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<ScopeRbacRoleBindingRoleArgs> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Role to bind to the principal.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder role(ScopeRbacRoleBindingRoleArgs role) {
            return role(Output.of(role));
        }

        /**
         * @param scopeId Id of the scope
         * 
         * @return builder
         * 
         */
        public Builder scopeId(@Nullable Output<String> scopeId) {
            $.scopeId = scopeId;
            return this;
        }

        /**
         * @param scopeId Id of the scope
         * 
         * @return builder
         * 
         */
        public Builder scopeId(String scopeId) {
            return scopeId(Output.of(scopeId));
        }

        /**
         * @param scopeRbacRoleBindingId The client-provided identifier of the RBAC Role Binding.
         * 
         * @return builder
         * 
         */
        public Builder scopeRbacRoleBindingId(@Nullable Output<String> scopeRbacRoleBindingId) {
            $.scopeRbacRoleBindingId = scopeRbacRoleBindingId;
            return this;
        }

        /**
         * @param scopeRbacRoleBindingId The client-provided identifier of the RBAC Role Binding.
         * 
         * @return builder
         * 
         */
        public Builder scopeRbacRoleBindingId(String scopeRbacRoleBindingId) {
            return scopeRbacRoleBindingId(Output.of(scopeRbacRoleBindingId));
        }

        /**
         * @param states State of the RBAC Role Binding resource.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder states(@Nullable Output<List<ScopeRbacRoleBindingStateArgs>> states) {
            $.states = states;
            return this;
        }

        /**
         * @param states State of the RBAC Role Binding resource.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder states(List<ScopeRbacRoleBindingStateArgs> states) {
            return states(Output.of(states));
        }

        /**
         * @param states State of the RBAC Role Binding resource.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder states(ScopeRbacRoleBindingStateArgs... states) {
            return states(List.of(states));
        }

        /**
         * @param uid Google-generated UUID for this resource.
         * 
         * @return builder
         * 
         */
        public Builder uid(@Nullable Output<String> uid) {
            $.uid = uid;
            return this;
        }

        /**
         * @param uid Google-generated UUID for this resource.
         * 
         * @return builder
         * 
         */
        public Builder uid(String uid) {
            return uid(Output.of(uid));
        }

        /**
         * @param updateTime Time the RBAC Role Binding was updated in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime Time the RBAC Role Binding was updated in UTC.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        /**
         * @param user Principal that is be authorized in the cluster (at least of one the oneof
         * is required). Updating one will unset the other automatically.
         * user is the name of the user as seen by the kubernetes cluster, example
         * &#34;alice&#34; or &#34;alice@domain.tld&#34;
         * 
         * @return builder
         * 
         */
        public Builder user(@Nullable Output<String> user) {
            $.user = user;
            return this;
        }

        /**
         * @param user Principal that is be authorized in the cluster (at least of one the oneof
         * is required). Updating one will unset the other automatically.
         * user is the name of the user as seen by the kubernetes cluster, example
         * &#34;alice&#34; or &#34;alice@domain.tld&#34;
         * 
         * @return builder
         * 
         */
        public Builder user(String user) {
            return user(Output.of(user));
        }

        public ScopeRbacRoleBindingState build() {
            return $;
        }
    }

}
