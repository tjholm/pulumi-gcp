// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class AwsClusterControlPlaneSshConfig {
    /**
     * @return The name of the EC2 key pair used to login into cluster machines.
     * 
     */
    private final String ec2KeyPair;

    @CustomType.Constructor
    private AwsClusterControlPlaneSshConfig(@CustomType.Parameter("ec2KeyPair") String ec2KeyPair) {
        this.ec2KeyPair = ec2KeyPair;
    }

    /**
     * @return The name of the EC2 key pair used to login into cluster machines.
     * 
     */
    public String ec2KeyPair() {
        return this.ec2KeyPair;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AwsClusterControlPlaneSshConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String ec2KeyPair;

        public Builder() {
    	      // Empty
        }

        public Builder(AwsClusterControlPlaneSshConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ec2KeyPair = defaults.ec2KeyPair;
        }

        public Builder ec2KeyPair(String ec2KeyPair) {
            this.ec2KeyPair = Objects.requireNonNull(ec2KeyPair);
            return this;
        }        public AwsClusterControlPlaneSshConfig build() {
            return new AwsClusterControlPlaneSshConfig(ec2KeyPair);
        }
    }
}
