// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.projects;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.projects.inputs.IAMAuditConfigAuditLogConfigArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class IAMAuditConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final IAMAuditConfigArgs Empty = new IAMAuditConfigArgs();

    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     * 
     */
    @Import(name="auditLogConfigs", required=true)
    private Output<List<IAMAuditConfigAuditLogConfigArgs>> auditLogConfigs;

    /**
     * @return The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     * 
     */
    public Output<List<IAMAuditConfigAuditLogConfigArgs>> auditLogConfigs() {
        return this.auditLogConfigs;
    }

    /**
     * The project id of the target project. This is not
     * inferred from the provider.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The project id of the target project. This is not
     * inferred from the provider.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
     * 
     */
    @Import(name="service", required=true)
    private Output<String> service;

    /**
     * @return Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
     * 
     */
    public Output<String> service() {
        return this.service;
    }

    private IAMAuditConfigArgs() {}

    private IAMAuditConfigArgs(IAMAuditConfigArgs $) {
        this.auditLogConfigs = $.auditLogConfigs;
        this.project = $.project;
        this.service = $.service;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IAMAuditConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IAMAuditConfigArgs $;

        public Builder() {
            $ = new IAMAuditConfigArgs();
        }

        public Builder(IAMAuditConfigArgs defaults) {
            $ = new IAMAuditConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param auditLogConfigs The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder auditLogConfigs(Output<List<IAMAuditConfigAuditLogConfigArgs>> auditLogConfigs) {
            $.auditLogConfigs = auditLogConfigs;
            return this;
        }

        /**
         * @param auditLogConfigs The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder auditLogConfigs(List<IAMAuditConfigAuditLogConfigArgs> auditLogConfigs) {
            return auditLogConfigs(Output.of(auditLogConfigs));
        }

        /**
         * @param auditLogConfigs The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder auditLogConfigs(IAMAuditConfigAuditLogConfigArgs... auditLogConfigs) {
            return auditLogConfigs(List.of(auditLogConfigs));
        }

        /**
         * @param project The project id of the target project. This is not
         * inferred from the provider.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project id of the target project. This is not
         * inferred from the provider.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param service Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
         * 
         * @return builder
         * 
         */
        public Builder service(Output<String> service) {
            $.service = service;
            return this;
        }

        /**
         * @param service Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
         * 
         * @return builder
         * 
         */
        public Builder service(String service) {
            return service(Output.of(service));
        }

        public IAMAuditConfigArgs build() {
            $.auditLogConfigs = Objects.requireNonNull($.auditLogConfigs, "expected parameter 'auditLogConfigs' to be non-null");
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            $.service = Objects.requireNonNull($.service, "expected parameter 'service' to be non-null");
            return $;
        }
    }

}
