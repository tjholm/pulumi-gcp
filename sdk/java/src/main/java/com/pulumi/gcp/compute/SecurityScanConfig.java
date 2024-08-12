// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.SecurityScanConfigArgs;
import com.pulumi.gcp.compute.inputs.SecurityScanConfigState;
import com.pulumi.gcp.compute.outputs.SecurityScanConfigAuthentication;
import com.pulumi.gcp.compute.outputs.SecurityScanConfigSchedule;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A ScanConfig resource contains the configurations to launch a scan.
 * 
 * To get more information about ScanConfig, see:
 * 
 * * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
 * * How-to Guides
 *     * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)
 * 
 * ## Example Usage
 * 
 * ### Scan Config Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Address;
 * import com.pulumi.gcp.compute.AddressArgs;
 * import com.pulumi.gcp.compute.SecurityScanConfig;
 * import com.pulumi.gcp.compute.SecurityScanConfigArgs;
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
 *         var scannerStaticIp = new Address("scannerStaticIp", AddressArgs.builder()
 *             .name("scan-basic-static-ip")
 *             .build());
 * 
 *         var scan_config = new SecurityScanConfig("scan-config", SecurityScanConfigArgs.builder()
 *             .displayName("scan-config")
 *             .startingUrls(scannerStaticIp.address().applyValue(address -> String.format("http://%s", address)))
 *             .targetPlatforms("COMPUTE")
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
 * ScanConfig can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/scanConfigs/{{name}}`
 * 
 * * `{{project}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, ScanConfig can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default projects/{{project}}/scanConfigs/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/securityScanConfig:SecurityScanConfig")
public class SecurityScanConfig extends com.pulumi.resources.CustomResource {
    /**
     * The authentication configuration.
     * If specified, service will use the authentication configuration during scanning.
     * Structure is documented below.
     * 
     */
    @Export(name="authentication", refs={SecurityScanConfigAuthentication.class}, tree="[0]")
    private Output</* @Nullable */ SecurityScanConfigAuthentication> authentication;

    /**
     * @return The authentication configuration.
     * If specified, service will use the authentication configuration during scanning.
     * Structure is documented below.
     * 
     */
    public Output<Optional<SecurityScanConfigAuthentication>> authentication() {
        return Codegen.optional(this.authentication);
    }
    /**
     * The blacklist URL patterns as described in
     * https://cloud.google.com/security-scanner/docs/excluded-urls
     * 
     */
    @Export(name="blacklistPatterns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> blacklistPatterns;

    /**
     * @return The blacklist URL patterns as described in
     * https://cloud.google.com/security-scanner/docs/excluded-urls
     * 
     */
    public Output<Optional<List<String>>> blacklistPatterns() {
        return Codegen.optional(this.blacklistPatterns);
    }
    /**
     * The user provider display name of the ScanConfig.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The user provider display name of the ScanConfig.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Controls export of scan configurations and results to Cloud Security Command Center.
     * Default value is `ENABLED`.
     * Possible values are: `ENABLED`, `DISABLED`.
     * 
     */
    @Export(name="exportToSecurityCommandCenter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> exportToSecurityCommandCenter;

    /**
     * @return Controls export of scan configurations and results to Cloud Security Command Center.
     * Default value is `ENABLED`.
     * Possible values are: `ENABLED`, `DISABLED`.
     * 
     */
    public Output<Optional<String>> exportToSecurityCommandCenter() {
        return Codegen.optional(this.exportToSecurityCommandCenter);
    }
    /**
     * The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
     * Defaults to 15.
     * 
     */
    @Export(name="maxQps", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxQps;

    /**
     * @return The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
     * Defaults to 15.
     * 
     */
    public Output<Optional<Integer>> maxQps() {
        return Codegen.optional(this.maxQps);
    }
    /**
     * A server defined name for this index. Format:
     * `projects/{{project}}/scanConfigs/{{server_generated_id}}`
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A server defined name for this index. Format:
     * `projects/{{project}}/scanConfigs/{{server_generated_id}}`
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
     * The schedule of the ScanConfig
     * Structure is documented below.
     * 
     */
    @Export(name="schedule", refs={SecurityScanConfigSchedule.class}, tree="[0]")
    private Output</* @Nullable */ SecurityScanConfigSchedule> schedule;

    /**
     * @return The schedule of the ScanConfig
     * Structure is documented below.
     * 
     */
    public Output<Optional<SecurityScanConfigSchedule>> schedule() {
        return Codegen.optional(this.schedule);
    }
    /**
     * The starting URLs from which the scanner finds site pages.
     * 
     * ***
     * 
     */
    @Export(name="startingUrls", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> startingUrls;

    /**
     * @return The starting URLs from which the scanner finds site pages.
     * 
     * ***
     * 
     */
    public Output<List<String>> startingUrls() {
        return this.startingUrls;
    }
    /**
     * Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
     * Each value may be one of: `APP_ENGINE`, `COMPUTE`.
     * 
     */
    @Export(name="targetPlatforms", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> targetPlatforms;

    /**
     * @return Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
     * Each value may be one of: `APP_ENGINE`, `COMPUTE`.
     * 
     */
    public Output<Optional<List<String>>> targetPlatforms() {
        return Codegen.optional(this.targetPlatforms);
    }
    /**
     * Type of the user agents used for scanning
     * Default value is `CHROME_LINUX`.
     * Possible values are: `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, `SAFARI_IPHONE`.
     * 
     */
    @Export(name="userAgent", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userAgent;

    /**
     * @return Type of the user agents used for scanning
     * Default value is `CHROME_LINUX`.
     * Possible values are: `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, `SAFARI_IPHONE`.
     * 
     */
    public Output<Optional<String>> userAgent() {
        return Codegen.optional(this.userAgent);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecurityScanConfig(java.lang.String name) {
        this(name, SecurityScanConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecurityScanConfig(java.lang.String name, SecurityScanConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecurityScanConfig(java.lang.String name, SecurityScanConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/securityScanConfig:SecurityScanConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecurityScanConfig(java.lang.String name, Output<java.lang.String> id, @Nullable SecurityScanConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/securityScanConfig:SecurityScanConfig", name, state, makeResourceOptions(options, id), false);
    }

    private static SecurityScanConfigArgs makeArgs(SecurityScanConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecurityScanConfigArgs.Empty : args;
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
    public static SecurityScanConfig get(java.lang.String name, Output<java.lang.String> id, @Nullable SecurityScanConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecurityScanConfig(name, id, state, options);
    }
}
