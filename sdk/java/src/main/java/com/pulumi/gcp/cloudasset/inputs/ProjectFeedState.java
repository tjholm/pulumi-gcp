// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudasset.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.cloudasset.inputs.ProjectFeedConditionArgs;
import com.pulumi.gcp.cloudasset.inputs.ProjectFeedFeedOutputConfigArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectFeedState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectFeedState Empty = new ProjectFeedState();

    /**
     * A list of the full names of the assets to receive updates. You must specify either or both of
     * assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
     * exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
     * See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
     * 
     */
    @Import(name="assetNames")
    private @Nullable Output<List<String>> assetNames;

    /**
     * @return A list of the full names of the assets to receive updates. You must specify either or both of
     * assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
     * exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
     * See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
     * 
     */
    public Optional<Output<List<String>>> assetNames() {
        return Optional.ofNullable(this.assetNames);
    }

    /**
     * A list of types of the assets to receive updates. You must specify either or both of assetNames
     * and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
     * the feed. For example: &#34;compute.googleapis.com/Disk&#34;
     * See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
     * supported asset types.
     * 
     */
    @Import(name="assetTypes")
    private @Nullable Output<List<String>> assetTypes;

    /**
     * @return A list of types of the assets to receive updates. You must specify either or both of assetNames
     * and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
     * the feed. For example: &#34;compute.googleapis.com/Disk&#34;
     * See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
     * supported asset types.
     * 
     */
    public Optional<Output<List<String>>> assetTypes() {
        return Optional.ofNullable(this.assetTypes);
    }

    /**
     * The project whose identity will be used when sending messages to the
     * destination pubsub topic. It also specifies the project for API
     * enablement check, quota, and billing. If not specified, the resource&#39;s
     * project will be used.
     * 
     */
    @Import(name="billingProject")
    private @Nullable Output<String> billingProject;

    /**
     * @return The project whose identity will be used when sending messages to the
     * destination pubsub topic. It also specifies the project for API
     * enablement check, quota, and billing. If not specified, the resource&#39;s
     * project will be used.
     * 
     */
    public Optional<Output<String>> billingProject() {
        return Optional.ofNullable(this.billingProject);
    }

    /**
     * A condition which determines whether an asset update should be published. If specified, an asset
     * will be returned only when the expression evaluates to true. When set, expression field
     * must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
     * expression &#34;temporal_asset.deleted == true&#34; will only publish Asset deletions. Other fields of
     * condition are optional.
     * Structure is documented below.
     * 
     */
    @Import(name="condition")
    private @Nullable Output<ProjectFeedConditionArgs> condition;

    /**
     * @return A condition which determines whether an asset update should be published. If specified, an asset
     * will be returned only when the expression evaluates to true. When set, expression field
     * must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
     * expression &#34;temporal_asset.deleted == true&#34; will only publish Asset deletions. Other fields of
     * condition are optional.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ProjectFeedConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * Asset content type. If not specified, no content but the asset name and type will be returned.
     * Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
     * 
     */
    @Import(name="contentType")
    private @Nullable Output<String> contentType;

    /**
     * @return Asset content type. If not specified, no content but the asset name and type will be returned.
     * Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
     * 
     */
    public Optional<Output<String>> contentType() {
        return Optional.ofNullable(this.contentType);
    }

    /**
     * This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
     * 
     */
    @Import(name="feedId")
    private @Nullable Output<String> feedId;

    /**
     * @return This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
     * 
     */
    public Optional<Output<String>> feedId() {
        return Optional.ofNullable(this.feedId);
    }

    /**
     * Output configuration for asset feed destination.
     * Structure is documented below.
     * 
     */
    @Import(name="feedOutputConfig")
    private @Nullable Output<ProjectFeedFeedOutputConfigArgs> feedOutputConfig;

    /**
     * @return Output configuration for asset feed destination.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ProjectFeedFeedOutputConfigArgs>> feedOutputConfig() {
        return Optional.ofNullable(this.feedOutputConfig);
    }

    /**
     * The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
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

    private ProjectFeedState() {}

    private ProjectFeedState(ProjectFeedState $) {
        this.assetNames = $.assetNames;
        this.assetTypes = $.assetTypes;
        this.billingProject = $.billingProject;
        this.condition = $.condition;
        this.contentType = $.contentType;
        this.feedId = $.feedId;
        this.feedOutputConfig = $.feedOutputConfig;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectFeedState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectFeedState $;

        public Builder() {
            $ = new ProjectFeedState();
        }

        public Builder(ProjectFeedState defaults) {
            $ = new ProjectFeedState(Objects.requireNonNull(defaults));
        }

        /**
         * @param assetNames A list of the full names of the assets to receive updates. You must specify either or both of
         * assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
         * exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
         * See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
         * 
         * @return builder
         * 
         */
        public Builder assetNames(@Nullable Output<List<String>> assetNames) {
            $.assetNames = assetNames;
            return this;
        }

        /**
         * @param assetNames A list of the full names of the assets to receive updates. You must specify either or both of
         * assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
         * exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
         * See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
         * 
         * @return builder
         * 
         */
        public Builder assetNames(List<String> assetNames) {
            return assetNames(Output.of(assetNames));
        }

        /**
         * @param assetNames A list of the full names of the assets to receive updates. You must specify either or both of
         * assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
         * exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
         * See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
         * 
         * @return builder
         * 
         */
        public Builder assetNames(String... assetNames) {
            return assetNames(List.of(assetNames));
        }

        /**
         * @param assetTypes A list of types of the assets to receive updates. You must specify either or both of assetNames
         * and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
         * the feed. For example: &#34;compute.googleapis.com/Disk&#34;
         * See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
         * supported asset types.
         * 
         * @return builder
         * 
         */
        public Builder assetTypes(@Nullable Output<List<String>> assetTypes) {
            $.assetTypes = assetTypes;
            return this;
        }

        /**
         * @param assetTypes A list of types of the assets to receive updates. You must specify either or both of assetNames
         * and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
         * the feed. For example: &#34;compute.googleapis.com/Disk&#34;
         * See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
         * supported asset types.
         * 
         * @return builder
         * 
         */
        public Builder assetTypes(List<String> assetTypes) {
            return assetTypes(Output.of(assetTypes));
        }

        /**
         * @param assetTypes A list of types of the assets to receive updates. You must specify either or both of assetNames
         * and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
         * the feed. For example: &#34;compute.googleapis.com/Disk&#34;
         * See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
         * supported asset types.
         * 
         * @return builder
         * 
         */
        public Builder assetTypes(String... assetTypes) {
            return assetTypes(List.of(assetTypes));
        }

        /**
         * @param billingProject The project whose identity will be used when sending messages to the
         * destination pubsub topic. It also specifies the project for API
         * enablement check, quota, and billing. If not specified, the resource&#39;s
         * project will be used.
         * 
         * @return builder
         * 
         */
        public Builder billingProject(@Nullable Output<String> billingProject) {
            $.billingProject = billingProject;
            return this;
        }

        /**
         * @param billingProject The project whose identity will be used when sending messages to the
         * destination pubsub topic. It also specifies the project for API
         * enablement check, quota, and billing. If not specified, the resource&#39;s
         * project will be used.
         * 
         * @return builder
         * 
         */
        public Builder billingProject(String billingProject) {
            return billingProject(Output.of(billingProject));
        }

        /**
         * @param condition A condition which determines whether an asset update should be published. If specified, an asset
         * will be returned only when the expression evaluates to true. When set, expression field
         * must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
         * expression &#34;temporal_asset.deleted == true&#34; will only publish Asset deletions. Other fields of
         * condition are optional.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder condition(@Nullable Output<ProjectFeedConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        /**
         * @param condition A condition which determines whether an asset update should be published. If specified, an asset
         * will be returned only when the expression evaluates to true. When set, expression field
         * must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
         * expression &#34;temporal_asset.deleted == true&#34; will only publish Asset deletions. Other fields of
         * condition are optional.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder condition(ProjectFeedConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param contentType Asset content type. If not specified, no content but the asset name and type will be returned.
         * Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
         * 
         * @return builder
         * 
         */
        public Builder contentType(@Nullable Output<String> contentType) {
            $.contentType = contentType;
            return this;
        }

        /**
         * @param contentType Asset content type. If not specified, no content but the asset name and type will be returned.
         * Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
         * 
         * @return builder
         * 
         */
        public Builder contentType(String contentType) {
            return contentType(Output.of(contentType));
        }

        /**
         * @param feedId This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
         * 
         * @return builder
         * 
         */
        public Builder feedId(@Nullable Output<String> feedId) {
            $.feedId = feedId;
            return this;
        }

        /**
         * @param feedId This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
         * 
         * @return builder
         * 
         */
        public Builder feedId(String feedId) {
            return feedId(Output.of(feedId));
        }

        /**
         * @param feedOutputConfig Output configuration for asset feed destination.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder feedOutputConfig(@Nullable Output<ProjectFeedFeedOutputConfigArgs> feedOutputConfig) {
            $.feedOutputConfig = feedOutputConfig;
            return this;
        }

        /**
         * @param feedOutputConfig Output configuration for asset feed destination.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder feedOutputConfig(ProjectFeedFeedOutputConfigArgs feedOutputConfig) {
            return feedOutputConfig(Output.of(feedOutputConfig));
        }

        /**
         * @param name The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The format will be projects/{projectNumber}/feeds/{client-assigned_feed_identifier}.
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

        public ProjectFeedState build() {
            return $;
        }
    }

}
