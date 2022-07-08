// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.accesscontextmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.accesscontextmanager.inputs.AccessLevelConditionDevicePolicyArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessLevelConditionState extends com.pulumi.resources.ResourceArgs {

    public static final AccessLevelConditionState Empty = new AccessLevelConditionState();

    /**
     * The name of the Access Level to add this condition to.
     * 
     */
    @Import(name="accessLevel")
    private @Nullable Output<String> accessLevel;

    /**
     * @return The name of the Access Level to add this condition to.
     * 
     */
    public Optional<Output<String>> accessLevel() {
        return Optional.ofNullable(this.accessLevel);
    }

    /**
     * Device specific restrictions, all restrictions must hold for
     * the Condition to be true. If not specified, all devices are
     * allowed.
     * Structure is documented below.
     * 
     */
    @Import(name="devicePolicy")
    private @Nullable Output<AccessLevelConditionDevicePolicyArgs> devicePolicy;

    /**
     * @return Device specific restrictions, all restrictions must hold for
     * the Condition to be true. If not specified, all devices are
     * allowed.
     * Structure is documented below.
     * 
     */
    public Optional<Output<AccessLevelConditionDevicePolicyArgs>> devicePolicy() {
        return Optional.ofNullable(this.devicePolicy);
    }

    /**
     * A list of CIDR block IP subnetwork specification. May be IPv4
     * or IPv6.
     * Note that for a CIDR IP address block, the specified IP address
     * portion must be properly truncated (i.e. all the host bits must
     * be zero) or the input is considered malformed. For example,
     * &#34;192.0.2.0/24&#34; is accepted but &#34;192.0.2.1/24&#34; is not. Similarly,
     * for IPv6, &#34;2001:db8::/32&#34; is accepted whereas &#34;2001:db8::1/32&#34;
     * is not. The originating IP of a request must be in one of the
     * listed subnets in order for this Condition to be true.
     * If empty, all IP addresses are allowed.
     * 
     */
    @Import(name="ipSubnetworks")
    private @Nullable Output<List<String>> ipSubnetworks;

    /**
     * @return A list of CIDR block IP subnetwork specification. May be IPv4
     * or IPv6.
     * Note that for a CIDR IP address block, the specified IP address
     * portion must be properly truncated (i.e. all the host bits must
     * be zero) or the input is considered malformed. For example,
     * &#34;192.0.2.0/24&#34; is accepted but &#34;192.0.2.1/24&#34; is not. Similarly,
     * for IPv6, &#34;2001:db8::/32&#34; is accepted whereas &#34;2001:db8::1/32&#34;
     * is not. The originating IP of a request must be in one of the
     * listed subnets in order for this Condition to be true.
     * If empty, all IP addresses are allowed.
     * 
     */
    public Optional<Output<List<String>>> ipSubnetworks() {
        return Optional.ofNullable(this.ipSubnetworks);
    }

    /**
     * An allowed list of members (users, service accounts).
     * Using groups is not supported yet.
     * The signed-in user originating the request must be a part of one
     * of the provided members. If not specified, a request may come
     * from any user (logged in/not logged in, not present in any
     * groups, etc.).
     * Formats: `user:{emailid}`, `serviceAccount:{emailid}`
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<String>> members;

    /**
     * @return An allowed list of members (users, service accounts).
     * Using groups is not supported yet.
     * The signed-in user originating the request must be a part of one
     * of the provided members. If not specified, a request may come
     * from any user (logged in/not logged in, not present in any
     * groups, etc.).
     * Formats: `user:{emailid}`, `serviceAccount:{emailid}`
     * 
     */
    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * Whether to negate the Condition. If true, the Condition becomes
     * a NAND over its non-empty fields, each field must be false for
     * the Condition overall to be satisfied. Defaults to false.
     * 
     */
    @Import(name="negate")
    private @Nullable Output<Boolean> negate;

    /**
     * @return Whether to negate the Condition. If true, the Condition becomes
     * a NAND over its non-empty fields, each field must be false for
     * the Condition overall to be satisfied. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> negate() {
        return Optional.ofNullable(this.negate);
    }

    /**
     * The request must originate from one of the provided
     * countries/regions.
     * Format: A valid ISO 3166-1 alpha-2 code.
     * 
     */
    @Import(name="regions")
    private @Nullable Output<List<String>> regions;

    /**
     * @return The request must originate from one of the provided
     * countries/regions.
     * Format: A valid ISO 3166-1 alpha-2 code.
     * 
     */
    public Optional<Output<List<String>>> regions() {
        return Optional.ofNullable(this.regions);
    }

    /**
     * A list of other access levels defined in the same Policy,
     * referenced by resource name. Referencing an AccessLevel which
     * does not exist is an error. All access levels listed must be
     * granted for the Condition to be true.
     * Format: accessPolicies/{policy_id}/accessLevels/{short_name}
     * 
     */
    @Import(name="requiredAccessLevels")
    private @Nullable Output<List<String>> requiredAccessLevels;

    /**
     * @return A list of other access levels defined in the same Policy,
     * referenced by resource name. Referencing an AccessLevel which
     * does not exist is an error. All access levels listed must be
     * granted for the Condition to be true.
     * Format: accessPolicies/{policy_id}/accessLevels/{short_name}
     * 
     */
    public Optional<Output<List<String>>> requiredAccessLevels() {
        return Optional.ofNullable(this.requiredAccessLevels);
    }

    private AccessLevelConditionState() {}

    private AccessLevelConditionState(AccessLevelConditionState $) {
        this.accessLevel = $.accessLevel;
        this.devicePolicy = $.devicePolicy;
        this.ipSubnetworks = $.ipSubnetworks;
        this.members = $.members;
        this.negate = $.negate;
        this.regions = $.regions;
        this.requiredAccessLevels = $.requiredAccessLevels;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessLevelConditionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessLevelConditionState $;

        public Builder() {
            $ = new AccessLevelConditionState();
        }

        public Builder(AccessLevelConditionState defaults) {
            $ = new AccessLevelConditionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessLevel The name of the Access Level to add this condition to.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(@Nullable Output<String> accessLevel) {
            $.accessLevel = accessLevel;
            return this;
        }

        /**
         * @param accessLevel The name of the Access Level to add this condition to.
         * 
         * @return builder
         * 
         */
        public Builder accessLevel(String accessLevel) {
            return accessLevel(Output.of(accessLevel));
        }

        /**
         * @param devicePolicy Device specific restrictions, all restrictions must hold for
         * the Condition to be true. If not specified, all devices are
         * allowed.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder devicePolicy(@Nullable Output<AccessLevelConditionDevicePolicyArgs> devicePolicy) {
            $.devicePolicy = devicePolicy;
            return this;
        }

        /**
         * @param devicePolicy Device specific restrictions, all restrictions must hold for
         * the Condition to be true. If not specified, all devices are
         * allowed.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder devicePolicy(AccessLevelConditionDevicePolicyArgs devicePolicy) {
            return devicePolicy(Output.of(devicePolicy));
        }

        /**
         * @param ipSubnetworks A list of CIDR block IP subnetwork specification. May be IPv4
         * or IPv6.
         * Note that for a CIDR IP address block, the specified IP address
         * portion must be properly truncated (i.e. all the host bits must
         * be zero) or the input is considered malformed. For example,
         * &#34;192.0.2.0/24&#34; is accepted but &#34;192.0.2.1/24&#34; is not. Similarly,
         * for IPv6, &#34;2001:db8::/32&#34; is accepted whereas &#34;2001:db8::1/32&#34;
         * is not. The originating IP of a request must be in one of the
         * listed subnets in order for this Condition to be true.
         * If empty, all IP addresses are allowed.
         * 
         * @return builder
         * 
         */
        public Builder ipSubnetworks(@Nullable Output<List<String>> ipSubnetworks) {
            $.ipSubnetworks = ipSubnetworks;
            return this;
        }

        /**
         * @param ipSubnetworks A list of CIDR block IP subnetwork specification. May be IPv4
         * or IPv6.
         * Note that for a CIDR IP address block, the specified IP address
         * portion must be properly truncated (i.e. all the host bits must
         * be zero) or the input is considered malformed. For example,
         * &#34;192.0.2.0/24&#34; is accepted but &#34;192.0.2.1/24&#34; is not. Similarly,
         * for IPv6, &#34;2001:db8::/32&#34; is accepted whereas &#34;2001:db8::1/32&#34;
         * is not. The originating IP of a request must be in one of the
         * listed subnets in order for this Condition to be true.
         * If empty, all IP addresses are allowed.
         * 
         * @return builder
         * 
         */
        public Builder ipSubnetworks(List<String> ipSubnetworks) {
            return ipSubnetworks(Output.of(ipSubnetworks));
        }

        /**
         * @param ipSubnetworks A list of CIDR block IP subnetwork specification. May be IPv4
         * or IPv6.
         * Note that for a CIDR IP address block, the specified IP address
         * portion must be properly truncated (i.e. all the host bits must
         * be zero) or the input is considered malformed. For example,
         * &#34;192.0.2.0/24&#34; is accepted but &#34;192.0.2.1/24&#34; is not. Similarly,
         * for IPv6, &#34;2001:db8::/32&#34; is accepted whereas &#34;2001:db8::1/32&#34;
         * is not. The originating IP of a request must be in one of the
         * listed subnets in order for this Condition to be true.
         * If empty, all IP addresses are allowed.
         * 
         * @return builder
         * 
         */
        public Builder ipSubnetworks(String... ipSubnetworks) {
            return ipSubnetworks(List.of(ipSubnetworks));
        }

        /**
         * @param members An allowed list of members (users, service accounts).
         * Using groups is not supported yet.
         * The signed-in user originating the request must be a part of one
         * of the provided members. If not specified, a request may come
         * from any user (logged in/not logged in, not present in any
         * groups, etc.).
         * Formats: `user:{emailid}`, `serviceAccount:{emailid}`
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members An allowed list of members (users, service accounts).
         * Using groups is not supported yet.
         * The signed-in user originating the request must be a part of one
         * of the provided members. If not specified, a request may come
         * from any user (logged in/not logged in, not present in any
         * groups, etc.).
         * Formats: `user:{emailid}`, `serviceAccount:{emailid}`
         * 
         * @return builder
         * 
         */
        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        /**
         * @param members An allowed list of members (users, service accounts).
         * Using groups is not supported yet.
         * The signed-in user originating the request must be a part of one
         * of the provided members. If not specified, a request may come
         * from any user (logged in/not logged in, not present in any
         * groups, etc.).
         * Formats: `user:{emailid}`, `serviceAccount:{emailid}`
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param negate Whether to negate the Condition. If true, the Condition becomes
         * a NAND over its non-empty fields, each field must be false for
         * the Condition overall to be satisfied. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder negate(@Nullable Output<Boolean> negate) {
            $.negate = negate;
            return this;
        }

        /**
         * @param negate Whether to negate the Condition. If true, the Condition becomes
         * a NAND over its non-empty fields, each field must be false for
         * the Condition overall to be satisfied. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder negate(Boolean negate) {
            return negate(Output.of(negate));
        }

        /**
         * @param regions The request must originate from one of the provided
         * countries/regions.
         * Format: A valid ISO 3166-1 alpha-2 code.
         * 
         * @return builder
         * 
         */
        public Builder regions(@Nullable Output<List<String>> regions) {
            $.regions = regions;
            return this;
        }

        /**
         * @param regions The request must originate from one of the provided
         * countries/regions.
         * Format: A valid ISO 3166-1 alpha-2 code.
         * 
         * @return builder
         * 
         */
        public Builder regions(List<String> regions) {
            return regions(Output.of(regions));
        }

        /**
         * @param regions The request must originate from one of the provided
         * countries/regions.
         * Format: A valid ISO 3166-1 alpha-2 code.
         * 
         * @return builder
         * 
         */
        public Builder regions(String... regions) {
            return regions(List.of(regions));
        }

        /**
         * @param requiredAccessLevels A list of other access levels defined in the same Policy,
         * referenced by resource name. Referencing an AccessLevel which
         * does not exist is an error. All access levels listed must be
         * granted for the Condition to be true.
         * Format: accessPolicies/{policy_id}/accessLevels/{short_name}
         * 
         * @return builder
         * 
         */
        public Builder requiredAccessLevels(@Nullable Output<List<String>> requiredAccessLevels) {
            $.requiredAccessLevels = requiredAccessLevels;
            return this;
        }

        /**
         * @param requiredAccessLevels A list of other access levels defined in the same Policy,
         * referenced by resource name. Referencing an AccessLevel which
         * does not exist is an error. All access levels listed must be
         * granted for the Condition to be true.
         * Format: accessPolicies/{policy_id}/accessLevels/{short_name}
         * 
         * @return builder
         * 
         */
        public Builder requiredAccessLevels(List<String> requiredAccessLevels) {
            return requiredAccessLevels(Output.of(requiredAccessLevels));
        }

        /**
         * @param requiredAccessLevels A list of other access levels defined in the same Policy,
         * referenced by resource name. Referencing an AccessLevel which
         * does not exist is an error. All access levels listed must be
         * granted for the Condition to be true.
         * Format: accessPolicies/{policy_id}/accessLevels/{short_name}
         * 
         * @return builder
         * 
         */
        public Builder requiredAccessLevels(String... requiredAccessLevels) {
            return requiredAccessLevels(List.of(requiredAccessLevels));
        }

        public AccessLevelConditionState build() {
            return $;
        }
    }

}
