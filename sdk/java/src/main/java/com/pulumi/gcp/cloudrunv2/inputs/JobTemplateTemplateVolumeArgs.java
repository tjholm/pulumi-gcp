// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudrunv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.cloudrunv2.inputs.JobTemplateTemplateVolumeCloudSqlInstanceArgs;
import com.pulumi.gcp.cloudrunv2.inputs.JobTemplateTemplateVolumeEmptyDirArgs;
import com.pulumi.gcp.cloudrunv2.inputs.JobTemplateTemplateVolumeSecretArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class JobTemplateTemplateVolumeArgs extends com.pulumi.resources.ResourceArgs {

    public static final JobTemplateTemplateVolumeArgs Empty = new JobTemplateTemplateVolumeArgs();

    /**
     * For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
     * Structure is documented below.
     * 
     */
    @Import(name="cloudSqlInstance")
    private @Nullable Output<JobTemplateTemplateVolumeCloudSqlInstanceArgs> cloudSqlInstance;

    /**
     * @return For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
     * Structure is documented below.
     * 
     */
    public Optional<Output<JobTemplateTemplateVolumeCloudSqlInstanceArgs>> cloudSqlInstance() {
        return Optional.ofNullable(this.cloudSqlInstance);
    }

    @Import(name="emptyDir")
    private @Nullable Output<JobTemplateTemplateVolumeEmptyDirArgs> emptyDir;

    public Optional<Output<JobTemplateTemplateVolumeEmptyDirArgs>> emptyDir() {
        return Optional.ofNullable(this.emptyDir);
    }

    /**
     * Volume&#39;s name.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Volume&#39;s name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
     * Structure is documented below.
     * 
     */
    @Import(name="secret")
    private @Nullable Output<JobTemplateTemplateVolumeSecretArgs> secret;

    /**
     * @return Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
     * Structure is documented below.
     * 
     */
    public Optional<Output<JobTemplateTemplateVolumeSecretArgs>> secret() {
        return Optional.ofNullable(this.secret);
    }

    private JobTemplateTemplateVolumeArgs() {}

    private JobTemplateTemplateVolumeArgs(JobTemplateTemplateVolumeArgs $) {
        this.cloudSqlInstance = $.cloudSqlInstance;
        this.emptyDir = $.emptyDir;
        this.name = $.name;
        this.secret = $.secret;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JobTemplateTemplateVolumeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JobTemplateTemplateVolumeArgs $;

        public Builder() {
            $ = new JobTemplateTemplateVolumeArgs();
        }

        public Builder(JobTemplateTemplateVolumeArgs defaults) {
            $ = new JobTemplateTemplateVolumeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cloudSqlInstance For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cloudSqlInstance(@Nullable Output<JobTemplateTemplateVolumeCloudSqlInstanceArgs> cloudSqlInstance) {
            $.cloudSqlInstance = cloudSqlInstance;
            return this;
        }

        /**
         * @param cloudSqlInstance For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cloudSqlInstance(JobTemplateTemplateVolumeCloudSqlInstanceArgs cloudSqlInstance) {
            return cloudSqlInstance(Output.of(cloudSqlInstance));
        }

        public Builder emptyDir(@Nullable Output<JobTemplateTemplateVolumeEmptyDirArgs> emptyDir) {
            $.emptyDir = emptyDir;
            return this;
        }

        public Builder emptyDir(JobTemplateTemplateVolumeEmptyDirArgs emptyDir) {
            return emptyDir(Output.of(emptyDir));
        }

        /**
         * @param name Volume&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Volume&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param secret Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secret(@Nullable Output<JobTemplateTemplateVolumeSecretArgs> secret) {
            $.secret = secret;
            return this;
        }

        /**
         * @param secret Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secret(JobTemplateTemplateVolumeSecretArgs secret) {
            return secret(Output.of(secret));
        }

        public JobTemplateTemplateVolumeArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
