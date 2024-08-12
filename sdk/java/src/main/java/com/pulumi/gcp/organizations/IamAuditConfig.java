// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.organizations;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.organizations.IamAuditConfigArgs;
import com.pulumi.gcp.organizations.inputs.IamAuditConfigState;
import com.pulumi.gcp.organizations.outputs.IamAuditConfigAuditLogConfig;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Allows management of audit logging config for a given service for a Google Cloud Platform Organization.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.IamAuditConfig;
 * import com.pulumi.gcp.organizations.IamAuditConfigArgs;
 * import com.pulumi.gcp.organizations.inputs.IamAuditConfigAuditLogConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var config = new IamAuditConfig("config", IamAuditConfigArgs.builder()
 *             .orgId("your-organization-id")
 *             .service("allServices")
 *             .auditLogConfigs(IamAuditConfigAuditLogConfigArgs.builder()
 *                 .logType("DATA_READ")
 *                 .exemptedMembers("user:joebloggs}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * IAM audit config imports use the identifier of the resource in question and the service, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig config &#34;your-organization-id foo.googleapis.com&#34;
 * ```
 * 
 */
@ResourceType(type="gcp:organizations/iamAuditConfig:IamAuditConfig")
public class IamAuditConfig extends com.pulumi.resources.CustomResource {
    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     * 
     */
    @Export(name="auditLogConfigs", refs={List.class,IamAuditConfigAuditLogConfig.class}, tree="[0,1]")
    private Output<List<IamAuditConfigAuditLogConfig>> auditLogConfigs;

    /**
     * @return The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     * 
     */
    public Output<List<IamAuditConfigAuditLogConfig>> auditLogConfigs() {
        return this.auditLogConfigs;
    }
    /**
     * The etag of iam policy
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return The etag of iam policy
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The numeric ID of the organization in which you want to manage the audit logging config.
     * 
     */
    @Export(name="orgId", refs={String.class}, tree="[0]")
    private Output<String> orgId;

    /**
     * @return The numeric ID of the organization in which you want to manage the audit logging config.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
     * 
     */
    @Export(name="service", refs={String.class}, tree="[0]")
    private Output<String> service;

    /**
     * @return Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
     * 
     */
    public Output<String> service() {
        return this.service;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamAuditConfig(java.lang.String name) {
        this(name, IamAuditConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamAuditConfig(java.lang.String name, IamAuditConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamAuditConfig(java.lang.String name, IamAuditConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:organizations/iamAuditConfig:IamAuditConfig", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IamAuditConfig(java.lang.String name, Output<java.lang.String> id, @Nullable IamAuditConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:organizations/iamAuditConfig:IamAuditConfig", name, state, makeResourceOptions(options, id), false);
    }

    private static IamAuditConfigArgs makeArgs(IamAuditConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IamAuditConfigArgs.Empty : args;
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
    public static IamAuditConfig get(java.lang.String name, Output<java.lang.String> id, @Nullable IamAuditConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamAuditConfig(name, id, state, options);
    }
}
