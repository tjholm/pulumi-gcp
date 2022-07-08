// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.logging;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BillingAccountBucketConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final BillingAccountBucketConfigArgs Empty = new BillingAccountBucketConfigArgs();

    /**
     * The parent resource that contains the logging bucket.
     * 
     */
    @Import(name="billingAccount", required=true)
    private Output<String> billingAccount;

    /**
     * @return The parent resource that contains the logging bucket.
     * 
     */
    public Output<String> billingAccount() {
        return this.billingAccount;
    }

    /**
     * The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     * 
     */
    @Import(name="bucketId", required=true)
    private Output<String> bucketId;

    /**
     * @return The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     * 
     */
    public Output<String> bucketId() {
        return this.bucketId;
    }

    /**
     * Describes this bucket.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Describes this bucket.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The location of the bucket.
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return The location of the bucket.
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    /**
     * Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
     * 
     */
    @Import(name="retentionDays")
    private @Nullable Output<Integer> retentionDays;

    /**
     * @return Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
     * 
     */
    public Optional<Output<Integer>> retentionDays() {
        return Optional.ofNullable(this.retentionDays);
    }

    private BillingAccountBucketConfigArgs() {}

    private BillingAccountBucketConfigArgs(BillingAccountBucketConfigArgs $) {
        this.billingAccount = $.billingAccount;
        this.bucketId = $.bucketId;
        this.description = $.description;
        this.location = $.location;
        this.retentionDays = $.retentionDays;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BillingAccountBucketConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BillingAccountBucketConfigArgs $;

        public Builder() {
            $ = new BillingAccountBucketConfigArgs();
        }

        public Builder(BillingAccountBucketConfigArgs defaults) {
            $ = new BillingAccountBucketConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param billingAccount The parent resource that contains the logging bucket.
         * 
         * @return builder
         * 
         */
        public Builder billingAccount(Output<String> billingAccount) {
            $.billingAccount = billingAccount;
            return this;
        }

        /**
         * @param billingAccount The parent resource that contains the logging bucket.
         * 
         * @return builder
         * 
         */
        public Builder billingAccount(String billingAccount) {
            return billingAccount(Output.of(billingAccount));
        }

        /**
         * @param bucketId The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
         * 
         * @return builder
         * 
         */
        public Builder bucketId(Output<String> bucketId) {
            $.bucketId = bucketId;
            return this;
        }

        /**
         * @param bucketId The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
         * 
         * @return builder
         * 
         */
        public Builder bucketId(String bucketId) {
            return bucketId(Output.of(bucketId));
        }

        /**
         * @param description Describes this bucket.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Describes this bucket.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param location The location of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param retentionDays Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
         * 
         * @return builder
         * 
         */
        public Builder retentionDays(@Nullable Output<Integer> retentionDays) {
            $.retentionDays = retentionDays;
            return this;
        }

        /**
         * @param retentionDays Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
         * 
         * @return builder
         * 
         */
        public Builder retentionDays(Integer retentionDays) {
            return retentionDays(Output.of(retentionDays));
        }

        public BillingAccountBucketConfigArgs build() {
            $.billingAccount = Objects.requireNonNull($.billingAccount, "expected parameter 'billingAccount' to be non-null");
            $.bucketId = Objects.requireNonNull($.bucketId, "expected parameter 'bucketId' to be non-null");
            $.location = Objects.requireNonNull($.location, "expected parameter 'location' to be non-null");
            return $;
        }
    }

}
