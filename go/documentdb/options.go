// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.7, generator: @autorest/go@4.0.0-preview.72)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package documentdb

// CassandraResourcesClientBeginCreateUpdateCassandraKeyspaceOptions contains the optional parameters for the CassandraResourcesClient.BeginCreateUpdateCassandraKeyspace
// method.
type CassandraResourcesClientBeginCreateUpdateCassandraKeyspaceOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginCreateUpdateCassandraTableOptions contains the optional parameters for the CassandraResourcesClient.BeginCreateUpdateCassandraTable
// method.
type CassandraResourcesClientBeginCreateUpdateCassandraTableOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginDeleteCassandraKeyspaceOptions contains the optional parameters for the CassandraResourcesClient.BeginDeleteCassandraKeyspace
// method.
type CassandraResourcesClientBeginDeleteCassandraKeyspaceOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginDeleteCassandraTableOptions contains the optional parameters for the CassandraResourcesClient.BeginDeleteCassandraTable
// method.
type CassandraResourcesClientBeginDeleteCassandraTableOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginMigrateCassandraKeyspaceToAutoscaleOptions contains the optional parameters for the CassandraResourcesClient.BeginMigrateCassandraKeyspaceToAutoscale
// method.
type CassandraResourcesClientBeginMigrateCassandraKeyspaceToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginMigrateCassandraKeyspaceToManualThroughputOptions contains the optional parameters for the
// CassandraResourcesClient.BeginMigrateCassandraKeyspaceToManualThroughput method.
type CassandraResourcesClientBeginMigrateCassandraKeyspaceToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginMigrateCassandraTableToAutoscaleOptions contains the optional parameters for the CassandraResourcesClient.BeginMigrateCassandraTableToAutoscale
// method.
type CassandraResourcesClientBeginMigrateCassandraTableToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginMigrateCassandraTableToManualThroughputOptions contains the optional parameters for the CassandraResourcesClient.BeginMigrateCassandraTableToManualThroughput
// method.
type CassandraResourcesClientBeginMigrateCassandraTableToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginUpdateCassandraKeyspaceThroughputOptions contains the optional parameters for the CassandraResourcesClient.BeginUpdateCassandraKeyspaceThroughput
// method.
type CassandraResourcesClientBeginUpdateCassandraKeyspaceThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientBeginUpdateCassandraTableThroughputOptions contains the optional parameters for the CassandraResourcesClient.BeginUpdateCassandraTableThroughput
// method.
type CassandraResourcesClientBeginUpdateCassandraTableThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CassandraResourcesClientGetCassandraKeyspaceOptions contains the optional parameters for the CassandraResourcesClient.GetCassandraKeyspace
// method.
type CassandraResourcesClientGetCassandraKeyspaceOptions struct {
	// placeholder for future optional parameters
}

// CassandraResourcesClientGetCassandraKeyspaceThroughputOptions contains the optional parameters for the CassandraResourcesClient.GetCassandraKeyspaceThroughput
// method.
type CassandraResourcesClientGetCassandraKeyspaceThroughputOptions struct {
	// placeholder for future optional parameters
}

// CassandraResourcesClientGetCassandraTableOptions contains the optional parameters for the CassandraResourcesClient.GetCassandraTable
// method.
type CassandraResourcesClientGetCassandraTableOptions struct {
	// placeholder for future optional parameters
}

// CassandraResourcesClientGetCassandraTableThroughputOptions contains the optional parameters for the CassandraResourcesClient.GetCassandraTableThroughput
// method.
type CassandraResourcesClientGetCassandraTableThroughputOptions struct {
	// placeholder for future optional parameters
}

// CassandraResourcesClientListCassandraKeyspacesOptions contains the optional parameters for the CassandraResourcesClient.NewListCassandraKeyspacesPager
// method.
type CassandraResourcesClientListCassandraKeyspacesOptions struct {
	// placeholder for future optional parameters
}

