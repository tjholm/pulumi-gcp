// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.organizations.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIAMPolicyAuditConfigAuditLogConfig extends com.pulumi.resources.InvokeArgs {

    public static final GetIAMPolicyAuditConfigAuditLogConfig Empty = new GetIAMPolicyAuditConfigAuditLogConfig();

    /**
     * Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
     * 
     */
    @Import(name="exemptedMembers")
    private @Nullable List<String> exemptedMembers;

    /**
     * @return Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
     * 
     */
    public Optional<List<String>> exemptedMembers() {
        return Optional.ofNullable(this.exemptedMembers);
    }

    /**
     * Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
     * 
     */
    @Import(name="logType", required=true)
    private String logType;

    /**
     * @return Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
     * 
     */
    public String logType() {
        return this.logType;
    }

    private GetIAMPolicyAuditConfigAuditLogConfig() {}

    private GetIAMPolicyAuditConfigAuditLogConfig(GetIAMPolicyAuditConfigAuditLogConfig $) {
        this.exemptedMembers = $.exemptedMembers;
        this.logType = $.logType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIAMPolicyAuditConfigAuditLogConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIAMPolicyAuditConfigAuditLogConfig $;

        public Builder() {
            $ = new GetIAMPolicyAuditConfigAuditLogConfig();
        }

        public Builder(GetIAMPolicyAuditConfigAuditLogConfig defaults) {
            $ = new GetIAMPolicyAuditConfigAuditLogConfig(Objects.requireNonNull(defaults));
        }

        /**
         * @param exemptedMembers Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
         * 
         * @return builder
         * 
         */
        public Builder exemptedMembers(@Nullable List<String> exemptedMembers) {
            $.exemptedMembers = exemptedMembers;
            return this;
        }

        /**
         * @param exemptedMembers Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
         * 
         * @return builder
         * 
         */
        public Builder exemptedMembers(String... exemptedMembers) {
            return exemptedMembers(List.of(exemptedMembers));
        }

        /**
         * @param logType Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
         * 
         * @return builder
         * 
         */
        public Builder logType(String logType) {
            $.logType = logType;
            return this;
        }

        public GetIAMPolicyAuditConfigAuditLogConfig build() {
            $.logType = Objects.requireNonNull($.logType, "expected parameter 'logType' to be non-null");
            return $;
        }
    }

}
