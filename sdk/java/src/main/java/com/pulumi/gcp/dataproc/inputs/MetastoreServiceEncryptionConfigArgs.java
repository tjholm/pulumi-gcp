// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataproc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class MetastoreServiceEncryptionConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final MetastoreServiceEncryptionConfigArgs Empty = new MetastoreServiceEncryptionConfigArgs();

    /**
     * The fully qualified customer provided Cloud KMS key name to use for customer data encryption.
     * Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
     * 
     */
    @Import(name="kmsKey", required=true)
    private Output<String> kmsKey;

    /**
     * @return The fully qualified customer provided Cloud KMS key name to use for customer data encryption.
     * Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
     * 
     */
    public Output<String> kmsKey() {
        return this.kmsKey;
    }

    private MetastoreServiceEncryptionConfigArgs() {}

    private MetastoreServiceEncryptionConfigArgs(MetastoreServiceEncryptionConfigArgs $) {
        this.kmsKey = $.kmsKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MetastoreServiceEncryptionConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MetastoreServiceEncryptionConfigArgs $;

        public Builder() {
            $ = new MetastoreServiceEncryptionConfigArgs();
        }

        public Builder(MetastoreServiceEncryptionConfigArgs defaults) {
            $ = new MetastoreServiceEncryptionConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param kmsKey The fully qualified customer provided Cloud KMS key name to use for customer data encryption.
         * Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
         * 
         * @return builder
         * 
         */
        public Builder kmsKey(Output<String> kmsKey) {
            $.kmsKey = kmsKey;
            return this;
        }

        /**
         * @param kmsKey The fully qualified customer provided Cloud KMS key name to use for customer data encryption.
         * Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
         * 
         * @return builder
         * 
         */
        public Builder kmsKey(String kmsKey) {
            return kmsKey(Output.of(kmsKey));
        }

        public MetastoreServiceEncryptionConfigArgs build() {
            $.kmsKey = Objects.requireNonNull($.kmsKey, "expected parameter 'kmsKey' to be non-null");
            return $;
        }
    }

}
