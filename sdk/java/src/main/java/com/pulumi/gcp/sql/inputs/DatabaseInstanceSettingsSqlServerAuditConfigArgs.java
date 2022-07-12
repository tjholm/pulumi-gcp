// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.sql.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseInstanceSettingsSqlServerAuditConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseInstanceSettingsSqlServerAuditConfigArgs Empty = new DatabaseInstanceSettingsSqlServerAuditConfigArgs();

    /**
     * The name of the destination bucket (e.g., gs://mybucket).
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return The name of the destination bucket (e.g., gs://mybucket).
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }

    /**
     * How long to keep generated audit files. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    @Import(name="retentionInterval")
    private @Nullable Output<String> retentionInterval;

    /**
     * @return How long to keep generated audit files. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    public Optional<Output<String>> retentionInterval() {
        return Optional.ofNullable(this.retentionInterval);
    }

    /**
     * How often to upload generated audit files. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    @Import(name="uploadInterval")
    private @Nullable Output<String> uploadInterval;

    /**
     * @return How often to upload generated audit files. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    public Optional<Output<String>> uploadInterval() {
        return Optional.ofNullable(this.uploadInterval);
    }

    private DatabaseInstanceSettingsSqlServerAuditConfigArgs() {}

    private DatabaseInstanceSettingsSqlServerAuditConfigArgs(DatabaseInstanceSettingsSqlServerAuditConfigArgs $) {
        this.bucket = $.bucket;
        this.retentionInterval = $.retentionInterval;
        this.uploadInterval = $.uploadInterval;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseInstanceSettingsSqlServerAuditConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseInstanceSettingsSqlServerAuditConfigArgs $;

        public Builder() {
            $ = new DatabaseInstanceSettingsSqlServerAuditConfigArgs();
        }

        public Builder(DatabaseInstanceSettingsSqlServerAuditConfigArgs defaults) {
            $ = new DatabaseInstanceSettingsSqlServerAuditConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The name of the destination bucket (e.g., gs://mybucket).
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of the destination bucket (e.g., gs://mybucket).
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param retentionInterval How long to keep generated audit files. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder retentionInterval(@Nullable Output<String> retentionInterval) {
            $.retentionInterval = retentionInterval;
            return this;
        }

        /**
         * @param retentionInterval How long to keep generated audit files. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder retentionInterval(String retentionInterval) {
            return retentionInterval(Output.of(retentionInterval));
        }

        /**
         * @param uploadInterval How often to upload generated audit files. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder uploadInterval(@Nullable Output<String> uploadInterval) {
            $.uploadInterval = uploadInterval;
            return this;
        }

        /**
         * @param uploadInterval How often to upload generated audit files. A duration in seconds with up to nine fractional digits, terminated by &#39;s&#39;. Example: &#34;3.5s&#34;.
         * 
         * @return builder
         * 
         */
        public Builder uploadInterval(String uploadInterval) {
            return uploadInterval(Output.of(uploadInterval));
        }

        public DatabaseInstanceSettingsSqlServerAuditConfigArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            return $;
        }
    }

}
