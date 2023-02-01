// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.deploymentmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.deploymentmanager.inputs.DeploymentLabelArgs;
import com.pulumi.gcp.deploymentmanager.inputs.DeploymentTargetArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DeploymentState extends com.pulumi.resources.ResourceArgs {

    public static final DeploymentState Empty = new DeploymentState();

    /**
     * Set the policy to use for creating new resources. Only used on
     * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
     * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
     * the deployment will fail. Note that updating this field does not
     * actually affect the deployment, just how it is updated.
     * Default value is `CREATE_OR_ACQUIRE`.
     * Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
     * 
     */
    @Import(name="createPolicy")
    private @Nullable Output<String> createPolicy;

    /**
     * @return Set the policy to use for creating new resources. Only used on
     * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
     * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
     * the deployment will fail. Note that updating this field does not
     * actually affect the deployment, just how it is updated.
     * Default value is `CREATE_OR_ACQUIRE`.
     * Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
     * 
     */
    public Optional<Output<String>> createPolicy() {
        return Optional.ofNullable(this.createPolicy);
    }

    /**
     * Set the policy to use for deleting new resources on update/delete.
     * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
     * resource is deleted after removal from Deployment Manager. If
     * `ABANDON`, the resource is only removed from Deployment Manager
     * and is not actually deleted. Note that updating this field does not
     * actually change the deployment, just how it is updated.
     * Default value is `DELETE`.
     * Possible values are `ABANDON` and `DELETE`.
     * 
     */
    @Import(name="deletePolicy")
    private @Nullable Output<String> deletePolicy;

    /**
     * @return Set the policy to use for deleting new resources on update/delete.
     * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
     * resource is deleted after removal from Deployment Manager. If
     * `ABANDON`, the resource is only removed from Deployment Manager
     * and is not actually deleted. Note that updating this field does not
     * actually change the deployment, just how it is updated.
     * Default value is `DELETE`.
     * Possible values are `ABANDON` and `DELETE`.
     * 
     */
    public Optional<Output<String>> deletePolicy() {
        return Optional.ofNullable(this.deletePolicy);
    }

    /**
     * Unique identifier for deployment. Output only.
     * 
     */
    @Import(name="deploymentId")
    private @Nullable Output<String> deploymentId;

    /**
     * @return Unique identifier for deployment. Output only.
     * 
     */
    public Optional<Output<String>> deploymentId() {
        return Optional.ofNullable(this.deploymentId);
    }

    /**
     * Optional user-provided description of deployment.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Optional user-provided description of deployment.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Key-value pairs to apply to this labels.
     * Structure is documented below.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<DeploymentLabelArgs>> labels;

    /**
     * @return Key-value pairs to apply to this labels.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<DeploymentLabelArgs>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Output only. URL of the manifest representing the last manifest that
     * was successfully deployed.
     * 
     */
    @Import(name="manifest")
    private @Nullable Output<String> manifest;

    /**
     * @return Output only. URL of the manifest representing the last manifest that
     * was successfully deployed.
     * 
     */
    public Optional<Output<String>> manifest() {
        return Optional.ofNullable(this.manifest);
    }

    /**
     * Unique name for the deployment
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Unique name for the deployment
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * If set to true, a deployment is created with &#34;shell&#34; resources that are not actually instantiated. This allows you to
     * preview a deployment. It can be updated to false to actually deploy with real resources. ~&gt;**NOTE:** Deployment Manager
     * does not allow update of a deployment in preview (unless updating to preview=false). Thus, Terraform will force-recreate
     * deployments if either preview is updated to true or if other fields are updated while preview is true.
     * 
     */
    @Import(name="preview")
    private @Nullable Output<Boolean> preview;

    /**
     * @return If set to true, a deployment is created with &#34;shell&#34; resources that are not actually instantiated. This allows you to
     * preview a deployment. It can be updated to false to actually deploy with real resources. ~&gt;**NOTE:** Deployment Manager
     * does not allow update of a deployment in preview (unless updating to preview=false). Thus, Terraform will force-recreate
     * deployments if either preview is updated to true or if other fields are updated while preview is true.
     * 
     */
    public Optional<Output<Boolean>> preview() {
        return Optional.ofNullable(this.preview);
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
     * Output only. Server defined URL for the resource.
     * 
     */
    @Import(name="selfLink")
    private @Nullable Output<String> selfLink;

    /**
     * @return Output only. Server defined URL for the resource.
     * 
     */
    public Optional<Output<String>> selfLink() {
        return Optional.ofNullable(this.selfLink);
    }

    /**
     * Parameters that define your deployment, including the deployment
     * configuration and relevant templates.
     * Structure is documented below.
     * 
     */
    @Import(name="target")
    private @Nullable Output<DeploymentTargetArgs> target;

    /**
     * @return Parameters that define your deployment, including the deployment
     * configuration and relevant templates.
     * Structure is documented below.
     * 
     */
    public Optional<Output<DeploymentTargetArgs>> target() {
        return Optional.ofNullable(this.target);
    }

    private DeploymentState() {}

    private DeploymentState(DeploymentState $) {
        this.createPolicy = $.createPolicy;
        this.deletePolicy = $.deletePolicy;
        this.deploymentId = $.deploymentId;
        this.description = $.description;
        this.labels = $.labels;
        this.manifest = $.manifest;
        this.name = $.name;
        this.preview = $.preview;
        this.project = $.project;
        this.selfLink = $.selfLink;
        this.target = $.target;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DeploymentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DeploymentState $;

        public Builder() {
            $ = new DeploymentState();
        }

        public Builder(DeploymentState defaults) {
            $ = new DeploymentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createPolicy Set the policy to use for creating new resources. Only used on
         * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
         * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
         * the deployment will fail. Note that updating this field does not
         * actually affect the deployment, just how it is updated.
         * Default value is `CREATE_OR_ACQUIRE`.
         * Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
         * 
         * @return builder
         * 
         */
        public Builder createPolicy(@Nullable Output<String> createPolicy) {
            $.createPolicy = createPolicy;
            return this;
        }

        /**
         * @param createPolicy Set the policy to use for creating new resources. Only used on
         * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
         * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
         * the deployment will fail. Note that updating this field does not
         * actually affect the deployment, just how it is updated.
         * Default value is `CREATE_OR_ACQUIRE`.
         * Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
         * 
         * @return builder
         * 
         */
        public Builder createPolicy(String createPolicy) {
            return createPolicy(Output.of(createPolicy));
        }

        /**
         * @param deletePolicy Set the policy to use for deleting new resources on update/delete.
         * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
         * resource is deleted after removal from Deployment Manager. If
         * `ABANDON`, the resource is only removed from Deployment Manager
         * and is not actually deleted. Note that updating this field does not
         * actually change the deployment, just how it is updated.
         * Default value is `DELETE`.
         * Possible values are `ABANDON` and `DELETE`.
         * 
         * @return builder
         * 
         */
        public Builder deletePolicy(@Nullable Output<String> deletePolicy) {
            $.deletePolicy = deletePolicy;
            return this;
        }

        /**
         * @param deletePolicy Set the policy to use for deleting new resources on update/delete.
         * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
         * resource is deleted after removal from Deployment Manager. If
         * `ABANDON`, the resource is only removed from Deployment Manager
         * and is not actually deleted. Note that updating this field does not
         * actually change the deployment, just how it is updated.
         * Default value is `DELETE`.
         * Possible values are `ABANDON` and `DELETE`.
         * 
         * @return builder
         * 
         */
        public Builder deletePolicy(String deletePolicy) {
            return deletePolicy(Output.of(deletePolicy));
        }

        /**
         * @param deploymentId Unique identifier for deployment. Output only.
         * 
         * @return builder
         * 
         */
        public Builder deploymentId(@Nullable Output<String> deploymentId) {
            $.deploymentId = deploymentId;
            return this;
        }

        /**
         * @param deploymentId Unique identifier for deployment. Output only.
         * 
         * @return builder
         * 
         */
        public Builder deploymentId(String deploymentId) {
            return deploymentId(Output.of(deploymentId));
        }

        /**
         * @param description Optional user-provided description of deployment.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Optional user-provided description of deployment.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param labels Key-value pairs to apply to this labels.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<DeploymentLabelArgs>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Key-value pairs to apply to this labels.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder labels(List<DeploymentLabelArgs> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels Key-value pairs to apply to this labels.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder labels(DeploymentLabelArgs... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param manifest Output only. URL of the manifest representing the last manifest that
         * was successfully deployed.
         * 
         * @return builder
         * 
         */
        public Builder manifest(@Nullable Output<String> manifest) {
            $.manifest = manifest;
            return this;
        }

        /**
         * @param manifest Output only. URL of the manifest representing the last manifest that
         * was successfully deployed.
         * 
         * @return builder
         * 
         */
        public Builder manifest(String manifest) {
            return manifest(Output.of(manifest));
        }

        /**
         * @param name Unique name for the deployment
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Unique name for the deployment
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param preview If set to true, a deployment is created with &#34;shell&#34; resources that are not actually instantiated. This allows you to
         * preview a deployment. It can be updated to false to actually deploy with real resources. ~&gt;**NOTE:** Deployment Manager
         * does not allow update of a deployment in preview (unless updating to preview=false). Thus, Terraform will force-recreate
         * deployments if either preview is updated to true or if other fields are updated while preview is true.
         * 
         * @return builder
         * 
         */
        public Builder preview(@Nullable Output<Boolean> preview) {
            $.preview = preview;
            return this;
        }

        /**
         * @param preview If set to true, a deployment is created with &#34;shell&#34; resources that are not actually instantiated. This allows you to
         * preview a deployment. It can be updated to false to actually deploy with real resources. ~&gt;**NOTE:** Deployment Manager
         * does not allow update of a deployment in preview (unless updating to preview=false). Thus, Terraform will force-recreate
         * deployments if either preview is updated to true or if other fields are updated while preview is true.
         * 
         * @return builder
         * 
         */
        public Builder preview(Boolean preview) {
            return preview(Output.of(preview));
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
         * @param selfLink Output only. Server defined URL for the resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(@Nullable Output<String> selfLink) {
            $.selfLink = selfLink;
            return this;
        }

        /**
         * @param selfLink Output only. Server defined URL for the resource.
         * 
         * @return builder
         * 
         */
        public Builder selfLink(String selfLink) {
            return selfLink(Output.of(selfLink));
        }

        /**
         * @param target Parameters that define your deployment, including the deployment
         * configuration and relevant templates.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<DeploymentTargetArgs> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target Parameters that define your deployment, including the deployment
         * configuration and relevant templates.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder target(DeploymentTargetArgs target) {
            return target(Output.of(target));
        }

        public DeploymentState build() {
            return $;
        }
    }

}
