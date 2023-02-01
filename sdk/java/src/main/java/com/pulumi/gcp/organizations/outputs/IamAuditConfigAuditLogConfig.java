// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.organizations.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class IamAuditConfigAuditLogConfig {
    /**
     * @return Identities that do not cause logging for this type of permission.  The format is the same as that for `members`.
     * 
     */
    private @Nullable List<String> exemptedMembers;
    /**
     * @return Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
     * 
     */
    private String logType;

    private IamAuditConfigAuditLogConfig() {}
    /**
     * @return Identities that do not cause logging for this type of permission.  The format is the same as that for `members`.
     * 
     */
    public List<String> exemptedMembers() {
        return this.exemptedMembers == null ? List.of() : this.exemptedMembers;
    }
    /**
     * @return Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
     * 
     */
    public String logType() {
        return this.logType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IamAuditConfigAuditLogConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> exemptedMembers;
        private String logType;
        public Builder() {}
        public Builder(IamAuditConfigAuditLogConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.exemptedMembers = defaults.exemptedMembers;
    	      this.logType = defaults.logType;
        }

        @CustomType.Setter
        public Builder exemptedMembers(@Nullable List<String> exemptedMembers) {
            this.exemptedMembers = exemptedMembers;
            return this;
        }
        public Builder exemptedMembers(String... exemptedMembers) {
            return exemptedMembers(List.of(exemptedMembers));
        }
        @CustomType.Setter
        public Builder logType(String logType) {
            this.logType = Objects.requireNonNull(logType);
            return this;
        }
        public IamAuditConfigAuditLogConfig build() {
            final var o = new IamAuditConfigAuditLogConfig();
            o.exemptedMembers = exemptedMembers;
            o.logType = logType;
            return o;
        }
    }
}