// CassandraResourcesClientListCassandraTablesOptions contains the optional parameters for the CassandraResourcesClient.NewListCassandraTablesPager
// method.
type CassandraResourcesClientListCassandraTablesOptions struct {
	// placeholder for future optional parameters
}

// CollectionClientListMetricDefinitionsOptions contains the optional parameters for the CollectionClient.NewListMetricDefinitionsPager
// method.
type CollectionClientListMetricDefinitionsOptions struct {
	// placeholder for future optional parameters
}

// CollectionClientListMetricsOptions contains the optional parameters for the CollectionClient.NewListMetricsPager method.
type CollectionClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// CollectionClientListUsagesOptions contains the optional parameters for the CollectionClient.NewListUsagesPager method.
type CollectionClientListUsagesOptions struct {
// An OData filter expression that describes a subset of usages to return. The supported parameter is name.value (name of
// the metric, can have an or of multiple names).
	Filter *string
}

// CollectionPartitionClientListMetricsOptions contains the optional parameters for the CollectionPartitionClient.NewListMetricsPager
// method.
type CollectionPartitionClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// CollectionPartitionClientListUsagesOptions contains the optional parameters for the CollectionPartitionClient.NewListUsagesPager
// method.
type CollectionPartitionClientListUsagesOptions struct {
// An OData filter expression that describes a subset of usages to return. The supported parameter is name.value (name of
// the metric, can have an or of multiple names).
	Filter *string
}

// CollectionPartitionRegionClientListMetricsOptions contains the optional parameters for the CollectionPartitionRegionClient.NewListMetricsPager
// method.
type CollectionPartitionRegionClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// CollectionRegionClientListMetricsOptions contains the optional parameters for the CollectionRegionClient.NewListMetricsPager
// method.
type CollectionRegionClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountRegionClientListMetricsOptions contains the optional parameters for the DatabaseAccountRegionClient.NewListMetricsPager
// method.
type DatabaseAccountRegionClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabaseAccountsClient.BeginCreateOrUpdate
// method.
type DatabaseAccountsClientBeginCreateOrUpdateOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DatabaseAccountsClientBeginDeleteOptions contains the optional parameters for the DatabaseAccountsClient.BeginDelete method.
type DatabaseAccountsClientBeginDeleteOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DatabaseAccountsClientBeginFailoverPriorityChangeOptions contains the optional parameters for the DatabaseAccountsClient.BeginFailoverPriorityChange
// method.
type DatabaseAccountsClientBeginFailoverPriorityChangeOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DatabaseAccountsClientBeginOfflineRegionOptions contains the optional parameters for the DatabaseAccountsClient.BeginOfflineRegion
// method.
type DatabaseAccountsClientBeginOfflineRegionOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DatabaseAccountsClientBeginOnlineRegionOptions contains the optional parameters for the DatabaseAccountsClient.BeginOnlineRegion
// method.
type DatabaseAccountsClientBeginOnlineRegionOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DatabaseAccountsClientBeginRegenerateKeyOptions contains the optional parameters for the DatabaseAccountsClient.BeginRegenerateKey
// method.
type DatabaseAccountsClientBeginRegenerateKeyOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DatabaseAccountsClientBeginUpdateOptions contains the optional parameters for the DatabaseAccountsClient.BeginUpdate method.
type DatabaseAccountsClientBeginUpdateOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DatabaseAccountsClientCheckNameExistsOptions contains the optional parameters for the DatabaseAccountsClient.CheckNameExists
// method.
type DatabaseAccountsClientCheckNameExistsOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientGetOptions contains the optional parameters for the DatabaseAccountsClient.Get method.
type DatabaseAccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientGetReadOnlyKeysOptions contains the optional parameters for the DatabaseAccountsClient.GetReadOnlyKeys
// method.
type DatabaseAccountsClientGetReadOnlyKeysOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientListByResourceGroupOptions contains the optional parameters for the DatabaseAccountsClient.NewListByResourceGroupPager
// method.
type DatabaseAccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientListConnectionStringsOptions contains the optional parameters for the DatabaseAccountsClient.ListConnectionStrings
// method.
type DatabaseAccountsClientListConnectionStringsOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientListKeysOptions contains the optional parameters for the DatabaseAccountsClient.ListKeys method.
type DatabaseAccountsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientListMetricDefinitionsOptions contains the optional parameters for the DatabaseAccountsClient.NewListMetricDefinitionsPager
// method.
type DatabaseAccountsClientListMetricDefinitionsOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientListMetricsOptions contains the optional parameters for the DatabaseAccountsClient.NewListMetricsPager
// method.
type DatabaseAccountsClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientListOptions contains the optional parameters for the DatabaseAccountsClient.NewListPager method.
type DatabaseAccountsClientListOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientListReadOnlyKeysOptions contains the optional parameters for the DatabaseAccountsClient.ListReadOnlyKeys
// method.
type DatabaseAccountsClientListReadOnlyKeysOptions struct {
	// placeholder for future optional parameters
}

