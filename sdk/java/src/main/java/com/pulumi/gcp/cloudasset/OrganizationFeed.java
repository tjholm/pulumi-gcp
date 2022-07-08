// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudasset;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudasset.OrganizationFeedArgs;
import com.pulumi.gcp.cloudasset.inputs.OrganizationFeedState;
import com.pulumi.gcp.cloudasset.outputs.OrganizationFeedCondition;
import com.pulumi.gcp.cloudasset.outputs.OrganizationFeedFeedOutputConfig;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Describes a Cloud Asset Inventory feed used to to listen to asset updates.
 * 
 * To get more information about OrganizationFeed, see:
 * 
 * * [API documentation](https://cloud.google.com/asset-inventory/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/asset-inventory/docs)
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * OrganizationFeed can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:cloudasset/organizationFeed:OrganizationFeed default organizations/{{org_id}}/feeds/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:cloudasset/organizationFeed:OrganizationFeed default {{org_id}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:cloudasset/organizationFeed:OrganizationFeed")
public class OrganizationFeed extends com.pulumi.resources.CustomResource {
    /**
     * A list of the full names of the assets to receive updates. You must specify either or both of
     * assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
     * exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
     * See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
     * 
     */
    @Export(name="assetNames", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> assetNames;

    /**
     * @return A list of the full names of the assets to receive updates. You must specify either or both of
     * assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
     * exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
     * See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
     * 
     */
    public Output<Optional<List<String>>> assetNames() {
        return Codegen.optional(this.assetNames);
    }
    /**
     * A list of types of the assets to receive updates. You must specify either or both of assetNames
     * and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
     * the feed. For example: &#34;compute.googleapis.com/Disk&#34;
     * See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
     * supported asset types.
     * 
     */
    @Export(name="assetTypes", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> assetTypes;

    /**
     * @return A list of types of the assets to receive updates. You must specify either or both of assetNames
     * and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
     * the feed. For example: &#34;compute.googleapis.com/Disk&#34;
     * See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
     * supported asset types.
     * 
     */
    public Output<Optional<List<String>>> assetTypes() {
        return Codegen.optional(this.assetTypes);
    }
    /**
     * The project whose identity will be used when sending messages to the
     * destination pubsub topic. It also specifies the project for API
     * enablement check, quota, and billing.
     * 
     */
    @Export(name="billingProject", type=String.class, parameters={})
    private Output<String> billingProject;

    /**
     * @return The project whose identity will be used when sending messages to the
     * destination pubsub topic. It also specifies the project for API
     * enablement check, quota, and billing.
     * 
     */
    public Output<String> billingProject() {
        return this.billingProject;
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
    @Export(name="condition", type=OrganizationFeedCondition.class, parameters={})
    private Output</* @Nullable */ OrganizationFeedCondition> condition;

    /**
     * @return A condition which determines whether an asset update should be published. If specified, an asset
     * will be returned only when the expression evaluates to true. When set, expression field
     * must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
     * expression &#34;temporal_asset.deleted == true&#34; will only publish Asset deletions. Other fields of
     * condition are optional.
     * Structure is documented below.
     * 
     */
    public Output<Optional<OrganizationFeedCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * Asset content type. If not specified, no content but the asset name and type will be returned.
     * Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
     * 
     */
    @Export(name="contentType", type=String.class, parameters={})
    private Output</* @Nullable */ String> contentType;

    /**
     * @return Asset content type. If not specified, no content but the asset name and type will be returned.
     * Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
     * 
     */
    public Output<Optional<String>> contentType() {
        return Codegen.optional(this.contentType);
    }
    /**
     * This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
     * 
     */
    @Export(name="feedId", type=String.class, parameters={})
    private Output<String> feedId;

    /**
     * @return This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
     * 
     */
    public Output<String> feedId() {
        return this.feedId;
    }
    /**
     * Output configuration for asset feed destination.
     * Structure is documented below.
     * 
     */
    @Export(name="feedOutputConfig", type=OrganizationFeedFeedOutputConfig.class, parameters={})
    private Output<OrganizationFeedFeedOutputConfig> feedOutputConfig;

    /**
     * @return Output configuration for asset feed destination.
     * Structure is documented below.
     * 
     */
    public Output<OrganizationFeedFeedOutputConfig> feedOutputConfig() {
        return this.feedOutputConfig;
    }
    /**
     * The format will be organizations/{organization_number}/feeds/{client-assigned_feed_identifier}.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The format will be organizations/{organization_number}/feeds/{client-assigned_feed_identifier}.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization this feed should be created in.
     * 
     */
    @Export(name="orgId", type=String.class, parameters={})
    private Output<String> orgId;

    /**
     * @return The organization this feed should be created in.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OrganizationFeed(String name) {
        this(name, OrganizationFeedArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrganizationFeed(String name, OrganizationFeedArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrganizationFeed(String name, OrganizationFeedArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudasset/organizationFeed:OrganizationFeed", name, args == null ? OrganizationFeedArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OrganizationFeed(String name, Output<String> id, @Nullable OrganizationFeedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudasset/organizationFeed:OrganizationFeed", name, state, makeResourceOptions(options, id));
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
    public static OrganizationFeed get(String name, Output<String> id, @Nullable OrganizationFeedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrganizationFeed(name, id, state, options);
    }
}
