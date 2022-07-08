// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BucketLoggingArgs extends com.pulumi.resources.ResourceArgs {

    public static final BucketLoggingArgs Empty = new BucketLoggingArgs();

    /**
     * The bucket that will receive log objects.
     * 
     */
    @Import(name="logBucket", required=true)
    private Output<String> logBucket;

    /**
     * @return The bucket that will receive log objects.
     * 
     */
    public Output<String> logBucket() {
        return this.logBucket;
    }

    /**
     * The object prefix for log objects. If it&#39;s not provided,
     * by default GCS sets this to this bucket&#39;s name.
     * 
     */
    @Import(name="logObjectPrefix")
    private @Nullable Output<String> logObjectPrefix;

    /**
     * @return The object prefix for log objects. If it&#39;s not provided,
     * by default GCS sets this to this bucket&#39;s name.
     * 
     */
    public Optional<Output<String>> logObjectPrefix() {
        return Optional.ofNullable(this.logObjectPrefix);
    }

    private BucketLoggingArgs() {}

    private BucketLoggingArgs(BucketLoggingArgs $) {
        this.logBucket = $.logBucket;
        this.logObjectPrefix = $.logObjectPrefix;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BucketLoggingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BucketLoggingArgs $;

        public Builder() {
            $ = new BucketLoggingArgs();
        }

        public Builder(BucketLoggingArgs defaults) {
            $ = new BucketLoggingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param logBucket The bucket that will receive log objects.
         * 
         * @return builder
         * 
         */
        public Builder logBucket(Output<String> logBucket) {
            $.logBucket = logBucket;
            return this;
        }

        /**
         * @param logBucket The bucket that will receive log objects.
         * 
         * @return builder
         * 
         */
        public Builder logBucket(String logBucket) {
            return logBucket(Output.of(logBucket));
        }

        /**
         * @param logObjectPrefix The object prefix for log objects. If it&#39;s not provided,
         * by default GCS sets this to this bucket&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder logObjectPrefix(@Nullable Output<String> logObjectPrefix) {
            $.logObjectPrefix = logObjectPrefix;
            return this;
        }

        /**
         * @param logObjectPrefix The object prefix for log objects. If it&#39;s not provided,
         * by default GCS sets this to this bucket&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder logObjectPrefix(String logObjectPrefix) {
            return logObjectPrefix(Output.of(logObjectPrefix));
        }

        public BucketLoggingArgs build() {
            $.logBucket = Objects.requireNonNull($.logBucket, "expected parameter 'logBucket' to be non-null");
            return $;
        }
    }

}