// DatabaseAccountsClientListUsagesOptions contains the optional parameters for the DatabaseAccountsClient.NewListUsagesPager
// method.
type DatabaseAccountsClientListUsagesOptions struct {
// An OData filter expression that describes a subset of usages to return. The supported parameter is name.value (name of
// the metric, can have an or of multiple names).
	Filter *string
}

// DatabaseClientListMetricDefinitionsOptions contains the optional parameters for the DatabaseClient.NewListMetricDefinitionsPager
// method.
type DatabaseClientListMetricDefinitionsOptions struct {
	// placeholder for future optional parameters
}

// DatabaseClientListMetricsOptions contains the optional parameters for the DatabaseClient.NewListMetricsPager method.
type DatabaseClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// DatabaseClientListUsagesOptions contains the optional parameters for the DatabaseClient.NewListUsagesPager method.
type DatabaseClientListUsagesOptions struct {
// An OData filter expression that describes a subset of usages to return. The supported parameter is name.value (name of
// the metric, can have an or of multiple names).
	Filter *string
}

// GremlinResourcesClientBeginCreateUpdateGremlinDatabaseOptions contains the optional parameters for the GremlinResourcesClient.BeginCreateUpdateGremlinDatabase
// method.
type GremlinResourcesClientBeginCreateUpdateGremlinDatabaseOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginCreateUpdateGremlinGraphOptions contains the optional parameters for the GremlinResourcesClient.BeginCreateUpdateGremlinGraph
// method.
type GremlinResourcesClientBeginCreateUpdateGremlinGraphOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginDeleteGremlinDatabaseOptions contains the optional parameters for the GremlinResourcesClient.BeginDeleteGremlinDatabase
// method.
type GremlinResourcesClientBeginDeleteGremlinDatabaseOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginDeleteGremlinGraphOptions contains the optional parameters for the GremlinResourcesClient.BeginDeleteGremlinGraph
// method.
type GremlinResourcesClientBeginDeleteGremlinGraphOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginMigrateGremlinDatabaseToAutoscaleOptions contains the optional parameters for the GremlinResourcesClient.BeginMigrateGremlinDatabaseToAutoscale
// method.
type GremlinResourcesClientBeginMigrateGremlinDatabaseToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginMigrateGremlinDatabaseToManualThroughputOptions contains the optional parameters for the GremlinResourcesClient.BeginMigrateGremlinDatabaseToManualThroughput
// method.
type GremlinResourcesClientBeginMigrateGremlinDatabaseToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginMigrateGremlinGraphToAutoscaleOptions contains the optional parameters for the GremlinResourcesClient.BeginMigrateGremlinGraphToAutoscale
// method.
type GremlinResourcesClientBeginMigrateGremlinGraphToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginMigrateGremlinGraphToManualThroughputOptions contains the optional parameters for the GremlinResourcesClient.BeginMigrateGremlinGraphToManualThroughput
// method.
type GremlinResourcesClientBeginMigrateGremlinGraphToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginUpdateGremlinDatabaseThroughputOptions contains the optional parameters for the GremlinResourcesClient.BeginUpdateGremlinDatabaseThroughput
// method.
type GremlinResourcesClientBeginUpdateGremlinDatabaseThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientBeginUpdateGremlinGraphThroughputOptions contains the optional parameters for the GremlinResourcesClient.BeginUpdateGremlinGraphThroughput
// method.
type GremlinResourcesClientBeginUpdateGremlinGraphThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// GremlinResourcesClientGetGremlinDatabaseOptions contains the optional parameters for the GremlinResourcesClient.GetGremlinDatabase
// method.
type GremlinResourcesClientGetGremlinDatabaseOptions struct {
	// placeholder for future optional parameters
}

