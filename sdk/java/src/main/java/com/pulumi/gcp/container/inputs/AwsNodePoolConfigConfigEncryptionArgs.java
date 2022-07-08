// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class AwsNodePoolConfigConfigEncryptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final AwsNodePoolConfigConfigEncryptionArgs Empty = new AwsNodePoolConfigConfigEncryptionArgs();

    /**
     * Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
     * 
     */
    @Import(name="kmsKeyArn", required=true)
    private Output<String> kmsKeyArn;

    /**
     * @return Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
     * 
     */
    public Output<String> kmsKeyArn() {
        return this.kmsKeyArn;
    }

    private AwsNodePoolConfigConfigEncryptionArgs() {}

    private AwsNodePoolConfigConfigEncryptionArgs(AwsNodePoolConfigConfigEncryptionArgs $) {
        this.kmsKeyArn = $.kmsKeyArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AwsNodePoolConfigConfigEncryptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AwsNodePoolConfigConfigEncryptionArgs $;

        public Builder() {
            $ = new AwsNodePoolConfigConfigEncryptionArgs();
        }

        public Builder(AwsNodePoolConfigConfigEncryptionArgs defaults) {
            $ = new AwsNodePoolConfigConfigEncryptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param kmsKeyArn Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(Output<String> kmsKeyArn) {
            $.kmsKeyArn = kmsKeyArn;
            return this;
        }

        /**
         * @param kmsKeyArn Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(String kmsKeyArn) {
            return kmsKeyArn(Output.of(kmsKeyArn));
        }

        public AwsNodePoolConfigConfigEncryptionArgs build() {
            $.kmsKeyArn = Objects.requireNonNull($.kmsKeyArn, "expected parameter 'kmsKeyArn' to be non-null");
            return $;
        }
    }

}
