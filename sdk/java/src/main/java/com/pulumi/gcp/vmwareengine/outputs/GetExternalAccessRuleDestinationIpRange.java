// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vmwareengine.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetExternalAccessRuleDestinationIpRange {
    private String externalAddress;
    private String ipAddressRange;

    private GetExternalAccessRuleDestinationIpRange() {}
    public String externalAddress() {
        return this.externalAddress;
    }
    public String ipAddressRange() {
        return this.ipAddressRange;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetExternalAccessRuleDestinationIpRange defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String externalAddress;
        private String ipAddressRange;
        public Builder() {}
        public Builder(GetExternalAccessRuleDestinationIpRange defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.externalAddress = defaults.externalAddress;
    	      this.ipAddressRange = defaults.ipAddressRange;
        }

        @CustomType.Setter
        public Builder externalAddress(String externalAddress) {
            if (externalAddress == null) {
              throw new MissingRequiredPropertyException("GetExternalAccessRuleDestinationIpRange", "externalAddress");
            }
            this.externalAddress = externalAddress;
            return this;
        }
        @CustomType.Setter
        public Builder ipAddressRange(String ipAddressRange) {
            if (ipAddressRange == null) {
              throw new MissingRequiredPropertyException("GetExternalAccessRuleDestinationIpRange", "ipAddressRange");
            }
            this.ipAddressRange = ipAddressRange;
            return this;
        }
        public GetExternalAccessRuleDestinationIpRange build() {
            final var _resultValue = new GetExternalAccessRuleDestinationIpRange();
            _resultValue.externalAddress = externalAddress;
            _resultValue.ipAddressRange = ipAddressRange;
            return _resultValue;
        }
    }
}
