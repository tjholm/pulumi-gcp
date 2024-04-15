// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class KeyRingIAMPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final KeyRingIAMPolicyArgs Empty = new KeyRingIAMPolicyArgs();

    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    @Import(name="keyRingId", required=true)
    private Output<String> keyRingId;

    /**
     * @return The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    public Output<String> keyRingId() {
        return this.keyRingId;
    }

    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Import(name="policyData", required=true)
    private Output<String> policyData;

    /**
     * @return The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    public Output<String> policyData() {
        return this.policyData;
    }

    private KeyRingIAMPolicyArgs() {}

    private KeyRingIAMPolicyArgs(KeyRingIAMPolicyArgs $) {
        this.keyRingId = $.keyRingId;
        this.policyData = $.policyData;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KeyRingIAMPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KeyRingIAMPolicyArgs $;

        public Builder() {
            $ = new KeyRingIAMPolicyArgs();
        }

        public Builder(KeyRingIAMPolicyArgs defaults) {
            $ = new KeyRingIAMPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param keyRingId The key ring ID, in the form
         * `{project_id}/{location_name}/{key_ring_name}` or
         * `{location_name}/{key_ring_name}`. In the second form, the provider&#39;s
         * project setting will be used as a fallback.
         * 
         * @return builder
         * 
         */
        public Builder keyRingId(Output<String> keyRingId) {
            $.keyRingId = keyRingId;
            return this;
        }

        /**
         * @param keyRingId The key ring ID, in the form
         * `{project_id}/{location_name}/{key_ring_name}` or
         * `{location_name}/{key_ring_name}`. In the second form, the provider&#39;s
         * project setting will be used as a fallback.
         * 
         * @return builder
         * 
         */
        public Builder keyRingId(String keyRingId) {
            return keyRingId(Output.of(keyRingId));
        }

        /**
         * @param policyData The policy data generated by
         * a `gcp.organizations.getIAMPolicy` data source.
         * 
         * @return builder
         * 
         */
        public Builder policyData(Output<String> policyData) {
            $.policyData = policyData;
            return this;
        }

        /**
         * @param policyData The policy data generated by
         * a `gcp.organizations.getIAMPolicy` data source.
         * 
         * @return builder
         * 
         */
        public Builder policyData(String policyData) {
            return policyData(Output.of(policyData));
        }

        public KeyRingIAMPolicyArgs build() {
            if ($.keyRingId == null) {
                throw new MissingRequiredPropertyException("KeyRingIAMPolicyArgs", "keyRingId");
            }
            if ($.policyData == null) {
                throw new MissingRequiredPropertyException("KeyRingIAMPolicyArgs", "policyData");
            }
            return $;
        }
    }

}