// GremlinResourcesClientGetGremlinDatabaseThroughputOptions contains the optional parameters for the GremlinResourcesClient.GetGremlinDatabaseThroughput
// method.
type GremlinResourcesClientGetGremlinDatabaseThroughputOptions struct {
	// placeholder for future optional parameters
}

// GremlinResourcesClientGetGremlinGraphOptions contains the optional parameters for the GremlinResourcesClient.GetGremlinGraph
// method.
type GremlinResourcesClientGetGremlinGraphOptions struct {
	// placeholder for future optional parameters
}

// GremlinResourcesClientGetGremlinGraphThroughputOptions contains the optional parameters for the GremlinResourcesClient.GetGremlinGraphThroughput
// method.
type GremlinResourcesClientGetGremlinGraphThroughputOptions struct {
	// placeholder for future optional parameters
}

// GremlinResourcesClientListGremlinDatabasesOptions contains the optional parameters for the GremlinResourcesClient.NewListGremlinDatabasesPager
// method.
type GremlinResourcesClientListGremlinDatabasesOptions struct {
	// placeholder for future optional parameters
}

// GremlinResourcesClientListGremlinGraphsOptions contains the optional parameters for the GremlinResourcesClient.NewListGremlinGraphsPager
// method.
type GremlinResourcesClientListGremlinGraphsOptions struct {
	// placeholder for future optional parameters
}

// LocationsClientGetOptions contains the optional parameters for the LocationsClient.Get method.
type LocationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// LocationsClientListOptions contains the optional parameters for the LocationsClient.NewListPager method.
type LocationsClientListOptions struct {
	// placeholder for future optional parameters
}

// MongoDBResourcesClientBeginCreateUpdateMongoDBCollectionOptions contains the optional parameters for the MongoDBResourcesClient.BeginCreateUpdateMongoDBCollection
// method.
type MongoDBResourcesClientBeginCreateUpdateMongoDBCollectionOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginCreateUpdateMongoDBDatabaseOptions contains the optional parameters for the MongoDBResourcesClient.BeginCreateUpdateMongoDBDatabase
// method.
type MongoDBResourcesClientBeginCreateUpdateMongoDBDatabaseOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginDeleteMongoDBCollectionOptions contains the optional parameters for the MongoDBResourcesClient.BeginDeleteMongoDBCollection
// method.
type MongoDBResourcesClientBeginDeleteMongoDBCollectionOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginDeleteMongoDBDatabaseOptions contains the optional parameters for the MongoDBResourcesClient.BeginDeleteMongoDBDatabase
// method.
type MongoDBResourcesClientBeginDeleteMongoDBDatabaseOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginMigrateMongoDBCollectionToAutoscaleOptions contains the optional parameters for the MongoDBResourcesClient.BeginMigrateMongoDBCollectionToAutoscale
// method.
type MongoDBResourcesClientBeginMigrateMongoDBCollectionToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginMigrateMongoDBCollectionToManualThroughputOptions contains the optional parameters for the MongoDBResourcesClient.BeginMigrateMongoDBCollectionToManualThroughput
// method.
type MongoDBResourcesClientBeginMigrateMongoDBCollectionToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginMigrateMongoDBDatabaseToAutoscaleOptions contains the optional parameters for the MongoDBResourcesClient.BeginMigrateMongoDBDatabaseToAutoscale
// method.
type MongoDBResourcesClientBeginMigrateMongoDBDatabaseToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginMigrateMongoDBDatabaseToManualThroughputOptions contains the optional parameters for the MongoDBResourcesClient.BeginMigrateMongoDBDatabaseToManualThroughput
// method.
type MongoDBResourcesClientBeginMigrateMongoDBDatabaseToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginUpdateMongoDBCollectionThroughputOptions contains the optional parameters for the MongoDBResourcesClient.BeginUpdateMongoDBCollectionThroughput
// method.
type MongoDBResourcesClientBeginUpdateMongoDBCollectionThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientBeginUpdateMongoDBDatabaseThroughputOptions contains the optional parameters for the MongoDBResourcesClient.BeginUpdateMongoDBDatabaseThroughput
// method.
type MongoDBResourcesClientBeginUpdateMongoDBDatabaseThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MongoDBResourcesClientGetMongoDBCollectionOptions contains the optional parameters for the MongoDBResourcesClient.GetMongoDBCollection
// method.
type MongoDBResourcesClientGetMongoDBCollectionOptions struct {
	// placeholder for future optional parameters
}

