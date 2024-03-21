// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.storage.outputs.TransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TransferJobTransferSpecAzureBlobStorageDataSource {
    /**
     * @return Credentials used to authenticate API requests to Azure block.
     * 
     */
    private @Nullable TransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials azureCredentials;
    /**
     * @return The container to transfer from the Azure Storage account.`
     * 
     */
    private String container;
    /**
     * @return Full Resource name of a secret in Secret Manager containing [SAS Credentials in JSON form](&lt;https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#azureblobstoragedata:~:text=begin%!w(MISSING)ith%!a(MISSING)%27/%!-(MISSING),credentialsSecret,-string&gt;). Service Agent for Storage Transfer must have permissions to access secret. If credentials_secret is specified, do not specify azure_credentials.`,
     * 
     */
    private @Nullable String credentialsSecret;
    /**
     * @return Root path to transfer objects. Must be an empty string or full path name that ends with a &#39;/&#39;. This field is treated as an object prefix. As such, it should generally not begin with a &#39;/&#39;.
     * 
     */
    private @Nullable String path;
    /**
     * @return The name of the Azure Storage account.
     * 
     */
    private String storageAccount;

    private TransferJobTransferSpecAzureBlobStorageDataSource() {}
    /**
     * @return Credentials used to authenticate API requests to Azure block.
     * 
     */
    public Optional<TransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials> azureCredentials() {
        return Optional.ofNullable(this.azureCredentials);
    }
    /**
     * @return The container to transfer from the Azure Storage account.`
     * 
     */
    public String container() {
        return this.container;
    }
    /**
     * @return Full Resource name of a secret in Secret Manager containing [SAS Credentials in JSON form](&lt;https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#azureblobstoragedata:~:text=begin%!w(MISSING)ith%!a(MISSING)%27/%!-(MISSING),credentialsSecret,-string&gt;). Service Agent for Storage Transfer must have permissions to access secret. If credentials_secret is specified, do not specify azure_credentials.`,
     * 
     */
    public Optional<String> credentialsSecret() {
        return Optional.ofNullable(this.credentialsSecret);
    }
    /**
     * @return Root path to transfer objects. Must be an empty string or full path name that ends with a &#39;/&#39;. This field is treated as an object prefix. As such, it should generally not begin with a &#39;/&#39;.
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }
    /**
     * @return The name of the Azure Storage account.
     * 
     */
    public String storageAccount() {
        return this.storageAccount;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TransferJobTransferSpecAzureBlobStorageDataSource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable TransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials azureCredentials;
        private String container;
        private @Nullable String credentialsSecret;
        private @Nullable String path;
        private String storageAccount;
        public Builder() {}
        public Builder(TransferJobTransferSpecAzureBlobStorageDataSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.azureCredentials = defaults.azureCredentials;
    	      this.container = defaults.container;
    	      this.credentialsSecret = defaults.credentialsSecret;
    	      this.path = defaults.path;
    	      this.storageAccount = defaults.storageAccount;
        }

        @CustomType.Setter
        public Builder azureCredentials(@Nullable TransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials azureCredentials) {

            this.azureCredentials = azureCredentials;
            return this;
        }
        @CustomType.Setter
        public Builder container(String container) {
            if (container == null) {
              throw new MissingRequiredPropertyException("TransferJobTransferSpecAzureBlobStorageDataSource", "container");
            }
            this.container = container;
            return this;
        }
        @CustomType.Setter
        public Builder credentialsSecret(@Nullable String credentialsSecret) {

            this.credentialsSecret = credentialsSecret;
            return this;
        }
        @CustomType.Setter
        public Builder path(@Nullable String path) {

            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder storageAccount(String storageAccount) {
            if (storageAccount == null) {
              throw new MissingRequiredPropertyException("TransferJobTransferSpecAzureBlobStorageDataSource", "storageAccount");
            }
            this.storageAccount = storageAccount;
            return this;
        }
        public TransferJobTransferSpecAzureBlobStorageDataSource build() {
            final var _resultValue = new TransferJobTransferSpecAzureBlobStorageDataSource();
            _resultValue.azureCredentials = azureCredentials;
            _resultValue.container = container;
            _resultValue.credentialsSecret = credentialsSecret;
            _resultValue.path = path;
            _resultValue.storageAccount = storageAccount;
            return _resultValue;
        }
    }
}
