// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vmwareengine.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExternalAccessRuleSourceIpRangeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ExternalAccessRuleSourceIpRangeArgs Empty = new ExternalAccessRuleSourceIpRangeArgs();

    /**
     * A single IP address.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return A single IP address.
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * An IP address range in the CIDR format.
     * 
     */
    @Import(name="ipAddressRange")
    private @Nullable Output<String> ipAddressRange;

    /**
     * @return An IP address range in the CIDR format.
     * 
     */
    public Optional<Output<String>> ipAddressRange() {
        return Optional.ofNullable(this.ipAddressRange);
    }

    private ExternalAccessRuleSourceIpRangeArgs() {}

    private ExternalAccessRuleSourceIpRangeArgs(ExternalAccessRuleSourceIpRangeArgs $) {
        this.ipAddress = $.ipAddress;
        this.ipAddressRange = $.ipAddressRange;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExternalAccessRuleSourceIpRangeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExternalAccessRuleSourceIpRangeArgs $;

        public Builder() {
            $ = new ExternalAccessRuleSourceIpRangeArgs();
        }

        public Builder(ExternalAccessRuleSourceIpRangeArgs defaults) {
            $ = new ExternalAccessRuleSourceIpRangeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipAddress A single IP address.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress A single IP address.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param ipAddressRange An IP address range in the CIDR format.
         * 
         * @return builder
         * 
         */
        public Builder ipAddressRange(@Nullable Output<String> ipAddressRange) {
            $.ipAddressRange = ipAddressRange;
            return this;
        }

        /**
         * @param ipAddressRange An IP address range in the CIDR format.
         * 
         * @return builder
         * 
         */
        public Builder ipAddressRange(String ipAddressRange) {
            return ipAddressRange(Output.of(ipAddressRange));
        }

        public ExternalAccessRuleSourceIpRangeArgs build() {
            return $;
        }
    }

}
