// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.logging.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FolderBucketConfigState extends com.pulumi.resources.ResourceArgs {

    public static final FolderBucketConfigState Empty = new FolderBucketConfigState();

    /**
     * The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     * 
     */
    @Import(name="bucketId")
    private @Nullable Output<String> bucketId;

    /**
     * @return The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
     * 
     */
    public Optional<Output<String>> bucketId() {
        return Optional.ofNullable(this.bucketId);
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
     * The parent resource that contains the logging bucket.
     * 
     */
    @Import(name="folder")
    private @Nullable Output<String> folder;

    /**
     * @return The parent resource that contains the logging bucket.
     * 
     */
    public Optional<Output<String>> folder() {
        return Optional.ofNullable(this.folder);
    }

    /**
     * The bucket&#39;s lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
     * 
     */
    @Import(name="lifecycleState")
    private @Nullable Output<String> lifecycleState;

    /**
     * @return The bucket&#39;s lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
     * 
     */
    public Optional<Output<String>> lifecycleState() {
        return Optional.ofNullable(this.lifecycleState);
    }

    /**
     * The location of the bucket.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location of the bucket.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The resource name of the bucket. For example: &#34;folders/my-folder-id/locations/my-location/buckets/my-bucket-id&#34;
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The resource name of the bucket. For example: &#34;folders/my-folder-id/locations/my-location/buckets/my-bucket-id&#34;
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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

    private FolderBucketConfigState() {}

    private FolderBucketConfigState(FolderBucketConfigState $) {
        this.bucketId = $.bucketId;
        this.description = $.description;
        this.folder = $.folder;
        this.lifecycleState = $.lifecycleState;
        this.location = $.location;
        this.name = $.name;
        this.retentionDays = $.retentionDays;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FolderBucketConfigState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FolderBucketConfigState $;

        public Builder() {
            $ = new FolderBucketConfigState();
        }

        public Builder(FolderBucketConfigState defaults) {
            $ = new FolderBucketConfigState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucketId The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
         * 
         * @return builder
         * 
         */
        public Builder bucketId(@Nullable Output<String> bucketId) {
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
         * @param folder The parent resource that contains the logging bucket.
         * 
         * @return builder
         * 
         */
        public Builder folder(@Nullable Output<String> folder) {
            $.folder = folder;
            return this;
        }

        /**
         * @param folder The parent resource that contains the logging bucket.
         * 
         * @return builder
         * 
         */
        public Builder folder(String folder) {
            return folder(Output.of(folder));
        }

        /**
         * @param lifecycleState The bucket&#39;s lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
         * 
         * @return builder
         * 
         */
        public Builder lifecycleState(@Nullable Output<String> lifecycleState) {
            $.lifecycleState = lifecycleState;
            return this;
        }

        /**
         * @param lifecycleState The bucket&#39;s lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
         * 
         * @return builder
         * 
         */
        public Builder lifecycleState(String lifecycleState) {
            return lifecycleState(Output.of(lifecycleState));
        }

        /**
         * @param location The location of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
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
         * @param name The resource name of the bucket. For example: &#34;folders/my-folder-id/locations/my-location/buckets/my-bucket-id&#34;
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The resource name of the bucket. For example: &#34;folders/my-folder-id/locations/my-location/buckets/my-bucket-id&#34;
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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

        public FolderBucketConfigState build() {
            return $;
        }
    }

}
