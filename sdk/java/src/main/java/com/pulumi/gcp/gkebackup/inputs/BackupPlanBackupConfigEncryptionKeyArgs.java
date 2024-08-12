// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkebackup.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class BackupPlanBackupConfigEncryptionKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final BackupPlanBackupConfigEncryptionKeyArgs Empty = new BackupPlanBackupConfigEncryptionKeyArgs();

    /**
     * Google Cloud KMS encryption key. Format: projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*
     * 
     */
    @Import(name="gcpKmsEncryptionKey", required=true)
    private Output<String> gcpKmsEncryptionKey;

    /**
     * @return Google Cloud KMS encryption key. Format: projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*
     * 
     */
    public Output<String> gcpKmsEncryptionKey() {
        return this.gcpKmsEncryptionKey;
    }

    private BackupPlanBackupConfigEncryptionKeyArgs() {}

    private BackupPlanBackupConfigEncryptionKeyArgs(BackupPlanBackupConfigEncryptionKeyArgs $) {
        this.gcpKmsEncryptionKey = $.gcpKmsEncryptionKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackupPlanBackupConfigEncryptionKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackupPlanBackupConfigEncryptionKeyArgs $;

        public Builder() {
            $ = new BackupPlanBackupConfigEncryptionKeyArgs();
        }

        public Builder(BackupPlanBackupConfigEncryptionKeyArgs defaults) {
            $ = new BackupPlanBackupConfigEncryptionKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param gcpKmsEncryptionKey Google Cloud KMS encryption key. Format: projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*
         * 
         * @return builder
         * 
         */
        public Builder gcpKmsEncryptionKey(Output<String> gcpKmsEncryptionKey) {
            $.gcpKmsEncryptionKey = gcpKmsEncryptionKey;
            return this;
        }

        /**
         * @param gcpKmsEncryptionKey Google Cloud KMS encryption key. Format: projects/*&#47;locations/*&#47;keyRings/*&#47;cryptoKeys/*
         * 
         * @return builder
         * 
         */
        public Builder gcpKmsEncryptionKey(String gcpKmsEncryptionKey) {
            return gcpKmsEncryptionKey(Output.of(gcpKmsEncryptionKey));
        }

        public BackupPlanBackupConfigEncryptionKeyArgs build() {
            if ($.gcpKmsEncryptionKey == null) {
                throw new MissingRequiredPropertyException("BackupPlanBackupConfigEncryptionKeyArgs", "gcpKmsEncryptionKey");
            }
            return $;
        }
    }

}
