// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.securitycenter;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.securitycenter.V2OrganizationMuteConfigArgs;
import com.pulumi.gcp.securitycenter.inputs.V2OrganizationMuteConfigState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Mute Findings is a volume management feature in Security Command Center
 * that lets you manually or programmatically hide irrelevant findings,
 * and create filters to automatically silence existing and future
 * findings based on criteria you specify.
 * 
 * To get more information about OrganizationMuteConfig, see:
 * 
 * * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v2/organizations.muteConfigs)
 * 
 * ## Example Usage
 * 
 * ### Scc V2 Organization Mute Config Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.securitycenter.V2OrganizationMuteConfig;
 * import com.pulumi.gcp.securitycenter.V2OrganizationMuteConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var default_ = new V2OrganizationMuteConfig("default", V2OrganizationMuteConfigArgs.builder()
 *             .muteConfigId("my-config")
 *             .organization("123456789")
 *             .location("global")
 *             .description("My custom Cloud Security Command Center Finding Organization mute Configuration")
 *             .filter("severity = \"HIGH\"")
 *             .type("STATIC")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * OrganizationMuteConfig can be imported using any of these accepted formats:
 * 
 * * `organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}`
 * 
 * * `{{organization}}/{{location}}/{{mute_config_id}}`
 * 
 * When using the `pulumi import` command, OrganizationMuteConfig can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:securitycenter/v2OrganizationMuteConfig:V2OrganizationMuteConfig default organizations/{{organization}}/locations/{{location}}/muteConfigs/{{mute_config_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:securitycenter/v2OrganizationMuteConfig:V2OrganizationMuteConfig default {{organization}}/{{location}}/{{mute_config_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:securitycenter/v2OrganizationMuteConfig:V2OrganizationMuteConfig")
public class V2OrganizationMuteConfig extends com.pulumi.resources.CustomResource {
    /**
     * The time at which the mute config was created. This field is set by
     * the server and will be ignored if provided on config creation.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The time at which the mute config was created. This field is set by
     * the server and will be ignored if provided on config creation.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * A description of the mute config.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the mute config.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * An expression that defines the filter to apply across create/update
     * events of findings. While creating a filter string, be mindful of
     * the scope in which the mute configuration is being created. E.g.,
     * If a filter contains project = X but is created under the
     * project = Y scope, it might not match any findings.
     * 
     */
    @Export(name="filter", refs={String.class}, tree="[0]")
    private Output<String> filter;

    /**
     * @return An expression that defines the filter to apply across create/update
     * events of findings. While creating a filter string, be mindful of
     * the scope in which the mute configuration is being created. E.g.,
     * If a filter contains project = X but is created under the
     * project = Y scope, it might not match any findings.
     * 
     */
    public Output<String> filter() {
        return this.filter;
    }
    /**
     * location Id is provided by organization. If not provided, Use global as default.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> location;

    /**
     * @return location Id is provided by organization. If not provided, Use global as default.
     * 
     */
    public Output<Optional<String>> location() {
        return Codegen.optional(this.location);
    }
    /**
     * Email address of the user who last edited the mute config. This
     * field is set by the server and will be ignored if provided on
     * config creation or update.
     * 
     */
    @Export(name="mostRecentEditor", refs={String.class}, tree="[0]")
    private Output<String> mostRecentEditor;

    /**
     * @return Email address of the user who last edited the mute config. This
     * field is set by the server and will be ignored if provided on
     * config creation or update.
     * 
     */
    public Output<String> mostRecentEditor() {
        return this.mostRecentEditor;
    }
    /**
     * Unique identifier provided by the client within the parent scope.
     * 
     * ***
     * 
     */
    @Export(name="muteConfigId", refs={String.class}, tree="[0]")
    private Output<String> muteConfigId;

    /**
     * @return Unique identifier provided by the client within the parent scope.
     * 
     * ***
     * 
     */
    public Output<String> muteConfigId() {
        return this.muteConfigId;
    }
    /**
     * Name of the mute config. Its format is
     * organizations/{organization}/locations/global/muteConfigs/{configId},
     * folders/{folder}/locations/global/muteConfigs/{configId},
     * or projects/{project}/locations/global/muteConfigs/{configId}
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the mute config. Its format is
     * organizations/{organization}/locations/global/muteConfigs/{configId},
     * folders/{folder}/locations/global/muteConfigs/{configId},
     * or projects/{project}/locations/global/muteConfigs/{configId}
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization whose Cloud Security Command Center the Mute
     * Config lives in.
     * 
     */
    @Export(name="organization", refs={String.class}, tree="[0]")
    private Output<String> organization;

    /**
     * @return The organization whose Cloud Security Command Center the Mute
     * Config lives in.
     * 
     */
    public Output<String> organization() {
        return this.organization;
    }
    /**
     * The type of the mute config.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the mute config.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Output only. The most recent time at which the mute config was
     * updated. This field is set by the server and will be ignored if
     * provided on config creation or update.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. The most recent time at which the mute config was
     * updated. This field is set by the server and will be ignored if
     * provided on config creation or update.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public V2OrganizationMuteConfig(java.lang.String name) {
        this(name, V2OrganizationMuteConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public V2OrganizationMuteConfig(java.lang.String name, V2OrganizationMuteConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public V2OrganizationMuteConfig(java.lang.String name, V2OrganizationMuteConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securitycenter/v2OrganizationMuteConfig:V2OrganizationMuteConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private V2OrganizationMuteConfig(java.lang.String name, Output<java.lang.String> id, @Nullable V2OrganizationMuteConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securitycenter/v2OrganizationMuteConfig:V2OrganizationMuteConfig", name, state, makeResourceOptions(options, id), false);
    }

    private static V2OrganizationMuteConfigArgs makeArgs(V2OrganizationMuteConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? V2OrganizationMuteConfigArgs.Empty : args;
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
    public static V2OrganizationMuteConfig get(java.lang.String name, Output<java.lang.String> id, @Nullable V2OrganizationMuteConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new V2OrganizationMuteConfig(name, id, state, options);
    }
}