// MongoDBResourcesClientGetMongoDBCollectionThroughputOptions contains the optional parameters for the MongoDBResourcesClient.GetMongoDBCollectionThroughput
// method.
type MongoDBResourcesClientGetMongoDBCollectionThroughputOptions struct {
	// placeholder for future optional parameters
}

// MongoDBResourcesClientGetMongoDBDatabaseOptions contains the optional parameters for the MongoDBResourcesClient.GetMongoDBDatabase
// method.
type MongoDBResourcesClientGetMongoDBDatabaseOptions struct {
	// placeholder for future optional parameters
}

// MongoDBResourcesClientGetMongoDBDatabaseThroughputOptions contains the optional parameters for the MongoDBResourcesClient.GetMongoDBDatabaseThroughput
// method.
type MongoDBResourcesClientGetMongoDBDatabaseThroughputOptions struct {
	// placeholder for future optional parameters
}

// MongoDBResourcesClientListMongoDBCollectionsOptions contains the optional parameters for the MongoDBResourcesClient.NewListMongoDBCollectionsPager
// method.
type MongoDBResourcesClientListMongoDBCollectionsOptions struct {
	// placeholder for future optional parameters
}

// MongoDBResourcesClientListMongoDBDatabasesOptions contains the optional parameters for the MongoDBResourcesClient.NewListMongoDBDatabasesPager
// method.
type MongoDBResourcesClientListMongoDBDatabasesOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PartitionKeyRangeIDClientListMetricsOptions contains the optional parameters for the PartitionKeyRangeIDClient.NewListMetricsPager
// method.
type PartitionKeyRangeIDClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// PartitionKeyRangeIDRegionClientListMetricsOptions contains the optional parameters for the PartitionKeyRangeIDRegionClient.NewListMetricsPager
// method.
type PartitionKeyRangeIDRegionClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// PercentileClientListMetricsOptions contains the optional parameters for the PercentileClient.NewListMetricsPager method.
type PercentileClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// PercentileSourceTargetClientListMetricsOptions contains the optional parameters for the PercentileSourceTargetClient.NewListMetricsPager
// method.
type PercentileSourceTargetClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// PercentileTargetClientListMetricsOptions contains the optional parameters for the PercentileTargetClient.NewListMetricsPager
// method.
type PercentileTargetClientListMetricsOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientBeginCreateUpdateClientEncryptionKeyOptions contains the optional parameters for the SQLResourcesClient.BeginCreateUpdateClientEncryptionKey
// method.
type SQLResourcesClientBeginCreateUpdateClientEncryptionKeyOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginCreateUpdateSQLContainerOptions contains the optional parameters for the SQLResourcesClient.BeginCreateUpdateSQLContainer
// method.
type SQLResourcesClientBeginCreateUpdateSQLContainerOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginCreateUpdateSQLDatabaseOptions contains the optional parameters for the SQLResourcesClient.BeginCreateUpdateSQLDatabase
// method.
type SQLResourcesClientBeginCreateUpdateSQLDatabaseOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginCreateUpdateSQLStoredProcedureOptions contains the optional parameters for the SQLResourcesClient.BeginCreateUpdateSQLStoredProcedure
// method.
type SQLResourcesClientBeginCreateUpdateSQLStoredProcedureOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginCreateUpdateSQLTriggerOptions contains the optional parameters for the SQLResourcesClient.BeginCreateUpdateSQLTrigger
// method.
type SQLResourcesClientBeginCreateUpdateSQLTriggerOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginCreateUpdateSQLUserDefinedFunctionOptions contains the optional parameters for the SQLResourcesClient.BeginCreateUpdateSQLUserDefinedFunction
// method.
type SQLResourcesClientBeginCreateUpdateSQLUserDefinedFunctionOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginDeleteSQLContainerOptions contains the optional parameters for the SQLResourcesClient.BeginDeleteSQLContainer
// method.
type SQLResourcesClientBeginDeleteSQLContainerOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginDeleteSQLDatabaseOptions contains the optional parameters for the SQLResourcesClient.BeginDeleteSQLDatabase
// method.
type SQLResourcesClientBeginDeleteSQLDatabaseOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginDeleteSQLStoredProcedureOptions contains the optional parameters for the SQLResourcesClient.BeginDeleteSQLStoredProcedure
// method.
type SQLResourcesClientBeginDeleteSQLStoredProcedureOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginDeleteSQLTriggerOptions contains the optional parameters for the SQLResourcesClient.BeginDeleteSQLTrigger
// method.
type SQLResourcesClientBeginDeleteSQLTriggerOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginDeleteSQLUserDefinedFunctionOptions contains the optional parameters for the SQLResourcesClient.BeginDeleteSQLUserDefinedFunction
// method.
type SQLResourcesClientBeginDeleteSQLUserDefinedFunctionOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginMigrateSQLContainerToAutoscaleOptions contains the optional parameters for the SQLResourcesClient.BeginMigrateSQLContainerToAutoscale
// method.
type SQLResourcesClientBeginMigrateSQLContainerToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginMigrateSQLContainerToManualThroughputOptions contains the optional parameters for the SQLResourcesClient.BeginMigrateSQLContainerToManualThroughput
// method.
type SQLResourcesClientBeginMigrateSQLContainerToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginMigrateSQLDatabaseToAutoscaleOptions contains the optional parameters for the SQLResourcesClient.BeginMigrateSQLDatabaseToAutoscale
// method.
type SQLResourcesClientBeginMigrateSQLDatabaseToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginMigrateSQLDatabaseToManualThroughputOptions contains the optional parameters for the SQLResourcesClient.BeginMigrateSQLDatabaseToManualThroughput
// method.
type SQLResourcesClientBeginMigrateSQLDatabaseToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginUpdateSQLContainerThroughputOptions contains the optional parameters for the SQLResourcesClient.BeginUpdateSQLContainerThroughput
// method.
type SQLResourcesClientBeginUpdateSQLContainerThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientBeginUpdateSQLDatabaseThroughputOptions contains the optional parameters for the SQLResourcesClient.BeginUpdateSQLDatabaseThroughput
// method.
type SQLResourcesClientBeginUpdateSQLDatabaseThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// SQLResourcesClientGetClientEncryptionKeyOptions contains the optional parameters for the SQLResourcesClient.GetClientEncryptionKey
// method.
type SQLResourcesClientGetClientEncryptionKeyOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientGetSQLContainerOptions contains the optional parameters for the SQLResourcesClient.GetSQLContainer method.
type SQLResourcesClientGetSQLContainerOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientGetSQLContainerThroughputOptions contains the optional parameters for the SQLResourcesClient.GetSQLContainerThroughput
// method.
type SQLResourcesClientGetSQLContainerThroughputOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientGetSQLDatabaseOptions contains the optional parameters for the SQLResourcesClient.GetSQLDatabase method.
type SQLResourcesClientGetSQLDatabaseOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientGetSQLDatabaseThroughputOptions contains the optional parameters for the SQLResourcesClient.GetSQLDatabaseThroughput
// method.
type SQLResourcesClientGetSQLDatabaseThroughputOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientGetSQLStoredProcedureOptions contains the optional parameters for the SQLResourcesClient.GetSQLStoredProcedure
// method.
type SQLResourcesClientGetSQLStoredProcedureOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientGetSQLTriggerOptions contains the optional parameters for the SQLResourcesClient.GetSQLTrigger method.
type SQLResourcesClientGetSQLTriggerOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientGetSQLUserDefinedFunctionOptions contains the optional parameters for the SQLResourcesClient.GetSQLUserDefinedFunction
// method.
type SQLResourcesClientGetSQLUserDefinedFunctionOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientListClientEncryptionKeysOptions contains the optional parameters for the SQLResourcesClient.NewListClientEncryptionKeysPager
// method.
type SQLResourcesClientListClientEncryptionKeysOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientListSQLContainersOptions contains the optional parameters for the SQLResourcesClient.NewListSQLContainersPager
// method.
type SQLResourcesClientListSQLContainersOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientListSQLDatabasesOptions contains the optional parameters for the SQLResourcesClient.NewListSQLDatabasesPager
// method.
type SQLResourcesClientListSQLDatabasesOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientListSQLStoredProceduresOptions contains the optional parameters for the SQLResourcesClient.NewListSQLStoredProceduresPager
// method.
type SQLResourcesClientListSQLStoredProceduresOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientListSQLTriggersOptions contains the optional parameters for the SQLResourcesClient.NewListSQLTriggersPager
// method.
type SQLResourcesClientListSQLTriggersOptions struct {
	// placeholder for future optional parameters
}

// SQLResourcesClientListSQLUserDefinedFunctionsOptions contains the optional parameters for the SQLResourcesClient.NewListSQLUserDefinedFunctionsPager
// method.
type SQLResourcesClientListSQLUserDefinedFunctionsOptions struct {
	// placeholder for future optional parameters
}

// TableResourcesClientBeginCreateUpdateTableOptions contains the optional parameters for the TableResourcesClient.BeginCreateUpdateTable
// method.
type TableResourcesClientBeginCreateUpdateTableOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// TableResourcesClientBeginDeleteTableOptions contains the optional parameters for the TableResourcesClient.BeginDeleteTable
// method.
type TableResourcesClientBeginDeleteTableOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// TableResourcesClientBeginMigrateTableToAutoscaleOptions contains the optional parameters for the TableResourcesClient.BeginMigrateTableToAutoscale
// method.
type TableResourcesClientBeginMigrateTableToAutoscaleOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// TableResourcesClientBeginMigrateTableToManualThroughputOptions contains the optional parameters for the TableResourcesClient.BeginMigrateTableToManualThroughput
// method.
type TableResourcesClientBeginMigrateTableToManualThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// TableResourcesClientBeginUpdateTableThroughputOptions contains the optional parameters for the TableResourcesClient.BeginUpdateTableThroughput
// method.
type TableResourcesClientBeginUpdateTableThroughputOptions struct {
// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// TableResourcesClientGetTableOptions contains the optional parameters for the TableResourcesClient.GetTable method.
type TableResourcesClientGetTableOptions struct {
	// placeholder for future optional parameters
}

// TableResourcesClientGetTableThroughputOptions contains the optional parameters for the TableResourcesClient.GetTableThroughput
// method.
type TableResourcesClientGetTableThroughputOptions struct {
	// placeholder for future optional parameters
}

// TableResourcesClientListTablesOptions contains the optional parameters for the TableResourcesClient.NewListTablesPager
// method.
type TableResourcesClientListTablesOptions struct {
	// placeholder for future optional parameters
}

